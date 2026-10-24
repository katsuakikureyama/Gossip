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
- Local fork checkout: `/Users/kureyamayuri/Desktop/Gossip-wordpress`.
  The older `/Users/kureyamayuri/Desktop/Gossip` checkout still points to
  `Gossip-prototype`; it is not the new fork. No existing directory was removed.

## Implemented and verified

- First original-source GoHP unit PASS at
  `1c84d0c736b682b81f3a267f4c5b2c6269461308`: five unchanged WordPress
  constant-return helpers were extracted, generated, hash-verified, compiled
  and compared against PHP. Stdout, empty stderr and exit 0 matched. The first
  attempt failed before translation due to a 16 MiB builder tmpfs; the corrected
  256 MiB configuration passed. See [evidence](evidence/return-helpers-2026-10-25.md).
  This does not prove return-type equivalence or full-file/core translation.
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
source unit is verified; next choose a bounded dependency-bearing unit or strengthen
the return-value observation before expanding the compatibility claim. Keep source
pins, isolated generation and per-unit differential results. Do not grow the independent CMS or
replace WordPress database semantics with the prototype JSON store.

Production work remains gated on owner approval, source/version inventory,
restorable data copy, dry run, compatibility validation and rollback evidence.
This baseline creates no new browser or agent operations; local test tooling
is explicitly excluded from WebMCP publication.

Track the baseline in Draft PR #1 in **this fork**, based on `master`; stack the
helper-unit follow-up on `codex/fork-baseline` until that PR is merged. No devboard
card is linked in this repository; none is marked Done by this report. Neither
the fork's default branch nor WordPress upstream is modified by the work branch.
