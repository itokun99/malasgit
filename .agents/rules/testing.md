# Testing

Canonical source: [`AGENTS.md` § Integration test conventions, Use
stretchr/testify, Demonstrating bugs before fixing them](../../AGENTS.md).
This file is a focused extract.

## Test commands

- `just unit-test` — `go test ./... -short`. Fast, runs in CI on every push.
- `just e2e` — run all integration tests headlessly.
- `just e2e <name>` — run one integration test headlessly.
- `just e2e-cli <name>` — run one with a visible UI (useful with `--sandbox`
  or `--slow`).

CI runs unit tests on Ubuntu and Windows across a git version matrix (2.32.0,
2.38.2, 2.44.0, latest) and integration tests on Ubuntu only. Match the matrix
when investigating version-specific failures.

## Integration test conventions

Don't bind views to local variables. Always chain method calls directly from
`t.Views().<View>()`:

```go
t.Views().Files().Focus().Lines(Equals("D  file03.txt"))
```

Patterns like `filesView := t.Views().Files().Focus()` followed by
`filesView.Lines(...)` are not how tests in this repo are written.

## Use stretchr/testify

Prefer `assert.Equal` and friends over hand-rolled `if` checks. Failure
messages are more useful and intent is clearer at a glance.

## Demonstrating bugs before fixing them

When fixing a defect, first land a commit that demonstrates the bug, then
fix in a follow-up. Use the `EXPECTED` / `ACTUAL` pattern.

The test asserts the current (wrong) behavior so it passes on broken code,
with the correct expectation preserved inline as a comment. The fix commit
swaps them: `EXPECTED` becomes the live assertion and `ACTUAL` is deleted.

Example shape:

```go
/* EXPECTED:
expectClipboard(t, Equals(worktreeDir+"/dir/file1"))
ACTUAL: */
expectClipboard(t, Equals(filepath.Dir(worktreeDir)+"/repo/dir/file1"))
```

The fix commit must be exactly "delete the markers and delete the `ACTUAL`
line" — no other edits. Structure the surrounding code so `EXPECTED` and
`ACTUAL` are drop-in replacements at the same syntactic position (usually by
placing the comment block between two adjacent chained calls).

Use this pattern only where it makes sense; don't apply by default.