package tq_test

import (
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tsvsheet "github.com/tsvsheet/go-tsvsheet"

	tq "github.com/tsvsheet/go-tq"
)

// TestReadTableComments asserts tsvsheet's comment semantics: a leading `#!`
// line and `# ` lines are skipped and occupy no row, while an error-value
// cell like `#N/A` is data.
func TestReadTableComments(t *testing.T) {
	t.Parallel()
	input := "#!/usr/bin/env tq\n# a comment\nname\tv\n# another\nx\t#N/A\n"
	table, err := tq.ReadTable(strings.NewReader(input), tq.Options{})
	require.NoError(t, err)
	assert.Equal(t, []string{"name", "v"}, table.Header)
	require.Len(t, table.Rows, 1)
	assert.Equal(t, []string{"x", "#N/A"}, table.Rows[0])
}

// TestReadTableHeaderModes asserts the header split: header-first by default,
// all rows data under IsHeaderless, and an empty input yielding an empty table in
// both modes.
func TestReadTableHeaderModes(t *testing.T) {
	t.Parallel()
	want := assert.New(t)

	headed, err := tq.ReadTable(strings.NewReader("h1\th2\na\tb\n"), tq.Options{})
	require.NoError(t, err)
	want.Equal([]string{"h1", "h2"}, headed.Header)
	want.Len(headed.Rows, 1)

	headless, err := tq.ReadTable(strings.NewReader("h1\th2\na\tb\n"), tq.Options{IsHeaderless: true})
	require.NoError(t, err)
	want.Nil(headless.Header)
	want.Len(headless.Rows, 2)

	empty, err := tq.ReadTable(strings.NewReader(""), tq.Options{})
	require.NoError(t, err)
	want.Nil(empty.Header)
	want.Empty(empty.Rows)
}

// TestRoundTrip asserts ReadTable→WriteTable is lossless for formula-free,
// comment-free tables (R7), ragged rows included.
func TestRoundTrip(t *testing.T) {
	t.Parallel()
	for _, input := range []string{
		"a\tb\n1\t2\n",
		"a\tb\tc\n1\n2\t3\n",
		"only\n",
		"",
	} {
		table, err := tq.ReadTable(strings.NewReader(input), tq.Options{})
		require.NoError(t, err)
		var b strings.Builder
		require.NoError(t, tq.WriteTable(&b, table))
		assert.Equal(t, input, b.String(), "input %q", input)
	}
}

// TestReadTableFailure asserts a reader failure passes through as
// go-tsvsheet's own read sentinel — an I/O failure, not a tq program error.
func TestReadTableFailure(t *testing.T) {
	t.Parallel()
	_, err := tq.ReadTable(iotest.ErrReader(assert.AnError), tq.Options{})
	require.Error(t, err)
	assert.ErrorIs(t, err, tsvsheet.ErrReadInput)
}

// failWriter fails every write.
type failWriter struct{}

// Write implements io.Writer, always failing.
func (failWriter) Write(_ []byte) (int, error) { return 0, assert.AnError }

// TestWriteTableFailure asserts a writer failure passes through as
// go-tsvsheet's own write sentinel.
func TestWriteTableFailure(t *testing.T) {
	t.Parallel()
	err := tq.WriteTable(failWriter{}, tq.Table{Header: []string{"a"}, Rows: [][]string{{"1"}}})
	require.Error(t, err)
	assert.ErrorIs(t, err, tsvsheet.ErrWriteFile)
}

// TestComputeFirstMalformedFormula asserts a malformed formula CELL surfaces
// as go-tsvsheet's syntax sentinel naming the cell — input data, not the
// program — and that Raw mode never parses it at all.
func TestComputeFirstMalformedFormula(t *testing.T) {
	t.Parallel()
	program, err := tq.Parse(tq.Query("limit 1"))
	require.NoError(t, err)
	table := tq.Table{Header: []string{"v"}, Rows: [][]string{{"=)"}}}

	_, err = program.Run(table, tq.Options{})
	require.Error(t, err)
	assert.ErrorIs(t, err, tsvsheet.ErrSyntax)
	assert.NotErrorIs(t, err, tq.ErrSyntax)

	out, err := program.Run(table, tq.Options{IsRaw: true})
	require.NoError(t, err)
	assert.Equal(t, [][]string{{"=)"}}, out.Rows)
}

// TestComputeFirstHeaderFormula asserts a formula in the HEADER row triggers
// the compute pass too, and the computed text becomes the column name.
func TestComputeFirstHeaderFormula(t *testing.T) {
	t.Parallel()
	got, err := runQuery(t, "select total", "=\"to\"&\"tal\"\n5\n", tq.Options{})
	require.NoError(t, err)
	assert.Equal(t, "total\n5\n", got)
}

// TestComputeSerializationMatchesWriter asserts the compute pass's grid
// serialization agrees with WriteTable byte for byte: computing a table read
// from TSV equals computing the sheet parsed from that same TSV directly.
func TestComputeSerializationMatchesWriter(t *testing.T) {
	t.Parallel()
	input := "name\tv\nx\t=1+1\ny\t3\n"
	table, err := tq.ReadTable(strings.NewReader(input), tq.Options{})
	require.NoError(t, err)
	var b strings.Builder
	require.NoError(t, tq.WriteTable(&b, table))
	require.Equal(t, input, b.String())

	sheet, err := tsvsheet.Parse([]byte(b.String()))
	require.NoError(t, err)
	grid := sheet.ComputeWith(tsvsheet.ComputeOptions{At: testClock()})

	program, err := tq.Parse(tq.Query("select v"))
	require.NoError(t, err)
	out, err := program.Run(table, tq.Options{At: testClock()})
	require.NoError(t, err)
	require.Len(t, out.Rows, 2)
	assert.Equal(t, grid[1][1], out.Rows[0][0])
	assert.Equal(t, grid[2][1], out.Rows[1][0])
}

