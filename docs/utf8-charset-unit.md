# Original UTF-8 charset-name predicate

This ninth standalone unit extracts `_is_utf8_charset` from the pinned
WordPress compat.php lines102–111 byte-for-byte. Its string guard, two
`strcasecmp` calls, strict zero comparisons, negation and short-circuit OR
remain original code. The whole file and exact excerpt are checked before
generation. It classifies a charset name; it does not validate, decode or
convert text bytes as UTF-8.

The [manifest](../gossip/tests/gohp/utf8-charset-helper.json) pins GoHP
`d9f6c7b4f288ef816c4ee3cb6a3458d2226e00d5`, based on the published #38
logical-comparison stack plus string-only ASCII `strcasecmp`. That compiler
does not contain the separate #36 isset/array-return additions used by the
fourteen-original bundle. Earlier standalone and bundle configurations are
unchanged; this is not a fifteen-original combined executable.

The 28 explicit inputs include six accepted ASCII case variants of UTF-8/UTF8,
fourteen rejected string spellings (space, suffix, NUL, non-ASCII dash/fullwidth
characters and other charset names), and eight non-string inputs. Both
observers require an actual Boolean return and record actual caller input
after invocation. Native PHP and generated Go must match each other and the
authored Boolean/unchanged-input expectations. Transport admits null, bool,
integers within ±1,000,000, strings and sequential nested lists (depth≤8), not
floats, JSON objects, associative keys, resources, references or cycles.

```sh
cd gossip
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressUTF8CharsetHelper$'
```

Without opt-in the test skips, not passes. The existing immutable images,
same-process PHP8.3.36 admission guard, LC_CTYPE=C, non-root execution,
read-only filesystem, disabled network and resource limits remain unchanged.
Report/source/generated hashes are checked before build. The
[execution record](evidence/utf8-charset-2026-10-25.md) separately binds the
actual test snapshot and measured outputs; generation alone is not PASS.

The public `is_utf8_charset` wrapper additionally needs default parameters,
null coalescing and get_option; those are not replaced by stubs or claimed
supported. `_canonical_charset`, encoding conversion, whole compat.php,
WordPress core/HTTP/DB/theme/plugin execution and PHP exception equivalence
remain unverified. No source rewrite, merge, deployment, release or license
decision is introduced.

Local developer capability `gossip.dev.wordpress-unit.translate-compare`
remains WebMCP-excluded. No browser/service/MCP operation or authentication
route changes. Devboard card unidentified/unsynchronized, not Done.
