package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var sayCommand = &cli.Command{
	Name:    "say",
	Aliases: []string{"s"},
	Usage:   "发言",
	Action: func(ctx *cli.Context) error {
		//if ctx.Args().Len() == 0 {
		//	ctx.App.Writer.Write([]byte("Error:发言内容不能为空！"))
		//	return nil
		//}
		//msg := ctx.Args().Slice()
		//c := werewolf.GetApp()
		//client := c.GetTcpClient()
		//go client.Send("[" +c.GetUser().Nickname + " said]" + strings.Join(msg, " "))
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
			ctx.App.Writer.Write([]byte("Error:需要指定名称"))
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
			ctx.App.Writer.Write([]byte("Error:需要指定名称"))
			return nil
		}
		// TODO: kill vote someone
		fmt.Println("你选择杀：", name)
		return nil
	},
}

var playerCommands = &cli.Command{
	Name:    		"player",
	Aliases:	 	[]string{"i"},
	Usage:   		"玩家操作",
	Before: func(c *cli.Context) error {
		fmt.Fprintf(c.App.Writer, "\n")
		return nil
	},
	Subcommands: 	[]*cli.Command{
		sayCommand,
		voteCommand,
		killCommand,
	},
}