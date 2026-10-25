# Original UTF-8 charset predicate — 2026-10-25

PASS: all 28 calls to unchanged `_is_utf8_charset` match native PHP 8.3.36
and authored expectations in actual Boolean return and post-call caller input.
Six charset-name variants return true; fourteen near-match/other strings and
eight non-string inputs return false. Empty stderr and exit0 are required.
This classifies charset names, not the validity of arbitrary UTF-8 text.

## Exact snapshots and execution

- Gossip test commit `fccd38e47c4f23ed0bb8b35a0cc974c9486f4d2f`, tree
  `64499078c9dd73381ba28bce7c3cb889392981cc`; branch
  `codex/gohp-utf8-charset` on #18. Committed before the first execution.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`, compat.php
  lines102–111. Whole-file SHA256
  `368f9b3ac117582698fd6df40b99f802479b26ddfc4a5f012f7ebd486df92487`;
  exact excerpt SHA256
  `241a7691e4461f13f8656fb12b8275fb4eaa71e34b983ef63301965286df5739`.
- GoHP `d9f6c7b4f288ef816c4ee3cb6a3458d2226e00d5`, tree
  `8acaff0e9073db725edd95652a1fc9dd9442c4f0`,
  [Draft #40](https://github.com/katsuakikureyama/GoHP/pull/40) on #38.
  Adds guarded string-only ASCII strcasecmp to the logical/comparison stack.
  Separate #36 isset/array-return changes are not integrated.
- Host Go1.22.1 darwin/arm64; builder Go1.25.14 linux/arm64;
  PHP8.3.36, LC_CTYPE=C. Existing immutable-image admission and same-process
  version guard, non-root/read-only/network-disabled/resource limits retained.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressUTF8CharsetHelper$'
go vet ./tests/gohp
```

Initial PASS37.69s (package37.898s); vet PASS. No failure/retry. Post-PASS changes
are documentation/evidence only; no code/configuration changes or repeated
runtime validation. This is local isolated execution, not Gossip CI.
The [machine record](utf8-charset-2026-10-25.json) retains all actual observations,
manifest/source/excerpt/corpus/observer/compiler/archive/generated/report/binary
hashes, exact command and compiler CI metadata.

## Compiler validation

At the same GoHP head, cached Go1.23.0 darwin/arm64 with
`GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0`:

```sh
go test -count=1 ./runtime ./internal/codegen ./internal/compat
go test -count=1 ./tests -run '^(TestUnsupported/strcasecmp_|TestStrCaseCmpRuntimeTypeGuards)$'
go vet ./...
```

Runtime/codegen/compat PASS0.251/0.390/0.592s; generated runtime guards
PASS1.866s; vet PASS. The second selector establishes generated guards, not
the slash-named static subtests; those are covered by the full CI below.

[Fixed PHP CI40158737961](https://github.com/katsuakikureyama/GoHP/actions/runs/40158737961)
(push, attempt1) passed all 296 catalog checks across49 features, vet,
CLI smoke and artifact publication. Exact clean head/tree match the pin;
actualPHP8.3.36/Go1.23.12 linuxamd64. Command:
`go test -count=1 -json ./cmd/... ./internal/... ./runtime ./tests`.
The automatically triggered PR run40158740634 also succeeded; no manual retry.
Artifact `compatibility-php8.3-40158737961-1`, ID12584808556, digest
`sha256:914d01bcc33d0993c533e9fd9b897763e8ca7463683020e955b7b83224dee2a0`,
expires2026-11-24. Retrieved to `/private/tmp/gohp-strcasecmp-ci.evz7ac/`.
The new strcasecmp fixture observes numeric results, ASCII/non-ASCII, prefix,
NUL, operand markers and guarded predicate composition. Compiler native schema
observes stdout/normal exit, not general stderr, exceptions or external effects.

## Separation and unresolved scope

This is the ninth independently pinned unit, not a fifteen-original bundle.
All older unit/bundle manifests, corpora and evidence stay unchanged and were
not rerun. In particular, the existing 55-case fourteen-original bundle still
uses its earlier compiler; its array-return stack is not silently replaced.
The [contract](../utf8-charset-unit.md) records modest-integer/sequential-list
transport bounds and rejected domains.

Public `is_utf8_charset` requires defaults/coalescing/get_option; no stubs or
rewrites are used to claim it. `_canonical_charset`, encoding conversion,
whole-file/core/HTTP/DB/theme/plugin execution remain unverified. Original
WordPress, SDK/hooks, PHP admission and master README-only publication are
unchanged. No implementation merge, deploy, release or license decision.
Local capability remains WebMCP-excluded; auth routes unchanged. Devboard
unidentified/unsynchronized, not Done.
