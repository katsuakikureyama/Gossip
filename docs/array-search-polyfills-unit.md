# Original array search polyfills

This separate unit pins GoHP #54 at `36191c57003408d001d9f7217805a3549fdd9f98`.
It targets three sibling WordPress originals sharing the same callback/iteration
contract: `array_any`, `array_find`, `array_find_key`. It is not an expansion of
the nineteen-original/192-case integrated bundle and does not replace its pins.

The [execution record](evidence/array-search-polyfills-2026-10-26.md) reports
30/30 actual comparisons and vet PASS on the exact committed configuration.

## Original source and bounded contract

Whole-source SHA-256 and the WordPress baseline are pinned before extraction.
From `wp-includes/compat.php`, declarations 612–620, 566–574, 589–597 are
compiled byte-for-byte, preserving array/callable annotations and Boolean return.
Native PHP executes complete guarded blocks 600–621, 554–575, 577–598.
Reflection checks all three are user-defined, not PHP8.4 builtins.
Extracting declarations is not support for top-level conditional loading.

Nine owned drivers call these originals with capture-free arrow callbacks:
strict value/key equality, an early-success numeric poison probe, and scalar
truthiness. The stop callback uses `$value === 'stop' || ($value + 0 === 1)`;
later nonnumeric poison must never execute once a match returns.
This asserts successful short circuit, not cross-language error equivalence.

Thirty authored cases (ten per helper) compare actual return type/value and
ordered caller input after the call against PHP and explicit expectations.
Inputs are flat null/bool/int/string values, integer keys/values bounded to
±1,000,000. Normalized duplicate keys, floats, objects, references and nested
arrays are excluded. No expected fields construct generated-runtime outputs.

Transport helpers are extracted unchanged from the existing numeric-array
observers, excluding their test execution/generated-function calls.
Markers must occur once; source and selected transport hashes are recorded.
The new observers perform their own original-function dispatch and measurement.
The old observer files, corpora, manifests and evidence are unchanged.

A single translation/build/binary covers this three-original unit. PHP8.3.36
minimum, same-process admission, C locale, non-root/read-only/network-off and
resource restrictions remain unchanged.

## Reproduce and remaining boundaries

From `gossip/`, non-root with cached pinned Docker images:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArraySearchPolyfills$'
go vet ./tests/gohp
```

Absent opt-in means SKIP, not PASS. Record the exact committed snapshot and
actual results before claiming compatibility. No compiler changes or reruns of
prior PASS configurations are needed.

`wp_ext2type` is not included: its arrow captures `$ext`, and its dependency
`wp_get_ext_types` invokes filters. Object/method-based interactivity callers
and broader KSES parsing are likewise outside this unit. General closures,
arbitrary callables, full-file/core bootstrap, HTTP, DB, themes and plugins
remain unverified. This is not a production/release or merge decision.

Capability `gossip.dev.wordpress-unit.translate-compare` remains explicitly
**WebMCP-excluded**: local developer fixture translation and comparison, with no
browser/API operation, authentication-route change or auto-submit.
Devboard is unidentified/unsynchronized, not Done.
