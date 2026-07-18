// Package exec executes planned pipeline stages over the table model
// (SPECIFICATION §3–§6): plain verbs over rows of raw cell texts, with every
// expression already resolved and compiled by internal/plan and evaluated
// through the go-tsvsheet seam. Executors never mutate their input rows — a
// Program is reusable across concurrent runs — and a reference to a cell a
// ragged row lacks reads an empty cell (§2).
package exec

import (
	"strings"

	tsvsheet "github.com/tsvsheet/go-tsvsheet"

	"github.com/tsvsheet/go-tq/internal/constants"
)

// Rows is the data-row stream a pipeline transforms; rows may be ragged.
type Rows [][]string

// Row is one data row of raw cell texts.
type Row []string

// Width is a stage's input column count — the schema width its expressions
// were compiled against.
type Width int

// Height is a group's row count — the shape a group aggregate is compiled for.
type Height int

// Position is a 0-based resolved column index.
type Position int

// Positions is an ordered list of resolved column indexes.
type Positions []Position

// Selection is a set of resolved column indexes (drop columns).
type Selection map[Position]bool

// Count is a limit/offset row count.
type Count int

// RowNumber is a 1-based row position within a stage's input, as strict-mode
// reports name it.
type RowNumber int

// ColumnLabel names the source of a value in a strict-mode report: the target
// column of a derive/aggregate assignment, or a where predicate's source text.
type ColumnLabel string

// cellText is one cell's raw text.
type cellText string

// Env is the per-run evaluation environment shared by every stage: the
// compute options handed to every expression evaluation, and strict mode.
type Env struct {
	Compute  tsvsheet.ComputeOptions
	IsStrict bool
}

// Compiled is one compiled stage expression with the label strict reports use.
type Compiled struct {
	Expr tsvsheet.Expr
	Name ColumnLabel
}

// Predicate is a compiled where predicate: the raw expression (whose value
// strict mode inspects), and the `(expr)=TRUE` form whose TRUE result proves
// the raw value is the boolean TRUE — not the string "TRUE", which the
// tsvsheet comparison rejects as #VALUE! — so "exactly TRUE" (§6) is decided
// by the engine's own typing, never by text.
type Predicate struct {
	Keep  tsvsheet.Expr
	Raw   Compiled
	Width Width
}

// Assign is one compiled derive assignment: the expression and its resolved
// target — an existing position to replace in place, or the row's current
// width to append (plan guarantees At is one of the two).
type Assign struct {
	Cell Compiled
	At   Position
}

// HeightCompiler compiles a group aggregate for a group of h rows, caching by
// height (the plan layer provides it, splicing whole-column ranges C1:Ch).
type HeightCompiler func(h Height) (tsvsheet.Expr, error)

// Agg is one compiled group aggregate.
type Agg struct {
	CompileAt HeightCompiler
	Name      ColumnLabel
}

// trueText is the canonical text of the boolean TRUE value.
const trueText = "TRUE"

// cellAt reads the cell at a column position; a ragged row reads empty (§2).
func cellAt(row Row, at Position) string {
	if int(at) < len(row) {
		return row[at]
	}
	return ""
}

// padded is row normalized to exactly width cells: short rows pad with empty
// cells, cells beyond the schema have no column and are dropped. The result is
// always a fresh slice.
func padded(row Row, width Width) []string {
	out := make([]string, width)
	copy(out, row)
	return out
}

// rowGrid is the one-row synthetic grid a row-stage expression evaluates
// against: the row's cells laid out as columns A1, B1, ….
func rowGrid(row Row, width Width) tsvsheet.Grid {
	return tsvsheet.Grid{padded(row, width)}
}

// gridOf is the synthetic grid of a whole group, one grid row per table row.
func gridOf(rows Rows, width Width) tsvsheet.Grid {
	g := make(tsvsheet.Grid, len(rows))
	for i, row := range rows {
		g[i] = padded(Row(row), width)
	}
	return g
}

