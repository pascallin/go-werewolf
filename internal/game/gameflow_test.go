package game

import (
	uuid "github.com/satori/go.uuid"
	"strconv"
	"testing"
)

func genPlayers(count int) []Player {
	var players []Player
	for i := 0; i < count; i++ {
		players = append(players, NewPlayer(uuid.NewV4(), "player"+strconv.Itoa(i)))
	}
	return players
}

func TestGameFlow(t *testing.T) {
	game := New()

	players := genPlayers(game.PlayersCount)
	for _, p := range players {
		game.JoinPlayer(p)
	}

	StartGame(game)
}