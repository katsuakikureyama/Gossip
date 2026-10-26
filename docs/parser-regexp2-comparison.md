# Regexp2 comparison: measured differences, not compatibility PASS

2026-10-27 JST: the fixed regexp2 string-API profile produced **22 equal outputs,
1 different outcome and 1 non-comparable limit case** against the same 24-case
[PHP reference](parser-token-reference.md). The harness completed successfully;
that is not a backend compatibility PASS. [Execution evidence](evidence/parser-regexp2-comparison-2026-10-27.md).

## Fixed candidate and original pattern

Candidate: [regexp2/v2 commit a6ef7b6](https://github.com/dlclark/regexp2/tree/a6ef7b6f88fff76c0c7d4f5fe832535de8cfda11).
Its go.mod requires Go 1.25; the probe uses cached Go1.25.14 darwin/arm64. This does
not raise Gossip's existing module floor or change any GoHP pin/dependency.
The owned [probe](../gossip/tests/gohp/testdata/parser-regexp2-probe.go) is built
in a temporary module with a local replacement pointing to a committed git
archive of that exact revision. No dirty checkout files or Git credentials are
mounted/compiled; archive/module/license hashes are retained. The temporary
require version is only a placeholder resolved by that local replacement, not
a claim to test a published v2.0.0 release. Upstream license metadata: MIT;
source and compiled binary are not vendored or published by this change.

The [harness](../gossip/tests/gohp/parser_regexp2_comparison_test.go) verifies the
original WordPress bytes, unchanged corpus and saved PHP evidence. It removes
only PHP pattern delimiters, selects `RE2|Singleline` (528), and does not rewrite
the pattern. RE2 mode admits original `(?P<name>)` groups and retains regexp2
lookaround/possessive support; it is not Go's standard regexp engine. This mode
also changes character classes, so syntax acceptance alone is not PCRE semantics.
See the fixed [upstream README](https://github.com/dlclark/regexp2/blob/a6ef7b6f88fff76c0c7d4f5fe832535de8cfda11/README.md).

## Measurements

| Selected surface | Result |
| --- | --- |
| 22 ordinary cases | Output equality: return shape, error code, complete ordered numeric/named captures, raw bytes, byte positions and optional capture presence |
| `offset_inside_utf8_prefix` | Different: PHP searches from byte 1 inside the prefix and finds the comment at byte 3; regexp2 returns an API error requiring a rune-aligned start |
| `backtrack_limit_one` | Not comparable: PHP returns false/error 2; candidate matches under its independent stack/time profile, which cannot express PCRE's match-count limit |

`FindStringMatchStartingAt` takes a byte index but rejects starts inside a valid
UTF-8 rune. No slicing, boundary rounding or input rewriting hides that difference.
For string input, `Capture.String()` slices the original string using `ByteRange()`;
the observed invalid UTF-8 attributes retain FF, and NUL/invalid-byte prefixes and
normal multibyte offsets also agree. Internal matching is still rune-based: these
observations do not prove arbitrary binary-string or byte-oriented PCRE semantics.
See [start validation](https://github.com/dlclark/regexp2/blob/a6ef7b6f88fff76c0c7d4f5fe832535de8cfda11/stringprefixfilter.go#L446)
and [capture implementation](https://github.com/dlclark/regexp2/blob/a6ef7b6f88fff76c0c7d4f5fe832535de8cfda11/match.go#L76).

Actual group names/numbers, participation, rune coordinates, byte coordinates and
base64 bytes are saved before fixture-only PHP shaping. Trailing unmatched groups
are omitted only by that explicit mapper; raw groups remain intact. The mapper
does not consume reference answers. Unknown engine errors remain observations;
they are not converted to PHP no-match or PREG_BACKTRACK_LIMIT_ERROR. Error text
is retained but not an equality field. PHP reference aliasing is not implemented.

Resource controls are stack 100000 integer slots plus a 100ms approximate timeout,
not PHP's pcre.backtrack_limit or recursion_limit. The limit=1 case is explicitly
non-comparable rather than tuned until an error appears. Even the 22 equal
outputs do not establish equivalent resource contracts. No timeout or stack-limit
error was observed in this corpus. See [upstream controls](https://github.com/dlclark/regexp2/blob/a6ef7b6f88fff76c0c7d4f5fe832535de8cfda11/options.go#L40).

## Reproduction and boundaries

From `GOSSIP_ROOT/gossip`, with the official source revision available locally:

```sh
GOSSIP_REGEXP2_SOURCE=/absolute/path/to/regexp2-clone \
GOSSIP_REGEXP2_GO=/absolute/path/to/go1.25.14/bin/go \
GOPROXY=off GOTOOLCHAIN=local \
go test -count=1 -v ./tests/gohp -run '^TestParserRegexp2Comparison$'
```

Missing opt-in means SKIP. The probe builds with CGO_ENABLED=0, GOWORK=off,
GOTOOLCHAIN=local, GOPROXY=off and GOSUMDB=off against a hash-recorded git archive,
not a downloaded module or vendor copy. Inputs are trusted regular files bounded
to 4096 bytes; each process has a five-second deadline and minimal environment.
This is host execution, not an OS/network sandbox or public arbitrary-regex service.
Build failure, nonzero exit, stderr or malformed evidence fails the harness.
Successful observation collection can include semantic differences by design.

`gossip.dev.wordpress-parser-regexp2.compare` is explicitly UI/WebMCP/MCP excluded:
local fixture only, no route/auth or submit operation. Inputs are pinned corpus,
reference, source-repository path and cached toolchain path. Inspect verifies pins;
prepare exports committed source and builds a temporary probe; execution emits
`gossip.parser-regexp2-comparison/v1` with compatibility_status=not-established.
No save/commit/publish API; evidence publication is a separate repository action.

## Decision after this comparison

Do not adopt this direct string-API regexp2 profile as a drop-in backend for the
original parser contract. The [PCRE2 experiment](parser-pcre2-comparison.md) remains
the better measured baseline for these exact cases (24/24 versus 22/1/1), not
proof of global superiority, production readiness or final dependency selection.
Next: scope a GoHP PCRE2 bridge/output-argument contract and review build/platform
and distribution conditions. A future byte-to-rune adapter would be a new profile,
not a retroactive fix to this result; general limits still require a policy.

No new PHP run, PHP8.4 result, GoHP runtime/compiler change, next_token/JSON/parse
execution, deployment or performance comparison. JIT/other patterns/flags,
whitespace-class edge cases, arbitrary raw-byte matching and packaging remain
unmeasured. Earlier fixtures/pins and successful PCRE2 evidence are unchanged.
