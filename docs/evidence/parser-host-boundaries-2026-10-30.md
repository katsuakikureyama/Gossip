# Original parser host boundaries — 2026-10-30

## Observed result

One new configuration **PASSes**, test2.65s/package2.900s, at clean Gossip
`ea45b88c8d3e2a2f2420dc98b7b6f16301c4fcab`, tree
`f60a0f9bddfc8901a22615645e4fcb76c7980932`. The
[machine record](parser-host-boundaries-2026-10-30.json) preserves all actual
observations, source/input hashes, exact sanitized command/environment and
run_count1. SHA256
`a26c81e17833e67b5af416d1d0fd999d7cf44af7f963d228ecf50fc7d3f85710`
(36633bytes); raw record SHA256
`b7a285e7bf7cff5af06cc088e335609780cf2442dd0ef08dc5f2d7c0e5a982d5`.

The [four separate fixed inputs](../../gossip/tests/gohp/parser-host-boundaries-cases.json)
exercise the actual original generated `parse` method:

| Input | Original call | Following empty call on the same instance |
| --- | --- | --- |
| 4096bytes, alternating NUL/invalid-UTF8 byte | Completes; freeform innerHTML preserves all bytes; PCRE becomes0 | Completes; PCRE0 |
| 4097 ASCII bytes | Explicit unsupported-subject panic; PCRE2/JSON4 preserved | Completes; PCRE0 |
| 4097 UTF8 bytes, only1367 characters | Same unsupported-subject boundary; limit is bytes | Completes; PCRE0 |
| 4097 binary bytes | Same unsupported-subject boundary, without byte coercion | Completes; PCRE0 |

The guard process completes8 original calls:5 completed calls and3 rejected
calls. Four direct caller regex seeds establish PCRE2 and are counted
separately; they are not original-pattern error evidence. JSON syntax4 is also
caller-established and remains4 throughout these calls. Caller strings are
unchanged. The unsupported wrapper panic is `*fmt.wrapError`, with message
`experimental preg83: input outside experimental preg83 profile: subject size or byte offset outside bounds`.
The observer recovers only around the actual original call to record this
explicit boundary, not to publish a fabricated PHP return value.

Parser fields are **partially written before rejection**: original `parse`
sets document/offset/output/stack before its regex call. Rejected calls retain
the new4097-byte document, offset0 and empty output/stack. They do not preserve
the constructor's four null properties. Following empty calls clear the parser
document/arrays normally. Execution-owned PCRE/JSON error preservation and
parser-state preservation are different contracts.

## Host deadline

The separate workload first completes one actual original parse of the new
4096-byte binary input and emits a small readiness record. Only then does the
host arm100ms; the child continuously executes further original parse calls,
without sleeping or waiting on stdin. The observed host interval is
101410333ns. Cancellation terminates the child with SIGKILL/exit-1,
`Wait` completes, runtime stderr is empty and no trailing observation is
published. The result is `host-timeout`, `complete_parse_observation=false`,
`php_error_result=not-produced`.

The repeat-call count and parser/error state at kill are unknown. This proves
the launched workload can be cancelled and reaped; it does **not** establish
interrupting one stalled PCRE2 call, a timeout result inside PHP, an equal
resource-failure point, or a process-memory/heap/depth contract. No successful
parse or native equivalence result is inferred from the killed process.

## Reused source and limits

The [new pin](../../gossip/tests/gohp/parser-host-boundaries.json) binds the
[new observer](../../gossip/tests/gohp/parser-host-boundaries-observer.go.txt)
and [new harness](../../gossip/tests/gohp/parser_host_boundaries_test.go).
They reuse the byte-identical retained generated source/report from
[Gossip76](parser-match1-generated-2026-10-30.md), pinned GoHP
`e6a71a5f267658e8c575ffab5db6712246e543b2`/`pcre2-php83-match1-v1`.
No generation, previous12-case comparison, PHP or Docker was rerun.

| Retained identity | SHA256 |
| --- | --- |
| Generated complete3class parser | `4c8648ef4b63f7862444b52c0c0f97c0898fc40bbe4f16698014ae7447f89842` |
| Build report | `ec3714aa49e36ecce45bacd9a0458c951833090d900f47b0a4f8932f4c54f3f6` |
| Earlier successful generated record | `fc5a4cc7467b5b6b1e5dbca119f9f9ba6850075573f5507d5586ee8133fdf794` |
| Pinned GoHP source archive | `1f65e066c382932f5d63bd2ad5a1a1cb1ec8bba10e16e1c294f4b21135234a03` |

The new build uses the existing cached Go1.23.0 Darwin/arm64, verified clang,
SDK15.5/staticPCRE10.47_1, CGO1/taggohp_pcre2, Match1/Depth10000/
Heap131072KiB/Subject4096. Parent test uses CGO0/offline/shared cache. The
temporary candidate binary is test-cleaned; retained parser.go/report.json
remain untouched. Linker duplicate-library warning is separate from runtime
stderr0. No new engine/dependency/image/daemon/CI resources or deployment.
Original WordPress and all previous corpora/observers/pins/evidence remain
unchanged. The resource gate remains partially-observed and G83-PARSER,
BOOT/HTTP/full WordPress acceptance remain open.

Static review checked8pin-file bindings plus2additional evidence bindings,
raw-observation equality, all4measured inputs unchanged, the11-file ownership
boundary, preservation of every previous plan field/gate, complete WordPress
non-overlay byte equality,5Markdown315relative links and whitespace: PASS.
Only documentation/evidence/status is added after the measured code commit;
successful runtime configurations are not repeated. Gossip has no configured
CI; no workflow or manual remote check was added.

## Capability

| Field | Contract |
| --- | --- |
| Stable ID | `gossip.tests.parser-host-boundaries` |
| Service/resource | Local developer inspection of the original generated parser's host boundaries |
| Input schema | `GOSSIP_GOHP_PARSER_HOST_BOUNDARIES=1`; explicit absolute GoHP source/PCRE prefix/retained parser+report directory, fresh evidence file; immutable4case pin; fixed100ms after readiness |
| Inspect / prepare | Verify original5source, retained source/report/proof/pin/profile, new cases/observer and existing toolchain; export pinned source and build a temporary candidate |
| Save / commit | Exclusive `GOSSIP_GOHP_HOST_EVIDENCE_OUT`; no overwrite of retained artifacts, binary test-cleaned; no application save/publish/send operation |
| Status / common result | Experimental/host-boundaries-observed; completed/unsupported-subject/host-timeout remain distinct; skip/build/observer/infrastructure failure are not compatible PHP results |
| WebMCP | `excluded`: local test, no browser/server/UI operation, auth or route change |

From `gossip/`, using the explicit environment recorded in the machine record:

```sh
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserHostBoundaries$'
```

[Scope6523160067](https://github.com/katsuakikureyama/GoHP/issues/1#issuecomment-6523160067).
GoHP scalar-request work is separately owned; this change adds no GoHP code or
feature count. No parity claim for PHP's unrestricted subject length or error
handling is made by these local host observations.
