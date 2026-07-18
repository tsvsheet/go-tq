package ast

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/tsvsheet/go-tq/internal/constants"
	grammar "github.com/tsvsheet/go-tq/src/grammar/tq"
)

// sourceText is any tq source handed to the ANTLR toolchain — a whole program
// or a standalone expression.
type sourceText string

// errorSink holds the first collected syntax error so an errorCollector can
// record it from a value-receiver method (the sink is shared by pointer).
type errorSink struct {
	err error
}

// errorCollector records the first syntax error as a sentinel. The mutable
// error lives behind the sink pointer, so the antlr ErrorListener callback is a
// value-receiver method whose write still persists.
type errorCollector struct {
	antlr.DefaultErrorListener
	sink *errorSink
}

// SyntaxError implements antlr.ErrorListener, converting the report into
// constants.ErrSyntax; only the first error is kept.
func (c errorCollector) SyntaxError(
	_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException,
) {
	if c.sink.err == nil {
		c.sink.err = constants.ErrSyntax.With(nil, "line", line, "column", column, "message", msg)
	}
}

// newParser wires the lifted lexer and parser over src with a fresh error
// collector replacing the default console listeners.
func newParser(src sourceText) (*errorCollector, *grammar.TqParser) {
	collector := &errorCollector{sink: &errorSink{}}

	lexer := grammar.NewTqLexer(antlr.NewInputStream(string(src)))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(collector)

	parser := grammar.NewTqParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	parser.RemoveErrorListeners()
	parser.AddErrorListener(collector)
	return collector, parser
}

// ParseProgram parses a whole tq program into the typed Program AST, or
// constants.ErrSyntax with line/column detail. There are no partial programs:
// any report from the lexer or parser fails the parse.
func ParseProgram(src ProgramText) (Program, error) {
	collector, parser := newParser(sourceText(src))
	tree := parser.Program()
	if collector.sink.err != nil {
		return Program{}, collector.sink.err
	}
	return buildProgram(src, tree)
}

// Span is one token's inclusive rune range within its expression source.
type Span struct {
	Start int
	Stop  int
}

// ColumnSpan is one column reference inside an expression: the classified
// column and the exact span of its `[...]` token, ready for splicing.
type ColumnSpan struct {
	Col  Column
	Span Span
}

// ExprInfo is the plan-facing dissection of one expression: every column
// reference with its token span, every string-literal span (whose bytes must
// pass through splicing untouched), and the first raw A1/sheet-qualified
// reference if any (its source text; "" when the expression has none).
type ExprInfo struct {
	Cols    []ColumnSpan
	CellRef string
	Strings []Span
}

// ParseExprInfo re-parses one expression (a stage's stored source text, or a
// builder-supplied one) and reports its column-reference spans, string-literal
// spans, and any raw cell reference. A malformed expression — possible only
// for builder-built programs, since parsed stages store text the grammar
// already accepted — is constants.ErrSyntax.
func ParseExprInfo(src ExprText) (ExprInfo, error) {
	collector, parser := newParser(sourceText(src))
	tree := parser.TqExpr()
	if collector.sink.err != nil {
		return ExprInfo{}, collector.sink.err
	}
	if parser.GetCurrentToken().GetTokenType() != antlr.TokenEOF {
		return ExprInfo{}, constants.ErrSyntax.With(nil, "message", "unexpected input after expression")
	}
	sink := &exprSink{}
	antlr.ParseTreeWalkerDefault.Walk(&exprListener{sink: sink}, tree)
	return sink.info, nil
}

// exprSink accumulates an ExprInfo during the listener walk (shared by pointer
// so the listener keeps value receivers).
type exprSink struct {
	info ExprInfo
}

// exprListener collects column-reference spans, string-literal spans, and raw
// cell references from an expression parse tree. The walk visits tokens in
// source order, so the collected spans are ordered and disjoint.
type exprListener struct {
	grammar.BaseTqParserListener
	sink *exprSink
}

// EnterTqColumnExpr records a `[...]` column operand: its classified column and
// the token's rune span.
func (l exprListener) EnterTqColumnExpr(ctx *grammar.TqColumnExprContext) {
	tok := ctx.COLUMN().GetSymbol()
	text := tok.GetText()
	l.sink.info.Cols = append(l.sink.info.Cols, ColumnSpan{
		Col:  ColumnOf(ColumnContent(text[1 : len(text)-1])),
		Span: Span{Start: tok.GetStart(), Stop: tok.GetStop()},
	})
}

// EnterTqStringExpr records a string literal's span: splicing must not
// normalize whitespace inside it.
func (l exprListener) EnterTqStringExpr(ctx *grammar.TqStringExprContext) {
	tok := ctx.STRING().GetSymbol()
	l.sink.info.Strings = append(l.sink.info.Strings, Span{Start: tok.GetStart(), Stop: tok.GetStop()})
}

// EnterTqRefExpr records a raw A1 or sheet-qualified reference — grammatically
// valid but a plan-time program error (ErrCellRef); the first one's text is
// kept for the report.
func (l exprListener) EnterTqRefExpr(ctx *grammar.TqRefExprContext) {
	if l.sink.info.CellRef == "" {
		l.sink.info.CellRef = ctx.GetText()
	}
}
