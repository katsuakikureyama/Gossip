# PHP 8.3 application acceptance

Owner direction, 2026-10-27: prioritize **one fixed Gossip application working on
PHP 8.3**, not PHP feature counts, broad version expansion or a storage rewrite.
This is the current delivery contract. Earlier bounded fixtures retain their
own pins and results; they are neither discarded nor added into an application
percentage. GoHP still implements general PHP semantics, not WordPress-name
special cases or handwritten replacements for original algorithms.

## Fixed first target

| Identity | Selection |
|---|---|
| WordPress source | `54106f6c545fc8b881f6b39b7963002bece6c28c`; version `7.2-beta1-64693` (development snapshot, not a production release) |
| Gossip baseline | `ec3dc36213ad501fd142de6ee207372ce881eaf3`; new fixtures are separately committed and hash-bound |
| PHP target / oracle | `8.3` / `8.3.36`, `php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd` |
| First unit | Complete `WP_Block_Parser` (all seven methods), `WP_Block_Parser_Block`, `WP_Block_Parser_Frame`, and the original two `require_once` edges |
| Candidate baseline | GoHP `9a7a9ce2ae0e514d80034dc09ae470a81fc3bb0d`; later candidates are new measurements, not silent repins |
| Native parser profile | 64-bit PHP, C LC_CTYPE, PCRE JIT off, backtrack limit 100000, recursion limit 10000, memory limit 128M; actual platform/engine recorded per run |
| Host/build | Record exact Go/OS/architecture and native dependencies per candidate. Existing optional PCRE2 evidence is Darwin ARM64 only; Linux enabled builds are not certified |

The [full-parser manifest](../gossip/tests/gohp/parser-full-reference.json)
pins complete source/license/version bytes and the new corpus/observer. A fixed
patch is a reproducibility choice, not a security guarantee or support for every
8.3 patch. A later oracle/source upgrade creates a new comparison series.
Existing GoHP 8.4 regression CI stays enabled; new minor-version work is deferred.
The separate `parser-source-closure.json` selects the same complete parser,
Block, and Frame files for the opt-in static-declaration closure; it does
not replace or repin the native-reference corpus.
Its fixed SHA256 is
`89f897c64146705559eae41bb50d49aff57abad0a11ee02d16d27fe78d64adaa`.

## Milestones and status

The eventual goal is the following **bounded single-site workflow**, not every
WordPress API, theme or third-party plugin. Later gates are requirements, not
already executable tests. Their DB/server/theme/platform/configuration and case
IDs must be pinned before execution; the whole application is not yet a fully
frozen runnable environment.

| Gate | Completion criterion | Current state |
|---|---|---|
| **G83-PARSER** (active) | Translate/build all three unchanged classes and their loading closure; run declared full `parse()` cases and compare typed ordered output, raw bytes, caller/state effects and diagnostics with pinned PHP | **Fixed-v1, carry-in and JSON key/type comparisons pass case-limited; acceptance remains open.** [Fresh fixed-v1 12/24](evidence/parser-full-fresh-2026-10-30.md) and [carry-in4/20](evidence/parser-carry-state-2026-10-30.md) unchanged. Separate [JSON key/type12/24](evidence/parser-json-key-types-2026-10-30.md) at clean2729405b8 matches actual PHP8.3.36 key order/types, signed64 extrema, bytes and all state/errors. [Boundary review](parser-full-acceptance-boundaries.md) still has unmeasured invalid-JSON/depth and original regex/resource paths. No whole-application acceptance |
| G83-BOOT | Original bootstrap and actual `plugin.php` / `WP_Hook`, including callbacks, globals and reentrant mutation; no mock hooks | Not run; source/dependency/config manifest still required |
| G83-DATA | A pinned relational DB adapter executes original data access and preserves schema, reads/writes, errors, transactions and post/options persistence across restart | Not run; DB product/version/config and adapter not selected; prototype JSON store is not a substitute |
| G83-HTTP | Original request flow: public pages, login/logout and protected admin, with correct status/body/headers/cookies, authorization and failure behavior | Not run; host/session/route contracts still required |
| G83-SITE | One pinned bundled theme, no third-party plugins: login → create/edit/save draft → publish → view public post → logout → restart/read persisted post; bounded local media upload/display | Not run; theme, fixtures, permissions and side-effect expectations still required |

Only after all required gates and scenarios pass at **one combined configuration**
may we say “the defined Gossip workflow works on PHP 8.3.” Passing isolated units
at different revisions does not satisfy that statement. Parser corpus v1 is the
first bounded input set, not every possible document, branch or resource limit.
Before admitting G83-PARSER, review the corpus for required error-limit/JSON/type
boundaries and explicitly retain any unsupported cases rather than removing them.

The [completed review and open probe plan](parser-full-acceptance-boundaries.md)
now record that all24 PCRE codes are0, JSON failure covers syntax4 only, raw
UTF-8/NUL is outside attributes, and JSON state is initialized before each call.
Original JSON capture is brace-bound; arbitrary scalar-root JSON is not a new
parser-callsite requirement. Do not equate observer failure or standalone builtin
success with original-parser behavior. The plan adds no executable cases and
does not modify the fixed12-case denominator, native records or evidence hashes.

