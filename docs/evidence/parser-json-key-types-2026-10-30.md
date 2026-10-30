# Original-parser JSON key/type boundaries — 2026-10-30

**PASS for twelve new documents / twenty-four actual original `parse()` calls
per backend. G83-PARSER and whole Gossip acceptance remain open.**

The [machine record](parser-json-key-types-2026-10-30.json) preserves both actual
observations, exact signed64 integers, hashes, command/environment and a single
measured attempt. This is a separate series from
[fixed-v1](parser-full-fresh-2026-10-30.md) and
[carry-in state](parser-carry-state-2026-10-30.md); neither was rerun or repinned.
No GoHP production change was needed.

## Measured cases

| Case | Native PHP and generated original agree |
|---|---|
| `duplicate_top_level` | Replacement retains first insertion position: first3, middle2, last4 |
| `escaped_duplicate` | `a` and escaped `\u0061` replace the same key in place |
| `canonical_numeric_keys` | `0`, `1`, `-1`, `2` become integer keys in input order |
| `noncanonical_numeric_keys` | `01`, `-0`, `+1`, `1.0`, `1e0`, space-prefixed1 remain string keys |
| `int64_key_bounds` | Signed64 extrema become integer keys; adjacent out-of-range **key strings** remain strings |
| `duplicate_numeric_keys` | Escaped1 replaces integer key1 in place; padded01 stays distinct |
| `int64_value_extrema` | Exact min/max signed64 values, zero and numeric negative-zero → integer0 |
| `empty_structures` | Associative decoding maps empty objects/lists to empty PHP arrays, including nested values |
| `nested_duplicates` | Replacement order and scalar types survive inside nested objects/lists |
| `list_mixed_types` | null, booleans, signed64 integers, strings, objects and lists remain typed and ordered |
| `decoded_string_bytes` | All standard escapes, decoded NUL/control bytes, UTF-8/emoji, empty/NUL/Unicode keys match byte-for-byte |
| `adjacent_attribute_isolation` | Two adjacent blocks retain distinct attrs, duplicate replacement and bool/null/empty-list values |

All documents are at most156bytes and contain brace-bound, valid JSON object
attributes. Out-of-range numeric **values**, fractions/exponents and invalid
JSON are absent. Escaped control characters decode successfully here; they do
not establish raw-control/invalid-encoding error behavior.

Each document runs with a fresh parser and a parser reused across all twelve
documents. All generated calls share one `Execution`. The unchanged observers
establish a caller JSON-null/error0 baseline before each call, as in fixed-v1;
this series does not replace the separate carry-in proof. Both sides return
JSON0/`No error` and PCRE0 throughout. The new pin additionally requires actual
named attribute blocks and `attrs` array type, so treating comments as freeform
or skipping JSON decode cannot pass. The final case returns two blocks; all
others return one. Complete typed output, key order, raw bytes, all parser
properties, caller preservation and fresh/reused state match.

## Identity and validation

- Clean measured Gossip commit `2729405b83646390fc7171ce90860e99b551deb5`, tree
  `be4da5e95b7dfa6eea47be2cdb9c8edd34eef227`; unchanged GoHP
  `b2d1482268a7710c515c07bf919af3d40b09d17c` and complete original WordPress
  Parser/Block/Frame with the existing five source/provenance hashes.
- New [corpus](../../gossip/tests/gohp/parser-json-key-types-cases.json) and
  [pin](../../gossip/tests/gohp/parser-json-key-types.json) bind the unchanged
  [PHP observer](../../gossip/tests/gohp/parser-full-reference-observer.php) and
  [Go observer](../../gossip/tests/gohp/parser-full-generated-observer.go.txt).
  No copied/rewritten observers or alternate parser. The new
  [shared JSON-boundary runner](../../gossip/tests/gohp/parser_json_boundaries_test.go)
  validates the pinned block/error/type contract, exact native profile and full
  typed case equality. Separate future corpora need their own pins and records.
- One configuration/attempt: `TestOriginalWordPressParserJSONKeyTypes`
  PASS2.57s/package2.788s,12cases/24actual calls per side,mismatch0, exits0 and
  runtime stderr0. Raw attempt SHA
  `f22bc6d64a329e2897a77570cf09dccade651b1d963382945d1ef32b991880de`.
- Canonical typed `cases` SHA on both sides:
  `3a49f9f142a45bc7e2d8167efd186edf0a8f7c8fd87be89c68ae76ed73efeac3`.
  Generated stdout67958bytes; PHP stdout68763bytes. Metadata envelopes differ;
  no whole-process stdout equality claim. The published record adds provenance
  to the same raw observations without rounding signed64 values.
- Existing checked generated Go/report from35524a6 and runtimeb2d1482 reused,
  after hash and empty implementation-diff checks. No transpilation or prior
  successful fixed-v1/carry-in suites/builds/CI/artifact downloads repeated.
  No new GoHP source/PR/CI.

Run from `gossip/` with the exact `commands.environment` in the machine record:

```sh
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserJSONKeyTypes$'
```

Opt-in is `GOSSIP_GOHP_PARSER_JSON_KEY_TYPES=1`, with explicit absolute source,
existing PCRE prefix, checked generation directory and fresh evidence path
`GOSSIP_GOHP_JSON_KEY_TYPES_EVIDENCE_OUT`, plus the exact cached PHP image.
Absent opt-in skips without claiming support. Offline cached Go1.23.0 runs the
harness with CGO0; only its child build uses CGO1/taggohp_pcre2, verified clang,
SDK15.5 and static PCRE10.47_1 on Darwin/arm64. Actual PHP8.3.36/PCRE10.42 runs
on Linux/aarch64, C locale, JIT0/backtrack100000/recursion10000/memory128M.
Engine/platform profiles differ, so this is not equal resource-failure proof.
The duplicate-library linker warning is retained separately from runtime
stderr0. Existing Docker25.0.3/image/engine/cache only; no install, startup,
pull, persistent binary or deployment. Test cleanup removes its executable.

## Capability and remaining scope

`gossip.tests.parser-json-key-types` is an experimental developer-local test
over the pinned original-parser resource. Its input schema is the explicit
opt-in/absolute paths/exact image/fresh evidence output above; corpus and
contracts are pinned rather than user-selected. Inspect/prepare verify source,
inputs, observers and toolchain. Save writes one exclusive local evidence file;
there is no application commit/publish operation. Results distinguish
PASS/FAIL/skip and retain phases/counts/hashes/diagnostics; application acceptance
remains not-established. WebMCP is **excluded** because this is local native
compilation/container testing, with no UI callsite, browser/server operation,
auth route or deployment change.

The key/type gate is now **passed-case-limited**. Invalid UTF-8/raw control/lone
surrogate/default-depth behavior belongs to separate scope6521354926 and is not
measured here. Original-pattern PCRE/resource errors, other public APIs,
references/floats/overflow values and arbitrary inputs remain unmeasured or
outside profile. G83-PARSER stays open; bootstrap/data/HTTP/site are not-run.
Issue1/current-state track progress; devboard remains unidentified.
