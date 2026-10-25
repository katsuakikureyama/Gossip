# Fifteen-original charset/array integration — 2026-10-25

PASS: fifteen byte-unchanged WordPress functions and nineteen owned drivers
were translated/built together once. All89 actual calls (55 direct/34 composed)
match PHP8.3.36 and authored expectations for return types/values, ordered
root keys where applicable, and caller arguments after invocation. Empty
stderr and exit0 are required. This is a new configuration, not relabelled
evidence from either compiler parent or earlier Gossip units.

## Exact snapshots

- Gossip test commit `b0835877b9e2172670582a71697f066833837257`, tree
  `27dd48a8bfcbb863b76cb483c04b93f21094c4af`, committed before execution;
  branch `codex/gohp-charset-helper-bundle`, based on #19.
- WordPress `54106f6c545fc8b881f6b39b7963002bece6c28c`. Whole-file hashes and
  selected ranges from functions.php, formatting.php and compat.php remain
  unchanged. Each original excerpt is copied byte-for-byte; slice retains
  its `: array` declaration, charset retains its negation/OR/string comparisons.
- GoHP `6305b9b486ee233e9aa4378ab8528f5b3d345f90`, tree
  `48c92101a7000794eda13cdfbbd0240b7327afb9`,
  [Draft #41](https://github.com/katsuakikureyama/GoHP/pull/41) on #40.
  Merge parents are #40 `d9f6c7b4f288ef816c4ee3cb6a3458d2226e00d5` and
  #36 `db3051f8a179078e2227c5d97f50a2f1611d1451`. Separate #39 if/else
  is not included. Reserved-name/fixture/spec conflicts retain both contracts.

The [machine record](charset-helper-bundle-2026-10-25.json) preserves every
source/excerpt/manifest/corpus/driver/observer/compiler/archive/generated/
report/binary hash, bundle offsets and all actual typed observations.

## Gossip execution

From nested `gossip/`, hostGo1.22.1 darwinarm64, builderGo1.25.14 linuxarm64,
actualPHP8.3.36/LC_CTYPE=C:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressCharsetHelperBundle$'
go vet ./tests/gohp
```

Initial PASS19.04s/package19.279s; vetPASS. No execution failure/retry.
Before the committed run, parent review corrected two new expected result
orders to selector order, included null/falsy selectors, and corrected direct
charset observer dispatch to accept non-string inputs. None was a post-PASS
change. Subsequent edits are documentation/evidence only; no rerun.
Existing immutable images, same-process native admission, non-root,
read-only/network-disabled/resource-limited execution are retained. This is a
local isolated comparison, not a Gossip CI run.

The corpus retains the prior55 associative-bundle cases unchanged as parsed
data, all28 standalone charset inputs/expectations, and adds six compositions.
They combine slicing with charset mapping, slicing with shallow stripping and
mapping, or slash/charset/Boolean conversion. Keys and null/falsy distinctions
are observed; nested arrays are not mistaken for charset-name strings.
See the [contract](../charset-helper-bundle-unit.md) for exact transport bounds.

## Merged compiler validation

At the exact compiler head above, cachedGo1.23.0 darwinarm64, prefix
`GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0`:

```sh
go test -count=1 ./runtime ./internal/frontend ./internal/codegen ./internal/compat
go test -count=1 ./tests -run '^(TestArrayReturnRuntimeGuards|TestArrayIssetRuntimeGuards|TestStrCaseCmpRuntimeTypeGuards)$'
go vet ./...
```

Runtime/frontend/codegen/compat PASS0.257/0.457/0.562/0.733s;
generated-binary guards PASS5.801s; vetPASS.
[Fixed PHP CI40159252743](https://github.com/katsuakikureyama/GoHP/actions/runs/40159252743)
(push, attempt1) passed all313 checks across51 features, vet/CLI smoke and
artifact publication. ActualPHP8.3.36/Go1.23.12 linuxamd64; exact clean head/tree.
Command `go test -count=1 -json ./cmd/... ./internal/... ./runtime ./tests`.
Automatic PR run40159255782 also succeeded; no manual retry or post-PASS
compiler code/configuration changes. The new shared native fixture is
`charset_array_integration.php`; existing checks from both parents remain.
Native compiler schema compares stdout/normal exits, not general stderr,
PHP warnings/exceptions or external effects.

Artifact `compatibility-php8.3-40159252743-1`, ID12586085024, digest
`sha256:576222737242c9cdc768f22181f0415e514251b949b4c42933326b47e56bcbf8`,
expires2026-11-24; retrieved to `/private/tmp/gohp-charset-array-ci.JkfCnY/`.

## Separation and remaining scope

Earlier nine independent units and ten/twelve/thirteen/fourteen-original
bundles retain their own pins/corpora/evidence and were not repinned or rerun.
These89 observations are not every earlier standalone input. Original files,
SDK/hooks, PHP admission, Gossip-prototype and master README-only state are
unchanged. Charset recognition is not UTF-8 text validation/conversion.
Public option-dependent wrappers, general object/reference/callback behavior,
nested associative transport, map_deep/wp_unslash, whole-file/core startup,
HTTP/DB/theme/plugin compatibility remain unsupported or unverified.
No implementation merge to master/main, deploy, release or license decision.
Local tooling remains WebMCP-excluded; auth routes unchanged. Devboard card
unidentified/unsynchronized; no broader Done claim.
