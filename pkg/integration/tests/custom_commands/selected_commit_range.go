package custom_commands

import (
	"github.com/itokun99/malasgit/pkg/config"
	. "github.com/itokun99/malasgit/pkg/integration/components"
)

var SelectedCommitRange = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Use the {{ .SelectedCommitRange }} template variable",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupRepo: func(shell *Shell) {
		shell.CreateNCommits(3)
	},
	SetupConfig: func(cfg *config.AppConfig) {
		cfg.GetUserConfig().CustomCommands = []config.CustomCommand{
			{
				Key:     config.Keybinding{"X"},
				Context: "global",
				Command: `git log --format="%s" {{.SelectedCommitRange.From}}^..{{.SelectedCommitRange.To}} > file.txt`,
			},
		}
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Commits().Focus().
			Lines(
				Contains("commit-03").IsSelected(),
				Contains("commit-02"),
				Contains("commit-01"),
			)

		t.GlobalPress(config.Keybinding{"X"})
		t.FileSystem().FileContent("file.txt", Equals("commit-03\n"))

		t.Views().Commits().Focus().
			Press(keys.Universal.RangeSelectDown)

		t.GlobalPress(config.Keybinding{"X"})
		t.FileSystem().FileContent("file.txt", Equals("commit-03\ncommit-02\n"))
	},
})
