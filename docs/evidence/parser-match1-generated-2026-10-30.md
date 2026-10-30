# Generated complete original parser at match1 — 2026-10-30

## Contract and initial status

Generated comparison is not yet measured. The [new pin](../../gossip/tests/gohp/parser-match1-generated.json)
selects merged [GoHP140](https://github.com/katsuakikureyama/GoHP/pull/140)
commit `e6a71a5f267658e8c575ffab5db6712246e543b2` and the explicit compile-time
profile `pcre2-php83-match1-v1`. The new
[harness](../../gossip/tests/gohp/parser_match1_generated_test.go) fresh-generates
the complete unchanged Parser/Block/Frame closure, verifies the distinct report
profile/runtime call, builds with the existing enabled engine and compares its
twelve documents/twenty-four actual parse calls against the retained
[match1 PHP reference](parser-resource-boundaries-2026-10-30.md).

The previous [native record](parser-resource-match1-2026-10-30.json) remains
unchanged at SHA256
`59d3e4cfe3553677f743b93d571966ffd538be9cbe1a0610c0788cc8488df60c`:
PHP8.3.36/PCRE10.42 Linux/aarch64/C, backtrack1/recursion10000/JIT0/process128M,
eight cases/sixteen calls at error2 and eight calls at0. Native PHP and Docker
are **not run** by this test. Native-only historical status is not overwritten.
The unchanged Go full observer, resource corpus, old pins/observers/source and
every prior successful configuration are reused as identities, not rerun.
This new profile needs its own fresh generation; the old default-profile
generated source is not edited or repurposed.

The candidate uses existing Go1.23.0Darwinarm64/clang/SDK15.5/staticPCRE10.47_1,
CGO1/taggohp_pcre2, fixed Match1/Depth10000/Heap131072KiB/Subject4096. Parent
harness/generation use CGO0/offline/shared cache. These engine/platform/process
budgets remain distinct from native PHP. Exact typed ordered cases, raw bytes,
parser fields, JSON/PCRE errors, caller preservation and fresh/reused parser
behavior are compared; whole stdout envelopes are not the equality contract.
Per-call JSON success baselines remain caller-established, not global-reset
proof. No PHP8.4 profile, arbitrary limits/ini support, heap/process equivalence
or whole application acceptance.

## Capability

| Field | Contract |
|---|---|
| Stable ID | `gossip.tests.parser-match1-generated` |
| Service/resource | Local developer inspection of pinned original parser and retained native error-reference observations |
| Input schema | Opt-in `GOSSIP_GOHP_PARSER_MATCH1_GENERATED=1`, absolute GoHP source/PCRE prefix, fresh absolute evidence file and artifact directory; pinned compiler/profile/corpus, not arbitrary user input |
| Inspect / prepare | Verify original5source/pin/corpus/observer/native/profile hashes, generate a separate named-profile unit, verify report/source and compile enabled candidate |
| Save / commit | One exclusive `GOSSIP_GOHP_MATCH1_EVIDENCE_OUT`; new `GOSSIP_GOHP_MATCH1_ARTIFACT_DIR` retains only generated parser.go/report.json, temporary binary is test-cleaned; no application/external publish operation |
| Status / common result | Experimental/unverified until measured; skip/failure/generation/build/execution/retained comparison are separate, fresh native explicitly not-run, application acceptance not-established |
| WebMCP | `excluded`: local compiler/native test, no browser/server operation/UI callsite/auth/route/deployment change |

Run only this new setting from `gossip/` with the explicit inputs above:

```sh
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserMatch1Generated$'
```

Scope6522495611. GoHP140's default-CGO0 tooling and enabled tiny owned-regex
test are separate evidence; its8.4 regression run does not admit an8.4 profile.
Original source/old manifests/corpora/observers/pins/evidence and BOOT75's
native baseline/source rejection stay intact. G83-PARSER resource host guards,
timeouts and other resource paths remain open, as does application acceptance.
No new dependencies/images/daemons/cache or persistent executable.
