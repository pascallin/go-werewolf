package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

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
	}

	console.Action = func(c *cli.Context) error {
		gameApp := werewolf.New(c.String("username"), os.Stdout)
		rl := commander.CreateReadline(c.String("username"))
		err, terminal := commander.CreateCliApp(gameApp, rl)
		if err != nil {
			return err
		}
		commander.ListenReadline(gameApp, terminal, rl)
		return nil
	}

	err := console.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
