# Gossip fork documentation

The inherited WordPress source remains authoritative for WordPress behavior.
These documents describe the experimental Gossip additions, not a production
support or compatibility promise.

| Document | Purpose |
| --- | --- |
| [Current state](current-state-fork-baseline.md) | Implemented, tested, unverified and next steps |
| [Reuse boundary](reuse-boundary.md) | Original source versus Go experiment and GoHP blockers |
| [Provenance](provenance.md) | Exact upstream/prototype references and unchanged source policy |
| [Hook comparison evidence](evidence/fork-hooks-2026-10-25.md) | Executed commit, runner, source/fixture hashes and exclusions |
| [Original-source translation units](translation-units.md) | Extraction, isolation, GoHP report verification and comparison contract |
| [Return-helper translation evidence](evidence/return-helpers-2026-10-25.md) | Five original helpers generated/built/compared; exact hashes and limitations |
| [Typed-return and rejection evidence](evidence/typed-helpers-2026-10-25.md) | Five actual return types/values matched; original empty-array helper explicitly rejected |

Start each work session with current state and this index. Keep new work in an
explicit branch/PR in `katsuakikureyama/Gossip`. Do not replace the root tree with
the prototype or merge to `master` without owner approval.
