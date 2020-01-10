package context

import (
	"sync"

	"github.com/pascallin/go-wolvesgame/game"
)

type context struct {
	game game.Game
}

func (c *context) SetGame(game game.Game) {
	c.game = game
}
func (c context) GetGame() game.Game {
	return c.game
}

var instance *context
var once sync.Once

func GetContext() *context {
	once.Do(func() {
		instance = &context{}
	})
	return instance
}