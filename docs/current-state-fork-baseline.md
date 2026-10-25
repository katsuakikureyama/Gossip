# Current state — WordPress fork baseline

Last updated: 2026-10-25 (JST).

## Latest: fifteen-original charset/array integration

- Branch `codex/gohp-charset-helper-bundle`, stacked on #19. Test commit
  `b0835877b9e2172670582a71697f066833837257`, tree
  `27dd48a8bfcbb863b76cb483c04b93f21094c4af`, committed before execution.
- Fifteen byte-unchanged originals from three pinned source files plus
  nineteen owned drivers translate/build once into one binary. All89 calls
  (direct55/composed34) match actual typed returns, ordered keys where relevant
  and caller arguments against PHP8.3.36 and authored expectations.
  Earlier55 associative cases plus28 direct charset cases and6 compositions.
- GoHP `6305b9b486ee233e9aa4378ab8528f5b3d345f90`, tree
  `48c92101a7000794eda13cdfbbd0240b7327afb9` (Draft #41 on #40), explicitly
  merges #40 string/logical comparison and #36 isset/typed-array returns.
  Both parents retained; separate #39 if/else excluded. Reserved-name,
  fixture and spec conflicts preserve both contracts. New cross-stack native
  fixture linked to existing catalog entries;51 features/313 checks.
- Local Gossip comparison first PASS19.04s/package19.279s, empty stderr/exit0,
  vetPASS. Compiler local runtime/frontend/codegen/compat/guards/vet PASS;
  fixed PHP CI40159252743 overall PASS313 checks, actualPHP8.3.36/Go1.23.12,
  exact clean head/tree, vet/CLI smoke/artifacts also successful.
  [Contract](charset-helper-bundle-unit.md) and [evidence](evidence/charset-helper-bundle-2026-10-25.md)
  retain revisions, commands, hashes, observations and pre-execution review
  corrections. No test failure/retry; post-PASS documentation only.
- Prior nine standalone and10/12/13/14-original configurations/evidence stay
  unchanged, not rerun. Original WordPress/SDK/hooks/PHP admission/prototype
  and master README-only state unchanged. No merge/deploy/release/license decision.
- Next: identify the next bounded original-helper dependency independently.
  Public option-dependent charset wrappers, general callbacks/objects/
  references, nested associative transport, map_deep/wp_unslash and whole
  core/HTTP/DB/theme/plugin operation remain unverified. Local developer
  capability WebMCP-excluded, auth routes unchanged; devboard unidentified/
  unsynchronized, not Done.

## Previous: original UTF-8 charset-name predicate

- Branch `codex/gohp-utf8-charset`, stacked on #18. Test commit
  `fccd38e47c4f23ed0bb8b35a0cc974c9486f4d2f`, tree
  `64499078c9dd73381ba28bce7c3cb889392981cc`, committed before execution.
- Unchanged `_is_utf8_charset`, compat.php102–111, executes on GoHP
  `d9f6c7b4f288ef816c4ee3cb6a3458d2226e00d5` (Draft #40 on #38).
  New guarded string-only strcasecmp plus existing !/||/is_string/===.
  Original source is neither rewritten nor replaced with an owned algorithm.
- All28 cases match actual bool returns and post-call inputs against PHP8.3.36
  and authored expectations; empty stderr/exit0. First local comparison
  PASS37.69s/package37.898s, vetPASS. No failure/retry. Post-PASS docs only.
- Compiler native CI40158737961, clean exact head/tree, PHP8.3.36 and Go1.23.12,
  passes all296 checks/49 features plus vet/CLI smoke/artifacts. Parent local
  runtime/codegen/compat and generated guard tests/vet also PASS. Exact
  [contract](utf8-charset-unit.md) and [evidence](evidence/utf8-charset-2026-10-25.md).
- Ninth separate unit, not a fifteen-original bundle. The #36 array-return/
  isset stack is absent from this compiler pin; earlier eight units and all
  bundles remain pinned and unchanged. Original files/SDK/hooks/PHP admission
  and master README-only publication are unchanged.
- Next: integrate the published compiler stacks in a separate change before
  a new fifteen-original bundle. Public is_utf8_charset/defaults/coalescing/
  get_option, _canonical_charset, text encoding conversion and full core/
  HTTP/DB/theme/plugin execution remain unverified. No merge/release/deploy
  or license decision. Local capability remains WebMCP-excluded; auth routes
  unchanged. Devboard unidentified/unsynchronized, not Done.

## Previous: fourteen-original associative bundle

- Branch `codex/gohp-associative-helper-bundle`, stacked on #17. Tested
  snapshot `9c14ce515c266d9db6d1f125a889c9e93cfc3dd1`, tree
  `c5ae3be90a0ed10feb4146f05bac6bb6f1af6879` (exact staged bytes committed
  after execution). Fourteen byte-unchanged originals and sixteen owned
  drivers translate/build once into one executable.
- All 55 cases PASS: previous 33 recursive-bundle cases, fourteen standalone
  slice cases and eight new compositions. Actual return type/value, ordered
  keys where applicable and both caller arguments agree with PHP 8.3.36 and
  authored expectations. Direct27/composed28; empty stderr/exit0. Local test
  8.41s/package8.629s and vet PASS, first comparison, no rerun.
- GoHP pin remains `db3051f8a179078e2227c5d97f50a2f1611d1451` (#36).
  Existing exact-commit CI40156201018 (238 checks/46 features) and earlier
  compiler local tests are explicitly reused, not rerun. Other compiler
  branches are not integrated. Prior eight standalone and ten/twelve/thirteen
  bundle configurations and evidence stay untouched.
- [Contract](associative-helper-bundle-unit.md) and
  [evidence](evidence/associative-helper-bundle-2026-10-25.md) retain exact
  hashes/observations. A pre-commit whitespace check failed on a trailing
  PHP EOF blank line; after successful runtime validation and exact-snapshot
  commit, only that cosmetic line was removed alongside documentation.
  No executable statement/configuration changed, so runtime PASS is reused.
- Next boundary: identify a concrete missing original-helper dependency
  before extending the compiler. Nested associative values, general objects/
  references/callbacks, map_deep/wp_unslash and full core/HTTP/DB/theme/plugin
  operation remain unverified. No implementation merge/release/deploy or
  license decision. Master README-only publication remains untouched.
  Local developer capability remains WebMCP-excluded; auth routes unchanged.
  Devboard unidentified/unsynchronized, not Done.

## Previous: original associative array slice

- Branch `codex/gohp-array-slice-assoc`, based on #15. Test commit
  `1b191a17b574f174e04249b6420e893ec57750c5`, tree
  `c0b85668909beca251efceb4d56351af08c3538e`.
- Exact original wp_array_slice_assoc, functions.php5173–5183 including
  `: array`, translates/builds/executes through pinned GoHP
  `db3051f8a179078e2227c5d97f50a2f1611d1451` ([Draft #36](https://github.com/katsuakikureyama/GoHP/pull/36)
  on #34). Compiler enforces global array returns including fallthrough,
  inherits one-level isset and indexed writes. Separate #33/#35 not included.
- Fourteen explicit cases use ordered root key/value pairs and selector lists.
  Missing/null versus falsy values, key normalization/order/collisions,
  duplicates, negative/NUL/UTF-8 keys, nested lists and both caller arguments
  after invocation are observed. Actual Go/PHP/expected comparisons all PASS,
  empty stderr/exit0, local8.42s/package8.619s/vetPASS. One earlier worker
  invocation failed at host cache setup before any test/Docker/PHP execution;
  first actual comparison passed. No post-PASS code/config change or rerun.
- Exact [contract](array-slice-assoc-unit.md) and
  [evidence](evidence/array-slice-assoc-2026-10-25.md) bind test/compiler/source/
  observer/corpus/report/binary hashes. Compiler local tests/vet and fixed
  CI40156201018 PASS238 checks/46 entries; actualPHP8.3.36/Go1.23.12,
  clean checkout equals pinned commit/tree.
- Eighth independent unit. Earlier seven standalone units and ten/twelve/
  thirteen-original bundles keep their pins/corpora/results; no combined
  fourteen-original test yet. Next: new pinned combined integration.
- Original files, SDK/hooks and PHP admission unchanged. Approved README-only
  master publication from #16 remains intact; implementation is Draft only,
  no merge/release/deployment/license decision. Core/HTTP/DB/themes/plugins,
  general object/reference/callback/map_deep/wp_unslash remain unverified.
  WebMCP-excluded local tooling, no auth changes; devboard unidentified/unsynced.

## Previous: thirteen-original recursive helper bundle

- Branch `codex/gohp-recursive-helper-bundle`, [Draft #15](https://github.com/katsuakikureyama/Gossip/pull/15)
  on #14; test commit `97ed80c3db187ce9ffb2a73c27dd5207f8594192`,
  tree `baef43b9bed5efd39b256d7fe705d93a133fafbf`.
- Previous twelve originals plus unchanged wp_slash and eleven owned PHP
  composition drivers translated/built once into one executable. New33-case
  configuration (direct13/composed20) includes string roundtrips, shallow list
  stripping, capture/newline and shape/keys/values compositions. Every actual
  observation also checks caller arguments after invocation remain unchanged.
- Initial local isolated comparison PASS8.19s/package8.387s, all33 actual
  return types/values/arguments match PHP8.3.36 and explicit expectations;
  empty stderr/exit0, LC_CTYPE=C, affected-package vetPASS. No failure/retry
  or implementation change after PASS. Exact [evidence](evidence/recursive-helper-bundle-2026-10-25.md)
  and [contract](recursive-helper-bundle-unit.md) retain all hashes/boundaries.
- Compiler/runtime unchanged at03d433e73503e32a84f310f023fc0f65c9b188cd
  (GoHP #27), existing exact-tree CI40152066817 PASS181 checks reused, not rerun.
  Earlier seven standalone units and ten-/twelve-function bundles retain their
  pins/corpora/results. Their complete earlier inputs are not relabelled as
  verified in this new combined configuration.
- wp_unslash/map_deep needs element/property mutation, object traversal and
  dynamic call_user_func; none is implemented here. Next: isolate a missing
  compiler contract such as array writes before claiming map_deep support.
  Full-core startup/HTTP/DB/themes/plugins remain unverified. Draft only,
  no merge/deploy/release/license decision. WebMCP-excluded local tooling,
  no auth changes; devboard unidentified/unsynced, not Done.

## Previous: original recursive wp_slash

- Branch `codex/gohp-wp-slash`, [Draft #14](https://github.com/katsuakikureyama/Gossip/pull/14)
  on #13; test commit `d3f02dc3625d331a477cd5fb8655f4d28e864432`,
  tree `2a6af2b6070e1bf36b18bc0cd93d75b484ee1720`.
- Complete unchanged formatting.php5933–5943 translated through compiler/runtime
  `03d433e73503e32a84f310f023fc0f65c9b188cd`,
  [GoHP Draft #27](https://github.com/katsuakikureyama/GoHP/pull/27) on #24,
  incorporating #25 is_array. Adds literal declared-global one-argument callback,
  single-input array_map and actual-string-only addslashes. No general callback
  system, weak coercion, object traversal or arbitrary-depth claim.
- Initial local PHP8.3.36 comparison PASS22 cases: actual return type/value and
  caller input after invocation match PHP and explicit expectations, empty
  stderr/exit0. PASS8.19s/package8.423s; vetPASS. Compiler local test typo caused
  one codegen FAIL, corrected before CI and only affected package rerun; final
  local tests/vetPASS. Fixed CI40152066817 PASS181 checks,42 catalog entries,
  PHP8.3.36/Go1.23.12; clean CI tree equals head.
- [Contract](wp-slash-unit.md) and [evidence](evidence/wp-slash-2026-10-25.md)
  preserve exact commands/revisions/hashes, failure/fix and unverified boundaries.
  Seventh standalone unit; previous six and both bundles unchanged/not rerun.
  No thirteen-function bundle, full WordPress startup, wp_unslash/map_deep,
  HTTP/DB/theme/plugin support. Next: select a concrete remaining original-source
  dependency, or test new composition under a separately pinned configuration.
- Draft only; no merge/deploy/release/license decision. Original files/SDK/hooks,
  PHP admission policy and other GoHP checkout unchanged. WebMCP-excluded local
  tooling, no auth changes; devboard unidentified/unsynced, not Done.

## Previous: twelve-original bundle on the integrated compiler

- Branch `codex/gohp-integrated-helper-bundle`, [Draft #13](https://github.com/katsuakikureyama/Gossip/pull/13)
  stacked on #12. Test commit `74e234dcd80674c03b7f718040e8804866c8b04f`,
  tree `ae04dc19a38c88556f56fc9e4de707a4849f2159`.
- Twelve unchanged WordPress functions plus seven owned PHP drivers are
  translated/built together into one executable. New24 observations:12 direct
  probes and12 compositions (captures→newline→unslash, array passthrough→keys
  or list shape/actual values, and four scalar chains).
- Compiler/runtime `e5ba8c1db81d235cc177cc83699868d812abff7e`,
  [GoHP Draft #24](https://github.com/katsuakikureyama/GoHP/pull/24), stacked
  on #22 and integrating #23. Full builtin reservations/fixtures/guards retained;
  shared semantic validation reconciled. No new runtime domain or PHP policy.
- Initial local Gossip PASS7.60s/package7.834s, actualPHP8.3.36, all24 typed
  values/expectations match, empty stderr/exit0; vetPASS. Compiler local tests
  and vetPASS; fixed [CI40149530122](https://github.com/katsuakikureyama/GoHP/actions/runs/40149530122)
  PASS166 checks across40 catalog entries, Go1.23.12/PHP8.3.36. Clean CI tree
  equals compiler head. No manual retry or implementation change after PASS.
- [Contract](integrated-helper-bundle-unit.md) and
  [evidence](evidence/integrated-helper-bundle-2026-10-25.md) record exact
  commands/pins/hashes. Existing six standalone units and ten-helper bundle
  retain their pins/results, not rerun or relabelled. Original files, SDK/hooks,
  native admission implementation and other session checkout unchanged.
- This verifies one bounded combined configuration, not whole files, recursive
  callbacks/unescaping, core startup, HTTP/DB/themes/plugins. Draft only,
  no default-branch merge/deploy/release/license decision. Local tooling remains
  WebMCP-excluded, auth routes unchanged; devboard unidentified/unsynced, not Done.

## Previous: string-only stripslashes helper

- Branch `codex/gohp-stripslashes-helper`, [Draft #12](https://github.com/katsuakikureyama/Gossip/pull/12)
  stacked on #11. Test commit `f4e65138298ae6763670d457c9c2f74615b8103a`
  translates unchanged `stripslashes_from_strings_only()` at formatting.php2934–2936.
- Compiler/runtime `0e49cc995913dfede0332566dd4dc11c5df16ff8`,
  [GoHP Draft #23](https://github.com/katsuakikureyama/GoHP/pull/23), on #21.
  Full three-arm ternary: condition once, selected arm only, both arms statically
  validated; one-actual-string stripslashes with static/runtime guards.
  Catalog35 entries/89 checks. No short ternary, coalescing, coercion or recursion.
- New21 typed cases matched PHP8.3.36 and expectations, empty stderr/exit0:
  initial PASS7.87s (package8.078s), vet PASS. Includes non-string passthrough
  and escaped strings inside nested lists deliberately left unchanged.
- Compiler local runtime/codegen/catalog tests and vet PASS; fixed
  [CI40121644555](https://github.com/katsuakikureyama/GoHP/actions/runs/40121644555)
  PASS89 checks, Go1.23.12 / PHP8.3.36. Clean CI tree equals compiler head.
  Exact revisions/commands/hashes in [evidence](evidence/stripslashes-helper-2026-10-25.md),
  with [contract](stripslashes-helper-unit.md). No manual rerun after PASS.
- Sixth standalone unit, not a twelve-function bundle or recursive
  `stripslashes_deep` / `wp_unslash` support. Earlier pins/evidence unchanged,
  no reruns; other GoHP stack and normal checkout unchanged. Draft only,
  no merge/deploy/release/license decision; WebMCP-excluded local capability,
  no auth route changes. Devboard card unidentified/unsynced, not Done.

## Previous: newline-preservation helper

- Branch `codex/gohp-newline-helper`, stacked on Draft #10; test commit
  `ff02bbd05a0dd5b75940724cda2352622c543030` translates unchanged
  `_autop_newline_preservation_helper()` from formatting.php lines 811–813.
- Compiler/runtime `dcb72c64d029c7a08a0e94256723ff4076209758`
  ([GoHP Draft #21](https://github.com/katsuakikureyama/GoHP/pull/21), stacked on #18)
  adds three-actual-string `str_replace`, with empty-search handling and static/
  runtime type guards. No array forms, count references, coercion or regex.
  Catalog: 33 entries / 83 checks. Other GoHP stacks/normal checkout unchanged.
- All 16 string-list inputs matched actual string return type/value against
  PHP 8.3.36 and expectations: 7.61s (package 7.789s), empty stderr/exit 0.
  The first command selected no test due to a wrong name and is not PASS evidence;
  the correct selector ran once. Vet PASS. No code changes or reruns after PASS.
- Compiler local runtime/codegen/catalog tests and vet PASS; fixed native
  [CI 40119695048](https://github.com/katsuakikureyama/GoHP/actions/runs/40119695048)
  PASSed all 83 checks on Go 1.23.12 / actual PHP 8.3.36, clean checkout tree
  identical to the compiler pin. Exact commands/revisions and hashes are in the
  [execution record](evidence/newline-helper-2026-10-25.md), with the
  [bounded contract](newline-helper-unit.md).
- This is an additional independent helper, not the full `wpautop` paragraph
  renderer, regex/callback execution, an eleven-function bundle or core startup.
  Earlier standalone/bundle pins and evidence remain untouched. Draft only,
  no merge/deploy/release/license decision; local capability WebMCP-excluded.

## Previous: one combined helper executable

- Branch `codex/gohp-helper-bundle`, stacked on Draft #9. Test commit
  `577791ff93543d8b5abdc1efbc47af99f4bf407c` bundles the ten unchanged originals
  from two source files plus five owned composition drivers, translating and
  building them together into one native Go executable.
- New 22-case integration corpus: ten direct probes and twelve composed probes.
  Actual return types/values matched PHP 8.3.36 and authored expectations, empty
  stderr/exit 0. Initial run PASS 8.89s (package 9.096s); affected-package vet PASS.
- Same GoHP compiler/runtime `e3db685dfb2a3602f94bbfa9df105db3db96cc89` as
  zeroise; no compiler change/CI rerun. Its prior fixed-PHP CI PASS is reused.
  Existing four unit pins/corpora/evidence are untouched and not rerun; their
  73 observations are not claimed as a combined run.
- See [contract](helper-bundle-unit.md) and
  [execution evidence](evidence/helper-bundle-2026-10-25.md). This establishes
  only the new combined helper cases, not full files, core startup, HTTP/DB or
  theme/plugin support. Draft only, no merge/deploy/release/license decision.
  Local developer capability remains WebMCP-excluded; devboard unlinked.

## Previous addition: original zeroise unit

- `codex/gohp-zeroise-helper` is stacked on Draft #8. Test commit
  `ba30252d7509a23f2f8b97de8f0b4ac748947c2b` pins a fourth independent unit,
  extracting original `zeroise()` without changing WordPress source.
- Its compiler/runtime pin is `e3db685dfb2a3602f94bbfa9df105db3db96cc89`
  ([GoHP Draft #18](https://github.com/katsuakikureyama/GoHP/pull/18), on #16).
  Bounded `sprintf('%0Ns', string|int64)` only, canonical width 0–4096;
  invalid static/dynamic inputs fail explicitly. Catalog: 32 entries/78 checks.
- PHP 8.3.36 final-pin comparison PASS: 18 typed string outputs, 10.67s
  (package 10.886s), empty stderr/exit 0. Affected-package vet PASS.
  Compiler runtime/codegen/catalog local tests and vet PASS. Native compiler
  CI initially failed on an unsupported negative literal in the owned fixture;
  fixture/doc-only correction keeps parser scope unchanged. Corrected
  [CI 40118450143](https://github.com/katsuakikureyama/GoHP/actions/runs/40118450143)
  PASSed all 78 checks on fixed Go 1.23.12 / PHP 8.3.36; checkout tree matches pin.
- Prior three manifests retain their previous pin/results, not rerun. This is
  ten selected functions across four units and two pins, not a combined program.
  Separate GoHP #17 work and the normal checkout remain untouched.
- See [unit contract](zeroise-helper-unit.md) and
  [execution evidence](evidence/zeroise-helper-2026-10-25.md). Draft only;
  no merge/deploy/release/license decision or linked devboard Done.

## Current native-oracle policy (supersedes historical runner pins)

New local oracle runs admit only the reviewed official PHP 8.3.36 image in
`gossip/internal/phporacle`, with stable 8.3 >=8.3.36 checked before any fixture
is required in the same PHP process. Other digests are refused before Docker
starts. PHP 7.4, lower patches, prereleases and unreviewed newer series are not
admitted. This applies to hooks, all six standalone GoHP units and both the
ten-helper and integrated twelve-helper bundle experiments.
See [policy](php-oracle-policy.md) and the
[current unified-compiler execution record](evidence/unified-compiler-2026-10-25.md).

The earlier PHP 8.3.13 Gossip results and PHP 7.4/8.3.30 compiler CI below are
historical evidence only, not current-policy passes. Their files are retained
without rewriting observations. GoHP compiler/runtime pins remain distinct and
unchanged for each unit in the previous runner migration. The subsequent
compiler integration below backports only the policy from the separate #14
stack; other features from that stack are not imported.

At `2cac2bf34ce66ef788975e0f2e6644cb815d33b5`, current-policy local execution
PASSed with actual PHP 8.3.36: hook comparison (5 cases, twice each), six
constant-return types/values, slash helpers (13 inputs/26 typed returns), and
Boolean helper (23 inputs). Policy unit tests and affected-package vet also
PASS. Each unit keeps its own compiler pin, and no general compatibility claim
is inferred. Exact commands/hashes/timings for that migration are in the
[prior floor execution record](evidence/php-83-floor-2026-10-25.md).

## Current unified compiler

- The three pre-zeroise manifests pin `4e1b5abfe019ec121e13d85e99e51c206e765533`:
  GoHP Draft #16 stacked on #13, policy backport from `948647c` only. Active
  compiler target is 8.3 and native floor 8.3.36. Legacy-array/rtrim/Boolean
  features are retained; Composer inventory/strlen/array_keys/array_values are
  not imported. Catalog: 31 entries / 73 checks, only active target 8.3.
- Gossip test commit `e943cbce1f74c93b5ca990de7a544caabb274aa5` PASSed all
  three original-source units once on PHP 8.3.36: six constant-return types/values
  (13.37s), slash 13 inputs / 26 typed returns (8.64s), Boolean 23 inputs
  (12.00s), package 34.204s. Same compiler/runtime identity, three separately
  built units; not a single nine-function bundle or full WordPress execution.
- Only three Gossip compiler pins changed. Source, cases, expected values,
  observers, PHP admission and hooks are unchanged; earlier hook/admission
  PASS is reused under the unchanged effective configuration, not rerun.
- Compiler local internal/CLI tests, vet and shell syntax PASS. Native
  [CI 40116685783](https://github.com/katsuakikureyama/GoHP/actions/runs/40116685783)
  PASSed on fixed Go 1.23.12 / PHP 8.3.36. Exact evidence and limitations are
  in the unified-compiler execution record. Historical files remain untouched.

## Repository identity

- `katsuakikureyama/Gossip` is a GitHub fork of `WordPress/WordPress`, repository
  ID `1432300019`, with default branch `master`.
- `katsuakikureyama/Gossip-prototype` is the preserved independent prototype,
  repository ID `1432085879`; its Draft PR #2 and historical commits remain there.
- Fork work uses branch `codex/fork-baseline`, based on upstream commit
  `54106f6c545fc8b881f6b39b7963002bece6c28c` (`7.2-beta1-64693`).
- The follow-up `codex/gohp-return-helpers` branch is stacked on that unmerged
  baseline branch; its Draft PR must target `codex/fork-baseline`, not upstream.
- `codex/gohp-typed-observations` is the next stacked change, targeting
  `codex/gohp-return-helpers`. No default-branch merge is included.
- `codex/gohp-empty-array` is stacked on `codex/gohp-typed-observations` and
  must target that branch, not `master`, WordPress upstream or the independent
  prototype.
- `codex/gohp-slash-helpers` is stacked on `codex/gohp-empty-array` and must
  target that branch. Its Draft PR is not a default-branch merge.
- Local fork checkout: `/Users/kureyamayuri/Desktop/Gossip-wordpress`.
- `codex/gohp-boolean-helper` is stacked on `codex/gohp-slash-helpers`.
  [Draft PR #6](https://github.com/katsuakikureyama/Gossip/pull/6) targets #5.
- `codex/php-83-oracle-floor` is stacked on `codex/gohp-boolean-helper`;
  [Draft PR #7](https://github.com/katsuakikureyama/Gossip/pull/7) targets #6,
  not the default branch or WordPress upstream.
- `codex/unified-php83-compiler` is stacked on `codex/php-83-oracle-floor`;
  [Draft PR #8](https://github.com/katsuakikureyama/Gossip/pull/8) targets #7.
  The older `/Users/kureyamayuri/Desktop/Gossip` checkout still points to
  `Gossip-prototype`; it is not the new fork. No existing directory was removed.

## Historical implementation and verification (prior oracle policy)

- Boolean-helper translation PASS at `8901d634e4d52e9be06c4c2d7f91714652974b28`:
  GoHP `14bb01053947f1bf2fb32edaff8085feb7227f48` (Draft #13, based on #10)
  translates unchanged `wp_validate_boolean()` from functions.php 7932–7942.
  All 23 mixed inputs match PHP 8.3.13 in actual bool return type/value and
  authored expectations; empty stderr/exit 0. PASS 8.38s (package 8.557s).
  The initial attempt failed its PHP locale guard after generation/build;
  the observer and compiler native harness now select and verify `LC_CTYPE=C`
  in the oracle process. WordPress code was never changed. This is a bounded
  one-arm conditional/short-circuit/scalar validation contract, not full PHP.
  See [Boolean evidence](evidence/boolean-helper-2026-10-25.md) for exact pins,
  failed attempt, corrected execution, compiler checks and limitations.
- GoHP #13 final head `4fabe11d2063b03fced8d0aabd62bbd20a10fb88` passes
  [fixed PHP 7.4.33/8.3.30 CI 40113520430](https://github.com/katsuakikureyama/GoHP/actions/runs/40113520430).
  It differs from the Gossip pin only in native-test evidence logging, fixing
  strict summary parsing without changing compiler/runtime implementation.
  Prior CI failures remain recorded; they are not presented as overall PASS.
- Slash-helper translation PASS at `a97f14f42e6f59049062d59695c74da9d56ee0d2`:
  pinned GoHP `ceee0fc71e1e7947f1c78f1dd18875236b9a180b` (Draft GoHP PR #10,
  based on `feat/legacy-array-syntax`) translates the unchanged original
  `trailingslashit()` and `untrailingslashit()` snippets from
  `wp-includes/formatting.php` lines 2883–2885 and 2898–2900. The single
  affected opt-in test PASS took 18.82 seconds (package 19.018 seconds): PHP
  8.3.13 and generated Go matched all 13 explicit string inputs and their 26
  typed return values, with empty stderr and exit 0. The inputs include empty,
  slash/backslash-only, repeated suffixes, Windows-looking paths, whitespace,
  Japanese text, URL-looking text, and embedded NUL/LF bytes. This is a
  bounded `rtrim` contract only: two arguments and a literal byte mask `/\\`,
  not ranges, non-string coercion, dynamic masks, general path handling, or
  full WordPress formatting compatibility. See
  [slash-unit evidence](evidence/slash-helpers-2026-10-25.md).
- GoHP PR #10 CI also PASSed for fixed PHP 7.4.33 and 8.3.30 with Go 1.23.12:
  [run 40112678059](https://github.com/katsuakikureyama/GoHP/actions/runs/40112678059)
  records `diff.rtrim`, nine negative checks, and two runtime checks as PASS
  for each target, with artifacts saved. This is compiler-project evidence for
  the bounded `rtrim` implementation, not full WordPress compatibility; this
  fork did not rerun its own test in that CI job.
- Empty-array translation PASS at `1596c5ac35846db59612098fe13d4f5736f23655`:
  pinned GoHP `c67fbefe34b1b098c0a9f475504d06bd014de23c` (Draft GoHP PR #8,
  based on `feat/compatibility-views`) translates six unchanged original
  helpers, including `__return_empty_array()` with its original `array()`
  expression. The single opt-in test PASS took 13.14 seconds (package 13.328
  seconds): PHP and generated Go matched five scalar echo observations and all
  six typed values/types, with empty stderr and exit 0. The echo oracle remains
  deliberately limited to the five scalar functions because PHP array coercion
  is not type-preserving. The Go observer accepts only an empty `rt.Array` and
  does not claim general array equivalence. See
  [current v3 evidence](evidence/empty-array-2026-10-25.md).
- GoHP PR #8 CI also PASSed for fixed PHP 7.4.33 and 8.3.30 with Go 1.23.12:
  [run 40112122164](https://github.com/katsuakikureyama/GoHP/actions/runs/40112122164)
  records `diff.legacy_arrays` and seven negative checks as PASS for each
  target, with artifacts saved. This is compiler-project evidence, not a
  full WordPress compatibility result; this fork did not rerun its own test.
- Historical typed-return comparison PASS at `b0f0ce82d0532e1ff0258478b95a158c29b84efc`:
  five original helpers match PHP in actual scalar type/value, stdout, empty
  stderr and exit 0. A separate unchanged `__return_empty_array()` returns an
  empty array in PHP but was rejected by the old GoHP pin `9c7389d` (`array()`
  was unsupported); no generated source/report was left. This historical boundary
  is superseded only for the new pin above; the old run is not relabeled. See
  [v2 evidence](evidence/typed-helpers-2026-10-25.md).
- First original-source GoHP unit PASS at
  `1c84d0c736b682b81f3a267f4c5b2c6269461308`: five unchanged WordPress
  constant-return helpers were extracted, generated, hash-verified, compiled
  and compared against PHP. Stdout, empty stderr and exit 0 matched. The first
  attempt failed before translation due to a 16 MiB builder tmpfs; the corrected
  256 MiB configuration passed. See [evidence](evidence/return-helpers-2026-10-25.md).
  That earlier result was echo-only; the later typed probe above covers these
  five constants, not general return-type equivalence or full-file/core translation.
- Original WordPress files/history retained. Changes only add the nested
  `gossip/` Go experiment and documentation; no inherited runtime, asset or
  license file is modified.
- Native Plugin SDK and five-case fixture copied unchanged from prototype
  `1b9c05b0557bcca3f0a4f5eebf046cbef67d18e2`.
- The comparison now runs verified source bytes from **this fork**, not the
  WordPress installation bundled in a container image. A manifest pins the
  upstream commit, version and five source/license SHA-256 values.
- At `a167e6c8c56294f2d702e7c25342c614bde364dd`, race-enabled tests PASS:
  native SDK tests, source baseline, and five owned hook comparisons (two
  invocations each) against PHP 8.3.13. See
  [exact commands and exclusions](evidence/fork-hooks-2026-10-25.md).

## Unimplemented / unverified

- Full-file/core GoHP translation or execution, themes or plugins. The only
  translated exceptions are six tiny constant-return helpers at their
  separately pinned `c67fbefe34b1b098c0a9f475504d06bd014de23c` evidence and
  the two slash helpers at `ceee0fc71e1e7947f1c78f1dd18875236b9a180b`; they
  were not tested together as a larger unit. The inspected GoHP subsets do not
  cover the original hook engine; see [reuse boundary](reuse-boundary.md).
- General WordPress hook parity: dynamic registration/removal, references,
  recursion, multiple arguments, negative priorities and core filter stacks.
- A Go server serving WordPress, PHP/native plugin bridging, block/theme
  rendering, original database mapping, production auth and CGI execution.
- Full-site runtime or browser validation of this development snapshot. The
  prototype's old WordPress 6.6.2 tests are not evidence for this fork baseline.
- Stable release selection, upstream-source update automation, production
  migration, public deployment, license finalization or default-branch merge.

## Next steps / migration gates

Use the original WordPress source as the behavioral reference. The Boolean and
slash units are separate bounded contracts from the six constant-return helpers; do not
combine their evidence or extrapolate it to general arrays, general `rtrim`,
paths, or core translation. Select the next dependency-bearing unit only with
its own explicit contract. The nearby `send_nosniff_header()` needs `header()`
and an HTTP side-effect contract; it was inspected only, not runtime-tested or
supported. Keep source pins, isolated generation and per-unit differential
results. Do not grow the independent CMS or replace WordPress database
semantics with the prototype JSON store.

Production work remains gated on owner approval, source/version inventory,
restorable data copy, dry run, compatibility validation and rollback evidence.
This baseline creates no new browser or agent operations; local test tooling
is explicitly excluded from WebMCP publication.

Track the baseline in Draft PR #1 in **this fork**, based on `master`; stack the
helper-unit Draft PR #2 on `codex/fork-baseline` until that PR is merged; stack
typed observations on that helper branch, then stack the empty-array change on
typed observations, the slash-helper change on empty-array, and Boolean on slash.
No devboard card
is linked in this repository; none is marked Done by this report. Neither the
fork's default branch nor WordPress upstream is modified by the work branch.
