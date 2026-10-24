# Original-source unit: stripslashes from strings only

This independent unit extracts `stripslashes_from_strings_only()` unchanged
from `wp-includes/formatting.php` lines 2934–2936 at WordPress baseline
`54106f6c545fc8b881f6b39b7963002bece6c28c`. The original conditional calls
`stripslashes` only for actual strings and returns other inputs unchanged.
It does not execute `stripslashes_deep`, `map_deep` or `wp_unslash`.

## Bounded input and compiler contract

Inputs are null, booleans, signed integers within the authored corpus, strings
and nested numeric lists of those values. Objects, associative maps, floats,
references and arbitrary PHP values are excluded. JSON integers are decoded
with `UseNumber` and checked as int64 before conversion to the runtime. Runtime
arrays must have ordered consecutive integer keys when observed as JSON lists.
Nested list strings stay escaped: the original helper is not recursive.

The compiler extension is full three-operand `condition ? yes : no`, with
supported scalar/array truthiness, one evaluation of the condition and only
the selected arm evaluated at runtime. Both arms must still belong to the
closed compile-time subset. Parentheses disambiguate nested operands; PHP 8
non-associativity is retained. Short `?:`, `??` and object truthiness are not
added. See the [PHP conditional operator contract](https://www.php.net/manual/en/language.operators.comparison.php#language.operators.comparison.ternary).

The new builtin accepts one actual string, with static/dynamic type guards.
It consumes backslash pairs byte-wise in one pass: escaped ASCII zero becomes
NUL, other following bytes survive without the backslash, and a trailing
backslash disappears. Existing NUL and unescaped multibyte bytes survive.
This is not `stripcslashes`, a C escape parser, sanitization or recursive
unescaping. Weak coercion and PHP warning/exception equivalence are excluded.
See the [manual](https://www.php.net/manual/en/function.stripslashes.php) and
[PHP 8.3 implementation](https://github.com/php/php-src/blob/PHP-8.3/ext/standard/string.c).

## Observation and separation

The owned cases contain explicit expected values, not calculated results.
Observers only convert/check types and dispatch to original/generated code.
Actual return type/value must match PHP 8.3.36 and the expectations, with empty
stderr and exit 0. The manifest pins the whole source hash, exact excerpt and
committed compiler/runtime; report hashes are checked before compilation.
The approved immutable PHP image, same-process version guard and existing
non-root, no-network, read-only and resource limits are unchanged.

```sh
cd gossip
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressStripslashesFromStringsOnlyHelper$'
```

Missing opt-in is SKIP, not PASS. Earlier standalone and combined-helper units
retain their pins, cases and evidence; this addition is not a twelve-function
bundle or a revalidation of earlier units on the new compiler.

## Capability exclusion

`gossip.dev.wordpress-unit.translate-compare` remains **excluded** from WebMCP:
local opt-in developer tooling, not a browser action or service operation.
No authentication route, cookie forwarding, automatic submit, production audit,
merge, deployment, release or license decision is introduced. HTTP, database,
themes, plugins, hook SDK integration and core startup remain unverified.
No devboard card is identified, synchronized or marked Done.
