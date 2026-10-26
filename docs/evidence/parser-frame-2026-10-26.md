# Original parser Frame + Block /12-case execution record

PASS on2026-10-26JST: two complete unchanged original classes/constructors, three
owned source drivers, one generated binary and12 Go/PHP8.3.36/C/expected comparisons.
Both executions exit0 with empty stderr. This is not a full parser or render result.
[Contract](../parser-frame.md); [all12 typed observations and hashes](parser-frame-2026-10-26.json).

## Exact Gossip snapshot and execution

- Clean commit `3fe02fe966ebdb7f2193ca3a2d4ebb86e0dfd63c`, tree
  `d45bddedbac2201f79e99b0754bf1e9e31e9d71f`.
- Worktree `/private/tmp/gossip-parser-frame`, branch `codex/gohp-parser-frame`,
  base master804c6390be43488dddab8935ccbfa920a436b66c.
- From nested gossip/:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-constructor-null-defaults GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserFrame$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

Test PASS9.72s/package9.960s; vet PASS on the same snapshot. First committed
snapshot passed. No product/fixture/configuration change or successful rerun
after PASS; only documentation/evidence follows. Raw log:
`/private/tmp/gossip-parser-frame-output.log`. Timing is not a controlled benchmark.

## Actual observations

Original WP_Block_Parser_Block15–90 and WP_Block_Parser_Frame16–79 execute
without editing their PHP. Frame's actual constructor has five untyped
parameters, three required and two trailing null defaults. Four cases each use
actual source new calls with3/4/5 arguments, covering omitted versus explicit
null/zero/nonzero values, sum/override offsets and leading-null/zero/nonzero.
Original class/property/method/constructor shape is confirmed by PHP Reflection.

The owned source creates an original Block, puts it into an original Frame,
takes an array snapshot, then changes the local Block.innerHTML. All12 executions
observe the same nested actual Block identifier through Frame.block, snapshot
and returned Block, with changed Unicode/NUL HTML visible. No concrete Go field
injection or substitute class is involved. Snapshot is typed as array and its
five property keys are observed in declaration order; Frame's four scalar
properties, all Block fields and nine caller arguments are measured separately.
Caller inputs remain unchanged. Negative synthetic offsets test data arithmetic,
not the validity of parsing such a document. Block.innerBlocks is empty.

Both source hashes/ranges, corpus hash, exact compiler archive, fixed images,
generated source/report/binary and original/adapted bounded transports are
recorded in JSON. Key hashes:

| Artifact | SHA256 |
| --- | --- |
| Frame whole source | e5883d82dff099a4b1a23786e9a53b7e29aa73654db83e7ddc444ba17df5e105 |
| Frame complete class | 32c9c3d1d5dea106941e5e13a22216b7a9bef293fd315ced3b6bf207db577a24 |
| Corpus | 55ec7334da0d8de3898f741b80ef9942547b0e1977e3eb16867323c8e3e5654f |
| Generated Go | a3fc4c50e600d77745169a4526bd3d32793cf828ed441dcfaf5f2c9bc2c9fc49 |
| Build report | bd3e4af7417b006198914342b0fa10b4905ea9b25d85d6e2776c2cedf5af2c98 |
| Binary | d8cbd2ef4c8a6df66931c973650f1b1803ce1894739e456eda66fd4e9029a5a6 |

## Compiler prerequisite and evidence reuse

[GoHP PR76](https://github.com/katsuakikureyama/GoHP/pull/76) is merged to main
`d6f32d90c686da8da4116d81f00da36ce71506d5`, tree
`773422be588abc30b748a742c2162cc460e76e12`, identical to published clean feature
head `e4ac8e60bd4364acc7248005c030e6288fe739e1`. Base main718d0d3 retains #75's
explicit Execution/strtok work and #74's constructors. Only semantic admission/
min-max arity changes: existing NullDefaultArgs and NewContext/MContext dispatch
are reused. No new runtime API, ordinary method defaults or general cast support.

First clean local code commit9858318ca359e90bf880eb486b5bef1a16a32b2c,
treeffcd8654502a7876f66ff32626acbdbc3de6c75c, Go1.23.0darwinarm64,
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0:

- `go test -count=1 ./internal/semantic ./internal/codegen ./internal/compat`:
  PASS codegen0.518s/compat0.345s; semantic has no tests.
- `go vet ./internal/semantic ./internal/codegen ./tests`: PASS.
- `GOHP_PHP_TARGET=8.3 go test -count=1 ./tests -run '^(TestConstructorNullDefaultsRuntimeGuards|TestUnsupported|TestMalformedIRFailsWithoutPanic)$'`:
  PASS8.316s.
- `GOHP_PHP_TARGET=8.3 go test -count=1 -v ./tests -run '^TestDifferential/constructor_null_defaults$'`:
  fixedPHP8.3.36 wrapper PASS fixture0.83s/package1.675s. Observes null/zero/defaults,
  all-optional zero arguments, case-insensitive constructor, argument/body order,
  fresh objects and caller:constructor:after strtok continuation.

The native comparison first stopped because a reused local wrapper mounted the
previous worktree and could not find this new fixture. Go execution/comparison
was not-run then. Only the wrapper mount was corrected; unchanged fixture/core
comparison resumed successfully. No successful checks were manually repeated.
One later compiler commit changes a historical documentation sentence only.
Compiler native harness compares stdout/normal exits, not stderr/exceptions/other
side effects; Gossip's new comparison additionally requires both stderr empty.

[Feature CI40271821016](https://github.com/katsuakikureyama/GoHP/actions/runs/40271821016)
artifact verifies exact clean feature head/tree, PHP8.3.36/Go1.23.12linuxamd64 and
all543 checks/75 features PASS (five new checks). PR CI40271851717 also PASS.
[Main CI40271997777](https://github.com/katsuakikureyama/GoHP/actions/runs/40271997777)
artifact verifies exact clean main/tree, the same543/75 and fixed versions, all
PASS. Artifacts are `compatibility-php8.3-<run>-1`, retained locally under
`/private/tmp/gohp-constructor-null-defaults{-main,}-ci`. Evidence is reused;
no manual full-suite rerun.

## Limits

Historical16/24/334 configurations, pins, corpora, original source and evidence
are unchanged/not rerun; case counts are not added. Full parser/tokenizer/stack,
WP_Block_Type/registry/render/IDs/CSS/HTML/ThemeJSON/bootstrap/HTTP/DB/themes/plugins
and general PHP error/coercion/reference semantics remain unverified. No
production target configured/created; no new release/deployment.
Capability gossip.dev.wordpress-unit.translate-compare remains WebMCP-excluded
(local developer fixture); route/auth unchanged. Devboard unidentified/not synced.
Parent owns design/core/integration/review/measurement; workers authored independent
fixtures only. Before commit/measurement the parent corrected empty-array expected
encoding, snapshot array type and Go key-order/domain checks; no post-PASS test
change. Unrelated root edits and the separate compiler owner's root docs are untouched.
