<p align="center">
  <img width="536" src="https://user-images.githubusercontent.com/8456633/174470852-339b5011-5800-4bb9-a628-ff230aa8cd4e.png">
</p>

<div align="center">

A simple terminal UI for git commands
<br/>

[![GitHub Releases](https://img.shields.io/github/downloads/itokun99/malasgit/total)](https://github.com/itokun99/malasgit/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/itokun99/malasgit)](https://goreportcard.com/report/github.com/itokun99/malasgit) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/f46416b715d74622895657935fcada21)](https://app.codacy.com/gh/itokun99/malasgit/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Codacy Badge](https://app.codacy.com/project/badge/Coverage/f46416b715d74622895657935fcada21)](https://app.codacy.com/gh/itokun99/malasgit/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_coverage) [![golangci-lint](https://img.shields.io/badge/linted%20by-golangci--lint-brightgreen)](https://golangci-lint.run/) [![GitHub tag](https://img.shields.io/github/v/tag/itokun99/malasgit?color=blue)](https://github.com/itokun99/malasgit/releases/latest) [![homebrew](https://img.shields.io/homebrew/v/malasgit?color=blue)](https://formulae.brew.sh/formula/malasgit)

</div>

## Credits

Malasgit is a fork of [jesseduffield/lazygit](https://github.com/jesseduffield/lazygit), originally created and maintained by Jesse Duffield. This fork is maintained by [@itokun99](https://github.com/itokun99).

## Features

### Stage individual lines

Press space on the selected line to stage it, `v` to start selecting a range, or `a` to select the entirety of the current hunk.

### Interactive Rebase

Press `i` to start an interactive rebase. Then squash (`s`), fixup (`f`), drop (`d`), edit (`e`), move up (`ctrl+k`) or move down (`ctrl+j`) any of the TODO commits, before continuing the rebase by bringing up the rebase options menu with `m` and selecting `continue`. The same actions work as a once-off without explicitly starting a rebase.

### Cherry-pick

Press `shift+c` on a commit to copy it and `shift+v` to paste (cherry-pick) it.

### Bisect

Press `b` in the commits view to mark a commit as good/bad and start a git bisect.

### Nuke the working tree

Press `shift+d` to bring up the reset options menu, then select 'nuke' to clear anything `git status` shows — including dirty submodules.

### Amend an old commit

Pressing `shift+a` on any commit amends that commit with the currently staged changes (running an interactive rebase in the background).

### Filter

Filter any view with `/` and press `enter` to drill in.

### Invoke a custom command

Malasgit has a flexible [custom command system](docs/Custom_Command_Keybindings.md) for adding your own actions.

### Worktrees

Press `w` in the branches view to create a worktree from the selected branch and switch to it, so you can keep multiple branches going at once without stashing or WIP commits.

### Rebase magic (custom patches)

Build a custom patch from an old commit, then remove the patch from the commit, split out a new commit, apply the patch in reverse to the index, and more. Hit `<enter>` on a commit to view its files, `<enter>` on a file to focus the patch, `<space>` to add lines to the custom patch, then `ctrl+p` for patch options. See the [Rebase magic tutorial](https://youtu.be/4XaToVut_hs) for a walkthrough.

### Rebase from marked base commit

Press `shift+b` on a commit to mark it as the base, then `r` on the target branch to rebase only the commits from your feature branch onto it.

### Undo

Press `z` to undo the last action and `shift+z` to redo. Undo uses the reflog, so it covers commits and branches but not the working tree or stash. See [docs/Undoing.md](docs/Undoing.md) for details.

### Commit graph

In an enlarged window (cycle screen modes with `+` and `_`), the commit graph is shown. Colours correspond to commit authors; the selected commit's parents are highlighted as you navigate.

### Compare two commits

Press `shift+w` on a commit or ref to mark it; the next commit you select will be diffed against it. Press `shift+w` again to access the diff menu (reverse direction, exit diff mode), or `<escape>` to exit.

### Show GitHub pull requests

In the branches panel, branches with an associated GitHub PR show a GitHub icon whose colour reflects PR state (open, merged, etc.). Press `shift-G` to open the PR in the browser. Works out of the box for github.com once the [`gh`](https://cli.github.com/) CLI is installed and `gh auth login` has been run; for GitHub Enterprise add a [`services` entry](docs/Config.md#custom-pull-request-urls) with the `github` provider.

## Installation

### Binary Releases

Download a binary for Windows, macOS (10.12+), or Linux from the [releases page](https://github.com/itokun99/malasgit/releases).

### Homebrew

```sh
brew install malasgit
```

### Go

```sh
go install github.com/itokun99/malasgit@latest
```

If `malasgit` is not found afterwards, add `~/go/bin` (macOS/Linux) or `%HOME%\go\bin` (Windows) to your `PATH`.

### Build from Source

```sh
git clone https://github.com/itokun99/malasgit.git
cd malasgit
go install
```

## Usage

Run `malasgit` inside a git repository:

```sh
$ malasgit
```

Optionally add an alias so you can call it as `lg`:

```sh
echo "alias lg='malasgit'" >> ~/.zshrc
```

## Configuration

See [docs/Config.md](docs/Config.md) for the full configuration reference and [docs/keybindings](docs/keybindings) for the keybinding cheatsheet.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). For contributor discussion, join the [discord channel](https://discord.gg/ehwFt2t4wt).

### Debugging Locally

Run `malasgit --debug` in one terminal and `malasgit --logs` in another to view the program and its log output side by side.

## FAQ

### What do the commit colors represent?

- Green: the commit is included in the master branch.
- Yellow: the commit is not included in the master branch.
- Red: the commit has not been pushed to the upstream branch.

## Alternatives

If malasgit doesn't fit your needs, consider:

- [GitUI](https://github.com/Extrawurst/gitui)
- [tig](https://github.com/jonas/tig)
- [GitArbor TUI](https://github.com/cadamsdev/gitarbor-tui)