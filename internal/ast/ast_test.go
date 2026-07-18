package ast_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tsvsheet/go-tq/internal/ast"
	"github.com/tsvsheet/go-tq/internal/constants"
)

// TestColumnOf asserts the §4 classification: digits-only content is a
// 1-based index (keeping its raw text), anything else — empty content and an
// int-overflowing digit run included — resolves as written.
func TestColumnOf(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	want.Equal(ast.Column{Name: "stars"}, ast.ColumnOf("stars"))
	want.Equal(ast.Column{Name: "full title"}, ast.ColumnOf("full title"))
	want.Equal(ast.Column{Name: ""}, ast.ColumnOf(""))
	want.Equal(ast.Column{Name: "2.5"}, ast.ColumnOf("2.5"))
	want.Equal(ast.Column{Name: "7", Index: 7, IsIndex: true}, ast.ColumnOf("7"))
	want.Equal(ast.Column{Name: "007", Index: 7, IsIndex: true}, ast.ColumnOf("007"))
	// An index too large for int keeps IsIndex with Index 0 — 1-based, so it
	// can never resolve and surfaces as ErrUnknownColumn at plan.
	huge := ast.ColumnOf("99999999999999999999999")
	want.True(huge.IsIndex)
	want.Zero(huge.Index)
}

// TestParseProgramAST asserts the full typed AST of a program spanning every
// stage kind, expression source text kept verbatim.
func TestParseProgramAST(t *testing.T) {
	t.Parallel()
	prog, err := ast.ParseProgram(
		"select name, [full title], 3 | drop notes | where [stars] > 1000 | " +
			"derive ratio = round([stars] / [forks], 2) | rename lang as language | " +
			"sort -stars, name | distinct | distinct lang | limit 10 | offset 2 | " +
			"group lang { total = sum([stars]) }",
	)
	require.NoError(t, err)
	want := ast.Program{Stages: []ast.Stage{
		{Verb: ast.VerbSelect, Columns: []ast.Column{
			{Name: "name"}, {Name: "full title"}, {Name: "3", Index: 3, IsIndex: true},
		}},
		{Verb: ast.VerbDrop, Columns: []ast.Column{{Name: "notes"}}},
		{Verb: ast.VerbWhere, Expr: ast.Expr{Src: "[stars] > 1000"}},
		{Verb: ast.VerbDerive, Assigns: []ast.Assignment{
			{Name: "ratio", Expr: ast.Expr{Src: "round([stars] / [forks], 2)"}},
		}},
		{Verb: ast.VerbRename, Renames: []ast.RenamePair{{From: ast.Column{Name: "lang"}, To: "language"}}},
		{Verb: ast.VerbSort, Keys: []ast.SortKey{
			{Col: ast.Column{Name: "stars"}, IsDescending: true}, {Col: ast.Column{Name: "name"}},
		}},
		{Verb: ast.VerbDistinct},
		{Verb: ast.VerbDistinct, Columns: []ast.Column{{Name: "lang"}}},
		{Verb: ast.VerbLimit, Count: 10},
		{Verb: ast.VerbOffset, Count: 2},
		{Verb: ast.VerbGroup, Columns: []ast.Column{{Name: "lang"}}, Assigns: []ast.Assignment{
			{Name: "total", Expr: ast.Expr{Src: "sum([stars])"}},
		}},
	}}
	assert.Equal(t, want, prog)
}

// TestParseProgramKeywordNames asserts verb keywords come back as ordinary
// names everywhere a name may appear (§7), and that a bracketed digits-only
// derive target is a NAME, never an index.
func TestParseProgramKeywordNames(t *testing.T) {
	t.Parallel()
	prog, err := ast.ParseProgram("select [select], sort | derive group = [limit] & \"x\", [3] = 1")
	require.NoError(t, err)
	want := ast.Program{Stages: []ast.Stage{
		{Verb: ast.VerbSelect, Columns: []ast.Column{{Name: "select"}, {Name: "sort"}}},
		{Verb: ast.VerbDerive, Assigns: []ast.Assignment{
			{Name: "group", Expr: ast.Expr{Src: `[limit] & "x"`}},
			{Name: "3", Expr: ast.Expr{Src: "1"}},
		}},
	}}
	assert.Equal(t, want, prog)
}

