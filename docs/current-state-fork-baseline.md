# Current state — WordPress fork baseline

Last updated: 2026-10-29 (JST).

## Latest: regex property inputs progressed; named class parameter blocks

GoHP [PR123](https://github.com/katsuakikureyama/GoHP/pull/123) admits named
property subject/offset values in the existing PHP8.3-only experimental regex
profile. Source-ordered once-only snapshots precede actual string/int guards;
direct-null local output, straight-line scope and legacy storage remain.
The unchanged full parser passes252/255 and now stops at named class parameter349.
[Probe](evidence/parser-preg-property-probe-2026-10-29.md) records the clean
compiler/source identity, exit1 and next dependencies. Complete semantic
validation, generated full-parser execution and application acceptance remain
blocked; P2 stack/object contracts and P3–P5 are not complete.

At clean code8403235, focused compiler/CLI/generated guards, catalog and vet
PASS once. Enabled DarwinARM64 PCRE2/static-library comparison against pinned
PHP8.3.36 PASS:15 ordered typed observations, including named property inputs,
mutation, byte offsets, capture copies and shared error-state reset. This is
structural typed equality, not serialized JSON byte equality. Initial Go
launcher/checksum setup failure occurred before execution and is retained.
Final headda57151 only adds evidence/docs; executable results are reused.
Default Linux CGO0 CI separately checks regressions, not enabled PCRE or PHP8.4
regex. Original files, consumer pins and old native reference results remain
unchanged. Gossip Actions remain unconfigured. No route/auth/deployment change;
local inspect/WebMCP exclusion stays. Issue1 tracks progress; devboard remains
unidentified/not synchronized.

## Previous: nested-property admission progressed; regex property input blocks

GoHP [PR122](https://github.com/katsuakikureyama/GoHP/pull/122), merged into
maina30709e51ecfe0231451afa121f19ab1ed247e97, adds guarded
named nested-property assignment, string concat and array append. The final
owner is resolved after RHS evaluation, retaining object identity and value
copies. Complete original source passes207–208 and stops at `preg_match`'s
property input at252. [Probe](evidence/parser-nested-property-probe-2026-10-29.md)
records clean commit2421daf, the exact source hash and exit1. Full source
semantic validation/generation/execution remains pending. P2 stack mutation,
unknown single-property writes/class parameters and P3–P5 remain open.

Focused local compiler/runtime/generated guards, existing public-property
checks, catalog and vet PASS once. Native PHP comparison is recorded in PR122;
an initial comparison mismatch fixed the codegen order to RHS-first for all
three operations. Only changed checks reran locally. Final headf01730b push
40731756226 / PR40731765069 CI both PASS: native8.3.36/8.4.27 independently
match385bytes of stdout, exit0 and quiet stderr. Seven new checks PASS;
765 total:8.3 747PASS+18 inapplicable /8.4 763PASS+2 inapplicable.
This is a generic language contract, not full-parser acceptance. Old native references and
consumer pins are retained without reruns. Gossip CI is unconfigured; no
route/auth/deployment or local inspect/WebMCP-exclusion change. Issue1 tracks
progress; devboard is unidentified/not synchronized.

## Previous: switch admission progressed; chained-property write blocks

GoHP [PR121](https://github.com/katsuakikureyama/GoHP/pull/121) adds nonnumeric
string switches with fallthrough, first duplicate dispatch, default and nearest
break. The complete unchanged parser advances past switch100 to the chained
property write at207. [Exact probe](evidence/parser-switch-probe-2026-10-29.md)
records exit1, before complete semantic validation or generated execution.
Next P2 nested property/stack mutation; P3–P5 remain open. Original source,
references and fixture pins are unchanged. Documentation-only checks here;
GoHP CI separately measures its generic 8.3/8.4 contract. Gossip CI is unconfigured.
Local inspect/WebMCP exclusion remains; no route/auth/deployment change.

## Previous: P1 list admission progressed; switch still blocks

GoHP [PR120](https://github.com/katsuakikureyama/GoHP/pull/120) adds bounded
positional list assignment to bare locals. The complete unchanged parser now
reaches switch at line100 instead of stopping at list assignment at94.
[Probe evidence](evidence/parser-list-probe-2026-10-29.md) binds the source,
compiler and exit1; it does not claim generated full-parser execution.
Next P1 switch, then P2–P5. Existing references, fixtures and pins remain unchanged.
Documentation checks only here; successful native comparisons were not rerun.
GoHP CI separately tracks the generic PHP8.3/8.4 language contract. Gossip CI
remains unconfigured. Issue tracking updated; devboard unidentified. No change
to existing local inspect capability/WebMCP exclusion or route/auth/deployment.

## Previous: P1 relational admission progressed; list assignment still blocks

GoHP [PR119](https://github.com/katsuakikureyama/GoHP/pull/119) clean source
1d7cd3e9bab88b16a7c9a26fdce40457b92c2421 adds bounded int64/null `<` and `>`.
The same complete original parser now rejects `list(...) = ...` at line94
instead of the lexer rejecting `>` at98. [Probe evidence](evidence/parser-relational-probe-2026-10-27.md)
records the exact source, command and exit1. This is progress in admission only:
no generated full-parser execution or application acceptance. Next is P1 list,
then switch; P2–P5 remain open. Original bytes, all fixtures and pins are unchanged.
Documentation-only changes require no repeat of native references or successful
consumer tests. GoHP CI tracks its generic PHP8.3/8.4 contract separately.
Gossip CI remains unconfigured. Issue1 tracks progress; devboard unidentified.
Existing developer-local capability/WebMCP exclusion; no route/auth/deployment.

## Previous: PHP 8.3 application-first acceptance

Owner direction, 2026-10-27: [G83-PARSER](php83-application-acceptance.md) is the
next delivery gate: all seven original parser methods, complete Block/Frame and
the original loading edges. Then actual bootstrap/hooks, DB, HTTP and a bounded
single-site workflow. Fix demonstrated PHP semantics before unrelated helpers,
generic COW expansion, benchmarks or new PHP targets; retain existing 8.4 CI.

WordPress54106f6 / PHP8.3.36 digest and original bytes remain fixed. Full unchanged
parser admission on GoHP9a7a9ce is **blocked**: `PHP002 line 98: unsupported token '>'`,
exit1, no generated source/report. Other inspected blockers: list/switch,
class-typed/nested property locations, stack mutation, preg property inputs,
JSON and require_once loading. No generated full-parser result is claimed.
The new native full-parse fixture/corpus is separate from old token/regex results.
At committed test inputs ed596cb94ae38631bd78744a92bb9adc260f4bc4, **12 cases / 24
original parse calls recorded** once (test1.05s/package1.289s), native8.3.36
Linux/aarch64/PCRE10.42, exit0/stderr0; fresh/reused output/state/error codes and
caller preservation agree. Malformed JSON gives error4/null attrs. `go vet`
PASS. This is native-reference-recorded, not generated-Go or application PASS.
See [baseline evidence](evidence/parser-full-baseline-2026-10-27.md) for execution
status; earlier successful configurations are not rerun or silently repinned.

Later DB/theme/server/platform configurations are not yet frozen or implemented.
Next implementation: P1 comparison/destructuring/control syntax at actual parser
callsites, with bounded PHP-native contracts and whole-file re-probes after
changes. UI/WebMCP/MCP excluded for the local fixture; no route/auth/deployment.
Issue1 tracks the gates; external devboard unidentified/not synchronized. Gossip
connector writes returned403; the existing authorized GitHub CLI path is usable
without changing credentials or permissions.

## Previous: generated regex consumer of GoHP #114

[Consumer contract](parser-preg83-generated.md), [evidence](evidence/parser-preg83-generated-2026-10-27.md).
Clean Gossip code99cb3f86f6c5ad5e55836b7ca7511d2d539a12ab: actual GoHP generation,
enabled native build and **23/23 retained PHP comparisons PASS** (6.41s). GoHP
compiler/runtime pin d7650944802a33cfc2dab0006bb410374d828245 (#114), exported
from its immutable Git archive; previous Gossip pins unchanged. The existing
24-case native record is reused; backtrack_limit_one is explicitly excluded
because generated B3 fixes limit100000, not1. No false 24/24 claim.

Original full-source hashes and line251 pattern bytes are verified; an owned
typed PHP driver re-encodes the literal without changing runtime regex bytes.
No original WordPress file, GoHP code or general module dependency is modified.
This is not next_token translation. Observations preserve ordered captures,
types/raw bytes/byte positions/presence/output-before/after/error codes.
Darwin ARM64 Go1.23.0/clang17/SDK15.5/PCRE2static10.47 only, explicit opt-in;
candidate exit0/stderr0, linker duplicate-archive warning retained. No dependency
installation or PHP rerun. Contract/source-baseline and vet PASS once, original
WordPress tree diff empty. Complete actual/expected JSON and hashes retained.

Next: original next_token dependency closure (property inputs, destructuring,
JSON and remaining source constructs), without rewriting the original method.
Full parser/general references/PHP8.4 regex/enabled Linux/packaging/performance
remain unmeasured. UI/WebMCP/MCP excluded; no route/auth/release/deployment.
Progress uses GoHP Issue1; devboard unidentified/not synced.

## Previous: staged PCRE2 bridge/output implementation contract

[Contract](pcre2-bridge-contract.md), design-only follow-up to #48. Read-only
GoHP audit pinned to fd6936a (#107): value-only call IR, no output destination,
scalar-only Frames/elements versus detached nested runtime.Array, Execution with
strtok but no regex error state, and default CGO=0 builds. Independently owned
#108 Slot-array work was in flight; no changes or assumptions about its result.

Next bounded implementation B1 is an optional engine-only Go/C bridge; keep PHP
builtin admission closed. B2 adds fresh PHP-shaped output/Execution state only
after its own evidence; B3 requires explicit local-output IR, a reviewed writable
local/nested-value contract and generated-source comparisons. No duplicate cell
engine, WordPress-specific compiler shortcut, forced CGO default or silent fallback.
Broader aliases/locations/flags/versions/JSON/full parser and distribution remain
separate gates. No new native run, code/corpus/pin change or compatibility count.
Documentation/source-link/original-preservation checks are the relevant validation;
reuse existing PASS rather than rerun it. No new capability/route/auth/deployment;
UI/WebMCP/MCP excluded. Progress uses Issue1; devboard unidentified/not synced.

## Previous: regexp2 comparison and explicit incompatibility boundary

The [isolated regexp2 profile](parser-regexp2-comparison.md) yields 22 equal
outputs, 1 difference (UTF-8-interior starting offset rejected) and 1 non-comparable
PCRE match-count limit. Raw captures retain invalid UTF-8 bytes for this corpus;
that does not establish general byte semantics. No pattern rewrite or silent
boundary rounding. Clean code `073e1218292c160a02f6e06ad2a4a83f2ea547e7`;
observation test PASS6.91s/package7.147s and vet PASS once. [Evidence](evidence/parser-regexp2-comparison-2026-10-27.md).

Official regexp2/v2 commit a6ef7b6 requires Go1.25; a temporary module uses cached
Go1.25.14 without changing Gossip's module floor or any GoHP pin. Saved PHP and
PCRE2 results remain unchanged/not rerun. Test PASS means observations collected,
not compatibility. Do not adopt this direct string-API profile as a drop-in backend.
PCRE2 remains the better measured baseline for these cases; next scope a GoHP
bridge/output-argument contract plus build/platform/distribution policy. No runtime
adoption or full parser/PHP8.4/JSON/performance result. UI/WebMCP/MCP excluded;
no route/auth/deployment. Progress uses GoHP Issue1; devboard unidentified/not synced.

## Previous: independent PCRE2 / retained PHP comparison

[PCRE2 10.47 C-engine comparison](parser-pcre2-comparison.md) matches all 24
retained PHP8.3.36/PCRE10.42 direct-regex observations, including ordered typed
captures/raw bytes/byte positions/presence and the backtracking-limit failure.
Clean code commit `c6d82db2faaad12592e3274d677c27c379dea9ac`; initial sandbox
attempt failed at the macOS compiler launcher before comparisons; same-code
retry with compiler access PASS (0.73s/package0.916s). Vet PASS; successful native
reference and candidate comparisons are not rerun. [Evidence](evidence/parser-pcre2-comparison-2026-10-27.md).

Only existing pinned local header/archive are linked in a temporary host probe;
no install/vendor/module dependency, GoHP/cgo integration or backend adoption.
Original WordPress and old pins remain unchanged. Candidate and PHP PCRE versions
differ explicitly; this is not cross-version general equivalence or full parser
execution. Next: regexp2 comparison and runtime/output-argument/build-policy gates.
PHP8.4/JIT-on/out-of-range offsets/other failure profiles remain unmeasured.
Local fixture is WebMCP/UI excluded; no route/auth/deployment. Progress uses GoHP
Issue1; devboard unidentified/not synced.

## Previous: original next_token / preg_match native reference

The [24-case native reference](parser-token-reference.md) passes its predeclared
PHP expectations at clean commit `29e6d2637e68cbc3a8d5764ba825e6e6600ddf74`:
17 matches, 6 non-matches, 1 backtracking-limit failure. PHP 8.3.36 / PCRE 10.42,
JIT explicitly off; full original Parser/Block/Frame loaded unchanged. Separate
direct preg_match and original next_token calls retain capture bytes/order/offsets,
error types and parser/caller state. Invalid UTF-8 attributes match the regex but
fail JSON decoding. [Execution record](evidence/parser-token-reference-2026-10-27.md)
includes complete machine-readable observations and hashes. Targeted native test
and vet PASS once; documentation-only additions reuse that code/corpus evidence.

This is not generated Go, a backend comparison, PHP 8.4, or full parse evidence.
GoHP code, prior compiler pins and original WordPress bytes are unchanged. GoHP
#105 source-token inventory is separate work, not duplicated here. Next: evaluate
the exact pattern/captures against PCRE2 first, then regexp2; out-of-range offsets,
JIT-on and other failure modes still need separate evidence. No backend adopted,
production target, route/auth or user operation added; WebMCP excluded. Progress
uses GoHP Issue1; devboard remains unidentified/not synced.

## Previous: project overview and responsibility map

Documentation-only follow-up to PR44, based on master
`ae218d3f78cc649e4f2884ceb24870e0fe2d76b7`, branch `docs/project-boundaries`.
The README is now an entry point with separate evidence rows, not a running
history. Prior README and reuse-boundary wording are retained under `history/`
with explicitly historical status and rebased links. Existing evidence JSON/MD,
original WordPress files, fixtures, Go code and compiler/oracle pins are unchanged.

The [current responsibility map](reuse-boundary.md) distinguishes GoHP's general
PHP semantics from Gossip's original-source acceptance and native integrations.
Original parser/WP_Hook translation remains the target; the native SDK is not a
substitute proof. DB/HTTP adapters are candidates, not chosen implementations;
include/autoload/themes are unverified, not implicitly discarded. Standing owner
authorization covers normal verified PR merges; this is not deployment approval
for an unspecified destination or a license decision.

[Regex backend evaluation gates](parser-backend-assessment.md) pin the actual
original next_token requirement and separate pattern syntax from byte offsets,
matches output, absent captures, errors/resource limits, JSON and parser-state
dependencies. PCRE2-first evaluation and regexp2 comparison are planned, not
implemented/measured/adopted. GoHP #103/#104 source-byte verification is not PHP
or WordPress application acceptance; it does not repin these Gossip fixtures.

Public evidence uses path placeholders instead of personal absolute paths. The
existing historical records remain unchanged; prior-path cleanup is separate.
No PHP behavior or compatibility totals are added by this documentation change.
No new capability/route/auth/SDK/server/deployment changes; WebMCP remains excluded.
Validation and publication results belong in this change's PR; do not rerun
successful historical native comparisons for a documentation-only update.
Devboard remains unidentified/not synced; cross-project progress uses GoHP Issue1.
Next: a pinned preg-match/next_token reference corpus and backend assessment,
alongside independently bounded original parser/storage dependencies.

## Previous: original add_freeform output path /16-case PASS

Unchanged add_freeform327–335/freeform314–316 plus complete Block15–90 execute
in one binary. Go selects original Parser prefix16–50/closing396; other five
methods are omitted, not stubbed. PHP loads complete original Parser/Block/Frame
files and verifies exact full class/line/default shapes. Owned source initializes
actual Parser state and invokes add_freeform; no observer injection/host Block.
All16 cases match PHP8.3.36/C/expectations:13 appended Block arrays,3 skipped,
null returns, ordered seed/output/properties, unchanged document/offset/stack/
caller inputs. NUL and split UTF-8 bytes remain exact via base64 observation.
[Contract](parser-add-freeform.md), [record](evidence/parser-add-freeform-2026-10-26.md).

Success43ebeba5710e4e4859109401b85ae48783839c06/tree5f17710e413b4a8a654a13b124c634b5d31e5bd3:
PASS10.48s/package10.690s, both exits0/stderr empty. Initial538abb9 on #80 pin
failed during generation at binary minus (no candidate build/run/oracle compare).
Original source/corpus/harness/observers unchanged; only pin/doc updated after
compiler #81. Initial vet PASS is reused for unchanged Go code, not rerun after
JSON pin change. Initial/success logs are separate; no successful comparison rerun.

GoHP #80 mergedmain4df47e0/tree83cc11e72f83675b5a6ab606bbd6782b8bfd2d57 adds
bounded method defaults/bare return/property append/factory snapshots/substr,
retaining parallel #79 exact scalar constructors. Feature/PR/main CI564checks/
79features PASS. #81 finalpin565e65e6c9d8e211b06779739bfba76b07c4d92d/
tree3b509a251cb2ab3d037c089c1ab0295a1e0a0d60 adds actual-int64 subtraction,
overflow rejection and intact decrement-token rejection; feature/PR/main CI569/
80 PASS with exact clean identities/fixedPHP8.3.36/Go1.23.12linuxamd64 artifacts.
Both initial local compiler negatives had one obsolete expectation; test-only
updates resumed only affected case, with no product-code fix or successful rerun.

The earlier property-left-coalesce audit was incorrect: the left operand is local
length, already supported. New dependency binary subtraction was found by actual
generation. Historical configurations/pins/corpora/evidence are unchanged/not
rerun or summed. Other five methods/parse/proceed/tokenizer/stack/render/CSS/HTML/
bootstrap/HTTP/DB/themes/plugins and generic PHP semantics remain unverified.
No production target configured/created; WebMCP excluded, route/auth unchanged.
Devboard unidentified/not synced. Next inspect actual stack/output dependencies.

## Previous: selected original freeform method /12 pair-case PASS

The unchanged original WP_Block_Parser::freeform314–316 now executes in Go with
its complete original Block dependency. Go selects the original Parser class/
four-property prefix16–50, complete method and closing brace396; other six
methods are omitted, not stubbed. The PHP oracle loads the complete original
Parser file and its real require_once Block/Frame files, not a reconstructed
Parser. Native Reflection confirms its seven-method shape and exact freeform
signature/lines. No complete Parser translation or parse() execution is claimed.
[Contract](parser-freeform.md), [record](evidence/parser-freeform-2026-10-26.md).

One owned source driver creates a Parser and calls freeform twice. All12 input
pairs/24 actual calls and Block constructions match PHP8.3.36/C/expectations:
five Block properties, distinct object identifiers even for identical HTML,
four still-null Parser properties and two unchanged caller strings. Empty/'0'/
HTML/literal block comments/CRLF/tab/quotes/backslash/NUL/Unicode/4096 bytes are
preserved. Observers neither construct Blocks nor inject generated fields.

Clean snapshotd218319afdf7327cd5001f1cd9ac0cd085e802b6, tree
0ece4798b033d6bbb56ba665746166e116e7df00: PASS12.29s/package12.498s, vet
PASS same snapshot, both exits0/stderr empty. First committed snapshot passed;
only docs/evidence follows, no successful rerun or code/config changes.

GoHP remains pinned to #76 merged maind6f32d90c686da8da4116d81f00da36ce71506d5,
tree773422be588abc30b748a742c2162cc460e76e12, reusing exact clean feature/main
543-check/75-feature CI evidence without compiler edits/tests/publish. Parallel
#77 standalone calls is merged main0e736a3 but is not adopted by this pin.
Tail-method source audit finds add_freeform327–335 the next smaller candidate,
still needing method defaults/bare return/property coalesce/substr/append.
This audit is not a new full-parser translation attempt or observed rejection.

Historical Frame12/data16/24/334 configurations/corpora/pins/evidence are unchanged,
not rerun or summed. Other Parser methods/tokenizer/stack, WP_Block_Type/registry/
render/IDs/CSS/HTML/ThemeJSON/bootstrap/HTTP/DB/themes/plugins and general error/
coercion/reference/cast semantics remain unverified. Frame is native-loaded only
in this configuration. WebMCP excluded, route/auth unchanged; no production
target configured/created. Devboard unidentified/not synced; PR/state tracking.

## Previous: original parser Frame + Block /12-case PASS

Two complete unchanged classes/constructors (Block15–90, Frame16–79) now run in
one generated binary. Three owned source drivers use actual Frame3/4/5-argument
new calls, then snapshot Frame to array and change its original nested Block.
All12 Go/PHP8.3.36/C/expected records match: actual original types, default/sum/
override offsets, leading-null/zero/value, five snapshot keys in declaration
order, nested object identity/changed HTML, Block fields and nine unchanged
caller arguments. No substitute class or observer field injection.
[Contract](parser-frame.md), [record](evidence/parser-frame-2026-10-26.md).

Clean snapshot3fe02fe966ebdb7f2193ca3a2d4ebb86e0dfd63c, tree
d45bddedbac2201f79e99b0754bf1e9e31e9d71f: PASS9.72s/package9.960s, vet
PASS same snapshot, both exits0/stderr empty. First committed snapshot passed;
only docs/evidence follows, no successful rerun. Reflection confirms original
public class/constructor shape; Frame's two optional parameters really default null.

GoHP #76 merged maind6f32d90c686da8da4116d81f00da36ce71506d5, tree
773422be588abc30b748a742c2162cc460e76e12, identical clean feature tree. Both
feature/main CI543 checks/75 features PASS, PR CI also PASS. Trailing untyped
null constructor defaults reuse existing NullDefaultArgs and execution dispatch;
ordinary method/non-null/typed defaults remain rejected. Existing #75 tokenizer/
Execution and #74 constructors are preserved. Initial local PHP wrapper mounted
the wrong worktree: only that configuration was corrected and the uncompleted
comparison resumed; successful checks were not repeated. Details are in evidence.

Historical16/24/334 configurations/pins/corpora/evidence are unchanged/not rerun
or summed. Block.innerBlocks is empty; Frame containing an actual Block is now
measured, not full parser/tokenizer/stack/WP_Block_Type/registry/render/IDs/CSS/
HTML/ThemeJSON/bootstrap/HTTP/DB/themes/plugins. General casts/errors/coercions/
references remain unverified. No production target configured/created.
WebMCP excluded, route/auth unchanged. Devboard unidentified/not synced; PR/state
tracking. Next inspect the actual parser's remaining constructor/stack operations,
without treating data-class support as parser execution or duplicating tokenizer work.

## Previous: original parser-block class → snapshot → decision /16-case PASS

Complete unchanged WP_Block_Parser_Block15–90, its actual five-argument constructor
and five public properties now translate with two unchanged original helpers:
wp_should_add_elements_class_name/_wp_array_get. One owned source glue creates
the real original class and returns object + (array) snapshot; no substitute class
or observer field injection. All16 actual class/property/snapshot/decision/argument
records match PHP8.3.36/C and expectations (11true/5false decisions), inputs unchanged.
[Contract](parser-block-data.md), [record](evidence/parser-block-data-2026-10-26.md).

Clean snapshotba74d40a0cd803fa3212d2a2768bb1cc1e5051b9, tree
151aa7a02a74a39497bf6f8ce1d3196698bd6f70: PASS9.85s/package10.049s, vet
PASS on same snapshot; both exits0/stderr empty. First committed snapshot passed;
only docs/evidence follow and no successful rerun. Source/class and full ranges
are hash-bound; native Reflection confirms the exact five-public-property class/
constructor shape and both user-defined original functions.

GoHP #74 merged maineb3b968d447ab7a11af4efd3b1a1976b2cd277e8, exact clean
tested treeef8060b2e6a32b7da5784ebf13652092208bb059. Feature CI528-check/73-feature
PASS, PR/main CI also PASS; parallel #73 while preserved. Fixed untyped public
constructors, implicit-null properties and known-class public snapshots only.
No typed/defaulted/returning constructors, inheritance, other magic/direct re-entry
or general casts. Compiler local first snapshot also passed, no manual reruns.

The original parser uses the same creation/array-cast shape at its line149, but
that parser method/tokenizer/stack algorithm is not executed. innerBlocks admits
empty/data-array entries only, not nested parser objects. Full parser/WP_Block_Type/
registry/render/IDs/CSS/HTML/Theme JSON/bootstrap/HTTP/DB/themes/plugins and general
PHP errors/coercions/references remain unverified. Old24/334 configurations/pins/
corpora/evidence remain unchanged/not rerun; no summed integration claim.
No production target configured/created; WebMCP excluded, route/auth unchanged.
Devboard unidentified/not synced, tracked in PR/current-state. Next choose another
concrete original dependency/class/caller, coordinating with the separate compiler
owner's tokenizer work rather than reimplementing it or repeating PASS runs.

## Previous: source-built input → originals / separate24-case PASS

The owned input holder's construction/property assignment now execute in the
same translated PHP source as the three original utility/class-decision helpers.
Go/PHP observers call the source factory instead of constructing a concrete
generated struct and injecting supports through its exported field. All24 retained
cases match PHP8.3.36/C and expectations, including independently measured factory
source input and supports/caller data after the call, unchanged.
[Contract](elements-source-input.md), [record](evidence/elements-source-input-2026-10-26.md).

Clean test commit7b58f064d7f5da4ad7389eff17245ce3c376a99c, tree
55cc0b8ef7bcc29424e640046d9754cc9a4459df: PASS80.09s/package80.475s, vet
PASS, both exits0/stderr empty. First committed snapshot passed; no successful
rerun or code change after PASS. One binary has3 originals,2 owned functions and
1 input class; the24-case corpus/source bytes stay unchanged. Elapsed time is
not a controlled performance comparison.

GoHP #71 is merged maince38319b4245ccea0a7ae627f53b4396acabfce6, identical
tested feature tree0e7c3a7706235f32de040adc956c2d6fec1aca5d. FixedPHP CI all509
checks/70 features PASS, PR/main CI also PASS; parallel #70 str_starts_with retained.
The known-class declared-public-property write is bounded; mutable property
reads retain actual-value runtime guards. Initial compiler test-only raw-IR null
error/correction is recorded separately in PR71/evidence, not hidden as PASS.

Earlier24/334 configurations/pins/evidence are unchanged/not rerun; do not sum
their case counts. The holder is still owned data, not original WP_Block_Type/
registry/constructors. Unknown/foreign receivers, dynamic/nested/array-property
writes, render/CSS/HTML/IDs and full WordPress runtime remain unverified. No
production target configured/created; WebMCP excluded, route/auth unchanged.
Devboard unidentified/not synced, tracked in PR/current-state. Next pick an
actual remaining dependency or advance the source-only object boundary without
substituting WordPress services or repeating successful configurations.

## Previous: real serialization utility → class decision / separate24-case PASS

Three unchanged originals now run together: wp_should_skip_block_supports_serialization,
_wp_array_get and wp_should_add_elements_class_name. One owned data class and
one owned driver connect actual skip options to actual class necessity. All24
fresh Go/PHP8.3.36/C observations match (16 direct/8 composed), including20 bool,
1 null,1 int and2 string returns; measured caller data stays unchanged.
[Contract](elements-serialization-path.md), [record](evidence/elements-serialization-path-2026-10-26.md).

Clean snapshot60624d1f6b2ad5a219b09f80e66bb1ecb86c33fa, tree
d283e538c2da65134939f074d375c598072e0e86: PASS8.82s/package9.030s, vet PASS,
both exits0/stderr empty. Initial9d0c05c68 failed before translation because the
utils.php range ended at38 rather than35; corrected range only, original source
unchanged. Failure log retained; no rerun after successful snapshot.

GoHP #69 is_object is merged main731ef251ac907326f060aac5f1b5b994ac8615bb,
tree145b905f1f7c5ea2c00a2ab3825ecfcd1069b2b8, exact clean feature CI498-check/
68-feature PASS; PR/main CI also PASS. Parallel #68 str_contains preserved.
Existing26-original334-case configuration stays unchanged/not rerun. This is
not a27-original358-case single-binary result.

Observers supply supports data to the owned generated class outside translation;
the class provides no WordPress service logic. Original WP_Block_Type/registry,
source property writes/array defaults, render hooks/IDs/CSS/HTML and full bootstrap/
HTTP/DB/themes/plugins remain unverified. No production target configured/created;
WebMCP excluded, route/auth unchanged. Devboard unidentified/not synced, tracked
in PR/current-state. Next choose another concrete render dependency, without
replacing original services or repeating successful configurations.

## Previous: real block element-class decision /26 originals334 cases PASS

Unchanged `wp_should_add_elements_class_name` now executes with its actual
unchanged `_wp_array_get` dependency, without resolver/class/service stubs.
All26 originals and48 retained owned drivers translate/build once into one
executable;334 fresh comparisons match PHP8.3.36/C and expectations (294
retained plus40 caller cases).
[Contract](elements-class-path.md), [record](evidence/elements-class-path-2026-10-26.md).

Clean test commit `32ac5e0060574c2f250c601d8e13d1bbe48cf1e9`, tree
`71731182f2193e9c997c6519a098aa5464cf7fd9`; PASS9.84s/package10.082s,
vet PASS, exit0 and empty stderr. First committed snapshot passed; only docs/
evidence follows. No repeated successful run. GoHP #67 main
`a3039414cff22689d214a6c90767f13dccf56ddb` adds the bounded nested isset/??
contracts and retains parallel #66 array_slice; exact combined CI488 checks PASS.

Four original source files/32 fragments and ten retained corpora remain hash-
bound and unchanged. The40 new cases cover all26 configured color paths and
missing/null/falsy/empty-array/skip behavior. Actual bool results and both block/
options argument states match, with inputs unchanged. Class necessity is
presence-based (non-null), not CSS validity or truthiness.

This does not execute the surrounding render hook or generate CSS/HTML/unique
class names. Full Theme JSON/resolvers/bootstrap/HTTP/DB/themes/plugins and
general PHP errors/coercions remain unverified. Parsed inputs are arrays with
array/null path intermediates; scalar/string/object traversal is excluded.
No production target created; WebMCP remains excluded and route/auth unchanged.
Devboard unidentified/not synced, tracked in PR/current-state. Old settings/pins/
corpora/evidence remain exact historical configurations. Next select a concrete
dependency in the surrounding render path without substituting WordPress logic.

## Previous:25-original294-case integrated Go/PHP PASS

All25 unchanged original declarations and48 owned drivers now translate/build
together into one executable on GoHP #64 main
`b8dc1757ed61f459d51362da004aa1d73d024116`. All294 fresh comparisons match
PHP8.3.36/C and expectations:262 retained edge-bundle cases,24 standalone lookup
cases and8 new settings compositions. This is not a sum of old PASS counts.
[Contract](settings-helper-bundle-unit.md),
[record](evidence/settings-helper-bundle-2026-10-26.md).

Test commit `24c780b361dd4d11f4e18b79cc90d80f5d267872`, tree
`56b9e3e1cf2938f615d0a4717b3532c31e969d43`; PASS17.62s/package17.816s,
vet PASS on the same snapshot, exit0 and empty stderr. The first test snapshot
passed; only documentation/evidence changes follow, with no successful rerun.
Compiler473-check CI evidence is reused for its identical integrated tree.

All three original source files/31 fragments and nine old corpora are hash-bound
and unchanged. Exact-once observer adapters retain measured results and distinct
paths. Lookup expectation bindings include source/fragment/corpus hashes; Go and
PHP still execute every retained call. New compositions exercise lookup→Boolean
validation or lookup→slice→first/last/numeric/all, including missing versus null,
false-like strings, selector order/duplicates/mixed keys and raw scalar values.
Typed returns and caller arguments match, with new nested arguments unchanged.

Next extend this bounded helper integration toward a concrete caller/dependency
path. Theme JSON/global-style/block callers, full WordPress bootstrap/HTTP/DB/
themes/plugins and general PHP value/error semantics remain unverified. No new
production target. Local capability stays WebMCP-excluded, route/auth unchanged;
devboard unidentified/not synced, tracked in PR/current-state. Old configurations
and evidence remain available with their original pins.

## Previous: unchanged nested settings lookup24-case Go/PHP PASS

Original `_wp_array_get` now generates/builds and matches PHP8.3.36/C on all24
retained cases with GoHP #64 main `b8dc1757ed61f459d51362da004aa1d73d024116`.
The complete functions.php5233–5272 declaration, case corpus and PHP observer
are unchanged. Actual return types/values, including nested arrays, and
input/path/default states match; all caller arguments remain unchanged.
[Contract](array-get-unit.md), [record](evidence/array-get-unit-2026-10-26.md).

Test commit `ce5318d414d0701f6d1d312d7d81ddbaeb4e514d`, tree
`1c67b687f6ede81a061ba800c891b44eb4cbb213`; PASS9.12s/package9.359s,
exit0 and empty stderr. Vet PASS from the same Go sources is reused after
manifest-only compiler repin. Only docs/evidence changes follow successful test.

GoHP #63 added bounded null defaults and bare/multiple isset, reusing #62
continue. Initial Gossip snapshotd2bc85ce6 generated/built but failed at
`null === $path_element` on an array value. #64 fixed literal-null comparison
without general array/object identity; its exact clean feature tree passed473
CI checks. The failed execution and obsolete callable-null negative CI test
correction are retained in the record, not erased or counted as success.

This is the25th successful original in a separate unit. The old24-original262-
case bundle is unchanged; a25-original combined binary has not run. Next combine
the new lookup with the existing bundle and useful bounded compositions before
claiming integrated coverage. Theme JSON/block callers, full WordPress startup,
HTTP/DB/themes/plugins and broad PHP value/error semantics remain unverified.
Local capability stays WebMCP-excluded, route/auth unchanged; no new deployment
target. Devboard unidentified/not synced; PR/current-state tracks progress.

## Previous: nested settings lookup reference and compiler gaps

The next target is original _wp_array_get, called60 times across12 files at the
fixed WordPress revision, including43 Theme JSON and6 global styles/settings
calls. A24-case PHP8.3.36 reference records actual scalar/nested-array results
and unchanged input/path/default arguments. Source5233–5272 is byte-unchanged.
[Contract](array-get-boundary.md), [record](evidence/array-get-boundary-2026-10-26.md).

On GoHP #59 main `c3420d8644642c2fee56d62a08c88a2dcedc45ac`, the original
rejects at its default parameter. Four isolated probes separately confirm
trailing defaults, bare/multiple isset and foreach continue parser gaps.
All five exit1 with exact expected diagnostics and no Go/report artifacts.
This is native_reference_pass_candidate_rejected, differential not-run.
The successful24-original262-case integration remains unchanged.

Test commit `79f94688f21369ee932c608cda5dcce3664f6d08`, tree
`771df55853ed38faccaee77598822202d8fa3f55`; boundary PASS13.31s/package13.521s,
vet PASS. Initial extraction omitted the closing brace; the range check stopped
before native execution and was corrected. Evidence retains both snapshots.
After PASS only docs/evidence changes follow.

Next implement the three bounded compiler contracts and compare the unchanged
original with a new pin. Theme JSON callers, block rendering, HTTP/DB/plugins
and broader values/errors remain unverified. Local capability is WebMCP-excluded;
no route/auth change. PR/current-state tracks progress, devboard unidentified.

## Previous: twenty-four-original262-case integration PASS

All24 original declarations and46 owned drivers translate/build together on
GoHP #58 main `440308c77da07cbcdce403fed7fff3945e0d2078`. A single executable
matches PHP8.3.36/C for all262 cases:230 retained search-bundle,24 edges and8
new cross-helper compositions. All old corpora and source bytes are preserved.
Actual typed returns and caller input/selector state match PHP and expectations.
[Contract](edge-helper-bundle-unit.md), [record](evidence/edge-helper-bundle-2026-10-26.md).

Successful test commit `c753c2d37ada08e2c412ea6996515f06dcf0b500`, tree
`69e40a5525a299f7c193fe4f52eae96e9961aad5`; PASS9.28s/package9.489s,
vet PASS, empty stderr/exit0. The initial snapshot failed to build an adapted
observer because its JSON type import was removed. Retaining that import fixed
the adapter; the comparison then completed. Evidence records both snapshots.
Only documentation/evidence changes follow PASS.

The eight new cases compose slice → Boolean mapping → first/last and search.
They distinguish empty/null from singleton false, no-match from key0, selector
order/deduplication and mixed/normalized keys. Older configurations remain exact
historical records; this is a newly measured integration rather than summed PASS
counts. WordPress startup/HTTP/DB/themes/plugins and broader callable/error
semantics remain future work. Next select a concrete original-source dependency
or application path to extend beyond the current bounded helper integration.
Local capability remains WebMCP-excluded; route/auth unchanged. PR/current-state
tracks progress, devboard card unidentified.

## Previous: original array_first/array_last24-case comparison PASS

Both originals now translate on GoHP main
`440308c77da07cbcdce403fed7fff3945e0d2078` (#58) and execute in one binary.
All24 retained cases (12 each) match PHP8.3.36/C, actual scalar return types/
values and ordered caller inputs. First PASS11.16s/package11.364s, empty stderr/
exit0; vet PASS with a writable cache after the default-cache sandbox denial.
Test commit `ecb89beed327dd14cd3772bfdb9e752074816a0f`, tree
`ed72977216cb8b70747b14a8c88804db74124507`. Original source/cases/PHP observer/
transports unchanged. No compiler change or repeated comparison.
[Contract](array-edge-unit.md), [record](evidence/array-edge-unit-2026-10-26.md).

The old rejected boundary and first-only configuration remain historical
evidence for their exact pins. This new two-original configuration does not
extend the22-original230-case integration automatically. A larger single-binary
integration is the next step. Full WordPress startup/HTTP/DB/themes/plugins
and broader value domains remain unverified. Developer capability remains
WebMCP-excluded, route/auth unchanged. PR/current-state tracks this work;
devboard card remains unidentified.

## Previous: original array_first comparison PASS

Original array_first now translates and matches PHP for all12 retained cases
using GoHP main `4a4fffb995e57daf592785b6626e45a7babe08da` (#57).
Test snapshot `fdf018beee72a4edb5e011ee8c3ed7e0afb0bae1`, tree
`66757ce66aece6ec82cb29e3558b77d94eb95f85`; first PASS14.95s/package15.187s,
vet PASS. Actual PHP8.3.36/C, empty stderr/exit0. Declaration and native loading
guard remain byte-unchanged; native Reflection confirms the user-defined polyfill.
Actual return type/value and ordered caller input match authored expectations.
[Contract](array-first-unit.md), [record](evidence/array-first-2026-10-26.md).

This unit uses the already-merged #57 implementation. The earlier locally saved
empty implementation is not adopted. Previous pins/corpora/rejection evidence are
retained, and the22-original/230-case integration remains its own configuration.
array_last still requires array_key_last; larger integration and WordPress
startup/HTTP/DB/theme/plugin behavior remain future work.

The owner authorized verified changes through PR and merge, and necessary
deployment during this development stage, without repeating approval prompts.
This change adds a developer comparison unit. No deployable runtime is introduced.
The local capability remains WebMCP-excluded. Devboard card is unidentified;
the PR and this record track progress.

## Previous: default-branch consolidation

The owner approved integrating the validated development stack into `master`.
The integration branch combines candidate `afaa7e38980a150fd9ae43bd8194c67b26ee266b`
with the README-only master `9b72be0e5d4dbb5a36e4c5aaaca4091a7b662d3c`.
Only README/current-state/index documentation is reconciled; implementation,
test code, compiler pins, original WordPress files and historical evidence stay
unchanged. Merge history is retained, not squashed or rebased.

The preceding GoHP integration is merged as
[`6e385ea`](https://github.com/katsuakikureyama/GoHP/pull/56), with the exact
source tree of its 420-check PHP 8.3.36 PASS. This does **not** repin Gossip:
the latest 22-original/230-case bundle still uses GoHP #54 `36191c5`.
Its existing PASS and the separate native-only 24-case boundary record below
remain the evidence; this documentation-only reconciliation does not rerun them.

Earlier "master unchanged", "Draft" and "no merge" statements below are
historical execution/publication status, not a continuing ban on integration.
New work should branch from the integrated master and merge at verified
milestones instead of indefinitely extending a Draft stack. No release,
deployment, branch deletion or prototype change is part of this consolidation.
The new bounded `empty` work remains local and unverified, outside this baseline.
Devboard card remains unidentified/unsynchronized; full WordPress compatibility
is not Done. PR merge state and the default-branch SHA are the publication record.

## Latest: array_first/last reference and explicit compiler blockers

- Branch `codex/gohp-array-edge-boundary` on #29. Test commit
  `6edeffa6cf568004babcdc59622e9fbd9662ab49`, tree
  `9a820ac45f284e679f8249338ef06dea7d76cf41`; clean commit before execution.
- Original guarded array_first/array_last native reference matches all24 typed
  return/input-after cases. GoHP #54 unchanged. Both original declarations reject
  with `PHP004 line 3: unsupported function empty`; an owned dependency probe
  rejects with `PHP004 line 2: unsupported function array_key_last`.
  CLI exit1/empty stdout/no generated/report artifacts required. Native exit0.
- First boundary-test PASS10.85s/package11.099s and vetPASS, actualPHP8.3.36/C.
  No unexpected failure/retry. Status is native_reference_pass_candidate_rejected,
  not differential PASS. Earlier22-original/230-case integration is unchanged.
  [Contract](array-edge-boundary.md), [record](evidence/array-edge-boundary-2026-10-26.md).
- Next: bounded initialized-array-variable empty support, then original array_first
  comparison on a new pin; separately array_key_last and original array_last.
  No source rewrite, compiler modification/retest or old configuration rerun.
  Original WordPress/SDK/hooks/oracle/prototype/master unchanged; no merge/deploy.
  Local capability WebMCP-excluded; devboard unidentified/unsynchronized.

## Previous: twenty-two-original integration PASS

- Branch `codex/gohp-search-helper-bundle`, based on Draft #28 head
  `cdbc55b8eb80363935dfca5a7d519066bfaf138c`; GoHP #54 pin unchanged.
  Test commit `97b264cfb7bf17d61bfec31c4103ecb9d67b17b8`, tree
  `037353ebeee9884f336fe9b8f151eba39bad7054`. Clean commit before execution.
- Twenty-two byte-unchanged original declarations/44 owned drivers translate
  together into one binary. All230 cases (192 retained numeric-bundle,30 search,
  8 new composed) match actual types/values and caller input-after against PHP.
  New selection/Boolean/search compositions distinguish matched false, no-match
  null and key0. All previous corpora/settings/evidence remain unchanged.
- First PASS20.52s/package20.766s, vetPASS, actualPHP8.3.36, C locale,
  empty stderr/exit0; no failures/retries. Only docs/evidence afterward;
  no repeated old configurations/compiler tests/new CI dispatch.
  [Contract](search-helper-bundle-unit.md) and
  [record](evidence/search-helper-bundle-2026-10-26.md) contain scope and hashes.
- Declaration extraction retains annotations/bodies; native executes complete
  original guarded polyfills. Conditional loading/full bootstrap remains outside
  the compiled unit. No general callbacks/captures/errors, HTTP/DB/theme/plugin
  compatibility or release readiness. WordPress originals/SDK/hooks/oracle policy/
  prototype/master unchanged. Local capability WebMCP-excluded; devboard unsynced.
- Next: select a concrete unsupported original-source dependency for a separate
  bounded change; retain this PASS without repeated verification.

## Previous: three-original search unit PASS; nineteen-original integration retained

- [Draft #27](https://github.com/katsuakikureyama/Gossip/pull/27) publishes the
  nineteen-original/192-case single-binary integration on GoHP #54.
  Test commit `9972a6f92ad3d3e889d8d71b128d893966d65682`, evidence head
  `baaac5474d751fc38fa55875ab862ffdb9564585`. All192 cases and vet PASS;
  [exact record](evidence/numeric-helper-bundle-2026-10-25.md). Published with
  owner approval on 2026-10-26, not merged. No retest on publication.
- The earlier two-original/31-case #53 unit also passed; see its
  [record](evidence/numeric-array-helper-2026-10-25.md). The rejected #44 boundary
  below is historical, not the current compiler result.
- Branch `codex/gohp-array-search-polyfills` adds a separate three-original
  search unit on the same #54 pin. It is not a22-original integrated PASS.
  [Contract](array-search-polyfills-unit.md), [execution evidence](evidence/array-search-polyfills-2026-10-26.md).
  Test commit `861814fcfecdc0eba0aca2e9f61ed6be1677bc03`, tree
  `ff3e64f9e016944c82c80d9a63a5e10b24788fc4`: first30/30 PASS9.23s/package9.568s,
  vetPASS, empty stderr/exit0, no retries. ActualPHP8.3.36, original guards retained
  for native; compiler extracts unchanged declarations. No compiler retest.
  Next: separately define a larger integration; do not add30 to192 as proof.
- WordPress originals, previous settings/evidence, SDK/hooks, oracle policy,
  prototype and master unchanged. General callback/capture/bootstrap/core/HTTP/
  DB/theme/plugin compatibility is not established. Local capability remains
  WebMCP-excluded; devboard unidentified/unsynchronized.

## Previous: numeric-array reference and rejected compiler boundary

- Branch `codex/gohp-numeric-array-boundary` on #24. Test commit
  `90870c68d01c9d399dcd6060b85ce38363096373`, tree
  `3d5e6e11af8f839beb373c916f683f1ea92c4011`; staged diff check and clean
  commit completed before the first execution.
- Selected unchanged `wp_is_numeric_array()` and the complete original guarded
  `array_all` polyfill for PHP8.3. All24 native reference cases match authored
  actual return/list/input-after expectations. Sparse integer-key arrays qualify
  despite not being lists; canonical numeric-string keys normalize first.
- GoHP unchanged at #44 `49a302e44464623dc4912bafbcf44192759d7878`.
  Separate unchanged declarations reject at the Boolean return type and typed
  parameter, both exit1 with no generated Go or report. Source review also
  identifies arrow/callable/dynamic-invocation dependencies. No candidate binary
  or Go/PHP differential run exists; this is NOT an eighteenth translated helper.
- First boundary test PASS7.60s/package7.839s; vetPASS, actualPHP8.3.36.
  Native stderr empty/exit0; compiler rejection stderr/exit1 recorded separately.
  No failure/retry. Post-PASS docs only. Existing oracle isolation/policy retained.
  [Contract](numeric-array-boundary.md) and
  [evidence](evidence/numeric-array-boundary-2026-10-25.md) retain exact commands,
  hashes, all24 native observations and the distinct rejected/not-run statuses.
- Next: implement bounded Boolean return declarations, then the typed-parameter
  and callback dependency chain with compiler-native evidence before attempting
  this unchanged original-source differential unit. No polyfill stub, source
  rewrite or PHP-series bypass. Survey was bounded, not an all-core inventory.
- Prior17-original/149-case integration is still the latest successful combined
  result. Older configurations/evidence, compiler source, WordPress originals,
  SDK/hooks, prototype and master unchanged; no old tests or CI rerun. No merge/
  deploy/release/license decision. Local developer capability WebMCP-excluded,
  auth routes unchanged; devboard unidentified/unsynchronized, not Done.

## Previous: seventeen-original yes/no integration

- Branch `codex/gohp-yn-helper-bundle` on #23. Test commit
  `22a5a4d9ab9dac84402e3c22d4d7a91ffa9963f3`, tree
  `6e1cec3c6132026cf623c6dee9d21ddba9db78a7`; committed before execution.
- Seventeen byte-unchanged originals and25 owned drivers translate/build once
  into one binary. All149 calls (direct103/composed46) match actual typed values,
  ordered keys and caller arguments against PHP8.3.36 and authored expectations.
  Previous119 cases plus24 direct yes/no and6 new compositions. The two original
  validators preserve their distinct results for the same string: yes→[false,true].
- First run PASS10.14s/package10.336s, empty stderr/exit0, vetPASS. Parent corrected
  a draft bool-return assertion and removed a duplicate helper before commit/
  first execution. No runtime failure/retry; post-PASS docs only.
  [Contract](yn-helper-bundle-unit.md) and
  [evidence](evidence/yn-helper-bundle-2026-10-25.md) retain all observations,
  hashes, exact commands and pre-execution review corrections.
- Compiler unchanged: GoHP `49a302e44464623dc4912bafbcf44192759d7878` (#44),
  tree `dc2046b28420a9f7a9da7447052a5e70d6543279`. Earlier CI40159749575
  PASS336 checks/52 features reused without compiler tests/CI rerun. Newer
  control-flow compiler branches are not adopted.
- New yes/no direct/pair inputs are strings; new slice roots contain only
  strings/null, with null removed by original slice before mapping. Prior HEIC
  scalar and nested-list domains retained. Eleven standalone and all old bundle
  settings remain unchanged/not rerun. WordPress originals/SDK/hooks/oracle
  policy/prototype/master unchanged. No merge/deploy/release/license decision.
- Next: select the next bounded original-helper dependency. General Boolean
  coercion, option-backed wrappers, array_all/arrow closures, whole core/HTTP/
  DB/themes/plugins remain unverified. Local developer capability WebMCP-excluded,
  auth routes unchanged; devboard unidentified/unsynchronized, not Done.

## Previous: original yes/no-label predicate

- Branch `codex/gohp-bool-yn` on #22. Test commit
  `0f236704a0fcb3140dd7f8c3699ed3aa8abfa615`, tree
  `98c7d07b8844a25eba9aa885ea235175670a8491`; committed before execution.
- Unchanged `bool_from_yn()`, functions.php1621–1623, runs using existing
  string-only strtolower and strict equality. All24 strings match actual bool
  returns and unchanged caller input against PHP8.3.36 and authored expectations.
  y/Y are true; words, digits, whitespace, NUL and non-ASCII cases are false.
  First run PASS13.23s/package13.483s, empty stderr/exit0, vetPASS. No retry;
  post-PASS docs only. [Contract](bool-yn-unit.md) and
  [evidence](evidence/bool-yn-2026-10-25.md) record exact hashes and observations.
- Compiler unchanged at GoHP `49a302e44464623dc4912bafbcf44192759d7878`
  (#44), tree `dc2046b28420a9f7a9da7447052a5e70d6543279`. Prior native
  CI40159749575 PASS336 checks/52 features reused without compiler test/CI rerun.
  Separate control-flow compiler integration is not adopted.
- Eleventh standalone unit/seventeenth distinct original helper, not a17-original
  integrated run. Earlier ten units and16-original119-case bundle remain
  unchanged and not rerun; original WordPress/SDK/hooks/oracle policy/prototype/
  master unchanged. No merge/deploy/release/license decision.
- Next: separately pin a17-original integration with the existing helpers.
  Non-string Boolean coercion, public option-dependent charset wrappers,
  array_all/arrow closures and full core/HTTP/DB/themes/plugins remain unverified.
  Local developer capability WebMCP-excluded, auth routes unchanged;
  devboard unidentified/unsynchronized, not Done.

## Previous: sixteen-original HEIC/charset/array integration

- Branch `codex/gohp-heic-helper-bundle` on #21. Tested tree
  `eabbffa481c23c04ed1c29380c2b2845ae8d1171`, recorded unchanged in commit
  `fb80796656c2628f73f086e8e53c53356e371b15` during execution. A pre-commit
  diff-check warning for PHP-observer EOF whitespace prevented the initial
  commit; staged bytes were unchanged. See evidence for launch-head/timing.
- Sixteen byte-unchanged originals from three source files and22 owned drivers
  translate/build once into one executable. All119 calls (direct79/composed40)
  match actual typed values, ordered keys and caller arguments against PHP8.3.36
  and authored expectations. Prior89 charset cases plus24 direct HEIC and6
  new slice/unescape/Boolean compositions. First runtime PASS32.63s/
  package32.878s, empty stderr/exit0, vetPASS. Post-PASS docs only.
- Compiler remains `49a302e44464623dc4912bafbcf44192759d7878` (GoHP #44),
  tree `dc2046b28420a9f7a9da7447052a5e70d6543279`. Existing CI40159749575
  PASS336 checks/52 features is reused, not rerun; #43 control-flow not adopted.
  [Contract](heic-helper-bundle-unit.md) and
  [evidence](evidence/heic-helper-bundle-2026-10-25.md) retain all hashes,
  actual119 observations, commands and the snapshot timing exception.
- New HEIC operands are scalar-only, including unselected values in new slice
  roots. Older nested-list cases remain. This is MIME-label/charset-name
  classification, not image contents, decoding, uploads or UTF-8 validation.
- Ten standalone and all earlier bundle configurations remain unchanged and
  not rerun. Original WordPress/SDK/hooks/oracle policy/prototype/master unchanged.
  No merge/deploy/release/license decision. Local developer capability
  WebMCP-excluded, auth routes unchanged; devboard unidentified/unsynchronized.
- Next: identify the next original-helper dependency and its bounded contract.
  General PHP, option-dependent wrappers and full core/HTTP/DB/themes/plugins
  remain unverified; this does not measure every historical standalone input.

## Previous: original HEIC/HEIF MIME-label predicate

- Branch `codex/gohp-heic-mime`, stacked on #20. Test commit
  `dde6850d859b0a51a4c05fefda19c4ef4cba8c61`, tree
  `ccdc46abd387ce9f3b31dc28266dd813db6eb4f4`, fixed before execution.
- WordPress functions.php9465–9474 is extracted byte-unchanged. All24 scalar
  cases match actual bool returns and caller input against PHP8.3.36 and
  authored expectations. MIME-label classification only, not image inspection.
- GoHP `49a302e44464623dc4912bafbcf44192759d7878`, tree
  `dc2046b28420a9f7a9da7447052a5e70d6543279`, Draft #44 on #41,
  adds strict scalar `in_array(..., true)` with whole-domain guards.
  Fixed PHP CI40159749575 passes336 checks/52 features plus vet/CLI smoke.
  Local runtime/codegen/catalog/static-negative/generated-guard/vet PASS.
- Gossip initial host compilation failed for one unused import; that import
  was removed before the successful run. Corrected snapshot PASS9.07s/
  package9.261s, vetPASS, empty stderr/exit0. After PASS docs only.
  [Contract](heic-mime-unit.md), [evidence](evidence/heic-mime-2026-10-25.md)
  record exact attempts, commands, pins, observations and compiler CI artifact.
- Tenth standalone, not16-original integration. Older nine units and all
  bundles including15-original/89-case remain pinned and not rerun. WordPress
  original/SDK/hooks/oracle policy/prototype/master unchanged; no merge/deploy/
  release/license decision. Browser/S2S/MCP authentication unchanged; local
  developer capability WebMCP-excluded; devboard unidentified/unsynchronized.
- Next: a separately pinned16-original integration can compose the new
  predicate with existing array selectors and Boolean helpers. No such run
  yet; general PHP, image contents and whole core/HTTP/DB/themes/plugins
  remain unverified. Separate if/else/elseif compiler stack is not integrated.

## Previous: fifteen-original charset/array integration

- Branch `codex/gohp-charset-helper-bundle`, stacked on #19. Test commit
  `b0835877b9e2172670582a71697f066833837257`, tree
  `27dd48a8bfcbb863b76cb483c04b93f21094c4af`, committed before execution.
- Fifteen byte-unchanged originals from three pinned source files plus
  nineteen owned drivers translate/build once into one binary. All89 calls
  (direct55/composed34) match actual typed returns, ordered keys where relevant
  and caller arguments against PHP8.3.36 and authored expectations.
  Earlier55 associative cases plus28 direct charset cases and6 compositions.
- GoHP `6305b9b486ee233e9aa4378ab8528f5b3d345f90`, tree
  `48c92101a7000794eda13cdfbbd0240b7327afb9` (Draft #41 on #40), explicitly
  merges #40 string/logical comparison and #36 isset/typed-array returns.
  Both parents retained; separate #39 if/else excluded. Reserved-name,
  fixture and spec conflicts preserve both contracts. New cross-stack native
  fixture linked to existing catalog entries;51 features/313 checks.
- Local Gossip comparison first PASS19.04s/package19.279s, empty stderr/exit0,
  vetPASS. Compiler local runtime/frontend/codegen/compat/guards/vet PASS;
  fixed PHP CI40159252743 overall PASS313 checks, actualPHP8.3.36/Go1.23.12,
  exact clean head/tree, vet/CLI smoke/artifacts also successful.
  [Contract](charset-helper-bundle-unit.md) and [evidence](evidence/charset-helper-bundle-2026-10-25.md)
  retain revisions, commands, hashes, observations and pre-execution review
  corrections. No test failure/retry; post-PASS documentation only.
- Prior nine standalone and10/12/13/14-original configurations/evidence stay
  unchanged, not rerun. Original WordPress/SDK/hooks/PHP admission/prototype
  and master README-only state unchanged. No merge/deploy/release/license decision.
- Next: identify the next bounded original-helper dependency independently.
  Public option-dependent charset wrappers, general callbacks/objects/
  references, nested associative transport, map_deep/wp_unslash and whole
  core/HTTP/DB/theme/plugin operation remain unverified. Local developer
  capability WebMCP-excluded, auth routes unchanged; devboard unidentified/
  unsynchronized, not Done.

## Previous: original UTF-8 charset-name predicate

- Branch `codex/gohp-utf8-charset`, stacked on #18. Test commit
  `fccd38e47c4f23ed0bb8b35a0cc974c9486f4d2f`, tree
  `64499078c9dd73381ba28bce7c3cb889392981cc`, committed before execution.
- Unchanged `_is_utf8_charset`, compat.php102–111, executes on GoHP
  `d9f6c7b4f288ef816c4ee3cb6a3458d2226e00d5` (Draft #40 on #38).
  New guarded string-only strcasecmp plus existing !/||/is_string/===.
  Original source is neither rewritten nor replaced with an owned algorithm.
- All28 cases match actual bool returns and post-call inputs against PHP8.3.36
  and authored expectations; empty stderr/exit0. First local comparison
  PASS37.69s/package37.898s, vetPASS. No failure/retry. Post-PASS docs only.
- Compiler native CI40158737961, clean exact head/tree, PHP8.3.36 and Go1.23.12,
  passes all296 checks/49 features plus vet/CLI smoke/artifacts. Parent local
  runtime/codegen/compat and generated guard tests/vet also PASS. Exact
  [contract](utf8-charset-unit.md) and [evidence](evidence/utf8-charset-2026-10-25.md).
- Ninth separate unit, not a fifteen-original bundle. The #36 array-return/
  isset stack is absent from this compiler pin; earlier eight units and all
  bundles remain pinned and unchanged. Original files/SDK/hooks/PHP admission
  and master README-only publication are unchanged.
- Next: integrate the published compiler stacks in a separate change before
  a new fifteen-original bundle. Public is_utf8_charset/defaults/coalescing/
  get_option, _canonical_charset, text encoding conversion and full core/
  HTTP/DB/theme/plugin execution remain unverified. No merge/release/deploy
  or license decision. Local capability remains WebMCP-excluded; auth routes
  unchanged. Devboard unidentified/unsynchronized, not Done.

## Previous: fourteen-original associative bundle

- Branch `codex/gohp-associative-helper-bundle`, stacked on #17. Tested
  snapshot `9c14ce515c266d9db6d1f125a889c9e93cfc3dd1`, tree
  `c5ae3be90a0ed10feb4146f05bac6bb6f1af6879` (exact staged bytes committed
  after execution). Fourteen byte-unchanged originals and sixteen owned
  drivers translate/build once into one executable.
- All 55 cases PASS: previous 33 recursive-bundle cases, fourteen standalone
  slice cases and eight new compositions. Actual return type/value, ordered
  keys where applicable and both caller arguments agree with PHP 8.3.36 and
  authored expectations. Direct27/composed28; empty stderr/exit0. Local test
  8.41s/package8.629s and vet PASS, first comparison, no rerun.
- GoHP pin remains `db3051f8a179078e2227c5d97f50a2f1611d1451` (#36).
  Existing exact-commit CI40156201018 (238 checks/46 features) and earlier
  compiler local tests are explicitly reused, not rerun. Other compiler
  branches are not integrated. Prior eight standalone and ten/twelve/thirteen
  bundle configurations and evidence stay untouched.
- [Contract](associative-helper-bundle-unit.md) and
  [evidence](evidence/associative-helper-bundle-2026-10-25.md) retain exact
  hashes/observations. A pre-commit whitespace check failed on a trailing
  PHP EOF blank line; after successful runtime validation and exact-snapshot
  commit, only that cosmetic line was removed alongside documentation.
  No executable statement/configuration changed, so runtime PASS is reused.
- Next boundary: identify a concrete missing original-helper dependency
  before extending the compiler. Nested associative values, general objects/
  references/callbacks, map_deep/wp_unslash and full core/HTTP/DB/theme/plugin
  operation remain unverified. No implementation merge/release/deploy or
  license decision. Master README-only publication remains untouched.
  Local developer capability remains WebMCP-excluded; auth routes unchanged.
  Devboard unidentified/unsynchronized, not Done.

## Previous: original associative array slice

- Branch `codex/gohp-array-slice-assoc`, based on #15. Test commit
  `1b191a17b574f174e04249b6420e893ec57750c5`, tree
  `c0b85668909beca251efceb4d56351af08c3538e`.
- Exact original wp_array_slice_assoc, functions.php5173–5183 including
  `: array`, translates/builds/executes through pinned GoHP
  `db3051f8a179078e2227c5d97f50a2f1611d1451` ([Draft #36](https://github.com/katsuakikureyama/GoHP/pull/36)
  on #34). Compiler enforces global array returns including fallthrough,
  inherits one-level isset and indexed writes. Separate #33/#35 not included.
- Fourteen explicit cases use ordered root key/value pairs and selector lists.
  Missing/null versus falsy values, key normalization/order/collisions,
  duplicates, negative/NUL/UTF-8 keys, nested lists and both caller arguments
  after invocation are observed. Actual Go/PHP/expected comparisons all PASS,
  empty stderr/exit0, local8.42s/package8.619s/vetPASS. One earlier worker
  invocation failed at host cache setup before any test/Docker/PHP execution;
  first actual comparison passed. No post-PASS code/config change or rerun.
- Exact [contract](array-slice-assoc-unit.md) and
  [evidence](evidence/array-slice-assoc-2026-10-25.md) bind test/compiler/source/
  observer/corpus/report/binary hashes. Compiler local tests/vet and fixed
  CI40156201018 PASS238 checks/46 entries; actualPHP8.3.36/Go1.23.12,
  clean checkout equals pinned commit/tree.
- Eighth independent unit. Earlier seven standalone units and ten/twelve/
  thirteen-original bundles keep their pins/corpora/results; no combined
  fourteen-original test yet. Next: new pinned combined integration.
- Original files, SDK/hooks and PHP admission unchanged. Approved README-only
  master publication from #16 remains intact; implementation is Draft only,
  no merge/release/deployment/license decision. Core/HTTP/DB/themes/plugins,
  general object/reference/callback/map_deep/wp_unslash remain unverified.
  WebMCP-excluded local tooling, no auth changes; devboard unidentified/unsynced.

## Previous: thirteen-original recursive helper bundle

- Branch `codex/gohp-recursive-helper-bundle`, [Draft #15](https://github.com/katsuakikureyama/Gossip/pull/15)
  on #14; test commit `97ed80c3db187ce9ffb2a73c27dd5207f8594192`,
  tree `baef43b9bed5efd39b256d7fe705d93a133fafbf`.
- Previous twelve originals plus unchanged wp_slash and eleven owned PHP
  composition drivers translated/built once into one executable. New33-case
  configuration (direct13/composed20) includes string roundtrips, shallow list
  stripping, capture/newline and shape/keys/values compositions. Every actual
  observation also checks caller arguments after invocation remain unchanged.
- Initial local isolated comparison PASS8.19s/package8.387s, all33 actual
  return types/values/arguments match PHP8.3.36 and explicit expectations;
  empty stderr/exit0, LC_CTYPE=C, affected-package vetPASS. No failure/retry
  or implementation change after PASS. Exact [evidence](evidence/recursive-helper-bundle-2026-10-25.md)
  and [contract](recursive-helper-bundle-unit.md) retain all hashes/boundaries.
- Compiler/runtime unchanged at03d433e73503e32a84f310f023fc0f65c9b188cd
  (GoHP #27), existing exact-tree CI40152066817 PASS181 checks reused, not rerun.
  Earlier seven standalone units and ten-/twelve-function bundles retain their
  pins/corpora/results. Their complete earlier inputs are not relabelled as
  verified in this new combined configuration.
- wp_unslash/map_deep needs element/property mutation, object traversal and
  dynamic call_user_func; none is implemented here. Next: isolate a missing
  compiler contract such as array writes before claiming map_deep support.
  Full-core startup/HTTP/DB/themes/plugins remain unverified. Draft only,
  no merge/deploy/release/license decision. WebMCP-excluded local tooling,
  no auth changes; devboard unidentified/unsynced, not Done.

## Previous: original recursive wp_slash

- Branch `codex/gohp-wp-slash`, [Draft #14](https://github.com/katsuakikureyama/Gossip/pull/14)
  on #13; test commit `d3f02dc3625d331a477cd5fb8655f4d28e864432`,
  tree `2a6af2b6070e1bf36b18bc0cd93d75b484ee1720`.
- Complete unchanged formatting.php5933–5943 translated through compiler/runtime
  `03d433e73503e32a84f310f023fc0f65c9b188cd`,
  [GoHP Draft #27](https://github.com/katsuakikureyama/GoHP/pull/27) on #24,
  incorporating #25 is_array. Adds literal declared-global one-argument callback,
  single-input array_map and actual-string-only addslashes. No general callback
  system, weak coercion, object traversal or arbitrary-depth claim.
- Initial local PHP8.3.36 comparison PASS22 cases: actual return type/value and
  caller input after invocation match PHP and explicit expectations, empty
  stderr/exit0. PASS8.19s/package8.423s; vetPASS. Compiler local test typo caused
  one codegen FAIL, corrected before CI and only affected package rerun; final
  local tests/vetPASS. Fixed CI40152066817 PASS181 checks,42 catalog entries,
  PHP8.3.36/Go1.23.12; clean CI tree equals head.
- [Contract](wp-slash-unit.md) and [evidence](evidence/wp-slash-2026-10-25.md)
  preserve exact commands/revisions/hashes, failure/fix and unverified boundaries.
  Seventh standalone unit; previous six and both bundles unchanged/not rerun.
  No thirteen-function bundle, full WordPress startup, wp_unslash/map_deep,
  HTTP/DB/theme/plugin support. Next: select a concrete remaining original-source
  dependency, or test new composition under a separately pinned configuration.
- Draft only; no merge/deploy/release/license decision. Original files/SDK/hooks,
  PHP admission policy and other GoHP checkout unchanged. WebMCP-excluded local
  tooling, no auth changes; devboard unidentified/unsynced, not Done.

## Previous: twelve-original bundle on the integrated compiler

- Branch `codex/gohp-integrated-helper-bundle`, [Draft #13](https://github.com/katsuakikureyama/Gossip/pull/13)
  stacked on #12. Test commit `74e234dcd80674c03b7f718040e8804866c8b04f`,
  tree `ae04dc19a38c88556f56fc9e4de707a4849f2159`.
- Twelve unchanged WordPress functions plus seven owned PHP drivers are
  translated/built together into one executable. New24 observations:12 direct
  probes and12 compositions (captures→newline→unslash, array passthrough→keys
  or list shape/actual values, and four scalar chains).
- Compiler/runtime `e5ba8c1db81d235cc177cc83699868d812abff7e`,
  [GoHP Draft #24](https://github.com/katsuakikureyama/GoHP/pull/24), stacked
  on #22 and integrating #23. Full builtin reservations/fixtures/guards retained;
  shared semantic validation reconciled. No new runtime domain or PHP policy.
- Initial local Gossip PASS7.60s/package7.834s, actualPHP8.3.36, all24 typed
  values/expectations match, empty stderr/exit0; vetPASS. Compiler local tests
  and vetPASS; fixed [CI40149530122](https://github.com/katsuakikureyama/GoHP/actions/runs/40149530122)
  PASS166 checks across40 catalog entries, Go1.23.12/PHP8.3.36. Clean CI tree
  equals compiler head. No manual retry or implementation change after PASS.
- [Contract](integrated-helper-bundle-unit.md) and
  [evidence](evidence/integrated-helper-bundle-2026-10-25.md) record exact
  commands/pins/hashes. Existing six standalone units and ten-helper bundle
  retain their pins/results, not rerun or relabelled. Original files, SDK/hooks,
  native admission implementation and other session checkout unchanged.
- This verifies one bounded combined configuration, not whole files, recursive
  callbacks/unescaping, core startup, HTTP/DB/themes/plugins. Draft only,
  no default-branch merge/deploy/release/license decision. Local tooling remains
  WebMCP-excluded, auth routes unchanged; devboard unidentified/unsynced, not Done.

## Previous: string-only stripslashes helper

- Branch `codex/gohp-stripslashes-helper`, [Draft #12](https://github.com/katsuakikureyama/Gossip/pull/12)
  stacked on #11. Test commit `f4e65138298ae6763670d457c9c2f74615b8103a`
  translates unchanged `stripslashes_from_strings_only()` at formatting.php2934–2936.
- Compiler/runtime `0e49cc995913dfede0332566dd4dc11c5df16ff8`,
  [GoHP Draft #23](https://github.com/katsuakikureyama/GoHP/pull/23), on #21.
  Full three-arm ternary: condition once, selected arm only, both arms statically
  validated; one-actual-string stripslashes with static/runtime guards.
  Catalog35 entries/89 checks. No short ternary, coalescing, coercion or recursion.
- New21 typed cases matched PHP8.3.36 and expectations, empty stderr/exit0:
  initial PASS7.87s (package8.078s), vet PASS. Includes non-string passthrough
  and escaped strings inside nested lists deliberately left unchanged.
- Compiler local runtime/codegen/catalog tests and vet PASS; fixed
  [CI40121644555](https://github.com/katsuakikureyama/GoHP/actions/runs/40121644555)
  PASS89 checks, Go1.23.12 / PHP8.3.36. Clean CI tree equals compiler head.
  Exact revisions/commands/hashes in [evidence](evidence/stripslashes-helper-2026-10-25.md),
  with [contract](stripslashes-helper-unit.md). No manual rerun after PASS.
- Sixth standalone unit, not a twelve-function bundle or recursive
  `stripslashes_deep` / `wp_unslash` support. Earlier pins/evidence unchanged,
  no reruns; other GoHP stack and normal checkout unchanged. Draft only,
  no merge/deploy/release/license decision; WebMCP-excluded local capability,
  no auth route changes. Devboard card unidentified/unsynced, not Done.

## Previous: newline-preservation helper

- Branch `codex/gohp-newline-helper`, stacked on Draft #10; test commit
  `ff02bbd05a0dd5b75940724cda2352622c543030` translates unchanged
  `_autop_newline_preservation_helper()` from formatting.php lines 811–813.
- Compiler/runtime `dcb72c64d029c7a08a0e94256723ff4076209758`
  ([GoHP Draft #21](https://github.com/katsuakikureyama/GoHP/pull/21), stacked on #18)
  adds three-actual-string `str_replace`, with empty-search handling and static/
  runtime type guards. No array forms, count references, coercion or regex.
  Catalog: 33 entries / 83 checks. Other GoHP stacks/normal checkout unchanged.
- All 16 string-list inputs matched actual string return type/value against
  PHP 8.3.36 and expectations: 7.61s (package 7.789s), empty stderr/exit 0.
  The first command selected no test due to a wrong name and is not PASS evidence;
  the correct selector ran once. Vet PASS. No code changes or reruns after PASS.
- Compiler local runtime/codegen/catalog tests and vet PASS; fixed native
  [CI 40119695048](https://github.com/katsuakikureyama/GoHP/actions/runs/40119695048)
  PASSed all 83 checks on Go 1.23.12 / actual PHP 8.3.36, clean checkout tree
  identical to the compiler pin. Exact commands/revisions and hashes are in the
  [execution record](evidence/newline-helper-2026-10-25.md), with the
  [bounded contract](newline-helper-unit.md).
- This is an additional independent helper, not the full `wpautop` paragraph
  renderer, regex/callback execution, an eleven-function bundle or core startup.
  Earlier standalone/bundle pins and evidence remain untouched. Draft only,
  no merge/deploy/release/license decision; local capability WebMCP-excluded.

## Previous: one combined helper executable

- Branch `codex/gohp-helper-bundle`, stacked on Draft #9. Test commit
  `577791ff93543d8b5abdc1efbc47af99f4bf407c` bundles the ten unchanged originals
  from two source files plus five owned composition drivers, translating and
  building them together into one native Go executable.
- New 22-case integration corpus: ten direct probes and twelve composed probes.
  Actual return types/values matched PHP 8.3.36 and authored expectations, empty
  stderr/exit 0. Initial run PASS 8.89s (package 9.096s); affected-package vet PASS.
- Same GoHP compiler/runtime `e3db685dfb2a3602f94bbfa9df105db3db96cc89` as
  zeroise; no compiler change/CI rerun. Its prior fixed-PHP CI PASS is reused.
  Existing four unit pins/corpora/evidence are untouched and not rerun; their
  73 observations are not claimed as a combined run.
- See [contract](helper-bundle-unit.md) and
  [execution evidence](evidence/helper-bundle-2026-10-25.md). This establishes
  only the new combined helper cases, not full files, core startup, HTTP/DB or
  theme/plugin support. Draft only, no merge/deploy/release/license decision.
  Local developer capability remains WebMCP-excluded; devboard unlinked.

## Previous addition: original zeroise unit

- `codex/gohp-zeroise-helper` is stacked on Draft #8. Test commit
  `ba30252d7509a23f2f8b97de8f0b4ac748947c2b` pins a fourth independent unit,
  extracting original `zeroise()` without changing WordPress source.
- Its compiler/runtime pin is `e3db685dfb2a3602f94bbfa9df105db3db96cc89`
  ([GoHP Draft #18](https://github.com/katsuakikureyama/GoHP/pull/18), on #16).
  Bounded `sprintf('%0Ns', string|int64)` only, canonical width 0–4096;
  invalid static/dynamic inputs fail explicitly. Catalog: 32 entries/78 checks.
- PHP 8.3.36 final-pin comparison PASS: 18 typed string outputs, 10.67s
  (package 10.886s), empty stderr/exit 0. Affected-package vet PASS.
  Compiler runtime/codegen/catalog local tests and vet PASS. Native compiler
  CI initially failed on an unsupported negative literal in the owned fixture;
  fixture/doc-only correction keeps parser scope unchanged. Corrected
  [CI 40118450143](https://github.com/katsuakikureyama/GoHP/actions/runs/40118450143)
  PASSed all 78 checks on fixed Go 1.23.12 / PHP 8.3.36; checkout tree matches pin.
- Prior three manifests retain their previous pin/results, not rerun. This is
  ten selected functions across four units and two pins, not a combined program.
  Separate GoHP #17 work and the normal checkout remain untouched.
- See [unit contract](zeroise-helper-unit.md) and
  [execution evidence](evidence/zeroise-helper-2026-10-25.md). Draft only;
  no merge/deploy/release/license decision or linked devboard Done.

## Current native-oracle policy (supersedes historical runner pins)

New local oracle runs admit only the reviewed official PHP 8.3.36 image in
`gossip/internal/phporacle`, with stable 8.3 >=8.3.36 checked before any fixture
is required in the same PHP process. Other digests are refused before Docker
starts. PHP 7.4, lower patches, prereleases and unreviewed newer series are not
admitted. This applies to hooks, all six standalone GoHP units and both the
ten-helper and integrated twelve-helper bundle experiments.
See [policy](php-oracle-policy.md) and the
[current unified-compiler execution record](evidence/unified-compiler-2026-10-25.md).

The earlier PHP 8.3.13 Gossip results and PHP 7.4/8.3.30 compiler CI below are
historical evidence only, not current-policy passes. Their files are retained
without rewriting observations. GoHP compiler/runtime pins remain distinct and
unchanged for each unit in the previous runner migration. The subsequent
compiler integration below backports only the policy from the separate #14
stack; other features from that stack are not imported.

At `2cac2bf34ce66ef788975e0f2e6644cb815d33b5`, current-policy local execution
PASSed with actual PHP 8.3.36: hook comparison (5 cases, twice each), six
constant-return types/values, slash helpers (13 inputs/26 typed returns), and
Boolean helper (23 inputs). Policy unit tests and affected-package vet also
PASS. Each unit keeps its own compiler pin, and no general compatibility claim
is inferred. Exact commands/hashes/timings for that migration are in the
[prior floor execution record](evidence/php-83-floor-2026-10-25.md).

## Current unified compiler

- The three pre-zeroise manifests pin `4e1b5abfe019ec121e13d85e99e51c206e765533`:
  GoHP Draft #16 stacked on #13, policy backport from `948647c` only. Active
  compiler target is 8.3 and native floor 8.3.36. Legacy-array/rtrim/Boolean
  features are retained; Composer inventory/strlen/array_keys/array_values are
  not imported. Catalog: 31 entries / 73 checks, only active target 8.3.
- Gossip test commit `e943cbce1f74c93b5ca990de7a544caabb274aa5` PASSed all
  three original-source units once on PHP 8.3.36: six constant-return types/values
  (13.37s), slash 13 inputs / 26 typed returns (8.64s), Boolean 23 inputs
  (12.00s), package 34.204s. Same compiler/runtime identity, three separately
  built units; not a single nine-function bundle or full WordPress execution.
- Only three Gossip compiler pins changed. Source, cases, expected values,
  observers, PHP admission and hooks are unchanged; earlier hook/admission
  PASS is reused under the unchanged effective configuration, not rerun.
- Compiler local internal/CLI tests, vet and shell syntax PASS. Native
  [CI 40116685783](https://github.com/katsuakikureyama/GoHP/actions/runs/40116685783)
  PASSed on fixed Go 1.23.12 / PHP 8.3.36. Exact evidence and limitations are
  in the unified-compiler execution record. Historical files remain untouched.

## Repository identity

- `katsuakikureyama/Gossip` is a GitHub fork of `WordPress/WordPress`, repository
  ID `1432300019`, with default branch `master`.
- `katsuakikureyama/Gossip-prototype` is the preserved independent prototype,
  repository ID `1432085879`; its Draft PR #2 and historical commits remain there.
- Fork work uses branch `codex/fork-baseline`, based on upstream commit
  `54106f6c545fc8b881f6b39b7963002bece6c28c` (`7.2-beta1-64693`).
- The follow-up `codex/gohp-return-helpers` branch is stacked on that unmerged
  baseline branch; its Draft PR must target `codex/fork-baseline`, not upstream.
- `codex/gohp-typed-observations` is the next stacked change, targeting
  `codex/gohp-return-helpers`. No default-branch merge is included.
- `codex/gohp-empty-array` is stacked on `codex/gohp-typed-observations` and
  must target that branch, not `master`, WordPress upstream or the independent
  prototype.
- `codex/gohp-slash-helpers` is stacked on `codex/gohp-empty-array` and must
  target that branch. Its Draft PR is not a default-branch merge.
- Local fork checkout: `/Users/kureyamayuri/Desktop/Gossip-wordpress`.
- `codex/gohp-boolean-helper` is stacked on `codex/gohp-slash-helpers`.
  [Draft PR #6](https://github.com/katsuakikureyama/Gossip/pull/6) targets #5.
- `codex/php-83-oracle-floor` is stacked on `codex/gohp-boolean-helper`;
  [Draft PR #7](https://github.com/katsuakikureyama/Gossip/pull/7) targets #6,
  not the default branch or WordPress upstream.
- `codex/unified-php83-compiler` is stacked on `codex/php-83-oracle-floor`;
  [Draft PR #8](https://github.com/katsuakikureyama/Gossip/pull/8) targets #7.
  The older `/Users/kureyamayuri/Desktop/Gossip` checkout still points to
  `Gossip-prototype`; it is not the new fork. No existing directory was removed.

## Historical implementation and verification (prior oracle policy)

- Boolean-helper translation PASS at `8901d634e4d52e9be06c4c2d7f91714652974b28`:
  GoHP `14bb01053947f1bf2fb32edaff8085feb7227f48` (Draft #13, based on #10)
  translates unchanged `wp_validate_boolean()` from functions.php 7932–7942.
  All 23 mixed inputs match PHP 8.3.13 in actual bool return type/value and
  authored expectations; empty stderr/exit 0. PASS 8.38s (package 8.557s).
  The initial attempt failed its PHP locale guard after generation/build;
  the observer and compiler native harness now select and verify `LC_CTYPE=C`
  in the oracle process. WordPress code was never changed. This is a bounded
  one-arm conditional/short-circuit/scalar validation contract, not full PHP.
  See [Boolean evidence](evidence/boolean-helper-2026-10-25.md) for exact pins,
  failed attempt, corrected execution, compiler checks and limitations.
- GoHP #13 final head `4fabe11d2063b03fced8d0aabd62bbd20a10fb88` passes
  [fixed PHP 7.4.33/8.3.30 CI 40113520430](https://github.com/katsuakikureyama/GoHP/actions/runs/40113520430).
  It differs from the Gossip pin only in native-test evidence logging, fixing
  strict summary parsing without changing compiler/runtime implementation.
  Prior CI failures remain recorded; they are not presented as overall PASS.
- Slash-helper translation PASS at `a97f14f42e6f59049062d59695c74da9d56ee0d2`:
  pinned GoHP `ceee0fc71e1e7947f1c78f1dd18875236b9a180b` (Draft GoHP PR #10,
  based on `feat/legacy-array-syntax`) translates the unchanged original
  `trailingslashit()` and `untrailingslashit()` snippets from
  `wp-includes/formatting.php` lines 2883–2885 and 2898–2900. The single
  affected opt-in test PASS took 18.82 seconds (package 19.018 seconds): PHP
  8.3.13 and generated Go matched all 13 explicit string inputs and their 26
  typed return values, with empty stderr and exit 0. The inputs include empty,
  slash/backslash-only, repeated suffixes, Windows-looking paths, whitespace,
  Japanese text, URL-looking text, and embedded NUL/LF bytes. This is a
  bounded `rtrim` contract only: two arguments and a literal byte mask `/\\`,
  not ranges, non-string coercion, dynamic masks, general path handling, or
  full WordPress formatting compatibility. See
  [slash-unit evidence](evidence/slash-helpers-2026-10-25.md).
- GoHP PR #10 CI also PASSed for fixed PHP 7.4.33 and 8.3.30 with Go 1.23.12:
  [run 40112678059](https://github.com/katsuakikureyama/GoHP/actions/runs/40112678059)
  records `diff.rtrim`, nine negative checks, and two runtime checks as PASS
  for each target, with artifacts saved. This is compiler-project evidence for
  the bounded `rtrim` implementation, not full WordPress compatibility; this
  fork did not rerun its own test in that CI job.
- Empty-array translation PASS at `1596c5ac35846db59612098fe13d4f5736f23655`:
  pinned GoHP `c67fbefe34b1b098c0a9f475504d06bd014de23c` (Draft GoHP PR #8,
  based on `feat/compatibility-views`) translates six unchanged original
  helpers, including `__return_empty_array()` with its original `array()`
  expression. The single opt-in test PASS took 13.14 seconds (package 13.328
  seconds): PHP and generated Go matched five scalar echo observations and all
  six typed values/types, with empty stderr and exit 0. The echo oracle remains
  deliberately limited to the five scalar functions because PHP array coercion
  is not type-preserving. The Go observer accepts only an empty `rt.Array` and
  does not claim general array equivalence. See
  [current v3 evidence](evidence/empty-array-2026-10-25.md).
- GoHP PR #8 CI also PASSed for fixed PHP 7.4.33 and 8.3.30 with Go 1.23.12:
  [run 40112122164](https://github.com/katsuakikureyama/GoHP/actions/runs/40112122164)
  records `diff.legacy_arrays` and seven negative checks as PASS for each
  target, with artifacts saved. This is compiler-project evidence, not a
  full WordPress compatibility result; this fork did not rerun its own test.
- Historical typed-return comparison PASS at `b0f0ce82d0532e1ff0258478b95a158c29b84efc`:
  five original helpers match PHP in actual scalar type/value, stdout, empty
  stderr and exit 0. A separate unchanged `__return_empty_array()` returns an
  empty array in PHP but was rejected by the old GoHP pin `9c7389d` (`array()`
  was unsupported); no generated source/report was left. This historical boundary
  is superseded only for the new pin above; the old run is not relabeled. See
  [v2 evidence](evidence/typed-helpers-2026-10-25.md).
- First original-source GoHP unit PASS at
  `1c84d0c736b682b81f3a267f4c5b2c6269461308`: five unchanged WordPress
  constant-return helpers were extracted, generated, hash-verified, compiled
  and compared against PHP. Stdout, empty stderr and exit 0 matched. The first
  attempt failed before translation due to a 16 MiB builder tmpfs; the corrected
  256 MiB configuration passed. See [evidence](evidence/return-helpers-2026-10-25.md).
  That earlier result was echo-only; the later typed probe above covers these
  five constants, not general return-type equivalence or full-file/core translation.
- Original WordPress files/history retained. Changes only add the nested
  `gossip/` Go experiment and documentation; no inherited runtime, asset or
  license file is modified.
- Native Plugin SDK and five-case fixture copied unchanged from prototype
  `1b9c05b0557bcca3f0a4f5eebf046cbef67d18e2`.
- The comparison now runs verified source bytes from **this fork**, not the
  WordPress installation bundled in a container image. A manifest pins the
  upstream commit, version and five source/license SHA-256 values.
- At `a167e6c8c56294f2d702e7c25342c614bde364dd`, race-enabled tests PASS:
  native SDK tests, source baseline, and five owned hook comparisons (two
  invocations each) against PHP 8.3.13. See
  [exact commands and exclusions](evidence/fork-hooks-2026-10-25.md).

## Unimplemented / unverified

- Full-file/core GoHP translation or execution, themes or plugins. The only
  translated exceptions are six tiny constant-return helpers at their
  separately pinned `c67fbefe34b1b098c0a9f475504d06bd014de23c` evidence and
  the two slash helpers at `ceee0fc71e1e7947f1c78f1dd18875236b9a180b`; they
  were not tested together as a larger unit. The inspected GoHP subsets do not
  cover the original hook engine; see [reuse boundary](reuse-boundary.md).
- General WordPress hook parity: dynamic registration/removal, references,
  recursion, multiple arguments, negative priorities and core filter stacks.
- A Go server serving WordPress, PHP/native plugin bridging, block/theme
  rendering, original database mapping, production auth and CGI execution.
- Full-site runtime or browser validation of this development snapshot. The
  prototype's old WordPress 6.6.2 tests are not evidence for this fork baseline.
- Stable release selection, upstream-source update automation, production
  migration, public deployment, license finalization or default-branch merge.

## Next steps / migration gates

Use the original WordPress source as the behavioral reference. The Boolean and
slash units are separate bounded contracts from the six constant-return helpers; do not
combine their evidence or extrapolate it to general arrays, general `rtrim`,
paths, or core translation. Select the next dependency-bearing unit only with
its own explicit contract. The nearby `send_nosniff_header()` needs `header()`
and an HTTP side-effect contract; it was inspected only, not runtime-tested or
supported. Keep source pins, isolated generation and per-unit differential
results. Do not grow the independent CMS or replace WordPress database
semantics with the prototype JSON store.

Production work remains gated on owner approval, source/version inventory,
restorable data copy, dry run, compatibility validation and rollback evidence.
This baseline creates no new browser or agent operations; local test tooling
is explicitly excluded from WebMCP publication.

Track the baseline in Draft PR #1 in **this fork**, based on `master`; stack the
helper-unit Draft PR #2 on `codex/fork-baseline` until that PR is merged; stack
typed observations on that helper branch, then stack the empty-array change on
typed observations, the slash-helper change on empty-array, and Boolean on slash.
No devboard card
is linked in this repository; none is marked Done by this report. Neither the
fork's default branch nor WordPress upstream is modified by the work branch.
