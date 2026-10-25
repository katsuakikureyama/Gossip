# Seventeen-original yes/no integration — 2026-10-25

PASS: all149 calls in one generated executable match native PHP8.3.36 and
authored expectations in actual return type/value, ordered keys where
applicable, and caller arguments after each call. Empty stderr/exit0.
Seventeen byte-unchanged WordPress originals and25 owned drivers;
direct103/composed46. This is bounded helper integration, not whole WordPress.

## Exact snapshots and execution

- Gossip test commit `22a5a4d9ab9dac84402e3c22d4d7a91ffa9963f3`, tree
  `6e1cec3c6132026cf623c6dee9d21ddba9db78a7`; branch
  `codex/gohp-yn-helper-bundle` on Draft #23. Staged diff check passed,
  commit succeeded, then the test was launched on that exact snapshot.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`; three
  whole-source hashes and all17 fragment ranges/hashes are retained in the
  [machine record](yn-helper-bundle-2026-10-25.json). New `bool_from_yn()`
  excerpt is functions.php1621–1623, SHA256
  `2fb383818a5529a24fc26b6b9960cd6140872a1aa86347c27741255cfa2c72f3`.
- GoHP unchanged at `49a302e44464623dc4912bafbcf44192759d7878`, tree
  `dc2046b28420a9f7a9da7447052a5e70d6543279`,
  [Draft #44](https://github.com/katsuakikureyama/GoHP/pull/44).
  Newer separate compiler branches are not adopted.
- HostGo1.22.1 darwinarm64; builderGo1.25.14 linuxarm64; PHP8.3.36,
  LC_CTYPE=C. Existing immutable images, same-process PHP guard, non-root,
  read-only, network-disabled and resource-limited execution retained.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressYNHelperBundle$'
go vet ./tests/gohp
```

Initial PASS10.14s/package10.336s; vetPASS. No runtime failure/retry.
Before committing or executing, parent review corrected a draft assertion
that mistakenly required a string return from bool_from_yn to require bool,
and removed a duplicate scalar-normalizer declaration in favor of the existing
helper. Neither draft error reached the executed snapshot. Post-PASS changes
are documentation/evidence only; no repeated runtime verification.
This is local isolated execution, not a new Gossip CI result.

## Corpus and observed distinctions

The119 previous HEIC-bundle cases retain their arguments, expected values and
post-call expectations. All24 direct yes/no cases are added, plus six new
compositions: slice→yes/no; slice→shallow-unescape→yes/no; and a two-result
comparison of the original yes/no and Boolean validators on the same string.
`Y` produces `[true,true]`, while `yes` produces `[false,true]`.

New slice inputs contain only strings/null at every root entry, even
unselected entries; original slice removes null/missing before mapping.
Actual string operands only reach bool_from_yn. Mixed/negative key order,
duplicate selectors, empty input/selection, escaped characters and NUL are
covered. Prior HEIC scalar-only and older nested-list domains remain intact.
See the [contract](../yn-helper-bundle-unit.md). Both observers derive results
and args_after from actual runtime values, never expected fields.

The machine record contains all149 observations and exact source/fragment/
manifest/corpus/driver/observer/compiler/archive/generated/report/binary hashes.

## Reused compiler evidence

Exact unchanged pin previously passed
[CI40159749575](https://github.com/katsuakikureyama/GoHP/actions/runs/40159749575):
all336 checks across52 features plus vet, CLI smoke and artifact publication.
ActualPHP8.3.36/Go1.23.12 linuxamd64, clean exact head/tree; command
`go test -count=1 -json ./cmd/... ./internal/... ./runtime ./tests`.
No compiler change, compiler tests or CI rerun in this change. This earlier
compiler evidence does not replace the fresh149-call Gossip run above.

Artifact `compatibility-php8.3-40159749575-1`, ID12586340125, digest
`sha256:4391dd90f542b6a2a9a5bbe5ff62f4cd63284e07406ce7f6662eaad0f5aa95e2`,
expiry2026-11-24; previously retrieved to `/private/tmp/gohp-in-array-ci.Gn0dnf/`.
Full reused metadata/local commands are in the machine record, traced through
[yes/no standalone evidence](bool-yn-2026-10-25.md). Compiler native schema
observes stdout/normal exit, not general stderr, exceptions or side effects.

## Unchanged and unverified

Eleven standalone configurations and all earlier bundles remain unchanged,
not rerun under old settings. Selected old cases are newly measured in this
combined configuration; it is not every historical standalone input.
WordPress originals, SDK/hooks, oracle admission, prototype and master
README-only state remain unchanged.

General Boolean coercion, image contents, UTF-8 text validation, option-backed
charset wrappers, array_all/arrow closures, nested associative values,
recursive unescaping, whole-file/core/HTTP/DB/theme/plugin behavior remain
unverified. No merge/deploy/release/license decision. Existing local developer
capability remains WebMCP-excluded; auth routes unchanged. Devboard remains
unidentified/unsynchronized, not Done.
