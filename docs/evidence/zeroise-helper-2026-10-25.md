# Original zeroise helper observation — 2026-10-25 JST

## Pinned targets

- Gossip final test commit: `ba30252d7509a23f2f8b97de8f0b4ac748947c2b`, tree
  `9a5c7087e808f670e234a52678b7c4cb2ae776cc`.
- GoHP compiler/runtime: `e3db685dfb2a3602f94bbfa9df105db3db96cc89`, tree
  `a62f5cf9fe8f36e8c95c1502ee46301360889876`.
- WordPress baseline: `54106f6c545fc8b881f6b39b7963002bece6c28c`, unchanged
  `wp-includes/formatting.php` lines 2848–2850 (`zeroise`).
- Actual native PHP: **8.3.36**, admitted immutable image and same-process guard
  unchanged. Host Go 1.22.1 darwin/arm64; builder Go 1.25.14 linux/arm64.

[GoHP Draft PR #18](https://github.com/katsuakikureyama/GoHP/pull/18) is stacked
on #16. It adds only bounded two-argument `sprintf('%0Ns', value)` with canonical
width 0–4096 and string/int64 values. The compiler and runtime share format
validation; known invalid literals fail compilation and dynamic violations fail
at runtime. Other formats/coercions remain unsupported. The separate #17 stack
and normal GoHP checkout were not modified or imported. Catalog: 32 entries and
78 unique checks, active PHP target 8.3 only.

## Gossip execution

From the nested `gossip/` module:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressZeroiseHelper$'
go vet ./tests/gohp
```

Final-pin run **PASS**, test 10.67 seconds, package 10.886 seconds. All 18 cases
matched in actual string return type/value between PHP, generated Go and authored
expectations, requiring empty stderr and exit 0. The cases include zero/exact/
smaller widths, negative and zero integers, empty/leading-zero/whitespace strings,
Japanese text, an emoji, NUL and signed int64 extrema. Observed widths are 0–22;
the runtime unit test, not this original-source fixture, checks the 4096 ceiling.
Vet PASSed on the unchanged observer/test code before the pin-only update.
No implementation changes or reruns followed the final-pin PASS.

The initial Gossip test commit `75d897e0f5a50a5f19dfc59c1f197a9485929d82`
also PASSed against compiler `151d2fa8657e4fe196309e0bdc97c0d09bd61b29`
(10.98s, package 11.192s). That [initial observation](zeroise-helper-initial-2026-10-25.json)
is retained separately. The compiler CI then found a fixture-only unsupported
negative integer literal; the fixture/doc correction changed the pinned archive
but no compiler/runtime implementation. Only the affected zeroise unit was run
once against the corrected pin. Generated/compiler/program/report hashes and all
18 observations remained identical; the archive and commit identities changed.

The [machine-readable record](zeroise-helper-2026-10-25.json) contains every
observation and the source, excerpt, case, compiler/archive, observer, report,
generated and binary hashes. Generation report hashes were verified before
building. The isolated export contains only the pinned committed compiler source.
No new source mutation or weakened container/admission setting was needed.

This is a fourth independently built unit with a newer compiler pin. The earlier
nine original functions keep their three unchanged manifests and earlier
[shared-compiler result](unified-compiler-2026-10-25.md). They were not rerun or
claimed to pass on the new compiler. Ten selected functions across four units
and two compiler pins do not establish a combined ten-function program or core
startup. Unchanged hooks/admission retain their prior exact-configuration PASS.

## Compiler verification

At initial implementation commit `151d2fa8657e4fe196309e0bdc97c0d09bd61b29`,
cached Go 1.23.0 darwin/arm64:

```sh
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./runtime ./internal/codegen ./internal/compat
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go vet ./...
```

PASS: runtime 0.439s, codegen 0.253s, compatibility catalog 0.645s; vet PASS.
The fixture/doc-only correction preserves these package results; they were not
repeated. The local command deliberately did not invoke host PHP. The native differential
result belongs to the fixed PHP 8.3.36 CI run, recorded separately below.

Initial [PR CI 40118327850](https://github.com/katsuakikureyama/GoHP/actions/runs/40118327850)
FAILed in `TestDifferential/sprintf` with `PHP002 line 7: unsupported expression
(at "42")`: the test fixture, not the original WordPress function, used `-42`
as a PHP literal. The existing parser does not support unary minus. The fixture
now uses the string `'-42'`; direct negative int64 runtime tests and the Gossip
typed-int cases remain unchanged. No parser scope was expanded. Initial failure
artifact `compatibility-php8.3-40118327850-1`, ID12569678130, is retained by CI
for 30 days.

Corrected [PR CI 40118450143](https://github.com/katsuakikureyama/GoHP/actions/runs/40118450143)
**PASSed** on fixed Go 1.23.12 linux/amd64 / PHP 8.3.36. Its clean merge checkout
`4c654dd8ef57c49cb135f304cf4b960356952b99` has tree
`a62f5cf9fe8f36e8c95c1502ee46301360889876`, identical to the pinned compiler
head. All 78 catalog checks PASS, including `diff.sprintf` and its four
compile/runtime checks. Native differential evidence compares stdout/exit only;
it does not claim stderr, exception or side-effect equivalence. Gossip's separate
typed test above also requires empty stderr. Artifact
`compatibility-php8.3-40118450143-1`, ID12569703294, is retained for 30 days.
The stored run records test exit 0 and actual PHP 8.3.36. No rerun after PASS.

## Limits and next steps

No default-branch merge, release, deployment, license finalization or upstream
WordPress change. General formatting, coercion, warnings/exceptions, full PHP,
HTTP/DB/plugins/themes and WordPress startup remain unverified. The local
developer capability remains explicitly WebMCP-excluded under the
[unit contract](../zeroise-helper-unit.md). No devboard card is linked, synced or
marked Done. Choose further source units by dependency and side-effect contract;
do not assume a compiler addition generalizes the existing observations.
