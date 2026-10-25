# Original yes/no-label predicate

The eleventh standalone unit extracts byte-unchanged `bool_from_yn()` from
WordPress functions.php1621–1623. The original is exactly:

```php
function bool_from_yn( $yn ) {
	return ( 'y' === strtolower( $yn ) );
}
```

The manifest pins WordPress, whole-file SHA256, exact range, compiler and
immutable builder/oracle images; the harness also compares complete excerpt
bytes. GoHP `49a302e44464623dc4912bafbcf44192759d7878` (Draft #44) already
supplies string-only ASCII `strtolower` and scalar strict equality. No
compiler change or new builtin is required. Its existing336-check native
CI evidence is reused, not rerun.

## Domain and observation

Only actual strings are admitted. Twenty-four authored cases include the
two true labels `y`/`Y`; `n`/`N`, words yes/no/true/false, digit strings,
empty and whitespace, leading/trailing space, newline/tab, embedded NUL,
Latin/fullwidth/Greek/Japanese text and repeated letters all return false.
There is no trimming, Unicode normalization, general truthiness or general
Boolean parsing. Non-string coercion, null, numbers, arrays, objects and
PHP warning/exception behavior are explicitly outside this unit.

Both isolated observers invoke the actual original/generated helper and
assert the actual result is bool, recording its value and actual caller
input after each call. Observed output never comes from expected fields.
The host requires equality with native PHP and independently authored
expectations, plus empty stderr/exit0.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressBoolFromYNHelper$'
go vet ./tests/gohp
```

Opt-in absent means skip, not PASS. The existing immutable-image admission,
same-process PHP8.3.36 guard and non-root/read-only/network-disabled/resource
limits remain. Source/report/generated hashes must match before building.
Generating/building alone is not compatibility evidence; results are recorded
separately after execution.

## Selection and separation

`_canonical_charset()` is not substituted or stubbed: its unchanged public
`is_utf8_charset()` dependency requires defaults, null coalescing and
`get_option()`. `wp_is_numeric_array()` requires a Boolean return declaration,
`array_all` and a two-argument arrow closure. Those contracts remain outside
this step; adopting another control-flow compiler alone would not resolve them.

This adds a seventeenth distinct original helper in a separate unit, not a
17-original integration result. The ten older standalone configurations and
16-original119-case bundle remain pinned and unchanged, not rerun.
WordPress original files, SDK/hooks, oracle policy, prototype and master
README-only state are not changed. Whole-core/HTTP/DB/theme/plugin behavior,
deployment and license resolution remain unverified; no merge/release.

Local developer capability `gossip.dev.wordpress-unit.translate-compare`
remains explicitly WebMCP-excluded. No browser/S2S/MCP/auth-route change.
Devboard is unidentified/unsynchronized, not Done.
