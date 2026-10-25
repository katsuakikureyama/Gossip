# Original-source unit: zeroise

This independent unit extracts unchanged `zeroise( $number, $threshold )` from
`wp-includes/formatting.php` lines 2848–2850 at the pinned WordPress baseline
`54106f6c545fc8b881f6b39b7963002bece6c28c`. Only the PHP opening tag is added.
The full source-file hash and exact original range are checked before generation.
It is not a rewrite of the function and does not translate the whole file.

## Closed input and execution contract

The owned case fixture supplies strings or signed 64-bit integers as `$number`
and nonnegative integer thresholds. GoHP accepts only two-argument `sprintf`
with format `%0Ns`, where N is canonical decimal 0–4096. It pads to a minimum
byte width with ASCII zeros, never truncating longer strings. Signs are part of
the string, not numeric padding directives. Other format directives, flags,
precision, argument counts, floats, bool/null coercion, arrays and objects are
outside the contract. Width 4096 is an implementation resource bound, not a PHP
limit. Known invalid literal formats/types are rejected by the compiler; dynamic
formats/types are checked at runtime.

The fixture includes empty and leading-zero strings, negative numbers,
multibyte text, NUL and signed integer boundaries. Expected outputs are explicit
fixture data, not computed by the observers. Both observers assert the declared
input type and threshold range before calling the original/generated function;
the Go observer uses JSON `UseNumber` and checked `Int64` conversion. Actual
string return types and values must match PHP and the authored expectations.

The existing hardened container/export/report verification workflow is reused.
Only committed pinned compiler/runtime source is exported, report hashes are
checked before building, and execution requires exit 0 and empty stderr. The
native oracle retains the approved PHP 8.3.36 image and same-process version
guard. Neither WordPress original source nor existing unit pins are changed.
The three earlier units retain their separate shared-compiler evidence; this
fourth unit has a new compiler pin and does not establish a ten-function bundle.

From `gossip/`, with pinned images cached:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressZeroiseHelper$'
```

Missing opt-in configuration yields SKIP, never a translation PASS. See the
[execution record](evidence/zeroise-helper-2026-10-25.md) for the actual tested
commits, cases, hashes and outcome. General PHP formatting, WordPress startup,
HTTP/DB/plugin/theme behavior and native hook integration remain unverified.

## Capability exclusion

`gossip.dev.wordpress-unit.translate-compare` remains explicitly **excluded**
from WebMCP: opt-in local developer tooling, no UI action, service route,
browser/MCP authentication, cookie forwarding, automatic submit, publication,
production audit endpoint or deployment. No devboard card is linked or Done.
