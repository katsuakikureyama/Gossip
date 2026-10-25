# Nineteen-original numeric helper integration — 2026-10-25

**PASS: all 192 cases matched actual PHP 8.3.36 in one generated binary.**
This is new integrated evidence, not a sum of previous independent PASS counts.
The old seventeen-original/149-case integration, two-original/31-case unit
and historical rejection evidence remain unchanged.

## Exact snapshot and execution

- Gossip branch `codex/gohp-numeric-array-bundle`, based on #26
  `6ad9b8d0cb6a22ff5855e59787990b21a30e52a7`.
- Test commit `9972a6f92ad3d3e889d8d71b128d893966d65682`, tree
  `38e1796f9f391bc28440e8abf6c9c3149649275a`. Staged diff check and clean
  commit preceded execution. Only documentation/evidence changed afterward.
- GoHP #54 `36191c57003408d001d9f7217805a3549fdd9f98`, tree
  `edbf8b45b3ca3c60c08e144b7725dfc24ae9ae08`. Exported exact git archive;
  no compiler edits, compiler test reruns or CI dispatch.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`.
  All nineteen declarations were extracted byte-for-byte from three hash-pinned
  source files. PHP executes the entire original guarded array_all block;
  compiler input extracts only its unchanged declaration, not its loading guard.
- Host Go 1.22.1 darwin/arm64; pinned builder Go 1.25.14 linux/arm64; actual
  PHP 8.3.36 with C locale. Non-root, read-only, network-disabled and resource
  restrictions plus same-process PHP admission policy remain unchanged.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressNumericHelperBundle$'
go vet ./tests/gohp
```

First run **PASS 13.63s / package 14.001s**, vet PASS, combined command exit0.
No failures or retries. No repeated old configurations or post-PASS rerun.
This is local Docker execution, not a new GitHub CI result.

## What was observed

One compiler build/translation and one generated binary run all 192 cases:
149 unchanged legacy cases, 24 unchanged numeric predicate cases, seven unchanged
direct-polyfill cases and twelve new compositions. The nineteen originals use
thirty-three owned driver functions; no original function body is replaced.

The new compositions combine associative selection, numeric/list/Boolean
classification, wp_slash, array_values, bool_from_yn and original array_all.
Empty selections, sparse/negative and string keys, selector order/duplicates/
missing entries, canonical decimal key normalization and Y/N/false labels match
authored expectations. Observers read actual Boolean-list returns and actual
ordered caller inputs/selectors after calls. The retained corpora also preserve
their prior typed, nested-list, scalar, key and short-circuit observations.

The old observer files remain unchanged. Exact-once adapters turn the two Go
init/exit observers into functions and isolate the legacy corpus path; PHP
adapters use require_once and that path. The raw evidence records original and
adapted hashes. Both languages observe the same complete original definition
bundle in one process. Generated source/report hashes match; stderr is empty
and runtime exit status is0.

Full actual observations and source/fragment/manifest/corpus/driver/adapter/
compiler/archive/generated/report/binary hashes are in the
[raw execution record](numeric-helper-bundle-2026-10-25.json).

## Limits and unchanged areas

Only these bounded declaration-level domains are established. New compositions
use flat string values and bounded keys. No general callback, error/exception,
conditional-declaration, full-file/bootstrap, full-core, HTTP, database, theme
or plugin compatibility is claimed. MIME tests classify labels, not image bytes.
Existing legacy nested-list coverage does not establish nested associative support.

WordPress originals, previous pins/corpora/evidence, SDK/hooks, PHP admission,
prototype and master are unchanged. No merge, deployment, release or licensing
decision. Local `gossip.dev.wordpress-unit.translate-compare` remains
WebMCP-excluded; no browser/API operation or authentication-route change.
Devboard is unidentified and unsynchronized, not Done.
