# Original nested settings lookup:24-case Go/PHP PASS

2026-10-26 JST. [Machine-readable actual observations](array-get-unit-2026-10-26.json)
and [contract](../array-get-unit.md). The declaration and24 native-reference
cases are unchanged. This is a newly measured one-original comparison, not a
25-original286-case combined integration or full WordPress compatibility.

## Exact successful snapshot

- Branch `codex/gohp-array-get-unit`, base Gossip master67984d7.
- Tested commit `ce5318d414d0701f6d1d312d7d81ddbaeb4e514d`, tree
  `1c67b687f6ede81a061ba800c891b44eb4cbb213`; clean before execution.
- GoHP #64 main `b8dc1757ed61f459d51362da004aa1d73d024116`, same tree as
  clean tested feature head18c4e67. Fixed PHP8.3.36 CI
  [40227365611](https://github.com/katsuakikureyama/GoHP/actions/runs/40227365611)
  passed all473 checks across64 catalog features. Earlier #63's470-check result
  is retained separately; compiler tests were not manually repeated for Gossip.
- Original functions.php5233–5272, full-source/fragment/corpus/observer/compiler/
  generated/report/binary/archive/image hashes are in the JSON record.
- Nested `gossip/` module command:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-null-identity GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArrayGetUnit$'
```

**24/24 differential PASS9.12s; package9.359s.** PHP8.3.36/C, builder
Go1.25.14 linux/arm64. Both executions exit0, stderr empty. Actual scalar/nested
array return types/values and all three caller argument states match PHP and
expectations. Two-argument calls really omit the default; observers never return
expected fields instead of invoking the helper. Existing nonroot/read-only/
network-none/resource limits and same-process PHP admission are retained.

`GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp`: PASS on initial
snapshotd2bc85ce6. The subsequent fix changes only the manifest compiler pin and
documentation, so this unchanged-Go-source vet result is reused. Only docs and
evidence follow PASS; no repeated successful comparison.

Raw success log `/private/tmp/gossip-array-get-unit-fixed-output.log`.

## Initial failure and bounded compiler correction

Snapshot `d2bc85ce66fa19146da47eaad6677c81bfd9e362`, tree
`7de097a3f8a711ff343deb94f88d052859734b59`, initially pinned #63 main
`1f10764ce7afec18276cd5b3da9662db575be7df`. Generation/build succeeded but
execution failed with `panic: strict equality outside supported scalar subset`
at `null === $path_element` for an array-valued path. FAIL8.42s/package8.703s;
comparison incomplete. Raw log `/private/tmp/gossip-array-get-unit-output.log`
is embedded in JSON. No source/case/observer workaround was applied.

[GoHP #63](https://github.com/katsuakikureyama/GoHP/pull/63) had supplied bounded
null defaults and bare/multiple isset while reusing #62 continue.
[GoHP #64](https://github.com/katsuakikureyama/GoHP/pull/64) fixes literal-null
strict equality/inequality as once-only canonical-null predicates. General
non-scalar identity remains unsupported. Its first CI failed an obsolete
callable-versus-null negative test; that negative now uses false, retaining the
old guard. Test-only correction and all473-check success are recorded in the PR.

## Remaining boundary

The previously verified24-original262-case bundle and old pins/evidence remain
unchanged. This adds a25th successful original in a separate unit. A combined
25-original configuration has not run. Theme JSON/global styles/block callers,
HTTP/DB/themes/plugins, general coercion/errors/objects/references/floats and
production operation remain unverified. No deployment target was invented.

Capability `gossip.dev.wordpress-unit.translate-compare` is explicitly
WebMCP-excluded, with no route/auth change. Devboard unidentified/not synced;
PR/current-state tracks this work. Parent owns core/design/integration/review;
workers supplied independent regression fixtures and the Go observer only.
