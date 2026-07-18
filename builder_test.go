package tq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	tq "github.com/tsvsheet/go-tq"
)

// builderPairs is the pairwise corpus: for every verb, a query and the
// builder program that must equal it — in value and in behavior (R2).
func builderPairs() map[string]struct {
	query   string
	program tq.Program
} {
	return map[string]struct {
		query   string
		program tq.Program
	}{
		"select": {
			query:   "select name, stars",
			program: tq.NewProgram(tq.Select("name", "stars")),
		},
		"select bracketed and indexed": {
			query:   "select [full title], 3, [2]",
			program: tq.NewProgram(tq.Select("[full title]", "3", "[2]")),
		},
		"drop": {
			query:   "drop forks",
			program: tq.NewProgram(tq.Drop("forks")),
		},
		"where": {
			query:   "where [stars] > 1000",
			program: tq.NewProgram(tq.Where("[stars] > 1000")),
		},
		"derive": {
			query: "derive ratio = round([stars] / [forks], 2), hot = [ratio] > 300",
			program: tq.NewProgram(tq.Derive(
				tq.Assignment{Name: "ratio", Expr: "round([stars] / [forks], 2)"},
				tq.Assignment{Name: "hot", Expr: "[ratio] > 300"},
			)),
		},
		"rename": {
			query: "rename name as repo, [4] as language",
			program: tq.NewProgram(
				tq.Rename(tq.RenamePair{Col: "name", As: "repo"}, tq.RenamePair{Col: "[4]", As: "language"}),
			),
		},
		"sort": {
			query:   "sort -stars, name",
			program: tq.NewProgram(tq.SortBy(tq.SortKey{Col: "stars", IsDescending: true}, tq.SortKey{Col: "name"})),
		},
		"distinct whole row": {
			query:   "distinct",
			program: tq.NewProgram(tq.Distinct()),
		},
		"distinct keyed": {
			query:   "distinct lang, forks",
			program: tq.NewProgram(tq.Distinct("lang", "forks")),
		},
		"limit": {
			query:   "limit 2",
			program: tq.NewProgram(tq.Limit(2)),
		},
		"offset": {
			query:   "offset 1",
			program: tq.NewProgram(tq.Offset(1)),
		},
		"group": {
			query: "group lang { total = sum([stars]), n = counta([name]) }",
			program: tq.NewProgram(tq.GroupBy(
				[]tq.Column{"lang"},
				tq.Assignment{Name: "total", Expr: "sum([stars])"},
				tq.Assignment{Name: "n", Expr: "counta([name])"},
			)),
		},
		"full pipeline": {
			query: "where [stars] > 500 | derive ratio = round([stars] / [forks], 2) | " +
				"select name, stars, ratio | sort -stars | limit 10",
			program: tq.NewProgram(
				tq.Where("[stars] > 500"),
				tq.Derive(tq.Assignment{Name: "ratio", Expr: "round([stars] / [forks], 2)"}),
				tq.Select("name", "stars", "ratio"),
				tq.SortBy(tq.SortKey{Col: "stars", IsDescending: true}),
				tq.Limit(10),
			),
		},
	}
}

// TestBuilderEqualsParsed asserts the builder produces Programs identical in
// VALUE to their parsed equivalents (R2) — same AST, same expression source.
func TestBuilderEqualsParsed(t *testing.T) {
	t.Parallel()
	for name, pair := range builderPairs() {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			parsed, err := tq.Parse(tq.Query(pair.query))
			require.NoError(t, err)
			assert.Equal(t, parsed, pair.program)
		})
	}
}

// TestBuilderRunsLikeParsed asserts the builder programs BEHAVE identically
// to their parsed equivalents over a real table, in both header modes where
// applicable.
func TestBuilderRunsLikeParsed(t *testing.T) {
	t.Parallel()
	input := "name\tstars\tforks\tlang\tfull title\n" +
		"alpha\t1500\t3\tgo\tAlpha One\n" +
		"beta\t900\t2\tgo\tBeta Two\n" +
		"gamma\t2000\t8\trust\tGamma Three\n"
	for name, pair := range builderPairs() {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			parsed, err := tq.Parse(tq.Query(pair.query))
			require.NoError(t, err)
			fromQuery, err := runProgram(t, parsed, input, tq.Options{})
			require.NoError(t, err)
			fromBuilder, err := runProgram(t, pair.program, input, tq.Options{})
			require.NoError(t, err)
			assert.Equal(t, fromQuery, fromBuilder)
		})
	}
}

// TestBuilderHeaderless asserts a positional builder program matches its
// parsed equivalent headerless too.
func TestBuilderHeaderless(t *testing.T) {
	t.Parallel()
	opts := tq.Options{IsHeaderless: true}
	parsed, err := tq.Parse(tq.Query("where [2] > 10 | select [1]"))
	require.NoError(t, err)
	built := tq.NewProgram(tq.Where("[2] > 10"), tq.Select("1"))
	assert.Equal(t, parsed, built)

	input := "a\t20\nb\t5\nc\t30\n"
	fromQuery, err := runProgram(t, parsed, input, opts)
	require.NoError(t, err)
	fromBuilder, err := runProgram(t, built, input, opts)
	require.NoError(t, err)
	assert.Equal(t, fromQuery, fromBuilder)
	assert.Equal(t, "a\nc\n", fromBuilder)
}

// TestBuilderEmptyDerive asserts a Derive with no assignments is the
// identity stage — the empty list stays nil, like a parsed program's absent
// fields.
func TestBuilderEmptyDerive(t *testing.T) {
	t.Parallel()
	input := "a\n1\n"
	out, err := runProgram(t, tq.NewProgram(tq.Derive()), input, tq.Options{})
	require.NoError(t, err)
	assert.Equal(t, input, out)
}

// TestBuilderUnbracketedSpacedName asserts the builder's permissive column
// form: a name with spaces needs no brackets programmatically.
func TestBuilderUnbracketedSpacedName(t *testing.T) {
	t.Parallel()
	input := "full title\tstars\nAlpha\t1\n"
	out, err := runProgram(t, tq.NewProgram(tq.Select("full title")), input, tq.Options{})
	require.NoError(t, err)
	assert.Equal(t, "full title\nAlpha\n", out)
}

// TestBuilderGroupByNoKeys pins the builder-only whole-table aggregation:
// with no keys every row falls into one group and the stage emits a single
// row of aggregates. The grammar requires at least one key, so this form has
// no query-text equivalent — it exists only through the builder.
func TestBuilderGroupByNoKeys(t *testing.T) {
	t.Parallel()
	input := "a\n1\n2\n3\n"
	program := tq.NewProgram(tq.GroupBy(nil, tq.Assignment{Name: "n", Expr: "counta([a])"}))
	out, err := runProgram(t, program, input, tq.Options{})
	require.NoError(t, err)
	assert.Equal(t, "n\n3\n", out)
}
