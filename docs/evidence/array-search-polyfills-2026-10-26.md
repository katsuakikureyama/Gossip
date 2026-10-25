# Original array search polyfills — 2026-10-26

**PASS: 30/30 cases against actual PHP8.3.36; vet PASS.**
Three original declarations execute in one generated binary. This independent
unit is not an expanded twenty-two-original integration; the earlier
nineteen-original/192-case bundle and evidence remain unchanged.

## Exact snapshot

- Branch `codex/gohp-array-search-polyfills`, based on Draft #27 head
  `baaac5474d751fc38fa55875ab862ffdb9564585`.
- Test commit `861814fcfecdc0eba0aca2e9f61ed6be1677bc03`, tree
  `ff3e64f9e016944c82c80d9a63a5e10b24788fc4`.
  Staged diff check and clean commit preceded execution.
  Only documentation/evidence changed afterward.
- GoHP #54 `36191c57003408d001d9f7217805a3549fdd9f98`,
  tree `edbf8b45b3ca3c60c08e144b7725dfc24ae9ae08`.
  Exact git archive, no compiler changes/retests/CI dispatch.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`.
  Original compat.php declarations: array_find566–574, array_find_key589–597,
  array_any612–620. Native guards:554–575,577–598,600–621.
  Source bytes/annotations/bodies unchanged. Native verifies user-defined origin.
- Fixed builder Go1.25.14 linux/arm64; PHP8.3.36; C locale.
  Same-process admission and non-root/read-only/network-off/resource limits
  remain unchanged.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArraySearchPolyfills$'
go vet ./tests/gohp
```

First run PASS9.23s/package9.568s, vet PASS, combined command exit0.
No failures/retries or repeated old configurations. This is local Docker
execution, not a new GitHub CI result.

## Observed behavior

Nine owned drivers exercise each original through strict value/key equality,
early-success arithmetic probes and scalar truthiness. Ten cases per original
cover empty/no-match, first/middle match, canonical numeric-string keys,
sparse/bounded integer keys, string keys and input preservation.

Actual outputs distinguish Boolean any results, find values including integer0
versus string1, and integer/string/null key results. Falsy null/false/0/'0'/empty
strings do not satisfy truth callbacks; the nonempty word 'false' does.
Later nonnumeric poison in the stop cases is skipped: evaluating it would fail
addition. Ordinary equality/truth cases containing the same word do not by
themselves prove short circuit; only the arithmetic probes do.

Go and native PHP both match authored expectations for actual typed scalar
returns and ordered caller arrays after calls, with empty stderr and exit0.
Expected fields never supply generated-runtime observations.
One compiler translation/build and one binary serve all30 cases.

The [raw record](array-search-polyfills-2026-10-26.json) contains all observations
and exact source/fragment/manifest/driver/corpus/compiler/archive/generated/
report/binary/observer/transport hashes. Shared transports are unchanged slices
of earlier observer files; old test dispatch and generated calls are excluded.

## Boundaries

Flat scalar values only, integer keys/values ±1,000,000; no normalized duplicate
keys, nested arrays, objects, floats or references. Matched-null and matched-false
value-return cases are not established by this corpus. No general callable,
capture, error/exception or conditional-loading compatibility is claimed.
wp_ext2type/filter/object-based callers are not included.

WordPress originals, old pins/corpora/evidence, SDK/hooks, PHP oracle policy,
prototype and master unchanged. No merge/deployment/release/licensing decision.
Capability `gossip.dev.wordpress-unit.translate-compare` remains WebMCP-excluded,
with no browser/API operation or authentication-route change.
Devboard unidentified/unsynchronized, not Done.
