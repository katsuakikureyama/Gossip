# Typed original-helper observations — 2026-10-25 JST

## Target and command

- Tested Gossip commit: `b0f0ce82d0532e1ff0258478b95a158c29b84efc`.
- Original WordPress: `54106f6c545fc8b881f6b39b7963002bece6c28c`.
- GoHP compiler/runtime: `9c7389d35032d623a429198a69cdfa5c11571e8e`.
- Host driver: Go 1.22.1 darwin/arm64; builder: Go 1.25.14 linux/arm64.
- PHP runner: PHP 8.3.13, using only this fork's extracted original functions,
  not image-bundled WordPress. Immutable image/source/artifact hashes are in
  the [machine-readable observation](typed-helpers-2026-10-25.json).

From `/Users/kureyamayuri/Desktop/Gossip-wordpress/gossip`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gossip-gohp.y33KQk/GoHP \
  go test -count=1 -v ./tests/gohp
```

PASS on the first run of this change: test 18.54 seconds, package 18.892 seconds.
No test was repeated after PASS. Only documentation/evidence changed afterward.
Unchanged SDK/hook tests were not rerun; their earlier results remain separately
recorded. Build tmpfs was 256 MiB throughout this run.

## Observed results

| Original function | PHP and generated Go observed return |
| --- | --- |
| `__return_true()` | bool `true` |
| `__return_false()` | bool `false` |
| `__return_zero()` | int `0` |
| `__return_null()` | null |
| `__return_empty_string()` | string `""` |

The test first verified generated source/report hashes and compiled the generated
program against the same pinned runtime. Echo stdout still matched the expected
five lines, with empty stderr and exit 0 in both runtimes. Separate observer
programs then called the actual functions; JSON records matched each other and
the owned expected fixture, again with empty stderr and exit 0. The Go observer
does not replace or edit generated functions; it exits before their echo main,
which is exercised independently. See [execution contract](../translation-units.md).

The unchanged original `__return_empty_array()` (lines 7222–7224) was exercised
separately. Native PHP returned `{"type":"array","value":[]}`, empty stderr,
exit 0. GoHP exited 1 with exactly:

```text
PHP004 line 3: unsupported function array
```

The compiler emitted no stdout, generated Go source or report. This expected
rejection passes the boundary test, but the function remains unsupported. The
original `array()` expression was not rewritten to the supported `[]` syntax.

## Isolation and limits

All executions used non-root, network-disabled, read-only-root containers with
resource limits, capability dropping and no-new-privileges. Only test-owned
temporary outputs were writable; containers and generated files were cleaned
automatically. No service, DB, host port, browser session or production resource
was involved. No original WordPress file or GoHP source was changed.

Typed results cover five constant returns only. Observer JSON/type reflection is
test instrumentation, not proof that GoHP translates those PHP builtins or
dynamic calls. Generated symbol naming is private to the pinned compiler.
Warnings/exceptions, nonconstant inputs, HTTP headers, hooks, themes, plugins,
whole-file/core conversion and a deployable Go WordPress server remain unverified.
No new user-facing operation is added; the local-test capability exclusion in
the translation contract continues to apply. No devboard card is linked or
marked Done. Draft review does not authorize deployment or default-branch merge.
