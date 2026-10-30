# Original-parser carry-in state — 2026-10-30

**PASS for four fixed sequences / twenty actual original `parse()` calls per
backend. G83-PARSER and whole Gossip acceptance remain open.**

The [machine record](parser-carry-state-2026-10-30.json) preserves both complete
observations, hashes, source commit, command/environment and a single measured
attempt. This is a separate series from the unchanged twelve-case/24-call
[fresh-v1 result](parser-full-fresh-2026-10-30.md); its denominator is not added
to that corpus. No GoHP production change was needed.

## Observed behavior

Each row runs once with a fresh parser for the whole sequence and once with a
parser reused across all four sequences. All generated calls share one
`Execution`, corresponding to PHP process-global JSON/PCRE state. Inputs are
decoded before any seed; observation JSON is encoded only after all calls.
State is seeded only at the beginning of a pass, never between its documents.

| Sequence | Actual incoming state | Document sequence | Observed JSON / PCRE state after each call |
|---|---|---|---|
| `syntax_carries_without_decode` | JSON4 / PCRE0 | empty, freeform, void block | 4/0 → 4/0 → 4/0 |
| `successful_decode_clears_carry` | JSON4 / PCRE0 | nested attributes, empty | 0/0 → 0/0 |
| `syntax_is_preserved_then_replaced` | JSON0 / PCRE0 | malformed attributes, freeform, nested attributes | 4/0 → 4/0 → 0/0 |
| `regex_error_resets_independently` | JSON4 / PCRE2 | empty, nested attributes | 4/0 → 0/0 |

JSON messages match (`Syntax error` / `No error`). Non-decoding parses preserve
the earlier JSON error. A successful attribute decode clears it. The original
regex execution clears the incoming PCRE error independently of JSON state.
Both passes match typed results, ordered raw-byte keys/values, every parser
property and before/after errors; caller documents are unchanged, output equals
the return value and the parser stack is empty after every call.

Seeds are real runtime operations, not writes to private state: `json_decode`
of `{` or `null`, followed by `/a/s` against an empty subject or
`/^(a+)+$/s` against forty `a` bytes plus `!`. The latter returns false/error2
under each fixed profile. **That is a caller seed, not a resource failure of the
original WordPress pattern.** The original regex-error/resource gate is still
unmeasured. This series covers JSON0/4 and incoming PCRE0/2 only.

## Identity, execution and evidence

- Clean measured Gossip commit `ed13d612505ad2241031dcba05c2b07e6c024434`, tree
  `7d197731427c6d1f3f0ec3da59bf65b5271db834`; GoHP
  `b2d1482268a7710c515c07bf919af3d40b09d17c`.
- Complete original Parser/Block/Frame, notices and WordPress provenance retain
  all five source hashes. Native reflection verifies the original methods and
  `parse` entry lines64–75. No extraction, source rewrite or substitute parser.
- New [pin](../../gossip/tests/gohp/parser-carry-state.json),
  [sequences](../../gossip/tests/gohp/parser-carry-state-cases.json),
  [PHP observer](../../gossip/tests/gohp/parser-carry-state-observer.php),
  [Go observer](../../gossip/tests/gohp/parser-carry-state-observer.go.txt) and
  [harness](../../gossip/tests/gohp/parser_carry_state_test.go) have separate
  immutable identities. Existing documents and all prior pins/observers/evidence
  remain unchanged.
- Existing generated Go/report from `35524a6` are hash-checked and reused after
  proving no compiler/runtime implementation diff through `b2d1482`. No new
  transpilation or old successful retained/fresh-v1 test is run.
- One test/build/candidate/native configuration, no failure or rerun:
  `TestOriginalWordPressParserCarryState` PASS2.60s / package2.806s, mismatch0.
  Generated stdout40598bytes; native stdout40880bytes. Both exit0/stderr0.
  Comparison is canonical typed `scenarios`, **not whole stdout equality**.
  Both scenario hashes:
  `c84456389053f42239b32b81949346b128761b328bd1e5b74a9117b92c967c0f`.
- Raw attempt SHA
  `8b1ee484b5cb5f30211cce4bb3b55bac7799b09b702e78d31a954e2cc2d9b8b7`.
  The published JSON adds committed-source/command/publication metadata to the
  same observation; it is not another execution.

Run from `gossip/` with the exact environment recorded under `commands.environment`
in the machine record:

```sh
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserCarryState$'
```

The harness requires explicit opt-in `GOSSIP_GOHP_PARSER_CARRY_STATE=1`, absolute
source/prefix/generated-directory/fresh evidence paths and the exact cached PHP
image. Missing opt-in skips without recording support. Offline cached Go1.23.0
runs the test with CGO0; the child candidate explicitly uses CGO1,
`gohp_pcre2`, verified clang/SDK15.5 and static PCRE10.47_1 on Darwin/arm64.
The native side is fixed PHP8.3.36/PCRE10.42 on Linux/aarch64, C locale,
JIT0/backtrack100000/recursion10000/process memory128M. These engine/platform
profiles differ; equal nominal limits do not prove equivalent failure points.
The duplicate static-library linker warning is retained separately from quiet
runtime stderr. Temporary candidate binaries are removed by test cleanup.
Existing Docker25.0.3/image/engine/cache only; no daemon startup, pulls,
dependency installs, persistent binaries or deployment.

## Capability and remaining scope

| Field | Contract |
|---|---|
| Stable ID | `gossip.tests.parser-carry-state` |
| Service/resource | Local developer test; pinned original parser and immutable observations |
| Input schema | Opt-in `1`, absolute GoHP source, PCRE prefix, checked generation directory, exact PHP image, fresh absolute evidence output; the pinned JSON corpus is not user-selectable |
| Inspect / prepare | Verify source, manifests, input/observer hashes and existing build profile |
| Save / commit | Write one exclusive local evidence file; no application save, publish or external commit operation |
| Status / result | Experimental, measured only for this configuration; PASS/FAIL/skip distinguished, phases/hashes/errors/counts recorded, application acceptance remains not-established |
| WebMCP | `excluded`: developer-local native compilation/container test, no browser/server operation or UI callsite; no browser/S2S/MCP/webhook/public route or authentication change |

Only the [carry-in gate](../parser-full-acceptance-boundaries.md) is now passed
for this case-limited series. JSON key/type and encoding/control/surrogate/depth
boundaries, the original PCRE resource-error path and other application gates
remain open. No general globals, nonstring inputs, float/overflow, arbitrary
public-method calls, references, PHP8.4 or full WordPress/Gossip acceptance is
established. Devboard is unidentified; tracking is Issue1 and the current-state
record.
