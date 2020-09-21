package game

import (
	"errors"
	"fmt"
)

type Status int
const (
	Waiting Status = iota	// wait for start
	Ready					// game started or restart
	WaitingPlayerAction
	Over					// game over
)

func (d Status) String() string {
	return [...]string{"Waiting", "Ready", "WaitingPlayerAction", "Over"}[d]
}

type Game struct {
	PlayersCount  int              	`json:"playersCount"` // game player needed
	Participants  int              	`json:"participants"`
	PlayerActions PlayerActions
	Players       []Player
	Lifecycle     chan Status
	RoundNumber   int
}

var (
	Roles12 = []RoleType{
		Villager, Villager, Villager, Villager,
		Werewolf, Werewolf, Werewolf , Werewolf,
		Seer, Witch, Idiot, Hunter,
	}
)

func New() Game {
	game := Game{
		PlayersCount:  	12,
		Participants:  	0,
		PlayerActions: 	NewActions(),
		Lifecycle:     	make(chan Status),
		RoundNumber: 	0,
	}

	return game
}

func (g *Game) PrintGameStatus() {
	fmt.Printf("%#v\n", g)
}

func (g *Game) JoinPlayer(player Player) {
	g.Participants += 1
	g.Players = append(g.Players, player)
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
	for _, player := range GetManLeft(g) {
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

func StartGame(game *Game) error {
	game.Lifecycle <- Waiting

	if game.PlayersCount != game.Participants {
		return errors.New("no enough players")
	}
	game.AssignRoles()

	game.Lifecycle <- Ready

	var i = 1
	for {
		game.RoundNumber = i
		RoundStart(i, game)
		if game.CheckGameOver() {
			break
		}
		i++
		continue
	}
	game.GameOver()
	return nil
}

func (g *Game) GameOver() {
	fmt.Println("============= game over =============")
	g.Lifecycle <- Over
}