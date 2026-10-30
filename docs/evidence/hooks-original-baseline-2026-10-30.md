# Original hook API baseline — 2026-10-30

**Native reference recorded; generated source rejected.** This series is a
native PHP reference and whole-source admission probe, not generated hook
compatibility or G83-BOOT acceptance. The [machine record](hooks-original-baseline-2026-10-30.json)
retains the actual native output and compiler diagnostic separately.

One clean configuration at Gossip `4e5e87879232607a13fe8633467aad4d878fb609`,
tree `cd1d274e944797a52d81aca4c8546acbd259b6a7`, pinned GoHP
`b2d1482268a7710c515c07bf919af3d40b09d17c`: test1.38s/package1.620s PASS.
All eight original-PHP workflows execute with exit0/stderr0. The single GoHP
probe exits1/stdout0, with **PHP002 at `plugin.php:25`, unsupported `require`**.
No Go source/report or other candidate artifacts are published. This expected
rejection is not a successful translation.

Native stdout43242bytes/SHA
`888257c7973e14db138d3a366e998dc4f6fb9ad5d5791f645abb377ad7599190`;
canonical typed workflows SHA
`548fdccd65e84bca83ef231effc0ec5e416e5dfae05a71e5ab5d7199cf0595d6`.
Raw attempt SHA
`0be81441d7243aaa84a4742a5ff9ef12e7f9768c874daecc4fbb2d0ad0623385`.
The public record preserves those observations and adds clean-source/command/
run_count1 metadata; it is not another execution.

## Source and observation contract

The [manifest](../../gossip/tests/gohp/hooks-source-closure.json) binds complete,
unchanged `plugin.php`, `class-wp-hook.php`, `class-wp-filter-sentinel.php`, the
original notice and version provenance. The only static file edges are original
`plugin.php:25-26` → the two classes. No declarations are extracted, rewritten,
stubbed or replaced. PHP's standard interfaces/functions supply the selected
workflows; other plugin-management functions retain external WordPress
dependencies and are not exercised. This is not the full bootstrap closure.

The owned [PHP observer](../../gossip/tests/gohp/hooks-original-observer.php)
requires original `plugin.php` once. It supplies an explicit empty `$wp_filter`
before loading and clears the four hook globals only between independent
workflows. Preinitialized hook migration and undefined-global startup are not
claimed. Reflection checks original function/class files. Callbacks and bounded
recursion guards are fixture inputs; hook scheduling and mutation are performed
only by original WordPress code.

| Workflow | Selected native behavior |
|---|---|
| `no_callbacks` | Empty filter/action return, counters and current-hook state |
| `priority_order` | Numeric priority and insertion order within a priority |
| `accepted_arguments` | Zero/one/two accepted args and propagation of callback returns |
| `callback_identity_removal` | Same/distinct closure registration, removal and lookup without exposing unstable object IDs |
| `actions` | Action args, counters/current action, ignored callback return |
| `reentrant_filter` | Same hook called from its callback, nested iteration state and cleanup |
| `mutation_during_iteration` | Remove a pending priority and add a new priority during original iteration |
| `sentinel_identity` | Original marker object preserved/replaced, distinct from null and false |

The [fixed case list](../../gossip/tests/gohp/hooks-original-cases.json),
[pin](../../gossip/tests/gohp/hooks-original-baseline.json) and
[harness](../../gossip/tests/gohp/hooks_original_baseline_test.go) have separate
identities from parser corpora. Each observation records initial state, actual
result, callback events and final state with ordered typed arrays/scalars and
base64 strings. Hook state includes public priority/callback counts and read-only
reflection of private iteration/priorities/nesting/action fields. Callback IDs
are intentionally not raw process-specific values; identity outcomes are booleans.
None of these observations are invented generated-Go expectations.

Actual callback order is10-first →10-second →20; accepted-arg counts are0/1/2.
Removing the pending20 callback and inserting15 during execution produces
10-mutator →15-added →30-tail. Reentry returns `inner|head|tail` and
`outer|head|tail`, with two filter invocations and cleaned iteration/global
stacks afterwards. The public `current_priority()` observation during the inner
tail reports10 while the outer invocation is at10; the record preserves this
original behavior rather than inferring the callback's registered20 priority.
Sentinel pass-through preserves identity, a new sentinel replaces it, and
null/false replacements remain distinguishable.

## Admission and next implementation work

The probe passes the complete original closure to pinned GoHP, without changing
bare `require` to `require_once` or skipping unsupported declarations. A rejected
probe must return an explicit file-bound compiler diagnostic and publish no Go
source/report. Native capture and source rejection remain separate outcomes;
test PASS only means this baseline was recorded as declared. Generated hook
build/execution and differential comparison remain not-run.

Source inspection identified these additional prerequisites; only the actual
first diagnostic will be called an observed compiler stopping point:

| Original source | Required semantics beyond small builtin additions |
|---|---|
| `plugin.php:25-57` | Original bare loading and execution-wide globals/initialization; declaration-only loading is insufficient |
| `plugin.php:175,499`; `WP_Hook:367-371` | Variadics, dynamic callable dispatch, argument lists and callback identity |
| `class-wp-hook.php:26-76`; sentinel class | Attributes/final/interfaces/visibility/default arrays/strict-types contracts |
| `WP_Hook:130-191,339-381,408` | References, ordered mutable callback tables, internal array pointers, reentrancy and iteration mutation |

These rows are an implementation plan, not supported features or proof that
fixing the first syntax alone makes hooks work. The native reference provides
bounded behavior for future implementation. General references/callable forms,
all-hook observers, same-priority mutation variants, exceptions, arbitrary
plugin management, full bootstrap and application acceptance remain unmeasured.

## Execution and capability

Run from `gossip/` with explicit opt-in and fixed paths/image:

```sh
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressHooksBaseline$'
```

`GOSSIP_GOHP_HOOKS_BASELINE=1`, absolute `GOSSIP_GOHP_HOOKS_SOURCE`, fresh absolute
`GOSSIP_GOHP_HOOKS_EVIDENCE_OUT` and exact `GOSSIP_PHP_RUNNER_IMAGE` are required.
Missing opt-in skips without evidence or support claims. Native profile is fixed
PHP8.3.36/Linux-aarch64/64-bit/C/JIT0/backtrack100000/recursion10000/memory128M.
The admission toolchain is existing offline Go1.23.0/Darwin-arm64/CGO0; no native
Go hook binary is produced. Existing running Docker/cache/image only; no setup,
dependency install, pull, daemon startup or deployment. Parser/resource/JSON
successes are reused without rerunning them.

Local stable capability `gossip.tests.hooks-original-baseline` is experimental.
Inputs are the opt-in/absolute source/exact image/fresh evidence path plus pinned
corpus/observer/manifest. Inspect/prepare verify immutable sources and profile;
save exclusively writes one local observation. There is no application commit
or publish operation. Results distinguish native-reference, compiler rejection,
infrastructure failure, skip and absent generated equivalence. WebMCP is
**excluded** because this is local native fixture/compiler testing, with no UI
callsite, browser/server operation, auth route or deployment change. Issue1 and
current-state track scope6522006068; devboard remains unidentified.
