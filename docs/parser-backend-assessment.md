# Block-parser regex backend: evaluation gates

Status: **design assessment only**, 2026-10-26 JST. No regex library, cgo build,
source rewrite, benchmark or new PHP compatibility result is introduced. This is
Gossip's consumer requirement for a general GoHP runtime contract, not a WordPress
special case in the compiler. Follow the [responsibility map](reuse-boundary.md).

Update, 2026-10-27: the [native reference corpus](parser-token-reference.md) now
has 24 passing PHP 8.3.36 cases with full capture/error/state observations.
The follow-up [independent PCRE2 10.47 probe](parser-pcre2-comparison.md) matches
all 24 retained direct-regex observations from PHP/PCRE10.42. Only this bounded
C-engine profile is measured: no GoHP/cgo adapter or backend adoption. Regexp2
remains unmeasured; the original assessment below remains the broader design gate.

## Fixed original requirement

Upstream WordPress `54106f6c545fc8b881f6b39b7963002bece6c28c`,
`wp-includes/class-wp-block-parser.php`, SHA-256
`8563621090de5e307431b1d29b78a9f75fd9ee238a0137b563a7db7b2e1f8d3b`.
The existing [add_freeform manifest](../gossip/tests/gohp/parser-add-freeform.json)
already binds whole-file and selected-fragment identities; its successful
data/output-path cases do **not** execute `next_token()`.

The [original next_token](https://github.com/WordPress/WordPress/blob/54106f6c545fc8b881f6b39b7963002bece6c28c/wp-includes/class-wp-block-parser.php#L239)
calls `preg_match` at lines 250–256 with named captures (`closer`, `namespace`,
`name`, `attrs`, `void`), positive/negative lookahead, possessive `*+`, `/s`,
`PREG_OFFSET_CAPTURE`, and a starting offset into the original document. There
is no `/u` modifier. Do not normalize strings to Unicode text or discard malformed
bytes before matching. Preserve the literal original pattern in acceptance tests;
do not edit WordPress to fit a chosen engine.

The caller distinguishes false and zero (both currently lead to no-more-tokens),
uses matched byte offsets/lengths and absent/unmatched named captures, then calls
`json_decode(..., true)` for attributes. It also uses list destructuring. Other
parser methods have switch/stack mutation and shared Block/Frame state; the file
requires real Block and Frame sources. These are separate dependency requirements,
not a claim that each is a newly observed compiler rejection. Regex success alone
would not establish a complete parser translation.

## Candidates, not an adoption decision

| Candidate | Reason to evaluate | Unresolved adoption conditions |
|---|---|---|
| PCRE2 8-bit via cgo | PHP uses PCRE2; same-family backend is the first semantic-baseline candidate | Pin PCRE2/wrapper revisions, compile options/JIT and limits; implement PHP wrapper behavior, ownership/errors; evaluate cgo packaging, reproducibility, supported OS/architectures and notices |
| Pure-Go regexp2 | Potential deployment simplicity without a C dependency; supports relevant advanced constructs | Pin exact release/commit and Go floor; evaluate named captures/options, raw-byte behavior, offsets, match failure and bounded resource use against the same PHP oracle |
| Go standard regexp | Useful only for explicitly supported regular-expression subsets elsewhere | Cannot accept this original lookahead/possessive syntax as-is; not a drop-in implementation for this selected pattern |
| Handwritten Go block tokenizer/parser | Possible separate native experiment if later deliberately chosen | Not original-source translation evidence; cannot satisfy the original-parser acceptance gate or hide missing GoHP semantics |

PHP's [installation documentation](https://www.php.net/manual/en/pcre.installation.php)
identifies PCRE2 as its backend. The [PCRE2 API](https://www.pcre.org/current/doc/html/pcre2api.html)
has an 8-bit interface, match offsets and resource controls, but it is not PHP's
`preg_match` wrapper. A PHP-free executable need not imply a pure-Go executable;
admitting cgo would be an explicit build/platform decision, not an embedded PHP
interpreter.

The inspected [regexp2 README](https://github.com/dlclark/regexp2/blob/a6ef7b6f88fff76c0c7d4f5fe832535de8cfda11/README.md)
documents possessive quantifiers in its default/RE2 modes, lookarounds, rune-based
matching and a byte-range API, plus timeout/backtracking controls. Therefore
mandatory `*+` rewriting is not a valid blanket assumption about current versions.
The byte-range API alone does not prove PHP non-UTF byte semantics. These are
upstream documentation observations, not local tests or selection of that commit
as a dependency. No conclusion about speed or full PCRE compatibility is made.

## Required comparison contract

Before coding a production adapter, pin the same unchanged pattern, input corpus
and driver for each candidate, along with native PHP target/version/image and
effective PCRE version/JIT/limits/locale. Record Go/compiler/runtime/backend/build
identities. Existing Gossip PHP 8.3.36 evidence does not cover PHP 8.4; any 8.4 run
needs explicit runner admission and its own results, not reuse of 8.3 observations.

The [PHP preg_match contract](https://www.php.net/manual/en/function.preg-match.php)
requires separate handling of matches, no match and failure; matches is an output
parameter and capture/start positions are byte offsets. The first bounded adapter
must specify accepted patterns, flags, offset ranges and output destinations.
General reference support or arbitrary preg_* support is not implied by a special
output-argument lowering. Unsupported forms must fail explicitly.

Measure at least the following before admitting this consumer path:

- Actual return type/value and every numeric/named capture entry: bytes, order,
  presence, unmatched offset and mutation of the matches output. Empty capture,
  absent capture and unmatched capture must not collapse into one state.
- Starting offsets at zero/end/outside the document and inside multibyte text;
  ASCII, UTF-8, NUL and invalid byte sequences. Do not implement offset by slicing
  the subject and silently changing assertion context or reported positions.
- Real block open/close/void comments, namespaces, attribute braces/newlines,
  malformed comments, unmatched groups and adjacent/nested-looking input.
- Long/adversarial malformed comments, no-match input and resource-limit/error
  behavior. A timeout/sandbox kill is not a PHP no-match or compatibility PASS.
  Preserve unobserved diagnostics/errors as unmeasured, not fabricated equality.

Separate engine matching from complete `next_token()` behavior (including JSON
semantics) and then complete `parse()` behavior with the original dependencies.
Keep originals, extraction scaffolding and owned observations identifiable.
Only the exact selected contract can pass; other preg functions/patterns remain
unsupported or unverified.

## Decision and next deliverable

Evaluate **PCRE2 first as the semantic baseline and regexp2 as the pure-Go
alternative**, using the same fixed oracle cases. This is evaluation order, not
permission to install an unreviewed wrapper or a final dependency choice. Report
rejections and differences as well as successful cases. Correctness/raw-byte
behavior, bounded failures, maintainability, licensing and platform costs precede
performance claims. If neither meets the contract, retain the blocker; do not
silently rewrite the source or replace the parser.

Next deliverable: regexp2 comparison and the remaining runtime/build-policy gates,
using the same versioned reference and explicit unmeasured cells. The first corpus covers in-range byte offsets,
malformed bytes, capture mutation and one backtracking-limit failure; out-of-range
offsets, JIT-on and other failure modes remain unmeasured. Work on other independently
bounded parser/storage dependencies need not wait for a final backend decision.
No user-facing capability, route or deployment is added.
