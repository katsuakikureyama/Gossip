# Original freeform method /12 pair-case execution record

PASS on2026-10-26JST:12 Go/PHP8.3.36/C/expected comparisons, with24 actual
original freeform calls and24 original Block constructions in one binary.
Both executions exit0 with empty stderr. [Contract](../parser-freeform.md);
[all12 typed observations and exact hashes](parser-freeform-2026-10-26.json).

## Snapshot and commands

Clean commit `d218319afdf7327cd5001f1cd9ac0cd085e802b6`, tree
`0ece4798b033d6bbb56ba665746166e116e7df00`, worktree
`/private/tmp/gossip-parser-freeform`, branch `codex/gohp-parser-freeform`,
base master72e9ad56045bfcf3d03923e9c751cf14ef29ad9c.

From nested gossip/:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-constructor-null-defaults GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserFreeform$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

Test PASS12.29s/package12.498s; vet PASS same snapshot. The first committed
snapshot passed, without later test/config changes or successful reruns.
Only documentation/evidence follows. Raw log:
`/private/tmp/gossip-parser-freeform-output.log`. Timing is not a benchmark.

## Source selection versus oracle

The Go candidate contains the complete unchanged original Block15–90, original
Parser prefix16–50 (class and four properties), complete original
freeform314–316 and closing brace396. It omits the other six Parser methods.
No body is rewritten or replaced with a stub; this is a selected-source unit,
not translation of the complete Parser.

The PHP oracle loads the full original class-wp-block-parser.php, including its
actual require_once of complete Block and Frame files. PHP does not use a
reconstructed/extracted Parser. Native Reflection confirms all seven original
methods, original property order, no constructor, exact freeform signature and
original line bounds. Frame is loaded only, not exercised. All full files and
candidate fragments are independently hash-bound in the manifest/record.

The identical small owned composition driver creates one Parser and calls its
actual freeform method twice. Observers invoke only that driver, not host-side
Block constructors or field injection. Both returned values must be actual
original Block objects with null blockName, empty arrays, unchanged innerHTML
and single-string innerContent. Observed getter/property values are emitted;
expected records are not substituted for execution.

Twelve input pairs cover empty/'0'/plain text/HTML, literal block comments,
CRLF/tab, quotes/backslashes, NUL/Unicode,4096-byte HTML and identical HTML.
In every pair the Blocks are different object identifiers, the Parser's four
properties remain null and both caller strings are unchanged. Empty input still
produces a Block with innerContent=[''], not an omitted block. Literal block
comments are preserved text: this helper does not parse them.

## Compiler evidence reused

GoHP pin remains [PR76](https://github.com/katsuakikureyama/GoHP/pull/76) merged
main d6f32d90c686da8da4116d81f00da36ce71506d5, tree
773422be588abc30b748a742c2162cc460e76e12. No compiler edits, push/PR, or compiler
test rerun this turn. Prior downloaded [feature CI40271821016](https://github.com/katsuakikureyama/GoHP/actions/runs/40271821016)
and [main CI40271997777](https://github.com/katsuakikureyama/GoHP/actions/runs/40271997777)
artifacts verify exact clean commits/tree,543 checks/75 features all PASS,
PHP8.3.36/Go1.23.12linuxamd64. Their retained evidence is linked from
[the previous Frame record](parser-frame-2026-10-26.json).

Parallel #77 standalone call statements is already merged main0e736a3, but the
candidate assigns both method results and does not need/adopt that newer pin.
No new tokenizer/runtime semantics are inferred from this unit.

Key SHA256:

| Artifact | SHA256 |
| --- | --- |
| Complete parser file (native oracle) |8563621090de5e307431b1d29b78a9f75fd9ee238a0137b563a7db7b2e1f8d3b|
| Original freeform314–316 |b54693c7de89a75c8dfa4139ecbc512533eba115be8f85ca5bcd5d438023830c|
| Corpus |23bf6bd18581421a6161398d0a83fbfd56694982e9d55f7a2b2af04c4f10163e|
| Generated Go |9f140a2e56e790f9ef0063620fa9c0af1e98ae775c1ed2814c82d6575e3407f2|
| Build report |9b53bd29520d2470119df7d7c416ab2a499de5ce5572c8663041d2ae48fab570|
| Binary |42a886372c73aab8a69d193e825ff780f98605f40cc9e016700e1b1cb85a46ff|

The JSON also records exact archive/compiler/driver/observer/image hashes and
all actual typed records. Fixed consumer Go image is1.25.14linuxarm64; all
containers remain readonly/nonroot/network-none with existing resource/time
limits and same-process PHP8.3.36/C admission.

## Not verified

The whole parser, other six Parser methods, parse/proceed/tokenizer/stack,
WP_Block_Type/registry/render/IDs/CSS/HTML/ThemeJSON/bootstrap/HTTP/DB/themes/plugins
and generic casts/errors/coercions/references remain unverified. The tail-method
dependency table is a source/contract audit, not a recorded compiler failure.
Historical12/16/24/334 configurations/corpora/pins/source/evidence are unchanged,
not rerun or summed. There is no production target configured/created.
Capability gossip.dev.wordpress-unit.translate-compare remains WebMCP-excluded
(local test), route/auth unchanged. Devboard unidentified/not synced.
Parent owns selection/design/harness/integration/review/measurement; worker
authored independent corpus/observers, explorer did read-only tail investigation.