// evalText evaluates a compiled expression against a grid and renders the
// canonical cell text — arrays reduced to their scalar-context value.
func evalText(expr tsvsheet.Expr, g tsvsheet.Grid, env Env) string {
	return tsvsheet.FormatValue(expr.Eval(g, env.Compute))
}

// isErrorText reports whether a computed cell text is a spreadsheet error
// value — exactly the engine's closed set, which round-trips through
// FormatValue by construction.
func isErrorText(s cellText) bool {
	switch tsvsheet.ErrorValue(s) {
	case tsvsheet.ErrRef, tsvsheet.ErrValue, tsvsheet.ErrName, tsvsheet.ErrDiv, tsvsheet.ErrCirc,
		tsvsheet.ErrNA, tsvsheet.ErrNum, tsvsheet.ErrNull, tsvsheet.ErrSpill, tsvsheet.ErrImport:
		return true
	default:
		return false
	}
}

// strictReport is the ErrStrict failure for one error value: the value, the
// column (or predicate) that produced it, and the row's 1-based position in
// the failing stage's input.
func strictReport(value cellText, column ColumnLabel, row RowNumber) error {
	return constants.ErrStrict.With(nil, "value", value, "column", column, "row", row)
}

// Project keeps exactly the listed columns, in order; duplicates repeat.
func Project(rows Rows, at Positions) Rows {
	out := make(Rows, len(rows))
	for i, row := range rows {
		cells := make([]string, len(at))
		for j, p := range at {
			cells[j] = cellAt(Row(row), p)
		}
		out[i] = cells
	}
	return out
}

// DropCols removes the selected columns; every other cell keeps its order.
func DropCols(rows Rows, dropped Selection) Rows {
	out := make(Rows, len(rows))
	for i, row := range rows {
		out[i] = dropRow(Row(row), dropped)
	}
	return out
}

// dropRow filters one row's cells by position.
func dropRow(row Row, dropped Selection) []string {
	cells := make([]string, 0, len(row))
	for i, c := range row {
		if !dropped[Position(i)] {
			cells = append(cells, c)
		}
	}
	return cells
}

// Filter keeps rows whose predicate value is exactly TRUE (§6). Any other
// value drops the row; under strict mode an error value aborts instead.
func Filter(rows Rows, p Predicate, env Env) (Rows, error) {
	out := make(Rows, 0, len(rows))
	for i, row := range rows {
		keep, err := keepRow(Row(row), p, env, RowNumber(i+1))
		if err != nil {
			return nil, err
		}
		if keep {
			out = append(out, row)
		}
	}
	return out, nil
}

// keepRow decides one row: the raw predicate value feeds strict mode, and the
// `(expr)=TRUE` form confirms the boolean — a string "TRUE" fails it.
func keepRow(row Row, p Predicate, env Env, at RowNumber) (bool, error) {
	g := rowGrid(row, p.Width)
	text := evalText(p.Raw.Expr, g, env)
	if env.IsStrict && isErrorText(cellText(text)) {
		return false, strictReport(cellText(text), p.Raw.Name, at)
	}
	if text != trueText {
		return false, nil
	}
	return evalText(p.Keep, g, env) == trueText, nil
}

// DeriveRows applies the assignments to every row, left to right, appending or
// replacing in place (§3); each row is normalized to the stage width first so
// an appended column lands at a consistent position.
func DeriveRows(rows Rows, width Width, assigns []Assign, env Env) (Rows, error) {
	out := make(Rows, 0, len(rows))
	for i, row := range rows {
		derived, err := deriveRow(Row(row), width, assigns, env, RowNumber(i+1))
		if err != nil {
			return nil, err
		}
		out = append(out, derived)
	}
	return out, nil
}

// deriveRow computes one row's assignments in order; later ones see earlier
// results because the row (and its evaluation grid) grows as it goes.
func deriveRow(row Row, width Width, assigns []Assign, env Env, at RowNumber) ([]string, error) {
	cur := padded(row, width)
	for _, a := range assigns {
		text := evalText(a.Cell.Expr, tsvsheet.Grid{cur}, env)
		if env.IsStrict && isErrorText(cellText(text)) {
			return nil, strictReport(cellText(text), a.Cell.Name, at)
		}
		cur = place(cur, a.At, cellText(text))
	}
	return cur, nil
}

