# Original WP_Block_Parser_Block:16 typed comparisons PASS

The [contract](../parser-block-data.md) and [complete actual record](parser-block-data-2026-10-26.json)
cover the whole unchanged original class15–90, its actual five-argument
constructor, unchanged wp_should_add_elements_class_name/_wp_array_get and one
owned source glue function. One translation/build and one binary execute all16
cases freshly. Actual class/type, all five properties, declaration-ordered array
snapshot, decision bool (11 true/5 false) and all six argument states match
PHP8.3.36/C and expectations. Caller arguments and snapshot remain unchanged.

The helper constructs the actual original class, then returns that object and
its `(array)` snapshot. Neither observer fills fields directly; no replacement
class, resolver or service is used. The original parser uses the same creation/
cast shape at class-wp-block-parser.php149, but its algorithm/method is not run.
This is data-class adoption, not full parser/render/HTML/CSS support.

## Exact snapshot and commands

Clean Gossip test snapshot `ba74d40a0cd803fa3212d2a2768bb1cc1e5051b9`, tree
`151aa7a02a74a39497bf6f8ce1d3196698bd6f70`. From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-block-data-constructor GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserBlockData$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

16/16 PASS9.85s/package10.049s; vet PASS on the same snapshot. Both exits0,
empty stderr. First committed snapshot passed; no successful rerun or code change
after PASS. Raw `/private/tmp/gossip-parser-block-data-output.log` retained.

GoHP [#74](https://github.com/katsuakikureyama/GoHP/pull/74) merged main
`eb3b968d447ab7a11af4efd3b1a1976b2cd277e8`, identical tested tree
`ef8060b2e6a32b7da5784ebf13652092208bb059`, feature head
`126cfbad6dcdfd213049175615f1a48c8c53ee1a`. [Feature CI](https://github.com/katsuakikureyama/GoHP/actions/runs/40270789346)
artifact confirms exact clean head/tree, PHP8.3.36/Go1.23.12linuxamd64 and all528
checks/73 features PASS. [PR CI](https://github.com/katsuakikureyama/GoHP/actions/runs/40270794585)
and [main CI](https://github.com/katsuakikureyama/GoHP/actions/runs/40270977204)
also PASS. Parallel #73 while is preserved. Initial local6195438 snapshot passed
all selected checks; commands/times are in PR74 body/comments. After union-only
merge its successful local evidence was reused, not manually repeated.

## Provenance and limits

Three full original source hashes and complete class/function selections are
bound to WordPress54106f6. Class full-source hash
`7e21f483cf0c9deddd2144e026842bf906c72983908f7bd163c0716e8f6f325e`;
class fragment `4a3d9629bae4dd309bc1e3f16fc1ff8ca985a1778c2b15ac25dc1a9db13cf0ad`.
Native Reflection requires an original user-defined class with the five public
properties in order, exactly one method (constructor) and five required args,
plus both user-defined original functions. Expectations assert actual execution,
not stand in for it. Shared transports and original/adapted hashes are retained.

Generated SHA256 `d333aefc4e2c0843a49b8d3eadd82af2605a89b855f67e150bfbc1feee0a1205`;
report `facab6f20874eb5405a2e45163cd887c75045708d44567e0464867ef2c7a73fd`;
binary `4560d263fb5bf28de65ab8c6288d0034ec5716e734e308ab1b385dd13ea44fbd`.
JSON stores all16 complete actual records and remaining source/fragment/corpus/
driver/observer/compiler/archive/image hashes, commands/revisions and boundaries.

`innerBlocks` is empty or data-array entries; innerContent contains string/null
fragments. No nested parser objects or full parsing algorithm is claimed. Fixed
PHP8.3.36/C, builderGo1.25.14linuxarm64, nonroot/read-only/network-none/resource
limits and same-process admission remain. Codec bounds depth12/256 entries/
4096-byte strings/int±1e6; reached class-decision paths are arrays/null only.
Original WP_Block_Type/registry/render/unique IDs/CSS/HTML/Theme JSON/bootstrap/
HTTP/DB/themes/plugins and general casts/constructors/errors/coercions/references
remain unverified. Old24/334 configurations/pins/corpora/evidence unchanged/not
rerun; do not add their counts. No production target configured/created. Capability
stays WebMCP-excluded, route/auth unchanged; devboard unidentified/not synced.
