# Twenty-five originals /294 cases: integrated PASS

2026-10-26 JST. [Contract](../settings-helper-bundle-unit.md),
[all294 actual observations and hashes](settings-helper-bundle-2026-10-26.json).

## Result and exact snapshot

**25 unchanged original declarations and48 owned drivers translated/built once
into one executable. All294 cases match fixed PHP8.3.36/C and expectations.**
Actual return types/values and ordered caller argument states are observed;
this is a new integrated run, not the sum of previous unit results.

- Branch `codex/gohp-settings-helper-bundle`, base master78b580a.
- Clean test commit `24c780b361dd4d11f4e18b79cc90d80f5d267872`, tree
  `56b9e3e1cf2938f615d0a4717b3532c31e969d43`.
- GoHP #64 main `b8dc1757ed61f459d51362da004aa1d73d024116`.
  Same tree as feature18c4e67; fixed-PHP
  [CI40227365611](https://github.com/katsuakikureyama/GoHP/actions/runs/40227365611)
  and [main CI40227482630](https://github.com/katsuakikureyama/GoHP/actions/runs/40227482630)
  PASS. Existing473-check compiler result reused, no compiler changes/retests.
- One manifest pins all three whole-source files and31 selected fragments
  (25 declarations, six full native guards). Original source is byte-unchanged.
- Fixed digest builder Go1.25.14 linux/arm64; PHP8.3.36/C. Existing nonroot,
  read-only, network-off/resource limits and same-process admission retained.

From the nested `gossip/` module:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-null-identity GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressSettingsHelperBundle$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

Comparison **PASS17.62s / package17.816s**, exit0 and empty stderr on both
executions. Vet PASS on the same snapshot. Raw log
`/private/tmp/gossip-settings-bundle-output.log`. The first committed snapshot
passed; only documentation/evidence changes follow. No repeated successful run.

## Retention and new compositions

The294 cases retain all262 edge-bundle observations and all24 standalone lookup
cases, plus8 new settings compositions. Nine old corpora are hash-locked and
unchanged. The lookup's retained expectation record is bound to the same source,
fragment and corpus hashes. Both new runtime executions still invoke all24
lookup calls; the retained record is an assertion, not a substitute observation.

Exact-once adapters turn old Go init observers into return functions and keep
their measured results. PHP wrappers capture independent observer output without
redeclaring helpers. Independent corpus paths remain distinct. Source/adapted
hashes, generated/report/compiler/binary/archive hashes and all observations are
recorded in JSON, including all48 driver declarations across seven owned files.

Four new cases compose lookup → Boolean validation: nested true, `'FALSE'`,
missing with `'0'` fallback and present-null despite true fallback. Four compose
lookup → associative slice → first/last/numeric-key predicate/all predicate:
missing uses `[]`; slice omits null; selector order/duplicates/mixed keys and
numeric-string selectors preserve intended insertion order. The first/last
results retain raw strings while the all predicate converts them. A numeric-key
array need not be a PHP list. The two initially authored expectations that
incorrectly converted first/last to bool were corrected in parent review before
the test snapshot. No runtime failure or code workaround was hidden.

All four new fixture argument states are measured and checked unchanged. The
parent owns design, harness/adapters, review/verification and integration; workers
supplied only the new driver/case and Go/PHP wrapper files. Old settings/pins/
corpora/evidence and WordPress source were not rewritten.

## Exclusions

This does not execute Theme JSON/global styles/block callers or start WordPress.
HTTP/DB/themes/plugins, general object/reference/float/error/coercion semantics
and production operation remain unverified. The new edge driver's retrieved
value is an array or a missing path using `[]`, not a present-null/scalar value.
Legacy fixture domains are not widened by the nested transport.

`gossip.dev.wordpress-unit.translate-compare` remains explicitly WebMCP-excluded,
an opt-in local developer fixture with no browser/server operation or route/auth
change. No unconfigured deployment target is created. Devboard unidentified/not
synced; PR/current-state tracks progress.
