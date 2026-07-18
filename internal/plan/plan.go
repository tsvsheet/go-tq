// Package plan resolves and compiles a parsed program against a concrete
// table shape — the phase between parse and execution. Every column reference
// resolves against the header (or arity, headerless) before any row is
// processed (ErrUnknownColumn), raw A1/sheet-qualified references are rejected
// (ErrCellRef), name-requiring forms fail headerless (ErrHeaderless), and each
// stage expression is spliced to A1 form and compiled ONCE through the
// go-tsvsheet seam. The result is an executable step list over internal/exec.
package plan

import (
	tsvsheet "github.com/tsvsheet/go-tsvsheet"

	"github.com/tsvsheet/go-tq/internal/ast"
	"github.com/tsvsheet/go-tq/internal/constants"
	"github.com/tsvsheet/go-tq/internal/exec"
)

// Shape is the input table's planning shape: its header (when present) and
// column count. Headerless width is the table's arity — its widest row.
type Shape struct {
	Header    []string
	Width     int
	HasHeader bool
}

// Options configures a plan: strict mode, the compute options every
// expression evaluation receives, and the expression compiler (nil selects
// the go-tsvsheet-backed Compile; tests inject failing compilers to exercise
// seam failure paths).
type Options struct {
	Compile  CompileFunc
	Compute  tsvsheet.ComputeOptions
	IsStrict bool
}

// step transforms the row stream; a step that can't fail wraps via pure.
type step func(rows exec.Rows) (exec.Rows, error)

// transform is a step that cannot fail.
type transform func(rows exec.Rows) exec.Rows

// pure lifts an infallible transform into a step.
func pure(f transform) step {
	return func(rows exec.Rows) (exec.Rows, error) { return f(rows), nil }
}

// Plan is a program resolved and compiled against one table shape: the
// ordered executable steps and the output header.
type Plan struct {
	steps  []step
	header []string
}

// New plans prog against shape. Every plan-time failure — ErrUnknownColumn,
// ErrCellRef, ErrHeaderless, and a builder expression's ErrSyntax — surfaces
// here, before any row is processed.
func New(prog ast.Program, shape Shape, opts Options) (Plan, error) {
	p := planner{env: exec.Env{Compute: opts.Compute, IsStrict: opts.IsStrict}, compile: opts.Compile}
	if p.compile == nil {
		p.compile = Compile
	}
	s := schema{names: shape.Header, width: columnCount(shape.Width), hasHeader: shape.HasHeader}
	steps := make([]step, 0, len(prog.Stages))
	for _, st := range prog.Stages {
		stp, next, err := p.stage(st, s)
		if err != nil {
			return Plan{}, err
		}
		if stp != nil {
			steps = append(steps, stp)
		}
		s = next
	}
	return Plan{steps: steps, header: s.headerRow()}, nil
}

// Run executes every step over rows, in order.
func (p Plan) Run(rows exec.Rows) (exec.Rows, error) {
	var err error
	for _, stp := range p.steps {
		if rows, err = stp(rows); err != nil {
			return nil, err
		}
	}
	return rows, nil
}

// Header is the output header row — nil in headerless mode, where no header
// is emitted.
func (p Plan) Header() []string { return p.header }

// planner carries the per-run environment through stage planning.
type planner struct {
	compile CompileFunc
	env     exec.Env
}

// stage plans one stage: its executable step (nil for the header-only
// rename) and the schema it hands the next stage.
func (p planner) stage(st ast.Stage, s schema) (step, schema, error) {
	switch st.Verb {
	case ast.VerbSelect:
		return p.selectStage(st, s)
	case ast.VerbDrop:
		return p.dropStage(st, s)
	case ast.VerbWhere:
		return p.whereStage(st, s)
	case ast.VerbDerive:
		return p.deriveStage(st, s)
	case ast.VerbRename:
		return p.renameStage(st, s)
	case ast.VerbSort:
		return p.sortStage(st, s)
	case ast.VerbDistinct:
		return p.distinctStage(st, s)
	case ast.VerbLimit, ast.VerbOffset:
		return p.countStage(st, s)
	default:
		return p.groupStage(st, s)
	}
}

