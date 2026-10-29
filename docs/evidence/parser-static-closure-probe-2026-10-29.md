# Complete original static closure admission — 2026-10-29

The original Parser/Block/Frame now pass per-file parsing and their two original
`require_once` edges resolve through the explicit checked closure. The next
rejection is **semantic**, entry209: an unknown generated-class receiver write.
This is not full-parser generation, native equivalence or G83-PARSER acceptance.

## Source and compiler identity

- WordPress/WordPress `54106f6c545fc8b881f6b39b7963002bece6c28c`,
  version7.2-beta1-64693.
- Separate new whole-file [closure manifest](../../gossip/tests/gohp/parser-source-closure.json),
  exact SHA256 `89f897c64146705559eae41bb50d49aff57abad0a11ee02d16d27fe78d64adaa`.
  Complete Parser11539bytes, Block2555bytes and Frame1999bytes plus license and
  version provenance retain every hash from the old native baseline. No source
  rewriting, method extraction, concatenation, substitute parser or changed pin.
- GoHP [PR130](https://github.com/katsuakikureyama/GoHP/pull/130), clean probe
  code `532e16c043d488d91269026626037b2c7d210e46`, tree
  `f30f55406ca24b891d47768dbff4bf01be4d6827`, base mainb1ba044.
- Cached Go1.23.0 DarwinARM64, CGO0, offline modules and the existing shared cache.

## Command and result

From the clean GoHP worktree, personal paths replaced by placeholders:

```sh
CGO_ENABLED=0 GOCACHE=CACHE GOPROXY=off GOSUMDB=off GOTOOLCHAIN=local GOWORK=off \
go run ./cmd/gohp -manifest GOSSIP_ROOT/gossip/tests/gohp/parser-source-closure.json \
  -source-root GOSSIP_ROOT -entry wp-includes/class-wp-block-parser.php \
  -php-version 8.3 -experimental-pcre2 -out FRESH/parser.go -report FRESH/report.json
```

Exit1, stdout0bytes (empty SHA256
`e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`),
stderr132bytes (SHA256
`f0d26b1d56accceea51bdbd3ec6ce3f2e3fd5036187fb9abe357f530f196238b`):

```text
wp-includes/class-wp-block-parser.php: PHP004 line 209: property assignment requires a known generated class variable
exit status 1
```

The last line comes from `go run`. No output/report/staging artifact exists; the
empty temporary directory was removed. The final manifest was probed once.
An initial invocation had the wrong upstream repository label
WordPress/wordpress-develop, manifest SHA1f0d46fb66a519921697c82dd20f6f2e6330046930501bc680d6aacc8e9320d5.
It gave the same error; provenance was corrected to WordPress/WordPress before
the recorded final-manifest invocation. No source/target/edge bytes changed.

## Interpretation and remaining gates

Original requires403/410 are no longer the current full-source blocker. All
three complete source files parse. After `$stack_top = array_pop($this->stack)`,
the write `$stack_top->prev_offset = ...` at209 lacks a statically known class.
Earlier parsed syntax is not proof of full semantic/runtime correctness.
P2 unknown-receiver writes, stack mutation and snapshots remain open, together
with tokenizer/JSON integration. No original parser Go was generated or built;
execution, full12-case/24-call comparison and G83-PARSER remain blocked/not-run.

GoHP adds an owned real-file fixture separately: nested diamond/repeated loads,
shared-object mutation, falsy property presence and copied arrays. Its manifest
SHA256 is `a849f22ebb68aac7854c1b3f851da7efab80a34e0b30144c28da2c7aebc17c97`.
Local focused checks pass after documented test/data fixes; native PHP8.3/8.4
results come from exact-candidate automatic CI, not this original-source probe.
See the [compiler evidence](https://github.com/katsuakikureyama/GoHP/blob/main/docs/evidence/static-declaration-closure-2026-10-29.md).

Final GoHP head `f731a688e96f5a78ea3d21e06b41c2a3273eca52`, tree
`83b834cef44d60ad15ff307b2752640ee2830e50`; automatic
[push40821198551](https://github.com/katsuakikureyama/GoHP/actions/runs/40821198551)
and [PR40821236205](https://github.com/katsuakikureyama/GoHP/actions/runs/40821236205)
attempt1 both SUCCESS. Saved push artifacts12874488326 (SHA256
`d62799f4ef48eaf016be53163e5b7f203134300e4624e679778af36728d2a40f`) and
12875035099 (SHA256 `712e3ba0311487de8cd462580788ce5d8975af8adf1a4d15ec818a5294af7d57`)
read once: exact clean head/tree, Go1.23.12Linuxamd64/CGO0, pinned
PHP8.3.36/PHP8.4.27. All9 new checks PASS on both. Owned fixture
parse/codegen/build/execute/compare PASS; native and Go stdout64bytes, SHA256
`f72a8b7c607552651560adc5f3cbc73d600e046c22ae55ff87787c7b3d1943a6`,
exit0, quiet stderr0. General exception/side-effect/stderr equivalence remains
not-checked.804 total checks:8.3 786PASS+18not-measured /8.4 802PASS+2not-measured.
The [PR result](https://github.com/katsuakikureyama/GoHP/pull/130#issuecomment-6488474050)
retains this evidence. GoHP merged at main7ee6ec39eed56d2926f308566eb5c2cf3b4ed064,
with identical tested feature tree; local main ff-only synchronized and clean.
Automatic main CI is tracked separately; no duplicate archive readback or manual
successful suite rerun. This completes bounded L0/L1, not G83-PARSER.

This Gossip increment adds only documentation and the separate measurement
manifest. Old native/consumer evidence is reused without rerun. Consumer pins,
original files, route/auth/deployment and local inspect/WebMCP exclusion remain
unchanged. Gossip CI is unconfigured. Devboard unidentified/not synchronized.
