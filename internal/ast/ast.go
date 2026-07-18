// Package ast turns tq source into the typed, immutable program AST. It is the
// covered seam over the lifted ANTLR parser in src/grammar/tq: a custom error
// listener converts every lexer/parser report into the constants.ErrSyntax
// sentinel (line/column via With), and the parse-tree walk produces plain value
// types — a Program of Stages whose embedded expressions keep their ORIGINAL
// SOURCE TEXT verbatim, because the plan layer rewrites column references by
// token-span splicing, never by pretty-printing.
package ast

import "strconv"

// ProgramText is the source of a whole tq program.
type ProgramText string

// ExprText is the source of one tq expression, exactly as authored.
type ExprText string

// ColumnContent is the text of a column reference without its brackets: the
// content between `[` and `]`, or a bare stage-position name or integer.
type ColumnContent string

// Verb names a pipeline stage; the values are the SPECIFICATION §3 keywords.
type Verb string

// The stage verbs, in specification order.
const (
	VerbSelect   Verb = "select"
	VerbDrop     Verb = "drop"
	VerbWhere    Verb = "where"
	VerbDerive   Verb = "derive"
	VerbRename   Verb = "rename"
	VerbSort     Verb = "sort"
	VerbDistinct Verb = "distinct"
	VerbLimit    Verb = "limit"
	VerbOffset   Verb = "offset"
	VerbGroup    Verb = "group"
)

// Column is one column reference: a 1-based index when IsIndex (digits-only
// content, SPECIFICATION §4), otherwise a header name. Name always keeps the
// raw content so error messages can quote what the author wrote.
type Column struct {
	Name    string
	Index   int
	IsIndex bool
}

// ColumnOf classifies bracket-free reference content per SPECIFICATION §4:
// digits-only content is a 1-based column index, anything else a header name.
// An index too large for int keeps IsIndex with Index 0, which can never
// resolve (indexes are 1-based), so it surfaces as ErrUnknownColumn at plan.
func ColumnOf(content ColumnContent) Column {
	if !digitsOnly(content) {
		return Column{Name: string(content)}
	}
	n, err := strconv.Atoi(string(content))
	if err != nil {
		n = 0
	}
	return Column{Name: string(content), Index: n, IsIndex: true}
}

// digitsOnly reports whether content is one or more ASCII digits.
func digitsOnly(content ColumnContent) bool {
	if content == "" {
		return false
	}
	for _, r := range content {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// Expr is one embedded expression, held as its original source text. Column
// spans and cell-reference rejection are recovered at plan time by re-parsing
// the text (ParseExprInfo), so parsed and builder-built programs are identical
// values.
type Expr struct {
	Src ExprText
}

// Assignment is one `name = expr` (derive or group aggregate). The name is
// introduced verbatim — bracket content included — never a positional index:
// per SPECIFICATION §4 a digits-only name is an (unusual) header name, which
// stays reachable by position or a later rename.
type Assignment struct {
	Name string
	Expr Expr
}

// RenamePair is one `col as name`: the source column reference and the new
// name, taken verbatim.
type RenamePair struct {
	To   string
	From Column
}

// SortKey is one sort key: the column and whether `-` reversed it.
type SortKey struct {
	Col          Column
	IsDescending bool
}

// Stage is one pipeline stage. Only the fields its Verb uses are set: Columns
// for select/drop/distinct and group keys, Expr for where, Assigns for
// derive and group aggregates, Renames for rename, Keys for sort, Count for
// limit/offset.
type Stage struct {
	Verb    Verb
	Columns []Column
	Expr    Expr
	Assigns []Assignment
	Renames []RenamePair
	Keys    []SortKey
	Count   int
}

// Program is a parsed, unplanned pipeline of stages — an immutable value,
// reusable and safe to share across concurrent runs.
type Program struct {
	Stages []Stage
}
