# Original array edges — execution record (2026-10-26 JST)

All **24/24 cases PASS** in one generated Go executable:12 array_first and12
array_last. Actual scalar return types/values and ordered caller inputs after
each call match actual PHP8.3.36 and authored expectations.
See the [contract](../array-edge-unit.md) and [machine record](array-edge-unit-2026-10-26.json).

- Test snapshot `ecb89beed327dd14cd3772bfdb9e752074816a0f`, tree `ed72977216cb8b70747b14a8c88804db74124507`.
- GoHP main `440308c77da07cbcdce403fed7fff3945e0d2078` (#58).
  Reuse [CI433-check PASS](https://github.com/katsuakikureyama/GoHP/actions/runs/40219429355);
  compiler source/configuration is not changed or manually retested.
- Original declarations array_first657–665 and array_last679–685 are byte-unchanged.
  Native executes complete original guards646–666 and668–686; Reflection verifies
  both functions are user-defined. Conditional loading is not compiled.
- Existing24 cases, PHP observer and numeric transports remain unchanged.
  New observer dispatches directly to generated originals. No owned PHP drivers.
- Actual PHP8.3.36/C locale; builder Go1.25.14 linux/arm64. Network-off,
  read-only, non-root, resource limits and same-process PHP admission retained.

Commands from nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-empty-array.nKn0kC/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArrayEdgeUnit$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

First comparison attempt **PASS11.16s/package11.364s**, command exit0.
Both candidates exited0 with empty stderr. Vet PASS. Initial vet used the default
host cache and could not access it under the sandbox; only vet was rerun with
the writable cache shown above. No implementation/test change or differential
retry. Staged diff check and clean commit preceded the comparison.

Generated SHA-256: `13022943e1f5ae8331a44e256186d9c0f1eb322b455742af7ff02c4a7e3e3f31`.
Report SHA-256: `3c58511089b71ee29d71e68c3fa653919b19ef6e62f76759b331fb98da7d6d6b`.
Binary SHA-256: `d69477652b59bdb1e041eb060758a4a23401e427210fd35244c590c53053c719`.

The JSON includes actual24 typed observations, ordered inputs, all fragment,
runner, compiler, bundle and observer hashes, commands and status.
After PASS only documentation/evidence is changed. Existing PASS configurations
are not rerun. This is local Docker evidence, not a new GitHub CI result.

This is a separate two-original configuration. The old #54 rejection record,
#57 array_first comparison and22-original230-case integration retain their own
pins and evidence. A24-original254-case integrated binary has not been tested.
Nested arrays, objects, references, error equivalence and full WordPress startup/
HTTP/DB/themes/plugins remain outside this result.
