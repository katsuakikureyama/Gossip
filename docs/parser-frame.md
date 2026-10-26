# Original parser Frame containing an original Block

This isolated configuration translates two complete unchanged WordPress classes:
WP_Block_Parser_Block15–90 and WP_Block_Parser_Frame16–79. Both real constructors
and all ten declared public properties are included. It does not execute the
parser/tokenizer/stack algorithm or any render hook.

[Execution record](evidence/parser-frame-2026-10-26.md):12 comparisons and vet
PASS on the first committed snapshot; both exits0/stderr empty.

Three small owned source drivers create the real Block, create the real Frame
with three/four/five arguments, snapshot the Frame using (array), then change
the known local Block's innerHTML. Both observers invoke this source instead of
injecting generated fields. Frame.block, snapshot['block'] and returned Block
must retain the same object identifier and expose the changed HTML. The array
snapshot remains an array, not a second Frame object; object identity is measured
separately from scalar snapshots. No substitute class or WordPress service exists.

GoHP [PR76](https://github.com/katsuakikureyama/GoHP/pull/76), merged main
d6f32d90c686da8da4116d81f00da36ce71506d5, admits only trailing untyped null
constructor defaults, preserving prior fixed-constructor bounds and caller
Execution. Existing NullDefaultArgs pads omitted slots; explicit null/zero is not
rewritten. Ordinary method defaults, typed/non-null defaults, direct constructor
re-entry, inheritance, general casts/errors/coercions/references remain excluded.
The exact clean feature tree773422be588abc30b748a742c2162cc460e76e12 passed543
checks/75 catalog features on fixedPHP8.3.36/Go1.23.12linuxamd64; main has the same
tree. No tokenizer implementation is duplicated.

The12 new cases have four each for actual Frame arity3/4/5: omitted/null/zero/
explicit offsets, ignored supplied glue values for omitted slots, negative and
zero sums, leading-null/zero/nonzero, Unicode/NUL HTML, null/empty/nested attrs
and string/null innerContent fragments. Negative synthetic offsets exercise
data-class arithmetic, not validity of a parser document. Block.innerBlocks is
empty in this unit. Expected previous offsets are literal hand-specified values.
Actual scalar values/types, ordered snapshot keys, Block fields, nested object
identity and all nine caller arguments after execution must match PHP and
expectations. Caller input values stay unchanged despite the owned object write.

One translation/build produces one binary with2 original classes/constructors,
zero original functions and3 owned source drivers. Go checks the generated
original Frame/Block concrete types and public getters, without layout/field
writes. PHP Reflection verifies both user-defined classes, five declared public
properties in order and constructor-only methods; Frame has five untyped
parameters, three required and two null defaults, Block has five required.

Manifest binds both whole-source hashes/ranges and corpus. Existing bounded
node transport is adapted exactly once with original/adapted hashes recorded:
depth12,256 entries,4096-byte strings,int±1e6. Objects never enter that generic
codec: their class/type/public data and explicit identity are observed separately.
Fixed images, same-process PHP8.3.36/C admission, readonly/nonroot/network-none
containers and resource/time limits are inherited without widening.

From nested gossip/:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserFrame$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

Historical16/24/334
configurations, corpora, compiler pins, original source and evidence remain
unchanged/not rerun; counts are not summed. Full parser/tokenizer/stack,
WP_Block_Type/registry/render/IDs/CSS/HTML/ThemeJSON/bootstrap/HTTP/DB/themes/plugins
remain unverified. No production target is configured or created.

Capability `gossip.dev.wordpress-unit.translate-compare` remains explicitly
WebMCP-excluded: local developer fixture, no browser/server operation, route/auth
or deployment API change. Devboard unidentified/not synced; tracked in PR/state.
