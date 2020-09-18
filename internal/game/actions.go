package game

import (
	"errors"
	"fmt"
)

func RoundStart(number int, game Game) {
	fmt.Println("============ round ============", number)

	// night fall
	fmt.Println("night fall")
	NightFall(game)
	// night actions

	// day
	fmt.Println("sunrise")
}

func StartGame(game Game) error {
	if game.PlayersCount != game.Participants {
		return errors.New("no enough players")
	}
	game.AssignRoles()

	var i = 1
	for {
		RoundStart(i, game)
		gameover := game.CheckGameOver()
		if gameover {
			OverGame()
			break
		}
		i++
		continue
	}
	return nil
}

func OverGame() {
	fmt.Println("============= game over =============")
}
