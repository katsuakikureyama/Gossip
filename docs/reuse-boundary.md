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
| PHP-to-Go translation | Separate GoHP project | Five original constant-return helpers pass output and typed-return comparison; original empty-array helper is explicitly rejected; full WordPress core and the hook engine remain unsupported |
| Earlier Go HTTP/editor/import/release prototype | Separate `Gossip-prototype` repository | Historical prototype only; not installed or activated by this fork |

## GoHP readiness check — 2026-10-25 JST

The inspected candidate is
[`feat/source-maps` at `9c7389d35032d623a429198a69cdfa5c11571e8e`](https://github.com/katsuakikureyama/GoHP/tree/9c7389d35032d623a429198a69cdfa5c11571e8e),
an unmerged draft stack ending in PR #5. Its documented closed subset includes
basic scalar operations, named functions, literal-array reads/foreach and
limited classes. This is not a stable WordPress translator interface.

The original hook engine requires features outside that subset: include/require,
global state, references, array mutation, dynamic callbacks, variadic arguments,
interfaces and other PHP builtins/object behavior. See the pinned GoHP
[compatibility report](https://github.com/katsuakikureyama/GoHP/blob/9c7389d35032d623a429198a69cdfa5c11571e8e/docs/compatibility.md)
and [integration contract](https://github.com/katsuakikureyama/GoHP/blob/9c7389d35032d623a429198a69cdfa5c11571e8e/docs/integration.md).
No full-file translation attempt or deployable WordPress artifact is claimed.
The later [helper-unit experiment](translation-units.md) converts only five
unchanged literal-return functions and records its own narrow evidence.

Do not replace the PHP originals to accommodate the translator. A later adapter
must pin its inputs, run in isolation, record generated/build/differential
outcomes separately, and keep a failed candidate inactive. A translated file
that compiles is not evidence of WordPress compatibility.

## Next bounded steps

1. The first five literal-return helpers now pass isolated output and typed-return
   comparison. The original empty-array helper is the next measured boundary:
   its `array()` call is rejected. Any support belongs in a separately reviewed
   GoHP change, followed by a pin update and differential test, never a rewrite
   of the WordPress original. The neighboring HTTP header helper requires a
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
executable user-facing capabilities. WebMCP support is not claimed.
