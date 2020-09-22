package cli

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pascallin/go-wolvesgame/internal/app"
	"github.com/pascallin/go-wolvesgame/internal/cli/command"
)

func createTerminal() error {
	cli.AppHelpTemplate = `{{if .Commands}}{{range .Commands}}{{if not .HideHelp}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}`
	terminal := cli.NewApp()
	terminal.Name = "wolves-game"
	terminal.Commands = command.GetCommands()
	terminal.Usage = "狼人杀命令行版"
	terminal.HelpName = "wolves-game"
	terminal.Version = "0.0.1"
	terminal.Action = func(c *cli.Context) error {
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
			terminal.Run([]string{"app", "exit"})
		}
		line = strings.TrimSpace(line)
		err = terminal.Run(strings.Fields("cmd " + line))
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func CreateCliApp(writer io.Writer) (error, *cli.App) {
	cli.AppHelpTemplate = `{{if .Commands}}{{range .Commands}}{{if not .HideHelp}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}`
	terminal := cli.NewApp()
	terminal.Name = "wolves-game"
	terminal.Commands = command.GetCommands()
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

func Start() {
	console := cli.NewApp()
	console.Usage = "狼人杀命令行版"
	console.Version = "0.0.1"
	console.HelpName = "wolves-game"
	console.Name = "wolves-game"

	console.Flags = []cli.Flag {
		&cli.StringFlag{
			Name: "username, u",
			Usage: "player nickname",
			Required: true,
		},
	}

	console.Action = func(c *cli.Context) error {
		app.GetApp().SetUser(app.NewUser(c.String("username")))
		err := createTerminal()
		return err
	}

	err := console.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
