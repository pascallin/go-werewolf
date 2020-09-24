package commander

import (
	"context"
	"io"

	"github.com/urfave/cli/v2"

	"github.com/pascallin/go-wolvesgame/internal/commander/commands"
	"github.com/pascallin/go-wolvesgame/internal/werewolf"
)

func CreateCliApp(app *werewolf.App, writer io.Writer) (error, *cli.App) {
	cli.AppHelpTemplate = `{{if .Commands}}{{range .Commands}}{{if not .HideHelp}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}`
	terminal := cli.App{
		Before: func(c *cli.Context) error {
			c.Context = context.WithValue(c.Context, "gameApp", app)
			return nil
		},
	}
	terminal.Name = "wolves-game"
	terminal.Commands = commands.GetCommands()
	terminal.Usage = "狼人杀命令行版"
	terminal.HelpName = "wolves-game"
	terminal.Version = "0.0.1"
	terminal.Action = func(c *cli.Context) error {
		c.App.Writer.Write([]byte("Command not found. Type 'help' for a list of command. \n"))
		return nil
	}
	// NOTE: define terminal writer
	if writer != nil {
		terminal.Writer = writer
	}
	return nil, &terminal
}
