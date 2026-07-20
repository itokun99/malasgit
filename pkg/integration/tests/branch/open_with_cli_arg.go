package branch

import (
	"github.com/itokun99/malasgit/pkg/config"
	. "github.com/itokun99/malasgit/pkg/integration/components"
)

var OpenWithCliArg = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Open straight to branches panel using a CLI arg",
	ExtraCmdArgs: []string{"branch"},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Branches().IsFocused()
	},
})
