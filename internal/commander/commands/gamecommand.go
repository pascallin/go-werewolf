package commands

import (
	"fmt"
	"github.com/pascallin/go-wolvesgame/internal/werewolf"
	"github.com/urfave/cli/v2"

	"github.com/pascallin/go-wolvesgame/internal/game"
)

var createCommand = &cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Usage:   "创建游戏",
	Flags:    []cli.Flag{
		&cli.IntFlag{
			Name:     "port",
			Value:    8080,
			Usage:    "端口",
			Required: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		gameApp := ctx.Context.Value("gameApp").(*werewolf.App)
		if gameApp.Game != nil {
			ctx.App.Writer.Write([]byte("Game has been created"))
			return nil
		}
		gameApp.Game = game.New()
		gameApp.User.JoinGame()
		go gameApp.TCPServer.Listen(ctx.String("port"))
		go gameApp.TCPClient.Dia("localhost:" + ctx.String("port"))
		return nil
	},
}

var joinCommand = &cli.Command{
	Name:    "join",
	Aliases: []string{"j"},
	Usage:   "加入游戏",
	Flags:    []cli.Flag{
		&cli.StringFlag{
			Name:     "url",
			Value:    "localhost:8080",
			Usage:    "游戏服务器IP+端口",
			Required: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		gameApp := ctx.Context.Value("gameApp").(*werewolf.App)
		if gameApp.Game != nil {
			ctx.App.Writer.Write([]byte("You have been join a game"))
			return nil
		}
		go gameApp.TCPClient.Dia(ctx.String("url"))
		return nil
	},
}

var statusCommand = &cli.Command{
	Name:    "status",
	Usage:   "显示游戏状态",
	Action: func(ctx *cli.Context) error {
		gameApp := ctx.Context.Value("gameApp").(*werewolf.App)
		if gameApp.Game == nil {
			ctx.App.Writer.Write([]byte("No game exist"))
			return nil
		}
		ctx.App.Writer.Write([]byte(gameApp.Game.GameStatusJSON()))
		return nil
	},
}

var startCommand = &cli.Command{
	Name:    "start",
	Usage:   "开始游戏",
	Action: func(ctx *cli.Context) error {
		gameApp := ctx.Context.Value("gameApp").(*werewolf.App)
		if gameApp.Game == nil {
			ctx.App.Writer.Write([]byte("You have not join any game yet"))
			return nil
		}
		game.StartGame(gameApp.Game)
		return nil
	},
}

var gameCommands = &cli.Command{
	Name:    		"game",
	Aliases:	 	[]string{"g"},
	Usage:   		"游戏操作",
	Before: func(c *cli.Context) error {
		fmt.Fprintf(c.App.Writer, "\n")
		return nil
	},
	Subcommands: 	[]*cli.Command{
		statusCommand,
		startCommand,
		createCommand,
		joinCommand,
	},
}