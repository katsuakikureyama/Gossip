# Original WordPress helper translation — 2026-10-25 JST

## Targets

- Gossip harness/fixture: `1c84d0c736b682b81f3a267f4c5b2c6269461308`.
  The exact tested source was committed without changes after PASS.
- Original WordPress: `54106f6c545fc8b881f6b39b7963002bece6c28c`,
  version `7.2-beta1-64693`. The inherited tree was not changed.
- GoHP compiler/runtime: `9c7389d35032d623a429198a69cdfa5c11571e8e`.
- Source unit: five literal-return helpers from `wp-includes/functions.php`,
  selected with exact original line ranges and whole-file/excerpt hashes.
- Host driver: Go 1.22.1, macOS arm64. Isolated builder: Go 1.25.14 linux/arm64.
- Exact source, image, compiler, generated report/source and binary hashes:
  [machine-readable observation](return-helpers-2026-10-25.json).

This is a new result, not a reuse of the native-SDK hook comparison. Here the
original PHP function bodies are translated by GoHP and their generated binary
is executed. The unit assembly and exclusions are in
[the contract](../translation-units.md).

## Commands and results

Pinned GoHP was cloned into `/private/tmp/gossip-gohp.y33KQk/GoHP`; its candidate
branch resolved to the commit above. The test exports the pinned commit object,
not the checkout's working files. Required Go/PHP images were already cached;
no image/toolchain/dependency download occurs inside the test.

From `/Users/kureyamayuri/Desktop/Gossip-wordpress/gossip`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gossip-gohp.y33KQk/GoHP \
  go test -count=1 -v ./tests/gohp
```

| Attempt / configuration | Result |
| --- | --- |
| Initial uncommitted harness, builder `/tmp` 16 MiB | Infrastructure FAIL: Go runtime compilation reported `no space left on device`; PHP translation had not started. No compatibility conclusion. |
| Final harness at `1c84d0c`, builder `/tmp` 256 MiB | PASS in 7.89 seconds for the integration test; generation, verified report hashes, build and PHP/generated execution succeeded. |

Only the temporary build-space limit changed after the initial failure. The
same unit was rerun once with the corrected limit; no tests were repeated after
PASS. Existing source-baseline/native-hook results remain at their own commits.

Final stage results:

1. `generated`: GoHP report schema/target/input/output hashes checked before
   compiling. Its report remains `generated`, not a compatibility assertion.
2. `compiled`: candidate linked against the same pinned GoHP runtime with Go
   module/toolchain network downloads disabled.
3. `differential_pass`: candidate and native PHP both exited 0, emitted empty
   stderr, and emitted exactly:

```text
true=1
false=
zero=0
null=
empty-string=
```

All lines end in a newline. This also matches the owned expected observation,
not just the other runtime's output.

Each invocation used a non-root, network-disabled, read-only-root, capability-
dropped, resource-limited container. Candidate files/caches were confined to the
Go test temporary directory. Containers used automatic removal; the test cleaned
temporary source, generated files and binaries. The pinned GoHP clone remains
in its explicitly named temporary folder for local reproduction. No source was
pushed to GoHP, no original WordPress file was changed, and no DB/host port/site
or production credential was involved.

## What this does not prove

Five dependency-free constant-return helpers are not WordPress core, a full
`functions.php` conversion, callbacks, plugin execution or a PHP runtime bridge.
No CMS request, mutation, theme, database or browser behavior is exercised.
Echo output alone cannot distinguish the return types of false/null/empty string;
return-type equivalence and warning/exception semantics remain unverified.
The original hook engine still needs unsupported GoHP language features.

Generated artifacts are disposable evidence, not a release or activation. A
test without `GOSSIP_GOHP_SOURCE` explicitly skips and must not be called PASS.
