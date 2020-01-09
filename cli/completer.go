package cli

import "github.com/chzyer/readline"

var completer = readline.NewPrefixCompleter(
	readline.PcItem("say"),
	readline.PcItem("help"),
	readline.PcItem("create"),
	// TODO: to be continue
	readline.PcItem("exist"),
)