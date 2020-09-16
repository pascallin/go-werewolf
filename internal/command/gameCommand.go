package command

import (
	"fmt"
	"strings"

	"github.com/pascallin/go-wolvesgame/context"
	"github.com/pascallin/go-wolvesgame/internal/transport/tcp"
	"github.com/urfave/cli/v2"
)

var statusCommand = &cli.Command{
	Name:    "status",
	Aliases: []string{"v"},
	Usage:   "显示游戏状态",
	Action: func(ctx *cli.Context) error {
		context.GetContext().GetGame().PrintGameStatus()
		return nil
	},
}

var startCommand = &cli.Command{
	Name:    "start",
	Aliases: []string{"v"},
	Usage:   "开始游戏",
	Action: func(ctx *cli.Context) error {
		// add game
		game := context.GetContext().GetGame()
		game.GameStart()

		// run socket server
		go tcp.NewServer()

		// create socket client
		c := context.GetContext()
		c.SetTcpClient(tcp.NewClient())

		return nil
	},
}

var sayCommand = &cli.Command{
	Name:    "say",
	Aliases: []string{"s"},
	Usage:   "发言",
	Action: func(ctx *cli.Context) error {
		if ctx.Args().Len() == 0 {
			fmt.Println("Error:发言内容不能为空！")
			return nil
		}
		msg := ctx.Args().Slice()
		c := context.GetContext()
		client := c.GetTcpClient()
		go client.Send(strings.Join(msg, " "))
		return nil
	},
}

var voteCommand = &cli.Command{
	Name:    "vote",
	Aliases: []string{"v"},
	Usage:   "投票",
	Action: func(ctx *cli.Context) error {
		name := ctx.Args().Get(0)
		if len(name) == 0 {
			fmt.Println("Error:需要指定名称")
			return nil
		}
		// TODO: vote someone
		fmt.Println("你投票给：", name)
		return nil
	},
}

var killCommand = &cli.Command{
	Name:    "kill",
	Aliases: []string{"k"},
	Usage:   "杀人",
	Action: func(ctx *cli.Context) error {
		name := ctx.Args().Get(0)
		if len(name) == 0 {
			fmt.Println("Error:需要指定名称")
			return nil
		}
		// TODO: kill vote someone
		fmt.Println("你选择杀：", name)
		return nil
	},
}
