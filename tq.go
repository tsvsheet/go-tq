package tq

import (
	"io"
	"slices"
	"strings"
	"time"

	tsvsheet "github.com/tsvsheet/go-tsvsheet"

	"github.com/tsvsheet/go-tq/internal/ast"
	"github.com/tsvsheet/go-tq/internal/constants"
	"github.com/tsvsheet/go-tq/internal/exec"
	"github.com/tsvsheet/go-tq/internal/plan"
)

// Query is tq source text — a pipeline of stages separated by `|`.
type Query string

// Program is a parsed, unplanned pipeline of stages: an immutable value,
// reusable and safe for concurrent Runs. Parse and the builder (NewProgram
// with the per-verb stage constructors) produce identical Programs.
type Program struct {
	prog ast.Program
}

// Table is the data model: an optional header row plus data rows of raw cell
// texts. A nil Header means headerless. Rows may be ragged; a reference to a
// column a row lacks reads an empty cell. Constructed by ReadTable, emitted
// by WriteTable.
type Table struct {
	Header []string
	Rows   [][]string
}

// Options is the per-run configuration. The zero value means: header-first,
// lenient, compute-first, unbounded input, default engine limits, and embeds
// and imports disabled (nil Loader/Fetcher).
type Options struct {
	// At is the compute-pass clock volatile functions read.
	At time.Time
	// Fetcher serves IMPORT* functions; nil disables them (#IMPORT!).
	Fetcher tsvsheet.Fetcher
	// Loader resolves SHEET(...) embeds; nil disables them (#REF!).
	Loader tsvsheet.Loader
	// Limits bounds engine allocations during the compute pass and expression
	// evaluation; the zero value selects the engine's defaults.
	Limits tsvsheet.Limits
	// MaxCells is the input ceiling ReadTable admits — rows × columns; 0 is
	// unbounded. Exceeding it is ErrLimit.
	MaxCells int
	// IsHeaderless treats every row as data: column references are positional
	// only, rename is ErrHeaderless, and no header is emitted.
	IsHeaderless bool
	// IsStrict aborts the run with ErrStrict on the first expression
	// evaluation that produces an error value.
	IsStrict bool
	// IsRaw skips the compute-first pass: every cell is verbatim text,
	// formula cells included.
	IsRaw bool
}

// compute is the tsvsheet ComputeOptions pass-through the compute-first pass
// and every expression evaluation share.
func (o Options) compute() tsvsheet.ComputeOptions {
	return tsvsheet.ComputeOptions{At: o.At, Limits: o.Limits, Loader: o.Loader, Fetcher: o.Fetcher}
}

// Parse parses and validates a whole pipeline; a failure is ErrSyntax with
// line/column detail. There are no partial programs.
func Parse(q Query) (Program, error) {
	prog, err := ast.ParseProgram(ast.ProgramText(q))
	if err != nil {
		return Program{}, err
	}
	return Program{prog: prog}, nil
}

// Run plans the program against t's header (or arity, headerless) and
// executes every stage. Plan-time failures — ErrUnknownColumn, ErrCellRef,
// ErrHeaderless, a builder expression's ErrSyntax — surface before any row is
// processed. When the input holds formula cells and IsRaw is unset, the sheet
// is computed first and the query runs over the computed values (ADR 0003).
// The returned Table shares no storage with t: mutating one never affects the
// other.
func (p Program) Run(t Table, opts Options) (Table, error) {
	tbl, err := computeFirst(tableFor(t, opts), opts)
	if err != nil {
		return Table{}, err
	}
	shape := shapeOf(tbl, opts)
	pl, err := plan.New(p.prog, shape, plan.Options{IsStrict: opts.IsStrict, Compute: opts.compute()})
	if err != nil {
		return Table{}, err
	}
	rows, err := pl.Run(exec.Rows(clipped(tbl.Rows, shape)))
	if err != nil {
		return Table{}, err
	}
	return Table{Header: slices.Clone(pl.Header()), Rows: cloneRows(rows)}, nil
}

// clipped bounds every data row to the header's width (§2): a cell beyond the
// last header column is outside the table model — no reference can reach it —
// so it never enters the pipeline. Headerless mode is untouched: there the
// schema width is the arity, the widest row.
func clipped(rows [][]string, s plan.Shape) [][]string {
	if !s.HasHeader {
		return rows
	}
	out := make([][]string, len(rows))
	for i, row := range rows {
		out[i] = row[:min(len(row), s.Width)]
	}
	return out
}

// cloneRows deep-copies the row structure (cell strings are immutable), so a
// returned Table shares no slice storage with the run's input.
func cloneRows(rows exec.Rows) [][]string {
	out := make([][]string, len(rows))
	for i, row := range rows {
		out[i] = slices.Clone(row)
	}
	return out
}

