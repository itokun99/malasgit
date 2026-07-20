# Commit Style

Canonical source: [`AGENTS.md` § When to commit, How to structure commits,
Iterate with `fixup!` commits, Surface mid-implementation decisions](../../AGENTS.md).
This file is a focused extract.

## When to commit

Commit as you go. Do not leave completed work uncommitted. Once a logical unit
is done and the tree is green, commit it.

A task implicitly includes "and commit your work" unless the user says
otherwise.

## Fine-grained history

Commits should be as small as possible while still being meaningful and
self-contained.

- **Every commit must compile and pass all tests.**
- **Every commit must be `gofumpt`-formatted** — run `just format` first.
- **Every commit must be lint-clean** — run `just lint` first.
- **Commit messages explain _why_, not _what_.** No paraphrasing the diff.
- **Wrap message body to 72 characters.** Subject up to 80 (or slightly more
  if needed).
- **Do not use conventional commits** (no `feat:`/`fix:`/`chore:` prefixes).
  Match the plain English imperative style of the existing history.
- **Separate preparatory refactorings from behavior changes.** Pure refactors
  in their own commit; behavior-change commit as small as possible. This
  applies even when the refactor is only discovered mid-change — stage hunks
  or reset and recommit.

## Iterate with `fixup!` commits

When refining committed work, use `git commit --fixup=<sha>` so the refinement
sits alongside its target for `git rebase --autosquash`.

- Use `fixup!` even when the target is HEAD. Do not `--amend`.
- If a fixup makes the target's message inaccurate, use `amend!` instead:
  `git commit --fixup=amend:<sha>`.
- **`amend!` message shape:**
  ```
  amend! <original subject>

  <new subject>

  <new body>
  ```
  The first line is the matcher; everything after the blank line is the
  complete replacement message and must begin with a subject line.
- **Never squash the fixups yourself.** Leave them as separate, reviewable
  commits. The user decides when to autosquash. Do not run
  `git rebase --autosquash`, do not `--amend` fixups into targets, do not
  reorder or collapse them.

## Surface mid-implementation decisions

When a real fork surfaces during implementation — a design choice, tradeoff,
scope cut, or a discovery that reshapes the work — stop, lay out the options
and your recommendation, and let the user weigh in.

This applies with equal force to unforeseen discoveries (latent bug, race,
wrong assumption, unhandled case), not just decisions you set out to make.
Don't quietly fold a self-directed fix into the branch.

Trivial mechanical choices with one sensible answer don't need a checkpoint.
Genuine forks where a reasonable person might pick differently do.

## Prefer the cleaner design over the smaller diff

When a task could go either by tacking onto existing code or by first
restructuring slightly, choose the restructuring. "Minimal change" is not a
goal in itself; a readable final state is.

This is not license for speculative abstraction. If the current change would
be clearer after extracting a method, splitting a function, or adjusting
names, that refactor is part of the task.

## Don't present "live with the bug" as an option

When investigating a defect, "accept the race / leave it as-is / document it
and move on" is not a valid option. Present actual fixes. If a real fix is
genuinely out of reach, say so plainly.

## Unify duplicated logic before changing it

When a fix or feature would land in logic duplicated across call sites, do the
behavior-preserving refactor that unifies them first, then make the change
once. Never sequence a branch so one commit introduces a divergence that a
later commit repairs.

If the need for the refactor surfaces mid-branch, suggest to the user to
rewrite history to move the refactor earlier — but don't do it without asking.