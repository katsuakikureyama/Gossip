# Full-parser admission after string switch — 2026-10-29

The complete original parser now passes the switch syntax at100 and stops at
the chained-property write at207. Generation remains blocked; full-parser
execution/comparison and application acceptance have not run.

Original WordPress `54106f6c545fc8b881f6b39b7963002bece6c28c`, whole input
`wp-includes/class-wp-block-parser.php`, SHA-256
`8563621090de5e307431b1d29b78a9f75fd9ee238a0137b563a7db7b2e1f8d3b`.
GoHP [PR121](https://github.com/katsuakikureyama/GoHP/pull/121) clean source
`3ecd1f4df5bfb1cbc72974ce82bf316885d3a676`, tree
`4662a190f6c1d657cf4b12c754a5eb2f76ad2d4e`, based on main81feba6.
Host Go1.23.0 Darwin ARM64, CGO0, offline cached modules/toolchain.

Personal local paths replaced by placeholders:

```sh
cd GOHP_ROOT
GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off GOWORK=off CGO_ENABLED=0 \
GOCACHE=CACHE go run ./cmd/gohp \
  -in GOSSIP_ROOT/wp-includes/class-wp-block-parser.php \
  -php-version 8.3 -experimental-pcre2
```

One invocation, exit1:

```text
PHP002 line 207: property assignment requires one declared property and = (at "->")
exit status 1
```

The second line is `go run`'s wrapper. No output/report destination was requested.
The full original file was used without extraction, concatenation or rewriting.
The rejected statement writes through `$stack_top->block->innerHTML`; support
for that location and its compound write belongs to P2. The probe stops before
complete semantic validation, so it does not certify all later P1 constructs.

At the same clean source, `go test -count=1 -run
'^Test(StringSwitch|SwitchString)' ./internal/codegen ./runtime ./tests` passed
six checks once (packages0.700s/0.911s/5.982s). Independent PHP8.3.36/8.4.27
native evidence and regression results are retained by PR121 CI. The owned
fixture covers first-match duplicates, default placement, fallthrough, selector
once, nested switches/loops, conditional returns and shared strtok state.
Numeric/nonstring comparison and switch-target continue remain excluded.

This updates the [previous switch rejection](parser-list-probe-2026-10-29.md).
Existing native parser references, successful consumers, original files and
all pins are unchanged and were not rerun. This Gossip changeset is documentation
only; next P2 nested property/stack mutation, then remaining P3–P5 dependencies.
Existing local inspect capability/WebMCP exclusion and route/auth/deploy stay unchanged.
