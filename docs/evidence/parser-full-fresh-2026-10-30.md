# Complete original parser fresh-PHP comparison — 2026-10-30 (JST)

## Result and adoption

The existing `fresh` mode executed **once**, at clean Gossip
`1f7a3b262a709203a799e07b12f636c58d44f7b5`, tree
`b3c6ceb96c4548ce79c7a660bb3fb737a0e57c45`, with unchanged GoHP
`b2d1482268a7710c515c07bf919af3d40b09d17c`. It passed in2.61s
(package2.843s), with12 cases/24 actual original parse calls per side and no
mismatch. A new pinned PHP execution and the complete generated original agree
on all canonical typed cases, not just a saved expectation file.

The later measurement owner found earlier boundary-review scope6519170236,
stopped overlapping review work and handed over the completed fresh result.
The canonical review adopted it **without another test/build/parse execution**.
[Coordination](https://github.com/katsuakikureyama/GoHP/issues/1#issuecomment-6519250519)
and [adoption](https://github.com/katsuakikureyama/GoHP/issues/1#issuecomment-6519323643)
preserve that distinction; there is only one publication changeset.

| Evidence | Exact identity |
|---|---|
| Raw measurement, preserved unchanged | SHA256 `3d83a0abc23d0441aade911fdcd900635761ecad5184ce012a7f7e925b9b31cf` |
| Handoff provenance | SHA256 `584a37bf4ac09b30fd0b7177e3f1216f6c293f5e342d6cc8d28350ca8249e7b5` |
| [Published machine record](parser-full-fresh-2026-10-30.json), including execution/publication provenance | SHA256 `6bf6a15e6930664c1181ba655be58930997955ae65086009f04e7ea56ddd1eab` |
| Native and Go canonical typed `cases` | SHA256 `3780c6a5ebd4eefda51bc548b2e95f585687c39f274eb6d3c8a622792607ca28` |
| New PHP process stdout | 60065bytes, SHA256 `c02b289131717834e4b927721738804ac641b3cb24affc21815f657cad1faf57`, exit0/stderr0 |
| Generated original stdout | 59260bytes, SHA256 `e0e5a47f2e36b7cbbcf293e16b3ced0417cd3eb72b8a63eba20ea85ca51deebb`, exit0/stderr0 |

The **new PHP stdout matches the old PHP stdout** byte-for-byte. PHP versus Go
whole stdout does **not** match: their metadata envelopes differ. The actual
PHP-versus-Go comparison is the typed `cases` projection, including ordered
keys, raw byte strings, parser fields, JSON/PCRE errors, caller preservation and
fresh/reused-parser state. Retained and fresh runs are separate configurations;
they do not become24 distinct cases or broaden the12-case input denominator.

## Execution configuration

The already-running Docker25.0.3 and exact cached oracle image were available
to the measurement owner; neither owner started a daemon, pulled an image or
installed a dependency. Image
`php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd`,
local image ID
`sha256:f39899ba9e42005403f535ec925a6e46a1c70707085b689e41a48219798521d7`,
Linux/arm64, PHP8.3.36/PCRE10.42, C LC_CTYPE and the original fixed settings.

The generated candidate uses the existing measured Darwin/arm64 Go1.23.0,
clang17/SDK15.5, CGO1/`gohp_pcre2`, static PCRE10.47_1 configuration from
[the retained execution](parser-full-generated-2026-10-30.md). The parent test
process uses CGO0. No original source, corpus, observer, harness, pin or runtime
changed. Existing checked generatedGo/report35524a6 were reused after the
unchanged implementation check; old generation/CGO0 build/retained-only suite/
CI results were not rerun. Temporary candidate binary was test-cleaned.
The duplicate-static-library linker warning is a build warning, not runtime
stderr; both runtime processes exited0 with empty stderr.

Exact effective command, personal paths replaced by verified named inputs:

```sh
GOSSIP_GOHP_FULL_PARSER_MODE=fresh \
GOSSIP_GOHP_FULL_PARSER_SOURCE="$GOHP_ROOT" \
GOSSIP_PCRE2_PREFIX="$VERIFIED_PCRE2_PREFIX" \
GOSSIP_GOHP_FULL_PARSER_GENERATED_DIR="$PRIOR_GENERATION_DIR" \
GOSSIP_GOHP_FULL_PARSER_EVIDENCE_OUT="$FRESH/attempt1.json" \
GOSSIP_PHP_RUNNER_IMAGE=php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd \
CC="$VERIFIED_CLANG" SDKROOT="$VERIFIED_MACOSX_SDK" \
GOCACHE="$SHARED_EXISTING_CACHE" GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off \
GOWORK=off CGO_ENABLED=0 \
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressFullParserGenerated$'
```

The existing harness hash
`88f9cd6baa1e113356b0531c6df98f1077db9148e0e3ec0ab44b510be499a825`
checked the exact source/corpus/native-observer/profile and compared fresh
cases. Run count1; no failure or retry in this fresh configuration. The published
record adds provenance to the unchanged observations; its different file hash
does not imply a second execution.

## Acceptance remains bounded

This resolves the fixed-v1 fresh comparison prerequisite. The
[boundary review](../parser-full-acceptance-boundaries.md) still identifies
unmeasured JSON/type/resource and carry-in-state cases, observer exclusions,
engine-profile differences and outside-profile APIs. All measured PCRE errors
are0; only JSON syntax4 is exercised. Do not infer global error resets, resource
equivalence, arbitrary attributes/documents or all parser branches.

G83-PARSER/application acceptance is not established. Bootstrap/DB/HTTP/site
gates remain not-run. Original sources/old manifests/corpus/observers/pins/native
and retained machine records, auth/routes/deployment and the local test/WebMCP
exclusion are unchanged. No GoHP change/PR or new CI/cloud resources.
