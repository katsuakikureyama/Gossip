# Nested settings lookup — reference record (2026-10-26 JST)

The unchanged original `_wp_array_get` matches all **24 native PHP expectations**,
including actual return type/value and unchanged input/path/default argument
states. GoHP rejects the original and four isolated feature probes as recorded
below. Go/PHP differential status is **not-run**.
[Contract](../array-get-boundary.md), [full record](array-get-boundary-2026-10-26.json).

- Successful test commit `79f94688f21369ee932c608cda5dcce3664f6d08`, tree `771df55853ed38faccaee77598822202d8fa3f55`.
- Fixed GoHP #59 main `c3420d8644642c2fee56d62a08c88a2dcedc45ac`; compiler source unchanged.
- WordPress baseline `54106f6c545fc8b881f6b39b7963002bece6c28c`; functions.php5233–5272
  extracted byte-for-byte, full source/fragment hashes recorded.
- Actual PHP8.3.36/C and Go1.25.14linuxarm64. Same-process admission,
  non-root/read-only/network-off execution and resource limits retained.
- The source survey identified60 actual calls across12 files, including43
  Theme JSON and6 global styles/settings calls. Those callers are not executed.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/private/tmp/gohp-empty-array.nKn0kC/GoHP go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressArrayGetBoundary$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

Boundary test PASS13.31s/package13.521s, command exit0, vet PASS.
Native exit0 with empty stderr. Every compiler probe exits1 with empty stdout
and no generated/report artifacts. Observed diagnostics:

| Probe | Diagnostic |
|---|---|
| Original _wp_array_get | PHP002 line2: expected comma at "=" |
| Owned trailing default | PHP002 line2: expected comma at "=" |
| Owned bare-variable isset | PHP002 line2: expected "[" at ")" |
| Owned multi-operand isset | PHP002 line2: expected "[" at "," |
| Owned continue | PHP002 line2: unsupported statement at "continue" |

Exact diagnostic bytes and each probe-source hash are in the JSON. These are
parser rejections, not successful translations or a runtime equivalence result.
No original function is rewritten to bypass its syntax.

The initial source-range snapshot `6afc943102e98af6abb9e7e66f6acdd89f91d0dc`
stopped at the return line5271 and failed the complete-declaration check before
any native execution. The corrected range includes closing brace5272. The
corrected test then completed its first native/probe execution. After PASS only
documentation/evidence is changed; no duplicate comparison or compiler test runs.
Raw logs: `/private/tmp/gossip-array-get.zAAtgm/test-output*.log`.

Cases cover nested values/array returns, fallback scalars/arrays, omitted default,
present null and falsy values, missing/scalar intermediates, empty/invalid paths,
normalized and unusual keys, rejected path elements and associative path order.
No reference/object/float/error behavior or full Theme JSON execution is claimed.
The successful24-original262-case Go/PHP integration remains unchanged.

Next implement bounded trailing defaults and call arity, initialized bare/multi
isset with short circuit, and foreach continue; verify those compiler contracts
before adopting a new pin for an unchanged-original comparison.
