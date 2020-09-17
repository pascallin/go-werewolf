package console

import (
	"github.com/chzyer/readline"
	"github.com/pascallin/go-wolvesgame/internal/app"
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
	user := app.GetApp().GetUser()
	rl, _ := readline.NewEx(&readline.Config{
		Prompt: "\033[31mWerwolf("+ user.Nickname + ")»\033[0m ",
		//HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:        completer,
		InterruptPrompt:     "^C",
		EOFPrompt:           "exit",
		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	return rl
}