// TestParseProgramMultilineExprKeepsSource asserts an expression spanning
// lines keeps its ORIGINAL bytes — newlines included — for span splicing.
func TestParseProgramMultilineExprKeepsSource(t *testing.T) {
	t.Parallel()
	prog, err := ast.ParseProgram("where [price] >\n\t4.50")
	require.NoError(t, err)
	require.Len(t, prog.Stages, 1)
	assert.Equal(t, ast.ExprText("[price] >\n\t4.50"), prog.Stages[0].Expr.Src)
}

// TestParseProgramErrors asserts every parse-failure class is
// constants.ErrSyntax, with line/column carried for a positioned report.
func TestParseProgramErrors(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	for _, src := range []string{
		"", "wc", "SELECT a", "select [x", "select a |", "where | select a",
		"rename old new", "group l total = 1", "where [x] | wc",
		"limit 1.5", "offset 0.5", "limit 99999999999999999999999",
		"select 2.5", "sort 2.5", "rename 2.5 as x", "distinct 2.5", "group 2.5 { x = 1 }",
	} {
		_, err := ast.ParseProgram(ast.ProgramText(src))
		require.Error(t, err, "source %q", src)
		want.True(errors.Is(err, constants.ErrSyntax), "source %q: %v", src, err)
	}

	_, err := ast.ParseProgram("select [x\ny")
	require.ErrorIs(t, err, constants.ErrSyntax)
	want.Contains(err.Error(), "line")
	want.Contains(err.Error(), "column")
}

// TestParseExprInfoSpans asserts the dissection the splicer depends on:
// column spans in source order with classified columns, and string-literal
// spans — rune-indexed, so multi-byte runes stay aligned.
func TestParseExprInfoSpans(t *testing.T) {
	t.Parallel()
	want := assert.New(t)

	info, err := ast.ParseExprInfo(`sum( [a] , [2] ) & "π✓"`)
	require.NoError(t, err)
	want.Empty(info.CellRef)
	require.Len(t, info.Cols, 2)
	want.Equal(ast.ColumnSpan{Col: ast.Column{Name: "a"}, Span: ast.Span{Start: 5, Stop: 7}}, info.Cols[0])
	want.Equal(
		ast.ColumnSpan{Col: ast.Column{Name: "2", Index: 2, IsIndex: true}, Span: ast.Span{Start: 11, Stop: 13}},
		info.Cols[1],
	)
	require.Len(t, info.Strings, 1)
	want.Equal(ast.Span{Start: 19, Stop: 22}, info.Strings[0])

	unicode, err := ast.ParseExprInfo(`"é" & [é col]`)
	require.NoError(t, err)
	require.Len(t, unicode.Cols, 1)
	want.Equal(ast.Column{Name: "é col"}, unicode.Cols[0].Col)
	want.Equal(ast.Span{Start: 6, Stop: 12}, unicode.Cols[0].Span)
}

// TestParseExprInfoCellRef asserts raw A1 and sheet-qualified references are
// reported with the FIRST offender's source text.
func TestParseExprInfoCellRef(t *testing.T) {
	t.Parallel()
	want := assert.New(t)

	info, err := ast.ParseExprInfo("B2 > B3")
	require.NoError(t, err)
	want.Equal("B2", info.CellRef)

	qualified, err := ast.ParseExprInfo(`"other.tsvt"!A1 + [x]`)
	require.NoError(t, err)
	want.Equal(`"other.tsvt"!A1`, qualified.CellRef)

	ranged, err := ast.ParseExprInfo("sum($A$1:C3)")
	require.NoError(t, err)
	want.Equal("$A$1:C3", ranged.CellRef)
}

// TestParseExprInfoErrors asserts a malformed expression and trailing input
// are constants.ErrSyntax.
func TestParseExprInfoErrors(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	for _, src := range []string{"sum(", ")", "1 +", `"unterminated`, "1 2", "", "x"} {
		_, err := ast.ParseExprInfo(ast.ExprText(src))
		require.Error(t, err, "source %q", src)
		want.True(errors.Is(err, constants.ErrSyntax), "source %q: %v", src, err)
	}
}
