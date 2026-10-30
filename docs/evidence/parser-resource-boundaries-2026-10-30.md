# Original-parser resource boundaries — 2026-10-30

## Contract and initial status

This new series is not yet measured. Twelve pinned documents / twenty-four
actual complete-original `parse()` calls per backend use the existing fixed
match100000/depth10000/subject4096 profile. Native PHP is captured first, so a
generated build/error cannot discard the original observations. A finite
matching sample is not proof that resource failures are unreachable.

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
| Status / common result | Experimental/unverified until measured; skips, failed phases, native/generated error, comparison and application acceptance are separate; phases/hashes/typed observations/counts retained |
| WebMCP | `excluded`: local native compilation/container test, no browser/server operation, UI callsite or auth/route/deployment change |

Scope6521976972. G83-PARSER/application acceptance stays open; original error
paths, over-bound guards and timeout/resource equivalence must not be marked
passed merely because a finite current-profile stress sample agrees.
