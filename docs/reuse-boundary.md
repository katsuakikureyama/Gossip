# Original translation and native integration boundary

Current design map, 2026-10-26 JST. This replaces the
[historical boundary overview](history/reuse-boundary-through-pr44.md), not its
evidence or source pins. Implemented, planned and unverified are separate states.

## Two projects, different responsibilities

**GoHP** owns general-purpose, version-aware PHP parsing/semantics, Go source
generation and PHP compatibility runtime operations. WordPress names, special
cases and handwritten replacements do not belong in its compiler core.

**Gossip** owns pinned original WordPress inputs, application cases, native Go
plugin experiments and future integration adapters. Demonstrated consumer gaps
become bounded GoHP requests. A newer GoHP main does not silently repin or certify
an existing Gossip configuration. GoHP currently admits separate 8.3/8.4 targets;
Gossip's retained comparisons here are **PHP 8.3.36 evidence only**.

## Responsibility map

| Surface | Intended path and owner | Current evidence / admission boundary |
|---|---|---|
| Selected WordPress helper functions | Original PHP → GoHP → generated Go; Gossip selects and compares | 26 originals / 334 cases in one recorded configuration; not whole-file or whole-core support |
| Block parser algorithm and data classes | Preserve original algorithm; GoHP supplies general PHP semantics | Complete Block and Frame classes and selected freeform/add_freeform paths have separate evidence; full parse/tokenizer/stack unverified |
| Original hooks: plugin.php / WP_Hook | Original behavior remains the target for translation and comparison | Native SDK's five owned hook cases do not establish WP_Hook parity; references, dispatch mutation and reentrancy require their own evidence |
| PHP builtins and extensions, including preg_match | GoHP compatibility runtime, possibly backed by an explicitly admitted external library | Backend choice is not PHP API equivalence; see [parser backend gates](parser-backend-assessment.md). No regex backend selected |
| Native Go plugin SDK | Gossip-owned trusted in-process API | Existing isolated registry experiment; not wired into WordPress request/hook lifecycle or a PHP plugin loader |
| DB and HTTP integrations | Candidate Gossip host adapters behind explicit contracts | No adapter selected or implemented here; original DB schema, query/transaction/error/side-effect contracts remain authoritative; not the prototype JSON store |
| Themes, template include, PHP autoload | PHP semantics in GoHP; source discovery, packaging and host integration in Gossip | Unverified. Not automatically reassigned to handwritten Go or claimed unnecessary because an SDK exists |
| Assets, Gutenberg markup and theme.json | Preserve originals; Gossip validates their eventual consumers | Files being present is not rendering, editor or theme compatibility |
| Browser sessions, authorization and routes | Future Gossip host integration with explicit route contracts | No new server route/authentication model or production service exists in this change |
| Old prototype editor, importer, JSON persistence and HTTP server | Separate Gossip-prototype history | Excluded from this fork's active runtime; not replacements for inherited WordPress |

Evidence links and exact configuration boundaries are in the [README](../README.md)
and [current state](current-state-fork-baseline.md). Rows marked candidate or
unverified are not committed implementation choices. There is no approved list
of WordPress APIs to discard. Unsupported is not silently replaced, stubbed or
excluded from a claimed complete unit.

## Original-source acceptance

1. Preserve inherited WordPress files, assets, notices and ancestry. Go experiments
   stay under `gossip/`; project README/docs are additions, not upstream rewrites.
2. Pin whole original files and selected excerpt/bundle/driver identities. Explicit
   selection is allowed as a bounded experiment, but omitted methods stay visible.
3. Keep original translated code, owned drivers/observers and native adapters
   distinguishable. A Go reimplementation may be a separately evaluated project
   component, never evidence that GoHP translated the original.
4. Admit the complete original parser only after its actual dependency closure and
   required cases pass; then original hooks and broader callers. GoHP's
   [roadmap](https://github.com/katsuakikureyama/GoHP/blob/e549d184e2e76bec4eb7c364ef51bef7bb40a7f6/docs/roadmap.md)
   requires the original WP_Hook, not an owned substitute.
5. Report source selection, generation, build, execution and application acceptance
   separately. Preserve failed/unmeasured outcomes. Do not sum distinct corpora or
   promote old compiler/oracle evidence to current main or another PHP target.

GoHP's [first whole-file corpus manifest](https://github.com/katsuakikureyama/GoHP/blob/e549d184e2e76bec4eb7c364ef51bef7bb40a7f6/docs/application-corpus.md)
fixes `functions.php` source identity. It does not measure PHP callsites, dependency
closure or compatibility and does not replace Gossip's per-experiment manifests.

## Publication, licensing and operating policy

Use branches and PRs in **this fork**, not upstream WordPress. Under the owner's
current authorization, ordinary verified changes may be merged without repeated
approval requests. Preserve historical restrictions as history, not current gates.
Deletion, new billing, unspecified deployment destinations and other material scope
expansion remain separate decisions. No production destination is configured here.

Retain upstream licensing/provenance; neither a fork label nor a repository split
settles generated-code, runtime, plugin or distribution obligations. GoHP's
[permissive preference remains unfinalized](https://github.com/katsuakikureyama/GoHP/blob/e549d184e2e76bec4eb7c364ef51bef7bb40a7f6/docs/licensing.md).
This document selects no license. Use the project name Gossip; describe WordPress
as the source/compatibility target, not an affiliation or an achieved replacement.
The [Foundation policy](https://wordpressfoundation.org/trademark-policy/) distinguishes
descriptive references from naming a product or implying endorsement; check it
again before branding or distribution decisions.

New public logs and PR bodies use `GOSSIP_ROOT`, `GOHP_ROOT` or `SOURCE_ROOT` in
place of personal absolute paths. Keep hashes, commands, versions and outcomes;
mark path substitution explicitly. Existing historical records are not silently
rewritten by this documentation change; old public-path cleanup is separate work.

## Capability and authentication boundary

This change adds documentation only. Existing local source inspection and
translation/comparison capabilities remain **excluded from WebMCP**; no new user
operation, UI callsite, server operation, automatic submit or persistent audit
service is introduced. Native Go plugin registration is not a browser capability.
No browser/S2S/MCP/webhook/public route is added or reclassified, and no sessions,
cookies or bearer tokens are transferred. Future adapters need their own capability,
authorization, validation, result-state and audit contracts before admission.
