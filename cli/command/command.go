package command

import (
	"github.com/urfave/cli/v2"
	"github.com/wsxiaoys/terminal/color"
)

var commands = []*cli.Command{
	{
		Name:    "say",
		Aliases: []string{"s"},
		Usage:   "发言",
		Action: func(ctx *cli.Context) error {
			msg := ctx.Args().Get(0)
			if len(msg) == 0 {
				color.Println("Error:发言内容不能为空！")
				return nil
			}
			// TODO: send msg
			//if err != nil {
			//	fmt.Println(err)
			//	os.Exit(1)
			//}
			color.Println("你说：", msg)
			return nil
		},
	},
}