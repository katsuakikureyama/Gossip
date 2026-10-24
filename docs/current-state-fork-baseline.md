# Current state — WordPress fork baseline

Last updated: 2026-10-25 (JST).

## Current native-oracle policy (supersedes historical runner pins)

New local oracle runs admit only the reviewed official PHP 8.3.36 image in
`gossip/internal/phporacle`, with stable 8.3 >=8.3.36 checked before any fixture
is required in the same PHP process. Other digests are refused before Docker
starts. PHP 7.4, lower patches, prereleases and unreviewed newer series are not
admitted. This applies to hooks and all three GoHP translation units.
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

- All three manifests now pin `4e1b5abfe019ec121e13d85e99e51c206e765533`:
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
