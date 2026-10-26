# Parser-token native reference execution — 2026-10-27 JST

Result: **24/24 predeclared native expectations PASS**. This is not a Go/PHP
differential result. [Contract](../parser-token-reference.md),
[full machine-readable evidence](parser-token-reference-2026-10-27.json).

## Exact execution snapshot

- Clean source/test commit: `29e6d2637e68cbc3a8d5764ba825e6e6600ddf74`.
- Tree: `7e8e33fe8d3c2870b3c9dc5b5c7400803285c3ac`.
- Base: Gossip master `f245f2b64a40dc10ed682788424650d256ca8d8f` (#45).
- Branch: `test/parser-token-reference`; subsequent additions are documentation
  and this retained execution evidence, not code/corpus changes.
- Host Go: go1.22.1 darwin/arm64; GOPROXY=off, GOTOOLCHAIN=local; local GOCACHE
  selected (personal/cache path omitted). No compiler/backend invocation.
- PHP image: `php@sha256:98d555ff97b3865140ad31032508c3ca089fa97a8944b14ca9dc03ea05d788cd`.
  Cached arm64 image ID: `sha256:f39899ba9e42005403f535ec925a6e46a1c70707085b689e41a48219798521d7`.
- Observed PHP 8.3.36, PCRE `10.42 2022-12-12`, PHP_INT_SIZE=8, LC_CTYPE=C.
  PCRE JIT supported but explicitly disabled; memory_limit=128M,
  recursion_limit=10000, backtrack_limit=100000 (limit case=1).

## Commands and outcomes

From `GOSSIP_ROOT/gossip`, with the image/environment above:

```sh
go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressParserTokenReference$'
go vet ./tests/gohp
```

Native test PASS once: test 1.13s, package 1.407s, command exit 0. Vet PASS once.
Formatting and `git diff --check` PASS. Native process exit 0, stderr 0 bytes,
stdout 76315 bytes, SHA-256
`d449a36d5aa897cb421ec659ac7132c80881c9b454ca210ff4ab80f810a640eb`.
Full source/fixture/observer/harness/shared-policy hashes and typed observations
are in the JSON. The retained evidence is extracted from that successful run;
no second native execution or candidate engine is implied. No committed GitHub
Actions workflow supplies an additional CI result for this change.

## Observed scope

24 separate direct preg_match calls and 24 original next_token calls: 17 matches
(integer 1), 6 non-matches (integer 0), 1 failure (Boolean false, PCRE error 2).
All predeclared match/error/token expectations and unchanged caller/parser-state
checks passed. Full captures, presence/order/raw bytes/offsets, before/after
matches, effective per-case profile and JSON errors are retained observations.
They have not yet been compared with another engine or Go implementation.

- Invalid JSON attributes: regex match, null attributes, JSON error 4.
- Invalid UTF-8 attributes: regex match preserving byte FF, null attributes,
  JSON error 5. Invalid bytes are not rejected merely because they are non-UTF-8.
- Multibyte prefix and interior-byte offset: returned comment starts at byte 3;
  NUL/invalid-byte prefix cases start at byte 1.
- Backtrack limit 1: false/error 2, matches changes from null to empty array;
  original method returns no-more-tokens, as for a normal non-match.
- Output/stack sentinel arrays and input/offset remain unchanged for every case.

Original WordPress Parser/Block/Frame, license and version are hash-verified;
no original source, old fixture or GoHP pin is changed. This corpus does not
execute parse(), generated Go or an alternate regex backend. PHP 8.4, out-of-range
offsets, JIT-on and additional error profiles are unmeasured, not assumed equal.
The next gate is a separately pinned backend comparison under this exact profile.
