# Original empty-array helper — 2026-10-25 JST

## Exact targets

- Gossip tested source: `1596c5ac35846db59612098fe13d4f5736f23655`.
  The test began on these exact uncommitted bytes, committed without changes
  while it ran. An initial `git add` used the wrong relative directory; it
  staged nothing and did not change the tested files.
- Original WordPress: `54106f6c545fc8b881f6b39b7963002bece6c28c`.
- GoHP compiler/runtime: `c67fbefe34b1b098c0a9f475504d06bd014de23c`,
  [Draft PR #8](https://github.com/katsuakikureyama/GoHP/pull/8).
- Host test driver: Go 1.22.1 darwin/arm64. Isolated Go builder:
  Go 1.25.14 linux/arm64. PHP oracle: PHP 8.3.13.
- Exact source, excerpt, image, compiler, observer, report and binary hashes:
  [machine-readable observation](empty-array-2026-10-25.json).

## Command and result

From `/Users/kureyamayuri/Desktop/Gossip-wordpress/gossip`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-legacy-array.eU8sLU/GoHP \
  go test -count=1 -v ./tests/gohp
```

PASS on the first test run: integration 13.14 seconds, package 13.328 seconds.
The source was not changed after PASS and the test was not repeated. Existing
SDK/hook tests were unchanged and not rerun.

Six complete original function excerpts were translated together. The added
`__return_empty_array()` is copied from `wp-includes/functions.php` lines
7222–7224, including its original `return array();` expression. Neither the
original file nor the extracted function text is rewritten to `[]`.

Both isolated runtimes returned the same typed observations and matched the
owned expected fixture: bool true, bool false, int 0, **array []**, null, and
string empty. Both exited 0 with empty stderr. The independent echo probe still
compares only the five scalar functions; array-to-string coercion is deliberately
not used as evidence of an array return.

The Go observer accepts an actual pinned-runtime `rt.Array` only when its Entries
length is zero; it emits a non-nil JSON empty list. Any nonempty array fails this
observer. This proves this empty-array result, not arbitrary array serialization
or general PHP array equivalence. JSON and type reflection are instrumentation,
not translated PHP builtins.

## Compiler boundary and history

The GoHP change parses case-insensitive `array(...)` through the same literal
item parser/IR as `[]`; it does not modify runtime or code generation semantics.
Its repository-owned fixture checks keys, mixed nesting, trailing commas and
iteration; seven negative cases retain explicit unsupported/malformed boundaries.
Compiler validation and version-specific CI are tracked in its Draft PR #8,
separately from this PHP 8.3.13 consumer result.

GoHP [PR CI run 40112122164](https://github.com/katsuakikureyama/GoHP/actions/runs/40112122164)
passed on Go 1.23.12 with PHP 7.4.33 and PHP 8.3.30 in separate jobs. Both
record `diff.legacy_arrays` and all seven added rejection checks as PASS. The
actual PR merge checkout was `3ea6fb9249f155f0d269e80cb72ca41ff1535f43`, clean,
tree `2139b66e473c0a71cdafc9d000fc8112a51838a1` (matching the compiler head).
Artifacts `compatibility-php74-40112122164-1` (12566498142) and
`compatibility-test-40112122164-1` (12567221196) retain the run evidence for 30 days.
These compiler fixtures do not establish WordPress compatibility on PHP 7.4.

The earlier [typed-helper rejection evidence](typed-helpers-2026-10-25.md) remains
valid for GoHP `9c7389d`; it is historical evidence, not the current candidate's
expected result. The current v3 observation replaces the v2 rejected-unit field
with a sixth successful typed observation. It does not relabel the old run.

All compiler builds, PHP execution and generated-code execution used isolated
non-root containers with network disabled, read-only roots, resource limits and
test-owned temporary outputs. Source report hashes were checked before build.
The same immutable GoHP revision supplied compiler and runtime. Containers and
temporary generated files were cleaned; the dedicated compiler clone remains
for reproduction. No WordPress source, production service or default branch was
changed. No merge, release, deployment or license finalization is included.

Whole-file/core translation, PHP hooks, themes, plugins, HTTP side effects,
database mapping and general warning/exception semantics remain unverified.
The local-test WebMCP exclusion remains in effect. No linked devboard card is
known; none is marked Done by this result.
