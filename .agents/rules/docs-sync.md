# Documentation Sync

Canonical source: [`AGENTS.md` § Don't edit files under `docs/`, Don't search
outside the working tree](../../AGENTS.md). This file is a focused extract.

## Don't edit files under `docs/`

`docs/` is the documentation rendered on GitHub for the current _release_.
Users read it as the reference for the version they're running. If a new
feature is documented in `docs/` in the same PR, the docs end up describing
features users don't yet have until the next release is cut.

So:

- Document new features in `docs-master/` only. The release process
  (`scripts/update_docs_for_release.sh`) copies `docs-master/` to `docs/` at
  release time.
- For changes to `userConfig` fields, don't edit `docs-master/Config.md` by
  hand either — the relevant section is auto-generated from the struct field
  doc comments. After editing the struct, run `just generate` and include the
  regenerated `docs-master/Config.md` (and `schema-master/config.json`) in
  the commit.
- Don't hard-wrap the doc comments on `userConfig` fields. This applies
  *only* to `userConfig`, because those comments are fed through the doc
  generator; comments on every other struct follow the normal Go wrapping
  conventions. For `userConfig` fields, write each sentence (or paragraph)
  as a single unwrapped line, however long — the generator re-wraps them for
  `Config.md` (see `wrapLine` in `pkg/jsonschema/generate_config_docs.go`).
  Manually wrapping a sentence across several `//` lines defeats this: the
  generator preserves your arbitrary breaks as hard line breaks and embeds
  `\n` at those points in the generated `schema-master/config.json`
  description. (Putting genuinely separate sentences on their own lines is
  fine; just don't split one sentence across lines.)

## When to regenerate

Run `just generate` whenever you:

- Add, remove, or rename an integration test (regenerates the test list).
- Change keybindings (regenerates the cheatsheets in
  `docs-master/keybindings/`).
- Edit a `userConfig` field's doc comment (regenerates `docs-master/Config.md`
  and `schema-master/config.json`).

CI fails if these are stale. Always include the regenerated files in the same
commit as the change that triggered the regeneration.