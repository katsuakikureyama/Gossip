# Original parser-token native reference

Status: 24 native PHP 8.3.36 expectations PASS, 2026-10-27 JST. This prepares
acceptance data for the [backend assessment](parser-backend-assessment.md); it
does not translate PHP or compare a Go implementation. See the
[execution record](evidence/parser-token-reference-2026-10-27.md).

## Original and owned boundaries

The [manifest](../gossip/tests/gohp/parser-token-reference.json) fixes WordPress
`54106f6c545fc8b881f6b39b7963002bece6c28c`, complete Parser/Block/Frame files,
license and version hashes. Native PHP loads the full original parser and its
real require_once dependencies. Reflection checks all seven methods and the
original next_token definition at lines 239–303. No source method is rewritten.

The [Go harness](../gossip/tests/gohp/parser_token_reference_test.go) verifies
those bytes before staging them. It extracts the exact single-quoted pattern
literal at line 251 for a separate direct preg_match probe: no evaluation or
pattern rewrite. This pinned literal contains neither PHP escaped quotes nor
double backslashes; stripping its quotes leaves runtime bytes unchanged.
The original method still uses its own original literal.

The [owned observer](../gossip/tests/gohp/parser-token-reference-observer.php)
sets parser input/offset and sentinel output/stack, invokes the original method,
and separately invokes preg_match with PREG_OFFSET_CAPTURE. The
[24 cases](../gossip/tests/gohp/parser-token-reference-cases.json) declare expected
match type/value, PCRE error and five-element token before execution. PHP strict
comparisons check them and unchanged parser/caller state. Full numeric/named
capture vectors and JSON error codes are **observations**, not independently
predeclared expected vectors or cross-implementation equality.

## Bounded contract

The fixed inputs are canonical base64 byte strings, at most 4096 bytes (longest
retained case: 2072), with offsets in [-byte_length, byte_length]. Cases cover
open/close/void comments, namespaces, JSON braces/newlines, adjacent comments,
malformed/unterminated comments, multibyte prefixes, NUL and invalid UTF-8,
end/negative/interior-byte offsets and one backtracking-limit failure.

PHP reports match positions in bytes and separates integer 1, integer 0 and
Boolean false; the matches argument is an output parameter. The record retains
these distinctions instead of collapsing failure into no-match.
See the [preg_match contract](https://www.php.net/manual/en/function.preg-match.php).
Strings use base64; arrays retain ordered typed keys/values, missing captures
remain missing, and present unmatched captures retain their empty bytes/-1
offsets. Both pre-call matches=null and post-call values are recorded.

The original next_token maps both no-match and regex failure to no-more-tokens;
the separate probe preserves the difference. Regex matching and attribute JSON
decoding are separate: invalid UTF-8 attributes can match the non-/u pattern yet
decode to null with JSON error 5. No Unicode normalization is applied.

Execution fixes pcre.jit=0, recursion_limit=10000, memory_limit=128M, LC_CTYPE=C;
backtrack_limit is 100000 except the limit=1 case. Effective values and PCRE
version are emitted, not inferred from image tags. See PHP's
[PCRE settings](https://www.php.net/manual/en/pcre.configuration.php) and
[PCRE errors](https://www.php.net/manual/en/function.preg-last-error.php).

## Reproduction and isolation

From the nested `gossip/` module, using the already-cached admitted image:

```sh
GOSSIP_PHP_RUNNER_IMAGE=php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd \
GOPROXY=off GOTOOLCHAIN=local \
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserTokenReference$'
```

No GoHP checkout is required. Unset runner image means SKIP, not reference PASS.
The existing [oracle policy](php-oracle-policy.md) performs same-process version
admission before loading the fixture. The container is networkless/read-only,
non-root, capability-dropped, no-new-privileges, limited to 768MiB/256 PIDs/2 CPUs,
with a 90-second deadline. Only verified inputs/owned driver are mounted read-only.
No image pull, credential mount or host source execution is needed. Nonzero exit,
stderr, timeout or killed container is not a successful reference result.

## Capability and exclusions

`gossip.dev.wordpress-parser-token.reference` is a local developer test, explicitly
**excluded from WebMCP/UI**: no server operation, browser route, auth or automatic
submission is introduced. Input is the pinned fixture plus admitted image env;
inspect verifies identities, prepare stages a private temporary input directory,
execution emits a versioned evidence result or fails. There is no save/publish
API; committing evidence remains a separate repository operation. Supported
status is native-reference-only for the measured profile, not PHP compatibility.

No generated Go, GoHP feature admission, other preg functions/patterns/flags,
out-of-range offsets, JIT-on, other PCRE failures, PHP 8.4, parse loop, bootstrap,
HTTP/DB, plugins or deployment is tested. Backend speed/licensing/build-platform
comparisons remain not-run. Earlier fixtures/pins and translation totals stay
unchanged; extend version-specific evidence explicitly rather than extrapolating.
