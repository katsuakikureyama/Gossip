# Historical README (through PR #44)

> Historical snapshot at `ae218d3f78cc649e4f2884ceb24870e0fe2d76b7`. It records
> evidence and instructions available at that point; it is not current project
> guidance. See the repository-root [`README.md`](../../README.md) and
> [`docs/INDEX.md`](../INDEX.md) for current material.

WordPress の原本と履歴を受け継ぎ、Go による実行環境と Go 製プラグインを探る、実験的な WordPress フォークです。

## 現在の到達点（2026-10-26）

検証済みの開発スタックを `master` に統合しています。
原文 **`add_freeform()` → `freeform()` → Block配列の追加**も[16ケースで完全なPHP原文と一致](../evidence/parser-add-freeform-2026-10-26.md)しました。長さ0のスキップと空Blockの追加、既存出力・NUL・UTF-8途中のバイト列を区別して確認しています。パーサ全体の解析処理ではありません。
パーサ原文の **`freeform()` メソッド**も選択的にGoへ変換し、[12組・24呼び出しで元の完全なPHPクラスと一致](../evidence/parser-freeform-2026-10-26.md)しました。HTML文字列の保持、返されるBlockの内容・独立性を確認したもので、パーサ全体の翻訳・ブロックコメント解析ではありません。
原文の **`WP_Block_Parser_Frame` と `WP_Block_Parser_Block` 全体**を一つのGo実行ファイルへ変換し、[別構成12ケースでPHPと一致](../evidence/parser-frame-2026-10-26.md)しました。Frameの省略引数・オフセット計算と、配列化後も同じBlockを共有することを確認しました。パーサの解析処理全体を実行した結果ではありません。
WordPress原文のデータクラス **`WP_Block_Parser_Block` 全体**をGoへ変換し、コンストラクタ→配列化→クラス要否判定まで[16ケースでPHPと一致](../evidence/parser-block-data-2026-10-26.md)しました。代用クラスは使っていません。ブロックパーサ全体やHTML/CSS描画の対応ではなく、以下の既存構成とは別の検証です。
選定した **26個の原文関数・334ケース**を一つのGo実行ファイルにまとめ、[PHPとの統合比較に成功](../evidence/elements-class-path-2026-10-26.md)しました。
実際のブロック要素スタイルのクラス要否判定（`wp_should_add_elements_class_name`）を、依存する`_wp_array_get`と原文のまま実行できています。CSS生成やレンダリング全体の対応ではありません。原文・既存ケース・過去の検証記録は保持しています。
さらに、[出力省略判定からクラス要否判定までの別構成24ケース](../evidence/elements-serialization-path-2026-10-26.md)もPHPと一致しました。原文3関数と入力データ用クラスを使った検証で、WordPressのブロック型登録や描画全体の代用ではありません。上記334ケースへ足し合わせた統合結果ではありません。
この入力データ用オブジェクトの生成・代入も[PHP→Go変換内で実行する経路](../evidence/elements-source-input-2026-10-26.md)へ進め、同じ24ケースが一致しました。新経路のGo観測コードからは、生成クラスのフィールドへの直接データ注入をなくしています。
WordPress全体の起動・DB・テーマ・既存プラグインの互換性や本番利用は未確認です。

[現在の状態](../current-state-fork-baseline.md) · [検証結果](../evidence/parser-add-freeform-2026-10-26.md) · [資料索引](../INDEX.md)

Gossip is an independent, experimental WordPress fork exploring a Go-powered
runtime and native Go plugins. This repository retains the original WordPress
source and history. It is not affiliated with or endorsed by the WordPress
project, and it is **not yet a Go replacement for WordPress**.

> 日本語: Gossip は WordPress の原本・履歴を保持した実験的なフォークです。
> Go / GoHP による実行と Go 製プラグインを検証しますが、現時点で WordPress
> 全体の Go 化や既存テーマ・プラグインとの互換性を実現したものではありません。

## What is here

