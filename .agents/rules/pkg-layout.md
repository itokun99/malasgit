# Package Layout

Canonical source: [`AGENTS.md` § gocui is in-tree, Don't search outside the
working tree](../../AGENTS.md). This file is a focused extract.

## gocui is in-tree, not a dependency

The `gocui` TUI library is a fork maintained directly in this repo under
`pkg/gocui` — it's an ordinary package, not a Go module dependency. Don't
look for it in `go.mod` / `go.sum` or the module cache (`$GOMODCACHE`); it
isn't there.

When you need to read or change gocui internals (the task manager, the event
loop, worker/UI-thread dispatch, view rendering), edit `pkg/gocui` directly.

## Top-level packages under `pkg/`

| Package | Role |
|---|---|
| `pkg/app` | Application bootstrap and lifecycle. |
| `pkg/commands` | Git command construction and execution. |
| `pkg/gui` | GUI controllers, views, list rendering. |
| `pkg/gocui` | In-tree TUI library fork. |
| `pkg/i18n` | Translations. Edit only `english.go`. |
| `pkg/integration` | Integration test harness (`TestIntegration/...`). |
| `pkg/jsonschema` | Auto-generation for `userConfig` docs. |
| `pkg/utils` | Shared helpers. |
| `pkg/config`, `pkg/env`, `pkg/theme`, `pkg/updates`, `pkg/snake`, ... | Cross-cutting concerns. |

## Don't search outside the working tree

Never run `find` (or similar) from `/` or other paths outside the project.
All third-party code is vendored under `vendor/`, so dependency sources are
reachable from inside the working tree.

## Don't read model state right after a `Refresh`

A `Refresh` (or `RefreshFromWorker`) does its git work on a worker and then
*enqueues* the model update onto the UI thread. When `Refresh` returns, the
model is **not** updated yet — the write is still queued. Reading a field
synchronously right after refreshing its scope reads the stale, pre-refresh
value (this is true even for SYNC refreshes).

Put the read in `RefreshOptions.Then` instead — it's queued after the scope's
model writes, so it sees the fresh value:

```go
self.c.Refresh(types.RefreshOptions{
    Scope: []types.RefreshableView{types.FILES},
    Then: func() error {
        files := self.c.Model().Files // fresh
        return nil
    },
})
```

`Then` is a `func() error` and works with any non-`ASYNC` mode.