package tq_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tsvsheet "github.com/tsvsheet/go-tsvsheet"

	tq "github.com/tsvsheet/go-tq"
)

// alignRows is the differential data: [a]=2 (number), [b]=x (string),
// [c]=3.5 (number); row 2 has an empty [a], a number [b], and an error [c].
func alignRows() [][]string {
	return [][]string{
		{"2", "x", "3.5"},
		{"", "7", "#N/A"},
	}
}

// alignCorpus pairs each tq expression with its sheet-formula template, where
// {r} is the 1-based sheet row — the same expression addressed by column in
// tq and by cell in a .tsvt. A non-nil want anchors the expected derived
// values, so the two routes cannot agree on a wrong answer unnoticed.
func alignCorpus() []struct {
	tq    string
	sheet string
	want  []string
} {
	return []struct {
		tq    string
		sheet string
		want  []string
	}{
		{tq: `[a]`, sheet: `A{r}`, want: []string{"2", ""}},
		{tq: `[b]`, sheet: `B{r}`, want: []string{"x", "7"}},
		{tq: `[c]`, sheet: `C{r}`, want: []string{"3.5", "#N/A"}},
		{tq: `1.5`, sheet: `1.5`, want: []string{"1.5", "1.5"}},
		{tq: `"hi"`, sheet: `"hi"`, want: []string{"hi", "hi"}},
		{tq: `TRUE`, sheet: `TRUE`, want: []string{"TRUE", "TRUE"}},
		{tq: `#DIV/0!`, sheet: `#DIV/0!`, want: []string{"#DIV/0!", "#DIV/0!"}},
		{tq: `-[a]`, sheet: `-A{r}`, want: []string{"-2", "-0"}},
		{tq: `+[c]`, sheet: `+C{r}`, want: []string{"3.5", "#N/A"}},
		{tq: `[a]%`, sheet: `A{r}%`, want: []string{"0.02", "0"}},
		{tq: `[c]^2`, sheet: `C{r}^2`, want: []string{"12.25", "#N/A"}},
		{tq: `[a]*[c]`, sheet: `A{r}*C{r}`, want: []string{"7", "#N/A"}},
		{tq: `[a]+[c]`, sheet: `A{r}+C{r}`, want: []string{"5.5", "#N/A"}},
		{tq: `[c]-[a]`, sheet: `C{r}-A{r}`, want: []string{"1.5", "#N/A"}},
		{tq: `[a]/4`, sheet: `A{r}/4`, want: []string{"0.5", "0"}},
		{tq: `[a]/0`, sheet: `A{r}/0`, want: []string{"#DIV/0!", "#DIV/0!"}},
		{tq: `[a]&[b]`, sheet: `A{r}&B{r}`, want: []string{"2x", "7"}},
		{tq: `[a]=2`, sheet: `A{r}=2`, want: []string{"TRUE", "#VALUE!"}},
		{tq: `[a]<>2`, sheet: `A{r}<>2`, want: []string{"FALSE", "#VALUE!"}},
		{tq: `[a]<3`, sheet: `A{r}<3`, want: []string{"TRUE", "#VALUE!"}},
		{tq: `[a]<=1`, sheet: `A{r}<=1`, want: []string{"FALSE", "#VALUE!"}},
		{tq: `[a]>1`, sheet: `A{r}>1`, want: []string{"TRUE", "#VALUE!"}},
		{tq: `[a]>=3`, sheet: `A{r}>=3`, want: []string{"FALSE", "#VALUE!"}},
		{tq: `[b]+1`, sheet: `B{r}+1`, want: []string{"#VALUE!", "8"}},
		{tq: `sum([a], [c])`, sheet: `sum(A{r}, C{r})`, want: []string{"5.5", "#N/A"}},
		{tq: `round([c], 0)`, sheet: `round(C{r}, 0)`, want: []string{"4", "#N/A"}},
		{tq: `if([a] > 1, "big", "small")`, sheet: `if(A{r} > 1, "big", "small")`, want: []string{"big", "#VALUE!"}},
		{tq: `left([b], 1)`, sheet: `left(B{r}, 1)`, want: []string{"x", "7"}},
		{tq: `len([b] & "y")`, sheet: `len(B{r} & "y")`, want: []string{"2", "2"}},
		{tq: `bogus([a])`, sheet: `bogus(A{r})`, want: []string{"#NAME?", "#NAME?"}},
		{tq: `today()`, sheet: `today()`, want: []string{"2026-07-18", "2026-07-18"}},
		// A bare dynamic array is excluded: in a sheet it SPILLS (a layout
		// phenomenon), while the seam reduces to scalar context — covered by
		// the golden suite. A scalar consumer keeps the array builtins here.
		{tq: `sum(sequence(2, 2))`, sheet: `sum(sequence(2, 2))`, want: []string{"10", "10"}},
		{tq: `(([a] - 1) * 10)%`, sheet: `((A{r} - 1) * 10)%`, want: []string{"0.1", "-0.1"}},
	}
}

// sheetColumn computes the expression as a formula cell per data row in a
// .tsvt sheet over the same grid, returning the computed formula column.
func sheetColumn(t *testing.T, rows [][]string, template string, opts tsvsheet.ComputeOptions) []string {
	t.Helper()
	var b strings.Builder
	for i, row := range rows {
		formula := strings.ReplaceAll(template, "{r}", strconv.Itoa(i+1))
		_, _ = b.WriteString(strings.Join(row, "\t") + "\t=" + formula + "\n")
	}
	sheet, err := tsvsheet.Parse([]byte(b.String()))
	require.NoError(t, err)
	grid := sheet.ComputeWith(opts)
	out := make([]string, len(rows))
	for i, row := range rows {
		out[i] = grid[i][len(row)]
	}
	return out
}

// deriveColumn runs `derive out = expr` over the same data and returns the
// derived column.
func deriveColumn(t *testing.T, rows [][]string, expr string, opts tq.Options) []string {
	t.Helper()
	program, err := tq.Parse(tq.Query("derive out = " + expr))
	require.NoError(t, err)
	result, err := program.Run(tq.Table{Header: []string{"a", "b", "c"}, Rows: rows}, opts)
	require.NoError(t, err)
	out := make([]string, len(result.Rows))
	for i, row := range result.Rows {
		out[i] = row[3]
	}
	return out
}

// TestExpressionAlignment is the acceptance differential: every corpus
// expression run via derive equals the same expression computed as a formula
// cell in a .tsvt sheet over the same data, cell for cell — and, where
// anchored, both match the contract's expected value.
func TestExpressionAlignment(t *testing.T) {
	t.Parallel()
	for _, tc := range alignCorpus() {
		t.Run(tc.tq, func(t *testing.T) {
			t.Parallel()
			derived := deriveColumn(t, alignRows(), tc.tq, tq.Options{At: testClock()})
			computed := sheetColumn(t, alignRows(), tc.sheet, tsvsheet.ComputeOptions{At: testClock()})
			assert.Equal(t, computed, derived, "derive and sheet diverge for %q", tc.tq)
			if tc.want != nil {
				assert.Equal(t, tc.want, derived, "anchored value for %q", tc.tq)
			}
		})
	}
}
