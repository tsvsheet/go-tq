package tq_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tsvsheet "github.com/tsvsheet/go-tsvsheet"

	tq "github.com/tsvsheet/go-tq"
)

// TestSentinelsDistinct asserts the six tq sentinels are pairwise distinct and
// distinct from every go-tsvsheet sentinel, so errors.Is tells the engines
// apart (R8).
func TestSentinelsDistinct(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	tqErrs := []error{tq.ErrSyntax, tq.ErrUnknownColumn, tq.ErrCellRef, tq.ErrHeaderless, tq.ErrStrict, tq.ErrLimit}
	for i, a := range tqErrs {
		for j, b := range tqErrs {
			if i != j {
				want.NotErrorIs(a, b)
			}
		}
	}
	for _, theirs := range []error{
		tsvsheet.ErrSyntax, tsvsheet.ErrInvalidValue, tsvsheet.ErrNotFound, tsvsheet.ErrReadInput, tsvsheet.ErrWriteFile,
	} {
		for _, ours := range tqErrs {
			want.NotErrorIs(ours, theirs)
			want.NotErrorIs(theirs, ours)
		}
	}
}

// parseErr parses a query expected to fail and returns the error.
func parseErr(t *testing.T, query string) error {
	t.Helper()
	_, err := tq.Parse(tq.Query(query))
	require.Error(t, err, "query %q must fail", query)
	return err
}

// TestParseSyntaxErrors asserts ErrSyntax on every parse-failure class, with
// line/column detail attached.
func TestParseSyntaxErrors(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	for _, query := range []string{
		"",
		"nope",
		"SELECT name",
		"select [name",
		"select name |",
		"where [price] | round(2) > 3",
		"limit 1.5",
		"offset 2.5",
		"limit 99999999999999999999999",
		"select 2.5",
		"rename 1.5 as x",
		"sort 1.5",
		"distinct 1.5",
		"drop 0.5",
		"group 1.5 { x = 1 }",
	} {
		err := parseErr(t, query)
		want.True(errors.Is(err, tq.ErrSyntax), "query %q: %v", query, err)
	}

	err := parseErr(t, "select [name\nmore")
	require.ErrorIs(t, err, tq.ErrSyntax)
	want.Contains(err.Error(), "line")
	want.Contains(err.Error(), "column")
}

// runErr runs a query over input expecting a run failure.
func runErr(t *testing.T, query, input string, opts tq.Options) error {
	t.Helper()
	_, err := runQuery(t, query, input, opts)
	require.Error(t, err, "query %q must fail", query)
	return err
}

// TestPlanUnknownColumn asserts ErrUnknownColumn for unresolved names and
// out-of-range indexes in every stage position.
func TestPlanUnknownColumn(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	for _, query := range []string{
		"select nope",
		"drop nope",
		"where [nope] > 1",
		"derive x = [nope]",
		"rename nope as x",
		"sort nope",
		"distinct nope",
		"group nope { n = counta([name]) }",
		"group lang { n = counta([nope]) }",
		"select [0]",
		"select [5]",
		"select [99999999999999999999999]",
		"select 5",
	} {
		err := runErr(t, query, repos, tq.Options{})
		want.True(errors.Is(err, tq.ErrUnknownColumn), "query %q: %v", query, err)
	}
}

// TestPlanCellRef asserts ErrCellRef for raw A1 references and sheet
// qualifiers anywhere in an expression (ADR 0002).
func TestPlanCellRef(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	for _, query := range []string{
		"where B2 > 0",
		"where [stars] > 0 | derive c = \"other.tsvt\"!A1",
		"derive c = sum(A1:A9)",
		"group lang { n = sum($B$2) }",
	} {
		err := runErr(t, query, repos, tq.Options{})
		want.True(errors.Is(err, tq.ErrCellRef), "query %q: %v", query, err)
	}
}

// TestPlanHeaderless asserts ErrHeaderless for rename and for any name
// reference in headerless mode (ADR 0003).
func TestPlanHeaderless(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	for _, query := range []string{
		"rename [1] as x",
		"select name",
		"where [name] = \"x\"",
		"sort name",
		"group lang { n = counta([1]) }",
	} {
		err := runErr(t, query, "a\tb\n", tq.Options{IsHeaderless: true})
		want.True(errors.Is(err, tq.ErrHeaderless), "query %q: %v", query, err)
	}
}

