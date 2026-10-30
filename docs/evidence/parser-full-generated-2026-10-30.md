# Complete generated Parser retained-reference execution — 2026-10-30 (JST)

## Status and scope

This is one opt-in, retained-reference execution of the complete generated
original `MContext_parse` entry.  It is the first actual invocation of the
generated complete Parser/Block/Frame closure; it is not a claim that Gossip,
the full PHP language, or the PHP 8.3 application gate is complete.

At clean Gossip commit
`29b9bd2a06d7100e00d5395beef6cd7bffb3501a` (tree
`e35090afbee70a0748bb7dc1b9d519d62b1c4b49`),
`TestOriginalWordPressFullParserGenerated` ran once in `retained` mode and
passed in 4.88 s (package: 5.122 s).  The matching machine-readable record is
[parser-full-generated-2026-10-30.json](parser-full-generated-2026-10-30.json)
(SHA-256 `d9acb80b2e9a42feec3016a9b5c52f162ee7d8092b586e113f445fbe315b9175`).

The test calls the complete generated entry 24 times: for each declared case,
first with a fresh parser and then with the parser reused across all 12 cases.
All calls share one GoHP `Execution`.  Before each call the harness installs
the JSON baseline produced by associative decoding of `null`; this is part of
the observed contract, not an alternative implementation of PHP parsing.

## Fixed inputs and identities

| Item | Identity / result |
| --- | --- |
| Current GoHP source pin | main `b2d1482268a7710c515c07bf919af3d40b09d17c`; Git archive SHA-256 `d20e5572b8aa5f5f57751623fe1a4f5631fbca8c8960ba18c8caa4be5b990653` |
| Reused checked generation | GoHP `35524a636073f4abbc8bb69b9e36263e71ac62d8`; generated Go SHA-256 `55c66c931ed9da4be7a0a263129242b169a68127e83bb526fbc1f6d247f402dd`; report SHA-256 `bbf61c482f333b6509cff72b8e27cd2c7833238bf81a4148dda3e4f08104a1df` |
| Production-source comparison | `35524a6…` to `b2d148…` is empty for `cmd`, `internal`, `runtime`, `go.mod`, and `go.sum`; no source compiler/runtime change is being attributed to this run |
| Closure manifest | SHA-256 `89f897c64146705559eae41bb50d49aff57abad0a11ee02d16d27fe78d64adaa` |
| Retained full native reference | PHP 8.3.36 record SHA-256 `1b34c7af6d7baeda2997e033c93727147bad4d36e6bfed676e183e6166ea43a7` |
| Retained reference manifest | SHA-256 `b3a80436c9498355305d654fe6ed127591c6d83f5ea76b70a6c595e405b6f5e5` |
| Canonical typed cases | Native and Go SHA-256 `3780c6a5ebd4eefda51bc548b2e95f585687c39f274eb6d3c8a622792607ca28` |
| Observer source | SHA-256 `5da30cc715689e79bf2354e29981a7f8d52f12dd3d0fa03527300a4324bee774` |

The prior CGO0 generation/build was not re-run.  This run explicitly built and
ran the candidate with CGO enabled and the existing static PCRE 10.47_1
archive on Darwin/arm64: Go 1.23.0, Apple clang 17, SDK 15.5.  The archive,
header, clang, and SDK-settings SHA-256 values are respectively
`ce40246eba1dfdd99e43142cfc33264610d7c965fab3062f386c68a6636a4a54`,
`058462dc030e845ee627e19ac36347f561f5cc44d89eaf1c0010f2b363ab68d5`,
`4c458256bcdf913774de1bb4a37768244d3b41f32ad55afb2dfec3432f46f4fe`, and
`58a133735f0a55a624a1703067059f6e78925e51725ec6e1f966072e142c9c42`.
The temporary executable (SHA-256
`82e9a08839151771a36653bc67c87c26efd5f3476f0f0a087cf31a3436683fb3`) was
removed automatically after observation.  No dependency installation, image
pull, daemon startup, persistent binary, or dedicated cache was created.

