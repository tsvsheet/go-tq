package tq

import "github.com/tsvsheet/go-tq/internal/constants"

// The closed program-error set (SPECIFICATION §8), matchable with errors.Is.
// Each sentinel is distinct from go-tsvsheet's; an embedded-expression compile
// failure is tq's ErrSyntax with the go-tsvsheet detail attached via With, so
// errors.Is reaches both. Everything else — coercion surprises, lookup
// misses, division by zero — is an error value in the data, never a Go error.
const (
	// ErrSyntax: the program does not parse (line/column attached) — including
	// a fractional limit/offset and an embedded-expression failure at compile.
	ErrSyntax = constants.ErrSyntax
	// ErrUnknownColumn: plan-time — a reference names no column or an index is
	// out of range.
	ErrUnknownColumn = constants.ErrUnknownColumn
	// ErrCellRef: plan-time — an expression contains a raw A1 reference or
	// sheet qualifier; tq addresses columns, never cells.
	ErrCellRef = constants.ErrCellRef
	// ErrHeaderless: plan-time — rename (or another name-requiring form) in
	// headerless mode (IsHeaderless).
	ErrHeaderless = constants.ErrHeaderless
	// ErrStrict: run-time under IsStrict — an expression produced an error
	// value, named with its column and 1-based input row.
	ErrStrict = constants.ErrStrict
	// ErrLimit: the input exceeded Options.MaxCells (raised by ReadTable).
	ErrLimit = constants.ErrLimit
)
