package game

import (
	"fmt"
)

func NightFall(game Game) {
	// kill goodman
	players := game.GetGoodPlayersLeft()
	fmt.Println("Kill =================> ", players[0])
	players[0].BeKilled()
	// god skills
	for _, p := range game.GetPlayersLeft() {
		if p.Role.Name == "预言家(Seer)" &&p.IsAlive() {
			fmt.Println("=============== Seer =================", p)
		}
		if p.Role.Name == "女巫(Witch)" &&p.IsAlive() {
			fmt.Println("=============== Witch =================", p)
		}
	}
}

func Sunrise(game Game) {
	// show dead man
	// random pick speaker
	// vote
}