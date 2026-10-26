# Real block element-class decision: integrated334-case PASS

2026-10-26 JST. [Contract](../elements-class-path.md),
[all334 measured observations/hashes](elements-class-path-2026-10-26.json).

**The unchanged real caller `wp_should_add_elements_class_name` now executes
with unchanged `_wp_array_get` and all retained helpers in one Go binary.**
All334 fresh comparisons match fixed PHP8.3.36/C and expectations:294 retained
plus40 caller cases. No resolver, class, service or PHP behavior stub replaces
the caller or its function dependency.

## Exact verification

- Branch `codex/gohp-elements-class-path`, base master9abd4b8.
- Clean tested commit `32ac5e0060574c2f250c601d8e13d1bbe48cf1e9`, tree
  `71731182f2193e9c997c6519a098aa5464cf7fd9`.
- GoHP #67 main `a3039414cff22689d214a6c90767f13dccf56ddb`, same tree
  `4a1b01fbbb7ab2fdd245bd3ffa6255241e5f3e81` as tested feature32adaa2.
  Fixed PHP8.3.36/Go1.23.12 amd64
  [CI40257395885](https://github.com/katsuakikureyama/GoHP/actions/runs/40257395885)
  artifact confirms clean exact head/tree and all488 checks PASS (66 features).
  [Main CI40257522353](https://github.com/katsuakikureyama/GoHP/actions/runs/40257522353)
  also PASS. Consumer verification reuses this compiler result, no manual rerun.
- Original caller: block-supports/elements.php33–97, byte-unchanged, whole-source
  hash bound. Total four source files/32 fragments,26 originals/48 owned drivers,
  ten unchanged retained corpora. One translation/build, one executable.

From the nested `gossip/` module:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-nested-array-probes GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressElementsHelperBundle$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

**334/334 PASS9.84s/package10.082s**, vet PASS on the same snapshot. Both runtime
exits0 and stderr empty. Fixed digest builderGo1.25.14 linux/arm64 and PHP8.3.36/C,
nonroot/read-only/network-off/resource limits and same-process admission retained.
The first committed snapshot passed; only documentation/evidence follows. Raw
log `/private/tmp/gossip-elements-class-path-output.log`, no repeated PASS run.

## Caller behavior measured

Each of26 original configured color paths is tested independently (button3,
link2, heading3, h1–h6 three each). Fourteen more cases exercise missing
attrs/style/elements, null intermediate/leaf, empty/unrecognized elements,
nonnull false/zero/empty string/empty array, skip true and skip strings `'0'`
versus `'false'`. Presence is **non-null**, not truthiness or valid CSS. Every
caller result is an actual bool, and actual block/options states remain unchanged.

Native Reflection verifies the original user-defined function. Go invokes
`F_wp_should_add_elements_class_name`; expected fields only assert the result.
Exact-once adapters retain current measurements from the294-case integration.
All334 results and caller states must agree with both current PHP and expected
records. Old pins, source, corpora and evidence are not rewritten.

## Compiler work

[GoHP #67](https://github.com/katsuakikureyama/GoHP/pull/67) adds defined-root
multi-level isset and lazy address-only `??`, preserving old one-level guards
and normal missing-key reads. Reached intermediates must be arrays/null; fallback
is lazy, non-null falsy values stay present, and ?? is right-associative below ||
and above ternary. Local tests/native fixture PASS on a0b2228; independent main
#66 array_slice was retained by merging and unioning only catalog/fixture lists.
The combined32adaa2 snapshot then passed488 CI checks. No local successful
configuration was manually repeated; exact commands/results are in the PR.

## Limits

This is the real **class-needed decision**, not the surrounding render hook,
CSS/HTML generation, unique class identifiers, Theme JSON/resolver loading,
WordPress bootstrap/HTTP/DB/themes/plugins or production operation. Malformed
scalar/string/object intermediates, general errors/coercions and undefined
variable suppression are excluded. Parsed block/options roots are arrays,
intermediates in the admitted corpus arrays/null; nested codec limits remain
depth12/256 entries/4096-byte strings/int±1e6. No deployment target was invented.

Capability `gossip.dev.wordpress-unit.translate-compare` stays explicitly
WebMCP-excluded with no browser/server operation or route/auth change. Devboard
unidentified/not synced, tracked in PR/current-state. Parent owns design/core/
harness/review/verification/integration; workers supplied independent tests,
fixtures/observers and scoped documents.
