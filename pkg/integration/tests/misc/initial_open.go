package misc

import (
	"github.com/itokun99/malasgit/pkg/config"
	. "github.com/itokun99/malasgit/pkg/integration/components"
)

var InitialOpen = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Confirms a popup appears on first opening Lazygit",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig: func(config *config.AppConfig) {
		config.GetUserConfig().DisableStartupPopups = false
	},
	SetupRepo: func(shell *Shell) {},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.ExpectPopup().Confirmation().
			Title(Equals("")).
			Content(Contains("Thanks for using malasgit!")).
			Confirm()

		t.Views().Files().IsFocused()
	},
})
