# Original newline-preservation helper — 2026-10-25 JST

## Outcome and pinned targets

**PASS:** the unchanged `_autop_newline_preservation_helper()` was translated,
built and compared with native PHP for 16 nonempty lists of string captures.
All actual string return values/types matched PHP and authored expectations,
with empty stderr and exit 0. This is one callback helper, not execution of the
enclosing `wpautop` renderer or a regex/callback engine.

- Gossip test commit `ff02bbd05a0dd5b75940724cda2352622c543030`, tree
  `2665b662a262214191c5e34f9db7a43cdae43301`.
- GoHP compiler/runtime `dcb72c64d029c7a08a0e94256723ff4076209758`, tree
  `f1e7cfda4526801fda8e2c07a803e18b2b556045`.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`, unchanged
  `wp-includes/formatting.php` lines 811–813, whole-file SHA
  `d44aa1329df2fc38bf2670ad6de570e6b9eb65db9ed5507f26c683e569b2919e`.
- Actual PHP **8.3.36**, unchanged approved immutable image/same-process guard;
  host Go 1.22.1 darwin/arm64; builder Go 1.25.14 linux/arm64.

The [machine-readable record](newline-helper-2026-10-25.json) contains the
source/excerpt/case/observer/compiler/archive/generated/report/binary hashes,
native version and all 16 observations. Exact original bytes were checked;
the report hashes were verified before compilation. Only pinned committed
compiler files were exported to the hardened container workflow.

## Gossip execution

From the nested `gossip/` module:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressAutopNewlinePreservationHelper$'
go vet ./tests/gohp
```

The actual selected test ran once: **PASS 7.61s**, package **7.789s**. Vet PASS.
An earlier command used `^TestOriginalWordPressNewlineHelper$`, which selected
no test (`[no tests to run]`, package 0.215s); its zero exit is **not** translation
evidence. The selector/documentation were corrected without changing test code.
Vet ran after that initial command and is reused for the identical code. There
were no failed observations, implementation changes or reruns after the actual
PASS. Subsequent changes are documentation only.

Cases cover empty and ordinary text, single/adjacent/leading/trailing LF, CR-only,
CRLF and mixed endings, tabs, literal backslash-n, Japanese, emoji, NUL, existing
placeholder text and an ignored additional capture. A CRLF becomes CR followed
by the marker; this helper only replaces LF and does not normalize line endings
as a whole. The first capture is an actual string at numeric key zero. Observers
dispatch the original/generated function, not a copied replacement algorithm.

The Go observer explicitly checks string element types before constructing
`rt.Array`, avoiding null-to-empty-string JSON decoding. PHP asserts a nonempty
list and string entries. Empty lists, maps, missing keys, non-string captures,
warnings and weak coercion are outside the experiment.

## Compiler change and verification

[GoHP Draft #21](https://github.com/katsuakikureyama/GoHP/pull/21), stacked on
#18, adds exactly three-actual-string `str_replace`. Empty search leaves the
subject unchanged; other replacements are byte-wise, case-sensitive and
nonoverlapping, with literal replacement data and no rescan of inserted bytes.
Known bad types/arity and redeclaration fail compilation; dynamic bad types fail
at runtime, including when the search string is empty. Array/count/reference/
coercion/regex forms and PHP warning/TypeError equivalence remain unsupported.
Catalog: 33 entries, 83 unique checks, PHP 8.3 active target only.

At the exact compiler commit above, cached Go 1.23.0 darwin/arm64:

```sh
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./runtime ./internal/codegen ./internal/compat
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go vet ./...
```

PASS runtime0.196s / codegen0.382s / compat0.588s; vet PASS. The owned native
fixture covers ordinary/empty search/nonoverlap/non-rescan/empty replacement,
case sensitivity, mixed-case call spelling, multibyte text and literal `$1`.
NUL and invalid UTF-8 byte substitution are direct runtime-unit cases; NUL
preservation around LF is additionally observed by the Gossip original-source
fixture. Unsupported PHP byte-escape syntax is not introduced to the parser.

[PR CI 40119695048](https://github.com/katsuakikureyama/GoHP/actions/runs/40119695048)
PASSed all **83 catalog checks** on fixed Go 1.23.12 linux/amd64 / actual PHP
8.3.36. The clean merge checkout `63cc3a6f90df36f97f0be36322a7c2552212283b`
has tree `f1e7cfda4526801fda8e2c07a803e18b2b556045`, identical to the pinned
compiler head. Test exit 0, native differential, vet, summary and artifact
publication succeeded. `diff.str_replace` and the four new compiler/runtime
check groups PASSed. Native CI compares stdout/normal exit, not general
stderr/exception/side-effect equivalence; the separate Gossip observation
additionally requires empty stderr.

Artifact `compatibility-php8.3-40119695048-1`, ID12569728782, retained for 30
days. Initial CI PASS, no retries or reruns after PASS. These are compiler CI
results, not CI execution of the local Gossip test above.

## Preserved boundaries

Earlier four standalone unit manifests and the ten-original combined helper
unit remain unchanged and were not rerun. This fifth standalone unit uses a
new compiler pin; it is not an eleven-function bundle or a revalidation of all
previous values on this compiler. The native SDK/hooks, WordPress originals,
oracle policy, other GoHP feature stacks and normal GoHP checkout are unchanged.

Gossip work is stacked on Draft #10. No default-branch/upstream merge, release,
deployment or license finalization. Full `wpautop`, regex callbacks, general PHP,
HTTP/DB/theme/plugin integration and core startup remain unverified. The local
developer capability remains WebMCP-excluded under the
[unit contract](../newline-helper-unit.md); no authentication route is added.
No devboard card is linked, synced or marked Done.
