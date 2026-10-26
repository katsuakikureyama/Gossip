# PCRE2 comparison execution — 2026-10-27 JST

Result: **24/24 selected preg_match comparisons PASS, 0 mismatches**.
[Contract](../parser-pcre2-comparison.md),
[complete raw and PHP-shaped observations](parser-pcre2-comparison-2026-10-27.json).

## Snapshot and commands

- Clean code commit `c6d82db2faaad12592e3274d677c27c379dea9ac`;
  tree `c8bba02deadacc981a306ee6fbe33e43c26c342f`.
- Base master `e391bae8a37c1b81ee7586f4ea03ee81e6e56330` (#46),
  branch `test/parser-pcre2-comparison`. Later changes add only docs/evidence.
- Go1.22.1 darwin/arm64; GOTOOLCHAIN=local, GOPROXY=off, private local GOCACHE.
- Existing PCRE2 10.47_1 arm64 header/archive (no installation); runtime reports
  `10.47 2025-10-21`. Apple clang 17.0.0 (clang-1700.0.13.5).
- Retained PHP8.3.36 / PCRE10.42 reference SHA-256
  `5cbf560247afcd2c5f1462cab2f95f16bb822ac2d4d1756be16221cffdc6dc60`.
  No new native PHP execution or PHP8.4 result.

From `GOSSIP_ROOT/gossip` with GOSSIP_PCRE2_PREFIX selecting the pinned install:

```sh
go test -count=1 -v ./tests/gohp -run '^TestParserPCRE2Comparison$'
go vet ./tests/gohp
```

The initial sandboxed attempt stopped at `clang --version`: the macOS compiler
launcher emitted FSEvents/cache-access diagnostics; the harness rejected stderr.
Test 3.32s / package 3.541s, FAIL **before candidate build/comparison**. No semantic
mismatch was observed. Retried the unchanged commit/profile with sandbox escalation
for compiler access: PASS, test 0.73s / package 0.916s. All 24 probe processes
exit 0/stderr 0. No successful comparison rerun. Vet and formatting/diff checks
PASS; worker's C syntax-only check also PASS before the committed execution.

Candidate executable SHA-256:
`9ddc487a06799992daba2ccc7f9846ca4fbacdf904da2e3837bc2b386f00b66f`.
Source/harness/header/archive/clang-entrypoint hashes and compile arguments are in
the JSON. `compiler_sha256` identifies the resolved clang entrypoint (Apple's
launcher), not every Xcode toolchain component. Temporary binary is not published;
its hash identifies this execution, not a reproducible-build guarantee.

## Observed result and exclusions

17 matches, 6 no-matches and 1 match-limit failure match the retained PHP shapes.
Full numeric/named capture order, raw bytes (including FF/NUL), byte positions,
middle-unmatched entries and omitted trailing entries agree. The limit case
returns raw -47, independently mapped to PHP false/error 2/empty matches; normal
non-match is -1 mapped to integer 0/error 0. Both raw and mapped results are saved.

DOTALL=32, LF=2, BSR_UNICODE=1 are checked from compiled-pattern metadata. JIT is
not invoked; limits are match 100000/1, depth 10000, heap 131072 KiB. C locale,
full byte subjects, no UTF/UCP. Candidate heap/process isolation is not equivalent
to the earlier PHP container: this is a bounded trusted local experiment.

Only the direct regex result is compared. PHP error text, reference aliasing,
JSON decoding, next_token/parse execution, GoHP/cgo, regexp2, other profiles,
platforms, PHP versions and performance are unmeasured. WordPress source/history,
earlier evidence/corpus/compiler pins and module dependencies remain unchanged.
No committed Actions workflow supplies CI for this fixture; local execution is
the gate. Subsequent documentation reuses the exact tested source bytes.
