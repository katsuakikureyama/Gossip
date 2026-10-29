# Parser array_pop admission probe — 2026-10-30 (JST record)

This evidence records one compiler-admission probe for the fixed Gossip PHP8.3
target. It does not record generated parser source, a native PHP/Go comparison,
or application acceptance.

## Candidate boundary

Merged GoHP137 admits only `array_pop($target)` as a standalone statement
or as the complete RHS of a bare-local result assignment. The target is an
existing bare local or a direct named public property of a defined bare local,
including `$this->stack`. For a property, generated code resolves the owner
before reading the actual target. It removes the final insertion entry without
renumbering the remaining keys, then binds the popped value (or `null` from an
empty array) last. Empty arrays publish no mutation. Popped nested arrays are
copied eagerly; child-object identity is retained.

Known non-array and undefined targets reject statically; inferred targets that
are actually non-arrays reject at runtime. Generic embedded calls, general
references, indexed/nested/dynamic/effectful targets, non-bare result targets,
builtin redeclaration, and general stack semantics remain outside this
increment. The shared index rewinds only when the removed integer key immediately
precedes its current next key. Negative keys retain that rule; minimum-int rewind
restores the sentinel and automatic0. An
occupied MAX key cannot be overwritten after a MAX-1 pop. This is an additive
IR0.1 statement, not a new storage/report/version lane or general reference
value model.

## Fixed inputs

| Item | Value |
| --- | --- |
| Initial clean compiler code | `1adc11ad053a19e22b98fcecefd898851830dd12` |
| Initial clean compiler tree | `00adfc030a89eea4fcf57f3fbd2eb2933925c4e8` |
| Fixed WordPress source | `54106f6c545fc8b881f6b39b7963002bece6c28c` (`7.2-beta1-64693`) |
| Closure | Complete unchanged `WP_Block_Parser`, `WP_Block_Parser_Block`, and `WP_Block_Parser_Frame`, with existing notice/provenance files |
| Manifest | [`parser-source-closure.json`](../../gossip/tests/gohp/parser-source-closure.json), SHA256 `89f897c64146705559eae41bb50d49aff57abad0a11ee02d16d27fe78d64adaa` |

The manifest and the five referenced source file sizes/hashes were verified
before the probe. No source extraction, concatenation, handwritten parser, or
interpreter fallback was used. The manifest, original source, and consumer pin
were not modified.

## One probe

```sh
go run ./cmd/gohp \
  -manifest GOSSIP_ROOT/gossip/tests/gohp/parser-source-closure.json \
  -source-root GOSSIP_ROOT \
  -entry wp-includes/class-wp-block-parser.php \
  -php-version 8.3 \
  -experimental-pcre2 \
  -out FRESH/parser.go \
  -report FRESH/report.json
```

The command was invoked once at the initial clean compiler revision. It exited
1 with stdout0 bytes and stderr103 bytes, SHA256
`97d870019d0d48ac7e09006c80a7b07877b61bad3e74c4b9491b984b95ee8f4c`:

```text
wp-includes/class-wp-block-parser.php: PHP004 line 283: unsupported function json_decode
exit status 1
```

The prior codegen rejection at original line205,
`array_pop($this->stack)`, was passed. This does not demonstrate successful
generation of the later original `array_pop` at line372. No Go source, report, or staging
artifact was produced; the empty temporary output directory was removed. The
next measured blocker is original `json_decode($matches['attrs'][0], true)` at line283.

## Candidate verification and limits

Final [GoHP PR137](https://github.com/katsuakikureyama/GoHP/pull/137) candidate
is `5fd2a5c52a72f2e1654553bbab3c21d1d5cfc218`, tree
`072ff4fde72bbc2a51622f5ecc447ce89910eb30`; its GoHP contract evidence is
`docs/evidence/array-pop-2026-10-30.md`. Generated fixture/host `TestArrayPop`
passed once in1.413s. Its fixture SHA256 is
`b5199cc408c31aa5d1aa89690e77c2689fe8f9d9e1bc7cc1a755dec96409b037`.
A first local command across arraykey/runtime/storage/frontend/codegen/compat
passed except for a source-map test fragment with same-named test locals;
test-only clean code `2325ea323b5a092f4ce11bf8313519a061e87c8c`, tree
`efecf327277de4f41382ae4a6c5eab5e9d14df74`, renames the locals. The affected
`TestArrayPopLoweringAndSourceMap` rerun passed in0.208s; prior successful
checks are reused. The final candidate adds ten checks. Its documentation has
166 checked local links.

A separate seven-case native PHP8.3.36 design diagnostic informed the bounded
return/null, key, and copy contract. It is not this fixture's PHP/Go comparison,
is not added to a comparison count, and does not establish original-parser
runtime behavior.

Push40928765568 and PR40928781175 attempt1 all four jobs succeeded at the exact
candidate tree, using Go1.23.12 Linux amd64/CGO0 and PHP8.3.36/8.4.27. All ten
new checks passed on both targets. The owned fixture's native and Go stdout are
equal at902 bytes, SHA256
`304031532fa6741bf1a9d0a921e72240971d105e571c81f61fcc736bd8ed9bc9`, with
exit0 and quiet stderr0. Parse/codegen/build/execute/differential pass for that
owned fixture; general stderr, exception, and side-effect equivalence remain
not-checked. The PHP8.3 push artifact12922101035 (published SHA256
`5a3296c59ce9795adb81dcead70f023c37bf5cce242db24878cdf27a7514a41d`) and
PHP8.4 artifact12921498050 (published SHA256
`9ec27b7aef5a4e6b1c7d361b937047fbb8ca1d1069874c38b20284ff360baaf4`) were
each read once. It is merged at main
`af3e9ada36728cb317096b35af6ca0f22e3492c4` with the same tested tree; its
[canonical native result](https://github.com/katsuakikureyama/GoHP/pull/137#issuecomment-6504793559)
is recorded on the PR. [Exact-main CI40929318757](https://github.com/katsuakikureyama/GoHP/actions/runs/40929318757)
attempt1 also succeeds at that merged main, with PHP8.3 job123643210418 and
PHP8.4 job123643210914. No same-tree successful artifacts or checks were
downloaded/run again. The original-source result remains only an unchanged-source
compiler-admission measurement. It does not prove original Parser/Block/Frame
generation, binary build, parser runtime, the declared 12 native reference
cases/24 original calls, differential output, or G83-PARSER acceptance.
`json_decode`, remaining stack/snapshot behavior, tokenizer/PCRE integration,
JSON attributes, and all later application gates remain open. General exception,
stderr, and side-effect equivalence are unmeasured.

Gossip has no configured CI. There is no service, resource, auth, route,
deployment, browser operation, or WebMCP publication: local
`gohp.cli.generate` remains explicitly excluded.
