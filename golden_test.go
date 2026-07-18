package tq_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	tq "github.com/tsvsheet/go-tq"
)

// testClock is the fixed compute clock every test injects, so volatile
// functions are deterministic.
func testClock() time.Time {
	return time.Date(2026, 7, 18, 12, 30, 0, 0, time.UTC)
}

// runQuery reads input as a table, runs the parsed query over it, and writes
// the result back to TSV.
func runQuery(t *testing.T, query, input string, opts tq.Options) (string, error) {
	t.Helper()
	program, err := tq.Parse(tq.Query(query))
	require.NoError(t, err, "query %q must parse", query)
	return runProgram(t, program, input, opts)
}

// runProgram runs a program over the input TSV and renders the output TSV.
func runProgram(t *testing.T, program tq.Program, input string, opts tq.Options) (string, error) {
	t.Helper()
	table, err := tq.ReadTable(strings.NewReader(input), opts)
	require.NoError(t, err)
	out, err := program.Run(table, opts)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	require.NoError(t, tq.WriteTable(&b, out))
	return b.String(), nil
}

// repos is the shared header-mode fixture.
const repos = "name\tstars\tforks\tlang\n" +
	"alpha\t1500\t3\tgo\n" +
	"beta\t900\t2\tgo\n" +
	"gamma\t2000\t8\trust\n" +
	"delta\t900\t2\tgo\n"

