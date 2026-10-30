# Original-parser resource boundaries — 2026-10-30

## Fixed-profile comparison

This new series PASSes twelve pinned documents / twenty-four
actual complete-original `parse()` calls per backend use the existing fixed
match100000/depth10000/subject4096 profile. Native PHP is captured first, so a
generated build/error cannot discard the original observations. A finite
matching sample is not proof that resource failures are unreachable. One run
at clean Gossip951b637c3178403f7c90c1b3aaa093f9e021b320/tree
9e9a17c8f4c7c26e5d8bdf71c2ad1740393625af, unchanged GoHPb2d1482:
PASS2.61s/package2.866s, mismatch0, exits0/runtime stderr0. All24 PCRE codes
are0; no original regex failure was found in this finite sample. JSON codes0/4
reflect freeform versus malformed attributes. The
[machine record](parser-resource-boundaries-2026-10-30.json) retains both
complete observations and source/command/run_count1 provenance, from raw SHA
`2fa4edf06566fa8fbfe595ce151e47c13743c758cf15cc13477c74331b78f8ca`.
Both canonical typedcases hash to
`51fad737f3e1457e6273cd2a7cdc03567e131ab506387c72af207cc351b64d52`.
Go stdout393640bytes and PHP394445bytes have different metadata envelopes.
Candidate binary was test-cleaned; previous successful configurations were
not repeated. Linker warning is distinct from runtime stderr0.

The new [corpus](../../gossip/tests/gohp/parser-resource-boundaries-cases.json),
[pin](../../gossip/tests/gohp/parser-resource-boundaries.json) and
[harness](../../gossip/tests/gohp/parser_resource_boundaries_test.go) reuse both
existing full observers unchanged. Ten documents are4096bytes; alternating
brace controls are1024/2048bytes. Cases cover plain/whitespace, unclosed and
repeated openers, alternating/dense braces, near delimiters, embedded comments,
one paired outer block and brace runs. No deep successful object transport or
unbounded search is introduced.

Original classes/closure, checked generation35524a6, GoHPb2d1482, old pins,
corpora/observers/evidence and previous successful configurations stay intact.
The generated entry fixes PCRE match100000/depth10000/heap131072KiB, maximum
subject4096. Only match-limit failure maps to PHP false/error2; other engine
failures and unsupported input are explicit Go errors/panics. No caller budget
override exists on the generated path. A direct regex with limits1 or modified
PHP settings cannot establish this profile's original generated error path.

Native PHP8.3.36/PCRE10.42Linuxaarch64/C/JIT0/backtrack100000/recursion10000/
process memory128M versus Go1.23.0Darwinarm64/CGO1/staticPCRE10.47_1/clangSDK15.5
are distinct profiles. Engine heap and total process memory are not equivalent.
Host timeouts, subject guards, observation failure and original PHP error state
remain distinct outcomes. No setup/install/startup/pull/new CI resources.

Run from `gossip/` with explicit source/prefix/generated/image/fresh-evidence
inputs and `GOSSIP_GOHP_PARSER_RESOURCE_BOUNDARIES=1`:

```sh
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserResourceBoundaries$'
```

## Capability

| Field | Contract |
|---|---|
| Stable ID | `gossip.tests.parser-resource-boundaries` |
| Service/resource | Local developer test; pinned complete-original parser, bounded stress corpus and observations |
| Input schema | Opt-in1, absolute GoHP source/PCRE prefix/checked generation directory/fresh `GOSSIP_GOHP_RESOURCE_EVIDENCE_OUT`, exact cached PHP image; immutable corpus |
| Inspect / prepare | Verify original/pin/corpus/observer/engine/toolchain identities; reuse checked generation |
| Save / commit | One exclusive local evidence file; no application save/publish/external commit operation |
| Status / common result | Experimental, measured case-limited; skips, failed phases, native/generated error, comparison and application acceptance are separate; phases/hashes/typed observations/counts retained |
| WebMCP | `excluded`: local native compilation/container test, no browser/server operation, UI callsite or auth/route/deployment change |

Scope6521976972. G83-PARSER/application acceptance stays open; original error
paths, over-bound guards and timeout/resource equivalence must not be marked
passed merely because a finite current-profile stress sample agrees.

## Separate match1 native reference: pending

The separate [match1 pin](../../gossip/tests/gohp/parser-resource-match1.json),
[PHP driver](../../gossip/tests/gohp/parser-resource-match1-observer.php) and
[native harness](../../gossip/tests/gohp/parser_resource_match1_test.go) use
profile `php83-original-parser-match1-v1`: backtrack1, recursion10000, JIT0,
process memory128M, same twelve documents and complete original source. The
new driver explicitly changes only the owned reference's budget; existing
observers and the fixed generated profile stay unchanged. This is not yet
measured and cannot be counted as a PHP-versus-Go comparison.

| Field | Separate native-reference contract |
|---|---|
| Stable ID | `gossip.tests.parser-resource-match1-reference` |
| Service/resource | Local developer native-only inspection of complete original parser under named match1 budget |
| Input schema | Opt-in `GOSSIP_PHP_PARSER_MATCH1_REFERENCE=1`, exact cached PHP image, fresh absolute `GOSSIP_PHP_PARSER_MATCH1_EVIDENCE_OUT`; pinned corpus/source/observer |
| Inspect / prepare | Verify source/pin/corpus/new observer; reuse cached image, no compiler/candidate build |
| Save / commit | One exclusive local native evidence file; no application or external save/publish operation |
| Status / common result | Experimental/unverified; native-reference-recorded, failed and skip distinct; generated comparison explicitly not-run-unsupported-budget |
| WebMCP | `excluded`: local native container test, no browser/server/UI operation or authentication/route/deployment change |

Run only this new setting from `gossip/`:

```sh
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserMatch1Reference$'
```
