package context

import (
	"sync"

	"github.com/pascallin/go-wolvesgame/game"
	"github.com/pascallin/go-wolvesgame/transport/tcp"
)

type context struct {
	game game.Game
	tcpClient *tcp.Client
	tcpServer *tcp.Server
}

// getter and setter
func (c *context) SetGame(game game.Game) {
	c.game = game
}
func (c context) GetGame() game.Game {
	return c.game
}
func (c *context) SetTcpClient(client *tcp.Client) {
	c.tcpClient = client
}
func (c context) GetTcpClient() *tcp.Client {
	return c.tcpClient
}
func (c *context) SetTcpServer(server *tcp.Server) {
	c.tcpServer = server
}
func (c context) GetTcpServer() *tcp.Server {
	return c.tcpServer
}

// instance method
var instance *context
var once sync.Once

func GetContext() *context {
	once.Do(func() {
		instance = &context{}
	})
	return instance
}