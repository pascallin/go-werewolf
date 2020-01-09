package command

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
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
	port := ctx.Int("port")
	people := ctx.Int("people")
	name := ctx.String("name")
	fmt.Println("开放端口:", port)
	fmt.Println("游戏名称:", name)
	fmt.Println("参与人数:", people)
	// TODO: create game
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
		// TODO: join game
		return nil
	},
}

var exitCommand = &cli.Command{
	Name:  "exit",
	Usage: "退出",
	Action: func(ctx *cli.Context) error {
		// TODO: exit
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
