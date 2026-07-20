package conflicts

import (
	"github.com/itokun99/malasgit/pkg/config"
	. "github.com/itokun99/malasgit/pkg/integration/components"
	"github.com/itokun99/malasgit/pkg/integration/tests/shared"
)

var ContinuePromptDismissedWhenResolvedExternally = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "When the prompt to continue a merge is showing and the merge is then continued outside malasgit, dismiss the prompt",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shared.CreateMergeConflictFile(shell)
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Common().PretendMergeOrRebaseStartedInLazygit()

		// Resolve the conflict and refresh so malasgit prompts us to continue.
		t.Views().Files().
			IsFocused().
			Lines(
				Contains("UU file").IsSelected(),
			).
			Tap(func() {
				t.Shell().UpdateFile("file", "resolved content")
			}).
			Press(keys.Universal.Refresh)

		t.ExpectPopup().Confirmation().
			Title(Equals("Continue")).
			Content(Contains("All merge conflicts resolved. Continue the merge?"))

		// While the prompt is up, the merge is continued outside malasgit (e.g. by
		// a coding agent).
		t.Shell().ContinueMerge()

		// Simulate malasgit noticing the change (as it would on its next refresh or
		// when the window regains focus); the stale prompt is dismissed.
		t.FocusIn()

		t.Views().Files().
			IsFocused().
			IsEmpty()

		t.Views().Information().Content(DoesNotContain("Merging"))
	},
})
