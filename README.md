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
  A separate [original-source translation test](docs/translation-units.md)
  converts six unchanged WordPress constant-return helpers through pinned
  GoHP, builds them, and compares five scalar echo observations plus all six
  actual return types/values with PHP. That narrow path has
  [passed](docs/evidence/empty-array-2026-10-25.md), including the original
  empty-array helper without rewriting its `array()` expression. This is not
  full-core conversion or general array compatibility.
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
skips unless enabled. To run it using the already pulled, immutable PHP runner:

```sh
GOSSIP_PHP_RUNNER_IMAGE=wordpress@sha256:04be81de853887fca707742785b32bf93fb4a692246dd1bc90b3790105071a9a \
  go test -race -count=1 -v ./...
```

The container's bundled WordPress is **not** the oracle. Only verified source
bytes from this checkout are staged into it. The container has no network, no
database, a read-only filesystem, non-root UID, resource limits and automatic
removal. Nothing pulls or executes third-party plugin ZIPs or Composer scripts.

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
