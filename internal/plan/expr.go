package plan

import (
	"strconv"
	"strings"

	tsvsheet "github.com/tsvsheet/go-tsvsheet"

	"github.com/tsvsheet/go-tq/internal/ast"
	"github.com/tsvsheet/go-tq/internal/constants"
	"github.com/tsvsheet/go-tq/internal/exec"
)

// SplicedText is a stage expression rewritten to pure tsvsheet form: column
// references replaced by A1 text, everything else the author's bytes.
type SplicedText string

// CompileFunc compiles one spliced expression through the go-tsvsheet seam.
type CompileFunc func(src SplicedText) (tsvsheet.Expr, error)

// Compile is the production CompileFunc: tsvsheet.CompileExpr, with a failure
// surfaced as tq's ErrSyntax carrying the go-tsvsheet detail via With, so
// errors.Is reaches both engines' sentinels.
func Compile(src SplicedText) (tsvsheet.Expr, error) {
	expr, err := tsvsheet.CompileExpr([]byte(src))
	if err != nil {
		return tsvsheet.Expr{}, constants.ErrSyntax.With(err, "expression", src)
	}
	return expr, nil
}

// srcPos is a rune position within one expression's source text.
type srcPos int

// resolvedRef is one column reference: its token span and resolved position.
type resolvedRef struct {
	span ast.Span
	at   exec.Position
}

// exprPlan is one stage expression ready to splice: the original source runes,
// the resolved column-reference spans, and the string-literal spans whose
// bytes must pass through untouched.
type exprPlan struct {
	src  []rune
	refs []resolvedRef
	strs []ast.Span
}

// exprPlan dissects and resolves one expression: re-parse for spans, reject
// raw A1/sheet-qualified references (ErrCellRef, ADR 0002 — tq addresses
// columns, never cells), and resolve every column reference against the
// schema.
func (p planner) exprPlan(e ast.Expr, s schema) (exprPlan, error) {
	info, err := ast.ParseExprInfo(e.Src)
	if err != nil {
		return exprPlan{}, err
	}
	if info.CellRef != "" {
		return exprPlan{}, constants.ErrCellRef.With(nil, "reference", info.CellRef, "expression", e.Src)
	}
	refs := make([]resolvedRef, 0, len(info.Cols))
	for _, c := range info.Cols {
		at, err := s.resolve(c.Col)
		if err != nil {
			return exprPlan{}, err
		}
		refs = append(refs, resolvedRef{span: c.Span, at: at})
	}
	return exprPlan{src: []rune(string(e.Src)), refs: refs, strs: info.Strings}, nil
}

// refRender renders one resolved column position as A1 text for a splice.
type refRender func(at exec.Position) string

// rowRef renders a row-stage reference: the position's cell in the one-row
// synthetic grid (`[stars]` → e.g. C1).
func rowRef(at exec.Position) string { return a1Col(at) + "1" }

// rangeRefAt renders a group-aggregate reference for a group of h rows: the
// position's whole column across the group (`[stars]` → C1:Ch).
func rangeRefAt(h exec.Height) refRender {
	return func(at exec.Position) string {
		c := a1Col(at)
		return c + "1:" + c + strconv.Itoa(int(h))
	}
}

// a1Col is the spreadsheet column letters for a 0-based position (A, B, … Z,
// AA, …) — bijective base 26.
func a1Col(at exec.Position) string {
	n := int(at) + 1
	buf := make([]byte, 0, 2)
	for n > 0 {
		n--
		buf = append([]byte{byte('A' + n%26)}, buf...)
		n /= 26
	}
	return string(buf)
}

// splice is the token-span rewrite of the expression seam contract: each
// column-reference span is replaced by its A1 rendering and every other byte
// of the author's expression passes through — no pretty-printing — except
// that inter-token whitespace runes the tq grammar allows but a tsvsheet
// formula does not (TAB, CR, LF; a tq program may span lines, §3) normalize
// to spaces. Bytes inside string literals are never touched.
func (e exprPlan) splice(render refRender) SplicedText {
	var b strings.Builder
	pos := srcPos(0)
	for _, r := range e.refs {
		_, _ = b.WriteString(e.normalized(pos, srcPos(r.span.Start)))
		_, _ = b.WriteString(render(r.at))
		pos = srcPos(r.span.Stop + 1)
	}
	_, _ = b.WriteString(e.normalized(pos, srcPos(len(e.src))))
	return SplicedText(b.String())
}

// normalized is src[from:to) with inter-token whitespace mapped to spaces.
func (e exprPlan) normalized(from, to srcPos) string {
	out := make([]rune, 0, to-from)
	for i := from; i < to; i++ {
		out = append(out, e.runeAt(i))
	}
	return string(out)
}

// runeAt is one source rune, with TAB/CR/LF outside string literals mapped to
// a space.
func (e exprPlan) runeAt(i srcPos) rune {
	c := e.src[i]
	if (c == '\t' || c == '\r' || c == '\n') && !e.inString(i) {
		return ' '
	}
	return c
}

// inString reports whether a position falls inside a string literal's span.
func (e exprPlan) inString(i srcPos) bool {
	for _, s := range e.strs {
		if int(i) >= s.Start && int(i) <= s.Stop {
			return true
		}
	}
	return false
}
