package game

import (
	"strconv"
	"testing"
)

func genPlayers(count int) []Player {
	var players []Player
	for i := 0; i < count; i++ {
		players = append(players, NewPlayer("player"+strconv.Itoa(i)))
	}
	return players
}

func TestGameFlow(t *testing.T) {
	game := NewGame()

	players := genPlayers(game.PlayersCount)
	for _, p := range players {
		game.JoinPlayer(p)
	}

	StartGame(game)
}