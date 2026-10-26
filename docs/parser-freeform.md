# Original parser freeform method / selected-source unit

[Execution record](evidence/parser-freeform-2026-10-26.md): all12 pair cases
and vet PASS on the first committed snapshot, both exits0/stderr empty.

This configuration advances from original data classes into one actual original
parser method: WP_Block_Parser::freeform314–316. It creates an original
WP_Block_Parser_Block with null blockName, empty attrs/innerBlocks, unchanged
innerHTML and one innerContent string (including an empty string).

The source selection is explicit, not a complete parser translation:

| Side | Source executed |
| --- | --- |
| Go candidate | Complete original Block15–90; original Parser class/property prefix16–50, complete freeform314–316 and closing brace396; one owned composition driver |
| PHP oracle | Complete original class-wp-block-parser.php, which require_once-loads complete original Block and Frame files; the identical owned driver |

Every selected byte and all three full source files are hash-bound. No method
body is rewritten, no method is replaced with a stub and no substitute WordPress
class/service is authored. The Go class deliberately lacks the other six Parser
methods; full Parser adoption is not claimed. PHP Reflection verifies the full
original seven-method Parser, its four public properties, lack of constructor
and freeform's exact original signature/file lines, plus the original Block shape.
The oracle does not run a second extracted/reconstructed Parser.

The owned driver creates a Parser and invokes its unchanged freeform method
twice. Only that method creates the two Blocks. Go/PHP observers call the driver,
then measure actual object types/public getters or properties, without creating
blocks themselves, injecting fields, using host substitutes or calling parse().
The two returned Blocks must be different identifiers even with identical input.
All five Block properties, both input strings after execution and the Parser's
four still-null properties must match PHP and independently constructed expectations.

Twelve new cases cover empty strings, '0', plain text, HTML, literal block-comment
delimiters, CRLF, tabs, quotes/backslashes, NUL, Unicode, 4096-byte HTML and equal
HTML inputs. This method preserves data; it does not tokenize, validate, sanitize
or render HTML. One binary runs24 actual method calls/Block constructions across
12 pair cases; this is not24 independent corpus cases or an addition to earlier
Frame/data/helper results. Transport is direct bounded UTF-8 JSON strings, not a
general object/node codec. Actual null and empty array/list shapes remain distinct.

GoHP stays pinned to already verified merged #76 main
d6f32d90c686da8da4116d81f00da36ce71506d5 / tree773422be588abc30b748a742c2162cc460e76e12.
Its exact feature/main543-check CI evidence is reused; no compiler code/config
change or repeated compiler test. Parallel #77 standalone calls is merged
main0e736a366117a860b6e0cca7526477cfaca78230, but this unit assigns method return
values and needs none of that change. It is not evidence for adopting #77.
Historical12/16/24/334 configurations, pins/corpora/source/evidence are unchanged
and not rerun or summed.

From nested gossip/:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserFreeform$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

Fixed PHP8.3.36/C and Go images,
same-process oracle admission, readonly/nonroot/network-none containers and
resource/time limits remain unchanged.

## Remaining parser dependencies (source audit, not a full-translation run)

| Original method | Selected source | Dependencies still outside the pinned compiler boundary |
| --- | --- | --- |
| add_freeform |327–335| Ordinary method null defaults, bare return, property-left coalesce, substr and property-array append |
| add_inner_block |349–361| Class-typed/defaulted method args, property-rooted indexing/chained writes, substr and string empty |
| add_block_from_stack |371–395| Method defaults, array_pop/substr, string empty, property isset, compound/append writes |

These are read-only source/contract findings, not new observed compiler failures.
parse/proceed/switch/list/preg_match/json_decode and the overall stack/tokenizer
algorithm are not translated or executed here. Frame is loaded only by the native
parser's original require_once, not exercised by this corpus. WP_Block_Type/
registry/render/IDs/CSS/HTML/ThemeJSON/bootstrap/HTTP/DB/themes/plugins and general
casts/errors/coercions/references remain unverified. No production target configured
or created. Capability gossip.dev.wordpress-unit.translate-compare remains explicitly
WebMCP-excluded (local developer test); route/auth/server operations unchanged.
Devboard unidentified/not synced; PR/current-state tracking.
