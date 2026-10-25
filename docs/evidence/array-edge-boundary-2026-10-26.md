# Array edge native reference / rejected compiler boundary — 2026-10-26

**Native reference PASS24/24; original Go translations REJECTED.**
The boundary test and vet passed. There is no generated candidate binary or
Go/PHP differential result. This does not increase the22-original/230-case count.

## Snapshot and execution

- Branch `codex/gohp-array-edge-boundary`, based on Draft #29 head
  `95034a00cd2877cd029d101e66467aed23fdeb4a`.
- Test commit `6edeffa6cf568004babcdc59622e9fbd9662ab49`, tree
  `9a820ac45f284e679f8249338ef06dea7d76cf41`.
  Staged diff check and clean commit preceded execution. Only docs/evidence afterward.
- GoHP #54 `36191c57003408d001d9f7217805a3549fdd9f98`,
  tree `edbf8b45b3ca3c60c08e144b7725dfc24ae9ae08`; exact git archive.
  Compiler source/pin unchanged; no compiler tests or CI dispatch.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`.
  Original array_first657–665 / array_last679–685 declarations unchanged.
  Native executes full compat.php guards646–666 /668–686 and verifies both
  functions are user-defined. Original annotations/indentation/bodies retained.
- Actual PHP8.3.36, C locale; fixed builder Go1.25.14 linux/arm64.
  Same-process admission/non-root/read-only/network-off/resource limits retained.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArrayEdgeBoundary$'
go vet ./tests/gohp
```

First boundary-test PASS10.85s/package11.099s, vetPASS, command exit0.
No unexpected failures/retries. Expected compiler exit1 is recorded separately.
No repeated old configuration, post-PASS rerun or new GitHub CI.

## Native behavior

Twenty-four actual return-type/value and ordered caller input-after observations
match the authored expectations (twelve per helper). Empty arrays return null;
singleton null/false/0/'0'/empty-string arrays return their actual typed values.
First/last mean insertion order, not numeric-key sorting. Sparse/negative/bounded
and canonical numeric-string keys, mixed keys and Unicode/NUL keys/values retain
their observed order and bytes. Caller inputs remain unchanged.

The native observer reads actual original-function returns, never substitutes
expectations. Shared bounded transport is extracted unchanged from the existing
numeric-array reference; source/selected hashes are recorded. Native stderr is
empty and native exit0.

## Actual compiler diagnostics

Each original declaration is probed unchanged, with the PHP opening on line1.
Both were parsed but rejected during code generation at the first empty call:

| Probe | Role | Actual diagnostic |
| --- | --- | --- |
| array_first | Original declaration | PHP004 line 3: unsupported function empty |
| array_last | Original declaration | PHP004 line 3: unsupported function empty |
| gossip_probe_array_key_last | Owned isolated dependency probe | PHP004 line 2: unsupported function array_key_last |

All three CLI executions returned exit1 with empty stdout and no generated Go
or report file. The third probe is not an original WordPress function or a
replacement implementation; it isolates the dependency hidden behind the first
array_last blocker. Shell/container failures are not accepted as these rejections.

The [raw record](array-edge-boundary-2026-10-26.json) contains exact diagnostics,
all24 native observations and source/fragment/corpus/observer/transport/manifest/
compiler/archive hashes. Its status is native_reference_pass_candidate_rejected,
and differential_status is not-run. No compiler output/runtime equivalence claimed.

## Next implementation and limits

First implement bounded empty on initialized actual-array variables/parameters,
with its own compiler/runtime evidence; then compare unchanged array_first under
a new pin. Separately add array_key_last's insertion-order key/null contract and
safe array indexing before comparing unchanged array_last. General undefined/
index/property empty semantics must not be accidentally claimed.

Only flat scalar values and bounded integer keys/values were referenced.
Nested arrays/objects/floats/references/normalized duplicate keys, general callable/
capture/error, bootstrap/full core/HTTP/DB/theme/plugin compatibility remain
outside scope. Previous success pins/evidence remain unchanged.
WordPress originals, SDK/hooks/oracle policy, prototype/master unchanged;
no merge/deploy/release/licensing decision. Local capability remains WebMCP-excluded.
Devboard unidentified/unsynchronized, not Done.