// selectStage projects and reorders to exactly the listed columns.
func (p planner) selectStage(st ast.Stage, s schema) (step, schema, error) {
	at, err := s.resolveAll(st.Columns)
	if err != nil {
		return nil, schema{}, err
	}
	return pure(func(rows exec.Rows) exec.Rows { return exec.Project(rows, at) }), s.project(at), nil
}

// dropStage removes the listed columns; everything else keeps its order.
func (p planner) dropStage(st ast.Stage, s schema) (step, schema, error) {
	at, err := s.resolveAll(st.Columns)
	if err != nil {
		return nil, schema{}, err
	}
	set := exec.Selection{}
	for _, a := range at {
		set[a] = true
	}
	return pure(func(rows exec.Rows) exec.Rows { return exec.DropCols(rows, set) }), s.drop(set), nil
}

// whereStage compiles the predicate twice through the seam: the raw form
// (whose value strict mode inspects) and the `(expr)=TRUE` form that proves
// the value is the boolean TRUE.
func (p planner) whereStage(st ast.Stage, s schema) (step, schema, error) {
	ep, err := p.exprPlan(st.Expr, s)
	if err != nil {
		return nil, schema{}, err
	}
	spliced := ep.splice(rowRef)
	keep, err := p.compile("(" + spliced + ")=TRUE")
	if err != nil {
		return nil, schema{}, err
	}
	raw, err := p.compile(spliced)
	if err != nil {
		return nil, schema{}, err
	}
	pred := exec.Predicate{
		Raw:   exec.Compiled{Expr: raw, Name: exec.ColumnLabel(st.Expr.Src)},
		Keep:  keep,
		Width: exec.Width(s.width),
	}
	return func(rows exec.Rows) (exec.Rows, error) { return exec.Filter(rows, pred, p.env) }, s, nil
}

// deriveStage compiles each assignment against the schema as it grows, so
// later assignments see earlier results (§3).
func (p planner) deriveStage(st ast.Stage, s schema) (step, schema, error) {
	width := exec.Width(s.width)
	assigns := make([]exec.Assign, 0, len(st.Assigns))
	cur := s
	for _, a := range st.Assigns {
		asg, next, err := p.assign(a, cur)
		if err != nil {
			return nil, schema{}, err
		}
		assigns = append(assigns, asg)
		cur = next
	}
	return func(rows exec.Rows) (exec.Rows, error) {
		return exec.DeriveRows(rows, width, assigns, p.env)
	}, cur, nil
}

// assign compiles one derive assignment and resolves its target: replace in
// place when the name already exists, append otherwise.
func (p planner) assign(a ast.Assignment, s schema) (exec.Assign, schema, error) {
	ep, err := p.exprPlan(a.Expr, s)
	if err != nil {
		return exec.Assign{}, schema{}, err
	}
	expr, err := p.compile(ep.splice(rowRef))
	if err != nil {
		return exec.Assign{}, schema{}, err
	}
	at, next := s.target(headerName(a.Name))
	return exec.Assign{Cell: exec.Compiled{Expr: expr, Name: exec.ColumnLabel(a.Name)}, At: at}, next, nil
}

// renameStage relabels without moving — a header-only change with no
// executable step. Headerless mode has no names to change (ErrHeaderless).
func (p planner) renameStage(st ast.Stage, s schema) (step, schema, error) {
	if !s.hasHeader {
		return nil, schema{}, constants.ErrHeaderless.With(nil, "verb", ast.VerbRename)
	}
	cur := s
	for _, r := range st.Renames {
		at, err := cur.resolve(r.From)
		if err != nil {
			return nil, schema{}, err
		}
		cur = cur.rename(at, headerName(r.To))
	}
	return nil, cur, nil
}

