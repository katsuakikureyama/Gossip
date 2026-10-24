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
