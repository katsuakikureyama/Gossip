# Original string-only stripslashes helper — 2026-10-25 JST

## Outcome and exact revisions

**PASS:** the unchanged `stripslashes_from_strings_only()` was translated and
built, then compared against native PHP 8.3.36 for 21 authored inputs. Actual
return types and values matched PHP and explicit expectations, with empty
stderr and exit 0. Non-string values, including nested lists containing escaped
strings, returned unchanged. This is not recursive WordPress unescaping.

- Gossip test commit `f4e65138298ae6763670d457c9c2f74615b8103a`, tree
  `e4e0c78a9cc0403f97c21d56fa82a990e7e13886`.
- GoHP compiler/runtime `0e49cc995913dfede0332566dd4dc11c5df16ff8`, tree
  `7aa14fc6a9f384298bdf94705dc80c236b85d37b`.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`, unchanged
  formatting.php lines 2934–2936, whole-file SHA
  `d44aa1329df2fc38bf2670ad6de570e6b9eb65db9ed5507f26c683e569b2919e`.
- Actual PHP 8.3.36; approved immutable image/same-process guard unchanged.
  Host Go1.22.1 darwin/arm64; builder Go1.25.14 linux/arm64.

The [machine record](stripslashes-helper-2026-10-25.json) contains every source,
excerpt, input, observer, compiler archive/binary, report, generated program and
executable hash plus all observations. Only committed compiler files were
exported. Report hashes were checked before building against that same runtime.

## Gossip execution

From the nested `gossip/` module, once at the test commit above:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressStripslashesFromStringsOnlyHelper$'
go vet ./tests/gohp
```

Initial PASS **7.87s**, package **8.078s**; vet PASS. No failures, retries or
implementation changes after PASS. Subsequent changes are documentation only.
Existing container isolation and native admission checks were preserved.

The corpus has 12 strings, five null/bool/integer values and four lists.
It covers quote escapes, paired/trailing/unknown backslashes, escaped ASCII
zero, actual NUL, Japanese, emoji, mixed escapes and nonrecursive passthrough.
Go transports integer values through `UseNumber`/`Int64`, validates actual
runtime types, and converts only arrays with consecutive int64 keys into JSON
lists. This corpus uses modest integers; it is not an arbitrary large-number
JSON comparison claim. Observers dispatch only, without copying the original
helper or calculating expected unescaping results.

## Compiler extension and local verification

[GoHP Draft #23](https://github.com/katsuakikureyama/GoHP/pull/23), stacked on
#21, adds full three-arm ternary plus exactly one actual-string `stripslashes`.
The condition evaluates once and only the chosen arm evaluates at runtime;
both arms are statically validated. PHP8 non-associativity and precedence are
retained; malformed IR arity, unsupported syntax, object conditions, invalid
static arguments and builtin redeclaration fail closed. Mixed branch types
remain unknown, while equal types are preserved. Dynamic non-string builtin
arguments are guarded at runtime. Catalog: **35 entries / 89 unique checks**.

At the pinned compiler commit, cached Go1.23.0 darwin/arm64:

```sh
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go test -count=1 ./runtime ./internal/codegen ./internal/compat
GOCACHE=/tmp/gohp-go-cache GOPROXY=off GOTOOLCHAIN=go1.23.0 go vet ./...
```

Initial PASS runtime0.419s / codegen0.600s / compat0.256s; vet PASS.
The native fixture observes precedence, nesting, condition/selected-arm echo
effects and guarded non-string values. NUL conversion is compared natively;
invalid UTF-8 byte handling is a direct runtime-only test.

[Fixed-PHP PR CI40121644555](https://github.com/katsuakikureyama/GoHP/actions/runs/40121644555)
PASSed all **89 checks** on Go1.23.12 linux/amd64 / actual PHP8.3.36.
Clean merge checkout `78a4178bf5ff7d6cea3e52db5df5f31840d79f34` has tree
`7aa14fc6a9f384298bdf94705dc80c236b85d37b`, identical to the compiler head.
Native test exit0, vet, summary and artifact publication all succeeded.
Artifact `compatibility-php8.3-40121644555-1`, ID12570818793, retained30days;
downloaded to `/private/tmp/gohp-ternary-ci.eomIpS/`. Initial CI success,
no manual retries/reruns. CI compares native stdout/normal exit, not general
stderr, exception or side-effect equivalence. Marker output establishes only
the fixture's visible evaluation order. The local Gossip observation separately
requires empty stderr; it is not a Gossip Actions result.

## Boundaries and remaining work

See the [unit contract](../stripslashes-helper-unit.md). Short ternary, null
coalescing, object truthiness, weak coercion, PHP warning/exception equivalence,
recursive `map_deep` / `stripslashes_deep` / `wp_unslash` and general PHP remain
unsupported. Earlier five standalone unit manifests and the ten-function bundle
retain their pins and results, without reruns. This sixth independent unit is
not a twelve-function bundle or full WordPress test.

Original WordPress files, native SDK/hooks, oracle policy implementation and
the separate GoHP feature stack/normal checkout are unchanged. No default-branch
merge, deployment, release or license finalization. HTTP/DB/theme/plugin/core
startup remain unverified. Local tooling remains WebMCP-excluded, with no auth
route change. No devboard card is linked, synced or marked Done.
