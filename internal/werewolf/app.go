package werewolf

import (
	"fmt"
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
			fmt.Println("new client join.....")
			app.TCPServer.BroadcastMessage("client 【" + client.ID.String() + "】joined")
			app.Game.JoinPlayer(game.NewPlayer(app.Game.Participants + 1, client.ID.String()))
		case client := <- app.TCPServer.Unregister:
			app.TCPServer.BroadcastMessage("client 【" + client.ID.String() + "】left")
		case message := <- app.TCPServer.Receiver:
			app.messageHandler(message)
		}
	}
}

func (app *App) messageHandler(message string) {
	ms := MessageDecode(message)
	// TODO: dispatch actions
	app.TCPServer.BroadcastMessage( ms.Username + " said: " + ms.Content)
}