# Generated complete original parser at match1 — 2026-10-30

## Result and contract

The generated comparison **PASSes twelve cases/twenty-four actual original
parse calls**, with sixteen resource error2 and eight error0 calls matching the
retained PHP8.3.36 record. Fresh generation/build/execution/comparison each ran
once, test2.90s/package3.137s, mismatch0, candidate exit0/runtime stderr0.
No native PHP or previous successful configuration was rerun.

The [new pin](../../gossip/tests/gohp/parser-match1-generated.json)
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

Clean measured Gossip commit `79e61bab8a00f23f779c1ece2808ae3df3e27129`, tree
`c81848cacb3e699932f6441e2d6833a4af16113d`; GoHP140 main above, tree
`6d4efcadbaf3b04af33e2916dafe79411c472242`. The
[new machine record](parser-match1-generated-2026-10-30.json) retains actual Go
observations plus the exact expected native cases and adds source/command/
run_count1 provenance. Published SHA256
`fc5a4cc7467b5b6b1e5dbca119f9f9ba6850075573f5507d5586ee8133fdf794`,
raw SHA256 `772744dd76b9ad90875fd7ed8bde98fdbf5103b8e9d5c10055b039f311dfb089`.

| Identity / observation | Exact value |
|---|---|
| New generated Go | SHA256 `4c8648ef4b63f7862444b52c0c0f97c0898fc40bbe4f16698014ae7447f89842` |
| New report | SHA256 `ec3714aa49e36ecce45bacd9a0458c951833090d900f47b0a4f8932f4c54f3f6` |
| Pinned GoHP source archive | SHA256 `1f65e066c382932f5d63bd2ad5a1a1cb1ec8bba10e16e1c294f4b21135234a03` |
| Canonical typed cases, Go and retained PHP | SHA256 `4f4d9fcd79d53704981a9ae8e1c51f1387971818d21fe0223e1092a89884a728` |
| Go stdout | 688192bytes, SHA256 `d7ac33e6b6f87a9698afbbf838a217abd8b265dc5a5ccd611a9fcd4c04165369` |
| Retained native stdout | 688992bytes, retained in the earlier native-only record; different metadata envelope |
| PCRE errors | Both sides:0 eight times,2 sixteen times; original false/no-more-tokens freeform output matches |

The generated source/report remain outside the source worktree as data;
only byte-identical source is copied into the temporary candidate build.
No generated source rewrite, extracted method or alternate parser is used.
The binary is test-cleaned. Duplicate-library linker warning is retained
separately from runtime stderr0. Parent/frontend/IR/report source identities
remain explicit, and all three complete original classes are translated.

GoHP140 push41054740243/PR41054754537/exact-main41055413020 CI succeeds;
the push artifacts were fetched once and matched exact cleanheadb1d87a9/tree,
catalog and all eight new checks. Total866:8.3 848PASS+18not-measured;
8.4 regression864PASS+2not-measured. That CGO0 tooling evidence is separate
from this enabled original-parser comparison and never admits an8.4 profile.

## Capability

| Field | Contract |
|---|---|
| Stable ID | `gossip.tests.parser-match1-generated` |
| Service/resource | Local developer inspection of pinned original parser and retained native error-reference observations |
| Input schema | Opt-in `GOSSIP_GOHP_PARSER_MATCH1_GENERATED=1`, absolute GoHP source/PCRE prefix, fresh absolute evidence file and artifact directory; pinned compiler/profile/corpus, not arbitrary user input |
| Inspect / prepare | Verify original5source/pin/corpus/observer/native/profile hashes, generate a separate named-profile unit, verify report/source and compile enabled candidate |
| Save / commit | One exclusive `GOSSIP_GOHP_MATCH1_EVIDENCE_OUT`; new `GOSSIP_GOHP_MATCH1_ARTIFACT_DIR` retains only generated parser.go/report.json, temporary binary is test-cleaned; no application/external publish operation |
| Status / common result | Experimental/retained-comparison-pass for this configuration; skip/failure/generation/build/execution distinct, fresh native explicitly not-run, application acceptance not-established |
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
