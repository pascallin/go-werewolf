package game

import "github.com/pascallin/go-wolvesgame/internal/game/roleplayer"

func (g *Game) GetWerewolfPlayersLeft() []*roleplayer.WerewolfPlayer {
	var playersLeft  []*roleplayer.WerewolfPlayer
	for i, player := range g.RolePlayers {
		switch player.(type) {
		case roleplayer.WerewolfPlayer:
			playersLeft = append(playersLeft, &g.RolePlayers[i])
			break
		}
	}
	return playersLeft
}

func (g *Game) GetGoodPlayersLeft() []*interface{} {
	var playersLeft []*interface{}
	for i, player := range g.RolePlayers {
		switch player.(type) {
		case roleplayer.WitchPlayer, roleplayer.SeerPlayer, roleplayer.HunterPlayer, roleplayer.IdiotPlayer, roleplayer.VillagerPlayer:
			playersLeft = append(playersLeft, &g.RolePlayers[i])
			break
		}
	}
	return playersLeft
}