# Block style serialization decision → class necessity

This separate configuration selects three complete unchanged declarations:

- `wp_should_skip_block_supports_serialization`, block-supports/utils.php22–38;
- its `_wp_array_get` dependency, functions.php5233–5272;
- `wp_should_add_elements_class_name`, block-supports/elements.php33–97.

Whole-source and fragment ranges/hashes bind the original WordPress bytes.
GoHP #69 main `731ef251ac907326f060aac5f1b5b994ac8615bb` supplies `is_object`
for generated class/arrow representations and retains #68 str_contains. The
exact integrated feature tree passed498 CI checks/68 catalog entries on fixed
PHP8.3.36. This manifest does not change any older compiler pin or result.

Sixteen direct cases measure nonobject/empty-feature early returns, missing/null
supports, true/false/null/int/string raw settings and strict list membership.
Despite its docblock, the original utility does not always return bool: a
nonarray setting is returned unchanged. Null, `0`, `'0'` and `'false'` remain
distinct values with distinct types. Eight owned compositions compute button/
link/heading skip options by calling that real utility, then call the real
class-decision helper. They cover Boolean/list/raw-string skip behavior, partial
skip, h6, hover, a nonnull false style leaf and absent styles.

## Input boundary, not a WordPress service stub

`GossipBlockSupportsInput` is an owned class with one public null-default
`supports` property. Both runtimes construct that class; their observers inject
the bounded supports data (Go exported field / PHP property assignment) outside
the translated source. The original utility is untyped and reads this property;
no fixture class computes WordPress behavior. This does not validate original
`WP_Block_Type`, its constructor/registry, source property writes or array
property defaults. The single owned driver models only the utility→options→
class-decision composition, not the original render hook or all its early exits.

One translation/build produces one binary with all three originals, one owned
class and one owned function. Go/PHP observers execute all24 cases, compare
actual type/value with each other and expectations, and record supports/nonobject
input, feature_set, feature and block after each call. Inputs must be unchanged.
The shared bounded codec retains depth12/256 entries/4096-byte strings/int±1e6;
reached nested suppressed-read intermediates are arrays/null, skip-list values
are scalars. Old334-case integration remains unchanged and is not rerun or
arithmetically promoted to a larger single-binary claim.

From the nested `gossip/` module:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressElementsSerialization$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

The normal container/oracle policy is retained: fixed PHP8.3.36/C and Go images,
nonroot/read-only/network-none/resource limits and same-process PHP admission.
Record actual execution evidence before reporting PASS. Full render hooks,
unique IDs, CSS generation/storage, HTML parsing, Theme JSON, WordPress bootstrap,
HTTP/DB/themes/plugins and general warning/error/coercion behavior are unverified.
No production target is configured/created. Capability
`gossip.dev.wordpress-unit.translate-compare` remains explicitly WebMCP-excluded:
local developer fixture, no browser/server operation or route/auth change.
Devboard unidentified/not synced; progress is tracked in PR/current-state.
