# Gossip fork documentation

The inherited WordPress source remains authoritative for WordPress behavior.
These documents describe the experimental Gossip additions, not a production
support or compatibility promise.

| Document | Purpose |
| --- | --- |
| [Current state](current-state-fork-baseline.md) | Implemented, tested, unverified and next steps |
| [Newline preservation unit](newline-helper-unit.md) | Original `wpautop` callback helper with bounded string replacement; not paragraph-renderer support |
| [Newline preservation evidence](evidence/newline-helper-2026-10-25.md) | Separate 16-case observation and pinned compiler evidence |
| [Combined helper unit](helper-bundle-unit.md) | Ten unchanged originals plus five owned composition drivers in one generated binary; 22-case integration contract |
| [Combined helper evidence](evidence/helper-bundle-2026-10-25.md) | Separate integration result; does not replace the four standalone unit records |
| [Zeroise-helper translation unit](zeroise-helper-unit.md) | Separate unchanged `zeroise()` contract with bounded string/integer zero-padding |
| [Zeroise-helper translation evidence](evidence/zeroise-helper-2026-10-25.md) | New unit execution record; earlier three-unit compiler pins remain unchanged |
| [Reuse boundary](reuse-boundary.md) | Original source versus Go experiment and GoHP blockers |
| [Provenance](provenance.md) | Exact upstream/prototype references and unchanged source policy |
| [PHP oracle policy](php-oracle-policy.md) | Admission floor for new Docker-oracle runs; not a compatibility conclusion |
| [PHP 8.3 floor evidence](evidence/php-83-floor-2026-10-25.md) | Previous PHP 8.3.36 execution record, retained with its earlier three compiler pins |
| [Unified-compiler evidence](evidence/unified-compiler-2026-10-25.md) | Current PHP 8.3.36 execution record for three independent GoHP units under one compiler pin |
| [Hook comparison evidence](evidence/fork-hooks-2026-10-25.md) | Historical PHP 8.3.13 execution record: commit, runner, source/fixture hashes and exclusions |
| [Original-source translation units](translation-units.md) | Extraction, isolation, GoHP report verification and comparison contract |
| [Return-helper translation evidence](evidence/return-helpers-2026-10-25.md) | Historical PHP 8.3.13 record: five original helpers generated/built/compared; exact hashes and limitations |
| [Typed-return and rejection evidence](evidence/typed-helpers-2026-10-25.md) | Historical PHP 8.3.13 v2 result: five actual return types/values matched; original empty-array helper was explicitly rejected |
| [Empty-array translation evidence](evidence/empty-array-2026-10-25.md) | Historical PHP 8.3.13 v3 result: six actual return types/values matched, including the unchanged original empty-array helper |
| [Slash-helper translation unit](slash-helper-unit.md) | Isolated contract for unchanged `untrailingslashit()` and `trailingslashit()` |
| [Slash-helper translation evidence](evidence/slash-helpers-2026-10-25.md) | Historical PHP 8.3.13 isolated result: 13 string inputs and 26 typed return values matched |
| [Boolean-helper translation unit](boolean-helper-unit.md) | Separate contract for unchanged `wp_validate_boolean()` across 23 typed inputs |
| [Boolean-helper translation evidence](evidence/boolean-helper-2026-10-25.md) | Historical PHP 8.3.13 record for the Boolean contract; independent of earlier helper evidence |

Start each work session with current state and this index. Keep new work in an
explicit branch/PR in `katsuakikureyama/Gossip`. Do not replace the root tree with
the prototype or merge to `master` without owner approval.
