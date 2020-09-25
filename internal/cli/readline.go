package cli

import (
	"fmt"
	"io"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/chzyer/readline"
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func CreateReadline(username string) *readline.Instance {
	rl, _ := readline.NewEx(&readline.Config{
		Prompt: "\033[31mWerwolf("+ username + ")»\033[0m ",
		//HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:        completer,
		InterruptPrompt:     "^C",
		EOFPrompt:           "exit",
		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	return rl
}

func ListenReadline(terminal *cli.App, rl *readline.Instance) {
	for {
		line, err := rl.Readline()
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
}