# Gossip

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
  [WordPress readme](readme.html) and [license](license.txt) remain unchanged.
- [`gossip/`](gossip/) is a separate experimental Go module. Initially it contains
  the native Plugin SDK and a differential harness that executes **this fork's
  original PHP hook files** and compares five owned cases with the Go registry.
  Separate original-source translation units convert six unchanged
  constant-return helpers and two unchanged slash helpers through pinned GoHP,
  build them, and compare narrowly defined behavior with PHP. The slash unit
  has [passed](docs/evidence/slash-helpers-2026-10-25.md) for 13 explicit
  string inputs (26 typed results); it verifies only `rtrim` with a literal
  two-byte slash/backslash mask. These are not full-core conversion, general
  `rtrim` compatibility, or WordPress path compatibility claims.
  A further, independently pinned contract targets the unchanged
  `wp_validate_boolean()` helper across 23 typed inputs. Its
  [contract](docs/boolean-helper-unit.md) and
  [execution record](docs/evidence/boolean-helper-2026-10-25.md) are separate
  from the earlier helper pins and do not extend them into a general PHP or
  WordPress compatibility claim.
  A fourth unit targets unchanged `zeroise()` with a bounded zero-padding
  `sprintf` subset. Its [contract](docs/zeroise-helper-unit.md) and
  [execution record](docs/evidence/zeroise-helper-2026-10-25.md) remain separate;
  this does not provide general PHP formatting or a combined ten-function test.
  The subsequent [combined-helper experiment](docs/helper-bundle-unit.md)
  translates these ten originals and five owned composition drivers together
  into one executable. Its new 22-case integration corpus and
  [execution record](docs/evidence/helper-bundle-2026-10-25.md) are separate
  from the standalone results; this still does not start WordPress core.
  A subsequent [newline-preservation unit](docs/newline-helper-unit.md)
  targets the original callback helper used by `wpautop` through bounded
  three-string replacement. Its separate
  [execution record](docs/evidence/newline-helper-2026-10-25.md) does not claim
  support for `wpautop`, regex callbacks or paragraph rendering as a whole.
  The [string-only stripslashes unit](docs/stripslashes-helper-unit.md) adds
  an original ternary helper, with non-string values returned unchanged.
  This is not recursive `stripslashes_deep` / `wp_unslash` support.
  A separate [integrated twelve-helper bundle](docs/integrated-helper-bundle-unit.md)
  brings all twelve selected originals and owned composition drivers into one
  executable on the combined conditional/string/array compiler. Its new corpus
  does not replace the independent unit or earlier ten-helper results.
  A seventh standalone [recursive wp_slash unit](docs/wp-slash-unit.md) targets
  the unchanged original through literal single-array callbacks and string-only
  `addslashes`. It observes nested lists and caller-input preservation, not a
  thirteen-function bundle or general callback/object traversal support.
  A later [thirteen-original bundle](docs/recursive-helper-bundle-unit.md) brings
  that recursive helper together with the previous twelve in one executable.
  Its 33-case corpus checks actual direct/composed returns and caller arguments;
  earlier standalone and bundle evidence remains separate. Shallow stripping of
  escaped lists is explicitly distinguished from unsupported `wp_unslash`.
  An eighth standalone [associative slice unit](docs/array-slice-assoc-unit.md)
  executes unchanged `wp_array_slice_assoc()`, including its `: array` return
  declaration. Fourteen cases observe ordered integer/string keys, null
  exclusion and caller-input preservation. Its [evidence](docs/evidence/array-slice-assoc-2026-10-25.md)
  remains a separate standalone result. The subsequent
  [fourteen-original bundle](docs/associative-helper-bundle-unit.md) integrates
  all fourteen originals and sixteen owned drivers into one executable.
  Its [55-case PASS](docs/evidence/associative-helper-bundle-2026-10-25.md)
  covers direct calls and compositions with ordered keys and caller arguments;
  it does not establish full-core or recursive unescaping compatibility.
  A ninth standalone [charset-name unit](docs/utf8-charset-unit.md) executes
  unchanged `_is_utf8_charset()` using ASCII case-insensitive comparison.
  Its [28-case result](docs/evidence/utf8-charset-2026-10-25.md) observes actual
  bool returns and unchanged inputs. This recognizes charset names, not valid
  UTF-8 text, and does not execute the option-dependent public wrapper.
  A subsequent [fifteen-original bundle](docs/charset-helper-bundle-unit.md)
  combines the published charset and typed-array compiler stacks and executes
  all fifteen originals plus nineteen drivers in one binary. Its new
  [89-case result](docs/evidence/charset-helper-bundle-2026-10-25.md) preserves
  the earlier configurations and adds slice/charset compositions; it does not
  establish whole-core compatibility.
  A tenth standalone [HEIC MIME-label unit](docs/heic-mime-unit.md) executes
  unchanged `wp_is_heic_image_mime_type()` with strict scalar membership.
  Its [24-case result](docs/evidence/heic-mime-2026-10-25.md) compares actual
  bool returns and unchanged inputs. This recognizes exact MIME labels only;
  it does not inspect, decode or convert image contents.
  The next [sixteen-original bundle](docs/heic-helper-bundle-unit.md) combines
  the HEIC predicate with the earlier helpers and22 owned drivers in one
  binary. Its separate119-case configuration covers direct calls and
  slice/unescape/Boolean compositions without changing earlier pins.
  The [119-case result](docs/evidence/heic-helper-bundle-2026-10-25.md) matches
  actual typed returns, ordered arrays and caller arguments against PHP.
  An eleventh standalone [yes/no-label unit](docs/bool-yn-unit.md) targets
  unchanged `bool_from_yn()` with24 actual-string inputs. It distinguishes
  the exact labels `y`/`Y` from words, digit strings, whitespace and non-ASCII
  text; this is not general Boolean parsing or a17-original integrated run.
  Its [24-case result](docs/evidence/bool-yn-2026-10-25.md) matches actual
  Boolean returns and unchanged inputs against PHP8.3.36.
  The [seventeen-original bundle](docs/yn-helper-bundle-unit.md) adds that
  original to one binary with25 owned drivers and149 cases. New compositions
  cover selected-array yes/no classification and the distinct behavior of
  `bool_from_yn()` and `wp_validate_boolean()` on the same input string.
  Its [149-case result](docs/evidence/yn-helper-bundle-2026-10-25.md) matches
  actual types/values, ordered keys and caller arguments against PHP.
  The next [numeric-array boundary](docs/numeric-array-boundary.md) establishes
  a24-case PHP-only reference using the original `array_all` polyfill. Both
  selected declarations are rejected by its older pinned GoHP; its
  [boundary record](docs/evidence/numeric-array-boundary-2026-10-25.md) is
  explicitly not an additional Go-compatible helper or integrated result.
  A separate [numeric-array helper unit](docs/numeric-array-helper-unit.md) now
  pins the arrow/callable compiler and translates both unchanged original
  declarations. Its [31-case comparison](docs/evidence/numeric-array-helper-2026-10-25.md)
  passes against PHP8.3.36 (24 predicate +7 direct polyfill cases), without
  changing the earlier17-original integrated result or claiming bootstrap support.
