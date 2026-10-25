# Array edge reference and compiler boundary

This separate configuration pins the same GoHP #54 revision as the existing
twenty-two-original bundle: `36191c57003408d001d9f7217805a3549fdd9f98`.
It establishes native PHP behavior and explicit translation blockers for
`array_first` and `array_last`, not an additional successful Go unit.
The earlier22-original/230-case integration remains unchanged.

The [execution record](evidence/array-edge-boundary-2026-10-26.md) reports
24/24 native matches and the three expected code-generation rejections.
This is not Go/PHP differential PASS.

## Exact original extraction and reference

WordPress baseline and whole compat.php SHA-256 are checked before extraction.
Original declarations are array_first657–665 and array_last679–685.
The PHP8.3 reference loads complete guarded blocks646–666 and668–686 and verifies
both functions are user-defined. Parameter types, function bodies, comments and
indentation are unchanged. The compiler probes use declarations only; excluding
top-level loading guards is an explicit boundary, not conditional-loading support.

Twenty-four authored cases (twelve per helper) measure actual return type/value
and ordered caller input after calls. Empty arrays are distinguished from arrays
containing null/false/0/'0'/empty string. Selection follows insertion order, not
numeric-key sorting. Sparse/negative/boundary keys, canonical numeric-string keys,
mixed string/integer keys and Unicode/NUL keys/values are included.
Only flat null/bool/int/string values and keys within the declared bounded
transport are used; integer keys/values ±1,000,000. No nested arrays, objects,
floats, references or normalized duplicate keys.

The native observer reuses unchanged transport functions selected by unique
markers from the previous numeric-array observer. Source and selected hashes
are recorded. Expected values never supply observed results.

## Compiler probes and next work

Two byte-unchanged original declaration probes are expected to reject at
`empty($array)`. A separate owned declaration calls `array_key_last($items)`
to expose the second dependency without rewriting array_last or treating a stub
as original-source compatibility. The roles are recorded separately.

The test requires exact CLI diagnostic, exit1, empty stdout and absence of
generated Go/report for each probe. Docker/shell failures do not count as expected
rejections. A passing test means native_reference_pass_candidate_rejected;
Go/PHP differential status remains not-run.

Next implementation order:

1. Add a bounded `empty` language-construct contract in the compiler, first for
   initialized actual-array variables/array parameters. It must not accidentally
   claim general undefined/index/property suppression or arbitrary PHP expressions.
   Prove its own positive/negative runtime and compiler behavior before changing pins.
2. Compare unchanged array_first against this reference in a new configuration.
3. Add `array_key_last` for actual arrays: insertion-order last int/string key,
   null for empty, no coercion/mutation, correct key use at indexing. It is separate
   from array values or key sorting. Then compare unchanged array_last separately.

Current is_int-only compiler work is not a substitute for these dependencies.
This boundary does not modify the compiler or adopt another revision's PASS.
General capture/filter/object-based callers remain out of scope.

## Reproduce and exclusions

From `gossip/`, non-root with cached immutable Docker images:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArrayEdgeBoundary$'
go vet ./tests/gohp
```

Absent opt-in is SKIP, not PASS. Same-process PHP8.3 admission, minimum8.3.36,
C locale, non-root/read-only/network-off and resource limits are unchanged.
No old PASS configuration or compiler test suite is rerun.

Capability `gossip.dev.wordpress-unit.translate-compare` stays **WebMCP-excluded**:
local developer reference/translation probe, not browser/API/authentication-route
changes or auto-submit. No WordPress originals, SDK/hooks/oracle policy,
prototype/master changes, merge/deploy/release/licensing or full-core/HTTP/DB/
theme/plugin compatibility claims. Devboard unidentified/unsynchronized, not Done.
