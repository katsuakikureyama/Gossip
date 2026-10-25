# Original-source recursive wp_slash unit

This independent unit extracts the complete original `wp_slash()` unchanged
from `wp-includes/formatting.php` at WordPress baseline
`54106f6c545fc8b881f6b39b7963002bece6c28c`. The original tests for arrays,
recursively maps its own name, quotes actual strings with `addslashes`, and
returns other values unchanged. No observer implements that recursion or
quoting algorithm.

## Input and dependency boundary

The corpus supplies null, booleans, modest signed integers, byte strings
transportable in JSON and finite nested numeric lists. It includes quotes,
backslashes, NUL, multibyte strings and already escaped strings, plus mixed
non-string values and empty lists. Expected results are explicitly authored.
Both actual return type/value and the input after the call must match native
PHP and the expected transformed/original values respectively.

Go JSON decoding uses `UseNumber` and checks `Int64`. Maps, floats, objects,
references and cyclic graphs are excluded. Output/input-after arrays must have
sequential int64 keys before recursive JSON normalization. General associative
array transport and arbitrary large-integer JSON precision are not claimed.
Finite corpus recursion does not establish arbitrary-depth behavior or PHP
stack-overflow/exception compatibility.

GoHP now accepts the exact `array_map('literal_name', $array)` form for one
declared global one-argument callback. Names resolve at compile time and can
refer forward or recursively; callback arguments/results use existing array
value-copy semantics. Dynamic/builtin/null/method/closure callbacks, multiple
input arrays and by-reference arguments are outside the subset. The integrated
`is_array` predicate separates arrays from other existing runtime values.

`addslashes` is actual-string-only and byte-oriented: quotes/backslashes gain
a backslash, NUL becomes backslash plus ASCII zero, other bytes remain unchanged.
This is WordPress data quoting, not SQL/HTML/shell sanitization; no database or
external execution is introduced. See the PHP
[array_map](https://www.php.net/manual/en/function.array-map.php) and
[addslashes](https://www.php.net/manual/en/function.addslashes.php) contracts.

## Execution and separation

The manifest pins the whole original source hash, exact excerpt and committed
compiler/runtime. Report source/generated hashes are verified before building.
The existing approved immutable PHP8.3.36 image, same-process version guard,
non-root/no-network/read-only execution and resource limits remain unchanged.
Generation/build are intermediate statuses, not compatibility PASS. Actual
typed comparison requires matching values, empty stderr and exit0.

```sh
cd gossip
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressWpSlashHelper$'
```

Missing opt-in is SKIP. Existing six standalone units and both ten-/twelve-helper
bundles retain their previous pins/corpora/evidence. This is a seventh standalone
unit, not a thirteen-function bundle or a revalidation of earlier inputs.
`wp_unslash`, `map_deep`, general callbacks, object traversal, whole PHP files,
WordPress core startup, HTTP/DB, themes/plugins and hook integration remain
unsupported or unverified.

## Capability exclusion

`gossip.dev.wordpress-unit.translate-compare` remains WebMCP-excluded local
opt-in developer tooling. No UI/server operation, browser/S2S/MCP/callback auth
route, cookie forwarding, automatic submit or production audit change. No
default-branch merge, release, deployment or license decision. No devboard card
is linked, synchronized or marked Done.
