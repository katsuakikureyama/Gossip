# Original array_first — execution record (2026-10-26 JST)

All **12/12 cases passed** against actual PHP8.3.36. The original declaration
translated, built and executed with GoHP main #57. Actual return types/values
and ordered caller inputs after the call matched PHP and authored expectations.
The [contract](../array-first-unit.md) defines the narrow input domain.

- Test commit: `fdf018beee72a4edb5e011ee8c3ed7e0afb0bae1`.
- Test tree: `66757ce66aece6ec82cb29e3558b77d94eb95f85`.
- Compiler pin: `4a4fffb995e57daf592785b6626e45a7babe08da` (GoHP main,
  [#57](https://github.com/katsuakikureyama/GoHP/pull/57)).
  Its [CI427-check PASS](https://github.com/katsuakikureyama/GoHP/actions/runs/40218487439)
  is reused, without compiler modification or manual retest.
- Declaration657–665 and native guard646–666 of `wp-includes/compat.php`
  are byte-unchanged. Whole-source, selected-fragment and bundle hashes are in
  the [machine record](array-first-2026-10-26.json).
- Exactly the12 earlier array_first reference cases, including empty/null/falsy
  values, ordered sparse/mixed keys, numeric-string normalization and Unicode/NUL.
  Reflection confirmed the native user-defined polyfill.
- Fixed builder: Go1.25.14 linux/arm64. PHP8.3.36/C locale, same-process
  admission, network-off, non-root, read-only and resource restrictions retained.

Run from the nested `gossip/` module:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-empty-array.nKn0kC/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArrayFirst$'
go vet ./tests/gohp
```

First attempt PASS14.95s (package15.187s), vet PASS, command exit0.
Native and Go stderr were empty and both execution commands exited0.
No unexpected failures or retries. The committed test snapshot was clean;
post-PASS edits concern documentation/evidence only, with no duplicate test run.
This is local Docker evidence, not a new GitHub Actions result.

Generated Go SHA-256: `06741abc537d97c2b000d30c63f27d3c090e2ddcf4ee06d055b1d88f8aeaac9a`.
Report SHA-256: `79506cd71046d83a761955ed35374e3e3d203ca42ae94a84ddc292867aaf4fac`.
Binary SHA-256: `94e4a4e83d97c251b778b7210ed262e9ae76c2562a56c43cfdd2c482c89cbf62`.

This adds one independently compared original. The existing22-original/230-case
single-binary integration stays at its own pin and scope;12 cannot be added to230
as evidence of a larger integrated configuration. Historical rejected boundary
records remain intact. array_last still requires array_key_last. Broader value
domains and WordPress startup/HTTP/DB/themes/plugins remain unverified.
