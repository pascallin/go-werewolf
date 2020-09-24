package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/pascallin/go-wolvesgame/internal/app"
	"github.com/pascallin/go-wolvesgame/internal/commander"
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
	}

	console.Action = func(c *cli.Context) error {
		app.GetApp().SetUser(app.NewUser(c.String("username")))
		err, terminal := commander.CreateCliApp(nil)
		if err != nil {
			return err
		}
		commander.ListenReadline(terminal)
		return nil
	}

	err := console.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
