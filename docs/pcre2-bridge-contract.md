# PCRE2 bridge and bounded PHP output contract

Implementation update, 2026-10-27: GoHP implemented bounded B1/B2/B3 in
[PR #110](https://github.com/katsuakikureyama/GoHP/pull/110),
[#112](https://github.com/katsuakikureyama/GoHP/pull/112) and
[#114](https://github.com/katsuakikureyama/GoHP/pull/114). The
[isolated Gossip consumer](parser-preg83-generated.md) pins #114, with 23 selected
fixed-limit cases and one explicitly excluded limit=1 case. General references,
the original next_token, PHP8.4 regex and native distribution remain outside scope.
The design-only baseline below is retained as historical requirements, not a
current claim that those three GoHP stages are still unimplemented.

Status: **design / implementation handoff only**, 2026-10-27 JST. No cgo bridge,
PHP builtin admission, new measurement, dependency adoption or GoHP code change
is introduced by this document. It specifies the next bounded work, not an
implemented API. Gossip owns consumer evidence; GoHP owns the generic runtime,
compiler and storage integration. WordPress names/pattern hashes must not become
compiler special cases.

## Evidence and fixed audit baseline

The [native reference](parser-token-reference.md) covers 24 PHP8.3.36 / PCRE10.42
cases. The [independent C probe](parser-pcre2-comparison.md) reproduces their
direct-regex outputs with PCRE2 10.47. The [regexp2 string profile](parser-regexp2-comparison.md)
has 22 equal outputs, one rune-boundary difference and one non-comparable limit.
These results justify PCRE2 as the next bridge candidate, not full PHP equivalence.
No successful native comparison needs rerunning for this documentation change.

GoHP was inspected read-only at
[`fd6936a4570e0f9f0bf522f8cb1f8802a6f70533`](https://github.com/katsuakikureyama/GoHP/tree/fd6936a4570e0f9f0bf522f8cb1f8802a6f70533)
(#107). The independently owned [#108](https://github.com/katsuakikureyama/GoHP/pull/108)
Slot-array work was in flight at audit time; do not infer its merge or behavior
from this snapshot, modify its branch, or duplicate its storage implementation.

| Existing seam at that commit | Consequence for this work |
| --- | --- |
| [IR call args are value expressions](https://github.com/katsuakikureyama/GoHP/blob/fd6936a4570e0f9f0bf522f8cb1f8802a6f70533/internal/ir/ir.go#L6); [call lowering reads values](https://github.com/katsuakikureyama/GoHP/blob/fd6936a4570e0f9f0bf522f8cb1f8802a6f70533/internal/codegen/generate.go#L492) | No builtin output-destination or reference mode exists. Standard PHP passes `$matches`, not a call-site `&$matches`; accepting the syntax alone does not implement its mutation. |
| [Generated scalar-frame / map split](https://github.com/katsuakikureyama/GoHP/blob/fd6936a4570e0f9f0bf522f8cb1f8802a6f70533/internal/codegen/generate.go#L185) | Scalar Frames cannot receive capture arrays. Array locals currently use detached-value map-root replacement. |
| [Ordered nested runtime.Array](https://github.com/katsuakikureyama/GoHP/blob/fd6936a4570e0f9f0bf522f8cb1f8802a6f70533/runtime/runtime.go#L14) | Numeric/string keys and `[bytes, offset]` values can be represented as values, but this is not a reference-aware writable location. |
| [Owned Slot arrays](https://github.com/katsuakikureyama/GoHP/blob/fd6936a4570e0f9f0bf522f8cb1f8802a6f70533/docs/owned-array-slot-values.md#L33) / [scalar-only elements](https://github.com/katsuakikureyama/GoHP/blob/fd6936a4570e0f9f0bf522f8cb1f8802a6f70533/internal/storage/scalar_array.go#L126) | A standalone Slot can own a scalar-element array; Frame locals and array elements remain scalar-only. Nested capture arrays cannot simply be inserted into this storage model. |
| [Execution state](https://github.com/katsuakikureyama/GoHP/blob/fd6936a4570e0f9f0bf522f8cb1f8802a6f70533/runtime/execution.go#L3) | Mutable state is per logical execution; only strtok state exists. No preg last-error state exists. |
| [Generated runtime import](https://github.com/katsuakikureyama/GoHP/blob/fd6936a4570e0f9f0bf522f8cb1f8802a6f70533/internal/codegen/generate.go#L106), [check script](https://github.com/katsuakikureyama/GoHP/blob/fd6936a4570e0f9f0bf522f8cb1f8802a6f70533/scripts/check.sh#L1), [Dockerfile](https://github.com/katsuakikureyama/GoHP/blob/fd6936a4570e0f9f0bf522f8cb1f8802a6f70533/Dockerfile#L1) | Generated code imports the pinned runtime; current testing/build configuration explicitly disables CGO. Adding a mandatory library to that path would change the default contract. |

The [reference/value roadmap](https://github.com/katsuakikureyama/GoHP/blob/fd6936a4570e0f9f0bf522f8cb1f8802a6f70533/docs/reference-value-model.md#L119)
remains authoritative for storage, evaluation and signature-aware calls. Internal
Slot tests or the new regex bridge must not be relabeled as R2/R4 completion.

## Separation and implementation order

```text
GoHP semantic admission + explicit local-output IR       [B3: not implemented]
              |
PHP wrapper: ordered result + per-Execution error state   [B2: not implemented]
              |
optional byte-only Go/C PCRE2 bridge                     [B1: next implementation]
              |
pinned PCRE2 8-bit library                               [C-probe evidence only]
```

**B1 is the immediate next change-set:** an opt-in engine-only bridge and its
tests in GoHP, with no compiler acceptance of `preg_match`. It can be built and
checked independently of unfinished nested storage. B2/B3 must not be smuggled
into B1 through a WordPress-specific helper or a second array/reference engine.
Each later stage has its own evidence and can fail without rewriting this history.

## B1: engine-only bridge boundary

Proposed interface data (not existing Go symbols): input pattern body bytes,
subject bytes, normalized nonnegative byte start, explicit compile options and
limits; output raw engine return code plus ordered numbered captures, names,
presence and byte spans. No PHP values, output locations, execution state or
JSON decoding belong in this layer. Unsupported errors remain errors, not no-match.

- Use PCRE2's **8-bit** interface, explicit lengths and whole subjects. Preserve
  NUL/invalid UTF-8; never use NUL-terminated C strings for payload lengths or
  normalize/slice the subject to implement starting offsets.
- First profile: DOTALL, LF newline, BSR_UNICODE, C tables, no UTF/UCP/JIT;
  limits match100000/1, depth10000, heap131072KiB, matching the recorded candidate.
  PHP's process memory limit is not the same as the engine heap bound.
- First bridge implementation is synchronous and stateless: no compiled-pattern
  cache, C callback into Go, retained Go pointer or implicit global last error.
  Copy inputs into C-owned buffers for the call, including valid empty-buffer
  handling; copy output bytes/metadata into Go-owned values before freeing C
  allocations, compiled code, match context and match data on every path.
- Enforce size/offset conversions and output bounds before exposing results;
  preserve unmatched versus zero-length matched captures and raw failure codes.
  No exposed pointer may outlive its documented owner. See the official
  [cgo pointer rules](https://pkg.go.dev/cmd/cgo#hdr-Passing_pointers).
- A Go context deadline does **not** interrupt an arbitrary synchronous C match.
  Use PCRE engine limits; keep test-only process deadlines as an outer harness
  guard, never report a timeout/killed process as PHP false or a matching error.
  Runtime cancellation requires a separately designed mechanism if later needed.

Build isolation proposal: an explicit tag such as `gohp_pcre2` plus `cgo` selects
the implementation; the disabled build returns a typed backend-unavailable
condition. Names are provisional. Do not silently fall back to RE2, regexp2, PHP,
a handwritten tokenizer or fabricated outputs. Default CGO=0 builds remain
unchanged and do not require headers/libraries. No new flag/tag is available yet.

First measured platform candidate is darwin/arm64 with the existing pinned
PCRE2 10.47 header/archive. Recheck the actual library identity in the cgo build:
the previous standalone C executable does not prove cgo linkage. Linux CI needs
its own pinned source/image/compiler/archive and native comparison before any
platform support claim. Keep the existing pure-Go CI lane; add an optional,
separate backend lane only with an explicitly prepared image. No surprise apt/brew
install, unpinned system library, binary publication or default static-link policy.

## B2: PHP wrapper, separate from the engine

The proposed first PHP-facing domain is a declared experimental target8.3
profile, not a promise about every PHP8.3 patch or PHP8.4. Pattern is a literal
slash-delimited byte string with `/s`; subject is an actual byte string; flags
are exactly PREG_OFFSET_CAPTURE; offset is an actual integer in
[-subject-byte-length, subject-byte-length]. No coercions or other preg functions.
Patterns are generic PHP syntax, not a WordPress name/hash whitelist. Invalid
pattern syntax and unmeasured errors are explicit unsupported/runtime diagnostics,
not fabricated PHP warnings/false. Syntax compilation and source publication
status must remain separate from a compatibility result.

For the measured outcomes, the wrapper produces:

| Engine outcome | PHP return | New matches value | Per-execution error |
| --- | --- | --- | --- |
| Match | integer 1 | Ordered PHP array: group0, then each named key before its numeric key; nested `[raw bytes, byte offset]`; unmatched middle pair empty/-1; trailing unmatched entries absent | 0 |
| No match | integer 0 | Empty array, replacing prior null for the admitted fixture | 0 |
| Match-limit failure | Boolean false | Empty array, replacing prior null for the admitted fixture | 2 |
| Unavailable backend, bad profile, malformed syntax, unknown engine failure | No compatible result | No compatible mutation claim | Distinct failure, never pretend it is PHP error2 |

The matches array must be fresh and owned, not a borrowed C view. Record raw
engine codes alongside translated PHP errors in evidence. PHP error messages,
general warning/exception behavior, other PCRE error mappings and arbitrary flags
need their own native evidence. See [preg_match](https://www.php.net/manual/en/function.preg-match.php)
and [preg_last_error](https://www.php.net/manual/en/function.preg-last-error.php).

Add last-error state only to the existing explicit Execution owner, propagated
through every admitted direct/nested call path by effect/signature analysis.
Success/no-match reset the state; measured failure sets it. Never use a package
global, goroutine-local surrogate or C-thread-global PHP state. Concurrent use
of one Execution remains outside its existing contract. Sequential failure→success,
two isolated executions and nested calls need independent regression evidence
before exposing `preg_last_error()`; current single-call captures are insufficient.

## B3: local output, not general reference support

Proposed first source shape:

```php
$matches = null;
$has_match = preg_match('/example/s', $subject, $matches,
                        PREG_OFFSET_CAPTURE, $offset);
```

The pattern is illustrative; the contract requires a literal expression, not
arbitrary runtime pattern construction. Initial admission requires an initialized, direct
local `$matches`; result destination must be a distinct direct local. All value
arguments must be proven side-effect-free reads/literals, with no read/output
aliasing. Original direct reads of known public data properties are a candidate
extension of that rule, not magic getters. Do not claim PHP's general argument
evaluation order based on Go's call evaluation order.

Calls need explicit signature/output metadata and validated writable-local IR;
do not disguise the third argument as an ordinary copied value expression, a raw
Go pointer, or call-site `&` syntax. Validate source **and** externally supplied
IR. A malformed destination or unsupported form must be rejected before emitting
an apparently supported call. Expression nesting, outputs in properties/elements,
dynamic locations, references/aliasing, variadics, callbacks and result/output
destination aliasing remain rejected until separately designed and compared.

Storage gate: the existing nested runtime.Array is a possible **value** for the
closed no-alias subset, but scalar Frames/ScalarArray cannot own it. A reviewed
whole-local publication through the established generated-local facade, or the
canonical nested-storage integration, must be completed and verified first.
Do not insert an `any`/nested value into scalar-only Slots, expose raw payloads,
copy/mutate/write back a shared array, or introduce an alternative cell engine.
The ordinary detached-value path may only be used while no references can be
admitted/hidden, with an explicit migration boundary owned by the GoHP reference
model. This document does not choose or implement that outstanding location API.

Keep result construction, output publication, return assignment and execution
state updates explicit and single-evaluation. Verify observable ordering using
the exact admitted no-alias form; do not broaden to effectful forms by inference.
The bounded builtin output must not advertise general by-reference parameters,
array COW, R2 or R4 support.

## Acceptance gates and evidence identities

| Gate | Required evidence before advancing | Status now |
| --- | --- | --- |
| B1 bridge | Same24 raw byte/capture/error cases through actual Go→C bridge; default-disabled behavior; memory ownership/empty/NUL/lifetime and conversion boundaries; exact linked library/toolchain/profile | Not implemented |
| B2 PHP wrapper | Typed ordered capture arrays and 1/0/false distinctions; explicit ownership; Execution reset/isolation/nested-call contract with fresh native cases | Not implemented |
| B3 generated call | PHP source→validated IR→Go→native PHP comparison; local output publication visible to caller; unsupported source/malformed-IR negatives; actual generated build with enabled backend | Blocked on output-location contract and B1/B2 |
| Original next_token | Unchanged selected original method, real dependencies, list destructuring/JSON/capture reads and remaining compiler gaps resolved and compared | Not established by regex matching |
| Complete parser | Original parse/stack/tokenizer dependency closure and full-method cases | Not implemented |

Attach compiler/runtime commit, target label, exact PHP patch/image/PCRE profile,
backend source/version/header/library hashes, Go/C toolchains, platform/ABI,
build tag/link mode, corpus/driver hashes and measured phases to each new record.
Backend identity must not be inferred from target8.3 or a system-installed name.
Existing evidence remains pinned; new GoHP main does not silently repin Gossip.
8.4 requires its own runner and evidence before admission, not copying 8.3 results.

Build/link/platform support, notices and license/distribution policy remain
separate gates. GoHP's project license is still undecided at the audit baseline;
this consumer document cannot settle it. The prior installed-library license
inspection is not a release authorization or a legal conclusion.

No user-facing capability/route/API is added here. Future local test tooling must
retain explicit UI/WebMCP/MCP exclusion; any new executable operation gets its
own capability record. No production target or deployment exists in this change.
