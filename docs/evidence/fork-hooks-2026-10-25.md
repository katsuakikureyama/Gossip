# Fork-source hook differential — 2026-10-25 JST

## Targets and provenance

- Gossip implementation/harness: `a167e6c8c56294f2d702e7c25342c614bde364dd`.
  The exact files tested were committed without any source changes afterward.
- WordPress source: parent commit
  `54106f6c545fc8b881f6b39b7963002bece6c28c`, version `7.2-beta1-64693`.
  This is a development snapshot, not a production version recommendation.
- The copied native SDK and fixture originate in `Gossip-prototype` commit
  `1b9c05b0557bcca3f0a4f5eebf046cbef67d18e2`. Their byte equality was checked.
- Fixture SHA-256:
  `b77b22bb03eb7639387e8ad0e70701db707affad19ba0ff43b647ecefa44ebed`.
- All five original-file hashes are recorded in
  [`gossip/wordpress-source.json`](../../gossip/wordpress-source.json).

## Execution and outcomes

Host: macOS arm64; Go `go1.22.1 darwin/arm64`. PHP runner:
`wordpress@sha256:04be81de853887fca707742785b32bf93fb4a692246dd1bc90b3790105071a9a`,
reporting PHP `8.3.13`. This already-cached immutable image supplies **PHP only**.
The WordPress 6.6.2 installation in the image is not loaded or compared.

From `/Users/kureyamayuri/Desktop/Gossip-wordpress/gossip`:

```sh
GOSSIP_PHP_RUNNER_IMAGE=wordpress@sha256:04be81de853887fca707742785b32bf93fb4a692246dd1bc90b3790105071a9a \
  go test -race -count=1 -v ./...
```

**PASS** in one run for this source/configuration:

| Check | Result |
| --- | --- |
| Native registry priority/tie ordering and failed registration | PASS |
| Original source/version/license hash check (5 files) | PASS |
| Empty hook | PASS; both calls returned `seed`, empty trace |
| Ascending priority | PASS; both traces `[1 2 0]`, `C(B(A(seed)))` |
| Equal-priority registration order | PASS; both traces `[0 1 2]`, `third(second(first(seed)))` |
| Priority 0/999 boundaries and ties | PASS; both traces `[1 2 0]`, `last(zero-second(zero-first()))` |
| Unicode string propagation | PASS; both traces `[1 0]`, `「前こんにちは & <content>後」` |

The harness first reads and hashes the original files, then stages those exact
bytes (not later reads from a possibly changed checkout). It loads
`/wordpress/wp-includes/plugin.php`, which requires the original `WP_Hook` and
filter sentinel classes. PHP reports the same source hashes and declared
WordPress version; Go checks both before accepting callback results.

The Docker invocation uses `--rm --pull=never --network=none --read-only`,
`--cap-drop=ALL`, `--security-opt=no-new-privileges`, UID `65532:65532`,
128 MiB memory, 32 PIDs and one CPU. Only staged public upstream files and the
owned fixture/harness directory are mounted read-only; no `.git`, credentials,
database, Docker socket or production data is mounted. The exited container
was automatically removed, and Go cleaned its temporary source directory.

The source-preservation check, run before committing additions, was:

```sh
git diff --exit-code 54106f6c545fc8b881f6b39b7963002bece6c28c -- .
```

It passed with no changes to tracked upstream files. The implementation commit
diff contains only added `gossip/` paths; the follow-up documentation adds only
README/docs. No build or behavioral test was repeated after PASS.

## Explicit exclusions

This comparison demonstrates five owned single-string filter cases in the
native SDK, not a full WordPress hook implementation. It does not cover dynamic
registration/removal, callback identity, recursion, references, multiple args,
negative priorities, sentinel behavior or the full core filter stack. The
sentinel class is loaded unchanged but its special semantics are not tested.

There is no PHP-to-Go translation, GoHP invocation, WP/native request bridge,
full WordPress startup, database, theme rendering, browser run or production
deployment in this evidence. An ordinary `go test ./...` without the environment
variable checks the source pin and SDK but explicitly skips the Docker test.
