# Sixteen-original HEIC/charset/array bundle

This independently pinned configuration translates sixteen byte-unchanged
WordPress helpers and twenty-two owned composition drivers together, then
builds one generated executable. It extends the15-original charset bundle
with `wp_is_heic_image_mime_type()` (functions.php9465–9474), not a replacement
implementation. The manifest binds three whole-file hashes, exact fragments,
compiler commit and immutable builder/oracle images. The HEIC, charset and
typed-array slice excerpts additionally have exact range/hash checks.

The compiler is unchanged GoHP `49a302e44464623dc4912bafbcf44192759d7878`
(Draft #44). Its prior fixed-PHP336-check PASS is reused, not rerun or described
as a new CI result. This bundle still needs its own original-source comparison;
the standalone and earlier bundle results do not establish integration.
Separate control-flow compiler #43 is not adopted.

## Corpus and transport

There are119 calls: all89 prior charset-bundle cases, all24 standalone HEIC
cases, and six new compositions. Direct79/composed40, with38 operations.
The prior case arguments, authored expectations and post-call expectations
are retained exactly. The new drivers are:

- `gossip_slice_heic`: original slice, then map the original MIME predicate.
- `gossip_slice_unescape_heic`: original slice, shallow string unescaping,
  then map the original MIME predicate.
- `gossip_heic_boolean`: original slash/unescape roundtrip, original MIME
  predicate, then original Boolean validation.

Two cases per new driver cover selected-key order, duplicate selectors,
missing/null exclusion, false/zero preservation, mixed/negative keys, empty
selection, escaped slash labels, nonmatching labels, and NUL roundtrip.
New HEIC direct/Boolean inputs and EVERY root value in new HEIC slice inputs
must be null/bool/int64/string, even unselected values. Array/object/float
needles are outside the compiler's strict membership contract.

Earlier operations retain the bounded ordered root `[key,value]` transport
and nested sequential lists. Keys are int/string, integer values/keys are
limited to ±1,000,000 and nested depth to8. No nested associative values,
objects, floats, references or cycles. New pair outputs preserve actual
selected keys/order. Both observers record actual type/value and caller args
after each call; expected values never construct observations.

## Execution and exclusions

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressHEICHelperBundle$'
go vet ./tests/gohp
```

No opt-in means skip, not PASS. Existing PHP8.3.36 admission and same-process
version guard, LC_CTYPE=C, immutable images, non-root/read-only/network-disabled
containers and resource limits remain. Generation/report hashes precede build;
PASS additionally requires matching PHP and authored expectations, empty
stderr and exit0. Execution evidence is recorded separately after the run.

This is MIME-label classification and charset-name recognition, not image
inspection/decoding/conversion, upload validation or UTF-8 text validation.
Earlier ten standalone units and10/12/13/14/15-original bundle configurations
remain unchanged. This does not include every historical standalone input.
Whole-file/core/HTTP/DB/theme/plugin behavior, option-dependent wrappers,
general callbacks, recursive unescaping and exception equivalence are unverified.
No master changes, merge, deployment, release or license decision.

Local developer capability `gossip.dev.wordpress-unit.translate-compare`
remains explicitly WebMCP-excluded; no browser/S2S/MCP/auth-route change.
Devboard remains unidentified/unsynchronized, not Done.
