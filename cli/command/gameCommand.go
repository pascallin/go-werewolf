package command

import (
	"fmt"
	"github.com/pascallin/go-wolvesgame/context"
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

var sayCommand = &cli.Command{
	Name:    "say",
	Aliases: []string{"s"},
	Usage:   "发言",
	Action: func(ctx *cli.Context) error {
		msg := ctx.Args().Get(0)
		if len(msg) == 0 {
			fmt.Println("Error:发言内容不能为空！")
			return nil
		}
		// TODO: send msg
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		fmt.Println("你说：", msg)
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
