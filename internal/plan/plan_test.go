package plan_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tsvsheet "github.com/tsvsheet/go-tsvsheet"

	"github.com/tsvsheet/go-tq/internal/ast"
	"github.com/tsvsheet/go-tq/internal/constants"
	"github.com/tsvsheet/go-tq/internal/exec"
	"github.com/tsvsheet/go-tq/internal/plan"
)

// TestCompileWrapsBothSentinels asserts a seam compile failure is tq's
// ErrSyntax AND still reaches go-tsvsheet's ErrSyntax through With, so
// errors.Is distinguishes the engines (tq-api contract).
func TestCompileWrapsBothSentinels(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	_, err := plan.Compile("(((")
	require.Error(t, err)
	want.True(errors.Is(err, constants.ErrSyntax), "tq sentinel: %v", err)
	want.True(errors.Is(err, tsvsheet.ErrSyntax), "tsvsheet sentinel: %v", err)

	expr, err := plan.Compile("1+1")
	require.NoError(t, err)
	want.Equal("2", tsvsheet.FormatValue(expr.Eval(nil, tsvsheet.ComputeOptions{})))
}

// failFrom is a CompileFunc that behaves like the real compiler until the
// n-th call, which fails — injected to exercise every seam-failure path.
func failFrom(n int) plan.CompileFunc {
	calls := 0
	return func(src plan.SplicedText) (tsvsheet.Expr, error) {
		calls++
		if calls >= n {
			return tsvsheet.Expr{}, assert.AnError
		}
		return plan.Compile(src)
	}
}

// program is a one-stage helper.
func program(stage ast.Stage) ast.Program {
	return ast.Program{Stages: []ast.Stage{stage}}
}

// headed is a header-mode shape over the given names.
func headed(names ...string) plan.Shape {
	return plan.Shape{Header: names, Width: len(names), HasHeader: true}
}

// TestCompilerFailurePaths asserts every plan-time compile call site
// propagates a seam failure: the where keep and raw forms, a derive
// assignment, and a group aggregate's eager height-1 compile.
func TestCompilerFailurePaths(t *testing.T) {
	t.Parallel()
	where := program(ast.Stage{Verb: ast.VerbWhere, Expr: ast.Expr{Src: "[a] > 1"}})
	derive := program(
		ast.Stage{Verb: ast.VerbDerive, Assigns: []ast.Assignment{{Name: "x", Expr: ast.Expr{Src: "[a]"}}}},
	)
	grouped := program(ast.Stage{
		Verb:    ast.VerbGroup,
		Columns: []ast.Column{{Name: "a"}},
		Assigns: []ast.Assignment{{Name: "s", Expr: ast.Expr{Src: "sum([a])"}}},
	})
	cases := map[string]struct {
		prog ast.Program
		from int
	}{
		"where keep form":       {prog: where, from: 1},
		"where raw form":        {prog: where, from: 2},
		"derive assignment":     {prog: derive, from: 1},
		"group eager aggregate": {prog: grouped, from: 1},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := plan.New(tc.prog, headed("a"), plan.Options{Compile: failFrom(tc.from)})
			require.Error(t, err)
			assert.ErrorIs(t, err, assert.AnError)
		})
	}
}

// TestGroupHeightCompileFailureAtRun asserts a seam failure for a LATER group
// height (the eager height-1 compile succeeded) surfaces from Run.
func TestGroupHeightCompileFailureAtRun(t *testing.T) {
	t.Parallel()
	grouped := program(ast.Stage{
		Verb:    ast.VerbGroup,
		Columns: []ast.Column{{Name: "a"}},
		Assigns: []ast.Assignment{{Name: "s", Expr: ast.Expr{Src: "sum([b])"}}},
	})
	p, err := plan.New(grouped, headed("a", "b"), plan.Options{Compile: failFrom(2)})
	require.NoError(t, err)
	_, err = p.Run(exec.Rows{{"k", "1"}, {"k", "2"}})
	require.Error(t, err)
	assert.ErrorIs(t, err, assert.AnError)
}

