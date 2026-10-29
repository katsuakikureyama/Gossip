# Original parser string-empty admission — 2026-10-29

The complete unchanged Parser/Block/Frame closure passes prior string-empty354.
The next measured semantic rejection is the unknown-class `(array)` snapshot
of `$stack_top->block` at394. No original parser Go/report/build/run or
G83-PARSER acceptance is produced.

## One probe and exact identities

- WordPress/WordPress `54106f6c545fc8b881f6b39b7963002bece6c28c`,
  version7.2-beta1-64693; complete original sources and existing notice/provenance.
- Unchanged [closure manifest](../../gossip/tests/gohp/parser-source-closure.json),
  SHA256 `89f897c64146705559eae41bb50d49aff57abad0a11ee02d16d27fe78d64adaa`.
- GoHP [PR133](https://github.com/katsuakikureyama/GoHP/pull/133), clean probe
  code `62cc41fdc45bf2126acd7c54d70c28034d4557ec`, tree
  `17e6ab69577bd640904405d85759596b9865961a`, base main5776a7b.
- Cached Go1.23.0 DarwinARM64/CGO0, offline modules, existing shared cache.

One invocation, personal paths replaced by stable placeholders:

```sh
CGO_ENABLED=0 GOCACHE=CACHE GOPROXY=off GOSUMDB=off GOTOOLCHAIN=local GOWORK=off \
go run ./cmd/gohp -manifest GOSSIP_ROOT/gossip/tests/gohp/parser-source-closure.json \
  -source-root GOSSIP_ROOT -entry wp-includes/class-wp-block-parser.php \
  -php-version 8.3 -experimental-pcre2 -out FRESH/parser.go -report FRESH/report.json
```

Exit1, stdout0bytes, stderr114bytes, SHA256
`0aa3c061c7896e4450a840ff8a41e16ac9c042c6b8e496d72eae9ac57bb41dba`:

```text
wp-includes/class-wp-block-parser.php: PHP004 line 394: array cast requires a known generated class
exit status 1
```

The last line comes from go run. No output/report/staging artifacts; the empty
temporary directory was removed. Original add_block_from_stack appends
`(array) $stack_top->block` to output. That value's inferred class is unknown;
the present snapshot contract requires a statically known generated class.
Earlier admission is not proof of every preceding method's semantics/runtime.
Original generation/build/execution,12-case/24-call comparison and G83-PARSER
stay blocked; stack/snapshot/tokenizer/JSON contracts still need full-unit proof.

## Bounded empty extension

Source empty accepts one defined bare local containing an actual array or
string. Exact empty bytes and single ASCII zero are empty strings; other strings
are not. No undefined/address suppression, other scalar/object domains,
coercion or generic-call bypass. Neutral additive empty_var/VariableEmpty is
separate from unchanged array-only legacy empty_array IR/ArrayEmpty. IR0.1,
report/source-map policy and storage optimization are not expanded.

Focused local frontend/codegen/runtime/catalog and generated guard/fixture
checks all PASS once at clean code62cc41f:0.362/0.465/0.202/0.703/3.805s.
Exact commands and boundaries are in the
[GoHP evidence](https://github.com/katsuakikureyama/GoHP/blob/main/docs/evidence/string-empty-2026-10-29.md).
New owned fixture SHA256
`5bb597705d84495bb3082d43eea9cbac716e9f2ba4a6fb99977c904950e4c2ec`
observes zero strings, typed/untyped/arrow inputs, substring/local/!empty,
reassignment/widening, singleton falsy arrays and short circuit. Its native result
belongs to each exact-candidate PHP8.3/8.4 automatic CI, not this original probe.

Final canonical head `97f5ff8a8de73e18845e0b6d5bac73f1670ac408`, tree
`f60026f9e266c84d92f76e7f2e941d8dce1f15e7`;
[push40847242945](https://github.com/katsuakikureyama/GoHP/actions/runs/40847242945)
and [PR40847383843](https://github.com/katsuakikureyama/GoHP/actions/runs/40847383843)
attempt1 all4 target jobs SUCCESS. Saved own push artifacts12884734901 (SHA256
`45dad1d554942d19c803731858a335b288edc3e90209cb0a6aed737276c84c86`) and
12884744866 (SHA256 `b86649fba466fcd267ced0a5332f13469315b4f15a3b64a1676bdc93e5d20bdf`)
read once: exact clean head/tree, Go1.23.12Linuxamd64/CGO0 and pinned
PHP8.3.36/PHP8.4.27. All8 new checks PASS. New fixture parse/codegen/build/
execute/compare PASS: native/Go stdout112bytes equal, SHA256
`ee3942adf7bf9c325e5a3457142e7873c0405694cad4b9ddef9a9b86b6365c39`,
exit0/quiet stderr0. General stderr/exception/side-effect equivalence remains
not-checked.820total:8.3 802PASS+18outside /8.4 818PASS+2outside.
[PR result6492670229](https://github.com/katsuakikureyama/GoHP/pull/133#issuecomment-6492670229)
retains this record. Main efcdbc47ffeb506d2bec1f9e6d84f0c67c37332c has the same
tested tree; local main ff-only synchronized. Automatic main CI tracked
separately; no duplicate archive readback or manual successful rerun.

Earlier scope6492331798/PR133 is canonical. Later duplicate132 closed per
[coordination](https://github.com/katsuakikureyama/GoHP/issues/1#issuecomment-6492622357)
without merging, cherry-picking or combining its alternative empty_value/
EmptyValue design,114-byte fixture or d723bfd probe evidence. Its unmerged
branch/work/evidence remain preserved; all results above belong solely to133.

Existing131 and130 successful native/consumer results remain at their old pins
and are not manually rerun. This Gossip increment changes documentation only;
originals, all manifests/pins, auth/routes/deployment and CLI/WebMCP exclusion
stay unchanged. Gossip CI unconfigured; devboard unidentified/not synchronized.
