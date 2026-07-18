package exec_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tsvsheet "github.com/tsvsheet/go-tsvsheet"

	"github.com/tsvsheet/go-tq/internal/exec"
)

// sortPair sorts a two-row single-column table ascending and returns the
// resulting first column — sort.SliceStable performs exactly one comparison,
// so every comparator branch is pinned deterministically.
func sortPair(a, b string) []string {
	rows := exec.SortRows(exec.Rows{{a}, {b}}, exec.Keys{{At: 0}})
	return []string{rows[0][0], rows[1][0]}
}

// TestSortTotalOrder pins the §5 total order pairwise: numbers before
// non-numbers, numeric order among numbers, byte order among texts, and NaN
// (a numeric coercion) kept total — equal to itself, after other numbers.
func TestSortTotalOrder(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	want.Equal([]string{"3", "5"}, sortPair("3", "5"))
	want.Equal([]string{"3", "5"}, sortPair("5", "3"))
	want.Equal([]string{"5", "5"}, sortPair("5", "5"))
	want.Equal([]string{"9", "10"}, sortPair("9", "10"))
	want.Equal([]string{"10", "1abc"}, sortPair("1abc", "10"))
	want.Equal([]string{"1abc", "9zz"}, sortPair("9zz", "1abc"))
	want.Equal([]string{"5", "NaN"}, sortPair("NaN", "5"))
	want.Equal([]string{"5", "NaN"}, sortPair("5", "NaN"))
	want.Equal([]string{"NaN", "NaN"}, sortPair("NaN", "NaN"))
	want.Equal([]string{"NaN", "x"}, sortPair("x", "NaN"))
}

// TestSortStableAndDescending asserts stability for equal keys and that `-`
// reverses the ENTIRE key order — non-numbers first descending.
func TestSortStableAndDescending(t *testing.T) {
	t.Parallel()
	want := assert.New(t)

	stable := exec.SortRows(exec.Rows{{"1", "first"}, {"1", "second"}, {"0", "third"}}, exec.Keys{{At: 0}})
	want.Equal(exec.Rows{{"0", "third"}, {"1", "first"}, {"1", "second"}}, stable)

	desc := exec.SortRows(exec.Rows{{"9"}, {"x"}, {"10"}}, exec.Keys{{At: 0, IsDescending: true}})
	want.Equal(exec.Rows{{"x"}, {"10"}, {"9"}}, desc)
}

// TestSortDoesNotMutateInput asserts the input slice order is untouched.
func TestSortDoesNotMutateInput(t *testing.T) {
	t.Parallel()
	rows := exec.Rows{{"2"}, {"1"}}
	_ = exec.SortRows(rows, exec.Keys{{At: 0}})
	assert.Equal(t, exec.Rows{{"2"}, {"1"}}, rows)
}

// TestProjectRagged asserts projection reads an empty cell for a column a
// ragged row lacks, and duplicates project twice.
func TestProjectRagged(t *testing.T) {
	t.Parallel()
	got := exec.Project(exec.Rows{{"a"}, {"b", "c"}}, exec.Positions{1, 0, 0})
	assert.Equal(t, exec.Rows{{"", "a", "a"}, {"c", "b", "b"}}, got)
}

// TestDropColsKeepsOrder asserts non-dropped cells keep their order, ragged
// tails included.
func TestDropColsKeepsOrder(t *testing.T) {
	t.Parallel()
	got := exec.DropCols(exec.Rows{{"a", "b", "c"}, {"a"}}, exec.Selection{1: true})
	assert.Equal(t, exec.Rows{{"a", "c"}, {"a"}}, got)
}

// TestDistinctRaggedEqualsEmpty asserts distinct compares by COLUMN, so a
// missing cell equals an empty one.
func TestDistinctRaggedEqualsEmpty(t *testing.T) {
	t.Parallel()
	got := exec.DistinctRows(exec.Rows{{"a"}, {"a", ""}, {"a", "x"}}, exec.Positions{0, 1})
	assert.Equal(t, exec.Rows{{"a"}, {"a", "x"}}, got)
}

// TestLimitOffsetBounds asserts the exact boundary behavior on both sides of
// the row count.
func TestLimitOffsetBounds(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	rows := exec.Rows{{"1"}, {"2"}}
	want.Equal(exec.Rows{{"1"}}, exec.LimitRows(rows, 1))
	want.Equal(rows, exec.LimitRows(rows, 2))
	want.Equal(rows, exec.LimitRows(rows, 3))
	want.Equal(exec.Rows{{"2"}}, exec.OffsetRows(rows, 1))
	want.Empty(exec.OffsetRows(rows, 2))
	want.Empty(exec.OffsetRows(rows, 3))
}

// TestGroupRowsCompileFailure asserts an aggregate whose height compiler
// fails surfaces the error from GroupRows.
func TestGroupRowsCompileFailure(t *testing.T) {
	t.Parallel()
	agg := exec.Agg{Name: "x", CompileAt: func(exec.Height) (tsvsheet.Expr, error) {
		return tsvsheet.Expr{}, assert.AnError
	}}
	_, err := exec.GroupRows(exec.Rows{{"k", "1"}}, 2, exec.Positions{0}, []exec.Agg{agg}, exec.Env{})
	require.Error(t, err)
	assert.ErrorIs(t, err, assert.AnError)
}
