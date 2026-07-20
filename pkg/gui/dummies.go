package gui

import (
	"github.com/itokun99/malasgit/pkg/commands/git_commands"
	"github.com/itokun99/malasgit/pkg/commands/oscommands"
	"github.com/itokun99/malasgit/pkg/common"
	"github.com/itokun99/malasgit/pkg/config"
	"github.com/itokun99/malasgit/pkg/updates"
)

func NewDummyUpdater() *updates.Updater {
	newAppConfig := config.NewDummyAppConfig()
	dummyUpdater, _ := updates.NewUpdater(common.NewDummyCommon(), newAppConfig, oscommands.NewDummyOSCommand())
	return dummyUpdater
}

// NewDummyGui creates a new dummy GUI for testing
func NewDummyGui() *Gui {
	newAppConfig := config.NewDummyAppConfig()
	dummyGui, _ := NewGui(common.NewDummyCommon(), newAppConfig, &git_commands.GitVersion{Major: 2, Minor: 0, Patch: 0}, NewDummyUpdater(), false, "", nil)
	return dummyGui
}
