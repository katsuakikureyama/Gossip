# Numeric-array original helper and array_all polyfill unit

This is a separate candidate configuration from the historical
[native-reference/rejection boundary](numeric-array-boundary.md). It pins GoHP
PR #53 at `bd419a2c6c72e645f0f0d7d09aa01eea15f70913`; previous pins, corpora and
evidence remain unchanged. Compiler CI is provenance, not proof of this unit.
PR #52's additional composition fixture is not included in this exact compiler
revision; no combined #52/#53 coverage is claimed.

The [execution record](evidence/numeric-array-helper-2026-10-25.md) reports
31/31 Go/PHP matches on the committed configuration, including actual Boolean
type/value, list status and caller input-after. This does not promote old pins.

## Extraction boundary

The original `wp_is_numeric_array` declaration is extracted byte-for-byte from
`wp-includes/functions.php:5448–5454`, and the original `array_all` declaration
from `wp-includes/compat.php:635–643`. The pinned whole-source hashes and baseline
revision are checked before extraction. Parameter/return annotations, arrow
callback, indentation and bodies are not rewritten or replaced by runtime stubs.

The native PHP 8.3 reference executes the complete original guarded block at
`compat.php:623–644`, then checks that `array_all` is user-defined, not a PHP8.4
builtin. The compiler receives only the unchanged declarations plus five owned
driver functions. Omitting the top-level `function_exists` loading wrapper is
an explicit extraction boundary, not support for conditional function declarations,
WordPress bootstrap or loading all of compat.php/functions.php.

## Observations

The previous 24-case `numeric-array-cases.json` and PHP reference observer are
reused unchanged. Seven non-array scalars and17 flat arrays distinguish integer
keys from sequential lists, negative/sparse/out-of-order keys and normalized
versus retained numeric-looking string keys. Integer values/keys are bounded to
±1,000,000; values are null/bool/int/string. Duplicate normalized keys, nested
values, floats, objects and references are outside this transport.

Seven separate direct-polyfill cases call owned drivers into the original
`array_all`: value/key correspondence, false results, empty arrays, all-true
iteration and unchanged caller inputs after callback-local reassignment.
Short-circuit cases place a nonnumeric poison string after the first/middle
false callback. If that later callback executes, numeric addition fails in PHP
and Go. The driver itself also uses logical short circuit to avoid adding the
stop marker. These are normal-exit assertions, not general error equivalence.

One compiler build translates both original declarations and drivers; the
generated report must match exact input/output hashes. One binary observes
all31 cases, calling the generated original functions rather than rewriting
the classification in Go. Actual Boolean type/value, native/runtime list flag
and typed ordered input-after values are compared against PHP and authored
expectations. Expected fields never supply Go observations. Stored keys are
observed without re-normalizing them. Both executions require empty stderr
and exit0. The old seventeen-original/149-case integrated binary is separate;
this new unit does not constitute a larger integrated bundle.

## Reproduction and capability exclusion

From `gossip/`, with cached immutable Docker images and a non-root user:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressNumericArrayHelper$'
go vet ./tests/gohp
```

The opt-in absent means skip, not PASS. Existing non-root/read-only/network-off,
resource limits, PHP8.3.36 floor and same-process admission guard remain intact.
Only the new configuration is run; old PASS/rejection configurations are not rerun.

Capability `gossip.dev.wordpress-unit.translate-compare` remains explicitly
**WebMCP-excluded**: local developer translation/comparison, temporary artifacts,
no browser/API operation, authentication route, external publication or auto-submit.
WordPress originals, SDK/hooks, PHP admission, prototype and master are unchanged.
No full-core/HTTP/DB/theme/plugin, deployment/release/license or main-merge claim.
Devboard remains unidentified/unsynchronized, not Done.