// TestGolden runs the (table, program, expected TSV) corpus: every verb, both
// header modes, compute-first and Raw, error-value rows, empty and
// header-only tables, ragged rows, and the §5 sort total order.
func TestGolden(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		query string
		input string
		want  string
		opts  tq.Options
	}{
		{
			name:  "select projects and reorders with duplicates",
			query: "select stars, name, stars",
			input: repos,
			want:  "stars\tname\tstars\n1500\talpha\t1500\n900\tbeta\t900\n2000\tgamma\t2000\n900\tdelta\t900\n",
		},
		{
			name:  "select by index and bracketed name",
			query: "select [2], [name]",
			input: repos,
			want:  "stars\tname\n1500\talpha\n900\tbeta\n2000\tgamma\n900\tdelta\n",
		},
		{
			name:  "drop removes and keeps order",
			query: "drop stars, [3]",
			input: repos,
			want:  "name\tlang\nalpha\tgo\nbeta\tgo\ngamma\trust\ndelta\tgo\n",
		},
		{
			name:  "where keeps exactly TRUE",
			query: "where [stars] > 1000",
			input: repos,
			want:  "name\tstars\tforks\tlang\nalpha\t1500\t3\tgo\ngamma\t2000\t8\trust\n",
		},
		{
			name:  "where non-boolean values drop rows",
			query: `where "TRUE"`,
			input: repos,
			want:  "name\tstars\tforks\tlang\n",
		},
		{
			name:  "where multiline expression",
			query: "where [stars] >\n\t1000",
			input: repos,
			want:  "name\tstars\tforks\tlang\nalpha\t1500\t3\tgo\ngamma\t2000\t8\trust\n",
		},
		{
			name:  "derive appends then later stages see it",
			query: "derive ratio = round([stars] / [forks], 2) | where [ratio] > 300 | select name, ratio",
			input: repos,
			want:  "name\tratio\nalpha\t500\nbeta\t450\ndelta\t450\n",
		},
		{
			name:  "derive replaces in place and applies left to right",
			query: "derive stars = [stars] * 2, big = [stars] > 3000 | select name, stars, big",
			input: repos,
			want:  "name\tstars\tbig\nalpha\t3000\tFALSE\nbeta\t1800\tFALSE\ngamma\t4000\tTRUE\ndelta\t1800\tFALSE\n",
		},
		{
			name:  "rename relabels without moving",
			query: "rename name as repo, [4] as language | select repo, language",
			input: repos,
			want:  "repo\tlanguage\nalpha\tgo\nbeta\tgo\ngamma\trust\ndelta\tgo\n",
		},
		{
			name:  "sort multi-key stable with reversal",
			query: "sort lang, -stars, name",
			input: repos,
			want:  "name\tstars\tforks\tlang\nalpha\t1500\t3\tgo\nbeta\t900\t2\tgo\ndelta\t900\t2\tgo\ngamma\t2000\t8\trust\n",
		},
		{
			name:  "sort total order numbers before text bytewise after",
			query: "sort k",
			input: "k\n10\t\n1abc\n9\n\nx\n", // the trailing TAB's overflow cell clips away: outside the 1-column header (§2)
			want:  "k\n9\n10\n\n1abc\nx\n",
		},
		{
			name:  "sort descending reverses the entire key order",
			query: "sort -k",
			input: "k\n10\n1abc\n9\n\nx\n",
			want:  "k\nx\n1abc\n\n10\n9\n",
		},
		{
			name:  "distinct whole row on raw text",
			query: "select stars, forks | distinct",
			input: repos,
			want:  "stars\tforks\n1500\t3\n900\t2\n2000\t8\n",
		},
		{
			name:  "distinct keyed keeps first per key",
			query: "distinct lang | select name, lang",
			input: repos,
			want:  "name\tlang\nalpha\tgo\ngamma\trust\n",
		},
		{
			name:  "limit bounds and beyond row count",
			query: "limit 2",
			input: repos,
			want:  "name\tstars\tforks\tlang\nalpha\t1500\t3\tgo\nbeta\t900\t2\tgo\n",
		},
		{
			name:  "limit beyond row count keeps all",
			query: "limit 99",
			input: repos,
			want:  repos,
		},
		{
			name:  "offset skips and beyond row count empties",
			query: "offset 3",
			input: repos,
			want:  "name\tstars\tforks\tlang\ndelta\t900\t2\tgo\n",
		},
		{
			name:  "offset beyond row count keeps none",
			query: "offset 99",
			input: repos,
			want:  "name\tstars\tforks\tlang\n",
		},
		{
			name:  "limit zero keeps none",
			query: "limit 0",
			input: repos,
			want:  "name\tstars\tforks\tlang\n",
		},
		{
			name:  "group sum counta avg in first-appearance order",
			query: "group lang { total = sum([stars]), n = counta([name]), mean = round(avg([stars]), 2) }",
			input: repos,
			want:  "lang\ttotal\tn\tmean\ngo\t3300\t3\t1100\nrust\t2000\t1\t2000\n",
		},
		{
			name:  "group by two keys emits keys then aggregates",
			query: "group lang, forks { n = counta([name]) }",
			input: repos,
			want:  "lang\tforks\tn\ngo\t3\t1\ngo\t2\t2\nrust\t8\t1\n",
		},
		{
			name:  "keywords as names",
			query: "derive group = [limit] & \"x\" | select [select], group",
			input: "select\tlimit\nA\t1\nB\t2\n",
			want:  "select\tgroup\nA\t1x\nB\t2x\n",
		},
		{
			name:  "keyword-named function calls",
			query: "derive s = sort([stars]) | select s | limit 1",
			input: repos,
			want:  "s\n1500\n",
		},
		{
			name:  "ragged rows read empty cells",
			query: "derive tag = [lang] & \"!\" | select name, lang, tag",
			input: "name\tstars\tforks\tlang\nalpha\t1\n",
			want:  "name\tlang\ttag\nalpha\t\t!\n",
		},
		{
			name:  "empty table passes through",
			query: "limit 5",
			input: "",
			want:  "",
		},
		{
			name:  "header-only table keeps its header",
			query: "select forks, name",
			input: "name\tstars\tforks\tlang\n",
			want:  "forks\tname\n",
		},
		{
			name:  "compute-first queries computed values",
			query: "where [total] > 5 | select name, total",
			input: "name\ta\tb\ttotal\nx\t2\t3\t=B2*C2\ny\t1\t4\t=B3*C3\n",
			want:  "name\ttotal\nx\t6\n",
		},
		{
			name:  "raw mode treats formulas as text",
			query: "select total",
			input: "name\ttotal\nx\t=SUM(B2)\n",
			opts:  tq.Options{IsRaw: true},
			want:  "total\n=SUM(B2)\n",
		},
		{
			name:  "computed error values are data and drop in where",
			query: "where [q] > 1 | select name",
			input: "name\ta\tq\nx\t0\t=2/B2\ny\t1\t=2/B3\n",
			want:  "name\ny\n",
		},
		{
			name:  "error value groups by its text",
			query: "group q { n = counta([name]) }",
			input: "name\ta\tq\nx\t0\t=1/B2\ny\t0\t=1/B3\n",
			want:  "q\tn\n#DIV/0!\t2\n",
		},
		{
			name:  "derive error value writes the error text",
			query: "derive bad = [stars] / 0 | select name, bad | limit 1",
			input: repos,
			want:  "name\tbad\nalpha\t#DIV/0!\n",
		},
		{
			name:  "derive array result reduces to scalar context",
			query: "derive first = sequence(2, 2) | select first | limit 1",
			input: repos,
			want:  "first\n1\n",
		},
		{
			name:  "headerless positional references no header emitted",
			query: "where [2] > 1000 | select [1], [2]",
			input: "alpha\t1500\nbeta\t900\ngamma\t2000\n",
			opts:  tq.Options{IsHeaderless: true},
			want:  "alpha\t1500\ngamma\t2000\n",
		},
		{
			name:  "headerless derive name is unemitted syntax reached by position",
			query: "derive double = [2] * 2 | select [1], [3]",
			input: "alpha\t1500\nbeta\t900\n",
			opts:  tq.Options{IsHeaderless: true},
			want:  "alpha\t3000\nbeta\t1800\n",
		},
		{
			name:  "headerless group keys then aggregates",
			query: "group [2] { n = counta([1]) }",
			input: "a\tgo\nb\tgo\nc\trust\n",
			opts:  tq.Options{IsHeaderless: true},
			want:  "go\t2\nrust\t1\n",
		},
		{
			name:  "headerless drop by position",
			query: "drop [2]",
			input: "a\tb\tc\nd\te\tf\n",
			opts:  tq.Options{IsHeaderless: true},
			want:  "a\tc\nd\tf\n",
		},
		{
			name:  "headerless sort and distinct",
			query: "sort -[1] | distinct",
			input: "1\n3\n3\n2\n",
			opts:  tq.Options{IsHeaderless: true},
			want:  "3\n2\n1\n",
		},
		{
			name:  "volatile functions read the injected clock",
			query: "derive today = today() | select today | limit 1",
			input: repos,
			opts:  tq.Options{At: testClock()},
			want:  "today\n2026-07-18\n",
		},
		{
			name:  "digits-only derive name is a name not an index",
			query: "derive [3] = [stars] | select [3]",
			input: "name\tstars\nx\t7\n",
			want:  "3\n7\n",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			opts := tc.opts
			if opts.At.IsZero() {
				opts.At = testClock()
			}
			got, err := runQuery(t, tc.query, tc.input, opts)
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
