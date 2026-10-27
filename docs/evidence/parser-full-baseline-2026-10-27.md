# Full parser baseline — 2026-10-27

Scope: [G83-PARSER](../php83-application-acceptance.md), complete original file
admission and a separate native `parse()` reference. Personal local paths below
are replaced by `GOHP_ROOT`, `GOSSIP_ROOT`, `PROBE_DIR`.

## Candidate admission: blocked, not PASS

- GoHP commit `9a7a9ce2ae0e514d80034dc09ae470a81fc3bb0d`, tree
  `51a138610a66066001d865080a372c2f1fd84ced`.
- Git tar archive SHA-256
  `9da00629c47aa958afe14f10fde87afd776530913a3648be496b4bc8c64e6453`.
- CLI binary SHA-256
  `f81cd335592bdb5a1b662cc48f7b1592a162a1b13e1ba0922e54f9a67ed799d7`.
- Original `wp-includes/class-wp-block-parser.php` SHA-256
  `8563621090de5e307431b1d29b78a9f75fd9ee238a0137b563a7db7b2e1f8d3b`.
- Go 1.23.0, Darwin ARM64, CGO disabled. Built from a clean GoHP code tree;
  documentation changes made afterwards do not alter the executable inputs.

```sh
cd GOHP_ROOT
CGO_ENABLED=0 GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOWORK=off \
GOTOOLCHAIN=go1.23.0 go build -trimpath -o PROBE_DIR/gohp ./cmd/gohp
PROBE_DIR/gohp -php-version 8.3 -experimental-pcre2 \
  -in GOSSIP_ROOT/wp-includes/class-wp-block-parser.php \
  -out PROBE_DIR/parser.go -report PROBE_DIR/parser.report.json
```

Build exit 0. Probe exit **1**, stdout 0 bytes, stderr 38 bytes:

```text
PHP002 line 98: unsupported token '>'
```

stderr SHA-256 `da11b51e216ddceec3268d792086b67269435e63b4f2214c2f6b19e5e8e6f870`;
empty stdout SHA-256 `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`.
Fresh output/report paths remained absent. Generation **blocked**;
generated binary build/execution/differential/application acceptance **not run**.
No method extraction, source concatenation, rewrite or alternate algorithm was
used. This rejection is the lexical first failure, not proof that the rest of
the source is accepted. Other blockers in the contract are inspection findings.

Setup attempts retained: `GOTOOLCHAIN=local` selected installed Go 1.22.1 and
could not build this module; explicitly choosing cached Go 1.23.0 with
`GOSUMDB=off` prevented toolchain verification. Removing that override succeeded
with `GOPROXY=off`. Neither failed setup ran the candidate. Docker image inspection
initially hit the filesystem sandbox; authorized read-only inspection succeeded,
confirming cached image `sha256:f39899ba9e42005403f535ec925a6e46a1c70707085b689e41a48219798521d7`, Linux ARM64.

## Native full-parser reference

Test inputs committed as Gossip `ed596cb94ae38631bd78744a92bb9adc260f4bc4`, tree
`676e09d956978bc5fe84dea3f721df50dbfd07db`. Only the new acceptance documentation
was uncommitted during execution; test/observer/corpus/manifest bytes were committed.

```sh
cd GOSSIP_ROOT/gossip
GOSSIP_PHP_RUNNER_IMAGE=php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd \
GOCACHE=/tmp/gossip-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 \
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressFullParserReference$'
GOCACHE=/tmp/gossip-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go vet ./tests/gohp
```

Each command ran once and exited **0**. Reference recording test PASS (1.05s;
package 1.289s), vet PASS. **12 native cases / 24 original parse calls recorded**:
fresh instance and instance reused across the corpus have equal typed results,
final state, PCRE/JSON error codes; caller document unchanged, stack empty after
each parse. The actual full parser and original require_once dependencies were
loaded, all seven original methods reflected. This is not hand-written expected
output, generated-Go comparison, exhaustive path coverage or application PASS.

Native PHP **8.3.36**, Linux/aarch64, integer size 8, PCRE **10.42**, C LC_CTYPE,
JIT0/backtrack100000/recursion10000/memory128M. Host Go1.23.0 Darwin ARM64.
Every final PCRE error code is 0; malformed JSON attributes record null attrs and
JSON error **4** on both fresh/reused calls. Other cases record JSON error0.
Float attributes, resource-limit failures and arbitrary documents remain outside
this first corpus; do not infer their outcome from these cases.

Native stdout 60065 bytes, SHA-256
`c02b289131717834e4b927721738804ac641b3cb24affc21815f657cad1faf57`, exit0/stderr0.
Full typed ordered results and final states are retained in
[machine evidence](parser-full-reference-2026-10-27.json), SHA-256
`1b34c7af6d7baeda2997e033c93727147bad4d36e6bfed676e183e6166ea43a7`.
The archive includes source/manifest/corpus/observer/harness/runner hashes, actual
platform/profile, each observation and the tested code commit. Status is
`native-reference-recorded`; compiler/differential/application acceptance are
`not-run` in this native record. The separate compiler probe above remains blocked.

No old next_token/regex/native/compiler suite was rerun. Original WordPress
tree differences outside README/docs/gossip and fork overlays remain empty;
existing fixture pins are unchanged. GitHub CI is not evidence for this opt-in
native fixture. No enabled Linux PCRE2 backend, full generated parser, DB/HTTP,
site workflow, PHP8.4 application result or deployment is claimed.
