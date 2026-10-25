# Twenty-four-original integration — execution record (2026-10-26 JST)

All **262 cases passed in one generated executable** containing24 unchanged
WordPress declarations and46 owned drivers. Actual typed return values and
caller input/selector state match PHP8.3.36 and authored expectations.
See the [contract](../edge-helper-bundle-unit.md) and
[full measured record](edge-helper-bundle-2026-10-26.json).

The measured262 comprise149 legacy,24 numeric predicates,7 array_all,12 numeric
compositions,30 search polyfills,8 retained search compositions,24 array edges
and8 new edge/search compositions. All seven retained corpora are hash-fixed;
this is a new full execution, not an addition of historical PASS counts.

## Revision and execution

- Successful test commit `c753c2d37ada08e2c412ea6996515f06dcf0b500`, tree `69e40a5525a299f7c193fe4f52eae96e9961aad5`.
- Original source baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`. Three whole-source hashes
  and all30 declaration/guard fragment hashes are retained in the JSON.
  Native uses complete guards and verifies all six polyfills are user-defined;
  compiler input retains declaration bytes/annotations but excludes loading guards.
- Compiler pin `440308c77da07cbcdce403fed7fff3945e0d2078` (GoHP #58 main).
  Reuse its [433-check CI](https://github.com/katsuakikureyama/GoHP/actions/runs/40219429355).
  No compiler modification or manual retest; #59 is not adopted by this manifest.
- Actual PHP8.3.36/C, builderGo1.25.14linuxarm64. Same-process admission,
  read-only/non-root/network-off execution and resource limits retained.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-empty-array.nKn0kC/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressEdgeHelperBundle$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

Comparison PASS9.28s/package9.489s, command exit0, vet PASS.
PHP and Go executions both exit0 with empty stderr. The final vet invocation
was corrected from the repository root (no go.mod) to the nested module.

The first snapshot `65f716e1c1968e5a50b33dab53387e8e341f38ad` translated all
declarations, then failed at Go build because the in-memory edge observer
adapter removed encoding/json while its case fields use json.RawMessage.
No differential comparison ran on that snapshot. The correction retains that
import and removes only os. Original source, cases, driver bodies and compiler
were unchanged. The corrected snapshot completed the first actual comparison.
After PASS only documentation/evidence is changed; no post-PASS repeat runs.

## Measured artifacts and scope

- Generated SHA-256: `ae8a82480f7aa38a969bc0e1b149969477c61cba1f356c4ae0237997599ce070`.
- Report SHA-256: `ea2fc59630da4ddbab892f38402e359cc618a41d6a24bb9dd38d2d9f8719f21b`.
- Binary SHA-256: `bb8dc939ce77cdbe9acf53b4a367b215549454c0a5e41b86220dd703442dc1cb`.

The JSON stores all262 observed typed results/caller states, corpus/driver/
observer hashes and exact in-memory observer adaptations. The raw local test
logs are retained in `/private/tmp/gossip-edge-bundle.bIElDy/test-output*.log`.

New8 cases compose wp_array_slice_assoc → Boolean mapping → array_first/last
and array_any/find/find_key. They distinguish empty/null from singleton false,
key0 from no-match null, selector order/deduplication, numeric-string and mixed
keys, first/last differences and positive/negative boundaries.

Earlier pins, corpora and evidence are preserved. The result covers the
contract's bounded domains; it does not establish WordPress startup/HTTP/DB/
theme/plugin compatibility or general callback/error semantics.
This is local Docker evidence, not a new GitHub Actions result.