Multisite, third-party plugins/themes, arbitrary Composer packages, external mail,
payments, automatic updates, public deployment and production readiness are not
part of this first workflow. They are not silently deleted from the original
source. No production destination is selected or changed here.

## Work backwards from the original parser

`parse` (line 64) resets state and repeatedly calls `proceed` (92).
`proceed` consumes `next_token` (239), then calls `freeform` (314),
`add_freeform` (327), `add_inner_block` (349) and `add_block_from_stack` (371).
Those paths construct/mutate real Block/Frame objects. Source loading is at
lines 403 and 410. These line anchors refer to the fixed source above.

| Slice | Actual original requirement | Evidence / exit condition |
|---|---|---|
| P1 — control/data syntax | `>` at 98, `<` at 128; `list` at 94/268; `switch` at 100 | Bounded relations (GoHP119), positional list assignment (GoHP120), string switches (GoHP121) added; [switch probe](evidence/parser-switch-probe-2026-10-29.md) passes switch and stops at207. Full source semantic/runtime acceptance is still pending |
| P2 — object locations and stack | Named class parameter at349; property `isset` at384; `array_push($this->stack, …)` at165, `array_pop` at205/372; chained property append and `.=` at207–208/351–361/380–381; object-to-array snapshots | Prior bounded admissions remain as documented. The new complete-parser run observes the stack/object behavior reached by its 12 cases in both fresh/reused passes and matches the retained typed record. It does **not** cover every P2 branch or make P2 complete; stack/snapshot boundaries outside the corpus remain open |
| P3 — tokenizer integration | Original `preg_match` at 250–256 reads document/offset properties, writes local nested capture array, then destructures/indexes it | The complete-parser run observes tokenizer behavior reached by its fixed 12 cases, including typed/raw-byte/state comparison against the retained record. This is not exhaustive `next_token` coverage: capture/error/resource and branches outside those cases remain unverified |
| P4 — attributes | `json_decode(..., true)` at 283 | GoHP138's bounded literal-true/default-depth/int64-only JSON admits original283. Fixed-v1 and separate key/type12/24 compare actual decoded attrs, ordered normalized keys, signed64 limits and raw bytes; carry-in remains its own4/20 record. Floats/flags/object mode and arbitrary inputs remain excluded or unverified; invalid-encoding/control/surrogate/depth error proof is not established by this key/type series |
| P5 — source closure and full unit | Original `require_once __DIR__ ...` at403/410 plus both complete dependency files | GoHP130 closure plus GoHP138 generation/build reaches the complete generated entry. Retained and fresh comparisons pass for 12 cases/24 calls per side. General include/require/autoload, dynamic/conditional loads, cycles, dependency initialization and all unexercised branches remain excluded or unverified |

The former compiler blocker was original `json_decode($matches['attrs'][0], true)`
at line283. The latest complete closure now generates, builds, and has one
enabled-native-PCRE retained-reference execution plus a separate fresh-PHP
comparison. These supply bounded runtime observations for P2/P3/P4/P5 only
where the fixed 12 cases reach them. They do not establish every branch or a
completed product goal. These rows remain dependencies of one application gate.
PR133's `empty` admission is deliberately narrow: defined bare locals whose
actual values are strings or arrays only, with strings empty exactly for `''`
or ASCII `'0'`. It does not make generic calls valid, suppress undefined/address
access, or broaden to other scalar/object domains; the legacy array-only IR
remains unchanged. The closure mode remains an opt-in compiler mode, not a
consumer pin, and excludes executable dependency initialization,
dynamic/conditional loads, cycles and autoload. Source flattening alone is not
evidence.
This parser does not use inheritance, closure captures or PHP reference-assignment
syntax. Those features can wait for a later demonstrated bootstrap/hook blocker.
Property mutation and builtin output arguments still need their real semantics;
their absence must not be hidden by eager copies or source rewrites.

## Reference and comparison boundary

The new `TestOriginalWordPressFullParserReference` loads the **complete original
files** in the admitted PHP container. It records typed ordered trees with base64
strings, parser state and reuse/reset behavior for a declared corpus of empty,
freeform, void/paired/nested/adjacent/malformed/unclosed documents, raw bytes and
JSON attributes. It does not implement a parser or supply synthetic expected Go
results. This is native reference capture, **not differential PASS**.
The [first execution](evidence/parser-full-baseline-2026-10-27.md) recorded all
12 cases with fresh/reused instances, exit0/stderr0 and preserved caller/state
contracts. Float attributes and resource-limit failure cases are not measured
in this initial corpus.

```sh
cd gossip
GOSSIP_PHP_RUNNER_IMAGE=php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd \
GOPROXY=off GOTOOLCHAIN=go1.23.0 \
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressFullParserReference$'
```

Use the cached admitted image and existing [oracle isolation](php-oracle-policy.md).
Missing runner means SKIP, not success. Never merge this new corpus with the
old 24-case **next_token** corpus or the 23-case owned regex driver denominator.

