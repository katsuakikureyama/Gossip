# Gossip

Gossip is an experimental fork of the official WordPress Git mirror. It keeps
the inherited WordPress source tree and history byte-for-byte intact while
isolating Go experiments under [`gossip/`](gossip/). The project explores a Go
runtime and native Go plugins; it is not a WordPress replacement, a production
deployment, or a claim of theme or plugin compatibility.

> 日本語: Gossip は WordPress 公式 Git ミラーを元にした実験的フォークです。
> WordPress 原本ツリーと履歴を変更せず、`gossip/` 配下に Go 実験を隔離します。
> WordPress 全体の Go 化、既存テーマ／プラグイン互換、本番利用は未実現です。

## Project boundaries

| Layer | Responsibility |
| --- | --- |
| WordPress root tree | Inherited original source, assets, history, and license notices; not a Gossip rewrite. |
| `gossip/` | Go runtime experiments, native Go Plugin SDK, fixtures, and differential harnesses. |
| [GoHP](https://github.com/katsuakikureyama/GoHP) | A general PHP-to-Go transpiler project. Gossip is its first measured application consumer, not its definition of PHP compatibility. |

The inherited [WordPress readme](readme.html) and [license](license.txt) remain
unchanged. Project documentation lives under [`docs/`](docs/), outside the
original source directories. Exact ownership, reuse, and exclusion boundaries
are recorded in [the reuse boundary](docs/reuse-boundary.md).

Gossip is independent and is not affiliated with or endorsed by WordPress or
the WordPress Foundation. “WordPress” is used only to identify the inherited
upstream project. See the WordPress Foundation's
[trademark policy](https://wordpressfoundation.org/trademark-policy/).

## Source baseline

The preserved baseline is WordPress commit
[`54106f6c545fc8b881f6b39b7963002bece6c28c`](https://github.com/WordPress/WordPress/tree/54106f6c545fc8b881f6b39b7963002bece6c28c),
whose `wp-includes/version.php` reports `7.2-beta1-64693`. It is a development
snapshot, not a current stable or production WordPress release. The exact
source hashes, retained prototype provenance, and verification procedure are
in [source provenance](docs/provenance.md).

## Measured evidence

Current priority: **make one fixed Gossip source work on PHP 8.3**, starting with
the complete original block parser. The [application gates and blocker list](docs/php83-application-acceptance.md)
define the scope and subsequent bootstrap/DB/HTTP/site workflows. **Full parser
translation is blocked**; native-reference capture is not generated-Go success.
The [new native full-parser reference](docs/evidence/parser-full-baseline-2026-10-27.md)
records 12 cases / 24 original calls; it is not added to the Go comparison counts.
General storage optimization and new PHP-target expansion are deferred.

The following are separate, bounded configurations. Counts must not be added
together: their fixtures, GoHP revisions, drivers, and retained cases differ.
Each record states its original source range, compiler revision, PHP oracle,
commands, hashes, and exclusions.

| Configuration | Observed result | Evidence |
| --- | --- | --- |
| Original-helper integration | 26 unchanged WordPress helpers; 334 PHP comparisons passed in one executable | [Element-class path](docs/elements-class-path.md) |
| Source-built input | 24 comparisons passed for translated object construction/property assignment | [Source-built input](docs/elements-source-input.md) |
| `WP_Block_Parser_Block` | 16 comparisons passed for the complete data class/constructor and bounded path | [Block data](docs/parser-block-data.md) |
| `WP_Block_Parser_Frame` | 12 comparisons passed for the complete Frame and nested Block classes | [Frame](docs/parser-frame.md) |
| `WP_Block_Parser::freeform` | 12 paired cases / 24 method calls passed using the complete original PHP parser class | [Freeform](docs/parser-freeform.md) |
| `WP_Block_Parser::add_freeform` | 16 comparisons passed for append/skip behavior with the original Block class | [Add freeform](docs/parser-add-freeform.md) |

These retained comparisons use PHP 8.3.36, not an inferred PHP 8.4 result.
They do not execute the complete WordPress bootstrap, parser loop, database,
HTTP, themes, or third-party plugins. The freeform unit omits six other parser
methods; the add_freeform unit includes freeform and omits the remaining five.
Omitted methods are a declared boundary, not replacement implementations.

Separately, the [parser-token reference](docs/parser-token-reference.md) records
24 native PHP 8.3.36 cases for the complete original `next_token()` and a direct
`preg_match()` probe using its unchanged pattern: 17 matches, 6 non-matches and
1 backtracking-limit failure. Capture bytes/positions, errors and parser state
are retained. This is **PHP reference evidence only**, not Go compatibility or a
regex-backend comparison, and is not added to the translation counts above.

The follow-up [PCRE2 experiment](docs/parser-pcre2-comparison.md) independently
links PCRE2 10.47 and matches all 24 retained direct-regex results, including
ordered capture bytes/offsets and the limit failure. This is a bounded C-engine
comparison against PHP's PCRE 10.42, **not** GoHP/cgo integration, full parser
execution or a backend adoption decision. It adds no generated-Go count.

The [regexp2 comparison](docs/parser-regexp2-comparison.md) records 22 equal
outputs, 1 UTF-8-interior-offset API difference and 1 non-comparable PCRE-limit
case. That direct string-API profile is not a drop-in backend; no dependency has
been adopted. These are engine observations, not additional translated functions.

The [bridge implementation contract](docs/pcre2-bridge-contract.md) separates an
optional PCRE2 Go/C bridge from PHP output-argument/state semantics and compiler
admission. GoHP has now implemented those bounded experimental stages through
[PR #114](https://github.com/katsuakikureyama/GoHP/pull/114); default admission
remains off and the CGO-disabled build/reference-storage boundaries are preserved.

The new [generated regex consumer](docs/parser-preg83-generated.md) connects a
Gossip-owned PHP driver to that pinned compiler/runtime using the unchanged
original pattern bytes: **23/23 fixed-limit comparisons passed**; the separate
limit=1 case is explicitly excluded. This is not an additional original WordPress
function or a translation of `next_token()`, and earlier fixture pins stay unchanged.

The current work state and links to each immutable execution record are in
[current state](docs/current-state-fork-baseline.md) and
[the documentation index](docs/INDEX.md). The extraction, source-byte checks,
and per-unit comparison contract are in
[original-source translation units](docs/translation-units.md).

## Local development

Use Go 1.22 or later from the nested module:

```sh
cd gossip
go test ./...
```

By default this runs local source-pin and SDK tests. The Docker PHP oracle is
intentionally skipped unless its explicitly admitted immutable runner image is
selected; a skipped oracle is not a compatibility result. The complete runner
admission, isolation, and invocation policy is in
[PHP oracle policy](docs/php-oracle-policy.md).

Each original-source fixture pins its own GoHP revision in its fixture metadata.
Do not reuse a compiler revision or command from another fixture: consult the
linked translation-unit document and its execution record for the exact
configuration.

## Direction

Gossip uses the active [PHP 8.3 application gate](docs/php83-application-acceptance.md)
to prioritize GoHP work. [GoHP PR130](https://github.com/katsuakikureyama/GoHP/pull/130), now merged to main, adds an opt-in,
checked static-declaration closure: it verifies the complete original parser,
Block, and Frame files through an explicit manifest/root/entry and admits their
original `require_once` directives at403/410 without source flattening. Its
[three-file admission probe](docs/evidence/parser-static-closure-probe-2026-10-29.md)
now reaches a semantic rejection at parser line209 (`$stack_top->prev_offset`
after `array_pop`), with no generated parser artifacts. A separate owned
multi-file fixture passes native PHP8.3/8.4 comparisons; it does not establish
general runtime include/autoload support or full-parser acceptance. Stack mutation/unknown receivers, snapshots,
tokenizer, and JSON integration remain required.
[The loading prerequisites](https://github.com/katsuakikureyama/GoHP/blob/main/docs/file-aware-locations.md)
retain per-file diagnostic/map identity; the new closure mode builds on that
infrastructure rather than treating it as executable-loading proof by itself.
The optional regex experiment alone does not complete that unit. Backend risks
remain documented in [the parser backend assessment](docs/parser-backend-assessment.md).

The project map distinguishes what remains original-source translation, what
is a native Go integration layer, and what is intentionally unsupported or
unmeasured. See [the reuse boundary](docs/reuse-boundary.md).

The inherited WordPress licensing remains with its original material. GoHP's
license for its own code and any downstream runtime/translation arrangements
remain under review; Gossip does not represent that policy as finalized or as
legal advice.

Changes are developed in this fork through scoped branches and evidence-bearing
pull requests. A merge records a bounded milestone; it is not a production
deployment. For the prior, longer operational narrative through PR #44, see
[the historical README](docs/history/readme-through-pr44.md).