// tableFor normalizes the run's input per the options: headerless, every row
// is data and any Header value is ignored (ReadTable never sets one).
func tableFor(t Table, opts Options) Table {
	if opts.IsHeaderless {
		return Table{Rows: t.Rows}
	}
	return t
}

// shapeOf is t's planning shape: the header and its width, or — headerless —
// the table's arity, its widest row.
func shapeOf(t Table, opts Options) plan.Shape {
	if opts.IsHeaderless {
		return plan.Shape{Width: int(arity(t.Rows))}
	}
	return plan.Shape{Header: t.Header, Width: len(t.Header), HasHeader: true}
}

// tableWidth is a table's column count.
type tableWidth int

// arity is the widest row's cell count.
func arity(rows tsvsheet.Grid) tableWidth {
	width := 0
	for _, row := range rows {
		width = max(width, len(row))
	}
	return tableWidth(width)
}

// computeFirst computes the sheet when the table holds formula cells and Raw
// is unset, so the query sees values, not formulas. A malformed formula cell
// surfaces as go-tsvsheet's own syntax error naming the cell — an input-data
// failure, not a tq program error.
func computeFirst(t Table, opts Options) (Table, error) {
	if opts.IsRaw || !hasFormula(t) {
		return t, nil
	}
	sheet, err := tsvsheet.Parse([]byte(tsvSource(t)))
	if err != nil {
		return Table{}, err
	}
	grid := sheet.ComputeWith(opts.compute())
	return splitGrid(grid, withHeader(t.Header != nil)), nil
}

// hasFormula reports whether any cell — header included — is a formula.
func hasFormula(t Table) bool {
	for _, row := range append([][]string{t.Header}, t.Rows...) {
		for _, cell := range row {
			if strings.HasPrefix(cell, "=") {
				return true
			}
		}
	}
	return false
}

// tsvSource serializes the table as TSV for the compute pass: cells joined by
// TAB, one newline-terminated line per row — byte-identical to WriteTable's
// output (asserted in tests) but total, cells being TAB- and newline-free by
// the grid model.
func tsvSource(t Table) string {
	var b strings.Builder
	for _, row := range gridOf(t) {
		_, _ = b.WriteString(strings.Join(row, "\t"))
		_, _ = b.WriteString("\n")
	}
	return b.String()
}

// withHeader tells splitGrid whether the grid's first row is a header.
type withHeader bool

// splitGrid rebuilds a Table from a grid, splitting off the header row when
// the grid carries one.
func splitGrid(grid tsvsheet.Grid, hasHeader withHeader) Table {
	if !hasHeader || len(grid) == 0 {
		return Table{Rows: grid}
	}
	return Table{Header: grid[0], Rows: grid[1:]}
}

// gridOf lays the table out as one grid, header row first when present.
func gridOf(t Table) tsvsheet.Grid {
	if t.Header == nil {
		return t.Rows
	}
	grid := make(tsvsheet.Grid, 0, len(t.Rows)+1)
	grid = append(grid, t.Header)
	return append(grid, t.Rows...)
}

// ReadTable reads a TSV grid with tsvsheet's comment semantics (a leading
// `#!` first line and any `# ` line are skipped), splits the header (all rows
// are data under IsHeaderless), and enforces opts.MaxCells (ErrLimit). A read failure passes
// through as go-tsvsheet's own sentinel.
func ReadTable(r io.Reader, opts Options) (Table, error) {
	grid, err := tsvsheet.ReadTSV(r)
	if err != nil {
		return Table{}, err
	}
	if err := checkCells(grid, cellCeiling(opts.MaxCells)); err != nil {
		return Table{}, err
	}
	return splitGrid(grid, withHeader(!opts.IsHeaderless)), nil
}

// cellCeiling is the MaxCells input ceiling; 0 or below is unbounded.
type cellCeiling int

// checkCells enforces the MaxCells input ceiling: rows × columns (the
// rectangular bound over the widest row).
func checkCells(grid tsvsheet.Grid, maxCells cellCeiling) error {
	if maxCells <= 0 {
		return nil
	}
	rows, cols := len(grid), int(arity(grid))
	if rows*cols > int(maxCells) {
		return constants.ErrLimit.With(nil, "rows", rows, "columns", cols, "max", maxCells)
	}
	return nil
}

// WriteTable writes the table as TSV — header row first when present, each
// row newline-terminated — via go-tsvsheet's grid writer. A write failure
// passes through as go-tsvsheet's own sentinel.
func WriteTable(w io.Writer, t Table) error {
	return tsvsheet.WriteTSV(w, gridOf(t))
}
