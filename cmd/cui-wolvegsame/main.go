package main

import (
	"log"

	"github.com/rivo/tview"

	"github.com/pascallin/go-wolvesgame/internal/commander"
	"github.com/pascallin/go-wolvesgame/internal/cui"
)

func main() {
	printPanel := cui.NewCommandResultPanel()
	err, terminal := commander.CreateCliApp(printPanel)
	if err != nil {
		log.Fatal(err)
	}
	commandPanel := cui.NewCommandInput(terminal)
	flexPanel := cui.NewFlexLayout(commandPanel, printPanel, cui.NewMessagePanel())
	view := tview.NewApplication()
	view.SetRoot(flexPanel, true).EnableMouse(true)

	if err := view.Run(); err != nil {
		panic(err)
	}
}