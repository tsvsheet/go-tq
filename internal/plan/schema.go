package plan

import (
	"slices"

	"github.com/tsvsheet/go-tq/internal/ast"
	"github.com/tsvsheet/go-tq/internal/constants"
	"github.com/tsvsheet/go-tq/internal/exec"
)

// headerName is one header cell's text.
type headerName string

// columnCount is a schema's column count.
type columnCount int

// schema is the column shape a stage plans against: the header names (header
// mode) and the width. Transitions always produce fresh name slices, so the
// caller's table header is never mutated.
type schema struct {
	names     []string
	width     columnCount
	hasHeader bool
}

// headerRow is the schema as an output header — nil in headerless mode.
func (s schema) headerRow() []string {
	if !s.hasHeader {
		return nil
	}
	return s.names
}

// resolve maps one column reference to its 0-based position: an index checks
// the width, a name matches the header exactly and case-sensitively, first
// match winning (§4). A name is unresolvable headerless (ErrHeaderless).
func (s schema) resolve(c ast.Column) (exec.Position, error) {
	if c.IsIndex {
		return s.resolveIndex(c)
	}
	if !s.hasHeader {
		return 0, constants.ErrHeaderless.With(nil, "column", c.Name)
	}
	for i, n := range s.names {
		if n == c.Name {
			return exec.Position(i), nil
		}
	}
	return 0, constants.ErrUnknownColumn.With(nil, "column", c.Name)
}

// resolveIndex checks a 1-based index against the width.
func (s schema) resolveIndex(c ast.Column) (exec.Position, error) {
	if c.Index < 1 || c.Index > int(s.width) {
		return 0, constants.ErrUnknownColumn.With(nil, "index", c.Name, "columns", s.width)
	}
	return exec.Position(c.Index - 1), nil
}

// resolveAll resolves a column list in order.
func (s schema) resolveAll(cols []ast.Column) (exec.Positions, error) {
	at := make(exec.Positions, 0, len(cols))
	for _, c := range cols {
		p, err := s.resolve(c)
		if err != nil {
			return nil, err
		}
		at = append(at, p)
	}
	return at, nil
}

// project is the schema after a select: exactly the listed columns, in order.
func (s schema) project(at exec.Positions) schema {
	if !s.hasHeader {
		return schema{width: columnCount(len(at))}
	}
	names := make([]string, len(at))
	for i, p := range at {
		names[i] = s.names[p]
	}
	return schema{names: names, width: columnCount(len(at)), hasHeader: true}
}

// drop is the schema after dropping the selected positions.
func (s schema) drop(set exec.Selection) schema {
	width := s.width - columnCount(len(set))
	if !s.hasHeader {
		return schema{width: width}
	}
	names := make([]string, 0, width)
	for i, n := range s.names {
		if !set[exec.Position(i)] {
			names = append(names, n)
		}
	}
	return schema{names: names, width: width, hasHeader: true}
}

// rename is the schema with the name at one position replaced.
func (s schema) rename(at exec.Position, to headerName) schema {
	names := slices.Clone(s.names)
	names[at] = string(to)
	return schema{names: names, width: s.width, hasHeader: true}
}

// target resolves a derive assignment's destination: an existing name is
// replaced in place; otherwise the column appends. Headerless names never
// match (they are unemitted syntax, ADR 0003), so every assignment appends.
func (s schema) target(name headerName) (exec.Position, schema) {
	if s.hasHeader {
		if i := slices.Index(s.names, string(name)); i >= 0 {
			return exec.Position(i), s
		}
	}
	return exec.Position(s.width), s.appendCol(name)
}

// appendCol is the schema with one column appended.
func (s schema) appendCol(name headerName) schema {
	next := schema{width: s.width + 1, hasHeader: s.hasHeader}
	if s.hasHeader {
		next.names = append(slices.Clone(s.names), string(name))
	}
	return next
}

// grouped is the schema a group stage emits: key columns first, then one
// column per aggregate assignment.
func (s schema) grouped(keys exec.Positions, aggs []ast.Assignment) schema {
	width := columnCount(len(keys) + len(aggs))
	if !s.hasHeader {
		return schema{width: width}
	}
	names := make([]string, 0, width)
	for _, at := range keys {
		names = append(names, s.names[at])
	}
	for _, a := range aggs {
		names = append(names, a.Name)
	}
	return schema{names: names, width: width, hasHeader: true}
}
