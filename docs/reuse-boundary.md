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
| PHP-to-Go translation | Separate GoHP project | Six original constant-return helpers pass isolated comparison: five scalar echo observations and all six typed returns, including one empty array. This is not general array support; full WordPress core and the hook engine remain unsupported |
| Earlier Go HTTP/editor/import/release prototype | Separate `Gossip-prototype` repository | Historical prototype only; not installed or activated by this fork |

## GoHP readiness check — 2026-10-25 JST

The inspected candidate is
[`c67fbefe34b1b098c0a9f475504d06bd014de23c`](https://github.com/katsuakikureyama/GoHP/tree/c67fbefe34b1b098c0a9f475504d06bd014de23c),
the unmerged Draft GoHP PR #8 stack based on `feat/compatibility-views`. Its
documented closed subset includes basic scalar operations, named functions,
literal-array reads/foreach and limited classes. PR #8 adds the original
legacy `array()` expression needed by `__return_empty_array()` through GoHP's
shared parser; it makes no Go runtime change. This is not a stable WordPress
translator interface or a claim of general array compatibility.

The original hook engine requires features outside that subset: include/require,
global state, references, array mutation, dynamic callbacks, variadic arguments,
interfaces and other PHP builtins/object behavior. See the pinned GoHP
[compatibility report](https://github.com/katsuakikureyama/GoHP/blob/c67fbefe34b1b098c0a9f475504d06bd014de23c/docs/compatibility.md)
and [integration contract](https://github.com/katsuakikureyama/GoHP/blob/c67fbefe34b1b098c0a9f475504d06bd014de23c/docs/integration.md).
No full-file translation attempt or deployable WordPress artifact is claimed.
The later [helper-unit experiment](translation-units.md) converts only six
unchanged literal-return functions and records its own narrow evidence. GoHP
PR #8's [CI run 40112122164](https://github.com/katsuakikureyama/GoHP/actions/runs/40112122164)
PASSed on fixed PHP 7.4.33 and 8.3.30 with Go 1.23.12, recording
`diff.legacy_arrays` and seven negative checks per target. It remains compiler
evidence, not a broad WordPress compatibility result.

Do not replace the PHP originals to accommodate the translator. A later adapter
must pin its inputs, run in isolation, record generated/build/differential
outcomes separately, and keep a failed candidate inactive. A translated file
that compiles is not evidence of WordPress compatibility.

## Next bounded steps

1. Six literal-return helpers now pass isolated comparison, including the
   unchanged original empty-array helper. That result is limited to an empty
   `rt.Array` return observed by the fixture, not arbitrary arrays, writes or
   mutations. Select the next dependency-bearing unit with its own explicit
   contract; the neighboring HTTP header helper still requires a separate
   side-effect contract and remains unverified.
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
executable user-facing capabilities. WebMCP support is not claimed.
