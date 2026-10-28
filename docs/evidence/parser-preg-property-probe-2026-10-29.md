# Regex property-input admission — 2026-10-29

The complete unchanged parser passes `preg_match` property inputs252/255 and
stops at named class parameter349. This advances source admission only; it is
not complete semantic validation, generated full-parser execution or acceptance
of [G83-PARSER](../php83-application-acceptance.md).

- WordPress `54106f6c545fc8b881f6b39b7963002bece6c28c`.
- Whole `wp-includes/class-wp-block-parser.php` SHA256
  `8563621090de5e307431b1d29b78a9f75fd9ee238a0137b563a7db7b2e1f8d3b`.
- GoHP [PR123](https://github.com/katsuakikureyama/GoHP/pull/123), clean executable
  commit `8403235e0e14fb460354bb73c9400b7e94708004`, tree
  `f0b982ac33665bdcb46f71ac9de9bbffb71d393e`, based on maina30709e.
- Go1.23.0 DarwinARM64, CGO0, cached/offline modules.

One invocation (personal paths replaced by stable placeholders):

```sh
cd GOHP_ROOT
CGO_ENABLED=0 GOCACHE=CACHE GOPROXY=off GOTOOLCHAIN=go1.23.0 \
go run ./cmd/gohp -in GOSSIP_ROOT/wp-includes/class-wp-block-parser.php \
  -php-version 8.3 -experimental-pcre2
```

Exit1, no output/report destination:

```text
PHP002 line 349: unsupported parameter declaration (at "block")
exit status 1
```

The final line is from `go run`. No source extraction, rewriting, concatenation,
wrapper or substitute algorithm. This is the next frontend stop, not proof
that all earlier original statements pass semantic/runtime checks. P2 class
parameters, stack mutation/single unknown-receiver writes/snapshots and the
remaining tokenizer, JSON and loading closure still need full-unit admission.

At that same clean code, focused compiler/CLI/generated-property-guard checks
PASS0.459/2.296/4.169s, catalogPASS0.896s, affected vet exit0. The separate
enabled static-PCRE2 versus fresh isolated PHP8.3.36 comparison PASS6.02s
(package6.341s),15 typed ordered observations. New property rows cover `$this`,
nested UTF-8/negative byte offset, changed inputs between calls, capture copies
and unchanged input properties; shared Execution resets error2 through a
property-input success. Both children exit0/stderr0. Comparison is structural
typed equality, not serialized JSON byte equality.

Exact native toolchain/image/limits, corrected launcher setup failure, fixture,
generated source and observer hashes, and raw native output are retained in
[GoHP's measurement](https://github.com/katsuakikureyama/GoHP/blob/da5715159016b79d4c4dac727c5d7efe5dfff230/docs/evidence/preg83-property-inputs-2026-10-29.md).
Final headda5715159016b79d4c4dac727c5d7efe5dfff230/tree
0e47ab7dedcaecb81e81cefc9c6649a9b3827dbe only adds evidence/docs; executable
successes and the source probe are reused, not repeated. Default Linux CGO0 CI
is a separate regression gate, not enabled PCRE or PHP8.4 regex evidence.

Original WordPress files, old manifests, fixture/consumer pins and native
reference results are unchanged. This Gossip increment is documentation-only;
old successful native/consumer tests were not rerun. Gossip CI is unconfigured.
Existing local inspect/UI/WebMCP/MCP exclusion remains, with no route,
authentication or deployment change. Devboard remains unidentified/not synced.
