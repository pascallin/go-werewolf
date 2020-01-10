package command

import (
	"fmt"
	"github.com/pascallin/go-wolvesgame/context"
	"github.com/pascallin/go-wolvesgame/game"
	"github.com/urfave/cli/v2"
	"os"
	"github.com/pascallin/go-wolvesgame/transport/tcpsocket"
)

var createFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "name",
		Value:    "狼人杀",
		Usage:    "游戏名称",
		Required: false,
	},
	&cli.IntFlag{
		Name:     "people",
		Aliases:  []string{"p"},
		Value:    9,
		Usage:    "参与人数",
		Required: false,
	},
	&cli.IntFlag{
		Name:     "port",
		Value:    8000,
		Usage:    "端口",
		Required: false,
	},
}

var createAction = func(ctx *cli.Context) error {
	c := context.GetContext()
	game := game.CreateGame()
	c.SetGame(game)
	go tcpsocket.NewServer()
	return nil
}

var createCommand = &cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Usage:   "创建游戏",
	Flags:   createFlags,
	Action:  createAction,
}

var joinCommand = &cli.Command{
	Name:    "join",
	Aliases: []string{"j"},
	Usage:   "加入游戏",
	Action: func(ctx *cli.Context) error {
		c := context.GetContext()
		c.SetSocket(tcpsocket.NewClient())
		return nil
	},
}

var exitCommand = &cli.Command{
	Name:  "exit",
	Usage: "退出",
	Action: func(ctx *cli.Context) error {
		fmt.Println("退出游戏")
		os.Exit(0)
		return nil
	},
}

var helpCommand = &cli.Command{
	Name:      "help",
	Aliases:   []string{"h"},
	Usage:     "展示帮助信息",
	ArgsUsage: "[command]",
	Action: func(c *cli.Context) error {
		args := c.Args()
		if args.Present() {
			err := cli.ShowCommandHelp(c, args.First())
			if err != nil {
				fmt.Println(err)
			}
			return nil
		}

		_ = cli.ShowAppHelp(c)
		return nil
	},
}
