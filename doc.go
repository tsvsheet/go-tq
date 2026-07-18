// Package tq is the canonical engine of tq — the query language of the
// tsvsheet ecosystem: a `|`-separated pipeline of relational verbs (select,
// drop, where, derive, rename, sort, distinct, limit, offset, group) over a
// TSV or .tsvt table, with every embedded expression written in the tsvsheet
// formula language. TSV in, TSV out; between cut and jq.
//
// Parse turns query text into an immutable Program; NewProgram and the
// per-verb stage constructors (Select, Where, Derive, …) build the identical
// Program programmatically. Program.Run plans the pipeline against the input
// table's header (or arity, headerless) — resolving every column reference,
// rejecting raw A1 references, and compiling each stage expression once
// through go-tsvsheet's expression seam — then executes the verbs over the
// table model. ReadTable and WriteTable delegate grid I/O (comment semantics
// included) to go-tsvsheet, so the two engines never disagree; when the input
// holds formula cells and IsRaw is unset, Run computes the sheet first and
// queries the computed values (ADR 0003).
//
// The engine is filesystem- and network-free by construction: the compute
// pass's clock, limits, sheet loader, and import fetcher arrive injected via
// Options. Errors returned to callers are the errs.Const sentinels
// re-exported from errors.go — the closed program-error set of
// SPECIFICATION §8 — matchable with errors.Is; expression evaluation
// failures are error values in the data, exactly as in a sheet.
package tq
