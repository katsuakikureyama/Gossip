# Combined original-helper experiment

This is a new integration configuration, not a rerun or replacement of the four
independent helper units. Ten unchanged WordPress functions from the pinned
`functions.php` and `formatting.php` are extracted into one source bundle,
translated once with GoHP `e3db685dfb2a3602f94bbfa9df105db3db96cc89`, and built
into one native Go executable. The compiler/runtime implementation is unchanged.

The manifest [`helper-bundle.json`](../gossip/tests/gohp/helper-bundle.json)
pins both complete source-file hashes, baseline commit, exact excerpt ranges,
compiler/runtime and immutable Go/PHP images. Each excerpt retains its bytes;
the bundle adds only the PHP opening tag, separator newlines, and explicitly
owned composition drivers. Source ranges, excerpt hashes and their bundle line
offsets are recorded. No WordPress source file is edited.

## What is combined

The ten originals are the six `__return_*` helpers, `trailingslashit`,
`untrailingslashit`, `wp_validate_boolean`, and `zeroise`. Five owned PHP drivers
call these original functions; they are translated in the same compilation,
not implemented by Go observers:

- slash normalization followed by zero-padding;
- empty-array return passed to Boolean validation;
- integer-zero return passed to zero-padding;
- false return passed to Boolean validation;
- empty-string return passed to trailing-slash handling.

The 22 explicit cases contain ten direct probes (one per original) and twelve
composition probes. They are a new, smaller integration corpus, **not** the 73
typed observations from the existing four units rerun in combination. The old
manifests, cases, observers, pins and execution records remain untouched. A
combined PASS applies only to these 22 cases, not all earlier input domains.

## Observation and scope

The Go/PHP observers only dispatch fixed named probes and report actual return
type/value; they never implement slash/Boolean/padding algorithms or derive
results from fixture expectations. Both validate input argument counts/types
and integer padding widths 0–4096. Go uses checked JSON `UseNumber`/`Int64`
conversion; return normalization admits bool/int/string/null and the empty
GoHP array only. Native PHP selects and verifies `LC_CTYPE=C`.

One committed compiler archive supplies generation and the linked runtime.
Generation report source/output hashes are checked before build. The existing
non-root, no-network, read-only, resource-limited containers and the admitted
PHP 8.3.36 image/same-process version guard are reused. Every result must equal
both native PHP and the authored fixture, with empty stderr and exit 0.

```sh
cd gossip
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressHelperBundle$'
```

Without opt-in configuration this test skips, not passes. Executed results and
exact revisions belong in the [execution record](evidence/helper-bundle-2026-10-25.md).
There is no generated HTTP server, core startup, database, theme/plugin loading,
general formatting/coercion support, or link to the native hook SDK. This does
not import the separate GoHP #17 feature stack or repin the standalone units.

## Capability exclusion

`gossip.dev.wordpress-unit.translate-compare` remains **excluded** from WebMCP:
opt-in local developer tooling only. No UI operation, route, session forwarding,
automatic submit, production audit endpoint, publication or deployment is added.
No devboard card is linked or marked Done. Default-branch merge, release,
deployment and license finalization remain outside this experiment.
