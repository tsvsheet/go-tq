// Package constants declares the tq engine's sentinel error values — the
// closed set of program-error classes in SPECIFICATION §8. The error mechanism
// (the matchable string type) lives in the shared gomatic/go-error library;
// these values are this package's own, and every string is distinct from
// go-tsvsheet's sentinels so errors.Is can tell the two engines apart (an
// embedded-expression failure carries the go-tsvsheet detail via With).
package constants

// Imported bare (the package is named error); this file declares only sentinels
// and uses no builtin error type, so each declaration reads errs.Const.
import errs "github.com/gomatic/go-error"

// Keep these constants sorted alphabetically.
const (
	ErrCellRef       errs.Const = "cell reference in a tq expression"
	ErrHeaderless    errs.Const = "name requires a header"
	ErrLimit         errs.Const = "input exceeds the cell ceiling"
	ErrStrict        errs.Const = "expression produced an error value"
	ErrSyntax        errs.Const = "tq syntax error"
	ErrUnknownColumn errs.Const = "unknown column"
)