- [`Gossip-prototype`](https://github.com/katsuakikureyama/Gossip-prototype) keeps
  the earlier independent CMS prototype, its PRs and historical evidence. Its
  custom editor, JSON store, importer and server have not replaced this fork's
  WordPress implementation.

The pinned baseline is WordPress commit
`54106f6c545fc8b881f6b39b7963002bece6c28c`, reporting `7.2-beta1-64693`.
This is a development snapshot, **not a stable production release**. The old
prototype's WordPress 6.6.2 results do not establish compatibility for this tree.
See [current state](docs/current-state-fork-baseline.md) for measured results.

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
same compiler. See the [unified-compiler execution record](docs/evidence/unified-compiler-2026-10-25.md)
and [current state](docs/current-state-fork-baseline.md) for the current run;
the earlier [PHP 8.3.36 record](docs/evidence/php-83-floor-2026-10-25.md) and
PHP 8.3.13 records remain historical evidence, not a substitute for this run.
The fourth, zeroise unit requires its own newer compiler pin from
[`zeroise-helper.json`](gossip/tests/gohp/zeroise-helper.json); it does not
repin or rerun the three unchanged units.
The newline-preservation unit additionally needs the pin in
[`newline-helper.json`](gossip/tests/gohp/newline-helper.json). Earlier standalone
and combined-unit pins remain separate and unchanged.
The string-only stripslashes unit uses its additional pin in
[`stripslashes-helper.json`](gossip/tests/gohp/stripslashes-helper.json).
The recursive wp_slash unit uses the additional pin in
[`wp-slash-helper.json`](gossip/tests/gohp/wp-slash-helper.json).
The associative slice unit uses its own pin in
[`array-slice-assoc-helper.json`](gossip/tests/gohp/array-slice-assoc-helper.json).
The fourteen-original bundle selects its configuration in
[`associative-helper-bundle.json`](gossip/tests/gohp/associative-helper-bundle.json).
The charset-name unit uses
[`utf8-charset-helper.json`](gossip/tests/gohp/utf8-charset-helper.json), based on
the separate logical/comparison compiler stack; earlier pins stay unchanged.
The new merged-stack bundle uses
[`charset-helper-bundle.json`](gossip/tests/gohp/charset-helper-bundle.json),
without repinning the older standalone or combined experiments.
The HEIC MIME-label unit separately pins its strict-membership compiler in
[`heic-mime-helper.json`](gossip/tests/gohp/heic-mime-helper.json).
The sixteen-original integration has its own
[`heic-helper-bundle.json`](gossip/tests/gohp/heic-helper-bundle.json), using
the same published strict-membership compiler without changing that compiler.
The yes/no-label unit uses
[`bool-yn-helper.json`](gossip/tests/gohp/bool-yn-helper.json) with that same
compiler, keeping all earlier unit and bundle pins separate.
The seventeen-original configuration is separately pinned in
[`yn-helper-bundle.json`](gossip/tests/gohp/yn-helper-bundle.json).

The official PHP image provides the interpreter, not a bundled WordPress
checkout. Only verified source bytes from this checkout are staged into it. The
container has no network, no database, a read-only filesystem, non-root UID,
resource limits and automatic removal. Nothing pulls or executes third-party
plugin ZIPs or Composer scripts. See the [PHP oracle policy](docs/php-oracle-policy.md)
for image admission, the native-version guard and the distinction between a
runner being allowed and a compatibility result.

## Direction and limits

Reuse original WordPress behavior and assets where validated, preserve its
provenance, and grow the Go boundary only behind explicit differential evidence.
GoHP is a separate compiler project; its current candidate cannot translate or
run the original WordPress hook engine, despite the narrow helper-unit PASS.
No production deployment, migration,
marketplace, full-theme rendering or general plugin compatibility is claimed.

Development changes go through a Draft PR **in this fork**, not an upstream
WordPress PR. See the [documentation index](docs/INDEX.md),
[reuse boundary](docs/reuse-boundary.md), and
[source provenance](docs/provenance.md).
