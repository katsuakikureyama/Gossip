# Original-parser JSON error boundaries — 2026-10-30

## Contract and status

This separate series **PASSes twelve pinned attribute documents / twenty-four
actual complete-original `parse()` calls per backend** against a new fixed
PHP8.3.36 execution. One configuration/attempt, test2.47s/package2.678s,
mismatch0; both processes exit0 with empty runtime stderr. It does not enlarge
or rerun the original twelve-case series or the four carry-in sequences.
G83-PARSER stays open.

The [machine record](parser-json-errors-2026-10-30.json) retains both complete
observations and adds source/command/run-count provenance to the same raw
measurement. Published SHA256
`4076f9005015d3b54fd0f3933f424f1a04c6afda0f484ae664fffea4d22cc1a0`;
raw SHA256 `b17f460d9b5faf9551373bff667eeed491982aec9f1b48987d55e0360a94b229`.

The new [corpus](../../gossip/tests/gohp/parser-json-errors-cases.json),
[pin](../../gossip/tests/gohp/parser-json-errors.json) and
[harness](../../gossip/tests/gohp/parser_json_errors_test.go) reuse the existing
[PHP observer](../../gossip/tests/gohp/parser-full-reference-observer.php) and
[Go observer](../../gossip/tests/gohp/parser-full-generated-observer.go.txt)
byte-for-byte. Their twelve-document contract, typed ordered/base64 transport,
fresh/reused parser calls, one shared Execution and per-call JSON success
baseline fit this series. The baseline deliberately isolates each decode error;
this is not another carry-in measurement.

| Family / observed JSON code | Separate case IDs | Input location |
|---|---|---|
| Invalid UTF-8 /5 | `invalid_utf8_ff_value`, `invalid_utf8_truncated_value`, `invalid_utf8_key`, `invalid_utf8_overlong_value` | Raw invalid bytes within a JSON attribute key/string, not freeform text;8calls per backend |
| Raw controls /3 | `raw_nul_in_string`, `raw_lf_in_string`, `raw_tab_in_string` | Unescaped NUL, LF or TAB within an attribute string;6calls per backend |
| UTF-16 escapes /10 | `lone_high_surrogate`, `lone_low_surrogate`, `high_surrogate_before_nonlow` | Incomplete/invalid surrogate pairing within an attribute string;6calls per backend |
| Default JSON depth /1 | `default_depth_nested_arrays`, `default_depth_nested_objects` | Outer attribute object plus512 nested containers; documents1054/3102bytes;4calls per backend |

Every document uses the original brace-bound attribute callsite at line283.
The harness requires nonzero JSON error state after the observer's success
baseline, successful tokenizer PCRE state and observation of all four error
families. Complete native typed cases remain the comparison authority; expected
parser trees are not manufactured from generated output.

The recorded error messages also match: malformed UTF-8, control-character,
unpaired UTF-16 surrogate and maximum-stack-depth errors. Every original token
has PCRE0; every block is `core/paragraph` with `attrs=null`, rather than
freeform fallback. Ordered keys, raw string bytes, complete parser fields,
caller preservation and fresh/reused parser results agree. Canonical typed
`cases` SHA256 on both sides:
`cdfa76d9c353c43a390ef44120d8f07e225e86979b2ec240b12ab652ef82933e`.
Generated stdout58737bytes/SHA256
`3a78f59ced69d2c9ca9c00e8ee1b09b5b3839ebde5e42cd771666d83f4e0af3b`;
PHP stdout59548bytes/SHA256
`eefdfe6c3ac5bd80eeaca24782446d823bde07c86566653d2840d2260f07f68c`.
Their metadata envelopes differ; this is not whole-stdout equality.

The depth cases concern JSON decode failure and its shallow `null` result.
They do not test successful values near the512 limit, observation depth over64,
or exact threshold/resource equivalence. An observer failure, process failure
or timeout cannot count as a matching PHP JSON error. JSON scalar roots, float/
overflow, arbitrary malformed inputs, original PCRE error/resource paths,
PHP8.4 and later application gates remain outside this series.

## Execution profile

Clean measured Gossip commit `62d9ea5573a88eb286b85b2904ca0fdf03020a4f`,
tree `ce553245b4a9ac67e13cdcad2139047681397250`. Harness SHA256
`f35f1e6106c223be55d82b1e196c3991505385ffd2a9363c1005b7b29451d379`,
corpus SHA256 `5ab1573c821adf04e120808f51708cc63b25e982988fd448539289e41c0713cb`,
pin SHA256 `5eeeaea21222bc2c29b45ad1e421460d29b412f2b500f0594da4e8f8476dd871`.
The measured harness/corpus/pin and both reused observers remain unchanged
after measurement; later edits only publish evidence/project status.

Use the existing checked generation/report35524a6 and unchanged GoHPb2d1482,
complete original Parser/Block/Frame and all prior pins/observers/evidence.
The harness checks hashes and no implementation diff between generation/runtime
commits; it does not run the compiler again. Exact Darwin/arm64 Go1.23.0,
clang/SDK15.5, static PCRE10.47_1 CGO1/taggohp_pcre2 candidate; parent CGO0.
Native fixed PHP8.3.36/PCRE10.42 Linux/aarch64, C locale, JIT0,
backtrack100000/recursion10000/process memory128M. Profiles differ.
Existing cached toolchain, image, running Docker and shared cache only;
no install/startup/pull, persistent binary or deployment is part of this test.
Candidate binary was removed by test cleanup. The duplicate-static-library
linker warning is recorded separately from runtime stderr0. No GoHP change/PR
or CI rerun was needed; prior successful configurations were not repeated.

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
| Status / common result | Experimental, measured only for this configuration; skip, failure and comparison PASS remain distinct; phases, hashes, typed observations, errors and counts are recorded; application acceptance not-established |
| WebMCP | `excluded`: local native compilation/container test, no browser/server operation or UI callsite; no browser/S2S/MCP/webhook/public route or authentication change |

Key/type boundaries belong to separate scope6521347316. This JSON-error scope
is6521354926, after superseded duplicate claim6521348620; no key/type work or
measurements were started under that claim. Devboard remains unidentified;
Issue1 and current-state records track progress.
