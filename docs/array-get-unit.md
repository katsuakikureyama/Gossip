# Original nested settings lookup comparison

The new `array-get-unit.json` pins integrated GoHP #63 main
`1f10764ce7afec18276cd5b3da9662db575be7df`. It reuses main #62 foreach continue
and adds trailing untyped null defaults plus defined bare/multi-address isset.
GoHP's exact feature snapshot passed470 catalog checks on fixed PHP8.3.36.
This is compiler evidence, not a WordPress result by itself.

The harness extracts the complete, byte-unchanged `_wp_array_get` declaration
from functions.php5233–5272, including the default parameter, both isset paths
and continue statements. The WordPress baseline, whole source and retained
24-case corpus are hash-pinned. Native Reflection verifies a user-defined
function. Old `array-get-boundary.json`, cases, PHP observer and historical
rejection evidence are unchanged.

Go and PHP execute the helper independently. Owned observers decode bounded
null/bool/integer/string or ordered-pair array nodes, invoke the original with
two or three arguments as specified, then encode actual return type/value and
input/path/default states. Expected fields only assert results; they are never
substituted for calls. Both observers check arguments unchanged. Array keys
normalize through each runtime and duplicate normalized keys are rejected.
Limits: depth12,256 entries per node,4096-byte strings and integers±1,000,000.

The24 retained cases cover nesting, null/falsy/missing values, scalar
intermediates, array fallback/results, omitted defaults, invalid paths,
integer/numeric-string/Unicode/NUL keys and ordered sparse paths. New comparison
results are recorded in [evidence](evidence/array-get-unit-2026-10-26.md).
This is a separate one-original unit, not a25-original286-case integration.

From the nested `gossip/` module:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArrayGetUnit$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

Containers retain fixed image digests, nonroot/read-only/network-none/resource
limits and PHP8.3.36/C admission. Full Theme JSON/global-style/block callers,
WordPress startup, HTTP/DB/themes/plugins, objects/references/floats and general
PHP errors remain unverified. No configured production deployment target exists.

Capability `gossip.dev.wordpress-unit.translate-compare` remains explicitly
WebMCP-excluded: opt-in local developer comparison only, no browser/server
operation or route/auth change. Devboard card unidentified; PR/current-state
tracks the change. Do not mark comparison PASS before actual observations.