## Exact execution shape

From the Gossip `gossip` module, with the cached Go1.23.0 bin also on PATH
(personal absolute paths replaced by the named verified inputs):

```sh
GOSSIP_GOHP_FULL_PARSER_MODE=retained \
GOSSIP_GOHP_FULL_PARSER_SOURCE="$GOHP_ROOT" \
GOSSIP_PCRE2_PREFIX="$VERIFIED_PCRE2_PREFIX" \
GOSSIP_GOHP_FULL_PARSER_GENERATED_DIR="$PRIOR_GENERATION_DIR" \
GOSSIP_GOHP_FULL_PARSER_EVIDENCE_OUT="$FRESH/attempt1.json" \
CC="$VERIFIED_CLANG" SDKROOT="$VERIFIED_MACOSX_SDK" \
GOCACHE="$SHARED_EXISTING_CACHE" CGO_ENABLED=0 GOTOOLCHAIN=local \
GOPROXY=off GOSUMDB=off GOWORK=off \
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressFullParserGenerated$'
```

The CGO0 setting belongs to the test harness; it explicitly builds the candidate
with CGO1, `gohp_pcre2`, the hash-checked static archive and the measured compiler/
SDK. The evidence output is exclusive/fresh, never an overwrite. Prior generated
source/report are hash-checked and reused; omitting the reuse directory permits
new generation for a later explicitly measured run, not an implicit repeat here.

## What passed

For every declared case in both fresh and reused-parser passes, the observer
matched the retained PHP reference for typed output, parser/global state, key
insertion order, raw string bytes, error status, caller-visible state, and the
reused-parser reset behavior.  This includes the retained reference's expected
result for the malformed JSON-attributes case.

The actual Go observer exited 0 with empty stderr.  Its stdout was 59,260 bytes
with SHA-256 `e0e5a47f2e36b7cbbcf293e16b3ced0417cd3eb72b8a63eba20ea85ca51deebb`.
The native record's complete raw stdout is 60,065 bytes because its metadata is
different.  Raw full stdout is therefore deliberately **not** compared; the
canonical typed `cases` value above is the equality contract.

The enabled build emitted a duplicate-static-archive linker warning.  It is a
retained build warning, not a runtime diagnostic: this execution's stderr is
empty.  No failure was observed within this narrowly defined execution scope.

## What remains unverified

The PHP container was not rerun: the Docker daemon was unavailable.  `fresh`
mode is implemented but unexecuted, and requires the exact existing PHP image.
Consequently this record compares Go only with the retained fixed PHP 8.3.36
reference; it is not a fresh native-PHP-versus-Go differential result.  Its
native and Go runtime profiles also differ (the retained reference uses its
fixed container profile, while this candidate uses Darwin/arm64 PCRE 10.47_1),
so this result must not be recast as byte-for-byte process-output equality.

`G83-PARSER` remains open pending a fresh native capture/comparison and the
required corpus-boundary review, including API/input negative cases.  The
bootstrap, data, HTTP, and bounded-site gates have not run.  The 12 cases are a
bounded runtime observation of the P2/P3/P4/P5 paths exercised by the corpus;
they do not demonstrate every branch of those slices.

This harness and observer are test-only. The stable local-operation exclusion
is `gossip.tests.original-parser-compare`: it covers retained/fresh comparison
and fresh-path evidence only, and has no browser, server, authentication, route,
or WebMCP capability. It does not add a PHP extraction, source rewrite, fallback
parser, builtin/compiler/runtime source change, deployment, or production target.
The existing compiler CLI remains unchanged and this test creates no WebMCP
exposure. Without the explicit mode, the test skips and produces no execution
proof; retained mode does not start Docker. Fresh mode also requires the exact
fixed `GOSSIP_PHP_RUNNER_IMAGE`; it was not executed here and does not authorize
daemon startup or image installation.
