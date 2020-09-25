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
	TCPClient 	*tcp.Client
	TCPServer 	*tcp.Server
	messageWriter io.Writer
}

func New(username string, messageWriter io.Writer) *App {
	app := &App{
		User: NewUser(username),
		TCPServer: tcp.NewServer(),
		messageWriter: messageWriter,
	}
	if messageWriter != nil {
		app.TCPClient = tcp.NewClient(messageWriter)
	} else {
		app.TCPClient = tcp.NewClient(os.Stdout)
	}
	go app.serverListen()
	return app
}

func (app *App) serverListen() {
	for {
		select {
		case client := <- app.TCPServer.Register:
			app.TCPServer.BroadcastMessage(client.ID.String() + "client joined")
			//fmt.Fprint(app.messageWriter, client.ID.String() + "client joined")
		case client := <- app.TCPServer.Unregister:
			//fmt.Fprint(app.messageWriter, client.ID.String() + "client left")
			app.TCPServer.BroadcastMessage(client.ID.String() + "client left")
		}
	}
}