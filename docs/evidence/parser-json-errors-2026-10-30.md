# Original-parser JSON error boundaries — 2026-10-30

## Contract and status

This separate series is **not yet measured**. It compares twelve pinned
attribute documents / twenty-four actual complete-original `parse()` calls per
backend against fixed PHP8.3.36. It does not enlarge or rerun the original
twelve-case series or the four carry-in sequences. G83-PARSER stays open.

The new [corpus](../../gossip/tests/gohp/parser-json-errors-cases.json),
[pin](../../gossip/tests/gohp/parser-json-errors.json) and
[harness](../../gossip/tests/gohp/parser_json_errors_test.go) reuse the existing
[PHP observer](../../gossip/tests/gohp/parser-full-reference-observer.php) and
[Go observer](../../gossip/tests/gohp/parser-full-generated-observer.go.txt)
byte-for-byte. Their twelve-document contract, typed ordered/base64 transport,
fresh/reused parser calls, one shared Execution and per-call JSON success
baseline fit this series. The baseline deliberately isolates each decode error;
this is not another carry-in measurement.

| Family | Separate case IDs | Input location |
|---|---|---|
| Invalid UTF-8 | `invalid_utf8_ff_value`, `invalid_utf8_truncated_value`, `invalid_utf8_key`, `invalid_utf8_overlong_value` | Raw invalid bytes within a JSON attribute key/string, not freeform text |
| Raw controls | `raw_nul_in_string`, `raw_lf_in_string`, `raw_tab_in_string` | Unescaped NUL, LF or TAB within an attribute string |
| UTF-16 escapes | `lone_high_surrogate`, `lone_low_surrogate`, `high_surrogate_before_nonlow` | Incomplete/invalid surrogate pairing within an attribute string |
| Default JSON depth | `default_depth_nested_arrays`, `default_depth_nested_objects` | Outer attribute object plus512 nested containers; documents1054/3102bytes |

Every document uses the original brace-bound attribute callsite at line283.
The harness requires nonzero JSON error state after the observer's success
baseline, successful tokenizer PCRE state and observation of all four error
families. Complete native typed cases remain the comparison authority; expected
parser trees are not manufactured from generated output.

The depth cases concern JSON decode failure and its shallow `null` result.
They do not test successful values near the512 limit, observation depth over64,
or exact threshold/resource equivalence. An observer failure, process failure
or timeout cannot count as a matching PHP JSON error. JSON scalar roots, float/
overflow, arbitrary malformed inputs, original PCRE error/resource paths,
PHP8.4 and later application gates remain outside this series.

## Execution profile

Use the existing checked generation/report35524a6 and unchanged GoHPb2d1482,
complete original Parser/Block/Frame and all prior pins/observers/evidence.
The harness checks hashes and no implementation diff between generation/runtime
commits; it does not run the compiler again. Exact Darwin/arm64 Go1.23.0,
clang/SDK15.5, static PCRE10.47_1 CGO1/taggohp_pcre2 candidate; parent CGO0.
Native fixed PHP8.3.36/PCRE10.42 Linux/aarch64, C locale, JIT0,
backtrack100000/recursion10000/process memory128M. Profiles differ.
Existing cached toolchain, image, running Docker and shared cache only;
no install/startup/pull, persistent binary or deployment is part of this test.

Run from `gossip/` with explicit source/prefix/checked-generation/image paths,
`GOSSIP_GOHP_PARSER_JSON_ERRORS=1` and a fresh absolute
`GOSSIP_GOHP_JSON_ERRORS_EVIDENCE_OUT`:

```sh
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserJSONErrors$'
```

## Capability

| Field | Contract |
|---|---|
| Stable ID | `gossip.tests.parser-json-errors` |
| Service/resource | Local developer test; unchanged original parser and pinned attribute-error observations |
| Input schema | Opt-in `1`; absolute GoHP source, PCRE prefix, checked generation directory and fresh evidence output; exact cached PHP image; immutable corpus, not user-selected input |
| Inspect / prepare | Verify source, corpus, observer, manifest and build-profile hashes; reuse checked generation |
| Save / commit | One exclusive local evidence file; no application save/publish/external commit operation |
| Status / common result | Experimental/unverified until measured; skip, failure and comparison PASS remain distinct; phases, hashes, typed observations, errors and counts are recorded; application acceptance not-established |
| WebMCP | `excluded`: local native compilation/container test, no browser/server operation or UI callsite; no browser/S2S/MCP/webhook/public route or authentication change |

Key/type boundaries belong to separate scope6521347316. This JSON-error scope
is6521354926, after superseded duplicate claim6521348620; no key/type work or
measurements were started under that claim. Devboard remains unidentified;
Issue1 and current-state records track progress.
