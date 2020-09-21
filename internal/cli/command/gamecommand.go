package command

import (
	"github.com/urfave/cli/v2"

	"github.com/pascallin/go-wolvesgame/internal/app"
	"github.com/pascallin/go-wolvesgame/internal/game"
	"github.com/pascallin/go-wolvesgame/pkg/tcp"
)

var createFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "name",
		Value:    "狼人杀",
		Usage:    "游戏名称",
		Required: false,
	},
	&cli.IntFlag{
		Name:     "people",
		Aliases:  []string{"p"},
		Value:    12,
		Usage:    "参与人数",
		Required: false,
	},
	&cli.IntFlag{
		Name:     "port",
		Value:    8000,
		Usage:    "端口",
		Required: false,
	},
}

var createCommand = &cli.Command{
	Name:    "create",
	Aliases: []string{"c"},
	Usage:   "创建游戏",
	Flags:   createFlags,
	Action: func(ctx *cli.Context) error {
		c := app.GetApp()
		game := game.New()
		c.SetGame(game)
		c.SetTcpServer(tcp.NewServer())
		c.SetTcpClient(tcp.NewClient())
		return nil
	},
}

var joinCommand = &cli.Command{
	Name:    "join",
	Aliases: []string{"j"},
	Usage:   "加入游戏",
	Action: func(ctx *cli.Context) error {
		// TODO: 判断是否在游戏中
		c := app.GetApp()
		c.SetTcpClient(tcp.NewClient())
		return nil
	},
}

var statusCommand = &cli.Command{
	Name:    "status",
	Usage:   "显示游戏状态",
	Action: func(ctx *cli.Context) error {
		app.GetApp().GetGame().PrintGameStatus()
		return nil
	},
}

var startCommand = &cli.Command{
	Name:    "start",
	Usage:   "开始游戏",
	Action: func(ctx *cli.Context) error {
		// add game
		g := app.GetApp().GetGame()
		game.StartGame(*g)

		// run socket server
		go tcp.NewServer()

		// create socket client
		c := app.GetApp()
		c.SetTcpClient(tcp.NewClient())

		return nil
	},
}

var gameCommands = &cli.Command{
	Name:    		"game",
	Aliases:	 	[]string{"g"},
	Usage:   		"游戏操作",
	Subcommands: 	[]*cli.Command{
		statusCommand,
		startCommand,
		createCommand,
		joinCommand,
	},
}