# Twelve-original integrated helper bundle — 2026-10-25 JST

## Outcome

**PASS:** twelve unchanged WordPress functions and seven owned PHP composition
drivers were translated together into one generated Go executable. All24 typed
observations matched PHP8.3.36 and authored expectations, with empty stderr and
exit0. This is new combined-compiler evidence, not an aggregation of previous
independent PASS results or a WordPress core startup test.

- Gossip test commit `74e234dcd80674c03b7f718040e8804866c8b04f`, tree
  `ae04dc19a38c88556f56fc9e4de707a4849f2159`.
- Combined GoHP compiler/runtime `e5ba8c1db81d235cc177cc83699868d812abff7e`,
  tree `0a82fee62457c2bc656bdd3e4d0a71c9d6d81e75`.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c` unchanged.
  functions.php SHA `456ab4cc7a6e6f05b061f2b1905f44233f8906774c461f90c5b80b9841303ccc`;
  formatting.php SHA `d44aa1329df2fc38bf2670ad6de570e6b9eb65db9ed5507f26c683e569b2919e`.
- Actual PHP8.3.36 with explicit LC_CTYPE=C guard; approved immutable images,
  same-process native-version guard and isolation controls unchanged.
  Host Go1.22.1 darwin/arm64; builder Go1.25.14 linux/arm64.

The [machine record](integrated-helper-bundle-2026-10-25.json) preserves all
source selections/ranges/excerpt hashes/bundle offsets, driver/corpus/observer
hashes, compiler/archive/report/generated/executable hashes, native versions,
all24 observations and compiler CI metadata. Only committed compiler files were
exported, and report source/output hashes were verified before compilation.

## Corpus and execution

Twelve direct probes exercise each selected original once: six constant-return
helpers, Boolean validation, trailing/untrailing slashes, zeroise, newline
preservation and string-only stripslashes. Twelve further probes compose:

- Four capture-list values → newline preservation → string-only unescaping,
  including quotes, NUL conversion, Japanese text and ignored extra captures.
- Two non-string array passthrough → array_keys probes.
- Two array passthrough → list-shape plus array_values probes, observing the
  actual escaped values inside nested lists rather than shape alone.
- Four scalar compositions: zero→padding, false→Boolean, empty string→slash,
  multibyte path→slash→padding.

The exact [contract](../integrated-helper-bundle-unit.md) excludes associative
JSON maps, floats, objects and references. Numbers are transported using
UseNumber/Int64; the corpus contains only modest integers, not an arbitrary
large-integer JSON equivalence claim. Array results require sequential int64
keys before recursive normalization. Observers dispatch and normalize types,
without copying original function bodies or computing expected results.

From nested `gossip/`, at the test commit above:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressIntegratedHelperBundle$'
go vet ./tests/gohp
```

Initial PASS **7.60s**, package **7.834s**; vetPASS. No failures, retries or
implementation changes after PASS. Subsequent changes are documentation only.
This is local isolated execution, not a Gossip Actions run.

## Combined compiler

[GoHP Draft #24](https://github.com/katsuakikureyama/GoHP/pull/24) is stacked
on #22 and combines parents `5be2e59c3bb05d98ddee2bcab700888a637b3965`
and `0e49cc995913dfede0332566dd4dc11c5df16ff8` (#23). Conflict resolution
retains all builtin reservations, both fixture registries and the shared
semantic argument-validation pass. No runtime domains, parser semantics,
native admission or CI isolation policy are broadened. Other checkout unchanged.

At the exact combined commit above, cached Go1.23.0 darwin/arm64, once:

```sh
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./runtime ./internal/codegen ./internal/compat
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go vet ./...
```

PASS runtime0.364s / codegen0.549s / compat0.748s; vetPASS.
Catalog40 entries/166 unique checks. The new native integration fixture
combines conditional unescaping with keyed arrays, null presence, keys/values/
list checks, formatting/replacement, visible lazy evaluation and original-value
preservation on the same compiler/runtime.

[Fixed PR CI40149530122](https://github.com/katsuakikureyama/GoHP/actions/runs/40149530122)
PASSed all166 checks on Go1.23.12 linux/amd64 / actualPHP8.3.36. Clean PR
checkout `5cbb87344104c906bf3595b4a8c54da09ff9612f` has the same tree as
the combined head. Native test exit0, vet, summary and artifact publication
succeeded. Artifact `compatibility-php8.3-40149530122-1`, ID12581503118,
retained30days and downloaded to `/private/tmp/gohp-integrated-ci.K6QcGR/`.
Initial success, no manual rerun. Native CI observes stdout/normal exit, not
general stderr/exception/side-effect equivalence. Gossip separately enforces
empty stderr for its24 cases.

## Preserved boundaries

Earlier six standalone configurations and the ten-function bundle retain their
pins, corpora and evidence and were not rerun. This new24-case combined result
does not relabel every earlier input as verified on the combined compiler.
No WordPress original, SDK/hook, default branch or other session's checkout was
changed. Recursive callbacks/unescaping, whole files, core startup, HTTP/DB,
themes/plugins and general PHP remain unsupported or unverified.

Draft only: no default-branch merge, deployment, release or license decision.
Local developer capability remains WebMCP-excluded, with no authentication-route
change. No devboard card is identified, synchronized or marked Done.