// TestNoHeaderIgnoresHeaderField asserts Run under IsHeaderless treats only Rows
// as the table, whatever a hand-built Header holds.
func TestNoHeaderIgnoresHeaderField(t *testing.T) {
	t.Parallel()
	program, err := tq.Parse(tq.Query("select [1]"))
	require.NoError(t, err)
	out, err := program.Run(
		tq.Table{Header: []string{"ignored"}, Rows: [][]string{{"a"}}},
		tq.Options{IsHeaderless: true},
	)
	require.NoError(t, err)
	assert.Nil(t, out.Header)
	assert.Equal(t, [][]string{{"a"}}, out.Rows)
}

// TestRunDoesNotMutateInput asserts Run never mutates the caller's table —
// rows and header are shared values a derive/rename must copy, not write.
func TestRunDoesNotMutateInput(t *testing.T) {
	t.Parallel()
	header := []string{"a", "b"}
	rows := [][]string{{"1", "2"}, {"3", "4"}}
	program, err := tq.Parse(tq.Query("derive a = [a] * 10 | rename b as c | sort -a"))
	require.NoError(t, err)
	_, err = program.Run(tq.Table{Header: header, Rows: rows}, tq.Options{})
	require.NoError(t, err)
	assert.Equal(t, []string{"a", "b"}, header)
	assert.Equal(t, [][]string{{"1", "2"}, {"3", "4"}}, rows)
}

// TestRunOutputSharesNoStorage asserts the returned table is independent of
// the input: mutating either side never shows through the other, even on
// pass-through stages (limit) that would otherwise alias the caller's slices.
func TestRunOutputSharesNoStorage(t *testing.T) {
	t.Parallel()
	header := []string{"a", "b"}
	rows := [][]string{{"1", "2"}, {"3", "4"}, {"5", "6"}}
	program, err := tq.Parse(tq.Query("limit 2"))
	require.NoError(t, err)
	out, err := program.Run(tq.Table{Header: header, Rows: rows}, tq.Options{})
	require.NoError(t, err)

	out.Header[0] = "mutated"
	out.Rows[0][0] = "mutated"
	assert.Equal(t, []string{"a", "b"}, header)
	assert.Equal(t, [][]string{{"1", "2"}, {"3", "4"}, {"5", "6"}}, rows)

	rows[1][1] = "mutated"
	assert.Equal(t, "4", out.Rows[1][1])
}

// TestWideRowsClipToHeader asserts header mode bounds every data row to the
// header's width (§2): a cell beyond the last header column is outside the
// table model, so every verb sees — and emits — the same clipped row.
func TestWideRowsClipToHeader(t *testing.T) {
	t.Parallel()
	in := tq.Table{Header: []string{"a", "b"}, Rows: [][]string{{"x", "y", "z"}, {"x", "y", "Q"}}}

	for query, want := range map[string][][]string{
		"limit 9":      {{"x", "y"}, {"x", "y"}},
		"drop b":       {{"x"}, {"x"}},
		"derive d = 1": {{"x", "y", "1"}, {"x", "y", "1"}},
		"distinct":     {{"x", "y"}},
		"select b":     {{"y"}, {"y"}},
	} {
		program, err := tq.Parse(tq.Query(query))
		require.NoError(t, err, query)
		out, err := program.Run(in, tq.Options{})
		require.NoError(t, err, query)
		assert.Equal(t, want, out.Rows, query)
	}
}

// TestWhereDerivedBooleanColumn pins the coercion contract (§5): cell text
// never reads back as a boolean, so a bare `where [flag]` over a derived
// TRUE/FALSE column drops every row — the comparison against the canonical
// text, `where [flag] = "TRUE"`, is the working idiom. If this test flips,
// the engine's cell-coercion model changed.
func TestWhereDerivedBooleanColumn(t *testing.T) {
	t.Parallel()

	empty, err := runQuery(t, "derive big = [stars] > 1000 | where [big] | select name", repos, tq.Options{})
	require.NoError(t, err)
	assert.Equal(t, "name\n", empty)

	kept, err := runQuery(t, "derive big = [stars] > 1000 | where [big] = \"TRUE\" | select name", repos, tq.Options{})
	require.NoError(t, err)
	assert.Equal(t, "name\nalpha\ngamma\n", kept)
}

// countingFetcher counts Fetch calls, serving a constant one-cell body.
type countingFetcher struct{ calls *int }

// Fetch serves the body "5" under the requested media type, counting calls.
func (f countingFetcher) Fetch(_ tsvsheet.ImportURL, accept tsvsheet.MediaType) (tsvsheet.FetchResult, error) {
	*f.calls++
	return tsvsheet.FetchResult{ContentType: accept, Body: []byte("5")}, nil
}

// TestWhereEvaluatesPredicateOncePerRow asserts the filter costs one predicate
// evaluation per row on the keep path — a side-effecting IMPORT* predicate
// fetches exactly once per input row, never twice.
func TestWhereEvaluatesPredicateOncePerRow(t *testing.T) {
	t.Parallel()
	calls := 0
	in := tq.Table{Header: []string{"a"}, Rows: [][]string{{"1"}, {"2"}, {"3"}}}
	program, err := tq.Parse(tq.Query(`where importcell("https://x.test/") = 5`))
	require.NoError(t, err)
	out, err := program.Run(in, tq.Options{Fetcher: countingFetcher{calls: &calls}})
	require.NoError(t, err)
	assert.Len(t, out.Rows, 3)
	assert.Equal(t, 3, calls)
}
