# Named-property isset admission — 2026-10-29

The complete unchanged original parser passes property isset384 and next stops
at its first `require_once`403. Source admission is not full semantic validation,
generation/build/execution or [G83-PARSER](../php83-application-acceptance.md)
acceptance. Original files, loading edges and algorithms remain unchanged.

- WordPress source `54106f6c545fc8b881f6b39b7963002bece6c28c`.
- Complete `wp-includes/class-wp-block-parser.php` SHA256
  `8563621090de5e307431b1d29b78a9f75fd9ee238a0137b563a7db7b2e1f8d3b`.
- GoHP [PR126](https://github.com/katsuakikureyama/GoHP/pull/126), clean code
  `53546a681006a42fe3cbde03bf2b84a15bb19843`, tree
  `a8bf5182398c63d83d0f99355b7b7c56be7618b1`, base maine5d32e2.
- Go1.23.0 DarwinARM64, CGO0, cached toolchain and offline modules.

One invocation; personal paths replaced by stable placeholders:

```sh
cd GOHP_ROOT
CGO_ENABLED=0 GOCACHE=CACHE GOPROXY=off GOSUMDB=off GOTOOLCHAIN=local GOWORK=off \
go run ./cmd/gohp -in GOSSIP_ROOT/wp-includes/class-wp-block-parser.php \
  -php-version 8.3 -experimental-pcre2
```

Exit1, no generated/report destination requested:

```text
PHP002 line 403: unsupported statement (at "require_once")
exit status 1
```

The last line is from `go run`. The first rejected original statement loads
`__DIR__ . '/class-wp-block-parser-block.php'`; the second Frame dependency is at410.
No extraction, concatenation, source rewriting or handwritten parser substitute
was used. Earlier syntactic progress is not proof of earlier semantic/runtime
acceptance. P2 stack/single-unknown-receiver/snapshot requirements and remaining
tokenizer/JSON integration are still open. A bounded real loading design is next.

Focused local new/affected compiler checks PASS0.351s, runtime0.532s, generated
external ABI2.810s; catalog0.328s, all once at code53546a6. Owned PHP fixture SHA256
`78dbb63a7a03b8d786e8adcb3d37141b81383b0555d9ddf9a2cd5faf87d4663c`.
Evidence/docs final GoHP head `2c6c2c873b08800f8f10dc12c06aa17fd067dca8`, tree
`0c6cc7b810bd06886c9ec03326be99629366d4a6`, retains the implementation/config.
[Final push40805015360](https://github.com/katsuakikureyama/GoHP/actions/runs/40805015360)
and [PR40805022142](https://github.com/katsuakikureyama/GoHP/actions/runs/40805022142)
attempt1 both SUCCESS. Saved push artifacts12867259317 /12867991731 read once:
exact clean head/tree, Go1.23.12 Linuxamd64/CGO0 and pinned PHP8.3.36/PHP8.4.27.
New seven checks PASS. Owned fixture parse/codegen/build/execute/compare PASS,
102stdout bytes equal each native target and Go, exit0/quiet stderr0; stdout SHA256
`9232e46d80e37ac1ff46ee678836a93b0d22547a7312ca80900f622c44b4b4bc`.
781 checks total:8.3 763PASS+18outside /8.4 779PASS+2outside. General PHP side
effects/exceptions remain unmeasured. This owned-fixture evidence is not original
parser acceptance. [PR result](https://github.com/katsuakikureyama/GoHP/pull/126#issuecomment-6486025944)
completes the pre-CI local checkpoint without another manual successful test run.
PR126 merged into main9eec4d502ae396e389f94bf40d986b8d933caac0; tree readback equals
`0c6cc7b810bd06886c9ec03326be99629366d4a6`, the tested feature tree. Automatic
main CI is tracked separately; identical-tree artifacts need no second download.
No enabled-PCRE or old successful consumer/native suite remeasurement.

This Gossip increment changes documentation only. Original-source non-overlay
equality, documentation links and diff checks are verified before publication.
Gossip CI remains unconfigured; no consumer repin or route/auth/deployment change.
Existing local inspect capability and WebMCP/MCP exclusion remain. Devboard is
unidentified/not synchronized.
