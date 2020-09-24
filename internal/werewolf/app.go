package werewolf

import (
	"io"
	"os"

	"github.com/pascallin/go-wolvesgame/internal/game"
	"github.com/pascallin/go-wolvesgame/pkg/tcp"
)

type App struct {
	User		*User
	Game      	*game.Game
	TCPClient 		*tcp.Client
	TCPServer 	*tcp.Server
}

func New(username string, messageWriter io.Writer) *App {
	app := &App{
		User: NewUser(username),
	}
	if messageWriter != nil {
		app.TCPServer = tcp.NewServer(messageWriter)
		app.TCPClient = tcp.NewClient(messageWriter)
	} else {
		app.TCPServer =  tcp.NewServer(os.Stdout)
		app.TCPClient = tcp.NewClient(os.Stdout)
	}
	return app
}