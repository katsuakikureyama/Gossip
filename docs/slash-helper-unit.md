# Parameterized original-source unit: trailing slashes

This unit adds two unchanged WordPress functions with a real dependency chain:
`trailingslashit($value)` calls `untrailingslashit($value)`, which calls
`rtrim($value, '/\\')`. The string result is observed directly, not through a
handwritten Go replacement of the WordPress functions.

## Source and scope

[`slash-helpers.json`](../gossip/tests/gohp/slash-helpers.json) pins WordPress
`54106f6c545fc8b881f6b39b7963002bece6c28c`, the complete `wp-includes/formatting.php`
hash, and exact excerpts: `trailingslashit` 2883–2885; `untrailingslashit` 2898–2900.
The bundle adds a PHP opening tag and separator newlines only. Function bodies,
spacing, parameter names and the original `rtrim` call remain unchanged.

GoHP compiler/runtime are pinned to
[`4e1b5abfe019ec121e13d85e99e51c206e765533`](https://github.com/katsuakikureyama/GoHP/tree/4e1b5abfe019ec121e13d85e99e51c206e765533),
with Gossip test commit `e943cbce1f74c93b5ca990de7a544caabb274aa5`. Its bounded
`rtrim` contract requires exactly two arguments, actual string input and a
literal byte mask without `..` ranges. It does not implement default masks,
dynamic masks, non-string coercion or all PHP trimming semantics. The prior
`ceee0fc71e1e7947f1c78f1dd18875236b9a180b` / Draft PR #10 pin remains only as
historical evidence; it does not select this current run.

The [case fixture](../gossip/tests/gohp/slash-cases.json) contains 13 explicit
input/expected pairs: empty and slash-only strings, repeated/mixed suffixes,
leading/interior slashes, Windows-style text, trailing whitespace, Japanese,
URL-looking text, embedded NUL and newline. Both functions are called for each
input, producing 26 observations. All return types must be string. A URL-looking
input is ordinary data: no URL parse, HTTP request or filesystem access is tested.

## Execution and evidence

The harness shares the existing isolated container/export helpers, not the
constant-return unit's mutable state. It exports only the pinned committed GoHP
source with `git archive`, verifies the WordPress source/excerpts, and stages
test-owned inputs/observers. It builds GoHP with downloads disabled, verifies
`gohp.build/v0.1` source/generated hashes before compilation, then builds against
the same pinned runtime. Compiler/report paths are fixed by the harness.

The separate Go `init()` observer calls the actual generated functions and exits
before generated main. The PHP observer calls the original functions. Both read
the same case data; neither uses expected outputs to calculate actual results.
The harness requires decoded JSON equality between Go, PHP and the authored
expected values, empty stderr and exit 0. JSON decoding, `get_debug_type` and
observation file reads are instrumentation, not translated PHP features.

Containers are non-root, network-disabled, read-only-root, capability-dropped
and resource-limited; only temporary candidate outputs are writable during build.
Execution mounts the input directory read-only and the candidate binary read-only.
The existing 90-second per-stage timeout and exact-container cleanup apply.
Source/license/image/observer/compiler/report/generated/binary hashes and all
typed observations are recorded under `gossip.slash-observation/v1`.

From the nested `gossip/` module, with pinned Docker images already cached:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressSlashHelpers$'
```

No environment variable means an explicit skip, not a verified translation.
The [constant-return unit](translation-units.md) and this unit now select one
compiler, but remain independently generated, built and compared. The current
PHP 8.3.36 result is recorded only in
[unified-compiler evidence](evidence/unified-compiler-2026-10-25.md); it does
not relabel the earlier per-pin evidence or turn the units into one bundle.
The fixture is intentionally string-only and does not validate weak scalar
coercion, warnings/exceptions, arbitrary byte input, full-file/core translation,
path canonicalization, HTTP behavior, PHP plugins, themes or database semantics.

## Capability exclusion

`gossip.dev.wordpress-unit.translate-compare` remains explicitly excluded from
WebMCP: this is opt-in local developer test tooling, not a user-facing service
operation. No UI action, authentication route, browser-cookie forwarding, MCP
endpoint, auto-submit, deployment, publication or production audit API is added.
