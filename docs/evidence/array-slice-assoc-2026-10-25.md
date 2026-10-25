# Original associative array-slice evidence — 2026-10-25

Result: **PASS, case-limited**. The exact original `wp_array_slice_assoc()`
including `: array` was translated, built and executed on fourteen inputs.
Actual return type/ordered entries and caller input/selectors after invocation
matched PHP8.3.36 and explicit expectations. Empty stderr and exit0 required.
This is a local isolated Docker result, not Gossip CI or whole-WordPress support.

## Revisions

- Gossip branch `codex/gohp-array-slice-assoc`, based on #15. Test commit
  `1b191a17b574f174e04249b6420e893ec57750c5`, tree
  `c0b85668909beca251efceb4d56351af08c3538e`.
- [GoHP Draft #36](https://github.com/katsuakikureyama/GoHP/pull/36) on #34:
  `db3051f8a179078e2227c5d97f50a2f1611d1451`, tree
  `3155927445663c82c783c5d062b009b7d1a13648`. Parallel #33/#35 are not included.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`, functions.php
  lines5173–5183. Full-file/excerpt/bundle, compiler/archive, generated output,
  report/binary, case and observer hashes are retained in
  [machine evidence](array-slice-assoc-2026-10-25.json).
- [Contract](../array-slice-assoc-unit.md): eighth standalone unit, not a
  fourteen-function bundle. Earlier pins/corpora/evidence are unchanged.

## Commands and observed results

Nested `gossip/`, hostGo1.22.1 darwin/arm64, builderGo1.25.14 linux/arm64,
PHP8.3.36 from the existing approved immutable image. Non-root/read-only/
no-network/resource restrictions and same-process oracle guard unchanged.

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArraySliceAssocHelper$'
go vet ./tests/gohp
```

PASS8.42s/package8.619s, all14 actual observations match; vetPASS. Before the
parent's final review/commit, a worker inadvertently invoked
`go test ./tests/gohp -run '^TestOriginalWordPressArraySliceAssocHelper$'`.
It failed at setup while opening the host Go build cache (`operation not
permitted`); no test body, Docker or PHP comparison ran. The committed test's
first actual isolated comparison passed. No rerun or code/config change after
PASS; subsequent changes are documentation only.

Compiler local cachedGo1.23.0 darwin/arm64, on the final pin:

```sh
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./runtime ./internal/frontend ./internal/codegen ./internal/compat
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./tests -run '^TestArrayReturnRuntimeGuards$'
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go vet ./...
```

runtimePASS0.602s/codegenPASS0.788s/compatPASS0.987s; frontend has no test files.
Three generated runtime-negative cases PASS1.622s; vetPASS, initial success.

[Fixed PHP CI40156201018](https://github.com/katsuakikureyama/GoHP/actions/runs/40156201018)
(push/attempt1) passed all238 checks across46 feature entries, actualPHP8.3.36 /
Go1.23.12 linux/amd64, clean checkout commit/tree matching the pin.
Command `go test -count=1 -json ./cmd/... ./internal/... ./runtime ./tests`;
whole job including vet/CLI smoke/artifact retention passed. Artifact
`compatibility-php8.3-40156201018-1`, ID12585095049, digest
`sha256:81ba128c07a3da3116d725b69cf3d17971934d0f8de25ee769b7f9031fc3e5e3`,
expires2026-11-24, downloaded to `/private/tmp/gohp-array-return-ci.CoMkuY/`.
Compiler native schema observes stdout/normal exits, not general stderr,
exceptions or external side effects. This is separate from the narrower Gossip
unit's explicit empty-stderr assertion. No manual CI retry or post-PASS retest.

## Remaining scope

No original PHP modification or implementation merge/release/deployment.
Master's approved README-only publication remains intact. Next: independently
integrate this fourteenth original with the earlier bundle on a new pinned
configuration, without reusing old PASS as proof of the combined revision.
Core/HTTP/DB/theme/plugin behavior and license decisions remain unverified.
Local developer capability is WebMCP-excluded; no authentication changes.
Devboard card remains unidentified/unsynchronized, not Done.
