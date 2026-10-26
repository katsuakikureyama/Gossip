# PCRE2 engine comparison: original parser pattern

Status, 2026-10-27 JST: **24/24 bounded engine comparisons PASS** against the
[retained PHP reference](parser-token-reference.md). An independently linked
PCRE2 10.47 C probe reproduces the selected preg_match results from PHP 8.3.36 /
PCRE 10.42. This is not a GoHP or cgo integration, backend adoption, complete
next_token execution, or a general equivalence claim across PCRE versions.
[Execution record](evidence/parser-pcre2-comparison-2026-10-27.md).

## What is compared

The [opt-in Go harness](../gossip/tests/gohp/parser_pcre2_comparison_test.go)
hash-checks the unchanged reference JSON, corpus and original WordPress files.
It supplies only the original pattern bytes, subject, normalized in-range byte
offset and match limit to the [C probe](../gossip/tests/gohp/testdata/parser-pcre2-probe.c).
The reference is never an input to the engine or mapping function. PHP delimiters
are removed and `/s` becomes PCRE2_DOTALL; no pattern syntax is rewritten.

The probe returns actual engine status, name table, numeric capture spans and
hex bytes. A fixture-only Go mapper orders named/numeric entries as PHP does,
retains unmatched middle entries with empty bytes/-1 and omits trailing unmatched
entries. PCRE2's positive result identifies the highest set capture plus one;
unset pairs are distinguishable from empty matched strings. Match/depth limits
are explicit. See the [PCRE2 API](https://www.pcre.org/current/doc/html/pcre2api.html).

The comparison checks typed return, error code, full ordered matches array,
named-capture presence and the fixture's null-to-output mapping. PHP's error text
is excluded, not manufactured from engine messages. PHP capture offsets are byte
offsets, and optional trailing captures are omitted without UNMATCHED_AS_NULL;
see [preg_match](https://www.php.net/manual/en/function.preg-match.php).

| Engine result | Fixture PHP-shaped result |
| --- | --- |
| Positive capture count | integer 1; ordered named/numeric captures |
| PCRE2_ERROR_NOMATCH (-1) | integer 0; empty matches array; error 0 |
| PCRE2_ERROR_MATCHLIMIT (-47) | Boolean false; empty matches array; error 2 |
| Any other result or invalid span/profile | Explicit test failure; no compatibility PASS |

This models an output value, not PHP reference aliasing or a compiler lowering.
The candidate does not execute next_token, JSON decoding or parser state changes;
those remain separate PHP-only observations from the earlier record.

## Fixed local build and execution profile

The candidate uses an already installed Homebrew PCRE2 10.47_1 arm64 static
archive and matching header, both SHA-pinned by the harness. No package install,
network fetch, source vendoring, shared Go-module dependency or binary publication
is introduced. The installed package's license metadata is
`BSD-3-Clause WITH PCRE2-exception`; its license hash is recorded. This is an
inspection record, not a completed downstream distribution/licensing decision.

Apple clang builds C11 with `-O2 -Wall -Wextra -Werror`. The engine is 8-bit,
DOTALL, LF newline, BSR_UNICODE, default C-locale tables, no UTF/UCP/JIT. Whole
subjects retain NUL/invalid UTF-8. Negative offsets are normalized within the
existing bounded input domain, without slicing the subject. Match limits are
100000/1, depth 10000, candidate heap limit 131072 KiB. That engine heap bound
is **not** asserted equivalent to PHP's 128M process memory setting.

The probe runs on the host, **not** inside the PHP oracle container. It receives
only fixture file paths/numeric options, a minimal environment and a five-second
deadline per case. Inputs are at most 4096 bytes; compilation has a 30-second
deadline. This is not an OS sandbox/network isolation claim or a safe public
arbitrary-pattern service. Only trusted pinned fixtures are admitted by the test.
Nonzero exit, stderr, timeout, malformed observations or unknown return codes fail.

From `GOSSIP_ROOT/gossip`, with the exact installed artifacts:

```sh
GOSSIP_PCRE2_PREFIX=/absolute/path/to/pcre2/10.47_1 \
GOPROXY=off GOTOOLCHAIN=local \
go test -count=1 -v ./tests/gohp -run '^TestParserPCRE2Comparison$'
```

Missing opt-in means SKIP, not PASS. This profile accepts darwin/arm64 only;
other builds need their own reviewed pins and evidence. PHP reference execution
is reused, not repeated. Library version differences are retained explicitly.

## Capability, remaining gates

`gossip.dev.wordpress-parser-pcre2.compare` is a local developer fixture, explicitly
excluded from UI/WebMCP/MCP: no server, auth route, automatic submit or public
input API. Inspect verifies pins, prepare stages trusted files/builds a temporary
probe, execution emits `gossip.parser-pcre2-comparison/v1` or fails. Input is the
pinned corpus/reference plus installed-prefix environment. No save/commit/publish
operation is exposed; evidence commits are separate repository operations.

Still unmeasured: regexp2; cgo/GoHP integration and general by-reference handling;
PHP 8.4; JIT-on, out-of-range offsets and other failures/flags/patterns; full
next_token/JSON/parse; cross-platform packaging and performance. Keep the
[backend evaluation gates](parser-backend-assessment.md): this result supports
PCRE2 as a candidate, not final adoption or substitution of the original parser.
