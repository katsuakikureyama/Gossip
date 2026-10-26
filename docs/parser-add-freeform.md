# Original add_freeform / selected parser output path

This unit executes unchanged WP_Block_Parser::add_freeform327–335 and its
unchanged freeform314–316 dependency, plus complete original Block15–90.
Go selects the original Parser class/property prefix16–50 and closing brace396.
Other five Parser methods are omitted, not stubbed/replaced. PHP loads the full
original Parser/Block/Frame files, not a reconstructed Parser; exact whole-file
and fragment hashes are bound by the manifest. Native Reflection verifies the
full seven-method class and both exact method line/signature/default shapes.

One owned source driver creates a Parser, sets document/offset/output/stack
through translated source, invokes the real add_freeform with an omitted or
explicit argument and returns the Parser plus raw return. It does not replace
parser logic. Neither observer initializes generated fields or creates Blocks.
The original method computes omitted/null length, skips exactly zero, otherwise
slices bytes, invokes real freeform, casts the new original Block and appends it.
Offset/document/stack are not advanced by this method.

Sixteen cases distinguish no append from an appended empty Block, omitted/null/
zero/positive/negative lengths, signed/outside/end offsets, existing seed output,
NUL and byte-aligned/split UTF-8. Explicit expected HTML is a literal canonical
base64 value (null means no append), not a reimplementation of substr in the
observer. All ordered output keys and Block property keys/types/values, null
return, Parser document/offset/stack and all caller inputs after execution must
match Go, full-original PHP8.3.36/C and expectations. Binary-safe base64 is used
only by observation transport, so splitting a UTF-8 character preserves its raw
byte instead of silently replacing it. Inputs are bounded UTF-8 strings4096bytes,
offset/length±1e6, seed list<=4strings; these are synthetic data-path cases, not
claims about valid full-parser offsets/documents.

The compiler pin is [GoHP PR81](https://github.com/katsuakikureyama/GoHP/pull/81)
merged main565e65e6c9d8e211b06779739bfba76b07c4d92d,
tree3b509a251cb2ab3d037c089c1ab0295a1e0a0d60. It adds actual-int64 binary subtraction
at addition precedence, with guarded overflow and no coercion/general unary
negation. Exact clean feature/PR CI passed569 checks/80 features on fixed versions.
It retains [PR80](https://github.com/katsuakikureyama/GoHP/pull/80) merged main
4df47e00cae2dfd25e4cc5bde299ef51cc56c4c5, tree83cc11e72f83675b5a6ab606bbd6782b8bfd2d57,
provides bounded public method null defaults, bare returns, known declared-array
property append, sole-return-new method factory knowledge for guarded snapshots,
and actual string/int/int-or-null byte substr. Caller Execution/copies and
exact scalar required constructors from parallel #79 are retained. Feature/
PR CI passed564 checks/79 features on exact clean head/tree, fixedPHP8.3.36/
Go1.23.12linuxamd64. No generic cast, autovivification, nested property mutation,
weak coercion or PHP error equivalence is claimed.

The initial clean consumer snapshot538abb9b16dc1ebfdc93750ea9958c20ba3321b4 on
the #80 pin failed during generation: PHP002 line117 expected semicolon at minus.
Build/run/oracle comparison were not-run then. The original's binary subtraction
was a newly measured dependency, not a PHP output mismatch. #81 corrects that
compiler gap; only the consumer pin/documentation changes before resuming.
The initial log remains /private/tmp/gossip-parser-add-freeform-output.log.
The initial snapshot's vet PASS applies to unchanged Go harness code; it is not
rerun solely for the JSON pin change. No prior successful comparison is repeated.

Correction to the earlier source audit: the coalesce left operand at328 is the
local length variable, not a property. Existing coalesce already covers it;
no new property-left probe was necessary or added.

From nested gossip/:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserAddFreeform$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

Executed results must be recorded before PASS. Existing fixed images, nonroot/
readonly/network-none/resource limits and same-process oracle admission remain.
Historical freeform/Frame/data/helper configurations/pins/corpora/source/evidence
are unchanged/not rerun or summed. parse/proceed/tokenizer/stack/other five Parser
methods, registry/render/CSS/HTML/ThemeJSON/bootstrap/HTTP/DB/themes/plugins and
general PHP errors/coercions/references remain unverified. Frame is native-loaded
only. No production target configured/created. Capability
gossip.dev.wordpress-unit.translate-compare remains WebMCP-excluded (local fixture);
route/auth/server operations unchanged. Devboard unidentified/not synced; PR/state.
