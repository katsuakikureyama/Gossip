# Full-parser admission after list assignment — 2026-10-29

Result: **blocked at switch**, not generated-parser acceptance. This advances
the [previous list rejection](parser-relational-probe-2026-10-27.md), without
rerunning or overwriting any native reference or consumer comparison.

- Original WordPress `54106f6c545fc8b881f6b39b7963002bece6c28c`.
- Complete input `wp-includes/class-wp-block-parser.php`, SHA-256
  `8563621090de5e307431b1d29b78a9f75fd9ee238a0137b563a7db7b2e1f8d3b`.
- GoHP [PR120](https://github.com/katsuakikureyama/GoHP/pull/120) clean probe
  commit `adb23a019599341f9d23783da3f37990fae65dcd`, tree
  `6ae49b61e0d637c568b93f306202f0a764c94329` (base `b4944f7`).
- Final feature `de8fa627c9518c546c0abc6fde3d25ebf3d4e224`, tree
  `a64df32d6eb1c173629a1b56576826e382f4b2ac`: only two compiler-test expectations
  changed after the probe; production code unchanged. No duplicate probe.
- Go1.23.0 Darwin ARM64, CGO0, cached/offline modules.

Personal local paths replaced with placeholders:

```sh
cd GOHP_ROOT
GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOWORK=off CGO_ENABLED=0 \
GOCACHE=CACHE go run ./cmd/gohp \
  -in GOSSIP_ROOT/wp-includes/class-wp-block-parser.php \
  -php-version 8.3 -experimental-pcre2
```

One invocation, exit **1**:

```text
PHP002 line 100: expected ; (at "{")
exit status 1
```

The second line is from `go run`. No output/report destination was requested.
No source extraction, wrapper, concatenation, rewrite or substitute algorithm.
List assignment at94 is now admitted syntactically; switch at100 stops parsing.
This does not establish semantic validation of the remainder, generated-source
success, native binary build, full-parser execution or application acceptance.

The separate GoHP command `go test -count=1 -run '^TestList(Assignment|Values)'
./internal/codegen ./runtime ./tests` initially had six successful checks and
one failed source-negative expectation: unknown property reads are already
admitted and runtime guarded, so expecting static rejection was incorrect.
Only the two stale-class test inputs were changed to property **writes**, which
require a statically known class. Only `TestListAssignmentUnsupportedSource`
was rerun and passed (0.329s); other PASS results were retained (runtime0.530s,
generated guards2.024s). Native 8.3/8.4 differential evidence is tracked by PR120
CI and is not a Gossip parser result.

This Gossip change is documentation only: no original source, fixture, manifest,
compiler pin or deployment change. Next P1 switch dispatch; P2–P5 remain open.
Existing developer-local inspect capability and WebMCP exclusion are unchanged.
