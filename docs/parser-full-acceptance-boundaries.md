# Complete-parser acceptance boundary review — 2026-10-30

## Decision

The fixed12-case/24-call result from [Gossip69](https://github.com/katsuakikureyama/Gossip/pull/69)
remains a **retained-reference comparison PASS**, not an arbitrary-input or
application acceptance result. The boundary review is now recorded; its open
items are not converted into exclusions that silently make G83-PARSER pass.
No compiler defect requiring a GoHP change was demonstrated by this review.

This is a source/evidence review, not a new execution or second verifier.
Reuse [the immutable generated execution](evidence/parser-full-generated-2026-10-30.md)
and [native capture](evidence/parser-full-baseline-2026-10-27.md). The existing
corpus, observers, pins and machine records are unchanged. The separate
[machine-readable boundary plan](../gossip/tests/gohp/parser-full-boundary-plan.json)
contains **planned, unexecuted probes**, not additional passed cases or an
executable extension of the current harness.

## What the current result establishes

| Dimension | Fixed cases / observation | Boundary that remains |
|---|---|---|
| Basic document structure | empty/freeform, single void/paired, nested/adjacent, leading/trailing freeform | A few complete paths, not every parser branch or arbitrary document |
| Block-comment recovery | `unclosed_opener`, `stray_closer` | Not JSON parse errors; all PCRE/JSON codes0 in these cases |
| Valid JSON attributes | `nested_attributes`: object/list, UTF-8 string, integer1, true/false/null | No integer extrema, duplicate/numeric-string keys or arbitrary nesting comparison in the complete parser; scalar roots cannot reach this JSON callsite |
| Invalid JSON attributes | `malformed_json_attributes`: `{"broken":}` produces null attrs and JSON4/message on both calls | Syntax error only; no complete-parser UTF-8/control/surrogate/depth error measurements |
| Raw string bytes | `raw_utf8_and_nul`: valid UTF-8 and NUL in freeform content | **Not invalid UTF-8 inside JSON attributes** |
| Parser state and reuse | same typed result and every parser field, unchanged caller, fresh/reused objects | JSON error is deliberately reset to0 before every call; arbitrary carry-in JSON/PCRE errors are not measured |
| Regex errors/resources | all24 recorded PCRE error codes0 | No original limit-error path, over-bound subject, heap/process-memory exhaustion, or timeout-equivalence measurement |
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
`no-more-tokens`; zero-error cases do not exercise that error path.

## Value and resource profile limits

The source/harness admits only the exact12 named, base64-bound strings, each
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
| Saved native | PHP8.3.36, Linux/aarch64, PCRE10.42, C LC_CTYPE, JIT off, backtrack100000/recursion10000, process memory128M | A PHP process memory limit is not a PCRE-only heap limit |
| Generated candidate | Darwin/arm64, Go1.23.0, static PCRE10.47_1, JIT/UTF-UCP off, match100000/depth10000/heap131072KiB/subject4096bytes | Equal nominal counters do not prove equal failure points across engine versions or total-process memory behavior |

The saved native and generated stdout envelopes have different metadata and
lengths. The comparison is the canonical typed `cases` value, not raw whole
process-output equality. Compiler/linker warnings, runtime stderr, test/harness
failure, parser errors and observation-transport failure are distinct states.

## Remaining gates and next probes

1. **Fresh fixed-PHP comparison — unverified.** Use the existing `fresh` mode,
   exact cached PHP image and unchanged original corpus/observer. The native
   runner must already be available; this review neither starts Docker nor
   installs an image/engine or creates remote CI resources. Retained mode PASS
   must not be relabeled as fresh. SKIP or unavailable environment is not PASS.
2. **Supported-profile JSON/type boundaries — unmeasured on the full parser.**
   A separate series must cover ordered duplicate/numeric keys, integer bounds,
   encoding/control/surrogate/depth errors at the actual brace-bound callsite.
   Each proposed case needs a complete original PHP observation, generated
   observation, exact profiles, all before/after fields and typed comparison.
3. **Original regex-error/resource path — unmeasured.** Select a reproducible
   in-profile input/error budget and observe the actual original parser's output,
   fields and PCRE status. Do not substitute direct preg_match or different
   engine experiments. Host guards/timeouts need a separate failure observation,
   not disappearance from the denominator or a matching nonzero process exit.
4. **Carry-in state — unmeasured.** In a new observer series, deliberately seed
   JSON and PCRE errors, then record both sides before/after original parse calls
   without the current JSON success baseline. Do not assume parse resets globals.
5. **Outside-profile API/value cases — explicit, not supported.** Nonstring
   input, float/overflow values, arbitrary public-method calls, references,
   visibility/dynamic properties and unbounded resource use remain excluded from
   the current result. Proposed negative records must preserve those rows; they
   must not replace the fixed12 original cases or imply PHP exception parity.

The plan's case families are review findings, **not newly captured native
expectations**. Keep the original12/24 result and its hashes unchanged. Each
extension gets a new manifest/corpus/observer/evidence identity, separate
denominator and supported/unsupported/unmeasured status. No expected values are
invented from generated output. G83-PARSER remains open; bootstrap/DB/HTTP/site
gates remain not-run.

## Capability and work performed

This change adds only the boundary plan and project documentation. The plan is
not an input consumed by `TestOriginalWordPressFullParserGenerated`; the current
harness still executes exactly its immutable12 cases. Existing local operation
`gossip.tests.original-parser-compare` remains experimental and WebMCP-excluded,
with unchanged retained/fresh schemas, explicit opt-in and fresh evidence output.
No new executable capability, public API, auth route or deployment is added.

No retained/fresh test, generation, build, native/Go execution or artifact
download was performed for this review. Prior unavailable Docker state is reused,
not repeatedly probed. Only changed JSON/Markdown validity, link targets, source
preservation and document diffs are checked. Root current-state/INDEX and Issue1
record the review separately from measurements; devboard remains unidentified.
