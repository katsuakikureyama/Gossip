# Regexp2 comparison execution — 2026-10-27 JST

Result: **22 equal, 1 different, 1 not-comparable**. Compatibility not established.
[Contract](../parser-regexp2-comparison.md),
[full observations](parser-regexp2-comparison-2026-10-27.json).

- Clean code commit `073e1218292c160a02f6e06ad2a4a83f2ea547e7`;
  tree `726730f558610e01e99badcbb226b0e069a03379`.
- Base Gossip master `d86a0dd52fb775581547be6cffc837590fe6e794` (#47),
  branch `test/parser-regexp2-comparison`; later additions are only docs/evidence.
- Regexp2 `a6ef7b6f88fff76c0c7d4f5fe832535de8cfda11`, source tree
  `ec945625899ff6c882fa546ae8193e358b7ad1b5`;
  git archive SHA `c1cbee3f2ce552bec7c649e6944dc49e4748d04622d745856a95c460687dd8da`.
- Parent harness Go1.22.1 darwin/arm64; isolated probe Go1.25.14 darwin/arm64,
  CGO_ENABLED=0, GOPROXY=off, GOWORK=off, GOSUMDB=off, GOTOOLCHAIN=local.
  Source/archive and toolchain identities fixed; no automatic dependency install.
- Original PHP8.3.36/PCRE10.42 reference SHA
  `5cbf560247afcd2c5f1462cab2f95f16bb822ac2d4d1756be16221cffdc6dc60` reused.

From `GOSSIP_ROOT/gossip` with installed source/toolchain environment from the contract:

```sh
go test -count=1 -v ./tests/gohp -run '^TestParserRegexp2Comparison$'
go vet ./tests/gohp
```

Observation harness PASS once (6.91s / package7.147s), vet PASS, format/diff checks
PASS. There was no failed native attempt or successful rerun. All 24 probe
processes exit0/stderr0. A regex API error is structured output, not a process
failure; test PASS means complete observations, **not** backend equivalence.
The earlier PHP/PCRE2 runs were not repeated. No committed Actions workflow or
CI success is claimed for this fixture.

Probe executable SHA
`86f1c9b873a81b63192f43b66e5185759ba814a1570b7113427ce2623ae313cd`.
The JSON records full source/harness/archive/license/toolchain hashes and all raw
groups/byte ranges/rune ranges alongside PHP-shaped output. The binary remains
temporary/not published; its hash does not promise cross-platform reproducibility.

Ordinary 22 cases agree, including FF attributes, NUL/invalid prefixes, ordered
captures, multibyte offsets and trailing-group omission. `offset_inside_utf8_prefix`
instead returns a rune-alignment API error (PHP returns a match starting at byte3).
`backtrack_limit_one` returns a candidate match, but has no equivalent PCRE
match-count setting: not-comparable, not error-equivalent or ordinary PASS.

No GoHP/cgo/runtime integration or dependency adoption. No full parser, JSON,
PHP8.4, arbitrary binary regex, extra whitespace/flag/profile or performance
claim. Original WordPress bytes/history, old module/fixture/compiler pins and
earlier evidence unchanged. The direct string-API profile is not acceptable as
a drop-in implementation of the retained original-parser contract.
