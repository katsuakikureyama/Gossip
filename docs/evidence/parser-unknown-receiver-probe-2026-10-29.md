# Original parser unknown-receiver admission — 2026-10-29

The complete unchanged Parser/Block/Frame closure passes the previous209
unknown receiver property-write rejection. The next measured semantic blocker
is **string empty at354**, not source loading or this property assignment.
No original parser Go, report, build or execution result is produced.

## Identities and one probe

- WordPress/WordPress `54106f6c545fc8b881f6b39b7963002bece6c28c`,
  version7.2-beta1-64693, complete three original sources plus notice/provenance.
- Existing [closure manifest](../../gossip/tests/gohp/parser-source-closure.json)
  SHA256 `89f897c64146705559eae41bb50d49aff57abad0a11ee02d16d27fe78d64adaa`,
  unchanged from the previous [loading probe](parser-static-closure-probe-2026-10-29.md).
- GoHP [PR131](https://github.com/katsuakikureyama/GoHP/pull/131), clean probe
  code `0f448b6609f15b90d6853be6dc2db4565edb7940`, tree
  `8e9403adea224a88e18e42c60b709673bfdd50e9`, base main7ee6ec3.
- Cached Go1.23.0 DarwinARM64/CGO0, offline modules, existing shared cache.

One invocation from the clean GoHP worktree; personal paths replaced by stable
placeholders:

```sh
CGO_ENABLED=0 GOCACHE=CACHE GOPROXY=off GOSUMDB=off GOTOOLCHAIN=local GOWORK=off \
go run ./cmd/gohp -manifest GOSSIP_ROOT/gossip/tests/gohp/parser-source-closure.json \
  -source-root GOSSIP_ROOT -entry wp-includes/class-wp-block-parser.php \
  -php-version 8.3 -experimental-pcre2 -out FRESH/parser.go -report FRESH/report.json
```

Exit1, stdout0bytes, stderr94bytes, SHA256
`f55bef7a8cb235bbac2fa946289eaaacfd440bb3b520db67ec66ddee3e488fd8`:

```text
wp-includes/class-wp-block-parser.php: PHP004 line 354: empty requires an array
exit status 1
```

The last line comes from go run. No output/report/staging artifacts; the empty
temporary directory was removed. In original add_inner_block, `$html` comes
from substr and `if (!empty($html))` tests that string. The existing empty
subset accepts arrays only. Earlier source progress is not proof of every
earlier method's semantic/runtime correctness. Full original Go generation,
build, execution,12-case/24-call comparison and G83-PARSER remain blocked.

## Bounded property implementation and verification

Defined unknown bare locals now admit named property assignment/actual-array
append, using one RHS then explicit nonnil PropertyObject and declared-slot
checks. Known wrong types/known absent class slots still reject statically.
Arrays retain value-copy separation and objects retain identity. No dynamic
properties/names, indexed/effectful receivers, direct concat, references or
general PHP exception equivalence. Fully implementing the structural host
protocol is explicit opt-in like existing nested writes, not generated-only
provenance or an authorization check.

Focused checks pass after documented legacy constructor/mapping spelling and
old static-rejection expectations were updated. Previously passing runtime/ABI
checks were reused; list/loop/switch now explicitly test generated rejection of
actual invalid scalar receivers. Exact commands/failure history are in the
[GoHP evidence](https://github.com/katsuakikureyama/GoHP/blob/main/docs/evidence/unknown-property-write-2026-10-29.md).
New owned fixture SHA256
`d66cd5f8ae71b8c84fffa2bc6cefc9d77eda77842a6b147253cc03237033aa76` covers
array-extracted objects, untyped helpers, aliases, falsy slots, detached nested
arrays, widened locals and same-slot RHS-reset append. Its exact native result
belongs to candidate-specific PHP8.3/8.4 automatic CI, not this original probe.

Final GoHP head `b8825cbb3a8eeca858da2e69ef9471b62c891b3b`, tree
`162d608c1663b6a4bc335c5d9ae60c7c2936b01b`;
[push40831935709](https://github.com/katsuakikureyama/GoHP/actions/runs/40831935709)
and [PR40831993242](https://github.com/katsuakikureyama/GoHP/actions/runs/40831993242)
attempt1 all4 target jobs SUCCESS. Saved push artifacts12878973325 (SHA256
`cb0aa1e359d8296386845aaa84453202b5ac5ec828b5ea5531ca98e714c8d01c`) and
12879450404 (SHA256 `f11189a3aab8d8f62233ac1c7888a18aee21aeee64fffb4ec9cb3a8d695c576c`)
read once: exact clean head/tree, Go1.23.12Linuxamd64/CGO0 and pinned
PHP8.3.36/PHP8.4.27. All8 new checks PASS on both. Owned fixture
parse/codegen/build/execute/compare PASS: native and Go stdout259bytes equal,
SHA256 `c5b4b9c74318477fb69889c1dde6df7833ca8b88539897d04219fd63e1bde230`,
exit0/quiet stderr0. General stderr/exception/side-effect equivalence stays
not-checked.812 total checks:8.3 794PASS+18outside /8.4 810PASS+2outside.
[PR result6490121401](https://github.com/katsuakikureyama/GoHP/pull/131#issuecomment-6490121401)
retains the evidence. Main5776a7b3a87725379f0445e2f737f7a560822927 has the same
tested feature tree; local main ff-only synchronized. Automatic main CI is
tracked separately, with no duplicate archive download or manual successful
suite/probe rerun. This is not the original parser's native comparison.

Existing130 closure/L1 native evidence remains unchanged. No old successful
native/consumer suite or enabled-PCRE run is repeated. Original files, existing
manifests/pins, auth/routes/deployment and CLI/WebMCP exclusion are unchanged;
Gossip CI is unconfigured and devboard unidentified/not synchronized. Next is the
measured empty-string value domain while remaining stack/snapshot/tokenizer/JSON
dependencies still require complete-unit proof.
