# Original add_freeform /16-case execution record

PASS on2026-10-26JST: all16 Go/PHP8.3.36/C/expected comparisons, both exits0/
stderr empty. Three cases skip;13 append an original freeform Block array,
including an empty slice on nonzero length. There are16 add_freeform calls,
13 freeform calls and13 Block constructions in this one binary.
[Contract](../parser-add-freeform.md); [all16 observations/hashes](parser-add-freeform-2026-10-26.json).

## Consumer snapshot, first failure and recovery

Success clean commit43ebeba5710e4e4859109401b85ae48783839c06,
tree5f17710e413b4a8a654a13b124c634b5d31e5bd3, branch
codex/gohp-parser-add-freeform, worktree /private/tmp/gossip-parser-add-freeform,
base master4b8f55ae1d7fd0cc72be5a140390c301be49a0f1. From nested gossip/:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-integer-subtraction GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserAddFreeform$'
```

PASS10.48s/package10.690s. Raw successful log:
/private/tmp/gossip-parser-add-freeform-success-output.log. No successful rerun
or code/config change after this PASS; only documentation/evidence follows.

Initial clean538abb9b16dc1ebfdc93750ea9958c20ba3321b4,
treee393e19aba827e2c849e9d8e54830d37760450b7, used #80 compiler main4df47e0 and
the same command with GOSSIP_GOHP_SOURCE=/private/tmp/gohp-parser-output.
Generation FAIL10.22s/package10.469s: PHP002 line117 expected semicolon at minus.
Candidate build/run/PHP comparison not-run; not an observed PHP result mismatch.
Original source/corpus/observers/harness were not changed. #81 added the missing
integer subtraction; consumer pin/documentation only changed before resuming.
Initial failure log remains /private/tmp/gossip-parser-add-freeform-output.log.

`GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp` PASS on initial538abb9.
The Go harness code is unchanged in43ebeba; only JSON compiler pin/doc changed.
That unaffected static-check evidence is reused, not falsely described as a new
vet run on the later commit. Timings are not controlled performance comparisons.

## What actually ran

Go selects original Parserprefix16–50/freeform314–316/add_freeform327–335/
closing396 plus complete original Block15–90. Other five Parser methods are
omitted, not stubbed. PHP loads full original Parser/Block/Frame files; native
Reflection verifies seven-method original class/line/default shapes. Frame is
loaded only. One owned driver initializes actual Parser state through source,
then invokes original add_freeform; no observer field injection or host Block
construction. Source bodies/whole files/selection ranges are hash-bound.

Observed return is null, document/offset/empty stack stay unchanged, existing
output seed remains ordered, appended Block arrays have all five original keys
in declaration order and caller input/seed remain unchanged. Omitted/null length
uses strlen(document)-offset, zero skips exactly, other signed lengths use byte
substr. Explicit nonzero with an empty slice still appends innerContent=[''].
NUL and split UTF-8 bytes are compared through base64 observation, not normalized
text. Literal expected bytes are independent of the actual execution.

## Compiler prerequisites / exact fixed CI

[GoHP #80](https://github.com/katsuakikureyama/GoHP/pull/80) main
4df47e00cae2dfd25e4cc5bde299ef51cc56c4c5/tree83cc11e72f83675b5a6ab606bbd6782b8bfd2d57
adds method null defaults/bare returns/declared-array property append/sole-new
method factory snapshots/actual byte substr. Parallel #79 exact scalar required
constructors retained by normal merge; only catalog/compatibility/fixture-list
conflicts union-resolved, auto-merged core reviewed. Feature05314c59ccc6af5177e53fe1ddf114cf95432fe7
has the same tree. [Feature CI40274408258](https://github.com/katsuakikureyama/GoHP/actions/runs/40274408258),
PR40274414594 and [main CI40274575307](https://github.com/katsuakikureyama/GoHP/actions/runs/40274575307)
PASS564 checks/79 features, exact clean commits/tree. Seven new checks.

[GoHP #81](https://github.com/katsuakikureyama/GoHP/pull/81) final pin
565e65e6c9d8e211b06779739bfba76b07c4d92d/tree3b509a251cb2ab3d037c089c1ab0295a1e0a0d60
adds actual-int64 binary subtraction, left association/addition precedence and
overflow rejection without negating MinInt64. No weak coercion/general unary
negation/decrement. Featureca5c6150444838960b78138430db6c66ba16c379 has the same
tree. [Feature CI40274995512](https://github.com/katsuakikureyama/GoHP/actions/runs/40274995512),
PR40275002284 and [main CI40275173481](https://github.com/katsuakikureyama/GoHP/actions/runs/40275173481)
PASS569 checks/80 features, exact clean commits/tree. Five new checks.
Both downloaded feature/main artifacts use actualPHP8.3.36/Go1.23.12linuxamd64.

Local compiler env: GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0:

| Snapshot | Commands / actual outcomes |
| --- | --- |
| #80 initial12021af8b52f81f9086e5a3c3eb204cc315bcafa/tree589475fd14320f1284419d7d562742442d992f83 | go test -count=1 ./internal/frontend ./internal/ir ./internal/semantic ./internal/codegen ./runtime ./internal/compat PASS(0.257/0.210/0.570/0.397s;ir/semantic no tests); go vet ./internal/frontend ./internal/ir ./internal/semantic ./internal/codegen ./runtime ./tests PASS |
| #80 runtime/negative | GOHP_PHP_TARGET=8.3 go test -count=1 ./tests -run '^(TestParserOutputRuntimeGuards\|TestUnsupported\|TestMalformedIRFailsWithoutPanic)$' had one obsolete property-append negative FAIL; runtime/malformed/other negatives passed. Test-only9022960 changes it to nested property append; affected TestUnsupported/array_append_property PASS0.219s |
| #80 native | Fixed wrapper PATH /private/tmp/gohp-parser-output-php.3s9gR4; GOHP_PHP_TARGET=8.3 go test -count=1 -v ./tests -run '^TestDifferential/parser_output$' PASSfixture0.70s/package1.426s; reset=2:reset:rhs confirms same-slot RHS ordering, caller:method:after confirms Execution sharing |
| #81 initiald54a85dee1dc0eed668bf6703c7c76e23f947520/tree126f0467f4fd837c048797d4aff6f6921db65374 | go test -count=1 ./internal/frontend ./internal/codegen ./runtime ./internal/compat PASS3.087/2.824/3.023/3.288s; go vet ./internal/frontend ./internal/semantic ./internal/codegen ./runtime ./tests PASS |
| #81 runtime/negative | GOHP_PHP_TARGET=8.3 go test -count=1 ./tests -run '^(TestIntegerSubtractionRuntimeGuards\|TestUnsupported\|TestMalformedIRFailsWithoutPanic)$' had obsolete valid1-2 negative FAIL; others passed. Test-only081f3fa changes it to noninteger operand; affected TestUnsupported/negative_integer_subtraction PASS0.201s |
| #81 native | Fixed wrapper PATH /private/tmp/gohp-integer-subtraction-php.GXdFkE; GOHP_PHP_TARGET=8.3 go test -count=1 -v ./tests -run '^TestDifferential/integer_subtraction$' PASSfixture0.56s/package2.961s; precedence/coalesce/property results match |

No product-code corrections after local PASS; only obsolete negative tests/docs
changed. Successful local checks were not manually repeated. Compiler native
harness compares stdout/normal exits only, not stderr/exceptions/other effects;
consumer additionally requires both stderr empty. PR bodies preserve fixture/
generated hashes; consumer JSON preserves all actual hashes/records/CI identities.

## Limits

Historical freeform/Frame/data/helper configurations/pins/corpora/source/evidence
unchanged/not rerun or summed. Other five Parser methods/parse/proceed/tokenizer/
stack, registry/render/CSS/HTML/ThemeJSON/bootstrap/HTTP/DB/themes/plugins, generic
casts/errors/coercions/references remain unverified. Corrected source audit:
coalesce left operand is local length, not a property; no property-left probe
extension. No production target configured/created. WebMCP local fixture excluded,
route/auth unchanged. Devboard unidentified/not synced. Parent core/harness/
integration/review/measurement; workers independent runtime/fixtures/observers.
