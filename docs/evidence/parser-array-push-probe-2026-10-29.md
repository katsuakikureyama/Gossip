# Parser array_push admission probe — 2026-10-29 (UTC record)

This evidence records one compiler-admission probe for the fixed Gossip PHP8.3
target. The record date is UTC; the surrounding Gossip current-state document
was updated on 2026-10-30 JST. It does not record a generated parser, a native
PHP/Go comparison, or application acceptance.

## Candidate boundary

[GoHP PR136](https://github.com/katsuakikureyama/GoHP/pull/136) admits only
`array_push($target, value, ...)` used as a standalone call or as the complete
RHS of a bare-local result assignment. The target is an existing bare local or a
direct named public property of a defined bare local, including `$this->stack`.
A property owner is captured before payload evaluation; payloads are captured
once left-to-right; the current target array is checked after those payloads;
successful appends publish in order; and the int count result is bound last.
This is an additive IR0.1 statement, not a general reference or ordinary
embedded-call model.

Known non-array and undefined targets reject statically; unknown targets that
are actually non-array reject at runtime. Generic embedded calls, general
references, indexed/nested/dynamic/effectful targets, non-bare result targets,
builtin redeclaration, `array_pop`, and general stack semantics remain outside
this increment.

## Fixed inputs

| Item | Value |
| --- | --- |
| Initial clean compiler code | `2d01aa04b09d926d48299baa4a2210065af435e8` |
| Initial clean compiler tree | `eff55f7dcdaba035382c04adb7a909f7335e2280` |
| Tested candidate | `b1344c8360a2f829ae4f7f2f9fff227b010cf514`, tree `8a7fcb8aed19f680c8a50f164c48eb39bc0642ab` |
| Merged main | `e61870a9d7fb8da8d9d03933a8a29b9351d8da28`, same tested tree |
| Fixed WordPress source | `54106f6c545fc8b881f6b39b7963002bece6c28c` (`7.2-beta1-64693`) |
| Closure | Complete unchanged `WP_Block_Parser`, `WP_Block_Parser_Block`, and `WP_Block_Parser_Frame`, with existing notice/provenance files |
| Manifest | [`parser-source-closure.json`](../../gossip/tests/gohp/parser-source-closure.json), SHA256 `89f897c64146705559eae41bb50d49aff57abad0a11ee02d16d27fe78d64adaa` |
| Probe toolchain | Cached Go1.23.0, Darwin ARM64, CGO0; `GOTOOLCHAIN=local`, `GOPROXY=off`, `GOSUMDB=off`, `GOWORK=off`; existing shared cache |

No source extraction, concatenation, handwritten parser or interpreter fallback
was used. The manifest, original source and consumer pin were not modified.

## One probe

```sh
go run ./cmd/gohp \
  -manifest GOSSIP_ROOT/gossip/tests/gohp/parser-source-closure.json \
  -source-root GOSSIP_ROOT \
  -entry wp-includes/class-wp-block-parser.php \
  -php-version 8.3 \
  -experimental-pcre2 \
  -out FRESH/parser.go \
  -report FRESH/report.json
```

The command was invoked once at the initial clean compiler revision. It exited
1 with stdout0 bytes and stderr101 bytes, SHA256
`246e5052200024fe11f7578fb32e1d9d2bd9215f7449d6c06bb020a0c0be9b00`:

```text
wp-includes/class-wp-block-parser.php: PHP004 line 205: unsupported function array_pop
exit status 1
```

The prior codegen rejection at original line165,
`array_push($this->stack, new WP_Block_Parser_Frame(...))`, was passed. No Go
source, report or staging artifact was produced, and the empty temporary output
directory was removed. The next measured blocker is original
`array_pop($this->stack)` at line205.

## Candidate verification and limits

Push40896987607 and PR40897000188 attempt1 succeeded on all four target jobs
at the tested candidate tree, using Go1.23.12 Linux amd64/CGO0 and PHP8.3.36 /
8.4.27. All nine new checks passed on both targets. The owned fixture's native
and Go stdout are equal at272 bytes, SHA256
`c0985590c4de999b83832f8421b22902f7a63138646dd96b552fe69732104800`, with
exit0 and quiet stderr0. Fixture SHA256 is
`5f6bd7cd6e5613a5aaf164b2c4191067ea5efa1e34a1a50d8d4cdd8df7ebc663`.

[Exact-main run40897806328](https://github.com/katsuakikureyama/GoHP/actions/runs/40897806328)
attempt1 also completed successfully at merged main
`e61870a9d7fb8da8d9d03933a8a29b9351d8da28`, the same tested tree. Its PHP8.3
job123531572125 and PHP8.4 job123531571699 both succeeded. No successful
artifact was downloaded again and no check was manually repeated.

This result is still only an unchanged-source compiler-admission measurement.
It does not prove original Parser/Block/Frame generation, binary build, parser
runtime, the declared 12 native reference cases/24 original calls,
differential output, or G83-PARSER acceptance. `array_pop`, remaining
stack/snapshot behavior, tokenizer/PCRE integration, JSON attributes and all
later application gates remain open. General exception, stderr and side-effect
equivalence are unmeasured.

Gossip has no configured CI. There is no service, resource, auth, route,
deployment, browser operation or WebMCP publication: local
`gohp.cli.generate` remains explicitly excluded.