For candidate admission, the earlier single-file, closure, and unknown-receiver
probes stay historical. The [string-empty probe](evidence/parser-string-empty-probe-2026-10-29.md)
uses the same explicit manifest/root/entry selection over the **unchanged
complete three-file closure**, not an excerpt or concatenation. At clean code
`62cc41fdc45bf2126acd7c54d70c28034d4557ec` / tree
`17e6ab69577bd640904405d85759596b9865961a`, it records exit1, stdout0,
stderr114 bytes (SHA256 `0aa3c061c7896e4450a840ff8a41e16ac9c042c6b8e496d72eae9ac57bb41dba`),
the line394 semantic rejection, and absent artifacts. It cannot be promoted to
parse execution even if its first rejection later disappears.

The [guarded-object-array probe](evidence/parser-guarded-object-array-probe-2026-10-29.md)
uses that same unchanged closure at clean code
`dbbcdc225247b0b826267ce0546814d9d7cf6e0a` / tree
`cc78cd70c8de172658a6be97937561c207127b3a`. It records exit1, stdout0,
stderr102 bytes (SHA256
`9db57cfa3d6c3e0411ef3e4725b12b1d18e38b5d6258c605bc15abb53d689c3e`),
the codegen line165 `array_push` rejection, and absent artifacts. The tested
candidate head `7457217dcc4594e26ad11449ef846417be38eba2`, tree
`1adeba9e3f998ba5a77a29e2f4ca71dcd1abd908`, is merged at main
`10086b9528e78ce19cf4cf3c1c128248b981c80f` with that same tree. Push/PR CI
passes and exact-main CI40888584852's two successful target jobs are compiler/
owned-fixture evidence, not this original-source probe. It remains neither original execution
nor a native PHP comparison.

The [array-push probe](evidence/parser-array-push-probe-2026-10-29.md) uses the
same unchanged closure at clean code `2d01aa04b09d926d48299baa4a2210065af435e8`
/ tree `eff55f7dcdaba035382c04adb7a909f7335e2280`. It records exit1, stdout0,
stderr101 bytes (SHA256
`246e5052200024fe11f7578fb32e1d9d2bd9215f7449d6c06bb020a0c0be9b00`), the
former line165 `array_push` rejection passed, and the next codegen rejection at
line205 `array_pop`; Go/report/staging artifacts are absent. The resulting
main tree is `8a7fcb8aed19f680c8a50f164c48eb39bc0642ab`; push/PR/exact-main CI covers its
owned fixture, not original parser execution or a native PHP comparison.

The [array-pop probe](evidence/parser-array-pop-probe-2026-10-30.md) uses the
same unchanged closure at clean code
`1adc11ad053a19e22b98fcecefd898851830dd12` / tree
`00adfc030a89eea4fcf57f3fbd2eb2933925c4e8`. It records exit1, stdout0,
stderr103 bytes (SHA256
`97d870019d0d48ac7e09006c80a7b07877b61bad3e74c4b9491b984b95ee8f4c`), passes
the former line205 `array_pop` rejection, and reaches the codegen line283
`json_decode` rejection; Go/report/staging artifacts are absent. [PR137](https://github.com/katsuakikureyama/GoHP/pull/137)'s push/PR CI and owned native
fixture comparison pass; it is merged at main `af3e9ada36728cb317096b35af6ca0f22e3492c4`, and exact-main CI also passes. This original-source probe remains
neither original execution nor a native PHP comparison.

Record stages separately: source verified → generation → native binary build →
execution → differential comparison → application acceptance. A blocked compiler
has no generated source/report/binary to certify. Native capture, expected
rejection, infrastructure failure, SKIP and unmeasured remain distinct states.
No feature-count or compatibility percentage is introduced here.

Latest [JSON-generation evidence](evidence/parser-json-generation-2026-10-30.md)
records clean35524a6's one complete unchanged generation plus one CGO0 build,
both exits0. Go/report exist and retain all three source hashes; no original
method was called. This supersedes the prior compiler stop, not the native
reference corpus or its acceptance criteria. Original execution, engine/driver
configuration and typed differential comparison remain unverified.

## Capability / progress management

The fixture is a developer-local read/inspect operation: stable ID
`gossip.dev.full-parser-reference`, input the fixed manifest plus admitted cached
runner, resource the pinned original parser corpus, stage `inspect`, status
experimental. Result states are native-reference-recorded, skipped or failed;
compiler/differential/application acceptance remain not-run/blocked as applicable.
Separate GoHP-local generation/build evidence uses its existing excluded local
compiler capability and is now observed; it does not change this native-reference
operation into a hosted execution or an application-acceptance operation.
Its structured log is local evidence, not a persistent audit service. UI, AI
Assist and WebMCP are **excluded**: there is no hosted operation, browser
callsite, server authorization contract or public execution API. It performs no
save/publish/send/pay operation on user application data. No browser/S2S/MCP/
webhook/public route is added, and no cookie/session/token is transferred.

Use this gate table as the current status and link evidence/PRs from Issue #1
in both repositories and [current state](current-state-fork-baseline.md).
The external devboard card has not been identified and is **not synchronized**.
Storage, performance and target-expansion history remains valid evidence, but
does not override this application-first delivery order.
