# Thirteen-original recursive helper bundle

This separate configuration translates thirteen complete, unchanged WordPress
functions together with eleven owned PHP composition drivers, builds one native
Go executable and compares it with the same source bundle running on PHP8.3.36.
The thirteen originals are the previous twelve-helper selection plus `wp_slash`.
The compiler/runtime pin is unchanged from the standalone wp_slash unit:
`03d433e73503e32a84f310f023fc0f65c9b188cd` (GoHP Draft #27).

## Corpus and observation

The 33 cases comprise thirteen direct probes and twenty compositions. Twenty-four
cases carry forward the earlier integrated corpus into this new configuration;
nine add the wp_slash direct probe and eight compositions:

- Three actual-string slash/strip roundtrips, including NUL, backslash, quotes,
  multibyte and empty strings.
- Three list slash/shallow-strip operations, including empty and mixed scalar
  lists. A nested list deliberately remains escaped after shallow stripping:
  this is not recursive wp_unslash support.
- One capture-list → recursive slash → newline preservation → string strip call.
- One recursive slash → array_is_list / array_keys / array_values composition.

All composition calls execute in translated PHP drivers, not observer-side Go
implementations. Observers dispatch only to named originals/drivers and normalize
typed results. Each actual observation includes all caller arguments after the
call; these must match the input fixture as well as the PHP observation.

Input transport admits null, bool, modest signed integers, JSON strings and
finite nested sequential list arrays only. Go uses UseNumber/Int64 and checks
sequential int64 keys on outputs and post-call arguments. No floats, associative
JSON maps, objects, references, cycles or arbitrary large-integer precision claim.
Explicit expected results are authored independently; roundtrip alone is not
used to establish quoting behavior (the direct probe and prior unit remain
separate evidence).

## Exact source and execution boundary

The new manifest binds the WordPress baseline, entire original-file hashes,
function ranges and compiler/runtime commit. Extraction records each excerpt's
hash and bundle line offset. Driver/case/observer/generated/report/binary hashes
are recorded, and the report hashes must match before compilation.
The existing immutable PHP image, allowlist, same-process version guard,
LC_CTYPE=C guard, non-root/read-only/no-network/resource-limited containers
remain unchanged. Matching return type/value, caller arguments, authored
expectations, empty stderr and exit0 are required; generation/build are only
intermediate statuses. Missing local opt-in is SKIP, not PASS.

```sh
cd gossip
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP \
  go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressRecursiveHelperBundle$'
go vet ./tests/gohp
```

Earlier seven standalone units and ten-/twelve-helper bundles retain their
original manifests, corpora, pins and execution records. This new result does
not claim every earlier input was revalidated. No compiler change or repeated
compiler CI is required; the exact-pin CI181-check PASS is reused.

## Remaining original-source dependency

`wp_unslash` delegates to `stripslashes_deep`, then `map_deep`. The latter's
unchanged body requires element/property assignment, object traversal and dynamic
call_user_func, which are outside this compiler subset. This bundle does not
rewrite that original into an array-only equivalent or claim to support it.
Whole-file/core startup, HTTP/DB, themes/plugins, SDK/hook integration, general
callbacks, arbitrary-depth recursion and object aliasing remain unverified.

## Capability exclusion

`gossip.dev.wordpress-unit.translate-compare` remains an opt-in local developer
capability explicitly excluded from WebMCP. No user-facing read/write operation,
route, browser/S2S/MCP/webhook auth, cookie forwarding, automatic submit or audit
API changes. No default-branch merge, deployment, release or license decision.
Devboard card is unidentified/unsynchronized, not Done.
