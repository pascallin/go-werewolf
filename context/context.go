package context

import (
	"net"
	"sync"

	"github.com/pascallin/go-wolvesgame/game"
)

type context struct {
	game game.Game
	socket net.Conn
}

func (c *context) SetGame(game game.Game) {
	c.game = game
}
func (c context) GetGame() game.Game {
	return c.game
}
func (c *context) SetSocket(socket net.Conn) {
	c.socket = socket
}
func (c context) GetSocket() net.Conn {
	return c.socket
}

var instance *context
var once sync.Once

func GetContext() *context {
	once.Do(func() {
		instance = &context{}
	})
	return instance
}