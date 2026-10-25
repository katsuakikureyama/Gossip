# Original-source unit: newline preservation

This unit extracts the unchanged WordPress
`_autop_newline_preservation_helper( $matches )` from
`wp-includes/formatting.php` lines 811–813, baseline
`54106f6c545fc8b881f6b39b7963002bece6c28c`. Its original body reads array key 0
and replaces line-feed bytes with the literal `<WPPreserveNewline />` marker.
It is one callback helper used by `wpautop`, **not** execution of `wpautop`, a
regular-expression callback engine or paragraph rendering as a whole.

## Input and compiler boundary

The owned fixture supplies nonempty JSON lists of strings, corresponding to
PHP numeric keys 0..n. Index 0 must exist and hold an actual string; extra string
captures are permitted and ignored by the original function. Empty lists,
associative maps, non-string values, missing-key warnings and coercion behavior
are outside this fixture. The observers check the contract before calling the
original/generated function and require a string return.

The required GoHP extension is exactly three-argument
`str_replace(search, replacement, subject)`, with all arguments actual strings.
The [PHP manual](https://www.php.net/manual/en/function.str-replace.php)
describes fixed-string, case-sensitive, binary-safe replacement. This candidate
uses nonoverlapping byte matches and treats the replacement as literal data,
not regex syntax. Empty search leaves the subject unchanged. Array forms,
fourth by-reference count, weak scalar conversion, regular expressions and
case-insensitive replacement remain unsupported. Known invalid argument types
and arity fail generation; unknown values are guarded at runtime.

The 16 explicitly authored input/expected cases cover ordinary/empty text,
single/adjacent/leading/trailing line feeds, carriage returns and CRLF, literal
backslash-n, whitespace, Japanese/emoji, NUL, existing marker text and extra
captures. The observers do not calculate expected replacement results. Go
converts lists to the existing ordered runtime array; it does not rewrite the
WordPress function or implement its replacement algorithm in the observer.

## Execution and separation

The manifest pins the full original-file hash, exact excerpt and compiler/runtime
commit. Export uses only committed compiler files. Report source/generated hashes
are checked before building, and the actual return type/value must match both
PHP 8.3.36 and the authored expectations, with empty stderr and exit 0.
Existing non-root/no-network/read-only/resource limits and the same-process
native PHP version guard remain in effect.

```sh
cd gossip
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressAutopNewlinePreservationHelper$'
```

Missing opt-in yields SKIP, not PASS. The
[execution record](evidence/newline-helper-2026-10-25.md) records actual outcomes,
revisions and hashes. Earlier standalone units and the ten-function combined
unit retain their existing pins/evidence, without being rerun or relabelled as
passing on this newer compiler. This helper is an additional independent unit,
not an eleven-function bundle or full WordPress conversion.

## Capability exclusion

`gossip.dev.wordpress-unit.translate-compare` remains **excluded** from WebMCP:
local opt-in developer tooling only. No UI action, service route, authentication,
cookie forwarding, automatic submit or production audit operation is added.
No default-branch merge, release, deployment or license finalization is implied.
No devboard card is linked, synced or marked Done. HTTP, database, themes,
plugins, core startup and hook SDK integration remain unverified.
