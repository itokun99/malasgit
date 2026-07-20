package custom_commands

import (
	"github.com/itokun99/malasgit/pkg/config"
	. "github.com/itokun99/malasgit/pkg/integration/components"
	"github.com/itokun99/malasgit/pkg/integration/tests/shared"
)

var CheckForConflicts = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Run a command and check for conflicts after",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupRepo: func(shell *Shell) {
		shared.MergeConflictsSetup(shell)
	},
	SetupConfig: func(cfg *config.AppConfig) {
		cfg.GetUserConfig().CustomCommands = []config.CustomCommand{
			{
				Key:     config.Keybinding{"m"},
				Context: "localBranches",
				Command: "git merge {{ .SelectedLocalBranch.Name | quote }}",
				After: &config.CustomCommandAfterHook{
					CheckForConflicts: true,
				},
			},
		}

		cfg.GetUserConfig().Git.LocalBranchSortOrder = "recency"
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Branches().
			Focus().
			TopLines(
				Contains("first-change-branch"),
				Contains("second-change-branch"),
			).
			NavigateToLine(Contains("second-change-branch")).
			Press(config.Keybinding{"m"})

		t.Common().AcknowledgeConflicts()
	},
})
