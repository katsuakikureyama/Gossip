# Original parser-block data class → array snapshot → class decision

This configuration translates the complete unchanged `WP_Block_Parser_Block`
class (class-wp-block-parser-block.php15–90), including all five untyped public
properties and its actual five-required-argument constructor. It also translates
unchanged `wp_should_add_elements_class_name` and its `_wp_array_get` dependency.
There is no replacement WordPress class or service implementation.

One small owned source helper constructs the original class and returns both
that actual object and `(array)$block`. The original parser uses this same
construction/cast shape at class-wp-block-parser.php149 before parsed block data
reaches render filters. Only the data class and owned glue are executed here,
not that parser method, tokenizer/stack algorithm, HTML parsing or render hook.

GoHP #74 main `eb3b968d447ab7a11af4efd3b1a1976b2cd277e8` provides fixed
untyped public constructors, implicit-null public
properties and a known-generated-class public-property array snapshot. Exact
arity, once-only constructor effects, declaration order/null retention and
array-copy/object-identifier boundaries are explicit; unknown/foreign/scalar/
array/Closure casts, typed/defaulted constructors, direct re-entry, other magic
methods, inheritance and visibility/dynamic-property semantics remain excluded.
Its exact clean integrated feature tree passed528 CI checks/73 catalog entries
on PHP8.3.36, retaining parallel #73 while. Older compiler pins are unchanged.

The16 cases measure actual object class/properties, generated array snapshot,
real class-decision bool and all six caller arguments after execution. Five
property names/order and values must match constructor inputs. Corpus variation
includes null/empty attributes, button colors/gradient, link/hover, heading/h6,
skip options, nonnull false/null style leaves, Unicode/NUL HTML, ordered mixed
keys and string/null fragments. `innerBlocks` is empty or data-array entries in
this bounded unit, not nested parser objects. It does not validate block comment
parsing or CSS validity. There is one source construction per case.

One translation/build produces one binary with one complete original class/
constructor, two original functions and one owned function. Both observers call
the generated/original source helper; neither populates fields directly. Go
observes the real generated original type and getter values; native Reflection
requires a user-defined original class with five public properties in original
order, only its constructor and exactly five required parameters, plus the two
user-defined original helpers. Expectations never replace actual execution.

Three whole-source hashes, complete fragment selections and corpus hash bind
the inputs. Shared recursive transport retains depth12/256 entries/4096-byte
strings/int±1e6 bounds, with reached class-decision paths restricted to array/
null intermediates. Earlier24/334 configurations/pins/corpora/evidence remain
unchanged and are not rerun or added to the new single-binary count.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserBlockData$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

Record executed results/snapshots before claiming PASS. Fixed PHP8.3.36/C and Go
images, nonroot/read-only/network-none/resource limits and same-process admission
remain active. Original full parser/WP_Block_Type/registry, unique IDs, render/
CSS/HTML/Theme JSON/bootstrap/HTTP/DB/themes/plugins and general error/coercion/
reference semantics remain unverified. No production target configured/created.
Capability `gossip.dev.wordpress-unit.translate-compare` stays WebMCP-excluded:
local developer fixture, no browser/server operation, route/auth or deployment
API change. Devboard unidentified/not synced; tracked in PR/current-state.
