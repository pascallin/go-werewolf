package cli

import (
	"github.com/chzyer/readline"
	"github.com/pascallin/go-wolvesgame/cli/command"
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func getReadline() *readline.Instance {
	rl, _ := readline.NewEx(&readline.Config{
		Prompt: "\033[31mWerwolf»\033[0m ",
		//HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:        command.Completer,
		InterruptPrompt:     "^C",
		EOFPrompt:           "exit",
		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	return rl
}
