package main

import (
	cliinterface "github.com/pascallin/go-wolvesgame/internal/cli"
	"github.com/pascallin/go-wolvesgame/internal/cui"
	"github.com/rivo/tview"
	"log"
)

func main() {
	printPanel := cui.NewCommandResultPanel()
	err, terminal := cliinterface.CreateCliApp(printPanel)
	if err != nil {
		log.Fatal(err)
	}
	commandPanel := cui.NewCommandInput(terminal)
	flexPanel := cui.NewFlexLayout(commandPanel, printPanel, cui.NewMessagePanel())
	app := tview.NewApplication()
	app.SetRoot(flexPanel, true).EnableMouse(true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}