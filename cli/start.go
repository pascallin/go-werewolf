package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var flags = []cli.Flag{
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


var action = func(ctx *cli.Context) error {
	port := ctx.Int("port")
	people := ctx.Int("people")
	name := ctx.String("name")
	fmt.Println("开放端口:", port)
	fmt.Println("游戏名称:", name)
	fmt.Println("参与人数:", people)
	// 创建交互式终端
	Console()
	return nil
}

var create = cli.Command{
	Name:    "start",
	Aliases: []string{"c"},
	Usage:   "创建一个游戏",
	Flags:   flags,
	Action:  action,
}
