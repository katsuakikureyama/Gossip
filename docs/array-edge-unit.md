# Original array_first and array_last comparison

This unit pins GoHP main `440308c77da07cbcdce403fed7fff3945e0d2078` (#58),
which provides bounded empty and insertion-order array_key_last. Its fixed
PHP8.3.36 CI433-check PASS is compiler evidence; this unit separately measures
the unchanged WordPress originals in one generated executable.

The manifest `gossip/tests/gohp/array-edge-unit.json` fixes the WordPress baseline,
whole compat.php SHA-256, compiler commit and immutable runner images.
Declarations array_first657–665 and array_last679–685 retain their exact bytes,
annotations and bodies. Native PHP executes complete guards646–666/668–686,
and Reflection confirms both are user-defined polyfills. Extraction of a
declaration does not compile top-level conditional loading.

The existing24 edge cases and PHP observer are reused unchanged. Each original
has12 cases: empty, singleton null/false/0/'0'/empty string, insertion order,
sparse/negative/mixed keys, canonical numeric-string keys and Unicode/NUL.
Values are flat scalars, integers bounded to ±1,000,000. The new Go observer
calls the actual generated F_array_first/F_array_last functions and measures
return type/value and ordered caller input after the call. The harness compares
these observations with PHP and authored expectations. Expected fields do not
create runtime results. No substitute bodies or extra PHP drivers are used.

Transport functions are selected unchanged from existing numeric observers.
The existing isolated harness retains PHP8.3.36 same-process admission and
C locale, non-root/read-only/network-off operation and resource limits.

From the nested `gossip/` module with the fixed images cached:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArrayEdgeUnit$'
go vet ./tests/gohp
```

Without compiler-source opt-in the test is SKIP. Record actual results and the
exact committed configuration before declaring a comparison PASS.

This two-original unit has its own compiler pin. The earlier array_first12-case
comparison, rejected-boundary record and22-original/230-case integration retain
their configurations. Passing this unit does not establish a24-original254-case
combined binary. Nested arrays, references, objects, errors and WordPress
startup/HTTP/DB/themes/plugins remain outside this contract.

Capability `gossip.dev.wordpress-unit.translate-compare` remains WebMCP-excluded
as a local developer fixture with no browser operation, service resource or auth
route. Progress is tracked by the PR and current-state; no devboard card has
been identified.
