# Complete original parser generation and build — 2026-10-30

One unchanged complete Parser/Block/Frame closure now generates Go and a v0.2
report, then builds as Go. This is **not** original parse invocation,
enabled-native-PCRE execution, the12-case/24-call differential comparison or
[G83-PARSER](../php83-application-acceptance.md) acceptance.

## Exact original and compiler identity

- WordPress/WordPress `54106f6c545fc8b881f6b39b7963002bece6c28c`, complete
  original classes, version7.2-beta1-64693, original notice/provenance bytes.
- Same [closure manifest](../../gossip/tests/gohp/parser-source-closure.json),
  SHA256 `89f897c64146705559eae41bb50d49aff57abad0a11ee02d16d27fe78d64adaa`.
  The compiler verified the selected sources/notices once before assembly.
- [GoHP138](https://github.com/katsuakikureyama/GoHP/pull/138), clean probe
  code `35524a636073f4abbc8bb69b9e36263e71ac62d8`, tree
  `2bec82d98e3ca2c36031d5aa5bd3f5e46c0f2a88`, base mainaf3e9ad.
- Cached Go1.23.0 Darwinarm64/CGO0/offline modules, existing shared cache.

One generation, followed by one build, personal paths replaced:

```sh
CGO_ENABLED=0 GOCACHE=CACHE GOPROXY=off GOSUMDB=off GOTOOLCHAIN=local GOWORK=off \
go run ./cmd/gohp -manifest GOSSIP_ROOT/gossip/tests/gohp/parser-source-closure.json \
  -source-root GOSSIP_ROOT -entry wp-includes/class-wp-block-parser.php \
  -php-version 8.3 -experimental-pcre2 -out FRESH/parser.go -report FRESH/report.json
go build -o FRESH/parser-build FRESH/parser.go
```

Both exits0. Generation stdout/stderr0bytes. The v0.2 report records all three
original source identities, once-only edges403/410 and per-file statement maps.
Generated Go SHA256 `55c66c931ed9da4be7a0a263129242b169a68127e83bb526fbc1f6d247f402dd`;
report SHA256 `bbf61c482f333b6509cff72b8e27cd2c7833238bf81a4148dda3e4f08104a1df`.
The temporary Mach-O binary was removed after the one build; generated Go,
report and logs remain outside source worktrees. No persistent binary/new cache.
The generated empty main was not executed as fake original-parser evidence.

The former json_decode283 rejection is gone, and the complete three-file
generated program builds, including declarations after that source line.
Runtime behavior of those original methods remains unverified. Next is the
real generated parse entry/observer plus explicitly enabled native PCRE
configuration and the fixed12-case/24-call comparison, with unchanged source.
Old bounded regex and method results cannot substitute for that combined run.

## Separate bounded JSON evidence

The [GoHP contract](https://github.com/katsuakikureyama/GoHP/blob/main/docs/json-associative-decode.md)
admits only actual-string/literal-true two-argument mode, default depth512 and
ordered nested Array/string/bool/null/int64 values. Valid supported input resets
execution-local JSON error; malformed input returns null and code/message.
Fraction/exponent/out-of-int64 tokens on valid JSON explicitly fail without
silent conversion or error publication. Flags/custom depth/object mode/general
floats and exception equivalence remain excluded. Error readers were included
because the original reference observer already records both code and message.

Initial cleanffdd4df codegen0.268/catalog0.304s PASS; runtime/generated-host build
failed on uint16 intermediate overflow in surrogate arithmetic. Widening before
combining at35524a6 fixes it; only affected runtime/host checks rerun0.216/0.627s
PASS. Other successful checks reused; exact commands/failed history in
[GoHP evidence](https://github.com/katsuakikureyama/GoHP/blob/main/docs/evidence/json-associative-decode-2026-10-30.md).
Ten new catalog checks include a separate owned fixture SHA256
`d4efaa01dbe151eabf146e89a5fe54e96134fda7fb8d9ede5e1a618b8b2f9600`.
Native exact-candidate PHP8.3/8.4 CI observations are separate from this original
generation/build. Local PHP/Docker unavailable; no new daemon/image configured.

Final candidate `e90460e8074ae3a09e22bd0939ad0cd2ca3e945f`, tree
`25914fb81df3e7ca9a545a6e459946f5f293c273`, merged at GoHP main
`b2d1482268a7710c515c07bf919af3d40b09d17c`.
[Push41021437634](https://github.com/katsuakikureyama/GoHP/actions/runs/41021437634)
and [PR41021444154](https://github.com/katsuakikureyama/GoHP/actions/runs/41021444154)
attempt1 all4 jobs SUCCESS before merge. Saved own push artifacts12960939634/
12962145737 read once: exact clean head/tree, Go1.23.12linuxamd64/CGO0 and pinned
PHP8.3.36/8.4.27, test_exit0/overallPASS. All10 new checks PASS independently;
855total:8.3 837PASS+18unmeasured /8.4 853PASS+2unmeasured.
Owned native/Go stdout812bytes, SHA256
`e9b26df6f07f7b25f0ab2668457038e76c48c7c8758d77c7bb110bf2ded4f76c`
equal, exit0/quietstderr0; all parse/codegen/build/execute/diff phases PASS.
General diagnostics/exception/sideeffect equivalence remains unmeasured.
[Exact PR result6517582871](https://github.com/katsuakikureyama/GoHP/pull/138#issuecomment-6517582871)
retains the artifact metadata and complete boundaries; no repeated archive or
successful test/probe/build run. This owned JSON result is not original execution.

Original files/all manifests/pins/auth/routes/deployment remain unchanged.
Existing137 and earlier successes/optional-native-PCRE evidence not rerun.
Gossip CI unconfigured, developer-local CLI/WebMCP exclusion unchanged;
devboard unidentified/not synchronized. No original acceptance or deployment.
