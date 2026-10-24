# Reuse and replacement boundary

This is an upstream-history-preserving WordPress fork, not an independent Go
CMS with a WordPress import option. All replacements below are proposals until
specific source, platform and behavioral evidence exists.

| Surface | Current source of behavior | Go / GoHP status |
| --- | --- | --- |
| Core hooks (`plugin.php`, `class-wp-hook.php`, filter sentinel) | Original, hash-pinned files in this fork | PHP execution is the oracle; native SDK compares only five owned single-string cases |
| Themes, assets, Gutenberg markup, `theme.json` | Inherited WordPress files and runtime contracts | No Go rendering port; presence of assets does not prove they work in a Go runtime |
| Posts, terms, options, metadata, DB schema | Original WordPress implementation | No replacement with the prototype JSON store; no migration or Go DB mapping here |
| Browser sessions, permissions and HTTP routes | Original WordPress implementation | No new routes or auth changes; no Go browser session claim |
| Native Go plugins | Isolated `gossip/pkg/pluginsdk` experiment | Trusted in-process registry; not linked into the inherited PHP request lifecycle |
| PHP-to-Go translation | Separate GoHP project | Six original constant-return helpers, two slash helpers and one Boolean helper remain three independently generated/built/compared units. The current run selects one shared compiler pin, not a nine-function bundle or full-core test. This is not general array, `rtrim`, scalar-coercion or full-core support; the hook engine remains unsupported |
| Earlier Go HTTP/editor/import/release prototype | Separate `Gossip-prototype` repository | Historical prototype only; not installed or activated by this fork |

## Docker PHP oracle admission — 2026-10-25 JST

New Gossip Docker-oracle runs for hooks, standalone units and the combined unit
use only the official PHP 8.3 image pinned at
`php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd`.
The shared oracle helper verifies that exact image allowlist before a container
starts, checks the actual interpreter version in the same PHP process before
requiring any original source, and records the observed version. The admission
floor is native PHP 8.3.36: PHP 7.4, older 8.3 patches, 8.4 or later,
pre-releases and malformed versions are rejected. This is runner admission, not
evidence that an admitted interpreter is compatible with a broader WordPress or
PHP surface; see the [PHP oracle policy](php-oracle-policy.md) and the separate
[PHP 8.3 floor execution record](evidence/php-83-floor-2026-10-25.md).

The earlier Gossip evidence files retain their PHP 8.3.13 observations as
historical records. The prior PHP 8.3.36 record also remains intact with its
then-distinct compiler pins. Neither is relabelled as the current unified run,
which is recorded independently.

The fourth [zeroise unit](zeroise-helper-unit.md) extends a descendant compiler
with bounded `sprintf('%0Ns', value)` for strings/integers only. Its new pin and
[execution evidence](evidence/zeroise-helper-2026-10-25.md) do not change the
earlier nine-function results below. Ten selected original functions across
four units and two compiler pins are not a combined WordPress program or general
formatting compatibility. Original source remains unchanged.

## GoHP readiness check — 2026-10-25 JST

The [string-only stripslashes unit](stripslashes-helper-unit.md) adds full
three-operand conditional expressions and a guarded string-only builtin.
The original returns non-strings unchanged, including escaped strings inside
lists; it does not execute recursive `stripslashes_deep`, `map_deep` or
`wp_unslash`. Earlier pins/evidence remain independent and unchanged.

The newer [newline-preservation unit](newline-helper-unit.md) adds an independent
original callback-helper observation using a three-string `str_replace` subset.
Its [separate result](evidence/newline-helper-2026-10-25.md) does not translate
the enclosing `wpautop` renderer, regex/callback execution or full HTML behavior.
The earlier ten-function bundle and standalone evidence retain their old pins.

The subsequent [combined helper experiment](helper-bundle-unit.md) translates
the ten selected originals together, adding five owned composition drivers and
22 new direct/composed probes. Its [separate result](evidence/helper-bundle-2026-10-25.md)
applies to one generated binary, not to the complete previous input corpora or
WordPress startup. The standalone pins and historical results below remain intact.

