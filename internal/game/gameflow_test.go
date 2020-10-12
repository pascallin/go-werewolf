package game

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func genPlayers(count int) []Player {
	var players []Player
	for i := 0; i < count; i++ {
		players = append(players, NewPlayer(i, "player"+strconv.Itoa(i)))
	}
	return players
}

func TestGameFlow(t *testing.T) {
	game := New()

	players := genPlayers(game.PlayersCount)
	for _, p := range players {
		game.JoinPlayer(p)
	}

	go game.Start()

	// NOTE: fake player user actions
	for {
		select {
		case s := <-game.Lifecycle:
			if s == WaitingNightPlayerAction {
				ai_werewolfAction(game)
				ai_seerAction(game)
				ai_poisonAction(game)
				ai_antidoteAction(game)
			}
			if s == WaitingDayPlayerAction {
				ai_talking(game)
				ai_voting(game)
			}
			if s == Over {
				fmt.Println(game.GameStatusJSON())
				os.Exit(0)
			}
		}
	}
}