// place writes text at an existing position, or appends it as the next column.
func place(cells Row, at Position, text cellText) []string {
	if int(at) < len(cells) {
		cells[at] = string(text)
		return cells
	}
	return append(cells, string(text))
}

// SortRows is the stable multi-key sort in the §5 total order. The input order
// is preserved for equal keys; the input slice is never mutated.
func SortRows(rows Rows, keys Keys) Rows {
	out := append(Rows{}, rows...)
	stableSort(out, keys)
	return out
}

// LimitRows keeps the first n rows; a bound beyond the row count keeps all.
func LimitRows(rows Rows, n Count) Rows {
	if int(n) >= len(rows) {
		return rows
	}
	return rows[:n]
}

// OffsetRows skips the first n rows; a bound beyond the row count skips all.
func OffsetRows(rows Rows, n Count) Rows {
	if int(n) >= len(rows) {
		return Rows{}
	}
	return rows[n:]
}

// DistinctRows keeps the first row per key, comparing raw cell text (§3); the
// plan expands whole-row distinct to every schema column.
func DistinctRows(rows Rows, at Positions) Rows {
	seen := map[string]bool{}
	out := make(Rows, 0, len(rows))
	for _, row := range rows {
		k := keyOf(Row(row), at)
		if seen[k] {
			continue
		}
		seen[k] = true
		out = append(out, row)
	}
	return out
}

// keyOf is a row's raw-text key over the given columns. Cells cannot contain
// TAB (the grid reader splits on it), so the join is collision-free.
func keyOf(row Row, at Positions) string {
	parts := make([]string, len(at))
	for i, p := range at {
		parts[i] = cellAt(row, p)
	}
	return strings.Join(parts, "\t")
}

// group is one partition: its rows and its first row's 1-based position.
type group struct {
	rows  Rows
	first RowNumber
}

// GroupRows partitions by key equality on raw cell text and emits one row per
// group in first-appearance order: key columns first, then one column per
// aggregate, each evaluated against the whole group's grid (§3, §5).
func GroupRows(rows Rows, width Width, keys Positions, aggs []Agg, env Env) (Rows, error) {
	groups := partition(rows, keys)
	out := make(Rows, 0, len(groups))
	for _, g := range groups {
		row, err := groupRow(g, width, keys, aggs, env)
		if err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, nil
}

// partition splits rows into groups by key, in first-appearance order.
func partition(rows Rows, keys Positions) []group {
	index := map[string]int{}
	groups := make([]group, 0)
	for i, row := range rows {
		k := keyOf(Row(row), keys)
		gi, ok := index[k]
		if !ok {
			gi = len(groups)
			index[k] = gi
			groups = append(groups, group{first: RowNumber(i + 1)})
		}
		groups[gi].rows = append(groups[gi].rows, row)
	}
	return groups
}

// groupRow emits one group's output row: the key cells from its first row,
// then each aggregate compiled for the group's height and evaluated against
// the group grid.
func groupRow(g group, width Width, keys Positions, aggs []Agg, env Env) ([]string, error) {
	grid := gridOf(g.rows, width)
	out := make([]string, 0, len(keys)+len(aggs))
	for _, at := range keys {
		out = append(out, cellAt(Row(g.rows[0]), at))
	}
	for _, a := range aggs {
		text, err := aggText(a, grid, env, g.first)
		if err != nil {
			return nil, err
		}
		out = append(out, text)
	}
	return out, nil
}

// aggText computes one aggregate cell, aborting under strict mode when the
// value is an error value (reported at the group's first input row).
func aggText(a Agg, grid tsvsheet.Grid, env Env, first RowNumber) (string, error) {
	expr, err := a.CompileAt(Height(len(grid)))
	if err != nil {
		return "", err
	}
	text := evalText(expr, grid, env)
	if env.IsStrict && isErrorText(cellText(text)) {
		return "", strictReport(cellText(text), a.Name, first)
	}
	return text, nil
}
