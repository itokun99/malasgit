package custom_commands

import (
	"github.com/itokun99/malasgit/pkg/config"
	. "github.com/itokun99/malasgit/pkg/integration/components"
)

var GlobalContext = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Ensure global context works",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupRepo: func(shell *Shell) {
		shell.EmptyCommit("my change")
	},
	SetupConfig: func(cfg *config.AppConfig) {
		cfg.GetUserConfig().CustomCommands = []config.CustomCommand{
			{
				Key:     config.Keybinding{"X"},
				Context: "global",
				Command: "touch myfile",
			},
		}
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		// commits
		t.Views().Commits().
			Focus().
			Press(config.Keybinding{"X"})

		t.Views().Files().
			Focus().
			Lines(Contains("myfile"))

		t.Shell().DeleteFile("myfile")
		t.GlobalPress(keys.Files.RefreshFiles)

		// branches
		t.Views().Branches().
			Focus().
			Press(config.Keybinding{"X"})

		t.Views().Files().
			Focus().
			Lines(Contains("myfile"))

		t.Shell().DeleteFile("myfile")
		t.GlobalPress(keys.Files.RefreshFiles)

		// files
		t.Views().Files().
			Focus().
			Press(config.Keybinding{"X"})

		t.Views().Files().
			Focus().
			Lines(Contains("myfile"))

		t.Shell().DeleteFile("myfile")
	},
})
