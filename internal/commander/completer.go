package commander

import "github.com/chzyer/readline"

var completer = readline.NewPrefixCompleter(
	// Common command
	readline.PcItem("help"),
	readline.PcItem("exit"),

	// Create
	readline.PcItem("create"),
	readline.PcItem("join"),

	// Game command
	readline.PcItem("say"),
	readline.PcItem("vote"),
	readline.PcItem("kill"),
	//readline.PcItem("exit"),
)