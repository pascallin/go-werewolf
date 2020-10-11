package main

import (
	"github.com/pascallin/go-wolvesgame/internal/cui"
	"github.com/rivo/tview"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"

	clipkg "github.com/pascallin/go-wolvesgame/internal/cli"
	"github.com/pascallin/go-wolvesgame/internal/commander"
	"github.com/pascallin/go-wolvesgame/internal/werewolf"
)

func main() {
	console := cli.NewApp()
	console.Usage = "狼人杀命令行版"
	console.Version = "0.0.1"
	console.HelpName = "wolves-game"
	console.Name = "wolves-game"

	console.Flags = []cli.Flag {
		&cli.StringFlag{
			Name: "username, u",
			Usage: "player nickname",
			Required: true,
		},
		&cli.BoolFlag{
			Name: "ui",
			Usage: "play with CUI",
			Value: false,
			Required: false,
		},
	}

	console.Action = func(c *cli.Context) error {
		withUI, err := strconv.ParseBool(c.String("ui"))
		if err != nil {
			return err
		}
		if withUI {
			initCUIApp(c.String("username"))
		} else {
			initCliApp(c.String("username"))
		}
		return nil
	}

	err := console.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func initCliApp(username string) {
	// init readline
	rl := clipkg.CreateReadline(username)

	// init game and terminal
	gameApp := werewolf.New(username, rl)
	err, terminal := commander.New(gameApp, rl)
	if err != nil {
		panic(err)
	}

	clipkg.ListenReadline(terminal, rl)
}

func initCUIApp(username string) {
	// init output panels
	printPanel := cui.NewCommandResultPanel()
	messagePanel := cui.NewMessagePanel()

	// init game and terminal
	gameApp := werewolf.New(username, messagePanel)
	err, terminal := commander.New(gameApp, printPanel)
	if err != nil {
		panic(err)
	}

	// init layout
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

