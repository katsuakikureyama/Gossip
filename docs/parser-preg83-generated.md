# Generated GoHP regex consumer — PHP 8.3 experiment

This is a new isolated Gossip consumer of GoHP's bounded B3 compiler profile.
It does **not** transpile `WP_Block_Parser::next_token()` or the complete parser.
It compiles a Gossip-owned PHP driver using the exact original regex bytes, then
observes the resulting Go executable against retained native PHP evidence.

## Inputs, pins and denominator

- WordPress baseline: `54106f6c545fc8b881f6b39b7963002bece6c28c`.
- GoHP compiler and runtime: `d7650944802a33cfc2dab0006bb410374d828245`
  ([merged B3 PR #114](https://github.com/katsuakikureyama/GoHP/pull/114)).
  This fixture alone adds that pin. Earlier Gossip fixtures retain their pins;
  later GoHP main changes are not silently consumed.
- Original source/reference/corpus identities come from the retained
  [parser-token reference](parser-token-reference.md). Its 24-case native PHP
  8.3.36 / PCRE 10.42 record remains immutable and is reused, not rerun.
- Selected: **23 cases** with backtrack limit 100000 (17 matches, 6 no-matches).
  Excluded: **1 case**, `backtrack_limit_one`, because GoHP B3 fixes the limit at
  100000. Running that input with a different limit is not the same case. This
  fixture never claims 24/24 generated compatibility or a consumer failure-case
  result. GoHP's separate generic failure/state evidence is not copied here.

The complete original parser file is hash-checked before extracting the literal
on line 251. Its runtime pattern hash and bytes must also equal the retained PHP
record. The owned driver re-encodes those bytes as a PHP single-quoted literal
(doubling backslashes); this changes source spelling, **not regex syntax or
runtime bytes**. The original WordPress file is never edited. No pattern/name/hash
special case is added to GoHP. Original property reads are not claimed supported:
this driver takes actual typed string/integer parameters instead.

The driver initializes its output local to null, snapshots the initial value,
executes the dedicated `preg_match` statement, then returns captures, typed match
result and `preg_last_error()`. Its observer only serializes returned values and
capture presence; it neither matches regexes nor receives expected outputs.

## Build and observation contract

The test exports only the immutable pinned GoHP Git archive, not the clone's
working tree, `.git`, credentials or untracked files. The CLI is built with CGO0,
then explicitly invoked with `-php-version 8.3 -experimental-pcre2`. Its fresh
provenance report must match the selected profile and source/generated hashes.
The emitted Go and owned observer are built against that same archived runtime
with `gohp_pcre2` and CGO1. No PHP interpreter is linked or spawned by the candidate.

This first optional lane uses the existing measured **Darwin ARM64 / Go 1.23.0 /
PCRE2 10.47 static archive** profile, explicit C compiler and SDK. Header, archive
and compiler identities are checked/recorded. It installs or downloads nothing.
Standard Gossip/GoHP CGO0 tests remain separate; no enabled Linux/other-platform
support or redistributable native package is inferred.

The candidate receives only selected names, base64 input bytes, offsets and the
fixed limit label. Its process runs without PHP on PATH. Comparison retains
actual scalar types, every ordered numeric/string key, raw capture bytes,
byte offsets, absent/unmatched captures, before/after output values and error
codes. PHP error-message text is excluded, not fabricated. JSON structural
equality is used; object serialization ordering is not semantic evidence.
NUL, invalid UTF-8, multibyte-interior offsets, named captures, JSON-looking text
and unterminated comments are in the fixed inputs. Attributes are regex captures,
**not** decoded JSON. Full method results and parser state remain unmeasured.

## Run locally

Use a cached exact Go 1.23.0 on PATH, a local GoHP clone containing the pin,
and the already installed pinned PCRE2 prefix. With explicit `CC`, `SDKROOT`,
`GOTOOLCHAIN=local`, `GOPROXY=off`, `GOSUMDB=off`, and `GOWORK=off`:

```sh
cd gossip
GOSSIP_PREG83_GOHP_SOURCE=/absolute/local/GoHP \
GOSSIP_PCRE2_PREFIX=/absolute/pinned/pcre2-prefix \
go test -count=1 -v ./tests/gohp -run '^TestParserPreg83GeneratedConsumer$'
```

Both opt-in variables must be provided. Without them the enabled comparison
explicitly skips; skip is not compatibility PASS. The ordinary contract test
checks source/pin/corpus selection without requiring native dependencies.
Do not set the general `GOSSIP_GOHP_SOURCE` to run this fixture: it would also
enable unrelated historical Docker fixtures in a broader test invocation.

UI / WebMCP / MCP capability: **excluded**, developer-only local test operation
with temporary build artifacts and read-only source inputs. No service route,
authentication, publication API, automatic submit or production target is added.
Licensing/notices, binary distribution, enabled Linux CI, general output locations,
JSON/list destructuring and original `next_token()` dependency closure remain
separate work. The next step is to measure those actual original-source gaps,
not replace the parser with this driver.
