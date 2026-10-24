# Parameterized original-source unit: Boolean validation

This separate unit translates one unchanged WordPress helper,
`wp_validate_boolean( $value )`. It is a bounded mixed-input observation, not
a port of WordPress Boolean semantics in general.

## Source and contract

[`boolean-helper.json`](../gossip/tests/gohp/boolean-helper.json) pins the
WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`, the complete
`wp-includes/functions.php` hash, and the exact original lines 7932–7942. The
staged source adds only the PHP opening tag; the function's spacing, parameter,
branches and body are not rewritten.

The contract pins GoHP
[`14bb01053947f1bf2fb32edaff8085feb7227f48`](https://github.com/katsuakikureyama/GoHP/tree/14bb01053947f1bf2fb32edaff8085feb7227f48)
and Gossip test commit `8901d634e4d52e9be06c4c2d7f91714652974b28`. It has its
own source, compiler and case hashes; it neither updates nor combines the
constant-return and slash-helper pins.

[`boolean-cases.json`](../gossip/tests/gohp/boolean-cases.json) owns 23
input/expected pairs. They cover null, both Boolean values, zero/positive/
negative integers, empty and non-empty strings, case variants of `false`,
leading/trailing spaces, NUL-containing text, Japanese text, and empty/non-empty
nested arrays. Each case declares its input type so that JSON transport is part
of the observation. Expected values are authored fixture data, never calculated
by either observer from a copied version of the WordPress algorithm.

The bounded GoHP surface required here is deliberately explicit: one-arm `if`,
early return, `&&` short circuit, scalar `===`, Boolean cast, `is_bool`,
`is_string`, and ASCII `strtolower` for strings only. PHP runs with `LC_ALL=C`;
the observer first selects `setlocale(LC_CTYPE, 'C')`, then queries
`setlocale(LC_CTYPE, 0)` and requires `C` before observing case values. The
GoHP native harness uses the same process-level selection and verification.
This makes the locale precondition visible rather than assuming a host's
case-conversion behavior.

Float and object inputs are outside this fixture and unsupported by this
candidate. No claim is made for arbitrary PHP coercion, Unicode case folding,
warnings/exceptions, whole-file/core translation, hooks, themes, plugins, HTTP,
filesystem or database behavior.

## Execution and record

The harness follows the existing isolated export/container design. It exports
only the pinned committed GoHP source with `git archive`, verifies the whole
WordPress-file and exact-excerpt hashes, and stages only fixture-owned inputs and
observers. Generation runs without network access, verifies the
`gohp.build/v0.1` report hashes before compilation, and builds against that same
pinned runtime.

The Go observer decodes JSON with `UseNumber`, permits only checked signed
integers, and recursively maps JSON lists to the GoHP runtime array type. It
rejects maps/objects and non-integer numeric values. The PHP observer decodes the
same case fixture, and both observers assert the declared input type before
calling the original or generated helper. Each emits only name, input type,
Boolean return type and Boolean value. The harness compares decoded Go output,
PHP output and the authored expected values, requiring empty stderr and exit 0.

From the nested `gossip/` module, with the pinned images already cached:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressBooleanHelper$'
```

The test is opt-in: an absent `GOSSIP_GOHP_SOURCE` produces a skip, not a
translation result. Its outcome, hashes and observed values are recorded
separately in the
[Boolean execution record](evidence/boolean-helper-2026-10-25.md). That record
is the only place to determine this contract's executed status; this document
defines scope and procedure, not a compatibility conclusion.

## Capability exclusion

`gossip.dev.wordpress-unit.translate-compare` remains explicitly excluded from
WebMCP publication. This Boolean unit is opt-in local developer tooling, not a
user-facing service operation. It adds no UI action, HTTP/MCP route, browser
session or cookie forwarding, automatic submit, deployment, publication or
production audit API.
