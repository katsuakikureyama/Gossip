# Unified PHP 8.3 compiler observations — 2026-10-25 JST

## Pinned targets

- Gossip test commit: `e943cbce1f74c93b5ca990de7a544caabb274aa5`.
- GoHP compiler/runtime: `4e1b5abfe019ec121e13d85e99e51c206e765533`, tree
  `69d7c6bbb86ddb4a76a62073e2bed8f5c395c7d0`.
- Original WordPress: `54106f6c545fc8b881f6b39b7963002bece6c28c`.
- Native oracle: PHP 8.3.36, existing approved immutable image and admission
  guard unchanged from the [previous floor migration](php-83-floor-2026-10-25.md).
- Host Go 1.22.1 darwin/arm64; builder Go 1.25.14 linux/arm64.

[GoHP Draft PR #16](https://github.com/katsuakikureyama/GoHP/pull/16) is stacked
on #13. It backports only policy commit `948647c2903a4a1b22069818861062df1f5bf414`
from the separate #14 stack. Conflicts preserve legacy arrays, restricted rtrim,
Boolean control and PHP 8.3 precedence; new compilation rejects 7.4 and unverified
series. The runtime and code generator implementation are unchanged from #13.
Composer inventory, strlen, array_keys and array_values implementations are not
imported. Catalog: 31 entries, 73 distinct checks, active target 8.3 only.

## Gossip execution

From `/Users/kureyamayuri/Desktop/Gossip-wordpress/gossip`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP \
  go test -count=1 -v ./tests/gohp
```

First run PASS, package **34.204 seconds**, with all three units using the same
compiler archive SHA and executable SHA. Report/generated hashes are checked
before build, and actual PHP/Go values match the owned expectations:

| Unit | Observations | Time |
| --- | --- | --- |
| Constant-return helpers | Six typed values including empty array; echo for five scalars | 13.37s |
| Slash helpers | 13 string inputs / 26 typed returns | 8.64s |
| Boolean helper | 23 mixed inputs / typed bool returns | 12.00s |

All require empty stderr and exit 0. The [full machine-readable record](unified-compiler-2026-10-25.json)
contains unchanged source/case/observer hashes, generated/binary hashes, actual
PHP version, compiler identity and every observation. These are **nine original
functions in three separately built units**, not one combined source bundle,
full-file translation or WordPress startup.

Only the three compiler pins changed in Gossip implementation. WordPress source,
test input, expected values, observers, PHP image/guard and native SDK/hooks did
not change. The hooks/admission results from the previous exact effective
configuration are retained, not rerun or relabelled as new execution. Older
per-compiler and old-oracle evidence files remain untouched. No test was rerun
after PASS; subsequent Gossip changes are documentation only.

## Compiler verification

At the GoHP commit above, cached Go 1.23.0 darwin/arm64:

```sh
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./internal/... ./cmd/...
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go vet ./...
sh -n scripts/check.sh
```

All PASS. Package times: codegen 0.280s, compat 0.471s, phpversion 0.635s,
CLI 1.749s, compatibility CLI 2.918s, view CLI 0.923s. Frontend/IR/semantic
packages have no standalone tests. Fixed native CI runs the complete configured
compiler test set under PHP 8.3.36, including the unchanged bounded fixtures.

[PR CI 40116685783](https://github.com/katsuakikureyama/GoHP/actions/runs/40116685783)
PASSed on fixed Go 1.23.12 linux/amd64 / PHP 8.3.36. All 73 catalog checks
PASS, including the four policy checks, `diff.legacy_arrays`, `diff.rtrim` and
`diff.booleans`; no live PHP 7.4 row is present. Vet, native test execution,
summary publication and artifact storage succeeded. The clean PR merge checkout
`24cc25db0a03c89b9f177ca36edf3902b7d33c7e` has tree
`69d7c6bbb86ddb4a76a62073e2bed8f5c395c7d0`, matching the pinned compiler head.
Artifact: `compatibility-php8.3-40116685783-1`, ID12569685340, 30-day retention.
No extra rerun or code change after PASS. This compiler CI is separate from
the local Gossip test and does not claim general WordPress compatibility.

## Limits

PHP 7.4 and below-floor historical archives remain readable but cannot approve
new execution. Admission is not proof of general PHP compatibility or a
vulnerability-free interpreter. Arbitrary arrays, coercion, warnings/exceptions,
HTTP/DB/theme/plugin behavior and full WordPress remain outside this result.
Future feature work starts from this unified policy-aware pin rather than
combining unrelated old results.

[Gossip Draft PR #8](https://github.com/katsuakikureyama/Gossip/pull/8) is stacked
on #7. No default branch merge, release, deployment, upstream WordPress change,
license finalization or new authentication route. Local developer capabilities
remain WebMCP-excluded; no devboard card is linked or marked Done.