// TestPlanFailsBeforeRows asserts a plan failure in a LATER stage surfaces
// even though an earlier stage would already empty the stream — planning
// happens before any row is processed (R3).
func TestPlanFailsBeforeRows(t *testing.T) {
	t.Parallel()
	err := runErr(t, "limit 0 | select nope", repos, tq.Options{})
	assert.ErrorIs(t, err, tq.ErrUnknownColumn)
}

// TestStrict asserts ErrStrict on the first error value under Strict — in a
// where predicate, a derive assignment, and a group aggregate — naming the
// value, the column, and the 1-based input row.
func TestStrict(t *testing.T) {
	t.Parallel()
	want := assert.New(t)

	err := runErr(t, "where [stars] / 0 > 1", repos, tq.Options{IsStrict: true})
	require.ErrorIs(t, err, tq.ErrStrict)
	want.Contains(err.Error(), "#DIV/0!")
	want.Contains(err.Error(), "row 1")

	err = runErr(t, "derive bad = if([name] = \"beta\", 1/0, 1)", repos, tq.Options{IsStrict: true})
	require.ErrorIs(t, err, tq.ErrStrict)
	want.Contains(err.Error(), "column bad")
	want.Contains(err.Error(), "row 2")

	err = runErr(t, "group lang { m = avg([name]) }", repos, tq.Options{IsStrict: true})
	require.ErrorIs(t, err, tq.ErrStrict)
	want.Contains(err.Error(), "column m")

	// Lenient mode lets the same programs run, the error values as data.
	_, err = runQuery(t, "derive bad = [stars] / 0", repos, tq.Options{})
	want.NoError(err)

	// A non-boolean predicate is no author error even under strict: the
	// `(expr)=TRUE` wrapper's #VALUE! is an artifact, so the row drops
	// without aborting.
	out, err := runQuery(t, "where [name]", repos, tq.Options{IsStrict: true})
	want.NoError(err)
	want.Equal("name\tstars\tforks\tlang\n", out)
}

// TestBuilderErrorsAtPlan asserts builder-only failure paths surface at Run's
// plan step: a malformed expression is ErrSyntax, a negative count is
// ErrSyntax (the grammar admits only non-negative literals).
func TestBuilderErrorsAtPlan(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	table, err := tq.ReadTable(strings.NewReader(repos), tq.Options{})
	require.NoError(t, err)

	for name, program := range map[string]tq.Program{
		"malformed where":       tq.NewProgram(tq.Where("1 +")),
		"trailing junk":         tq.NewProgram(tq.Where("1 2")),
		"malformed derive":      tq.NewProgram(tq.Derive(tq.Assignment{Name: "x", Expr: "sum("})),
		"malformed aggregate":   tq.NewProgram(tq.GroupBy([]tq.Column{"lang"}, tq.Assignment{Name: "x", Expr: ")"})),
		"negative limit count":  tq.NewProgram(tq.Limit(-1)),
		"negative offset count": tq.NewProgram(tq.Offset(-2)),
	} {
		_, err := program.Run(table, tq.Options{})
		require.Error(t, err, name)
		want.True(errors.Is(err, tq.ErrSyntax), "%s: %v", name, err)
	}
}

// TestReadTableLimit asserts ErrLimit when the input exceeds MaxCells — the
// rectangular rows × columns bound — and admission at or under it.
func TestReadTableLimit(t *testing.T) {
	t.Parallel()
	want := assert.New(t)

	_, err := tq.ReadTable(strings.NewReader(repos), tq.Options{MaxCells: 19})
	require.Error(t, err)
	want.True(errors.Is(err, tq.ErrLimit), "got %v", err)

	_, err = tq.ReadTable(strings.NewReader(repos), tq.Options{MaxCells: 20})
	want.NoError(err)

	_, err = tq.ReadTable(strings.NewReader(repos), tq.Options{})
	want.NoError(err)
}
