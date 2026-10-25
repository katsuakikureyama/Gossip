# Original array_first comparison

This unit uses GoHP main `4a4fffb995e57daf592785b6626e45a7babe08da`, which
includes the bounded array-variable empty construct from GoHP #57.
Its existing fixed PHP8.3.36 CI passed all427 compiler checks. Gossip compares
the actual original polyfill separately rather than deriving compatibility from
compiler-owned fixtures. The earlier local empty implementation is not used.

The WordPress baseline and whole `wp-includes/compat.php` SHA-256 are pinned
in `gossip/tests/gohp/array-first-unit.json`. Declaration lines657–665, including
the array parameter annotation, comments and body, are extracted byte-for-byte
for Go compilation. Native PHP executes the complete original guard646–666;
Reflection confirms user-defined origin. Conditional loading is not compiled.

The twelve array_first cases from the earlier24-case native boundary are retained
unchanged, including expectations and order. They cover empty, singleton
null/false/0/'0'/empty string, sparse/negative/mixed keys, insertion order,
canonical numeric-string keys and Unicode/NUL. Values are flat scalars and
integer keys/values are bounded to ±1,000,000 by the existing transport.
The test compares actual return type/value and ordered caller input after the
call against PHP and explicit expectations. No expected field creates the Go
result. Transport functions are selected unchanged from the existing numeric
observers with exact marker counts; the new Go observer calls F_array_first
directly. There are no substitute polyfill bodies or owned PHP drivers.

Run from `gossip/` as a non-root user with the pinned Docker images cached:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArrayFirst$'
go vet ./tests/gohp
```

An absent compiler-source opt-in is SKIP. The execution record must state the
tested commit and measured result. PHP8.3.36 admission runs in the same process
as the oracle; C locale, network-off, read-only, non-root and resource limits
use the existing harness.

This is a separate single-original configuration. It does not extend the
22-original/230-case single-binary result. Previous pins, corpora and evidence
remain historical records of their own revisions; the earlier rejected empty
probe is not rewritten. array_last still requires array_key_last. Nested arrays,
objects, references, PHP errors and WordPress startup/HTTP/DB/themes/plugins
are outside this unit.

Capability `gossip.dev.wordpress-unit.translate-compare` remains WebMCP-excluded
because it is a local developer compiler fixture with no browser operation,
service resource or authentication route. Devboard card remains unidentified;
progress is tracked by repository PR and current-state documentation.
