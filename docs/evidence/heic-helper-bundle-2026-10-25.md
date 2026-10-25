# Sixteen-original HEIC/charset/array integration — 2026-10-25

PASS: all119 calls in one generated executable match native PHP8.3.36 and
authored expectations in actual return type/value, ordered keys where
applicable, and caller arguments after each call. Empty stderr/exit0.
Sixteen byte-unchanged WordPress originals and22 owned composition drivers;
direct79/composed40. This is MIME-label and charset-name classification,
not image-content inspection, decoding, upload or UTF-8 text validation.

## Snapshots and execution

- Tested Gossip tree `eabbffa481c23c04ed1c29380c2b2845ae8d1171`, recorded
  unchanged in commit `fb80796656c2628f73f086e8e53c53356e371b15`; branch
  `codex/gohp-heic-helper-bundle` on Draft #21.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`.
  Three original files and all16 fragment ranges/hashes are in the
  [machine record](heic-helper-bundle-2026-10-25.json). The new HEIC fragment
  is functions.php9465–9474, hash
  `de4452e658e48a6cac0d02b501b425bb8ac60872fdbbe9ab200a107f98223585`.
- Unchanged GoHP `49a302e44464623dc4912bafbcf44192759d7878`, tree
  `dc2046b28420a9f7a9da7447052a5e70d6543279`,
  [Draft #44](https://github.com/katsuakikureyama/GoHP/pull/44).
  Separate #43 control-flow integration is not adopted.
- HostGo1.22.1 darwinarm64, builderGo1.25.14 linuxarm64, PHP8.3.36,
  LC_CTYPE=C. Existing immutable-image admission and same-process guard,
  non-root/read-only/network-disabled containers and resource limits retained.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressHEICHelperBundle$'
go vet ./tests/gohp
```

Initial runtime PASS32.63s/package32.878s; vetPASS. No runtime failure/retry.
This is local isolated execution, not a new Gossip CI result. Post-PASS changes
are documentation/evidence only; no code/config changes or repeated execution.

Snapshot timing exception: the pre-commit `git diff --cached --check` exited2
for a trailing blank line at PHP-observer EOF, preventing the initially chained
commit. The test still launched against staged working-tree bytes with launch
HEAD `d167cc33fd9d2fbc82acae2453a782f85f0744f9`. During execution, no unstaged
differences were present; `git write-tree` recorded the tested tree above and
those exact bytes were committed without editing. The blank line is retained.
This is not a claim that the test commit existed before launch. Actual
observer/manifest/source hashes bind the compared bytes independently.

## Corpus and boundaries

All89 prior charset-bundle cases are retained with unchanged arguments,
expectations and post-call expectations; all24 direct HEIC cases are added.
Six new compositions cover slice→HEIC, slice→shallow-unescape→HEIC and
slash/unescape→HEIC→Boolean. Mixed/negative keys, selector order and duplicates,
null/missing removal, false/zero values, escaped slashes and NUL are observed.

New HEIC inputs are scalar-only, including every root value of the new slice
operations even if unselected. Older operations retain their bounded nested
sequential-list transport. Actual output and input-after observations never
come from expected fields. See the [contract](../heic-helper-bundle-unit.md).

## Reused compiler evidence, not another CI run

The exact unchanged GoHP pin already passed
[CI40159749575](https://github.com/katsuakikureyama/GoHP/actions/runs/40159749575):
all336 checks across52 features plus vet, CLI smoke and artifact publication.
ActualPHP8.3.36/Go1.23.12 linuxamd64, clean exact commit/tree. Command:
`go test -count=1 -json ./cmd/... ./internal/... ./runtime ./tests`.
No compiler changes, tests or CI retries in this change. This reused compiler
result does not substitute for the new119-call Gossip integration above.

Artifact `compatibility-php8.3-40159749575-1`, ID12586340125, digest
`sha256:4391dd90f542b6a2a9a5bbe5ff62f4cd63284e07406ce7f6662eaad0f5aa95e2`,
expiry2026-11-24; previously retrieved to `/private/tmp/gohp-in-array-ci.Gn0dnf/`.
Full reused metadata and local commands remain in the machine record and
[earlier HEIC evidence](heic-mime-2026-10-25.md).
Compiler native schema observes stdout/normal exit, not general stderr,
PHP exceptions or external side effects.

## Unchanged and unverified

Ten standalone units and all earlier bundle manifests/corpora/evidence remain
unchanged, not rerun in their old configurations. Their selected cases are
newly measured together under this new configuration; this is not every
historical standalone input. WordPress originals, SDK/hooks, oracle policy,
Gossip-prototype and master README-only state remain unchanged.

General PHP/loose membership, array/object/float HEIC inputs, nested associative
transport, option-dependent charset wrappers, recursive unescaping, whole-file/
core/HTTP/DB/theme/plugin behavior remain unverified. No merge/deploy/release/
license decision. Local developer capability remains WebMCP-excluded; auth
routes unchanged. Devboard unidentified/unsynchronized, not Done.