- The root `wp-admin/`, `wp-includes/`, `wp-content/` and PHP entry points are
  upstream WordPress, not rewritten Gossip equivalents. The inherited
  [WordPress readme](../../readme.html) and [license](../../license.txt) remain unchanged.
- [`gossip/`](../../gossip/) is a separate experimental Go module. Initially it contains
  the native Plugin SDK and a differential harness that executes **this fork's
  original PHP hook files** and compares five owned cases with the Go registry.
  Separate original-source translation units convert six unchanged
  constant-return helpers and two unchanged slash helpers through pinned GoHP,
  build them, and compare narrowly defined behavior with PHP. The slash unit
  has [passed](../evidence/slash-helpers-2026-10-25.md) for 13 explicit
  string inputs (26 typed results); it verifies only `rtrim` with a literal
  two-byte slash/backslash mask. These are not full-core conversion, general
  `rtrim` compatibility, or WordPress path compatibility claims.
  A further, independently pinned contract targets the unchanged
  `wp_validate_boolean()` helper across 23 typed inputs. Its
  [contract](../boolean-helper-unit.md) and
  [execution record](../evidence/boolean-helper-2026-10-25.md) are separate
  from the earlier helper pins and do not extend them into a general PHP or
  WordPress compatibility claim.
  A fourth unit targets unchanged `zeroise()` with a bounded zero-padding
  `sprintf` subset. Its [contract](../zeroise-helper-unit.md) and
  [execution record](../evidence/zeroise-helper-2026-10-25.md) remain separate;
  this does not provide general PHP formatting or a combined ten-function test.
  The subsequent [combined-helper experiment](../helper-bundle-unit.md)
  translates these ten originals and five owned composition drivers together
  into one executable. Its new 22-case integration corpus and
  [execution record](../evidence/helper-bundle-2026-10-25.md) are separate
  from the standalone results; this still does not start WordPress core.
  A subsequent [newline-preservation unit](../newline-helper-unit.md)
  targets the original callback helper used by `wpautop` through bounded
  three-string replacement. Its separate
  [execution record](../evidence/newline-helper-2026-10-25.md) does not claim
  support for `wpautop`, regex callbacks or paragraph rendering as a whole.
  The [string-only stripslashes unit](../stripslashes-helper-unit.md) adds
  an original ternary helper, with non-string values returned unchanged.
  This is not recursive `stripslashes_deep` / `wp_unslash` support.
  A separate [integrated twelve-helper bundle](../integrated-helper-bundle-unit.md)
  brings all twelve selected originals and owned composition drivers into one
  executable on the combined conditional/string/array compiler. Its new corpus
  does not replace the independent unit or earlier ten-helper results.
  A seventh standalone [recursive wp_slash unit](../wp-slash-unit.md) targets
  the unchanged original through literal single-array callbacks and string-only
  `addslashes`. It observes nested lists and caller-input preservation, not a
  thirteen-function bundle or general callback/object traversal support.
  A later [thirteen-original bundle](../recursive-helper-bundle-unit.md) brings
  that recursive helper together with the previous twelve in one executable.
  Its 33-case corpus checks actual direct/composed returns and caller arguments;
  earlier standalone and bundle evidence remains separate. Shallow stripping of
  escaped lists is explicitly distinguished from unsupported `wp_unslash`.
  An eighth standalone [associative slice unit](../array-slice-assoc-unit.md)
  executes unchanged `wp_array_slice_assoc()`, including its `: array` return
  declaration. Fourteen cases observe ordered integer/string keys, null
  exclusion and caller-input preservation. Its [evidence](../evidence/array-slice-assoc-2026-10-25.md)
  remains a separate standalone result. The subsequent
  [fourteen-original bundle](../associative-helper-bundle-unit.md) integrates
  all fourteen originals and sixteen owned drivers into one executable.
  Its [55-case PASS](../evidence/associative-helper-bundle-2026-10-25.md)
  covers direct calls and compositions with ordered keys and caller arguments;
  it does not establish full-core or recursive unescaping compatibility.
  A ninth standalone [charset-name unit](../utf8-charset-unit.md) executes
  unchanged `_is_utf8_charset()` using ASCII case-insensitive comparison.
  Its [28-case result](../evidence/utf8-charset-2026-10-25.md) observes actual
  bool returns and unchanged inputs. This recognizes charset names, not valid
  UTF-8 text, and does not execute the option-dependent public wrapper.
  A subsequent [fifteen-original bundle](../charset-helper-bundle-unit.md)
  combines the published charset and typed-array compiler stacks and executes
  all fifteen originals plus nineteen drivers in one binary. Its new
  [89-case result](../evidence/charset-helper-bundle-2026-10-25.md) preserves
  the earlier configurations and adds slice/charset compositions; it does not
  establish whole-core compatibility.
  A tenth standalone [HEIC MIME-label unit](../heic-mime-unit.md) executes
  unchanged `wp_is_heic_image_mime_type()` with strict scalar membership.
  Its [24-case result](../evidence/heic-mime-2026-10-25.md) compares actual
  bool returns and unchanged inputs. This recognizes exact MIME labels only;
  it does not inspect, decode or convert image contents.
  The next [sixteen-original bundle](../heic-helper-bundle-unit.md) combines
  the HEIC predicate with the earlier helpers and22 owned drivers in one
  binary. Its separate119-case configuration covers direct calls and
  slice/unescape/Boolean compositions without changing earlier pins.
  The [119-case result](../evidence/heic-helper-bundle-2026-10-25.md) matches
  actual typed returns, ordered arrays and caller arguments against PHP.
  An eleventh standalone [yes/no-label unit](../bool-yn-unit.md) targets
  unchanged `bool_from_yn()` with24 actual-string inputs. It distinguishes
  the exact labels `y`/`Y` from words, digit strings, whitespace and non-ASCII
  text; this is not general Boolean parsing or a17-original integrated run.
  Its [24-case result](../evidence/bool-yn-2026-10-25.md) matches actual
  Boolean returns and unchanged inputs against PHP8.3.36.
  The [seventeen-original bundle](../yn-helper-bundle-unit.md) adds that
  original to one binary with25 owned drivers and149 cases. New compositions
  cover selected-array yes/no classification and the distinct behavior of
  `bool_from_yn()` and `wp_validate_boolean()` on the same input string.
  Its [149-case result](../evidence/yn-helper-bundle-2026-10-25.md) matches
  actual types/values, ordered keys and caller arguments against PHP.
  The next [numeric-array boundary](../numeric-array-boundary.md) establishes
  a24-case PHP-only reference using the original `array_all` polyfill. Both
  selected declarations are rejected by its older pinned GoHP; its
  [boundary record](../evidence/numeric-array-boundary-2026-10-25.md) is
  explicitly not an additional Go-compatible helper or integrated result.
  A separate [numeric-array helper unit](../numeric-array-helper-unit.md) now
  pins the arrow/callable compiler and translates both unchanged original
  declarations. Its [31-case comparison](../evidence/numeric-array-helper-2026-10-25.md)
  passes against PHP8.3.36 (24 predicate +7 direct polyfill cases), without
  changing the earlier17-original integrated result or claiming bootstrap support.
  A new [nineteen-original bundle](../numeric-helper-bundle-unit.md) integrates
  both on the combined GoHP compiler stack. Its
  [192-case result](../evidence/numeric-helper-bundle-2026-10-25.md) passes in
  one binary:149 retained legacy,31 numeric/polyfill and12 cross-helper cases.
  This is bounded declaration-level integration, not full WordPress startup.
  A separate [array-search unit](../array-search-polyfills-unit.md) translates
  unchanged `array_any`, `array_find` and `array_find_key` polyfills with
  [30 passing cases](../evidence/array-search-polyfills-2026-10-26.md).
  It checks actual scalar return types, first-match behavior and caller inputs;
  it is not yet a larger integrated bundle.
  A new [twenty-two-original bundle](../search-helper-bundle-unit.md) now
  integrates all of those declarations in one binary with44 owned drivers.
  Its [230-case result](../evidence/search-helper-bundle-2026-10-26.md) includes
  the retained192 and30 cases plus8 new selection/Boolean/search compositions,
  distinguishing matched false, no-match null and integer key0. This does not
  establish WordPress bootstrap, HTTP, database, theme or plugin compatibility.
  The next [array-edge boundary](../array-edge-boundary.md) fixes24 native
  PHP reference cases for unchanged `array_first`/`array_last`. Its
  [record](../evidence/array-edge-boundary-2026-10-26.md) confirms missing
  `empty` and `array_key_last` translation support; neither is counted as an
  additional Go-compatible original.
  A later [array_first unit](../array-first-unit.md) adopts the integrated
  empty-capable GoHP main and [passes all12 retained cases](../evidence/array-first-2026-10-26.md).
  The historical boundary above still describes its older compiler pin;
  array_last was still blocked on that pin. The22-original/230-case integrated result retains
  its own configuration.
  The subsequent [two-original edge unit](../array-edge-unit.md) uses GoHP #58
  and [passes all24 cases](../evidence/array-edge-unit-2026-10-26.md) for both
  array_first and array_last together. Both original declarations and the native
  guarded blocks are preserved unchanged.
  The [twenty-four-original integration](../edge-helper-bundle-unit.md) now
  combines all24 declarations and46 owned drivers into one executable on
  GoHP #58. Its [262-case result](../evidence/edge-helper-bundle-2026-10-26.md)
  includes all230 earlier cases,24 edge cases and8 new Boolean/edge/search
  compositions. Earlier standalone and integrated evidence retains its pins.
  The [nested settings lookup unit](../array-get-unit.md) now executes unchanged
  _wp_array_get used by Theme JSON and block settings. Its retained24 cases
  [match Go/PHP](../evidence/array-get-unit-2026-10-26.md), including actual
  nested returns and unchanged caller arguments. GoHP #63/#64 supplied bounded
  defaults/isset and the literal-null fix exposed by the first execution.
  This adds a25th successful original separately, not a25-original integration.
  The earlier [rejection boundary](../array-get-boundary.md) keeps its old pin.
  The new [twenty-five-original integration](../settings-helper-bundle-unit.md)
  combines those originals and48 owned drivers into one executable on GoHP #64.
  All [294 fresh comparisons](../evidence/settings-helper-bundle-2026-10-26.md)
  pass:262 retained edge-bundle cases,24 lookup cases and8 new settings
  compositions. This replaces neither the earlier pins/evidence nor the original
  WordPress implementation, and does not establish full Theme JSON execution.
  The [real element-class decision path](../elements-class-path.md) adds the
  unchanged `wp_should_add_elements_class_name`, without resolver/class/service
  stubs. All [334 integrated cases](../evidence/elements-class-path-2026-10-26.md)
  pass in one binary:294 retained and40 actual caller cases. This measures class
  necessity only, not CSS/HTML generation or the surrounding WordPress render hook.
