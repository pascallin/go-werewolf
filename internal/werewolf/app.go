package app

import (
	"github.com/pascallin/go-wolvesgame/internal/game"
	"github.com/pascallin/go-wolvesgame/pkg/tcp"
)

type App struct {
	User		*User
	Game      	*game.Game
	Client 		*tcp.Client
	TCPServer 	*tcp.Server
}

func New(username string) *App {
	return &App{
		User: NewUser(username),
	}
}