# Declared class-parameter admission — 2026-10-29

The complete unchanged parser passes its named class parameter at349 and next
stops at property-expression `isset`384. This is source admission only, not
full semantic validation, generated execution or [G83-PARSER](../php83-application-acceptance.md)
acceptance. Original source, loading edges and algorithms are not rewritten.

- WordPress source commit `54106f6c545fc8b881f6b39b7963002bece6c28c`.
- Complete `wp-includes/class-wp-block-parser.php` SHA256
  `8563621090de5e307431b1d29b78a9f75fd9ee238a0137b563a7db7b2e1f8d3b`.
- GoHP [PR124](https://github.com/katsuakikureyama/GoHP/pull/124), clean code
  `35bfcaac799f4cce8a3f0c5fae71cc9acd4cdd60`, tree
  `bdb6d517712789aade29aea07001246ee06bff67`, base main2045ab2.
- Go1.23.0 DarwinARM64, CGO0, cached/offline modules.

One invocation (local paths replaced with placeholders):

```sh
cd GOHP_ROOT
CGO_ENABLED=0 GOCACHE=CACHE GOPROXY=off GOSUMDB=off GOTOOLCHAIN=local GOWORK=off \
go run ./cmd/gohp -in GOSSIP_ROOT/wp-includes/class-wp-block-parser.php \
  -php-version 8.3 -experimental-pcre2
```

Exit1, no output/report destination requested:

```text
PHP002 line 384: expected , (at "->")
exit status 1
```

The final diagnostic is the `go run` wrapper. The failing original expression
is `isset($stack_top->leading_html_start)`. Source parsing stopped before the
full unit could reach class resolution, generation or execution. In particular,
passing the parameter syntax does not mean the Block dependency is loaded or
that all earlier parser statements pass semantic checks. Property presence,
stack writes, remaining object/snapshot requirements, JSON and loading remain.

Focused compiler/source/raw-IR and affected existing parameter checks PASS0.288s;
catalog tests PASS0.414s. Delegated external ABI test PASS1.208s before commit,
reused without rerunning unchanged code. Final documentation-only GoHP head
`148198c2abbace26eacbc7f5fe87c1ba58dca2a4`, tree
`1ce495eb562416da94c165f83ae47dc5768e6b96`, retains those successes without
another manual run. [Push CI40797216017](https://github.com/katsuakikureyama/GoHP/actions/runs/40797216017)
and [PR CI40797252113](https://github.com/katsuakikureyama/GoHP/actions/runs/40797252113)
both attempt1 SUCCESS on pinned PHP8.3.36 and PHP8.4.27 independently.
Saved push artifacts12864980251 /12863319778 were read once: exact clean head/tree,
Go1.23.12 Linux/amd64, CGO0. All five new checks PASS. The owned fixture's
parse/codegen/build/execute/compare phases pass,146 stdout bytes equal Go on each
target, exit0/quiet stderr0; stdout SHA256
`d8da26a1eafd57ae95742c7e1a0fbc888bbd12d183d7a854d0705ac52fd6bf16`.
Catalog774 total:8.3 756PASS+18 inapplicable /8.4 772PASS+2 inapplicable.
General side effects and PHP exception equivalence remain unmeasured. This
generic owned-fixture evidence does not certify the original parser. No
enabled-PCRE comparison rerun. [PR result](https://github.com/katsuakikureyama/GoHP/pull/124#issuecomment-6484881929)
completes the pre-CI local measurement chronology without rerunning successes.
PR124 merged into maine5d32e2087b63c13ecfe5f9ef785a60d65ed5b9c; readback tree
`1ce495eb562416da94c165f83ae47dc5768e6b96` equals the checked feature tree.
Automatic main CI is tracked separately; no successful suite/probe rerun or
same-tree artifact download is needed to confirm this integration.

This Gossip change is documentation-only. Old native references and all
consumer/fixture pins are retained. Gossip CI is unconfigured. Existing local
CLI inspect capability and WebMCP/MCP exclusion remain; no route/auth/deploy
changes. External devboard remains unidentified/not synchronized.
