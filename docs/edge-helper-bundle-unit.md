# Twenty-four-original integrated helper bundle

This configuration translates24 unchanged WordPress declarations and46 owned
composition drivers once, builds one executable, and compares262 cases with
PHP8.3.36. It selects GoHP #58 main
`440308c77da07cbcdce403fed7fff3945e0d2078`, the compiler already proven by the
two-original edge unit. The subsequent #59 feature is not needed for this
configuration. Its previous pins and evidence stay separate.

The [execution record](evidence/edge-helper-bundle-2026-10-26.md) reports
262/262 typed/caller-state comparisons and vet PASS on the exact configuration.

The manifest `gossip/tests/gohp/edge-helper-bundle.json` retains the22-original
search bundle's exact source fragments and adds array_first657–665 and
array_last679–685 from compat.php. Source files and fragments are hash-checked;
declarations are copied byte-for-byte, including annotations and comments.
Native PHP executes all six complete guarded array polyfills, and the observers
verify user-defined origin. The compiler consumes declarations only.

## Cases and observations

The single executable runs all230 prior search-bundle cases, all24 edge cases
and8 new composed cases. The seven retained corpus hashes are fixed in the new
harness. Prior corpora, manifests, observers and evidence are unchanged.

The new drivers select keys through wp_array_slice_assoc, map either bool_from_yn
or wp_validate_boolean, then observe array_first, array_last and the three
search polyfills together. Eight cases distinguish:

- empty selections from singleton false;
- first and last values after selector ordering and deduplication;
- missing/null exclusion, numeric-string normalization and mixed keys;
- a matched false and key0 from a no-match null;
- all-true/no-match inputs and bounded negative/positive keys.

Actual return types/values and ordered caller arguments are compared against
PHP and authored expectations. Expected fields never construct generated results.
New compositions use flat string/null inputs, int/string selectors and integers
within ±1,000,000. Retained cases keep their own original domains.

Observers from the prior units are adapted in memory: init/exit becomes a
return, obsolete imports are removed, PHP uses require_once/shared transport,
and case paths are separated. Every replacement must match exactly once;
original/adapted hashes are recorded. The new outer observer collects all
results in one process. These adapters do not modify originals or generated
function bodies.

## Reproduction and limits

From nested `gossip/`, non-root with cached fixed Docker images:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressEdgeHelperBundle$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

Absent compiler-source opt-in means SKIP. Record the exact committed snapshot
and actual run before declaring success. The harness retains same-process
PHP8.3.36 admission, C locale, network-off/read-only/non-root execution and
resource limits. No compiler changes or duplicate historical runs are needed.

This is declaration-level integration. WordPress startup/HTTP/DB/themes/plugins,
general callbacks/captures/error equivalence and broader data domains remain
future work. Capability `gossip.dev.wordpress-unit.translate-compare` remains
WebMCP-excluded as a local developer operation; no route/auth change occurs.
Devboard card is unidentified; the PR and current-state track progress.
