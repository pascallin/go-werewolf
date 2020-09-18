package game

import (
	"fmt"
)

func NightFall(game Game) {
	goodman := game.GetGoodPlayersLeft()
	fmt.Println(goodman)

	// werewolf action
	// TODO: random kill
	badguy := game.GetWerewolfPlayersLeft()
	fmt.Println(badguy)

	//var actionWerewolf = badguy[0]
	//var theFirstGoodMan = goodman[0]

	//actionWerewolf.KillPlayer(theFirstGoodMan)

	//for _, p := range game.GetGoodPlayersLeft() {
	//	// seer skills
	//	if p.Type == roleplayer.Seer &&p.IsAlive() {
	//
	//	}
	//}
}

func Sunrise(game Game) {
	// show dead man
	// random pick speaker
	// vote
}