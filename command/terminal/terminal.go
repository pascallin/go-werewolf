package terminal

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/urfave/cli/v2"
	"io"
	"strings"
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
	console := cli.NewApp()
	console.Commands = commands
	console.Action = func(c *cli.Context) error {
		fmt.Println("Command not found. Type 'help' for a list of commands.")
		return nil
	}
	l, _ := readline.NewEx(&readline.Config{
		Prompt: "\033[31mWerwolf»\033[0m ",
		//HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    completer,
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