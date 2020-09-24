package commander

import (
	"github.com/urfave/cli/v2"
	"io"

	"github.com/pascallin/go-wolvesgame/internal/commander/commands"
)

func CreateCliApp(writer io.Writer) (error, *cli.App) {
	cli.AppHelpTemplate = `{{if .Commands}}{{range .Commands}}{{if not .HideHelp}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}`
	terminal := cli.NewApp()
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
	return nil, terminal
}
