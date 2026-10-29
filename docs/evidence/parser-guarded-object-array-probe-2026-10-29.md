# Original parser guarded object-array admission — 2026-10-29

The complete unchanged Parser/Block/Frame closure passes the former unknown
object-to-array cast rejection at394. The next measured compiler rejection is
original `array_push($this->stack, ...)` at165. No original parser Go/report/
build/run or G83-PARSER acceptance is produced.

## One probe and exact identities

- WordPress/WordPress `54106f6c545fc8b881f6b39b7963002bece6c28c`,
  version7.2-beta1-64693; complete original sources and existing notice/provenance.
- Unchanged [closure manifest](../../gossip/tests/gohp/parser-source-closure.json),
  SHA256 `89f897c64146705559eae41bb50d49aff57abad0a11ee02d16d27fe78d64adaa`.
- GoHP [PR135](https://github.com/katsuakikureyama/GoHP/pull/135), clean probe
  code `dbbcdc225247b0b826267ce0546814d9d7cf6e0a`, tree
  `cc78cd70c8de172658a6be97937561c207127b3a`.
- Cached Go1.23.0 DarwinARM64/CGO0, offline modules, and the existing shared
  cache.

One invocation, personal paths replaced by stable placeholders:

```sh
CGO_ENABLED=0 GOCACHE=CACHE GOPROXY=off GOSUMDB=off GOTOOLCHAIN=local GOWORK=off \
go run ./cmd/gohp -manifest GOSSIP_ROOT/gossip/tests/gohp/parser-source-closure.json \
  -source-root GOSSIP_ROOT -entry wp-includes/class-wp-block-parser.php \
  -php-version 8.3 -experimental-pcre2 -out FRESH/parser.go -report FRESH/report.json
```

Exit1, stdout0bytes, stderr102bytes, SHA256
`9db57cfa3d6c3e0411ef3e4725b12b1d18e38b5d6258c605bc15abb53d689c3e`:

```text
wp-includes/class-wp-block-parser.php: PHP004 line 165: unsupported function array_push
exit status 1
```

The final line is from `go run`. No output/report/staging artifacts exist; the
empty temporary directory was removed. The prior line394 cast gate is therefore
only a source-admission advance. Traversal reaches a lower source line after the
semantic gate changes; it does not establish original parser execution, runtime
behavior, or a regression in the preserved source.

## Bounded guarded cast

PR135 retains the existing `object_array_cast` lowering and `ObjectToArray`
runtime API. An inferred-unknown expression may reach it only when the actual value
is a nonnil `PublicPropertyObject`; known non-object operands and undefined
reads still reject statically. Unknown null, scalar, array, closure,
partial/invalid-marker, and typed-nil generated objects reject at runtime. Evaluation is
once-only; declared-property order, null/falsy values, nested-array copy
behavior, and nested-object identity are retained.

This adds neither an IR/runtime API nor general structural host-object
conversion. Complete structural hosts explicitly opt into the existing snapshot
ABI; this does not claim unforgeable host provenance. The owned fixture's
extracted local uses an indexed frame, not `array_pop`; it is not a replacement
or rewrite of the original parser. Full stack mutation, snapshot behavior,
tokenizer, and JSON integration remain required.

Focused local codegen/compat/runtime checks passed once at the probe revision
in 0.386s/0.601s/0.421s. An initial new-fixture run failed because it exercised
an unsupported nested write and external-module setup, not the cast contract.
After test-only one-level read-copy-replace repairs and Go1.23.0 offline
`-mod=mod` setup, only the earlier failing fixture/host-entry tests were rerun;
they passed once in0.771s at clean code
`af2cf093f71a311ad83a87c1a41ae5c687f87e21`, tree
`7bcf1453ede92d9ae515a6f2dd69d55e9142adb8`. No already-passing production
checks were rerun.

Tested candidate head `7457217dcc4594e26ad11449ef846417be38eba2`, tree
`1adeba9e3f998ba5a77a29e2f4ca71dcd1abd908`, is merged at main
`10086b9528e78ce19cf4cf3c1c128248b981c80f` with the same tree. Push40887941868
and PR40887952996 attempt1 all4 target jobs SUCCESS. Saved own artifacts
12904335380 (PHP8.3, SHA256
`57679c4ec7ac2c4a5d7297db380de06b2ac8b0404f9ce49932e57534661e0bf8`) and
12903342470 (PHP8.4, SHA256
`6ab4dd2ffb8d1c97530c6a9e1865c62d5cfb1c7d0e99182cafd88009f19bd8e3`) match
the clean head/tree, Go1.23.12Linuxamd64/CGO0, actual PHP8.3.36/8.4.27,
testexit0, and overall PASS. All six new checks pass on both targets.

The owned fixture SHA256 is
`b15f5915e3ebc3ea410758305e894ccce1e3b10edf04b2daeb202e47e6ab37ec`.
Native/Go stdout is280bytes equal, SHA256
`6ec53193fc676aa19075315f22625c6c3aa72a080926e8e79fa6f0736d311d85`, exit0
and quiet stderr0. General stderr/exception/side-effect equivalence is
not-checked. 826total: PHP8.3 808PASS+18not-measured; PHP8.4
824PASS+2not-measured. See the [canonical PR result](https://github.com/katsuakikureyama/GoHP/pull/135#issuecomment-6498911020).
Exact-main automatic [CI40888584852](https://github.com/katsuakikureyama/GoHP/actions/runs/40888584852)
attempt1 is COMPLETED/SUCCESS, both jobs123498867749/123498868073 SUCCESS.
See the [integration result](https://github.com/katsuakikureyama/GoHP/pull/135#issuecomment-6498996066).
No duplicate same-tree artifact readback or successful manual check/probe rerun.
Duplicate134 is CLOSED with branch/commits/seven-check/151-byte fixture evidence
preserved independently; none is combined into canonical135's six-check/
280-byte evidence.

Earlier PR133/PR131 evidence remains at its own pins. This Gossip increment
changes documentation only. Original source, manifests, pins, auth/routes,
deployment, and CLI/WebMCP exclusion are unchanged; Gossip CI is unconfigured
and the external devboard is unidentified/not synchronized.
