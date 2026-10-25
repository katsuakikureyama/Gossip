# Thirteen-original recursive bundle evidence — 2026-10-25

Result: **PASS, case-limited**. Thirteen original WordPress functions and eleven
owned PHP composition drivers were translated/built together into one binary.
All33 actual return type/value and post-call argument observations matched
PHP8.3.36 and the explicit expected results/unchanged inputs. Empty stderr/exit0.
This is a new local isolated Docker configuration, not Gossip CI or core startup.

## Revision and measured configuration

- [Gossip Draft #15](https://github.com/katsuakikureyama/Gossip/pull/15), branch
  `codex/gohp-recursive-helper-bundle`, stacked on #14.
- Test commit `97ed80c3db187ce9ffb2a73c27dd5207f8594192`, tree
  `baef43b9bed5efd39b256d7fe705d93a133fafbf`.
- Compiler/runtime unchanged at
  `03d433e73503e32a84f310f023fc0f65c9b188cd`
  ([GoHP Draft #27](https://github.com/katsuakikureyama/GoHP/pull/27)),
  tree `4e94e4b6e826436a858bdd29fcc7a79af43ba8e9`.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`.
  Whole source hashes, original ranges, excerpt hashes/bundle offsets and
  driver/manifest/case/observer/compiler/archive/generated/report/binary hashes
  are captured in [machine evidence](recursive-helper-bundle-2026-10-25.json).
- The existing immutable PHP/Go images, allowlist, same-process native version
  guard, LC_CTYPE=C guard and non-root/read-only/no-network/resource limits are
  unchanged. ActualPHP8.3.36, builderGo1.25.14 linux/arm64; hostGo1.22.1 darwin/arm64.

## Corpus and result

Thirteen direct probes and twenty compositions. The prior24-case integrated
corpus is carried into this new configuration; nine additions cover direct
nested wp_slash and eight composition inputs (string roundtrip3, list shallow
strip3, capture/newline1, array shape/keys/values1). Every case now also observes
the actual caller arguments after invocation and requires their preservation.
Observers do not implement the escaping/stripping algorithm.

A flat mixed list roundtrips, while the nested-list shallow-strip case retains
the escaped inner string. This distinction is deliberate and matches PHP; it
does not establish recursive unescaping/wp_unslash support. Direct probes and
explicit intermediate-like array-shape output supplement roundtrips, so the
evidence is not solely two algorithms potentially cancelling each other's errors.

Working directory: nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressRecursiveHelperBundle$'
go vet ./tests/gohp
```

Initial run PASS8.19s/package8.387s; affected-package vetPASS. No failed test,
manual retry or code change after PASS; subsequent changes are documentation.
See [contract](../recursive-helper-bundle-unit.md) for the exact transport and
capability boundaries. Earlier seven standalone units and ten-/twelve-function
bundles retain their pins/corpora/results and were not rerun. This new bundle
does not claim to revalidate every prior standalone input.

## Reused compiler evidence

No compiler/runtime/configuration changes were made. Reuse
[PR CI40152066817](https://github.com/katsuakikureyama/GoHP/actions/runs/40152066817):
PASS181 checks across42 catalog entries, actualPHP8.3.36/Go1.23.12 linux/amd64.
Clean checkout `60d0b142b55d8654aa19910106c778d1102cbc81` has the exact tree of
the compiler pin. Artifact12582837295 / compatibility-php8.3-40152066817-1,
downloaded at `/private/tmp/gohp-array-map-ci.fV1ibO/`, expires2026-11-24.
Metadata/digests are retained in machine evidence with reused=true/rerun=false.
The earlier compiler local test failure/fix history remains in
[wp_slash evidence](wp-slash-2026-10-25.md), not relabelled as a new run.
Compiler native schema compares stdout/normal exit only, not general
stderr/exceptions/side effects.

## Unverified and next dependency

Original wp_unslash (formatting.php5964–5966) → stripslashes_deep (2918–2920) →
map_deep (5268–5284) needs array element assignment, object/property mutation,
is_object/get_object_vars and dynamic call_user_func. Those are not implemented
by this integration. No original body is substituted with an array-only helper.
A next compiler change should isolate a specific missing contract (for example
array writes with alias/key rules) before claiming support for map_deep.

No associative JSON transport, floats, objects/references/cycles, arbitrary
large-integer precision/depth, whole PHP files/core startup, HTTP/DB/themes/plugins
or SDK/hook integration claim. Original WordPress files and the other session's
GoHP checkout are unchanged. No default-branch merge, deployment, release or
license decision. Local developer capability remains WebMCP-excluded; no auth
route change. Devboard card unidentified/unsynchronized, not Done.
