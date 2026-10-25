# Twenty-two-original helper integration — 2026-10-26

**PASS: 230/230 cases against actual PHP8.3.36; vet PASS.**
Twenty-two unchanged original declarations execute together in one generated
binary. This is a new integrated run, not a sum of prior independent PASS counts.
The nineteen-original192-case and three-original30-case evidence remain unchanged.

## Exact snapshot and execution

- Branch `codex/gohp-search-helper-bundle`, based on Draft #28 head
  `cdbc55b8eb80363935dfca5a7d519066bfaf138c`.
- Test commit `97b264cfb7bf17d61bfec31c4103ecb9d67b17b8`, tree
  `037353ebeee9884f336fe9b8f151eba39bad7054`.
  Staged diff check and clean commit preceded execution.
  Only documentation/evidence changed afterward.
- GoHP #54 `36191c57003408d001d9f7217805a3549fdd9f98`,
  tree `edbf8b45b3ca3c60c08e144b7725dfc24ae9ae08`.
  Exact git archive; no compiler changes/retests/CI dispatch.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`.
  Exact source and selected declaration bytes are checked.
  Native PHP executes complete original guards for the four array polyfills
  and verifies user-defined provenance. The compiler extracts only unchanged
  declarations; top-level conditional loading/WordPress bootstrap is excluded.
- Fixed builder Go1.25.14 linux/arm64; actualPHP8.3.36; C locale.
  Non-root/read-only/network-off/resource restrictions and same-process PHP
  admission are unchanged.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressSearchHelperBundle$'
go vet ./tests/gohp
```

First run PASS20.52s/package20.766s, vet PASS, combined command exit0.
No failures/retries, no post-PASS rerun or repeated old configurations.
This is local Docker evidence, not a new GitHub CI result.

## Actual observations

One translation/build and one generated binary execute 230 cases:
149 legacy +24 numeric predicate +7 direct array_all +12 numeric compositions
+30 search-polyfill +8 new search compositions. All five retained corpora have
fixed hashes. Forty-four owned driver declarations exercise twenty-two originals.

New compositions select ordered entries, map bool_from_yn/wp_validate_boolean,
then call array_any/array_find/array_find_key; Y/N also calls wp_is_numeric_array.
Matched false is returned as an actual Boolean, not null; no match returns null.
Key0 is an integer distinct from null. Empty/null-only selections, canonical
numeric-string keys, mixed/sparse keys, selector order/dedup/missing entries and
different Y/N versus general Boolean interpretation match explicit expectations.

The observer measures actual scalar-valued lists and caller input/selectors after
execution. Native PHP and authored expectations both match these observations;
expected fields never construct runtime results. Existing retained observations
also execute in this same process, preserving their declared typed/list/nested/
scalar and short-circuit domains. Both runtimes emit empty stderr and exit0.

Existing observers are adapted by exact-once replacements solely for return-based
aggregation, corpus path separation, unused os imports and require_once/shared
transport loading. Original observer files are unchanged. Original/adapted hashes
and every observation, source/fragment/corpus/manifest/driver/compiler/archive/
generated/report/binary hash are in the
[raw execution record](search-helper-bundle-2026-10-26.json).

## Remaining limits

This is bounded declaration-level integration, not full WordPress startup.
New composed input values are flat string/null; keys are bounded int/string.
Matched-null value return, general callbacks/captures/errors, arbitrary nested
associative values, HTTP/DB/theme/plugin compatibility remain unverified.
wp_ext2type/filter callers are outside this unit. MIME cases classify labels,
not image bytes. No production/release/licensing or master-merge claim.

WordPress originals, previous pins/corpora/evidence, SDK/hooks/oracle policy,
prototype/master unchanged. Capability `gossip.dev.wordpress-unit.translate-compare`
remains WebMCP-excluded; no browser/API/authentication-route change.
Devboard unidentified/unsynchronized, not Done.
