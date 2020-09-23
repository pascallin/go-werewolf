package cui

import (
	"github.com/rivo/tview"
)

func GetTerminalView(commandChan chan string) (*tview.Application, *tview.TextView){
	app := tview.NewApplication()

	flexPanel, printPanel := NewFlexLayout(app, commandChan)
	return app.SetRoot(flexPanel, true).EnableMouse(true), printPanel
}