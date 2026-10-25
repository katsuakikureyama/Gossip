# Fifteen-original charset/array bundle

This independently pinned configuration integrates fifteen byte-unchanged
WordPress helpers and nineteen owned composition drivers in one translation,
one build and one generated executable. It combines the fourteen-original
associative bundle with `_is_utf8_charset` from compat.php102–111.
The [manifest](../gossip/tests/gohp/charset-helper-bundle.json) binds three
whole-file hashes, exact source ranges and the merged GoHP compiler pin.
Slice and charset excerpts additionally have their exact hashes checked.

GoHP combines published #40 strcasecmp/logical-comparison and #36 isset/
array-return stacks. Both contracts are present in the new configuration;
older compiler pins are not modified, and parent PASS does not substitute
for fresh merged-compiler and Gossip integration evidence.

## Corpus

| Part | Calls | Boundary |
| --- | ---: | --- |
| Previous associative bundle | 55 | Same cases, arguments and expectations under the newly merged compiler |
| Direct charset predicate | 28 | All standalone charset-name cases, including non-string rejection |
| New compositions | 6 | Slice/classify, slice/shallow-unescape/classify, slash/classify/Boolean |

Total89 calls: direct55/composed34. The new drivers are `gossip_slice_charset`,
`gossip_slice_unescape_charset` and `gossip_charset_boolean`. They call original
helpers and admitted builtins, never replace original bodies. Mapping preserves
selected root keys/order; null/missing values are excluded by original slice,
while false and nested lists reach the charset predicate and return false.
Shallow stripping does not recursively decode nested lists. The two scalar
Boolean compositions use string inputs only. Charset recognition is not UTF-8
text validation or encoding conversion.

Both observers use explicit operation dispatch and actual post-call arguments.
The two new slice operations use the existing ordered `[key,value]` root-array
transport plus selector list and serialize actual results back to ordered
pairs; the new direct predicate retains mixed scalar/sequential-list inputs.
Expected fields never construct observed results. The existing bounded
integer(±1,000,000)/depth≤8/int-or-string-key rules remain; nested values are
sequential lists, not nested associative arrays or JSON objects. Floats,
references, cycles and arbitrary precision/depth remain excluded.

## Execution

```sh
cd gossip
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressCharsetHelperBundle$'
```

Without opt-in the test skips. Immutable images, same-process PHP8.3.36 guard,
LC_CTYPE=C and non-root/read-only/network-disabled/resource-limited execution
remain unchanged. Source/generated/report hashes are checked before build.
PASS requires actual return types/values and caller arguments to match native
PHP and authored expectations, plus empty stderr and exit0. The
[execution record](evidence/charset-helper-bundle-2026-10-25.md) binds the
new configuration; earlier nine independent units and ten/twelve/thirteen/
fourteen-original bundles retain their separate pins/corpora/results.

This is not all historic standalone inputs. Public is_utf8_charset/defaults/
coalescing/get_option, _canonical_charset, nested associative transport,
general objects/references/callbacks, map_deep/wp_unslash, whole-file/core,
HTTP/DB/theme/plugin operation and PHP exception equivalence remain unverified.
No implementation merge to master, deployment, release or license decision.
Local capability `gossip.dev.wordpress-unit.translate-compare` remains
WebMCP-excluded; authentication routes unchanged. Devboard unidentified and
unsynchronized, not Done.
