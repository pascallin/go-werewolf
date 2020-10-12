package game

import (
	"errors"
	"fmt"
)

func (g *Game) JoinPlayer(player Player) {
	g.Participants += 1
	g.Players = append(g.Players, player)
}

func (g *Game) RemovePlayer(player Player) {
	g.Participants -= 1
	for i, v := range g.Players {
		if player.ID == v.ID {
			// remove player
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
		}
	}
}

func (g *Game) AssignRoles() {
	randomList, _ := genRandomList(g.Participants)
	for i := range g.Players {
		g.Players[i].InitRole(Roles12[randomList[i]])
	}
}

func (g *Game) CheckGameOver() bool {
	var villagerCount int
	var werewolfCount int
	var godCount int
	for _, player := range g.GetManLeft() {
		if player.Type == Werewolf {
			werewolfCount++
		}
		if player.Type == Villager {
			villagerCount++
		}
		if player.Type != Villager && player.Type != Werewolf {
			godCount++
		}
	}
	// side killed
	if villagerCount == 0 || werewolfCount == 0 || godCount == 0 {
		return true
	}
	return false
}

func (g *Game) Start() error {
	g.Lifecycle <- Waiting

	if g.PlayersCount != g.Participants {
		return errors.New("no enough players")
	}
	g.AssignRoles()

	g.Lifecycle <- Ready

	var i = 1
	for {
		g.RoundNumber = i
		g.RoundStart(i)
		if g.CheckGameOver() {
			break
		}
		i++
		continue
	}
	g.Over()
	return nil
}

func (g *Game) Over() {
	fmt.Println("============= game over =============")
	g.Lifecycle <- Over
}