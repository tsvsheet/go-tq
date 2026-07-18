package ast

import (
	"strconv"
	"strings"

	"github.com/tsvsheet/go-tq/internal/constants"
	grammar "github.com/tsvsheet/go-tq/src/grammar/tq"
)

// tokenText is one token's raw source text.
type tokenText string

// buildProgram converts the program parse tree into the typed Program.
func buildProgram(src ProgramText, tree grammar.IProgramContext) (Program, error) {
	all := tree.AllStage()
	stages := make([]Stage, 0, len(all))
	for _, sc := range all {
		stage, err := buildStage(src, sc)
		if err != nil {
			return Program{}, err
		}
		stages = append(stages, stage)
	}
	return Program{Stages: stages}, nil
}

// buildStage converts one stage parse node; each verb has its own builder.
func buildStage(src ProgramText, ctx grammar.IStageContext) (Stage, error) {
	switch c := ctx.(type) {
	case *grammar.SelectStageContext:
		return columnsStage(VerbSelect, c.ColumnList())
	case *grammar.DropStageContext:
		return columnsStage(VerbDrop, c.ColumnList())
	case *grammar.WhereStageContext:
		return Stage{Verb: VerbWhere, Expr: exprOf(src, c.TqExpr())}, nil
	case *grammar.DeriveStageContext:
		return Stage{Verb: VerbDerive, Assigns: buildAssigns(src, c.AllAssignment())}, nil
	case *grammar.RenameStageContext:
		return buildRename(c)
	case *grammar.SortStageContext:
		return buildSort(c)
	case *grammar.DistinctStageContext:
		return buildDistinct(c)
	case *grammar.LimitStageContext:
		return countStage(VerbLimit, tokenText(c.NUMBER().GetText()))
	case *grammar.OffsetStageContext:
		return countStage(VerbOffset, tokenText(c.NUMBER().GetText()))
	default:
		return buildGroup(src, ctx.(*grammar.GroupStageContext))
	}
}

// columnsStage builds a stage that is just a verb and a column list.
func columnsStage(verb Verb, list grammar.IColumnListContext) (Stage, error) {
	cols, err := buildColumns(list)
	if err != nil {
		return Stage{}, err
	}
	return Stage{Verb: verb, Columns: cols}, nil
}

// buildColumns converts a column list, one item at a time.
func buildColumns(list grammar.IColumnListContext) ([]Column, error) {
	all := list.AllColumnItem()
	cols := make([]Column, 0, len(all))
	for _, item := range all {
		col, err := buildColumnItem(item)
		if err != nil {
			return nil, err
		}
		cols = append(cols, col)
	}
	return cols, nil
}

// buildColumnItem converts one stage-position column: a bracketed reference, a
// bare name, or a bare 1-based index. The grammar admits any NUMBER here; a
// fractional one is a syntax error at program build (SPECIFICATION §3).
func buildColumnItem(ctx grammar.IColumnItemContext) (Column, error) {
	c := ctx.(*grammar.ColumnItemContext)
	switch {
	case c.COLUMN() != nil:
		text := c.COLUMN().GetText()
		return ColumnOf(ColumnContent(text[1 : len(text)-1])), nil
	case c.BareName() != nil:
		return Column{Name: c.BareName().GetText()}, nil
	default:
		return numberColumn(tokenText(c.NUMBER().GetText()))
	}
}

// numberColumn converts a bare NUMBER stage-position column into a 1-based
// index, rejecting a fractional literal.
func numberColumn(text tokenText) (Column, error) {
	if strings.Contains(string(text), ".") {
		return Column{}, constants.ErrSyntax.With(nil, "message", "fractional column index", "number", text)
	}
	return ColumnOf(ColumnContent(text)), nil
}

