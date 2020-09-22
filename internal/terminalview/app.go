package terminalview

import "github.com/rivo/tview"

func GetTerminalView() *tview.Application {
	app := tview.NewApplication()
	return app.SetRoot(NewFlexLayout(app), true).EnableMouse(true)
}