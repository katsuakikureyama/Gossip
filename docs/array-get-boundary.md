# Nested settings lookup: original _wp_array_get

The next application-oriented target is `_wp_array_get`, which retrieves a
nested value from an ordered key path. The fixed WordPress source uses it in
Theme JSON (43 callsites), global styles/settings (6) and block support lookup
(blocks.php:2745), plus ten other callsites. This survey excludes its declaration
and doc example:60 actual calls across12 files. It does not execute those callers.

Examples at the fixed source revision:

- class-wp-theme-json.php:1023 passes a settings path and omits the third argument.
- global-styles-and-settings.php:44 passes the full settings array as fallback.
- blocks.php:2745 passes the caller's default value for a block-support path.

## Original and native reference

The manifest pins the WordPress baseline, whole functions.php hash and the exact
declaration5233–5271, including `$default_value = null`. Native PHP loads those
bytes without rewriting. Reflection verifies user-defined origin.
The compiler pin is GoHP #59 main `c3420d8644642c2fee56d62a08c88a2dcedc45ac`.

Twenty-four authored cases cover three-level lookup, array returns and defaults,
missing/scalar intermediates, present-null versus missing, false/0/'0'/empty
string, empty/non-array paths, non-array input, integer/canonical numeric-string/
leading-zero/empty/NUL/Unicode keys, rejected bool/null/array path elements,
ordered sparse/associative paths, and omitted defaults.

The owned transport encodes arrays as `{"array":[[key,node],...]}` so nested
integer/string keys and insertion order remain explicit. Scalars are null/bool/
integer/string. The observer rejects duplicate normalized keys, unsupported
types and values outside depth12,256 entries,4096-byte strings and integers
±1,000,000. It uses PHP's own array assignment for key normalization.
Expected nodes must be canonical. Actual return type/value and input/path/
default argument states are recorded; each argument is checked unchanged.

## Compiler work required

The original and four isolated owned probes identify these boundaries separately:

1. A trailing parameter default and argument count range. Both original2-arg
   calls (default null) and explicit mixed fallback values need preservation.
2. Bare-variable and multi-operand isset. The original's first operand must
   short-circuit the indexed second operand, and null must stay distinct from
   a missing key. Current indexed-only isset does not cover these forms.
3. Continue inside foreach, including through nested if bodies. General level
   expressions, other loops and arbitrary control flow are separate scope.

The probe inputs are independent owned functions. The original is never patched
to remove its default, split isset calls or replace continue. Every rejection
requires the pinned diagnostic, exit1, empty stdout and no generated/report
artifacts. This is a native reference plus explicit parser blockers, not a new
Go-compatible original or an extension of24-original262-case integration.

From the nested `gossip/` module:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArrayGetBoundary$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

The existing PHP8.3.36 same-process admission, C locale, non-root/read-only/
network-off execution and resource limits remain. Record actual measurements
before marking the reference PASS; Go/PHP comparison remains not-run until a
new compiler supports the unchanged original. Full Theme JSON, blocks, database
and rendering are not covered by testing this helper.

Capability `gossip.dev.wordpress-unit.translate-compare` remains explicitly
WebMCP-excluded as a local developer fixture; no browser/server operation or
authentication route is added. PR/current-state tracks progress; devboard card
is not identified.
