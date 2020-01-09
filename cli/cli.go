package cli

import (
	"github.com/urfave/cli/v2"
	"os"
)

func Cli() {
	app := cli.NewApp()
	app.Name = "werewolf-cli"
	app.Usage = "狼人杀命令行版"
	app.Version = "0.0.1"
	app.HelpName = "werewolf-cli"
	app.HideHelp = true
	app.HideVersion = true
	cli.HelpFlag = &cli.BoolFlag{
		Name: "help,h",
		//Aliases: []string{"h"},
		Usage: "展示帮助信息",
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "显示版本信息",
		Hidden:  true,
	}
	app.Commands = []*cli.Command{
		&create,
	}

	app.Run(os.Args)

}
