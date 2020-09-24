package main

import (
	"github.com/pascallin/go-wolvesgame/internal/werewolf"
	"log"

	"github.com/rivo/tview"

	"github.com/pascallin/go-wolvesgame/internal/commander"
	"github.com/pascallin/go-wolvesgame/internal/cui"
)

func main() {
	// initialize panels
	printPanel := cui.NewCommandResultPanel()
	messagePanel := cui.NewMessagePanel()
	// TODO: get username
	gameApp := werewolf.New("Pascal", messagePanel)
	err, terminal := commander.CreateCliApp(gameApp, printPanel)
	if err != nil {
		log.Fatal(err)
	}
	commandPanel := cui.NewCommandInput(terminal)
	gamePanel := cui.NewGamePanel()
	flexPanel := cui.NewFlexLayout(commandPanel, printPanel, messagePanel, gamePanel.Table)

	// start tview app
	view := tview.NewApplication()
	view.SetRoot(flexPanel, true).EnableMouse(true)

	if err := view.Run(); err != nil {
		panic(err)
	}
}