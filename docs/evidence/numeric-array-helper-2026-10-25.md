# Numeric-array original helper / polyfill differential — 2026-10-25

**PASS: 31 cases against actual PHP8.3.36.** Two byte-unchanged original
declarations were translated and executed together. This is a new standalone
unit, not a larger integrated WordPress bundle or a bootstrap compatibility claim.
The old17-original/149-case integration and #25 rejection evidence are unchanged.

## Snapshot and execution

- Gossip branch `codex/gohp-numeric-array-helper`, based on #25
  `1743fa1b80e1b2bdbdbea8001415b7558531e96c`.
- Test commit `5fa46a8ac5d0441c784dd7d86a63c889fccac164`, tree
  `3a028ecc8766b9c604b3cd81ed8ad87d33cdd1d6`. Staged diff check and clean commit
  preceded execution. Documentation/evidence additions afterward do not change
  tested code, source, corpus, pin or configuration.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`.
  `functions.php:5448–5454` and `compat.php:635–643` declarations unchanged.
  Native PHP loads the full guarded polyfill at `compat.php:623–644`; the
  observer verifies user-defined provenance. Compiler input excludes only that
  top-level loading wrapper, not parameter/return types or function bodies.
- GoHP PR #53 `bd419a2c6c72e645f0f0d7d09aa01eea15f70913`, tree
  `481c23cfdd84c1ac75995db912d8e488faedf39b`; no compiler edits/retests/CI dispatch.
  PR #52's separate composition fixture is not included or claimed here.
- Host Go1.22.1 darwin/arm64; fixed builder Go1.25.14 linux/arm64; actual
  PHP8.3.36, C locale. Non-root/read-only/network-disabled/resource restrictions
  and the same-process PHP admission guard are unchanged.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressNumericArrayHelper$'
go vet ./tests/gohp
```

First run PASS14.05s/package14.290s; vetPASS, command exit0. No failures or
retries, no repeated old configurations. This is local Docker execution, not a
new GitHub CI result. Compiler #53's CI is prior compiler provenance only.

## Actual observations

The unchanged24-case reference corpus matched generated Go, native PHP and
explicit expectations: actual bool results, list flags and actual typed ordered
caller inputs after every call. Sparse/negative/out-of-order integer keys remain
distinct from lists, and canonical decimal-string keys normalize while
01/+1/-0/empty-string keys remain strings. Seven scalar inputs return false.

Seven new direct-polyfill cases matched value/key argument correspondence,
empty=true, first/middle false termination, all-true iteration and unchanged
input after callback-local assignments. Nonnumeric poison values after the
stopping point were not evaluated: evaluation would fail numeric addition.
Both sides exited0 with empty stderr. Every result was read from actual generated
function calls; expectations were not used to construct the Go observations.

One compiler build and one generated binary serve all31 cases. Five owned driver
functions exercise the original polyfill without replacing its implementation.
The generated-status report's source/output hashes are checked before building.
Full actual observations, source/fragment/driver/corpus/pin/compiler/archive/
generated/report/binary/observer hashes are in the [raw record](numeric-array-helper-2026-10-25.json).

## Remaining boundary

Flat null/bool/int/string values, integer range±1,000,000, no nested values,
objects, floats, references or duplicate normalized keys. General callback
side effects, error/exception compatibility, conditional loading, WordPress
bootstrap/full core/HTTP/DB/themes/plugins are not established. A subsequent
expanded single-binary integration needs its own pin/corpus/evidence; do not
add this unit's case count to the old149 and call that an integrated PASS.

WordPress originals, all previous pins/tests/evidence, SDK/hooks/oracle policy,
prototype and master unchanged. No main merge/deploy/release/license decision.
Local `gossip.dev.wordpress-unit.translate-compare` stays WebMCP-excluded;
no browser/API operation or auth-route change. Devboard unidentified/unsynced.
