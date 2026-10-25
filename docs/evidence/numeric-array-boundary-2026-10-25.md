# Numeric-array native reference and compiler boundary — 2026-10-25

**Native reference PASS, GoHP candidate REJECTED.** No generated candidate
binary exists and no Go/PHP differential test has passed for this helper.
The earlier seventeen-original/149-case integration remains the latest successful
combined result; this work does not increase that count.

## Exact snapshot and execution

- Gossip branch `codex/gohp-numeric-array-boundary`, based on Draft #24 head
  `aed73f1398799d2c6f7cd6e204db979f528df7ab`.
- Test commit `90870c68d01c9d399dcd6060b85ce38363096373`, tree
  `3d5e6e11af8f839beb373c916f683f1ea92c4011`. Staged diff check and commit
  succeeded before the first execution; the checkout was clean.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`.
  `functions.php:5448–5454` is the unchanged predicate; `compat.php:623–644`
  is the complete unchanged guarded polyfill. Its declaration at635–643
  is also submitted independently to GoHP, without changing indentation.
- GoHP unchanged at `49a302e44464623dc4912bafbcf44192759d7878`, tree
  `dc2046b28420a9f7a9da7447052a5e70d6543279` (Draft #44). One compiler
  build serves both rejection probes. No compiler source changes, compiler
  test reruns, CI dispatch or adoption of other compiler branches occurred.
- Host Go1.22.1 darwin/arm64; builder Go1.25.14 linux/arm64; actual PHP8.3.36.
  Immutable images, same-process PHP admission guard, C locale and existing
  non-root/read-only/network-disabled/resource limits are retained.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressNumericArrayBoundary$'
go vet ./tests/gohp
```

First run boundary-test PASS7.60s/package7.839s; vetPASS, command exit0.
No failure/retry. After PASS, only documentation/evidence changed and the
same configuration was not rerun. This is local execution, not a new CI result.

## Native observations

All24 authored cases matched actual Boolean return type/value, list status and
typed ordered caller input after the call. Seven scalar inputs return false;
17 arrays exercise empty/list/sparse/negative/out-of-order keys, PHP's canonical
numeric-string key normalization, retained string keys, mixed ordering and
null/false/zero/empty-string values. All values are flat scalars; integer values
and keys stay within ±1,000,000. Native stderr is empty and exit is0.

The PHP observer confirms the active `array_all` is user-defined, not a newer
native builtin. The source bundle contains the original WordPress compatibility
implementation and guard; there is no stub. Sparse and negative-key arrays
demonstrate why this predicate must not be replaced by `array_is_list`.

## Actual compiler rejection

Each declaration is translated as an independent unchanged source fragment:

| Original | Exit | First diagnostic | Generated Go / report |
| --- | --- | --- | --- |
| `wp_is_numeric_array` | 1 | `PHP002 line 2: only array return declarations are supported (at "{")` | Neither exists |
| `array_all` | 1 | `PHP002 line 2: expected var (at "array")` | Neither exists |

Line2 refers to the extracted translation unit, not the WordPress file.
Stdout is empty for both. The harness checks actual CLI exit/stderr and absence
of both artifacts; an infrastructure failure is not accepted as a rejection.
These first-error results do not dynamically prove every later blocker. Source
review additionally identifies the arrow callback, callable parameter and
dynamic value/key invocation. The full guarded block was executed only in PHP,
not submitted as a combined GoHP translation unit in this step.

The [machine record](numeric-array-boundary-2026-10-25.json) stores all24 actual
observations, exact diagnostics, source/fragment/manifest/corpus/observer/
compiler/archive hashes, commits, commands and distinct statuses:
`native_reference_pass_candidate_rejected`, `differential_status: not-run`.
Existing compiler CI evidence is referenced as historical provenance only.

## Remaining work

The [contract](../numeric-array-boundary.md) specifies the next Boolean-return,
typed-parameter and callback dependency work before a true original-source
differential unit can pass. Neither successful generation nor a full-core
compatibility result is claimed. Older standalone/bundle pins, corpora and
evidence were not modified or rerun. Original WordPress files, SDK/hooks,
PHP admission, prototype and master are unchanged.

Local developer capability remains explicitly WebMCP-excluded; no auth route,
UI action or publication/deployment API is added. No merge/deploy/release/license
decision. Devboard is unidentified/unsynchronized, not Done.
