# Original-source associative array slice

This eighth independent translation unit extracts `wp_array_slice_assoc()`
unchanged from `wp-includes/functions.php` lines 5173–5183, including the original
`: array` return declaration. WordPress baseline is
`54106f6c545fc8b881f6b39b7963002bece6c28c`. The original selection loop, isset
probe, indexed write and return execute as generated Go; observers contain
transport and observation code, not a replacement selection algorithm.

## Contract and observation

The fourteen owned cases transport root associative arrays as ordered
`[key, value]` pairs, so JSON object-key coercion cannot hide integer/string
key differences. Selectors are sequential lists of int/string values. Values
are null, bool, modest signed integers, strings or nested sequential lists.
Integers are bounded to ±1,000,000 and value depth to eight. JSON objects,
floats, nested associative values, references and cycles are excluded.

Actual root return type, ordered key/value entries, caller input entries and
selector array after invocation must equal native PHP and independently authored
expectations. Missing/null entries are omitted; false/zero/empty strings/empty
arrays remain. Request order and first position of duplicated selectors are
preserved. Cases also cover integer/numeric-string normalization, noncanonical
numeric keys, collisions overwritten with null, negative keys, UTF-8/NUL/empty
keys, nested list values and sparse output from list input.

Input construction itself follows PHP key normalization and replacement. The
explicit `input_after` expectation describes that constructed array, not the
raw pair list with duplicate keys. Observers do not normalize result pairs by
reconstructing the helper algorithm. JSON scalars distinguish strings, booleans,
null and bounded integers; nested array keys must remain sequential.

The compiler pin adds global array-return declarations with checks at explicit
returns and fallthrough; unknown non-arrays fail rather than erasing the type.
It inherits the bounded one-level isset and indexed assignment contracts. This
does not add general isset, typed parameters/methods, arbitrary return types,
TypeError equivalence or nested writes. See the pinned
[GoHP return contract](https://github.com/katsuakikureyama/GoHP/blob/db3051f8a179078e2227c5d97f50a2f1611d1451/docs/array-return-boundary.md).

## Execution and separation

The [manifest](../gossip/tests/gohp/array-slice-assoc-helper.json) pins whole-file
hash, exact excerpt, committed compiler/runtime and approved immutable images.
Generation reports are hash-checked before build. Native PHP version guard,
non-root/read-only/no-network/resource-limited execution remain unchanged.
PASS additionally requires empty stderr, exit0 and exact comparisons.

```sh
cd gossip
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArraySliceAssocHelper$'
```

Missing opt-in is SKIP. [Execution evidence](evidence/array-slice-assoc-2026-10-25.md)
records the measured revision/configuration. Earlier seven standalone units and
ten/twelve/thirteen-original bundles retain their existing pins/corpora/results;
this is not a fourteen-function combined test or a rerun of old inputs.

Whole-file/core startup, hooks, HTTP/DB, themes/plugins, map_deep/wp_unslash,
general object/reference/callback behavior remain unsupported or unverified.

## Capability exclusion

`gossip.dev.wordpress-unit.translate-compare` remains WebMCP-excluded local,
opt-in developer tooling. No UI/service action, authentication route, browser
cookie forwarding, external save/commit, automatic submit or production audit
change. No implementation merge, release, deployment or license decision.
Devboard card is unidentified/unsynchronized, not Done.
