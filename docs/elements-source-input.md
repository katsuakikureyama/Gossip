# Source-built input → original serialization/class decision

The [execution record](evidence/elements-source-input-2026-10-26.md) reports
all24 actual typed/caller-state comparisons and vet PASS on the first committed
snapshot, without repeating earlier successful configurations.

This separate configuration removes the observer field-write bridge from
[the prior24-case unit](elements-serialization-path.md). All24 corpus bytes and
the three original WordPress declarations remain unchanged. Only the input
construction path, owned driver, observers, harness and compiler pin change.

The owned `gossip_elements_source_input($input)` function creates the owned
`GossipBlockSupportsInput`, assigns `$holder->supports = $input` in PHP source,
then returns the holder. Both native PHP and generated Go execute that exact
driver; neither observer constructs the holder or writes its exported field.
Go's observer has no dependency on a generated concrete struct or field layout.
It only calls the generated factory and observes the actual supports getter.

GoHP #71 main `ce38319b4245ccea0a7ae627f53b4396acabfce6` supplies simple
assignment to an existing public untyped property on
a defined, statically known generated class variable, including known aliases
and method `$this`. Unknown function-parameter receivers, dynamic/undeclared/
chained/indexed property writes and constructors remain excluded. Arrays detach
at setter/getter boundaries; object/arrow identifiers retain identity. Mutable
property reads use actual-value guards rather than permanently assuming the
default's type. This permits a null-default property to hold input array data
without adding array property defaults or arbitrary host-object writes.
The identical clean integrated feature tree passed509 CI checks/70 catalog
entries on fixed PHP8.3.36, including parallel #70 str_starts_with. Older compiler
pins remain unchanged.

The unchanged originals are `_wp_array_get` (functions.php5233–5272),
`wp_should_skip_block_supports_serialization` (block-supports/utils.php22–35)
and `wp_should_add_elements_class_name` (block-supports/elements.php33–97).
Sixteen direct cases call the real skip utility; eight owned compositions pass
actual utility-derived options to the real class-necessity helper. The native
originals are verified user-defined, their complete ranges/source hashes bound.
One translation/build yields one binary with three originals, two owned functions
and one owned input class. Earlier24/334 configurations/pins/evidence stay exact
and are not rerun or summed into a larger integrated result.

Observers measure actual return type/value plus factory source input, supports/
nonobject input, feature_set, feature and parsed block after the call. They must
match PHP and authored expectations and remain unchanged. Existing corpus covers
raw bool/null/int/string skip values, strict scalar lists, partial/all skips,
h6/hover/non-null false styles and absent styles. Factory source-input state is
a new separate observation, not copied from expectations. Shared transport
adapters/hashes and bounded node limits (depth12/256 entries/4096-byte strings/
int±1e6) are retained, with array/null suppressed-read intermediates only.

From nested `gossip/`:

```sh
GOSSIP_GOHP_SOURCE=/absolute/path/to/GoHP GOCACHE=/private/tmp/gossip-go-cache go test -count=1 -v ./tests/gohp -run '^TestOriginalWordPressElementsSourceInput$'
GOCACHE=/private/tmp/gossip-go-cache go vet ./tests/gohp
```

Record actual results/snapshots before claiming PASS. Fixed PHP8.3.36/C and Go
images, nonroot/read-only/network-none/resource limits and same-process oracle
admission remain active. The owned class is input data, not original WP_Block_Type,
its registry/constructor, a service replacement or a complete render hook. CSS/
HTML/unique IDs/bootstrap/HTTP/DB/themes/plugins and general error/coercion/
reference behavior remain unverified. No production target configured/created.
Capability `gossip.dev.wordpress-unit.translate-compare` stays WebMCP-excluded:
local developer fixture, no browser/server operation or route/auth change.
Devboard unidentified/not synced; tracked in PR/current-state.
