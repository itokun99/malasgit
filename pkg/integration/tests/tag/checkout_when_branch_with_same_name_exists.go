package tag

import (
	"github.com/itokun99/malasgit/pkg/config"
	. "github.com/itokun99/malasgit/pkg/integration/components"
)

var CheckoutWhenBranchWithSameNameExists = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Checkout a tag when there's a branch with the same name",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.EmptyCommit("one")
		shell.NewBranch("tag")
		shell.Checkout("master")
		shell.EmptyCommit("two")
		shell.CreateLightweightTag("tag", "HEAD")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Tags().
			Focus().
			Lines(
				Contains("tag").IsSelected(),
			).
			PressPrimaryAction() // checkout tag

		t.Views().Branches().IsFocused().Lines(
			Contains("HEAD detached at tag").IsSelected(),
			Contains("master"),
			Contains("tag"),
		)
	},
})
