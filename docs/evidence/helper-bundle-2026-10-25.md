# Combined helper execution — 2026-10-25 JST

## Outcome and exact targets

**PASS:** ten unchanged WordPress functions and five owned PHP composition
drivers were translated together and built into **one native Go executable**.
All 22 typed observations matched native PHP and authored expectations, with
empty stderr and exit 0. This is a bounded integration result, not WordPress
startup or a rerun of the complete standalone input corpora.

- Gossip test commit: `577791ff93543d8b5abdc1efbc47af99f4bf407c`, tree
  `30a343150571ea85d7833148ce0347facb6f7d36`.
- GoHP compiler/runtime: `e3db685dfb2a3602f94bbfa9df105db3db96cc89`, unchanged
  from the final zeroise unit. Source archive hash
  `208a11f282a64522f3e1006007f8d1927bf0e916fd03f1148f23040cd10fc19e`.
- Original WordPress: `54106f6c545fc8b881f6b39b7963002bece6c28c`; source files
  `wp-includes/functions.php` and `wp-includes/formatting.php` remain unchanged.
- Actual PHP **8.3.36**, existing approved immutable image and same-process
  admission guard. `LC_CTYPE=C` selected and verified in the oracle process.
- Host Go 1.22.1 darwin/arm64; builder Go 1.25.14 linux/arm64.

The [machine-readable record](helper-bundle-2026-10-25.json) preserves both
complete source hashes, ten selected range/excerpt hashes and bundle offsets,
the manifest/driver/case/observer hashes, compiler/archive identity, generated/
report/binary hashes and every actual observation. Report hashes were verified
before building the program. The exact bundle contract is
[documented separately](../helper-bundle-unit.md).

## Execution

From `/Users/kureyamayuri/Desktop/Gossip-wordpress/gossip`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressHelperBundle$'
go vet ./tests/gohp
```

Initial run PASS: test **8.89 seconds**, package **9.096 seconds**. Vet PASS.
No failed attempt, code change or rerun after PASS; subsequent changes are
documentation only. No new compiler CI run was needed: the unchanged pinned
compiler already has fixed PHP 8.3.36 / Go 1.23.12
[CI 40118450143](https://github.com/katsuakikureyama/GoHP/actions/runs/40118450143)
PASS for all 78 catalog checks. That is prior compiler evidence, not a new CI
execution of this Gossip integration test. Gossip evidence here is local,
isolated Docker execution.

The ten direct probes cover each original function once, including actual
bool/int/string/null/empty-array returns. Twelve composition probes run five
owned PHP drivers that call the originals: eight slash-then-padding cases and
one each for empty-array-to-Boolean, zero-to-padding, false-to-Boolean and
empty-string-to-slash handling. These entry-probe counts are not counts of
every nested invocation. The driver bodies are compiled along with the
originals; the observers do not reproduce their algorithms.

The path composition corpus covers empty, slash/backslash-only and mixed
suffixes, UTF-8 byte width, NUL preservation, no truncation and width zero.
The direct Boolean input `FaLsE` exercises ASCII lowering and string handling.
This small integration corpus does not cover every earlier standalone value,
arbitrary byte strings, arbitrary arrays or coercion domains.

## Preserved boundaries

The four earlier standalone manifests, their compiler pins, cases, observers
and evidence remain unchanged and were not rerun. Their 73 typed observations
are not relabelled as passing in the combined executable. The combined result
belongs only to these new 22 cases. No WordPress original implementation,
GoHP source, dependency, native SDK/hook or oracle-policy change was needed.
The separate GoHP #17 stack and normal checkout remain untouched.

The change is stacked on Gossip Draft #9, not merged to the default branch or
WordPress upstream. No release, deployment, license finalization, HTTP server,
DB/theme/plugin integration or core startup is claimed. The local developer
capability stays WebMCP-excluded; no authentication route is added. No devboard
card has been linked, synced or marked Done. Next work should select a concrete
missing source dependency or runtime integration boundary; this PASS does not
justify treating the whole WordPress runtime as translated.
