package cli

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/pascallin/go-wolvesgame/cli/command"
	"github.com/urfave/cli/v2"
	"io"
	"strings"

	"github.com/pascallin/go-wolvesgame/game"
	)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func Console() {
	game := game.CreateGame()
	game.PrintGameStatus()

	console := cli.NewApp()
	console.Commands = command.commands
	console.Action = func(c *cli.Context) error {
		fmt.Println("Command not found. Type 'help' for a list of command.")
		return nil
	}
	l, _ := readline.NewEx(&readline.Config{
		Prompt: "\033[31mWerwolf»\033[0m ",
		//HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    cli2.completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		console.Run(strings.Fields("cmd " + line))
	}
}