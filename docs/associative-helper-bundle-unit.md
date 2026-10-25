# Fourteen-original associative helper bundle

This independently pinned configuration translates fourteen unchanged WordPress
helpers and sixteen owned composition drivers together, builds one Go binary,
and compares 55 actual calls against PHP8.3.36 and explicit expectations.
It combines the thirteen-original recursive bundle with `wp_array_slice_assoc`,
including its original `: array` declaration. No original function is rewritten.

## Source, compiler and corpus

The [manifest](../gossip/tests/gohp/associative-helper-bundle.json) binds the two
original files by whole-file hash and selected line ranges. Each selected
function is copied byte-for-byte; names must be unique and all fourteen must be
present. The slice excerpt additionally has its exact known hash checked.
Owned drivers only call original helpers or supported builtins; they do not
substitute for the original selection, escaping or normalization algorithms.

Compiler/runtime pin is `db3051f8a179078e2227c5d97f50a2f1611d1451` (GoHP #36),
unchanged from the standalone slice unit. Its existing CI238-check result is
reused with exact revision/configuration metadata, not rerun or relabelled as a
new compiler verification. The new combined executable gets its own measured
source/output/report/binary hashes and observations.

| Corpus part | Calls | What is exercised |
| --- | ---: | --- |
| Previous recursive bundle | 33 | Same thirteen direct and twenty composed inputs/expectations under this new configuration |
| Original slice helper | 14 | All standalone slice inputs, ordered keys, null exclusion, collisions and caller input/selectors |
| New compositions | 8 | Slice then slash, slash then slice, shallow roundtrip, sparse shape and array-to-Boolean |

Total: 27 direct calls and 28 composed calls. New composition drivers are
`gossip_slice_slash`, `gossip_slash_slice`, `gossip_slice_shallow_roundtrip`,
`gossip_slice_shape` and `gossip_slice_boolean`. Together with the eleven prior
drivers, all use the same generated originals in a single process/binary.
The nested shallow-roundtrip case deliberately retains escaped nested strings;
it is not `wp_unslash` or recursive unescaping support. Shape observation checks
sparse/reordered keys separately from reindexed values. A nonempty slice whose
only value is false remains truthy as an array.

## Transport and observation boundary

Operations are explicitly whitelisted, not inferred from DOM, reflection or
arbitrary callback names. The six slice-related operations take an ordered
`[key,value]` pair list as argument zero and a selector list as argument one.
Only this argument is constructed as an associative PHP/GoHP array. Other
arguments retain the prior scalar/sequential-list transport.

Direct slice and the three escaping compositions serialize actual root array
entries back to ordered pairs, preserving integer versus string keys. Shape
and Boolean results use ordinary list/scalar observation. Actual post-call
arguments use the same operation-specific transport; explicit `args_after`
expectations describe the constructed input (including PHP key normalization),
not the raw pair list with duplicate keys. Expected values never construct an
observed result. The input-encoding rule does not implement the slice algorithm.

JSON values are null/bool/modest signed integer/string/finite lists. Numbers are
checked as signed integers within ±1,000,000, and the host validates depth≤8.
Root associative keys are int/string only; nested values remain sequential
lists. JSON objects, floats, nested associative values, references, cycles and
arbitrary integer precision are excluded. Root return type and actual values
and post-call arguments must match both native PHP and authored expectations.

## Execution and separation

```sh
cd gossip
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressAssociativeHelperBundle$'
```

No opt-in is SKIP. Approved immutable images, same-process PHP version guard,
PHP LC_CTYPE=C, non-root/read-only/network-disabled/resource-limited execution
remain in force. Source/generated report hashes are verified before build.
Generation or compilation alone is not compatibility PASS; the comparison
requires exact observations, empty stderr and exit0.

The [execution record](evidence/associative-helper-bundle-2026-10-25.md) binds
this configuration. Previous eight standalone units and ten/twelve/thirteen-
original bundles retain their pins/corpora/evidence. Prior PASS is not evidence
for a new configuration, and this corpus is not all prior standalone inputs.

Whole-file/core startup, hooks, HTTP/DB, themes/plugins, general object/reference/
callback behavior, map_deep/wp_unslash and PHP exception equivalence remain
unsupported or unverified. No deployment, release, implementation merge or
license decision is introduced.

`gossip.dev.wordpress-unit.translate-compare` remains WebMCP-excluded local
opt-in developer tooling. No service/browser action, auth route, cookie
forwarding, automatic submission or production audit change. Devboard card
unidentified/unsynchronized; no broader Done claim.
