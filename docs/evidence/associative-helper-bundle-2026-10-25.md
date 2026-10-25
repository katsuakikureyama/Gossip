# Fourteen-original bundle execution — 2026-10-25

PASS: one translation/build and one generated executable matched PHP 8.3.36
and explicit expectations for all 55 calls (27 direct, 28 composed). Actual
return types/values, ordered root keys where applicable, and caller arguments
after invocation matched; stderr was empty and exit status was zero.

## Exact configuration

- Gossip tested snapshot: `9c14ce515c266d9db6d1f125a889c9e93cfc3dd1`, tree
  `c5ae3be90a0ed10feb4146f05bac6bb6f1af6879`. The exact staged bytes were run
  before commit, then committed unchanged with this tree.
- Branch `codex/gohp-associative-helper-bundle`, stacked on slice branch #17
  (`18f4e80af6e1b28fbef7d9f66012203bd46344b1`).
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`: fourteen
  byte-unchanged original functions from functions.php and formatting.php,
  plus sixteen owned composition drivers. Whole-file and excerpt hashes,
  selected ranges and bundle offsets are retained in the JSON record.
- GoHP/compiler/runtime `db3051f8a179078e2227c5d97f50a2f1611d1451`, tree
  `3155927445663c82c783c5d062b009b7d1a13648`, unchanged from standalone slice.
- Host Go 1.22.1 darwin/arm64; builder Go 1.25.14 linux/arm64;
  PHP 8.3.36, LC_CTYPE=C. Immutable image digests, source archive/compiler/
  manifest/corpus/driver/observers/generated/report/binary hashes and all
  actual observations are in the [machine record](associative-helper-bundle-2026-10-25.json).

## Commands and results

Run from nested `gossip/` with the existing approved non-root, read-only,
network-disabled, resource-limited Docker configuration and same-process
PHP version guard:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressAssociativeHelperBundle$'
go vet ./tests/gohp
```

First actual comparison PASS 8.41s (package 8.629s); vet PASS. This is local
isolated execution, not a Gossip CI run. No functional test failure or retry.

Before execution, `git diff --cached --check` reported a trailing blank line
at PHP observer EOF and stopped the chained commit command. The separate test
command still executed against the staged tree identified above. After PASS,
those exact bytes were committed, then the trailing empty line inside open PHP
code was removed. Publication adds only documentation/evidence and that
cosmetic deletion; no executable statement or effective configuration changed.
The JSON observer hash deliberately identifies the tested, pre-normalization
bytes. The unaffected PASS is reused, not claimed as a run of the later full
documentation tree; tests were not repeated. Final diff whitespace checking
is separate from runtime validation.

## Reused compiler evidence

[GoHP CI 40156201018](https://github.com/katsuakikureyama/GoHP/actions/runs/40156201018)
already passed all 238 checks across 46 features, plus vet/CLI smoke/artifact
steps, on the exact clean compiler commit/tree above. PHP 8.3.36 / Go 1.23.12
linux/amd64, push attempt 1. Artifact `compatibility-php8.3-40156201018-1`,
ID 12585095049, digest
`sha256:81ba128c07a3da3116d725b69cf3d17971934d0f8de25ee769b7f9031fc3e5e3`.
Metadata and earlier local compiler results are copied with `reused: true`
from the standalone slice evidence. Neither compiler tests nor CI were rerun
for this bundle. Compiler native schema observes stdout/normal exit, not
general stderr, exceptions or external side effects.

## Scope and remaining work

The corpus carries forward the previous 33 recursive-bundle cases and all 14
standalone slice cases, then adds eight compositions: slicing and escaping in
both orders, flat and deliberately shallow nested stripping, sparse shape,
and empty/nonempty array Boolean conversion. This new run validates those
inputs under the new configuration; it is not all historic standalone inputs.
Older eight standalone units and ten/twelve/thirteen-original bundles remain
unchanged and were not rerun or repinned.

See the [contract](../associative-helper-bundle-unit.md) for transport bounds.
Nested associative values, floats, objects, references, cycles, general
callbacks, map_deep/wp_unslash, whole-file/core startup, HTTP/DB and themes/
plugins remain unsupported or unverified. Original files, SDK/hooks and PHP
admission are unchanged. No merge, release, deployment or license decision.
Local capability `gossip.dev.wordpress-unit.translate-compare` remains
WebMCP-excluded; no auth-route change. Devboard is unidentified/unsynchronized.
