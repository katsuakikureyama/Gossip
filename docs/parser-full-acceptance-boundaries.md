# Complete-parser acceptance boundary review — 2026-10-30

## Decision

The fixed12-case/24-call result from [Gossip69](https://github.com/katsuakikureyama/Gossip/pull/69)
remains a retained-reference PASS. A separately measured
[fresh-PHP comparison](evidence/parser-full-fresh-2026-10-30.md) is now adopted
as a **case-limited fresh comparison PASS**, not arbitrary-input or application
acceptance. The boundary review is now recorded; its open
items are not converted into exclusions that silently make G83-PARSER pass.
No compiler defect requiring a GoHP change was demonstrated by this review.

A later, separate [carry-in-state measurement](evidence/parser-carry-state-2026-10-30.md)
now passes four sequences/twenty original calls per backend. JSON0/4 and incoming
PCRE0/2 match before/after original calls without per-call success resets.
Its separate [JSON key/type series](evidence/parser-json-key-types-2026-10-30.md)
now also passes twelve new documents/twenty-four original calls per backend:
duplicate/numeric keys/order, signed64 bounds, nested types and decoded bytes.
A separate [JSON-error series](evidence/parser-json-errors-2026-10-30.md) now
passes twelve documents/twenty-four original calls per backend: invalid attribute
UTF-8, raw controls, unpaired UTF-16 escapes and default-depth null/error paths.
The [resource series](evidence/parser-resource-boundaries-2026-10-30.md) adds
twelve fixed-limit stress documents/twenty-four calls per backend that agree
with all PCRE0. A separately named PHP-only match1 reference observes original
error2 handling; the generated entry has no matching budget selector and is
not compared under that setting. Resource gates are partially observed and
still open, as is G83-PARSER. These extensions do not alter the fixed12-case
result described below.

The original PR70 boundary review itself is a source/evidence review, not a new
execution or second verifier. The later carry-in series has its own measured
source, observers, command and evidence.
Reuse [the immutable generated execution](evidence/parser-full-generated-2026-10-30.md)
and [native capture](evidence/parser-full-baseline-2026-10-27.md). The existing
corpus, observers, pins and machine records are unchanged. The separate
[machine-readable boundary plan](../gossip/tests/gohp/parser-full-boundary-plan.json)
separates passed fresh-v1, carry-in, key/type and JSON-error identities from **remaining planned,
unexecuted probes**. It is not an executable extension of either harness.

## What the current result establishes

| Dimension | Fixed cases / observation | Boundary that remains |
|---|---|---|
| Basic document structure | empty/freeform, single void/paired, nested/adjacent, leading/trailing freeform | A few complete paths, not every parser branch or arbitrary document |
| Block-comment recovery | `unclosed_opener`, `stray_closer` | Not JSON parse errors; all PCRE/JSON codes0 in these cases |
| Valid JSON attributes | Fixed-v1 `nested_attributes`, plus separate key/type12/24 with signed64 extrema, duplicate/escaped/numeric-string keys, ordered nested values and decoded bytes | Case-limited supported profile, not arbitrary nesting, float/overflow values or scalar-root JSON; scalar roots cannot reach this callsite |
| Invalid JSON attributes | Fixed-v1 syntax4 plus separate JSON-error12/24: invalid attribute UTF-8/controls/unpaired surrogates/default-depth produce null attrs and codes5/3/10/1 with matching messages | Case-limited errors, not arbitrary malformed input, successful deep output or exact failure-threshold/resource equivalence |
| Raw string bytes | Fixed-v1 valid UTF-8/NUL freeform; separate JSON-error corpus observes invalid UTF-8 and raw NUL/LF/TAB inside attributes | The separate denominators remain distinct; fixed-v1 freeform alone did not establish attribute errors |
| Parser state and reuse | fixed-v1 has same typed result and every parser field, unchanged caller, fresh/reused objects | Fixed-v1 resets JSON before every call. The separate carry-in series now measures JSON0/4 and incoming PCRE0/2; arbitrary globals remain unmeasured |
| Regex errors/resources | Fixed-profile stress12/24 has PCRE0, including ten4096-byte inputs; separate native-only match1 observes16original error2/freeform calls | No generated match1 support/comparison, default-profile original nonzero error witness, over-bound subject, heap/process-memory failure or timeout-equivalence measurement |
| Entry API | `parse(document)` is always passed a string decoded from the fixed corpus | No null/bool/int/array/object coercion or PHP warning/exception equivalence; the other six public methods are not directly API-tested |

Case IDs and contents are fixed in
[the original corpus](../gossip/tests/gohp/parser-full-reference-cases.json).
The PHP observer resets JSON at lines89/96 and the Go observer calls
`fullParserResetJSON` before each observation. Therefore `reuse_reset=true`
does **not** mean the original parser clears PHP-global error state itself.
That caller-established baseline must remain visible in every description of
the result. No PCRE state is manually reset by this observer, but a sequence
of all-zero PCRE cases still says nothing about a prior nonzero error.

The [original next_token pattern](../wp-includes/class-wp-block-parser.php)
captures attributes only inside braces. A comment containing a scalar/list
instead of an attribute object is a tokenizer/input-recovery question, **not a
scalar-root json_decode call in this parser**. Generic JSON scalar-root support
must not be made a new G83 requirement just because the builtin supports it.
The same original method has an explicit preg_match-false path returning
`no-more-tokens`; fixed-profile zero-error cases do not exercise that path.
The separate native-only match1 reference now observes it under its declared
different budget; it is not a matching generated-profile execution.

## Value and resource profile limits

The fixed-v1 source/harness admits only its exact12 named, base64-bound strings, each
no more than4096bytes. Both typed observers reject unsupported value types
and observation nesting over64. A float attribute is not a successful PHP
rejection case: the native observation transport itself cannot encode floats.
GoHP138 separately rejects valid fraction/exponent/out-of-int64 numeric tokens.
Do not add float cases to the existing observer and treat its failure as a
parser verdict. A versioned negative observation must separate PHP's value,
the GoHP profile rejection and transport failure without coercing any of them.

Likewise a very deeply nested successful result can hit the observer limit
before the JSON512 or parser resource boundaries. A JSON depth-error probe
must record the native null/error result; observation overflow is not a JSON
error PASS. Existing standalone GoHP JSON/PCRE fixtures are feature evidence,
not the same complete-parser/source/engine/observer configuration.

| Runtime side | Actual measured profile | Non-equivalence to retain |
|---|---|---|
| Saved native and byte-identical fresh capture | PHP8.3.36, Linux/aarch64, PCRE10.42, C LC_CTYPE, JIT off, backtrack100000/recursion10000, process memory128M | A PHP process memory limit is not a PCRE-only heap limit |
| Generated candidate | Darwin/arm64, Go1.23.0, static PCRE10.47_1, JIT/UTF-UCP off, match100000/depth10000/heap131072KiB/subject4096bytes | Equal nominal counters do not prove equal failure points across engine versions or total-process memory behavior |

The saved native and generated stdout envelopes have different metadata and
lengths. The comparison is the canonical typed `cases` value, not raw whole
process-output equality. Compiler/linker warnings, runtime stderr, test/harness
failure, parser errors and observation-transport failure are distinct states.

## Remaining gates and next probes

1. **Fresh fixed-PHP comparison — passed, case-limited.** The existing `fresh`
   mode ran once at clean1f7a3b2 with the exact already-cached image and unchanged
   corpus/observer; this review adopted its handoff without rerunning. No daemon,
   image/engine or remote resources were installed. The old retained result is
   unchanged; a separate machine record binds the new actual PHP execution.
2. **Supported-profile JSON/key/type — passed, case-limited separate series.**
   The [key/type12/24 measurement](evidence/parser-json-key-types-2026-10-30.md)
   covers duplicate/escaped/numeric keys/order, signed64 min/max, empty/nested
   arrays/objects and scalar/string bytes at the actual brace-bound callsite.
   Unchanged full observers and a new pinned named-block/attrs-type contract
   distinguish decoding from input recovery; all typed values/state/errors match
   actual fixed PHP.
3. **JSON-error boundaries — passed, case-limited separate series.** A separate
   [JSON-error12/24 measurement](evidence/parser-json-errors-2026-10-30.md)
   now covers invalid encoding/raw control/lone-surrogate/default-depth errors
   at original283, with null attrs and matching codes/messages5/3/10/1.
   It reuses the exact observers and their JSON success baseline. Valid escaped
   controls/surrogate pairs and invalid forms remain separate case families;
   successful deep output, arbitrary malformed inputs and resource equivalence
   are not established.
4. **Original regex-error/resource path — partially observed, still open.**
   The new finite stress series compares twelve documents/twenty-four original
   calls at unchanged limits, all PCRE0. A separate named match1 PHP reference
   records eight cases/sixteen original error2/freeform calls, with full fields
   and raw bytes. Generated execution under match1 is explicitly unsupported/
   not-run, not a cross-runtime PASS or a defect in the fixed profile. No PHP
   source change or direct-regex seed substitutes for the original path.
   A future generated failure observation needs either an actual current-profile
   witness or a separately reviewed explicit profile contract that preserves
   the default. Host guards/timeouts still need separate failure observations,
   not disappearance from the denominator or matching nonzero process exits.
   Finite zero-error samples do not prove all errors unreachable; PHP process
   memory, PCRE heap and nominal counters remain distinct.
5. **Carry-in state — passed, case-limited separate series.** The new
   [four-sequence measurement](evidence/parser-carry-state-2026-10-30.md) seeds
   actual JSON0/4 and PCRE0/2, then records before/after twenty original parse
   calls per backend without per-call resets. JSON persists until a decode
   replaces it; original regex execution clears PCRE independently. All typed
   values/state/caller/reuse match. Other globals/error codes remain unmeasured;
   a direct PCRE caller seed is not the original regex resource-error path.
6. **Outside-profile API/value cases — explicit, not supported.** Nonstring
   input, float/overflow values, arbitrary public-method calls, references,
   visibility/dynamic properties and unbounded resource use remain excluded from
   the current result. Proposed negative records must preserve those rows; they
   must not replace the fixed12 original cases or imply PHP exception parity.

The plan's remaining unmeasured families are review findings, **not captured
native expectations**. Keep the original12/24 result and its hashes unchanged. Each
extension gets a new manifest/corpus/observer/evidence identity, separate
denominator and supported/unsupported/unmeasured status. No expected values are
invented from generated output. G83-PARSER remains open; bootstrap/DB/HTTP/site
gates remain not-run.

## Original PR70 review capability and work performed

This change adds the boundary plan, project documentation and adopted fresh
evidence. The plan is
not an input consumed by `TestOriginalWordPressFullParserGenerated`; the current
harness still executes exactly its immutable12 cases. Existing local operation
`gossip.tests.original-parser-compare` remains experimental and WebMCP-excluded,
with unchanged retained/fresh schemas, explicit opt-in and fresh evidence output.
No new executable capability, public API, auth route or deployment is added.

No retained/fresh test, generation, build, native/Go execution or artifact
download was repeated by this review. A separate owner's one successful fresh
configuration was adopted with source/run provenance after its existing runner
became available; it was not silently treated as the old retained measurement.
Only changed JSON/Markdown validity, link targets, source preservation and document
diffs are checked locally. Root current-state/INDEX and Issue1
record the review separately from measurements; devboard remains unidentified.
