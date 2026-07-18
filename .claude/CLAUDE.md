# go-tq

The canonical engine of the tq query language (module `github.com/tsvsheet/go-tq`, root package `tq`): a `|`-separated pipeline of relational verbs over a TSV/tsvt table, every embedded expression evaluated through go-tsvsheet's expression seam. The normative language definition lives in [tsvsheet/tq](https://github.com/tsvsheet/tq) (SPECIFICATION.md + grammars); the API contract and SDD tree live in `tsvsheet/_projects/specs/tq/`.

## Layout

- [src/grammar/tq](../src/grammar/tq/) — the ANTLR-generated parser LIFTED from `tsvsheet/tq` `gen/go` (regenerate there with `make go`, Docker required, then re-copy). Committed, `DO NOT EDIT`, excluded from the coverage/vet gates via [Makefile.local](../Makefile.local).
- [internal/ast](../internal/ast/) — the covered seam over the generated parser: syntax-error listener → `ErrSyntax` (line/column via `With`), parse-tree walk → typed immutable AST. Expressions keep their ORIGINAL SOURCE TEXT: the plan layer rewrites column references by token-span splicing, never by pretty-printing.
- [internal/plan](../internal/plan/) — resolution + compilation: header/arity resolution (`ErrUnknownColumn`), raw A1/sheet-qualifier rejection (`ErrCellRef`), headerless rules (`ErrHeaderless`), token-span splice of `[col]` → A1 (row stages → `C1` against a one-row grid; group aggregates → `C1:Cn`, compiled per height and cached), one `CompileExpr` per unique (expression, shape). The compiler is injected (`Options.Compile`) so seam failure paths are honestly coverable.
- [internal/exec](../internal/exec/) — verb execution over the table model (SPECIFICATION §3–§6): stable total-order sort, first-wins distinct on raw text, first-appearance groups, TRUE-only where (`(expr)=TRUE` disambiguates boolean TRUE from the string), derive left-to-right with in-place replacement, strict-mode aborts (`ErrStrict`).
- Root package — the facade per the tq-api contract: `Query`/`Program`/`Table`/`Options`, `Parse`, `Program.Run`, `ReadTable`/`WriteTable` (grid I/O delegated to go-tsvsheet; `MaxCells` → `ErrLimit`), and the builder (`Select` … `GroupBy`, `NewProgram`) producing Programs value-identical to parsed ones.

## Rules

- Do not redesign semantics here: SPECIFICATION.md and the ADRs in `tsvsheet/_projects/specs/tq/decisions/` are normative. tq queries computed VALUES (compute-first unless `Raw`), addresses columns never cells, and `|` always means "next stage".
- Sentinels are `errs.Const` in [internal/constants](../internal/constants/), each distinct from go-tsvsheet's strings; never `fmt.Errorf`/`errors.New`.
- Value receivers only; named parameter types; gocognit ≤ 7; 100.0% aggregate coverage with `src/grammar` excluded. `make check` must exit 0 (`make tools` first to populate `${GOBIN}`).
- The corpus under [testdata/corpus](../testdata/corpus/) is copied verbatim from `tsvsheet/tq/testdata` — refresh it when the grammar repo's corpus changes, never edit it here.
- Shared `Makefile`, `.golangci.yaml`, `.editorconfig`, `.gitignore`, `.github/` are distributed by `nicerobot/tools.repository` — never edit in-tree; repo-local gate scoping goes in `Makefile.local`.
