# Original slash-helper observations — 2026-10-25 JST

## Targets and command

- Tested Gossip implementation: `a97f14f42e6f59049062d59695c74da9d56ee0d2`.
- Original WordPress: `54106f6c545fc8b881f6b39b7963002bece6c28c`.
- GoHP compiler/runtime: `ceee0fc71e1e7947f1c78f1dd18875236b9a180b`,
  [Draft PR #10](https://github.com/katsuakikureyama/GoHP/pull/10), based on #8.
- Host test driver Go 1.22.1 darwin/arm64; isolated builder Go 1.25.14 linux/arm64;
  oracle PHP 8.3.13. Source/images/artifacts are pinned by SHA in the
  [machine-readable observation](slash-helpers-2026-10-25.json).

From `/Users/kureyamayuri/Desktop/Gossip-wordpress/gossip`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressSlashHelpers$'
```

PASS on the first run: 18.82 seconds (package 19.018 seconds). No code/config
changed or test rerun after PASS; later changes are documentation/evidence only.
The existing constant-return and native SDK/hook units were unchanged and not
rerun. Their older pins/results are separate evidence, not implicitly upgraded.

## Observations

The unchanged `trailingslashit` and `untrailingslashit` excerpts from
`wp-includes/formatting.php` were translated in one bundle. The original
dependency chain reaches GoHP's new restricted `rtrim` implementation.
Generated source/report hashes were verified before compilation against the
same pinned runtime; no WordPress implementation was replaced by a Go wrapper.

All 13 inputs produced matching PHP/Go string types and values for both
functions: **26 actual typed returns**, empty stderr and exit 0. Results also
matched the independently authored expectations in `slash-cases.json`.
Coverage includes empty/slash-only strings, repeated mixed slash suffixes,
leading/interior separators, Windows-style text, whitespace, Japanese,
URL-looking text, embedded NUL and newline. Full values are in the JSON record.
No input causes a filesystem or network operation.

Observers call actual original/generated functions. JSON/type inspection and
case-file reads are instrumentation, not evidence of translated PHP JSON or
file APIs. The domain is actual string inputs only. See the
[unit contract](../slash-helper-unit.md) for isolation and exclusions.

## Compiler validation (separate evidence)

At the exact GoHP commit above, with cached Go 1.23.0, host darwin/arm64:

```sh
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./runtime -run '^TestRTrim'
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./tests -run '^TestUnsupported/rtrim_'
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./internal/compat
```

All PASS (0.220, 0.194, 0.240 seconds respectively). `git diff --check` PASS.
[PR CI run 40112678059](https://github.com/katsuakikureyama/GoHP/actions/runs/40112678059)
passed both Go 1.23.12 / PHP 7.4.33 and PHP 8.3.30 jobs, including vet, native
differential execution, summary publication and artifact storage. Both summaries
record `diff.rtrim`, nine exact rejection checks and two runtime checks PASS.
The `rtrim` status is case-limited for each tested target, not arbitrary-PHP support.

Actual PR merge checkout `0a0a883cbaf43b2cb27c7222293d5a32f2b3b287` was clean;
tree `d00e16b86e5ddc700c5837bf56ddfd6ae1b2b881` matches the compiler head.
Artifacts (30-day retention): `compatibility-php74-40112678059-1` ID12566864216,
`compatibility-test-40112678059-1` ID12567480496. No extra rerun after PASS.
This compiler matrix is not a WordPress PHP 7.4 compatibility result.

## Limits and state

The compiler supports exactly two-argument `rtrim` with string input and a literal
byte mask without `..`. Default masks, range masks, dynamic masks and weak scalar
coercion remain outside the contract. Runtime rejection of dynamic non-string
values is not PHP warning/exception equivalence. The byte-split native fixture
tests byte semantics, not Unicode normalization.

No full `formatting.php`/WordPress translation, path canonicalization, HTTP
behavior, plugin/theme support, database mapping or production operation is
verified. Both projects remain on separate Draft PRs; no default branch merge,
deployment, release or license finalization. WordPress originals are unchanged.
Local developer-test/compiler capabilities remain WebMCP-excluded. No linked
devboard card is known or marked Done by this result.
