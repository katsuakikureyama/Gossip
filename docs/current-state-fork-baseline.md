# Current state — WordPress fork baseline

Last updated: 2026-10-25 (JST).

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
- Local fork checkout: `/Users/kureyamayuri/Desktop/Gossip-wordpress`.
  The older `/Users/kureyamayuri/Desktop/Gossip` checkout still points to
  `Gossip-prototype`; it is not the new fork. No existing directory was removed.

## Implemented and verified

- Typed-return comparison PASS at `b0f0ce82d0532e1ff0258478b95a158c29b84efc`:
  five original helpers match PHP in actual scalar type/value, stdout, empty
  stderr and exit 0. A separate unchanged `__return_empty_array()` returns an
  empty array in PHP but is explicitly rejected by pinned GoHP (`array()` is
  unsupported); no generated source/report is left. This is a verified boundary,
  not sixth-function support. See [new evidence](evidence/typed-helpers-2026-10-25.md).
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

- Full-file/core GoHP translation or execution, themes or plugins. Five tiny
  extracted constant-return helpers are the sole translated exception. The
  inspected GoHP candidate is now `9c7389d35032d623a429198a69cdfa5c11571e8e`,
  not the earlier prototype report's PR #2 ref. Its supported subset does not
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

Use the original WordPress source as the behavioral reference. The first tiny
source unit now has typed observations. The next concrete compiler boundary is
the unchanged empty-array helper: resolve `array()` support in a separately
reviewed GoHP change before updating the pin and attempting a positive comparison.
The nearby `send_nosniff_header()` needs `header()` and an HTTP side-effect
contract; it was inspected only, not runtime-tested or supported. Keep source
pins, isolated generation and per-unit differential results. Do not grow the independent CMS or
replace WordPress database semantics with the prototype JSON store.

Production work remains gated on owner approval, source/version inventory,
restorable data copy, dry run, compatibility validation and rollback evidence.
This baseline creates no new browser or agent operations; local test tooling
is explicitly excluded from WebMCP publication.

Track the baseline in Draft PR #1 in **this fork**, based on `master`; stack the
helper-unit Draft PR #2 on `codex/fork-baseline` until that PR is merged; stack
typed observations on that helper branch. No devboard
card is linked in this repository; none is marked Done by this report. Neither
the fork's default branch nor WordPress upstream is modified by the work branch.
