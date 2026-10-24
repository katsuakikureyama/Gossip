# Original Boolean-helper observations — 2026-10-25 JST

## Targets and command

- Tested Gossip: `8901d634e4d52e9be06c4c2d7f91714652974b28`.
- Original WordPress: `54106f6c545fc8b881f6b39b7963002bece6c28c`.
- GoHP compiler/runtime: `14bb01053947f1bf2fb32edaff8085feb7227f48`,
  [Draft PR #13](https://github.com/katsuakikureyama/GoHP/pull/13), stacked on #10.
- Host Go 1.22.1 darwin/arm64; isolated builder Go 1.25.14 linux/arm64;
  oracle PHP 8.3.13, explicitly selected and verified `LC_CTYPE=C`.

From `/Users/kureyamayuri/Desktop/Gossip-wordpress/gossip`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressBooleanHelper$'
```

PASS: 8.38 seconds (package 8.557 seconds). All 23 actual return types and
values agree between PHP, generated Go and independently authored expectations,
with empty stderr and exit 0. See the [full hash/value record](boolean-helper-2026-10-25.json).
No code/config changed or test rerun after this PASS; subsequent changes are docs.

The initial run at Gossip `018a3841d040d68e40ebc4deda0d7af6ecbe9861`, using
GoHP `b994e23b774b1fa04ea1b60b0a7ecf867839d3db`, failed after generation and
compilation (7.32 seconds, package 7.522 seconds). The PHP observer rejected
its startup `LC_CTYPE`; setting `LC_ALL=C` alone did not establish the required
PHP locale. The observer now explicitly sets and verifies `C` in the process
that calls the original helper. GoHP's native harness received the same fix;
compiler/frontend/runtime implementation bytes did not change. This is a
test-environment correction, not a rewrite of WordPress behavior.

## Observations and limits

The exact original `wp_validate_boolean()` body at functions.php 7932–7942 is
extracted without substitutions. Cases cover bool, null, signed integers,
strings (including `false` variants, surrounding spaces, `0` versus `00`,
Japanese and NUL), empty/nonempty and nested arrays. Input types are checked
after JSON transport; array inputs use actual `rt.Array` values. Observers call
the actual original/generated functions, not copied algorithms.

Generation/report hashes are checked before compilation with the same pinned
runtime. Containers have no network, non-root UID, read-only root, restricted
mounts, capabilities and resource limits. JSON decoding, fixture reads and
typed observation are instrumentation, not translated JSON/file APIs.

This unit is separate from previous six constant-return and two slash helpers;
their pins and evidence remain unchanged and were not rerun. No combined nine-
function validation is claimed. Float/object/resource inputs, weak coercion,
non-C PHP 7.4 locales, array identity, general conditionals, full-file/core,
HTTP/DB/theme/plugin behavior and production use remain unverified/unsupported.
See the [unit contract](../boolean-helper-unit.md).

## Compiler checks

At implementation commit `b994e23b774b1fa04ea1b60b0a7ecf867839d3db`, cached
Go 1.23.0 darwin/arm64, all of the following PASS:

```sh
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./internal/codegen ./internal/compat
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./runtime -run '^(TestBool|TestIdentical|TestStrToLower)'
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./tests -run '^TestUnsupported$'
```

Package times: codegen 0.237, compat 0.443, runtime 0.191, tests 0.195 seconds.
These unchanged implementation checks were not repeated locally after the
native-oracle-only fix. Fixed PHP 7.4.33/8.3.30 CI status is recorded below.

At the pinned `14bb010` commit, [CI 40113406271](https://github.com/katsuakikureyama/GoHP/actions/runs/40113406271)
ran both targets: raw Go test events contain no failed test and all seven
native differential fixtures, including `booleans`, passed. However, the CI
job failed during summary validation: the strict versioned `GOHP_CASE` parser
rejected the new `php_lc_ctype` field, so `diff.booleans` was incomplete. That
job is **not** recorded as CI PASS. Earlier CI 40113353477 also failed.

Compiler head `4fabe11d2063b03fced8d0aabd62bbd20a10fb88` fixes only this
evidence-format mismatch: it leaves the existing JSON schema intact and logs
the selected/verified locale separately in the archived raw test events.
It does not change frontend, semantic analysis, code generation or runtime.
The Gossip pin remains `14bb010` to retain the exact successful original-source
observation; no post-PASS repinning or redundant Gossip execution was done.

Final [PR CI 40113520430](https://github.com/katsuakikureyama/GoHP/actions/runs/40113520430)
PASSed both PHP 7.4.33 and 8.3.30 with Go 1.23.12 linux/amd64. Raw events
record the same-process `LC_CTYPE=C` check and matching `booleans` output;
summaries record `diff.booleans`, individual syntax/arity/domain rejection
checks and three runtime boundaries as PASS. Vet, test execution, summary and
artifact upload succeeded. No extra rerun after PASS.

Actual clean PR merge checkout: `6d255dc5c27f7062fcf5f644b754a91f2c0f6c0f`;
tree `cf61000eeef37da5a2a200b36e95243e61b63d11` equals compiler head.
Artifacts (30-day retention): `compatibility-php74-40113520430-1` ID12567557182,
`compatibility-test-40113520430-1` ID12567343135. This compiler matrix does
not establish WordPress PHP 7.4 compatibility or general PHP support.

Both changes remain Draft/unmerged. No deployment, release, license finalization
or upstream WordPress change. Local developer tooling is WebMCP-excluded; no
new authentication route or user-facing operation. No devboard card is linked
or marked Done.