- [`Gossip-prototype`](https://github.com/katsuakikureyama/Gossip-prototype) keeps
  the earlier independent CMS prototype, its PRs and historical evidence. Its
  custom editor, JSON store, importer and server have not replaced this fork's
  WordPress implementation.

The pinned baseline is WordPress commit
`54106f6c545fc8b881f6b39b7963002bece6c28c`, reporting `7.2-beta1-64693`.
This is a development snapshot, **not a stable production release**. The old
prototype's WordPress 6.6.2 results do not establish compatibility for this tree.
See [current state](../current-state-fork-baseline.md) for measured results.

## Local development

Run from the nested Go module (Go 1.22 or later):

```sh
cd gossip
go test ./...
```

This verifies the source pin and SDK tests; the Docker differential explicitly
skips unless enabled. New Docker-oracle runs require the admitted, immutable
official PHP 8.3 runner, pulled explicitly before the network-isolated tests:

```sh
docker pull php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd
GOSSIP_PHP_RUNNER_IMAGE=php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd \
  go test -race -count=1 -v ./...
```

This enables hooks; GoHP units additionally require `GOSSIP_GOHP_SOURCE` to
point at a clone containing the existing three-unit compiler pin
`4e1b5abfe019ec121e13d85e99e51c206e765533`. The three GoHP units retain
separate fixtures and comparison boundaries even though they now select that
same compiler. See the [unified-compiler execution record](../evidence/unified-compiler-2026-10-25.md)
and [current state](../current-state-fork-baseline.md) for the current run;
the earlier [PHP 8.3.36 record](../evidence/php-83-floor-2026-10-25.md) and
PHP 8.3.13 records remain historical evidence, not a substitute for this run.
The fourth, zeroise unit requires its own newer compiler pin from
[`zeroise-helper.json`](../../gossip/tests/gohp/zeroise-helper.json); it does not
repin or rerun the three unchanged units.
The newline-preservation unit additionally needs the pin in
[`newline-helper.json`](../../gossip/tests/gohp/newline-helper.json). Earlier standalone
and combined-unit pins remain separate and unchanged.
The string-only stripslashes unit uses its additional pin in
[`stripslashes-helper.json`](../../gossip/tests/gohp/stripslashes-helper.json).
The recursive wp_slash unit uses the additional pin in
[`wp-slash-helper.json`](../../gossip/tests/gohp/wp-slash-helper.json).
The associative slice unit uses its own pin in
[`array-slice-assoc-helper.json`](../../gossip/tests/gohp/array-slice-assoc-helper.json).
The fourteen-original bundle selects its configuration in
[`associative-helper-bundle.json`](../../gossip/tests/gohp/associative-helper-bundle.json).
The charset-name unit uses
[`utf8-charset-helper.json`](../../gossip/tests/gohp/utf8-charset-helper.json), based on
the separate logical/comparison compiler stack; earlier pins stay unchanged.
The new merged-stack bundle uses
[`charset-helper-bundle.json`](../../gossip/tests/gohp/charset-helper-bundle.json),
without repinning the older standalone or combined experiments.
The HEIC MIME-label unit separately pins its strict-membership compiler in
[`heic-mime-helper.json`](../../gossip/tests/gohp/heic-mime-helper.json).
The sixteen-original integration has its own
[`heic-helper-bundle.json`](../../gossip/tests/gohp/heic-helper-bundle.json), using
the same published strict-membership compiler without changing that compiler.
The yes/no-label unit uses
[`bool-yn-helper.json`](../../gossip/tests/gohp/bool-yn-helper.json) with that same
compiler, keeping all earlier unit and bundle pins separate.
The seventeen-original configuration is separately pinned in
[`yn-helper-bundle.json`](../../gossip/tests/gohp/yn-helper-bundle.json).

The official PHP image provides the interpreter, not a bundled WordPress
checkout. Only verified source bytes from this checkout are staged into it. The
container has no network, no database, a read-only filesystem, non-root UID,
resource limits and automatic removal. Nothing pulls or executes third-party
plugin ZIPs or Composer scripts. See the [PHP oracle policy](../php-oracle-policy.md)
for image admission, the native-version guard and the distinction between a
runner being allowed and a compatibility result.

## Direction and limits

Reuse original WordPress behavior and assets where validated, preserve its
provenance, and grow the Go boundary only behind explicit differential evidence.
GoHP is a separate compiler project; its current candidate cannot translate or
run the original WordPress hook engine, despite the narrow helper-unit PASS.
No production deployment, migration,
marketplace, full-theme rendering or general plugin compatibility is claimed.

Development changes go through a PR **in this fork**, not an upstream
WordPress PR. Use Draft while work is incomplete; integrate verified changes
into master at bounded milestones. A merge is not a production deployment. See the [documentation index](../INDEX.md),
[reuse boundary](../reuse-boundary.md), and
[source provenance](../provenance.md).
