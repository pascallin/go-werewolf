package terminalview

import (
	"github.com/rivo/tview"
)

func GetTerminalView() (*tview.Application, *tview.TextView){
	app := tview.NewApplication()

	flexPanel, printPanel := NewFlexLayout(app)
	return app.SetRoot(flexPanel, true).EnableMouse(true), printPanel
}