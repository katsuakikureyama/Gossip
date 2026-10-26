# Gossip fork documentation

The inherited WordPress source remains authoritative for WordPress behavior.
These documents describe the experimental Gossip additions, not a production
support or compatibility promise.

| Document | Purpose |
| --- | --- |
| [Current responsibility map](reuse-boundary.md) | Original translation, native Go integration candidates and unsupported surfaces; current owner operating policy |
| [Parser backend assessment](parser-backend-assessment.md) | Fixed original next_token requirements, PCRE2/regexp2 comparison gates; no backend adopted or measured |
| [Historical README](history/readme-through-pr44.md) | Prior long narrative through PR44, preserved with rebased links; not current commands or policy |
| [Historical boundary overview](history/reuse-boundary-through-pr44.md) | Previous boundary/pin/approval wording retained as history, superseded by current map |
| [Original add_freeform](parser-add-freeform.md) | Two selected original methods/complete Block versus full original PHP Parser; raw-byte output append and skip boundaries |
| [Add-freeform evidence](evidence/parser-add-freeform-2026-10-26.md) |16 cases,13 appends/3 skips PASS; initial minus generation failure retained, GoHP #80/#81 main CI564/569 checks |
| [Original parser freeform](parser-freeform.md) | Selected unchanged parser method versus complete original PHP parser; explicit extraction/native boundaries and remaining tail dependencies |
| [Freeform evidence](evidence/parser-freeform-2026-10-26.md) |12 pair cases/24 actual method calls and vet PASS, typed data/identity/caller state; prior GoHP #76 evidence reused |
| [Original Frame + Block](parser-frame.md) | Two complete unchanged classes/constructors, Frame3/4/5 arguments and nested object identity after snapshot |
| [Frame evidence](evidence/parser-frame-2026-10-26.md) |12 actual typed/default/offset/identity/caller-state comparisons and vet PASS; GoHP #76 pin/main543-check CI PASS |
| [Original parser-block data class](parser-block-data.md) | Complete original class/constructor → array snapshot → real class decision, no substitute class |
| [Parser-block data evidence](evidence/parser-block-data-2026-10-26.md) |16 actual class/property/snapshot/decision/argument-state comparisons and vet PASS; GoHP #74 pin |
| [Source-built block input](elements-source-input.md) | Object creation/property assignment inside translated PHP, no observer field injection, same24-case corpus |
| [Source-built input evidence](evidence/elements-source-input-2026-10-26.md) |24 actual typed/factory-input/supports/caller-state comparisons and vet PASS; GoHP #71 pin |
| [Serialization utility → class decision](elements-serialization-path.md) | Three unchanged originals, explicit input-object boundary,16 direct +8 owned composition cases |
| [Serialization decision evidence](evidence/elements-serialization-path-2026-10-26.md) |24 typed/caller-state comparisons and vet PASS; initial extraction failure retained; old334 configuration unchanged |
| [Current state](current-state-fork-baseline.md) | Current documentation/ownership map plus original add_freeform16-case evidence; prior configurations remain separately pinned |
| [Real element-class decision](elements-class-path.md) | wp_should_add_elements_class_name + _wp_array_get without resolver stubs;40 caller cases +294 retained |
| [Element-class decision evidence](evidence/elements-class-path-2026-10-26.md) | All334 actual typed/caller-state comparisons and vet PASS; GoHP #67 exact pin and limits |
| [Twenty-five-original integration](settings-helper-bundle-unit.md) | 262 retained edge cases +24 lookup cases +8 new settings compositions, new GoHP #64 configuration |
| [Twenty-five-original evidence](evidence/settings-helper-bundle-2026-10-26.md) | All294 actual typed/caller-state observations and vet PASS; exact commands, snapshot and hashes |
| [Nested settings lookup unit](array-get-unit.md) | Unchanged original, retained24 cases, GoHP #64 pin; bounded recursive node transport |
| [Nested settings lookup evidence](evidence/array-get-unit-2026-10-26.md) | Actual typed return/caller states PASS; initial runtime failure and compiler correction retained |
| [Nested settings lookup boundary](array-get-boundary.md) | Original helper used by Theme JSON/global styles; default/isset/continue prerequisites |
| [Nested settings reference evidence](evidence/array-get-boundary-2026-10-26.md) | Native24 PASS, five parser rejections, exact snapshots/commands/hashes; Go/PHP not-run |
| [Twenty-four-original integration](edge-helper-bundle-unit.md) | All24 originals/46 drivers;230 retained +24 edges +8 new composed cases |
| [Twenty-four-original evidence](evidence/edge-helper-bundle-2026-10-26.md) | Actual typed/caller-state262-case PASS, exact snapshots, commands and hashes |
| [Original array edges](array-edge-unit.md) | Both unchanged declarations in one Go binary;24 retained cases on integrated GoHP main |
| [Array edge comparison evidence](evidence/array-edge-unit-2026-10-26.md) | Actual PHP8.3.36 return/input-after matches; commands, hashes and observations |
| [Original array_first](array-first-unit.md) | Separate12-case comparison on GoHP #57; original declaration and guard retained |
| [array_first evidence](evidence/array-first-2026-10-26.md) | Actual return/input-after matches PHP8.3.36; exact commit, hashes and local commands |
| [Array edge boundary](array-edge-boundary.md) | Historical PHP reference and empty/array_key_last rejection on older GoHP #54 |
| [Array edge boundary evidence](evidence/array-edge-boundary-2026-10-26.md) | Native24/24 PASS,2 original +1 dependency rejection; not Go/PHP differential PASS |
| [Twenty-two-original search bundle](search-helper-bundle-unit.md) | One binary,22 unchanged originals/44 drivers/230 cases; retained settings unchanged |
| [Twenty-two-original bundle evidence](evidence/search-helper-bundle-2026-10-26.md) | 230 typed/input-after comparisons and vet PASS on exact committed configuration |
| [Array search polyfills](array-search-polyfills-unit.md) | Separate three-original search unit;30 bounded cases PASS, not an expanded integrated bundle |
| [Array search evidence](evidence/array-search-polyfills-2026-10-26.md) | 30 actual type/value/input-after comparisons and vet PASS, exact snapshot and limits |
| [Nineteen-original numeric bundle](numeric-helper-bundle-unit.md) | GoHP #54,19 originals/33 drivers, single-binary192-case integration |
| [Nineteen-original bundle evidence](evidence/numeric-helper-bundle-2026-10-25.md) | 192 actual comparisons and vet PASS; Draft #27, not full WordPress bootstrap |
| [Numeric-array helper unit](numeric-array-helper-unit.md) | Two originals on GoHP #53 with31 cases; previous rejection remains historical |
| [Numeric-array helper evidence](evidence/numeric-array-helper-2026-10-25.md) | 31 actual comparisons PASS, separate from integrated bundle |
| [Historical numeric-array boundary](numeric-array-boundary.md) | Original predicate/polyfill dependencies and24-case PHP-only reference; old GoHP #44 translation rejected |
| [Numeric-array boundary evidence](evidence/numeric-array-boundary-2026-10-25.md) | Native reference PASS and two actual compiler rejections; explicitly not differential PASS |
| [Seventeen-original yes/no bundle](yn-helper-bundle-unit.md) | One binary,17 unchanged originals/25 owned drivers/149 cases; separate input domains and Boolean contracts |
| [Seventeen-original bundle evidence](evidence/yn-helper-bundle-2026-10-25.md) | 149-call typed/input-after PASS; exact pre-execution commit and reused compiler336-check evidence |
| [Yes/no-label unit](bool-yn-unit.md) | Eleventh standalone unit: unchanged bool_from_yn, string-only y/Y recognition; no general Boolean coercion |
| [Yes/no-label evidence](evidence/bool-yn-2026-10-25.md) | 24-case bool/input-after PASS; exact pre-execution snapshot and reused compiler evidence |
| [Sixteen-original HEIC/charset/array bundle](heic-helper-bundle-unit.md) | Separate119-case integration with16 unchanged originals and22 owned drivers; strict scalar HEIC domain |
| [Sixteen-original bundle evidence](evidence/heic-helper-bundle-2026-10-25.md) | 119 typed/input-after comparisons PASS; exact tested bytes and reused compiler336-check evidence |
| [HEIC MIME-label unit](heic-mime-unit.md) | Tenth standalone original helper; bounded strict scalar membership, not image inspection |
| [HEIC MIME-label evidence](evidence/heic-mime-2026-10-25.md) | 24-case bool/input-after PASS, exact pins and compiler validation |
| [Fifteen-original charset/array bundle](charset-helper-bundle-unit.md) | One binary with15 originals/19 drivers, typed arrays plus charset comparison; separate merged pin |
| [Fifteen-original bundle evidence](evidence/charset-helper-bundle-2026-10-25.md) | 89-call typed/input-after PASS and exact merged-compiler CI313-check result |
| [UTF-8 charset-name unit](utf8-charset-unit.md) | Ninth standalone unit: unchanged _is_utf8_charset with guarded ASCII comparison; not UTF-8 text validation |
| [UTF-8 charset evidence](evidence/utf8-charset-2026-10-25.md) | 28-case bool/input-after PASS; exact pins and new compiler CI296-check result |
| [Fourteen-original associative bundle](associative-helper-bundle-unit.md) | One generated binary with fourteen unchanged originals and sixteen owned drivers; bounded ordered-key transport |
| [Fourteen-original bundle evidence](evidence/associative-helper-bundle-2026-10-25.md) | 55-case direct/composed PASS, exact tested tree, actual caller arguments and reused compiler CI238-check evidence |
| [Associative slice unit](array-slice-assoc-unit.md) | Eighth standalone unit: unchanged wp_array_slice_assoc including typed return; ordered key/value transport |
| [Associative slice evidence](evidence/array-slice-assoc-2026-10-25.md) | 14-case original-source PASS, actual return and both caller arguments; compiler CI238-check evidence |
| [Thirteen-original recursive bundle](recursive-helper-bundle-unit.md) | One generated binary,33 direct/composed observations with caller-argument preservation; no recursive unescaping claim |
| [Thirteen-original bundle evidence](evidence/recursive-helper-bundle-2026-10-25.md) | Initial33-case PASS, exact source/test/compiler pins and reused compiler CI181-check evidence |
| [Recursive wp_slash unit](wp-slash-unit.md) | Seventh standalone original-source unit; literal self-callback and nested-list input preservation |
| [Recursive wp_slash evidence](evidence/wp-slash-2026-10-25.md) | 22 actual return/input-after comparisons; exact pins and compiler CI181-check result |
| [Integrated twelve-helper bundle](integrated-helper-bundle-unit.md) | One generated executable using the combined conditional/string/array compiler; separate direct/composed corpus |
| [Integrated twelve-helper evidence](evidence/integrated-helper-bundle-2026-10-25.md) | New24 typed observations and combined compiler CI166-check result; earlier configurations retained |
| [String-only stripslashes unit](stripslashes-helper-unit.md) | Original conditional helper, lazy string-only builtin and non-string passthrough; not recursive unescaping |
| [String-only stripslashes evidence](evidence/stripslashes-helper-2026-10-25.md) | Separate 21-case native PHP comparison and exact compiler/test pins |
| [Newline preservation unit](newline-helper-unit.md) | Original `wpautop` callback helper with bounded string replacement; not paragraph-renderer support |
| [Newline preservation evidence](evidence/newline-helper-2026-10-25.md) | Separate 16-case observation and pinned compiler evidence |
| [Combined helper unit](helper-bundle-unit.md) | Ten unchanged originals plus five owned composition drivers in one generated binary; 22-case integration contract |
| [Combined helper evidence](evidence/helper-bundle-2026-10-25.md) | Separate integration result; does not replace the four standalone unit records |
| [Zeroise-helper translation unit](zeroise-helper-unit.md) | Separate unchanged `zeroise()` contract with bounded string/integer zero-padding |
| [Zeroise-helper translation evidence](evidence/zeroise-helper-2026-10-25.md) | New unit execution record; earlier three-unit compiler pins remain unchanged |
| [Provenance](provenance.md) | Exact upstream/prototype references and unchanged source policy |
| [PHP oracle policy](php-oracle-policy.md) | Admission floor for new Docker-oracle runs; not a compatibility conclusion |
| [PHP 8.3 floor evidence](evidence/php-83-floor-2026-10-25.md) | Previous PHP 8.3.36 execution record, retained with its earlier three compiler pins |
| [Unified-compiler evidence](evidence/unified-compiler-2026-10-25.md) | Historical PHP 8.3.36 execution record for three independent GoHP units under their recorded compiler pin |
| [Hook comparison evidence](evidence/fork-hooks-2026-10-25.md) | Historical PHP 8.3.13 execution record: commit, runner, source/fixture hashes and exclusions |
| [Original-source translation units](translation-units.md) | Extraction, isolation, GoHP report verification and comparison contract |
| [Return-helper translation evidence](evidence/return-helpers-2026-10-25.md) | Historical PHP 8.3.13 record: five original helpers generated/built/compared; exact hashes and limitations |
| [Typed-return and rejection evidence](evidence/typed-helpers-2026-10-25.md) | Historical PHP 8.3.13 v2 result: five actual return types/values matched; original empty-array helper was explicitly rejected |
| [Empty-array translation evidence](evidence/empty-array-2026-10-25.md) | Historical PHP 8.3.13 v3 result: six actual return types/values matched, including the unchanged original empty-array helper |
| [Slash-helper translation unit](slash-helper-unit.md) | Isolated contract for unchanged `untrailingslashit()` and `trailingslashit()` |
| [Slash-helper translation evidence](evidence/slash-helpers-2026-10-25.md) | Historical PHP 8.3.13 isolated result: 13 string inputs and 26 typed return values matched |
| [Boolean-helper translation unit](boolean-helper-unit.md) | Separate contract for unchanged `wp_validate_boolean()` across 23 typed inputs |
| [Boolean-helper translation evidence](evidence/boolean-helper-2026-10-25.md) | Historical PHP 8.3.13 record for the Boolean contract; independent of earlier helper evidence |

Start each work session with current state and this index. Keep new work in an
explicit branch/PR in `katsuakikureyama/Gossip`. Do not replace the root tree with
the prototype. Ordinary verified merges are covered by the owner's standing
authorization; see the [current operating policy](reuse-boundary.md#publication-licensing-and-operating-policy).
