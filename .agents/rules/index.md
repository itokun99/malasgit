# Malasgit AI Workflow Rules

Entry point for AI coding agents working on this repository.

The canonical guide is [`AGENTS.md`](../AGENTS.md) at the repo root — read it
first. This directory splits that guide into focused rule files so agents can
load only what's relevant to the current task.

## Fork identity

`malasgit` is a fork of [`jesseduffield/lazygit`](https://github.com/jesseduffield/lazygit).
Remote is `git@github.com:itokun99/malasgit.git`. See
[`malasgit-specific.md`](malasgit-specific.md) for fork-only rules.

## When to read each rule file

| File | Read when... |
|---|---|
| [`commit-style.md`](commit-style.md) | About to create any commit, fixup, or amend. |
| [`testing.md`](testing.md) | Writing or modifying tests; demonstrating a bug. |
| [`pkg-layout.md`](pkg-layout.md) | Touching `pkg/` structure, imports, or `pkg/gocui/`. |
| [`i18n.md`](i18n.md) | Editing `pkg/i18n/english.go` or adding user-facing strings. |
| [`docs-sync.md`](docs-sync.md) | Editing `docs-master/`, `schema-master/`, or `userConfig` fields. |
| [`malasgit-specific.md`](malasgit-specific.md) | Anything that touches fork-only behavior (PRs, remote, identity). |

## Convention for rule files

- Each file is a focused extract from `AGENTS.md`. Content stays aligned with
  the canonical guide; if they drift, `AGENTS.md` wins.
- Do not duplicate sections already present in `AGENTS.md` — link instead.
- Files are written for AI consumption: terse, imperative, no prose padding.