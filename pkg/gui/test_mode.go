package gui

import (
	"log"
	"os"
	"time"

	"github.com/itokun99/malasgit/pkg/gocui"
	"github.com/itokun99/malasgit/pkg/gui/popup"
	"github.com/itokun99/malasgit/pkg/gui/types"
	"github.com/itokun99/malasgit/pkg/integration/components"
	"github.com/itokun99/malasgit/pkg/utils"
)

type IntegrationTest interface {
	Run(*GuiDriver)
}

func (gui *Gui) handleTestMode() {
	test := gui.integrationTest
	if os.Getenv(components.SANDBOX_ENV_VAR) == "true" {
		return
	}

	if test != nil {
		waitUntilIdle := func() {
			gui.c.GocuiGui().WaitUntilIdle()
		}

		go func() {
			waitUntilIdle()

			toastChan := make(chan string, 100)
			gui.PopupHandler.(*popup.PopupHandler).SetToastFunc(
				func(message string, kind types.ToastKind) { toastChan <- message })

			test.Run(&GuiDriver{gui: gui, toastChan: toastChan, headless: Headless()})

			gui.g.Update(func(*gocui.Gui) error {
				return gocui.ErrQuit
			})

			// Wait for the event loop to actually exit.
			<-gui.g.LoopExited()
		}()

		if os.Getenv(components.WAIT_FOR_DEBUGGER_ENV_VAR) == "" {
			go utils.Safe(func() {
				time.Sleep(time.Second * 40)
				log.Fatal("40 seconds is up, malasgit recording took too long to complete")
			})
		}
	}
}

func Headless() bool {
	return os.Getenv("LAZYGIT_HEADLESS") != ""
}
