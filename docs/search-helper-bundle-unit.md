# Twenty-two-original search helper integration

This new configuration integrates the nineteen-original numeric bundle and the
three-original array-search unit on GoHP #54
`36191c57003408d001d9f7217805a3549fdd9f98`.
Their old manifests, corpora, observer files and evidence remain unchanged.
Previous independent PASS counts are not proof of this configuration.

## What runs together

The manifest checks the WordPress baseline, three whole-source SHA-256 values
and exact line ranges. Twenty-two declarations compile byte-for-byte, retaining
annotations/bodies/indentation. Native PHP executes the complete original guarded
blocks for array_all, array_any, array_find and array_find_key and verifies
user-defined provenance. The compiler extracts declarations only: no conditional
loading, full-file or WordPress bootstrap claim.

One translation/build and one generated binary run all230 cases:
192 retained numeric-bundle cases,30 retained search cases and8 new compositions.
All five retained corpus hashes are fixed. Forty-four owned drivers call the
original functions without substituting their bodies.

Two new drivers combine wp_array_slice_assoc with bool_from_yn or
wp_validate_boolean, then array_any/array_find/array_find_key; the Y/N composition
also checks wp_is_numeric_array. Authored cases distinguish matched false from
no-match null, key0 from null, selection order, duplicate/missing selectors,
normalized numeric-string keys, mixed keys and null values excluded by slicing.
New inputs contain flat string/null values only; bounded integer keys ±1,000,000.
Retained cases keep their separate previously defined input domains.

The new Go observer measures actual scalar-valued list returns and actual ordered
caller input/selectors after the call. It never uses expected fields as output.
Native PHP and authored expectations must agree with the actual generated results.

## Observer reuse and isolation

Existing observers are reused by exact-once textual adapters: init/exit becomes
a returning function, standalone-only os imports are removed, corpus paths are
separated, and native helpers load once. The search observer reuses the transport
already loaded by the retained observer. Original and adapted hashes are recorded.
The previous observer files are not edited; one top-level observer controls the
single process. Expected data is independently constructed from the corpora.

Same-process PHP8.3 admission, minimum8.3.36, C locale, immutable images,
non-root/read-only/network-off and resource limits remain unchanged.

## Reproduce and boundaries

From `gossip/`, non-root with cached pinned Docker images:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressSearchHelperBundle$'
go vet ./tests/gohp
```

Absent opt-in means SKIP, not PASS. Record exact committed snapshot and actual
results before claiming differential_pass. Do not repeat old PASS configurations.

This establishes only bounded declaration-level behavior. Matched-null return,
general callable/capture/error semantics, arbitrary nested associative inputs,
full core/HTTP/database/themes/plugins remain unverified. wp_ext2type and filter
callers are excluded. No merge/deployment/release/licensing decision is made.

Capability `gossip.dev.wordpress-unit.translate-compare` stays explicitly
**WebMCP-excluded**: local developer fixture translation/comparison, no browser
operation/API/authentication-route change or auto-submit.
Devboard remains unidentified/unsynchronized, not Done.
