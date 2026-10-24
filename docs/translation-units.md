# Original WordPress translation units

The first GoHP experiment translates **unchanged excerpts of actual WordPress
functions**, not a handwritten Go equivalent or a rewritten PHP approximation.
It does not translate or execute all of `functions.php` or WordPress.

## Pinned unit

The manifest is
[`gossip/tests/gohp/return-helpers.json`](../gossip/tests/gohp/return-helpers.json).
It pins the original WordPress commit and entire `wp-includes/functions.php`
SHA-256 before selecting these complete functions:

| Function | Original lines (inclusive) | Echo observation / typed return |
| --- | --- | --- |
| `__return_true` | 7181–7183 | `true=1` / bool `true` |
| `__return_false` | 7196–7198 | `false=` / bool `false` |
| `__return_zero` | 7209–7211 | `zero=0` / int `0` |
| `__return_null` | 7235–7237 | `null=` / null |
| `__return_empty_string` | 7250–7252 | `empty-string=` / string `""` |

Each observation is newline-terminated. Original function text, indentation and
inline comments are copied byte-for-byte. The bundle adds only a PHP opening
tag, separator newlines and the owned call suffix `calls.php`. No WordPress
function is changed to satisfy GoHP. Each excerpt hash and its original/bundle
line location are recorded in the result. Upstream files remain untouched.

`__return_empty_array` (7222–7224) is a separate rejection fixture, not a
successfully translated function. PHP must return an empty array; pinned GoHP
must exit 1 with `PHP004 line 3: unsupported function array`, produce no stdout,
and leave neither a generated Go file nor a report. Its original `array()` is
never rewritten to `[]`. Dynamic callbacks, plugin APIs,
global state, references, includes, side effects and full WordPress startup
are also outside it. Successful conversion of these literal-return helpers does
not remove the original hook engine's GoHP blockers.

## Execution boundary

1. Require the pinned complete WordPress source hash and matching fork baseline.
2. Export the pinned GoHP commit with `git archive`; do not build a caller's
   modified working tree. Stage regular files/directories only, not `.git`,
   credentials or symlinks. No remote fetch is performed by the test.
3. In a network-disabled, non-root, read-only-root container, build GoHP and
   generate a fresh Go file/report. Mount only staged sources/input read-only
   and a test-owned candidate directory writable. Never invoke Composer or
   repository shell scripts from the compiler source.
4. Require successful generation, recognize `gohp.build/v0.1`, and verify both
   source/generated hashes before using the generated source. Ignore additive
   report fields; never use report-supplied paths as filesystem authority.
5. Build against the **same pinned GoHP runtime**, with network/toolchain
   downloads disabled. Record `compiled` separately from `generated`.
6. Run the resulting binary and the PHP bundle in separate isolated containers.
   Require exit 0, empty stderr and exactly matching stdout, also matching the
   owned expected observation.
7. Compile a separate Go observer beside the unchanged generated file. Its
   `init()` calls the actual generated functions, records their return types and
   values as JSON, then exits before the echo-based main. Independently observe
   the unchanged PHP definitions. Require both parsed records to match each other
   and `typed-expected.json`, with empty stderr and exit 0.
8. Require the original empty-array fixture's PHP success and explicit compiler
   rejection described above. Only then report `differential_pass` for the five
   scalar helpers, while recording the sixth function separately as `rejected`.

GoHP commit: `9c7389d35032d623a429198a69cdfa5c11571e8e` (draft candidate).
Go and PHP runner images are immutable digests in the unit manifest. GoHP needs
Go 1.23+, so the existing cached Go 1.25 image is used inside the builder; the
host can continue running the Gossip test driver with Go 1.22.

Every container uses no network, no added capabilities, no-new-privileges,
non-root host UID/GID, 768 MiB memory, 256 PIDs, two CPUs and a 90-second stage
timeout. Builder `/tmp` is a 256 MiB noexec tmpfs. Compiler caches/artifacts live
only in the test's temporary candidate directory. On timeout the harness removes
only its exact randomly named container. Successful containers use `--rm`; Go
cleans the test directories. There are no persistent services or host ports.

## Reproduction

Obtain the pinned GoHP source and runner images deliberately; the test itself
will neither pull images nor download a compiler/dependency. From `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp
```

The clone must contain the manifest's commit object; the checked-out branch
does not select the compiler version. Without `GOSSIP_GOHP_SOURCE` this test
explicitly skips. A skipped test is not translation evidence.

The JSON observation printed by a successful run records original source and
excerpts, source bundle, GoHP commit/archive/compiler binary, generated source,
compiler report, candidate binary, runner images, toolchain and compared output.
The compiler's own report remains `generated`; the independent observation is
the narrow differential result. Neither status activates a plugin or release.

The v2 observation additionally records typed observations, observer/definitions/
typed-binary hashes and the rejected unit. The observers are test instrumentation:
PHP `get_debug_type`/`json_encode` and dynamic calls are executed only by PHP, not
translated. Go's `F___return_*` symbols are a private convention of the pinned
compiler, not a stable integration API. The adapter accepts only bool, int64,
string and nil, mapping them to PHP type labels; unexpected Go types fail.

The typed probe distinguishes false, null and empty string for these five actual
returns. It does not prove general return-type equivalence, PHP warnings/exception
behavior, arbitrary inputs, mutation or complete language semantics. It is a
narrow original-source translation path, not production compatibility.

## Capability exclusion

`gossip.dev.wordpress-unit.translate-compare` is explicitly **excluded** from
WebMCP: it is an opt-in local developer integration test, not a browser, MCP or
S2S operation. It adds no auth/session/cookie routing, UI callsite, auto-submit,
publication, install/activation, audit API or production migration operation.
Generated artifacts are temporary test outputs, not a supported deployment API.
