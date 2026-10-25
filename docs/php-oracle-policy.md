# Docker PHP oracle policy

This policy defines which native PHP interpreters may execute new Gossip Docker
oracle runs. It is an admission control for test infrastructure, not a PHP or
WordPress compatibility result.

## Current admission floor

New hook-comparison and GoHP-unit oracle runs use the official PHP image pinned
to:

```text
php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd
```

The admitted language series is PHP 8.3 and the native minimum is 8.3.36. The
image digest is checked against the exact allowlist before the container starts.
Inside the same PHP process that will require the original source, the oracle
checks `PHP_VERSION` before that require. Hooks record the version from their
oracle response; the GoHP units record a guarded version probe using the same
immutable image and repeat the guard in every comparison process. A separate
probe alone never authorizes the interpreter that executes the comparison.

The guard rejects PHP 7.4, every 8.3 patch below 8.3.36, 8.4 and later,
pre-release versions, malformed versions and any image digest outside the
allowlist. An allowed image still does not prove the requested code translates,
compiles, behaves identically, or is secure for production.

The selected floor follows the [official PHP 8.3.36 release](https://www.php.net/releases/8_3_36.php).
Moving the floor or admitting a new major/minor series requires a new explicit
policy, image pin, native guard, execution evidence and documentation update in
the same change set.

## Scope and evidence

The shared `internal/phporacle` helper applies this admission process to all
Gossip Docker PHP oracles: the hook comparison, seven standalone original-source
GoHP units (constant/array returns, slash helpers, Boolean validation, zeroise,
newline preservation, string-only stripslashes and recursive wp_slash), and the separate ten-helper
and integrated twelve-helper bundle experiments. Each
unit retains its own WordPress excerpt, fixture and behavior contract. The
current unified compiler pin is `4e1b5abfe019ec121e13d85e99e51c206e765533`,
with Gossip test commit `e943cbce1f74c93b5ca990de7a544caabb274aa5`; using one
compiler does not merge the three units into one source bundle or whole-core
comparison. It backports only the PHP 8.3 policy commit
`948647c2903a4a1b22069818861062df1f5bf414` from separate GoHP PR #14 to Draft
PR #16, which is stacked on #13. Composer inventory, `strlen`, `array_keys`
and `array_values` work are not included. These details describe the first three
units. The fourth [zeroise unit](zeroise-helper-unit.md) pins a newer descendant
compiler with bounded `sprintf` support and reuses the same admission policy;
it has its own [execution record](evidence/zeroise-helper-2026-10-25.md).
The later [combined helper unit](helper-bundle-unit.md) reuses the same image,
admission guard and compiler as zeroise, with a separate 22-case integration
configuration. This does not alter the standalone unit pins or policy.
The later [newline-preservation unit](newline-helper-unit.md) also uses this
same admission policy and approved image. Its new compiler pin and 16-case
corpus have separate evidence; the earlier units are not repinned.
The [string-only stripslashes unit](stripslashes-helper-unit.md) uses the same
guard and image with another independent compiler pin and typed corpus.
The [integrated twelve-helper bundle](integrated-helper-bundle-unit.md) retains
the same admission policy/image and has a new combined-compiler pin and corpus.
It does not change the six standalone or previous ten-helper configuration.
The later [recursive wp_slash unit](wp-slash-unit.md) reuses the same admission
policy and image, with its own compiler pin and 22-case corpus. Earlier pins
and evidence remain unchanged.

Historical records executed with PHP 8.3.13 remain retained as historical
evidence. They are useful provenance, but are below the current native floor
and must not be reported as a current-policy run. The new PHP 8.3.36 record is
separate in [PHP 8.3 floor evidence](evidence/php-83-floor-2026-10-25.md).
The new unified-compiler run is recorded separately in
[unified-compiler evidence](evidence/unified-compiler-2026-10-25.md). Until
that record contains an executed result for its stated scope, this policy does
not establish an outcome for the affected comparisons.

## Capability exclusion

PHP-oracle admission and its Docker execution are opt-in local developer test
infrastructure. They are explicitly excluded from WebMCP publication: they add
no user-facing operation, route, browser session, cookie forwarding, automatic
submit, deployment, publication, permission model or audit API.
