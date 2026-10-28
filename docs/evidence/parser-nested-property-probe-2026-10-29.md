# Nested-property admission — 2026-10-29

The unchanged complete parser now passes chained-property updates at207–208
and stops at the property input to `preg_match` at252. This is source-admission
progress for [G83-PARSER](../php83-application-acceptance.md); complete semantic
validation, generated binary execution and application acceptance remain pending.

- WordPress `54106f6c545fc8b881f6b39b7963002bece6c28c`.
- Whole `wp-includes/class-wp-block-parser.php` SHA-256
  `8563621090de5e307431b1d29b78a9f75fd9ee238a0137b563a7db7b2e1f8d3b`.
- GoHP [PR122](https://github.com/katsuakikureyama/GoHP/pull/122), clean commit
  `2421dafb0eec3baea0a778b921082949dc0bda9d`, tree
  `86ba970ebce04efe0d711410e32e7e422f3e4ca2`, based on mainf6b6bbc.
- Go1.23.0 Darwin ARM64, CGO0, cached/offline modules.

Personal local paths are replaced with `GOHP_ROOT`, `GOSSIP_ROOT` and `CACHE`:

```sh
cd GOHP_ROOT
CGO_ENABLED=0 GOCACHE=CACHE GOPROXY=off GOTOOLCHAIN=go1.23.0 \
go run ./cmd/gohp -in GOSSIP_ROOT/wp-includes/class-wp-block-parser.php \
  -php-version 8.3 -experimental-pcre2
```

One invocation, exit **1**:

```text
PHP002 line 252: expected , (at "->")
exit status 1
```

The final line is from `go run`. No output/report destination was requested.
No extraction, rewrite, concatenation, wrapper or replacement algorithm was
used. Earlier syntax can now be parsed; this is not semantic or runtime proof
for those original methods. The next measured frontend boundary is the original
`$this->document` regex input. P2 stack mutation, single unknown-receiver writes,
class parameters and snapshots also remain open; P3–P5 are incomplete.

At the same clean commit, the focused local checks ran once:

```sh
CGO_ENABLED=0 GOCACHE=CACHE GOPROXY=off GOTOOLCHAIN=go1.23.0 \
go test -count=1 -run '^(TestNestedProperty|TestPublicPropertyWrite)' ./internal/codegen ./runtime ./tests
CGO_ENABLED=0 GOCACHE=CACHE GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./internal/compat
CGO_ENABLED=0 GOCACHE=CACHE GOPROXY=off GOTOOLCHAIN=go1.23.0 go vet ./internal/frontend ./internal/semantic ./internal/codegen ./runtime
```

PASS: codegen0.620s, runtime0.414s, generated guards3.726s, catalog0.426s; vet
exit0. The old compiler-only null-intermediate negative is now covered by the
new runtime guard; the retained source negative checks an undeclared prefix.
All other existing single-property admission remains unchanged.

## Native evaluation-order correction

Initial push40730860690 / PR40730925185 failed at commit2421daf. Retained PHP8.3
artifact12833924894 showed exactly two failing leaf tests: a stale
`TestParserOutputUnsupportedSource/append_chained` rejection and a real
`TestDifferential/nested_property` output difference. The original plain
assignment targeted the replacement child in PHP, while generated Go targeted
the old child. Native stdout314bytes/SHA
`5f91a48d872c0e5200b65d39ce40d751e68c4b46285f6bd5e01ee72e43f16b73`;
candidate323bytes/SHA
`1a0ee4c1ac7b1b673886973161b077542743e586f04d69f7516a97f2452294b0`,
both exit0/stderr0. The artifact is also retained locally for failure analysis.

A single isolated PHP8.3.36 probe with a declared old/new child and a RHS method
that replaces the child recorded:

| Operation | Old child after RHS replacement | New child |
|---|---|---|
| `=` | text remains `old` | text `rhs` |
| `.=` | text remains `old` | text `newrhs` |
| `[] =` | items remain empty | items contain `rhs` |

This is the measured named-chain contract: RHS runs before writable-path
resolution for all three operations. The probe used the cached admitted digest,
explicit PHP_VERSION8.3.36 guard, read-only/networkless/nonroot container,
display-errors stderr and exit0. It makes no claim for general PHP evaluation
order, magic paths, indexed addresses or other versions.

Codegen was corrected in6d5551b31d5c1ccded419931940a8a3bf4089534. The final
headf01730bc099b39522c219670673818abb0726a0a/tree
6799c9e49621ea8a2a6afa6a42dcd8da758d1d11 also clarifies fixture comments.
All three child-replacement operations are now in the owned differential
fixture. Existing unsupported chained append now checks an undeclared prefix;
null/intermediate failures belong to runtime guards. At final head, only changed
`TestNestedPropertyLowering` / `TestParserOutputUnsupportedSource` reran locally,
PASS0.386s. Runtime/other local successes were reused. Frontend and semantic
files are unchanged from the original probe, so that source-admission result is
retained without another invocation.

The owned differential fixture hash is
`7457885d9199ca47ae288f740683356cc8b5a8d8a3adb33bf6a643b1531cc04e`.
Independent PHP8.3/8.4 native comparison is tracked by PR122 CI; it is not a
Gossip full-parser result.

## Final native comparison

At the final clean head/tree above, [push40731756226](https://github.com/katsuakikureyama/GoHP/actions/runs/40731756226)
and PR40731765069 completed successfully, attempt1. Saved push artifacts were
read back once: Go1.23.12 Linux/amd64, CGO0; independent native PHP8.3.36 and
8.4.27 with the fixed target image digests. The owned nested-property fixture
passed parse, codegen, Go build, execution and differential comparison for both
targets. PHP and Go stdout are each385bytes, SHA-256
`1b8e50c3376d33af55cc375a673bef03289f1deb0ee12b6966c037b20f3b6747`,
both exit0 and quiet-success stderr0. General stderr/error/exception and
external side-effect equivalence remain unclaimed.

All seven new catalog checks passed. Across765 catalog checks, PHP8.3 has
747PASS +18 target-inapplicable; PHP8.4 has763PASS +2 target-inapplicable.
Artifact8.3 ID12835760958, ZIP digest
`sha256:b275bab76d721beab48b12dae7b5918a02b17d4b39b4ffbfed30a6c40f10e0e8`;
artifact8.4 ID12834967541, ZIP digest
`sha256:e00ffa87668f0e16fee49cf152c5d5c41d9f599160c450f26d8d20e273727657`.
Both expire2026-11-27. The initial failed comparison is retained above rather
than overwritten. These generic language fixtures do not establish execution
of the unchanged full WordPress parser.

PR122 merged into maina30709e51ecfe0231451afa121f19ab1ed247e97 after the
expected head and clean mergeability check. Main tree is checked against the
validated feature tree during integration; no additional manual suite is needed
for an identical tree.

Original files, existing manifests and consumer pins are unchanged. This Gossip
change adds documentation only; old native references/consumer PASS are reused.
Gossip CI remains unconfigured. Existing local inspect/WebMCP exclusion stays;
no route/authentication or deployment change.
