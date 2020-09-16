package command

import (
	"fmt"
	"os"

	"github.com/pascallin/go-wolvesgame/context"
	"github.com/pascallin/go-wolvesgame/internal/game"
	"github.com/pascallin/go-wolvesgame/internal/transport/tcp"
	"github.com/urfave/cli/v2"
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

var createCommand = &cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Usage:   "创建游戏",
	Flags:   createFlags,
	Action: func(ctx *cli.Context) error {
		c := context.GetContext()
		game := game.CreateGame()
		c.SetGame(game)
		c.SetTcpServer(tcp.NewServer())
		c.SetTcpClient(tcp.NewClient())
		return nil
	},
}

var joinCommand = &cli.Command{
	Name:    "join",
	Aliases: []string{"j"},
	Usage:   "加入游戏",
	Action: func(ctx *cli.Context) error {
		// TODO: 判断是否在游戏中
		c := context.GetContext()
		c.SetTcpClient(tcp.NewClient())
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
