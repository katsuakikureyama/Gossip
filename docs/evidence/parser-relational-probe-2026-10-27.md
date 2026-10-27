# Full-parser admission after relational syntax — 2026-10-27

Result: **blocked at list assignment**, not generated-parser acceptance.
This supersedes the first-rejection location in the historical
[baseline probe](parser-full-baseline-2026-10-27.md); its native reference remains
unchanged. No PHP reference or successful consumer comparison was rerun.

- Original WordPress commit: `54106f6c545fc8b881f6b39b7963002bece6c28c`.
- Complete input: `wp-includes/class-wp-block-parser.php`, SHA-256
  `8563621090de5e307431b1d29b78a9f75fd9ee238a0137b563a7db7b2e1f8d3b`.
- GoHP [PR #119](https://github.com/katsuakikureyama/GoHP/pull/119) tested clean
  source `1d7cd3e9bab88b16a7c9a26fdce40457b92c2421`, tree
  `c333e5ddcb27698a5b0740f09ea948141046cd99`, based on `b424667`.
- Host: Go1.23.0 Darwin ARM64, CGO0, cached module/toolchain only.

Personal local paths are replaced by `GOHP_ROOT`, `GOSSIP_ROOT`, `CACHE`:

```sh
cd GOHP_ROOT
GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOWORK=off CGO_ENABLED=0 \
GOCACHE=CACHE go run ./cmd/gohp \
  -in GOSSIP_ROOT/wp-includes/class-wp-block-parser.php \
  -php-version 8.3 -experimental-pcre2
```

One invocation, exit **1**. Diagnostic:

```text
PHP002 line 94: expected ; (at "=")
exit status 1
```

The second line comes from `go run`. No output/report destination was requested;
generation failed, so no generated parser binary or comparison exists. This
probe used the full original file, with no selected-method wrapper, concatenation,
source rewrite or substitute parser. The earlier lexer rejection at line98 is
gone; the parser now reaches line94's `list` assignment. Earlier line number does
not mean regression: lexing formerly failed before statement parsing began.

GoHP's owned unit/compiler/generated-guard command at the same clean source:

```sh
go test -count=1 -run '^TestRelational' ./internal/codegen ./runtime ./tests
```

PASS once (0.624s / 0.431s / 1.962s), six top-level checks. It covers guarded
int64/null `<` and `>`; it does not certify general ordering or the parser.
The independent native 8.3/8.4 comparison is tracked in PR119 CI, not relabelled
as a Gossip full-parser result. This Gossip changeset is documentation only:
no fixture/compiler pin, manifest, source, runtime or deployment change.

Next: P1 list destructuring, then switch dispatch. P2 object/stack mutation,
P3 regex properties, P4 JSON and P5 original loading remain open.
Existing local inspect capability and WebMCP exclusion remain unchanged.
