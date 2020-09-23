package app

import (
	"github.com/pascallin/go-wolvesgame/internal/game"
	"sync"

	"github.com/pascallin/go-wolvesgame/pkg/tcp"
)

type App struct {
	Game      		*game.Game
	User			*User
	TCPClient 		*tcp.TCPClient
	TCPServer 		*tcp.Server
}

// getter and setter
func (c *App) SetGame(game game.Game) {
	c.Game = &game
}
func (c App) GetGame() *game.Game {
	return c.Game
}
func (c *App) SetUser(user User) {
	c.User = &user
}
func (c App) GetUser() *User {
	return c.User
}
func (c *App) SetTcpClient(client *tcp.TCPClient) {
	c.TCPClient = client
}
func (c App) GetTcpClient() *tcp.TCPClient {
	return c.TCPClient
}
func (c *App) SetTcpServer(server *tcp.Server) {
	c.TCPServer = server
}
func (c App) GetTcpServer() *tcp.Server {
	return c.TCPServer
}

// instance method
var instance *App
var once sync.Once

func GetApp() *App {
	once.Do(func() {
		instance = &App{
		}
	})
	return instance
}