// TestGroupHeightCompileCaches asserts one CompileExpr per unique group
// height: with the real compiler allowed exactly two calls (eager height 1 +
// height 2), three groups of heights 2, 1, and 2 still run.
func TestGroupHeightCompileCaches(t *testing.T) {
	t.Parallel()
	grouped := program(ast.Stage{
		Verb:    ast.VerbGroup,
		Columns: []ast.Column{{Name: "a"}},
		Assigns: []ast.Assignment{{Name: "s", Expr: ast.Expr{Src: "sum([b])"}}},
	})
	p, err := plan.New(grouped, headed("a", "b"), plan.Options{Compile: failFrom(3)})
	require.NoError(t, err)
	rows, err := p.Run(exec.Rows{{"x", "1"}, {"x", "2"}, {"y", "5"}, {"z", "3"}, {"z", "4"}})
	require.NoError(t, err)
	assert.Equal(t, exec.Rows{{"x", "3"}, {"y", "5"}, {"z", "7"}}, rows)
}

// TestSpliceWideColumns asserts the A1 rendering past column Z (bijective
// base 26): a 28-column table's last references splice to AA1/AB1.
func TestSpliceWideColumns(t *testing.T) {
	t.Parallel()
	derive := program(ast.Stage{Verb: ast.VerbDerive, Assigns: []ast.Assignment{
		{Name: "x", Expr: ast.Expr{Src: "[27] & [28]"}},
	}})
	p, err := plan.New(derive, plan.Shape{Width: 28}, plan.Options{})
	require.NoError(t, err)
	row := make([]string, 28)
	row[26], row[27] = "AA-val", "AB-val"
	rows, err := p.Run(exec.Rows{row})
	require.NoError(t, err)
	require.Len(t, rows, 1)
	assert.Equal(t, "AA-valAB-val", rows[0][28])
}

// TestSpliceNormalizesWhitespaceOutsideStrings asserts the splice contract:
// newlines/tabs BETWEEN tokens (legal in tq, illegal in a tsvsheet formula)
// become spaces, while a TAB INSIDE a string literal passes through verbatim.
func TestSpliceNormalizesWhitespaceOutsideStrings(t *testing.T) {
	t.Parallel()
	where := program(ast.Stage{Verb: ast.VerbWhere, Expr: ast.Expr{Src: "[a]\t<>\n\r\"x\ty\""}})
	p, err := plan.New(where, headed("a"), plan.Options{})
	require.NoError(t, err)
	rows, err := p.Run(exec.Rows{{"q"}, {"x\ty"}})
	require.NoError(t, err)
	assert.Equal(t, exec.Rows{{"q"}}, rows)
}

// TestHeaderTravel asserts Plan.Header reflects the pipeline's final schema
// in header mode and stays nil headerless.
func TestHeaderTravel(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	prog := ast.Program{Stages: []ast.Stage{
		{Verb: ast.VerbRename, Renames: []ast.RenamePair{{From: ast.Column{Name: "a"}, To: "z"}}},
		{Verb: ast.VerbDerive, Assigns: []ast.Assignment{{Name: "d", Expr: ast.Expr{Src: "1"}}}},
	}}
	p, err := plan.New(prog, headed("a", "b"), plan.Options{})
	require.NoError(t, err)
	want.Equal([]string{"z", "b", "d"}, p.Header())

	headless, err := plan.New(program(ast.Stage{Verb: ast.VerbLimit, Count: 1}), plan.Shape{Width: 2}, plan.Options{})
	require.NoError(t, err)
	want.Nil(headless.Header())
}

// TestNegativeCount asserts a builder-only negative limit/offset is rejected
// at plan as ErrSyntax, mirroring the grammar's non-negative literals.
func TestNegativeCount(t *testing.T) {
	t.Parallel()
	for _, verb := range []ast.Verb{ast.VerbLimit, ast.VerbOffset} {
		_, err := plan.New(program(ast.Stage{Verb: verb, Count: -1}), headed("a"), plan.Options{})
		require.Error(t, err)
		assert.ErrorIs(t, err, constants.ErrSyntax)
	}
}
