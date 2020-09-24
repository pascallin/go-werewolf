package commands

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

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
	Before: func(c *cli.Context) error {
		fmt.Fprintf(c.App.Writer, "\n")
		return nil
	},
	Action: func(c *cli.Context) error {
		args := c.Args()
		if args.Present() {
			err := cli.ShowCommandHelp(c, args.First())
			if err != nil {
				fmt.Println(err)
			}
			return nil
		}
		cli.ShowAppHelp(c)
		return nil
	},
}