// sortStage resolves the multi-key list for the stable §5 total-order sort.
func (p planner) sortStage(st ast.Stage, s schema) (step, schema, error) {
	keys := make(exec.Keys, 0, len(st.Keys))
	for _, k := range st.Keys {
		at, err := s.resolve(k.Col)
		if err != nil {
			return nil, schema{}, err
		}
		keys = append(keys, exec.Key{At: at, IsDescending: k.IsDescending})
	}
	return pure(func(rows exec.Rows) exec.Rows { return exec.SortRows(rows, keys) }), s, nil
}

// distinctStage resolves the key columns; no columns means whole-row, which
// expands to every schema column.
func (p planner) distinctStage(st ast.Stage, s schema) (step, schema, error) {
	at, err := s.resolveAll(st.Columns)
	if err != nil {
		return nil, schema{}, err
	}
	if len(st.Columns) == 0 {
		at = allPositions(s.width)
	}
	return pure(func(rows exec.Rows) exec.Rows { return exec.DistinctRows(rows, at) }), s, nil
}

// countStage plans limit/offset. The grammar only admits non-negative
// literals; a negative builder count is rejected at program build (§3).
func (p planner) countStage(st ast.Stage, s schema) (step, schema, error) {
	if st.Count < 0 {
		return nil, schema{}, constants.ErrSyntax.With(
			nil,
			"message",
			"negative count",
			"verb",
			st.Verb,
			"number",
			st.Count,
		)
	}
	n := exec.Count(st.Count)
	if st.Verb == ast.VerbLimit {
		return pure(func(rows exec.Rows) exec.Rows { return exec.LimitRows(rows, n) }), s, nil
	}
	return pure(func(rows exec.Rows) exec.Rows { return exec.OffsetRows(rows, n) }), s, nil
}

// groupStage resolves the key columns and compiles each aggregate with a
// per-height compiler (validated eagerly at height 1, so a seam failure
// surfaces at plan time, and cached — one CompileExpr per unique height).
func (p planner) groupStage(st ast.Stage, s schema) (step, schema, error) {
	keys, err := s.resolveAll(st.Columns)
	if err != nil {
		return nil, schema{}, err
	}
	width := exec.Width(s.width)
	aggs := make([]exec.Agg, 0, len(st.Assigns))
	for _, a := range st.Assigns {
		agg, err := p.agg(a, s)
		if err != nil {
			return nil, schema{}, err
		}
		aggs = append(aggs, agg)
	}
	return func(rows exec.Rows) (exec.Rows, error) {
		return exec.GroupRows(rows, width, keys, aggs, p.env)
	}, s.grouped(keys, st.Assigns), nil
}

// agg plans one aggregate assignment: column references denote whole-column
// ranges across the group (§5), compiled per group height.
func (p planner) agg(a ast.Assignment, s schema) (exec.Agg, error) {
	ep, err := p.exprPlan(a.Expr, s)
	if err != nil {
		return exec.Agg{}, err
	}
	comp := p.heightCompiler(ep)
	if _, err := comp(1); err != nil {
		return exec.Agg{}, err
	}
	return exec.Agg{Name: exec.ColumnLabel(a.Name), CompileAt: comp}, nil
}

// heightCompiler builds the caching per-height compiler for one aggregate.
func (p planner) heightCompiler(ep exprPlan) exec.HeightCompiler {
	cache := map[exec.Height]tsvsheet.Expr{}
	return func(h exec.Height) (tsvsheet.Expr, error) {
		if expr, ok := cache[h]; ok {
			return expr, nil
		}
		expr, err := p.compile(ep.splice(rangeRefAt(h)))
		if err != nil {
			return tsvsheet.Expr{}, err
		}
		cache[h] = expr
		return expr, nil
	}
}

// allPositions is 0..width-1 — the whole-row key set.
func allPositions(width columnCount) exec.Positions {
	at := make(exec.Positions, width)
	for i := range at {
		at[i] = exec.Position(i)
	}
	return at
}
