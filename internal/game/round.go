package game

import (
	"fmt"
)

func RoundStart(number int, game *Game) {
	manLeft := GetManLeft(game)
	for i := range manLeft {
		manLeft[i].clearRoundVoted()
	}
	// NOTE: night actions
	game.Lifecycle <- WaitingNightPlayerAction
	fmt.Println("============ round ============", number)
	killed := <-game.PlayerActions.WerewolfKill
	fmt.Println("WerewolfAction kill ============", killed)
	killed.BeKilled()

	saw := <-game.PlayerActions.SeerCheck
	fmt.Println("Seer check ============", saw.IsWerewolf())
	poison := <-game.PlayerActions.UsePoison
	if poison != nil {
		fmt.Println("UsePoison kill ============", poison)
		poison.BeKilled()
	}
	antidote := <-game.PlayerActions.UseAntidote
	if antidote {
		fmt.Println("UseAntidote save ============")
		killed.BeSaved()
	}
	// NOTE: day actions
	game.Lifecycle <- WaitingDayPlayerAction
	listenTalking(game)
}

func listenTalking(game *Game) {
	talking := GetManLeft(game)
	var talked = 0
	var talkingEnd = false
	fmt.Println("Talking ============", len(talking))
	for !talkingEnd {
		select {
		case playerTalked := <-game.PlayerActions.TalkedCount:
			talked += playerTalked
			if talked == len(talking) {
				talkingEnd = true
			}
		}
	}
	fmt.Println("Talking End ============")
}