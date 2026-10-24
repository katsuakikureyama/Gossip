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
| PHP-to-Go translation | Separate GoHP project | Six original constant-return helpers have their own isolated evidence, and two original slash helpers separately match 13 string inputs/26 typed results. A third, separately pinned Boolean-helper contract is tracked independently. None of these is general array, `rtrim`, scalar-coercion or full-core support; the hook engine remains unsupported |
| Earlier Go HTTP/editor/import/release prototype | Separate `Gossip-prototype` repository | Historical prototype only; not installed or activated by this fork |

## GoHP readiness check — 2026-10-25 JST

The slash-helper unit inspects
[`ceee0fc71e1e7947f1c78f1dd18875236b9a180b`](https://github.com/katsuakikureyama/GoHP/tree/ceee0fc71e1e7947f1c78f1dd18875236b9a180b),
the unmerged Draft GoHP PR #10 stack based on `feat/legacy-array-syntax`. It
adds only the bounded `rtrim` form needed by that unit: two arguments and a
literal byte mask. It does not establish ranges, non-string coercion, dynamic
masks, or general PHP builtin compatibility. The separately pinned
[`c67fbefe34b1b098c0a9f475504d06bd014de23c`](https://github.com/katsuakikureyama/GoHP/tree/c67fbefe34b1b098c0a9f475504d06bd014de23c)
remains the evidence pin for the six constant-return helpers and Draft GoHP PR
#8's legacy `array()` expression. Neither candidate is a stable WordPress
translator interface or a claim of general array compatibility.

The original hook engine requires features outside that subset: include/require,
global state, references, array mutation, dynamic callbacks, variadic arguments,
interfaces and other PHP builtins/object behavior. See the pinned GoHP
[compatibility report](https://github.com/katsuakikureyama/GoHP/blob/c67fbefe34b1b098c0a9f475504d06bd014de23c/docs/compatibility.md)
and [integration contract](https://github.com/katsuakikureyama/GoHP/blob/c67fbefe34b1b098c0a9f475504d06bd014de23c/docs/integration.md).
No full-file translation attempt or deployable WordPress artifact is claimed.
The [helper-unit experiment](translation-units.md) converts only six unchanged
literal-return functions under its own pin; the separate
[slash-helper unit](slash-helper-unit.md) records a different pin and contract.
GoHP PR #8's [CI run 40112122164](https://github.com/katsuakikureyama/GoHP/actions/runs/40112122164)
PASSed on fixed PHP 7.4.33 and 8.3.30 with Go 1.23.12, recording
`diff.legacy_arrays` and seven negative checks per target. It remains compiler
evidence, not a broad WordPress compatibility result.

GoHP PR #10's [CI run 40112678059](https://github.com/katsuakikureyama/GoHP/actions/runs/40112678059)
PASSed on fixed PHP 7.4.33 and 8.3.30 with Go 1.23.12, recording `diff.rtrim`,
nine negative checks, and two runtime checks per target. It is scoped compiler
evidence for the bounded `rtrim` addition, not a broad PHP or WordPress
compatibility result.

The separate Boolean-helper contract pins GoHP
[`14bb01053947f1bf2fb32edaff8085feb7227f48`](https://github.com/katsuakikureyama/GoHP/tree/14bb01053947f1bf2fb32edaff8085feb7227f48)
and Gossip test commit `8901d634e4d52e9be06c4c2d7f91714652974b28`. It extracts
only `wp_validate_boolean()` from `wp-includes/functions.php` lines 7932–7942
and records a 23-case mixed-input contract. Its generated/build/differential
outcome belongs only in the
[Boolean execution record](evidence/boolean-helper-2026-10-25.md); it does not
rerun, revise or combine the older return-helper and slash-helper evidence.
The bounded compiler work covers the helper's one-arm early return, `&&`
short-circuiting, scalar `===`, Boolean cast, `is_bool`, `is_string`, and
string-only ASCII `strtolower` under an explicitly selected and verified
`LC_CTYPE=C`. Float and object behavior, and whole-file/core translation,
remain unsupported.

Do not replace the PHP originals to accommodate the translator. A later adapter
must pin its inputs, run in isolation, record generated/build/differential
outcomes separately, and keep a failed candidate inactive. A translated file
that compiles is not evidence of WordPress compatibility.

## Next bounded steps

1. Six literal-return helpers pass under their own pin, including the unchanged
   original empty-array helper. Separately, two slash helpers pass the 13-case
   string contract under the newer pin. Neither result is a combined eight-
   function test or evidence for arbitrary arrays, writes, mutations, paths,
   or general `rtrim`. The Boolean helper has its own 23-case contract and
   execution record, rather than being folded into either earlier result.
   Select the next dependency-bearing unit with its own explicit contract; the
   neighboring HTTP header helper still requires a separate side-effect contract
   and remains unverified.
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
WebMCP publication. WebMCP support is not claimed.
