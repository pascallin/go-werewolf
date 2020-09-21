package command

import "github.com/urfave/cli/v2"

func GetCommands() []*cli.Command {
	var commands = []*cli.Command{
		helpCommand,
		exitCommand,
		gameCommands,
		playerCommands,
	}
	return commands
}