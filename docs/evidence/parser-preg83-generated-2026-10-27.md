# Generated GoHP regex consumer — evidence, 2026-10-27

Result: **23/23 selected cases PASS**, zero mismatches. One of the original
24 cases is explicitly excluded: `backtrack_limit_one` uses limit 1, while
GoHP's admitted profile fixes limit 100000. This is a generated **owned driver**
with unchanged WordPress pattern bytes, not a translated original method.

## Immutable identities

- Gossip clean code: `99cb3f86f6c5ad5e55836b7ca7511d2d539a12ab`,
  tree `085890abb070dc363a41ec5b3b3680db54977aa5`.
- GoHP compiler/runtime: `d7650944802a33cfc2dab0006bb410374d828245` (#114), exported from Git at that
  exact revision, independent of local checkout head or uncommitted changes.
  Archive SHA256 `d97e025f4835c70ec95a3e4526000010cd2b70d32c4158ec5f36ed48be918395`.
- Original WordPress: `54106f6c545fc8b881f6b39b7963002bece6c28c`.
  Full parser/Block/Frame/license/version file identities are in the
  [machine-readable result](parser-preg83-generated-2026-10-27.json).
- Retained native reference SHA256 `5cbf560247afcd2c5f1462cab2f95f16bb822ac2d4d1756be16221cffdc6dc60`.
  The existing PHP 8.3.36 / PCRE 10.42 record is reused, **not rerun**.
- Reference image:
  `php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd`.
  Native C locale, JIT0/depth10000/match100000 for the 23 selected cases.
- Candidate: Go1.23.0 Darwin ARM64, Apple clang17, macOS SDK15.5, existing
  PCRE2 10.47 static archive. Compiler/header/archive/Go-tool/SDK-settings hashes
  and sanitized commands are retained in the machine-readable result. No install,
  network fetch, pkg-config fallback or binary release occurred.

## Commands and observed phases

With the cached Go1.23.0 first on PATH, an existing writable build cache,
`GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOWORK=off CGO_ENABLED=0`:

```sh
cd gossip
go test -v -count=1 \
  -run '^(TestParserPreg83GeneratedContract|TestWordPressSourceBaseline)$' \
  ./tests/gohp ./tests/compat

# Additionally set CC to the pinned clang executable, SDKROOT to macOS15.5,
# and these two absolute local paths. The harness fixes the child CGO flags.
GOSSIP_PREG83_GOHP_SOURCE=/absolute/local/GoHP \
GOSSIP_PCRE2_PREFIX=/absolute/pinned/pcre2-prefix \
go test -v -count=1 -run '^TestParserPreg83GeneratedConsumer$' ./tests/gohp

go vet ./tests/gohp
```

The contract and original-source baseline checks PASS; the actual generated
consumer PASS in 6.41 seconds (package 8.910s); vet exit0 with no output. One
execution of each, no failed attempts and no repeated successful tests. Gofmt
and diff whitespace checks pass. Original tree comparison against the WordPress
baseline (excluding only Gossip overlays: gossip/, docs/, README.md and
.gitignore) is empty. No earlier fixtures or their pins changed.

The pinned compiler is built with CGO0, the owned PHP source is compiled using
the explicit PHP8.3 regex profile, and report/source/generated hashes are checked.
Actual emitted Go plus the observer are then built with CGO1 and `gohp_pcre2`
against the **same** archived runtime. The binary executes with empty PATH:
exit0, stderr0. The native linker emits only a duplicate-static-archive warning;
its sanitized text and original stderr hash are retained, not conflated with
runtime diagnostics. No interpreter or handwritten matcher supplies its result.

All 23 returned observations structurally equal the retained PHP projection:
17 matches and 6 non-matches, with ordered typed captures, raw bytes, byte
positions, named presence, before/after output values and error codes. PHP error
message strings, the limit1 failure and full next_token/parser state are excluded.
The observer receives selected inputs only, not expected results. The driver
receives a source-escaped literal with the exact original runtime pattern;
no regex rewrite or compiler WordPress whitelist is used.

## Selected SHA256 records

| Artifact | SHA256 |
| --- | --- |
| [Full generated/native comparison result](parser-preg83-generated-2026-10-27.json) | `e80444887226d5bbd7a016b44e9eb240184636bb4c2c8c39045fa9536e7a1996` |
| Complete 24-case corpus | `ac92d6eb7cbf2fef591d73749e13c8110fe00934cb96a7e825dc7e5727ff83ef` |
| Selected 23 inputs without oracle expectations | `b57af0c58a4d485e00953bac5e50bac45992c59d506e36271807a4e1dd7ca64b` |
| Original runtime pattern | `3394500101eae41893b0978b3003f891dfafab54e7d7ebcbf820422753d8446e` |
| Owned PHP driver after literal substitution | `3ec461f297cad266bfd84dc360dd082b11afad7dcb77e0fd4083b4040c740476` |
| Generated Go | `85cae3924a0a815c8c62a85974130ce468782e77f7f5eeee79edc1a327769359` |
| Observer | `4924fbfe296e8923ec5a157b5aa3f88e77aaf242023aa326d63f4a70b22d32d1` |
| Harness | `9b9213303d1d164f9d8c5cc5034880257fa87e59ff38e011f93c78f9126f992d` |
| Candidate executable | `617085feaf6ee0b0eba033f54bf123ee114483d1d2e43ef0bac56d27171bcf3d` |
| Candidate stdout | `b273a16543a12597d2ab020e0d0d3b09847652c70a4d6e0e53603106b22e4a91` |
| Local raw consumer log | `bdb0dae548e96b81dc92d30cf2defc99414af6fa1b5b5a4f81fe54c43dede927` |
| Local raw contract log | `3761b0cc44a555f57153fe6bfc469a95a397f5a1b9ec349f52cfb9de80900557` |

The full result retains both actual and expected observations plus provenance.
The raw logs remain local; paths containing usernames are not published. The
compiler report hash is recorded but the path-bearing report is not published.

## Remaining boundary

This is the first Gossip generated-code connection to the bounded regex
profile, not 23 newly translated WordPress functions. The full original
`next_token()`, list destructuring, JSON decoding, general property/output
locations, complete parser stack/control flow, PHP8.4 regex, other backends,
enabled Linux CI/packaging, performance and distribution remain unmeasured.
GoHP's earlier generic failure/Execution tests are separate evidence; the excluded
consumer limit1 case is not promoted by them. Existing original-method evidence
and older compiler pins are preserved.

UI/WebMCP/MCP capability is explicitly excluded: local developer fixture only,
no route/authentication change, service operation or deployment.
