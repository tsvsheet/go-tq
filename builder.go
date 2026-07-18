package tq

import (
	"strings"

	"github.com/tsvsheet/go-tq/internal/ast"
)

// Column is a stage-position column reference, written as it would appear in
// a program: a bare header name (`stars`), a bare 1-based index (`2`), or the
// bracketed form (`[full title]`, `[2]`) — brackets are stripped, and
// digits-only content is an index (SPECIFICATION §4). Unlike the parser, the
// builder accepts any name unbracketed.
type Column string

// ColumnName is a column name being introduced — a derive target, a rename
// target, or a group aggregate name — taken verbatim (digits-only text
// introduces a digits-only name, never an index).
type ColumnName string

// Expression is one tq expression in tsvsheet formula syntax (the pipe sugar
// excluded — `|` belongs to stages), exactly as it would appear in a program.
// A builder expression is validated at Run's plan step (ErrSyntax).
type Expression string

// RowCount is a limit/offset row count; a negative count is rejected at
// Run's plan step (ErrSyntax), matching the grammar's non-negative literals.
type RowCount int

// Assignment is one `name = expr` (derive or group aggregate).
type Assignment struct {
	Name ColumnName
	Expr Expression
}

// RenamePair is one `col as name`.
type RenamePair struct {
	Col Column
	As  ColumnName
}

// SortKey is one sort key; IsDescending reverses the key's entire order
// (`-col`).
type SortKey struct {
	Col          Column
	IsDescending bool
}

// Stage is one built pipeline stage, produced by the per-verb constructors.
type Stage struct {
	stage ast.Stage
}

// NewProgram builds a Program from stages — identical in value and behavior
// to parsing the equivalent query text.
func NewProgram(stages ...Stage) Program {
	built := make([]ast.Stage, 0, len(stages))
	for _, s := range stages {
		built = append(built, s.stage)
	}
	return Program{prog: ast.Program{Stages: built}}
}

// Select projects and reorders to exactly the listed columns (duplicates
// allowed).
func Select(cols ...Column) Stage {
	return Stage{stage: ast.Stage{Verb: ast.VerbSelect, Columns: astColumns(cols)}}
}

// Drop removes the listed columns; everything else keeps its order.
func Drop(cols ...Column) Stage {
	return Stage{stage: ast.Stage{Verb: ast.VerbDrop, Columns: astColumns(cols)}}
}

// Where keeps rows whose predicate value is exactly TRUE.
func Where(expr Expression) Stage {
	return Stage{stage: ast.Stage{Verb: ast.VerbWhere, Expr: astExpr(expr)}}
}

// Derive appends computed columns, or replaces in place on name collision,
// left to right.
func Derive(assignments ...Assignment) Stage {
	return Stage{stage: ast.Stage{Verb: ast.VerbDerive, Assigns: astAssigns(assignments)}}
}

// Rename relabels columns without moving them.
func Rename(pairs ...RenamePair) Stage {
	renames := make([]ast.RenamePair, 0, len(pairs))
	for _, p := range pairs {
		renames = append(renames, ast.RenamePair{From: astColumn(p.Col), To: string(p.As)})
	}
	return Stage{stage: ast.Stage{Verb: ast.VerbRename, Renames: renames}}
}

// SortBy sorts stably by the keys, in the total order of SPECIFICATION §5.
func SortBy(keys ...SortKey) Stage {
	built := make([]ast.SortKey, 0, len(keys))
	for _, k := range keys {
		built = append(built, ast.SortKey{Col: astColumn(k.Col), IsDescending: k.IsDescending})
	}
	return Stage{stage: ast.Stage{Verb: ast.VerbSort, Keys: built}}
}

// Distinct keeps the first row per key — whole row when no columns are given
// — comparing raw cell text.
func Distinct(cols ...Column) Stage {
	return Stage{stage: ast.Stage{Verb: ast.VerbDistinct, Columns: astColumns(cols)}}
}

// Limit keeps the first n rows.
func Limit(n RowCount) Stage {
	return Stage{stage: ast.Stage{Verb: ast.VerbLimit, Count: int(n)}}
}

// Offset skips the first n rows.
func Offset(n RowCount) Stage {
	return Stage{stage: ast.Stage{Verb: ast.VerbOffset, Count: int(n)}}
}

// GroupBy partitions by the key columns and emits one row per group — key
// columns first, then one column per aggregate — in first-appearance order.
func GroupBy(keys []Column, aggs ...Assignment) Stage {
	return Stage{stage: ast.Stage{Verb: ast.VerbGroup, Columns: astColumns(keys), Assigns: astAssigns(aggs)}}
}

// astColumns converts a builder column list; an empty list stays nil so a
// built stage is value-identical to its parsed equivalent (`distinct`).
func astColumns(cols []Column) []ast.Column {
	if len(cols) == 0 {
		return nil
	}
	built := make([]ast.Column, 0, len(cols))
	for _, c := range cols {
		built = append(built, astColumn(c))
	}
	return built
}

// astColumn classifies one builder column: brackets are stripped, then
// digits-only content is a 1-based index, anything else a name.
func astColumn(c Column) ast.Column {
	text := string(c)
	if len(text) >= 2 && strings.HasPrefix(text, "[") && strings.HasSuffix(text, "]") {
		text = text[1 : len(text)-1]
	}
	return ast.ColumnOf(ast.ColumnContent(text))
}

// astExpr wraps a builder expression as its verbatim source text.
func astExpr(expr Expression) ast.Expr {
	return ast.Expr{Src: ast.ExprText(expr)}
}

// astAssigns converts builder assignments; an empty list stays nil.
func astAssigns(assignments []Assignment) []ast.Assignment {
	if len(assignments) == 0 {
		return nil
	}
	built := make([]ast.Assignment, 0, len(assignments))
	for _, a := range assignments {
		built = append(built, ast.Assignment{Name: string(a.Name), Expr: astExpr(a.Expr)})
	}
	return built
}
