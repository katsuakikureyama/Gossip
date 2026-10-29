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
| **G83-PARSER** (active) | Translate/build all three unchanged classes and their loading closure; run the declared full `parse()` cases and compare typed ordered output, raw bytes, caller/state effects and diagnostics with pinned PHP | **Blocked**: merged GoHP133 passes original string-empty354, then rejects at394: `PHP004: array cast requires a known generated class` for `(array) $stack_top->block`. Its owned string-empty fixture passes native PHP8.3/8.4 comparison, but the three-file probe is admission only: no original generation/execution or acceptance. The12 native reference cases /24 original calls remain reference-only |
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
| P2 — object locations and stack | Named class parameter at349; property `isset` at384; `array_push($this->stack, …)` at165, `array_pop` at205/372; chained property append and `.=` at207–208/351–361/380–381; object-to-array snapshots | GoHP122 adds nested updates,124 declared-class guards,126 quiet property presence. PR131's bounded direct named-property admission passes original line209; PR133's [string-empty probe](evidence/parser-string-empty-probe-2026-10-29.md) then reaches the unknown-class snapshot cast at394. Full-unit stack mutation and snapshots remain open; P2 is not complete |
| P3 — tokenizer integration | Original `preg_match` at 250–256 reads document/offset properties, writes local nested capture array, then destructures/indexes it | GoHP123 adds named property inputs and the [full-source probe](evidence/parser-preg-property-probe-2026-10-29.md) now reaches349. Enabled owned fixture comparison records15 typed observations at its own pin; existing driver23/23 remains at its old pin. Neither is **original next_token translation**. Capture/raw-byte/error/full-method integration remains pending |
| P4 — attributes | `json_decode(..., true)` at 283 | Preserve associative keys/types/nesting and malformed-input null/error semantics; do not substitute a JSON string or hard-coded attrs |
| P5 — source closure and full unit | Original `require_once __DIR__ ...` at403/410 plus both complete dependency files | GoHP130 connects L0's checked graph and file-aware locations into explicit manifest/root/entry assembly and v0.2 reports. The complete unchanged parser, Block, and Frame pass both original loading directives; its [L2 probe records the prior semantic rejection](evidence/parser-static-closure-probe-2026-10-29.md) at209. L1's owned real-file fixture passes native PHP8.3/8.4 comparisons. General include/require/autoload, dynamic/conditional loads, cycles and dependency initialization remain excluded; this is not original-parser generation or acceptance |

The measured next semantic blocker is `(array) $stack_top->block` at line394,
which requires an unknown-class snapshot; P2's remaining stack/snapshot
dependencies also need implementation before full-unit acceptance. These rows
define dependencies of the same gate, not five new standalone product goals.
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
probes stay historical. The new [string-empty probe](evidence/parser-string-empty-probe-2026-10-29.md)
uses the same explicit manifest/root/entry selection over the **unchanged
complete three-file closure**, not an excerpt or concatenation. At clean code
`62cc41fdc45bf2126acd7c54d70c28034d4557ec` / tree
`17e6ab69577bd640904405d85759596b9865961a`, it records exit1, stdout0,
stderr114 bytes (SHA256 `0aa3c061c7896e4450a840ff8a41e16ac9c042c6b8e496d72eae9ac57bb41dba`),
the line394 semantic rejection, and absent artifacts. It cannot be promoted to
parse execution even if its first rejection later disappears.

Record stages separately: source verified → generation → native binary build →
execution → differential comparison → application acceptance. A blocked compiler
has no generated source/report/binary to certify. Native capture, expected
rejection, infrastructure failure, SKIP and unmeasured remain distinct states.
No feature-count or compatibility percentage is introduced here.

## Capability / progress management

The fixture is a developer-local read/inspect operation: stable ID
`gossip.dev.full-parser-reference`, input the fixed manifest plus admitted cached
runner, resource the pinned original parser corpus, stage `inspect`, status
experimental. Result states are native-reference-recorded, skipped or failed;
compiler/differential/application acceptance remain not-run/blocked as applicable.
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
