package gameround

import (
	"fmt"
	"github.com/pascallin/go-wolvesgame/internal/game/palyer"
)

func NightFall(playersLeft []palyer.Player) {
	var goodman []palyer.Player
	for _, player := range playersLeft {
		if player.Role.IsGood() {
			goodman = append(goodman, player)
		}
	}
	fmt.Println("=============== kill =================", goodman[0])
	goodman[len(goodman) - 1].BeKilled()
}