// countStage builds a limit/offset stage, rejecting a fractional count — the
// grammar admits any NUMBER; integrality is enforced at program build
// (SPECIFICATION §3). An integer too large for int is likewise rejected.
func countStage(verb Verb, text tokenText) (Stage, error) {
	if strings.Contains(string(text), ".") {
		return Stage{}, constants.ErrSyntax.With(nil, "message", "fractional count", "verb", verb, "number", text)
	}
	n, err := strconv.Atoi(string(text))
	if err != nil {
		return Stage{}, constants.ErrSyntax.With(err, "message", "count out of range", "verb", verb, "number", text)
	}
	return Stage{Verb: verb, Count: n}, nil
}

// buildAssigns converts `name = expr` assignments (derive, group aggregates).
func buildAssigns(src ProgramText, all []grammar.IAssignmentContext) []Assignment {
	assigns := make([]Assignment, 0, len(all))
	for _, ac := range all {
		c := ac.(*grammar.AssignmentContext)
		assigns = append(assigns, Assignment{
			Name: columnNameText(c.ColumnName()),
			Expr: exprOf(src, c.TqExpr()),
		})
	}
	return assigns
}

// columnNameText reads an introduced name (derive target, rename target, group
// aggregate): bracket content verbatim, or the bare identifier. Digits-only
// content introduces a digits-only NAME here, never an index (SPECIFICATION §4:
// such a header stays reachable by position).
func columnNameText(ctx grammar.IColumnNameContext) string {
	c := ctx.(*grammar.ColumnNameContext)
	if c.COLUMN() != nil {
		text := c.COLUMN().GetText()
		return text[1 : len(text)-1]
	}
	return c.BareName().GetText()
}

// buildRename converts a rename stage's `col as name` pairs.
func buildRename(c *grammar.RenameStageContext) (Stage, error) {
	all := c.AllRenamePair()
	pairs := make([]RenamePair, 0, len(all))
	for _, pc := range all {
		p := pc.(*grammar.RenamePairContext)
		from, err := buildColumnItem(p.ColumnItem())
		if err != nil {
			return Stage{}, err
		}
		pairs = append(pairs, RenamePair{From: from, To: columnNameText(p.ColumnName())})
	}
	return Stage{Verb: VerbRename, Renames: pairs}, nil
}

// buildSort converts a sort stage's `-?col` keys.
func buildSort(c *grammar.SortStageContext) (Stage, error) {
	all := c.AllSortKey()
	keys := make([]SortKey, 0, len(all))
	for _, kc := range all {
		k := kc.(*grammar.SortKeyContext)
		col, err := buildColumnItem(k.ColumnItem())
		if err != nil {
			return Stage{}, err
		}
		keys = append(keys, SortKey{Col: col, IsDescending: k.DASH() != nil})
	}
	return Stage{Verb: VerbSort, Keys: keys}, nil
}

// buildDistinct converts a distinct stage; no column list means whole-row.
func buildDistinct(c *grammar.DistinctStageContext) (Stage, error) {
	if c.ColumnList() == nil {
		return Stage{Verb: VerbDistinct}, nil
	}
	return columnsStage(VerbDistinct, c.ColumnList())
}

// buildGroup converts a group stage: key columns and aggregate assignments.
func buildGroup(src ProgramText, c *grammar.GroupStageContext) (Stage, error) {
	keys, err := buildColumns(c.ColumnList())
	if err != nil {
		return Stage{}, err
	}
	return Stage{Verb: VerbGroup, Columns: keys, Assigns: buildAssigns(src, c.AllAssignment())}, nil
}

// exprOf slices an expression's ORIGINAL source text out of the program by its
// start/stop token rune indexes — verbatim, whitespace included — so the plan
// layer can splice column references by token span without re-emission.
func exprOf(src ProgramText, ctx grammar.ITqExprContext) Expr {
	runes := []rune(string(src))
	return Expr{Src: ExprText(runes[ctx.GetStart().GetStart() : ctx.GetStop().GetStop()+1])}
}
