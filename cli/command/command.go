package command

import (
	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	// Cli Command
	createCommand,
	joinCommand,
	// Game Command
	sayCommand,
	voteCommand,
	killCommand,

	helpCommand,
	exitCommand,
}