# Seventeen-original yes/no, HEIC, charset and array bundle

This independently pinned configuration translates seventeen byte-unchanged
WordPress helpers with twenty-five owned composition drivers, then builds one
generated executable. It extends the16-original HEIC bundle with
`bool_from_yn()` (functions.php1621–1623), not a replacement algorithm.
The manifest pins three full source hashes, exact fragments, compiler commit
and immutable builder/oracle images. The new yes/no fragment is also checked
against its exact range/hash, alongside the existing guarded excerpts.

Compiler GoHP `49a302e44464623dc4912bafbcf44192759d7878` (Draft #44) is
unchanged. Its existing336-check native CI evidence is reused without another
compiler test/CI run. This new Gossip configuration requires fresh integration
evidence; individual or older bundle PASS results do not substitute for it.
The separate newer control-flow compiler stack is not adopted here.

## Corpus and observation

Total149 calls: all119 previous HEIC-bundle cases with unchanged arguments,
expectations and post-call expectations; all24 standalone yes/no cases; and
six new compositions. Direct103/composed46 across42 operations.

The three new owned drivers each have two cases:

- `gossip_slice_yn`: original slice, then map the original yes/no predicate.
- `gossip_slice_unescape_yn`: original slice, shallow string unescaping,
  then map the original yes/no predicate.
- `gossip_yn_boolean_pair`: call both original predicates on the same string
  and return their actual results in order. `Y` gives `[true,true]`, while
  `yes` gives `[false,true]`; the helpers have distinct contracts.

Direct yes/no and Boolean-pair operands must be actual strings. Every root
value of the new yes/no slice operations must be string or null, including
unselected entries. The original slice excludes null/missing entries before
mapping, so no non-string reaches `bool_from_yn`. Bool, integer, nested array
and object values are outside those new input domains. Existing HEIC scalar
guards and older nested sequential-list transports remain unchanged.

The new cases cover ordered mixed/negative keys, duplicate selectors,
null/missing removal, empty input/selection, escaped characters and NUL.
Actual typed returns, pair ordering and caller arguments after each call must
match PHP and authored expectations. Observers never construct results from
expected fields. Pair arrays use explicit `[key,value]` transport; integer
keys/values remain bounded to ±1,000,000, nested list depth to8. No nested
associative values, objects, floats, references or cycles.

## Execution and exclusions

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressYNHelperBundle$'
go vet ./tests/gohp
```

Missing opt-in means skip, not PASS. Existing immutable-image admission,
same-process PHP8.3.36 guard, LC_CTYPE=C, non-root/read-only/network-disabled
containers and resource limits remain. Source/generated/report hashes must
agree before build. PASS additionally requires actual observations to match
the oracle and expectations with empty stderr/exit0. Results are recorded
separately after execution.

Eleven older standalone units and all previous bundles remain pinned and
unchanged; this is not every historical standalone input. General Boolean
coercion, MIME image-content processing, UTF-8 text validation, option-backed
charset wrappers, array_all/arrow closures and whole-file/core/HTTP/DB/theme/
plugin behavior remain unverified. No master change, merge, deploy, release
or license decision.

Local capability `gossip.dev.wordpress-unit.translate-compare` remains
explicitly WebMCP-excluded. No browser/S2S/MCP/auth-route changes.
Devboard remains unidentified/unsynchronized, not Done.
