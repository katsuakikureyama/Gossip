# Next original helper: integer-key array classification

This is a **native PHP reference and rejected-translation boundary**, not a
twelfth successful standalone Go unit or an eighteen-original integration.
The existing seventeen-original/149-case result remains separate and unchanged.

## Exact source and dependencies

The selected original is `wp_is_numeric_array()` at
`wp-includes/functions.php:5448–5454`. It tests the keys, not the values, and
does not require a zero-based consecutive list. Empty arrays and arrays with
sparse, negative or out-of-order integer keys qualify. A surviving string key
does not. PHP first normalizes canonical decimal string keys when constructing
the array; strings such as `01`, `+1` and `-0` remain string keys.

The unchanged function contains a `: bool` return declaration and a two-argument
arrow callback passed to `array_all`. On the admitted PHP 8.3 runner, the native
reference must load WordPress's original polyfill, not an invented replacement.
The harness extracts the complete guarded block at `compat.php:623–644`,
including its original `array`/`callable` parameters, Boolean return declaration,
callback invocation, short circuit and empty-array result. The observer requires
`array_all` to be user-defined, excluding accidental use of a newer native builtin.

Whole-source hashes, exact extraction ranges, WordPress revision, compiler and
immutable container images are pinned in `numeric-array-boundary.json`. Source
fragments, including comments and indentation, are never rewritten. The same
polyfill declaration (`compat.php:635–643`) is also probed independently so the
helper's earlier diagnostic does not conceal its dependency boundary.

## Reference contract

Twenty-four authored cases use root ordered key/value pairs or scalar inputs.
Array values are flat null/bool/integer/string values; integers and integer keys
are bounded to ±1,000,000. The cases distinguish integer-key arrays from actual
lists and cover normalized versus retained numeric-looking string keys and
non-array inputs. Nested values, objects, floats, references and duplicate
normalized keys are outside this transport contract.

The PHP observer calls the actual original helper and records actual Boolean
return type/value, `array_is_list` when applicable, and actual caller input after
the call with typed ordered keys. It does not construct outputs from expected
fields. The host compares the observations with authored expectations and
requires empty native stderr/exit0. This alone does not validate a Go candidate.

## Compiler boundary and next implementation

GoHP remains pinned to `49a302e44464623dc4912bafbcf44192759d7878` (#44).
The helper and polyfill declaration are each submitted unchanged to one built
compiler. Both must exit1 with the expected parser diagnostic and publish
neither generated Go nor a generated-status report. These are first-error
observations, not an exhaustive dynamic measurement of every dependency.
There is no generated candidate binary, Go behavior observation or differential
PASS in this step. The evidence status is deliberately
`native_reference_pass_candidate_rejected`; `differential_status` is `not-run`.

The next implementation sequence is:

1. Define bounded Boolean return-declaration behavior, including explicit
   mismatch/fallthrough rejection and native evidence; do not erase the type.
2. Define the two-argument arrow/callable contract and the unchanged polyfill's
   typed parameters, callback value/key ordering, early termination and empty
   input. The retained corpus does not yet prove callback side effects or copy
   behavior; those need focused compiler evidence when implemented.
3. Execute both original declarations under a separately pinned compiler and
   compare actual Go returns/input-after with this reference. Only then add a
   successful standalone unit and later a separate integrated bundle.

The full conditional `function_exists` wrapper is needed by the native reference.
Translating original function declarations alone may omit that top-level loading
wrapper only as an explicitly documented extraction boundary; it cannot pretend
to support WordPress bootstrap or arbitrary conditional declarations. Do not
raise the admitted PHP series, stub `array_all`, rewrite the helper, or silently
adopt a newer compiler to bypass this dependency.

Other candidates were not selected: current `absint()` (`load.php:1546–1574`)
contains float/finite/range/constant/cast behavior rather than the historical
one-line implementation; `_wp_array_get()` has optional arguments, additional
type predicates, scalar/multiple-operand `isset` and `continue`; the public
charset wrapper depends on defaults, null coalescing and `get_option`.
This was a bounded survey, not an exhaustive inventory of every WordPress helper.

## Reproduction and capability exclusion

From the nested `gossip/` directory:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressNumericArrayBoundary$'
go vet ./tests/gohp
```

Absent opt-in means skip, not PASS. Existing image admission, same-process
PHP8.3.36 floor, C locale, non-root/read-only/network-disabled and resource-limited
containers are retained. Only the new boundary test runs; old successful
configurations are not repinned or rerun. The
[execution record](evidence/numeric-array-boundary-2026-10-25.md) records24 native
reference matches and two expected compiler rejections, not differential PASS.

Local capability `gossip.dev.wordpress-unit.translate-compare` remains explicitly
**WebMCP-excluded**: this adds an opt-in developer boundary/reference fixture,
not a browser/S2S/MCP operation, UI action, session forwarding, publication,
deployment or production API. Temporary artifacts are local test outputs.
WordPress originals, SDK/hooks, oracle policy, prototype and master remain
unchanged. Whole-core/HTTP/DB/themes/plugins, merge/deploy/release and licensing
remain unresolved. Devboard is unidentified/unsynchronized, not Done.
