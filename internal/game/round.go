package game

import (
	"fmt"
)

func RoundStart(number int, game *Game) {
	game.Lifecycle <- WaitingPlayerAction
	fmt.Println("============ round ============", number)
	killed := <- game.PlayerActions.WerewolfKill
	fmt.Println("============ WerewolfAction kill ============", killed)
	killed.BeKilled()

	saw := <- game.PlayerActions.SeerCheck
	fmt.Println("============ Seer check ============", saw.IsWerewolf())
	poison := <- game.PlayerActions.UsePoison
	if poison != nil {
		fmt.Println("============ UsePoison kill ============", poison)
		poison.BeKilled()
	}
	antidote := <- game.PlayerActions.UseAntidote
	if antidote {
		fmt.Println("============ UseAntidote save ============")
		killed.BeSaved()
	}
}