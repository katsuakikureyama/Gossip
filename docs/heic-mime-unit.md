# HEIC/HEIF MIME-label predicate

The tenth standalone original-source unit extracts unchanged WordPress
`wp_is_heic_image_mime_type()`, functions.php lines9465–9474. The manifest
pins WordPress, full-file SHA256, exact fragment, compiler commit and immutable
builder/oracle images. The test additionally requires byte-exact original
source; the implementation is not replaced with a Go equivalent.

GoHP adds only `in_array(needle, haystack, true)` for null/bool/int64/string
needles and scalar-valued arrays. It uses strict type-and-value comparison;
the complete haystack is guarded before lookup. No loose or dynamic strictness,
array/object/float needle, nested-value membership or coercion support.

The owned corpus contains24 inputs: four accepted exact labels, fourteen
rejected strings (case, whitespace, parameters, suffix, NUL, Unicode and URL
included), plus null/true/false/0/1/-1. JSON transport is scalar-only with
integers bounded to ±1,000,000; arrays, objects and floats are rejected.
Both observers call the actual original/generated helper and record actual
Boolean returns and the caller input after each call. Results must equal
native PHP and independent authored expectations, with empty stderr/exit0.

The harness retains same-process PHP8.3.36 admission, hash-bound generation
report and isolated non-root/read-only/network-disabled resource-limited
containers. Status progresses generated → compiled → differential_pass; the
first two are not compatibility evidence. Run from nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressHEICMimeHelper$'
go vet ./tests/gohp
```

This classifies a MIME string only. It does not inspect image content, decode,
convert or validate uploads/media metadata. This new pin builds on compiler
#41 but does not update or rerun the earlier nine standalone units or the
15-original/89-case bundle. It is not a16-original integration result.
Whole functions.php, WordPress core/HTTP/DB/theme/plugin execution, general
PHP semantics, deployment and license resolution remain unverified.

Local developer capability `gossip.dev.wordpress-unit.translate-compare` remains explicitly
WebMCP-excluded: no browser action, public agent operation or auth-route change.
No master change, merge, release or deployment. Devboard remains unidentified
and unsynchronized, not Done. Execution evidence is recorded separately after
the pinned snapshot is run.
