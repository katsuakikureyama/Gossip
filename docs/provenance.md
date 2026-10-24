# Source provenance

## WordPress baseline

- Parent and source repository: `WordPress/WordPress`.
- Baseline commit: `54106f6c545fc8b881f6b39b7963002bece6c28c`.
- Version declared by that source: `7.2-beta1-64693` (development snapshot).
- `gossip/wordpress-source.json` records SHA-256 for the exact PHP hook engine,
  version file and inherited license. Tests refuse source drift.
- All inherited WordPress files and history are retained. This change adds
  `README.md`, `docs/` and `gossip/` only; it neither replaces nor rewrites the
  WordPress core, bundled themes/assets, `readme.html` or `license.txt`.

Do not silently update the pin to make a failing test pass. Review upstream
changes and establish new evidence when moving the source baseline. This fork
does not imply endorsement, stable release status or successful Go conversion.
The inherited license/copyright notices remain in place; no licensing
finalization or separate permissive grant is made for the experimental module.

## Prototype-derived additions

Source repository:
[`katsuakikureyama/Gossip-prototype`](https://github.com/katsuakikureyama/Gossip-prototype/tree/1b9c05b0557bcca3f0a4f5eebf046cbef67d18e2),
commit `1b9c05b0557bcca3f0a4f5eebf046cbef67d18e2`.

- `pkg/pluginsdk/sdk.go`, `pkg/pluginsdk/sdk_test.go` and
  `tests/compat/hooks-v1.json` are copied unchanged under `gossip/`.
- `tests/compat/hooks_test.go` and `hooks-oracle.php` are adapted: new nested
  module path, fork-source staging/hashing, and a PHP-runner variable replace
  the old image-bundled WordPress 6.6.2 oracle assumption.
- `gossip/tests/compat/source_test.go` and `wordpress-source.json` are new.

The prototype's old commit hashes and reports remain with that repository;
they are not reassigned to this fork. Only new results recorded against this
fork count toward this fork's compatibility evidence. The custom CMS runtime,
editor, local JSON persistence/import and release CLI are deliberately not
copied into the WordPress runtime.

The native fixture uses `NOASSERTION` as test manifest metadata; this is not
a licensing decision for WordPress or a distributable plugin.

## Original helper-unit extraction

`gossip/tests/gohp/return-helpers.json` pins the entire original
`wp-includes/functions.php`, five complete function line ranges, GoHP commit and
runner image digests. The test copies those original excerpts unchanged into a
temporary PHP bundle, adding only the opening tag, separator newlines and owned
observation calls. The original license is staged alongside it. It does not
commit rewritten or generated replacements into the WordPress tree.

The GoHP compiler and runtime come from a `git archive` of the same pinned
commit, excluding the checkout's working edits and `.git` credentials/config.
Exact excerpt, bundle, archive, compiler, generated source/report and executable
hashes are in [the unit observation](evidence/return-helpers-2026-10-25.json).