The current unified compiler candidate is
[`4e1b5abfe019ec121e13d85e99e51c206e765533`](https://github.com/katsuakikureyama/GoHP/tree/4e1b5abfe019ec121e13d85e99e51c206e765533),
with Gossip test commit `e943cbce1f74c93b5ca990de7a544caabb274aa5`. It selects
the same compiler/runtime for three independent units: six constant-return
helpers, two slash helpers and `wp_validate_boolean()`. Each unit retains its
own source excerpt, fixture, expected values, generated program and PHP oracle;
this is nine functions across three comparisons, not one bundle or whole-core
translation. Its execution outcome belongs only in
[unified-compiler evidence](evidence/unified-compiler-2026-10-25.md).

This candidate backports only GoHP PHP 8.3 policy commit
`948647c2903a4a1b22069818861062df1f5bf414` from separate PR #14 to Draft PR
#16, which is stacked on #13. It does not import Composer inventory, `strlen`,
`array_keys` or `array_values` work. The earlier
`c67fbefe34b1b098c0a9f475504d06bd014de23c`,
`ceee0fc71e1e7947f1c78f1dd18875236b9a180b` and
`14bb01053947f1bf2fb32edaff8085feb7227f48` pins remain historical explanations
for their existing evidence; they do not select this run.

The original hook engine requires features outside that subset: include/require,
global state, references, array mutation, dynamic callbacks, variadic arguments,
interfaces and other PHP builtins/object behavior. See the current pinned GoHP
[compatibility report](https://github.com/katsuakikureyama/GoHP/blob/4e1b5abfe019ec121e13d85e99e51c206e765533/docs/compatibility.md)
and [integration contract](https://github.com/katsuakikureyama/GoHP/blob/4e1b5abfe019ec121e13d85e99e51c206e765533/docs/integration.md).
No full-file translation attempt or deployable WordPress artifact is claimed.
The [helper-unit experiment](translation-units.md),
[slash-helper unit](slash-helper-unit.md) and
[Boolean-helper unit](boolean-helper-unit.md) share the current compiler pin
but retain separate contracts and observations.
GoHP PR #8's [CI run 40112122164](https://github.com/katsuakikureyama/GoHP/actions/runs/40112122164)
PASSed on fixed PHP 7.4.33 and 8.3.30 with Go 1.23.12, recording
`diff.legacy_arrays` and seven negative checks per target. It remains compiler
evidence, not a broad WordPress compatibility result.

GoHP PR #10's [CI run 40112678059](https://github.com/katsuakikureyama/GoHP/actions/runs/40112678059)
PASSed on fixed PHP 7.4.33 and 8.3.30 with Go 1.23.12, recording `diff.rtrim`,
nine negative checks, and two runtime checks per target. It is scoped compiler
evidence for the bounded `rtrim` addition, not a broad PHP or WordPress
compatibility result.

Do not replace the PHP originals to accommodate the translator. A later adapter
must pin its inputs, run in isolation, record generated/build/differential
outcomes separately, and keep a failed candidate inactive. A translated file
that compiles is not evidence of WordPress compatibility.

## Next bounded steps

1. Keep the shared-compiler run to its three independent unit boundaries: six
   literal-return helpers, two slash helpers and one Boolean helper. It is not
   a nine-function bundle or evidence for arbitrary arrays, writes, mutations,
   paths or general `rtrim`. Select the next dependency-bearing unit with its
   own explicit contract; the neighboring HTTP header helper still requires a
   separate side-effect contract and remains unverified.
2. Expand the comparison contract only for demonstrated behavior. Negative
   priorities, references, removal during dispatch, recursion and dynamic hook
   registration remain outside the current native SDK evidence.
3. Design the runtime integration and DB-copy plan before adding CMS features.
   A/B binary rollback must not be described as DB migration rollback.

No production database, secret, site deployment, upstream PR, default-branch
merge or license finalization is authorized by this baseline work.

## Capabilities and authentication

`gossip.dev.source.inspect` and `gossip.dev.hooks.compare` are explicitly excluded
from WebMCP publication: they are local developer tests, not user-facing site
operations. No HTTP/MCP operation, automatic submit, browser cookie forwarding,
permission model or audit service is introduced. The SDK event names are not
executable user-facing capabilities. The Boolean-helper fixture remains within
the same explicit local-test exclusion; it introduces no capability, route or
WebMCP publication. PHP oracle admission is also local test infrastructure, not
a user-facing operation or WebMCP capability. WebMCP support is not claimed.
