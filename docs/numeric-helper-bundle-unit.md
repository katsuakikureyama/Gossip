# Nineteen-original numeric helper bundle

This separate configuration pins GoHP PR #54 at
`36191c57003408d001d9f7217805a3549fdd9f98`. It integrates the
[seventeen-original bundle](yn-helper-bundle-unit.md) with the
[two-original numeric unit](numeric-array-helper-unit.md).
Old manifests, corpora, rejection records and evidence remain unchanged.
Compiler CI and the previous independent results are not proof of this bundle.

The [execution record](evidence/numeric-helper-bundle-2026-10-25.md) reports
192/192 matches in this exact integrated configuration, with empty stderr,
exit0 and vet PASS. This does not promote or replace previous pins.

## Extraction and observation

The manifest pins the WordPress baseline, whole-source hashes and exact ranges.
All nineteen original declarations retain their bytes, annotations and bodies.
Native PHP executes the complete original guarded `array_all` block and verifies
user-defined polyfill provenance. The compiler extracts its unchanged declaration:
this is not conditional-declaration, full-file loading or bootstrap support.

One compiler invocation translates nineteen originals and thirty-three owned
drivers. One generated binary observes all 192 cases: the unchanged 149 legacy
cases, 24 numeric predicate cases, seven direct polyfill cases and twelve new
cross-helper cases. Retained corpus SHA-256 values are checked explicitly.

The new drivers combine associative slicing with numeric-key/list/Boolean
classification, slashing, value reindexing, string callbacks and Y/N predicates.
Their twelve authored expectations cover empty selection, sparse/negative keys,
mixed string keys, selector order/duplicates/misses, normalized numeric-looking
keys and y/Y/false values. They accept bounded flat arrays of strings.
Actual Boolean lists and actual ordered caller inputs/selectors after execution
are compared with native PHP and explicit fixture expectations.

Existing Go observers are adapted only from init-and-exit to returning their
observations, with a distinct legacy corpus path. Existing PHP observers use
require_once and that distinct path. Each textual adaptation must match exactly
once; original and adapted observer hashes are recorded. Existing observer files
are not edited. No expected field supplies generated-runtime observations.

## Reproduction

From `gossip/`, using a non-root user and cached immutable Docker images:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressNumericHelperBundle$'
go vet ./tests/gohp
```

The absent opt-in skips, not passes. Exact committed configuration must be
recorded with results before claiming differential_pass. Same-process PHP 8.3
admission, minimum 8.3.36, C locale, non-root, read-only, network-off and resource
limits remain unchanged. Generated source/report hashes and empty stderr/exit0
are required. No rerun of previous PASS configurations is needed.

## Capability and remaining boundaries

Capability `gossip.dev.wordpress-unit.translate-compare` remains explicitly
**WebMCP-excluded**: local developer translation/comparison with temporary
artifacts, no browser/API operation, new authentication route or auto-submit.
This test does not authorize external publication, merge or deployment.

This is bounded declaration-level integration, not general callback, full-core,
HTTP, database, theme, plugin, release or production compatibility.
WordPress originals, SDK/hooks, PHP admission, prototype and master are unchanged.
Devboard is unidentified/unsynchronized, not Done.
