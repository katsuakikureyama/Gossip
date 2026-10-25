# Twelve-original helper bundle on the integrated compiler

This new independent configuration translates twelve unchanged WordPress
functions together, plus owned composition drivers, into one generated Go
executable. It adopts the combined conditional/string/array compiler, not a
collection of results from different compiler revisions.

The originals are six constant-return helpers, `wp_validate_boolean`,
`trailingslashit`, `untrailingslashit`, `zeroise`,
`_autop_newline_preservation_helper` and `stripslashes_from_strings_only`.
Their source remains at baseline
`54106f6c545fc8b881f6b39b7963002bece6c28c`; exact whole-file hashes, ranges,
excerpt hashes and positions in the generated source bundle are recorded.
No original body is rewritten or replaced by an observer implementation.

## New integration contract

The owned corpus contains twelve direct probes, one for each original, and
twelve composition probes. The compositions include extraction of capture
values followed by newline preservation and string-only unescaping, non-string
array passthrough followed by key/list operations, and earlier scalar helpers
chained with padding, Boolean validation or slash handling. These drivers are
test-owned PHP, not additional WordPress originals.

JSON inputs are limited to null, booleans, modest signed integers, strings and
numeric lists of those values. Go decodes numbers with `UseNumber` and `Int64`;
only runtime arrays with sequential int64 keys are normalized to JSON lists.
Argument guards retain the narrower requirements of individual operations,
including nonempty string captures and bounded padding widths. Typed root
results and nested list values must match native PHP and explicit expectations.
No expected results are calculated by the observers.

This does not add associative-map JSON transport, float/object/reference
support, recursive unescaping or arbitrary integer JSON precision guarantees.
The compiler's separate native fixture covers additional keyed-array cases;
that is not the same input contract as this Gossip observation.

## Execution and history

The manifest pins a committed combined compiler/runtime and the unchanged
approved Go/PHP images. Existing non-root, no-network, read-only and resource
limits, same-process PHP8.3.36 guard and report source/output hash validation
remain mandatory. Generation/build are intermediate statuses, not PASS.
Success requires all typed values and expectations to match with empty stderr
and exit0. Missing explicit local opt-in is SKIP, not PASS.

Earlier six standalone units and the ten-function bundle are retained with
their original compiler pins, cases and evidence. They are not repinned,
rerun or relabelled as passing on the new compiler. This new corpus provides
its own combined evidence; it is not a rerun of every earlier individual input.

```sh
cd gossip
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressIntegratedHelperBundle$'
```

See the [execution record](evidence/integrated-helper-bundle-2026-10-25.md)
for exact revisions, commands, all24 observations and compiler CI provenance.

Whole PHP files, WordPress startup, hooks, HTTP, databases, themes and plugins
remain unverified. No default-branch merge, deployment, release or license
decision is implied.

## Capability exclusion

`gossip.dev.wordpress-unit.translate-compare` remains **excluded** from WebMCP:
opt-in local developer tooling, without a UI or service operation. No browser,
S2S, MCP or callback authentication path changes, cookie forwarding, automatic
submit or production audit operation is introduced. No devboard card is linked,
synchronized or marked Done.
