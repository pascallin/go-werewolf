package console

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/pascallin/go-wolvesgame/internal/command"
	"github.com/urfave/cli/v2"
)

var (
	console cli.App
)

func terminal() error {
	cli.AppHelpTemplate = `{{if .Commands}}{{range .Commands}}{{if not .HideHelp}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}`
	console := cli.NewApp()
	console.Name = "wolvesgame"
	console.Commands = command.Commands
	console.Usage = "狼人杀命令行版"
	console.HelpName = "wolvesgame"
	console.Version = "0.0.1"
	console.Action = func(c *cli.Context) error {
		fmt.Println("Command not found. Type 'help' for a list of command.")
		return nil
	}
	l := getReadline()

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			console.Run([]string{"app", "exit"})
		}
		line = strings.TrimSpace(line)
		err = console.Run(strings.Fields("cmd " + line))
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func Console() {
	app := cli.NewApp()
	app.Usage = "狼人杀命令行版"
	app.Version = "0.0.1"
	app.HelpName = "wolvesgame"
	app.Name = "wolvesgame"
	app.Action = func(c *cli.Context) error {
		err := terminal()
		return err
	}
	app.Run(os.Args)
}
