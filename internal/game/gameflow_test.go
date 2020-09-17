package game

import (
	"strconv"
	"testing"

	"github.com/pascallin/go-wolvesgame/internal/game/palyer"
)

//for temporary
func genPlayers(count int) []palyer.Player {
	var players []palyer.Player
	for i := 0; i < count; i++ {
		players = append(players, palyer.NewPlayer("player"+strconv.Itoa(i)))
	}
	return players
}

func TestGameFlow(t *testing.T) {
	game := CreateGame()

	players := genPlayers(game.playersCount)
	for _, p := range players {
		game.JoinPlayer(p)
	}

	game.GameStart()
}