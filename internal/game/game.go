package game

import (
	"errors"
	"fmt"

	"github.com/pascallin/go-wolvesgame/internal/game/player"
	"github.com/pascallin/go-wolvesgame/internal/game/roleplayer"
)

type Status int

const (
	Waiting Status = iota	// wait for start
	Ready					// game started or restart
	Block					// waiting for player action
	Over					// game over
)

func (d Status) String() string {
	return [...]string{"Waiting", "Ready", "Block", "Over"}[d]
}

type Game struct {
	Status       Status           	`json:"status"`
	PlayersCount int              	`json:"playersCount"` // game player needed
	Participants int              	`json:"participants"`
	Players      []player.Player  	`json:"players"`
	RolePlayers	 []interface{} 		`json:"rolePlayers"`
	//Roles 		 Roles12Position
}

type Roles12Position struct {
	WerewolfPlayers [4]roleplayer.WerewolfPlayer
	VillagerPlayers [4]roleplayer.VillagerPlayer
	HunterPlayer 	roleplayer.HunterPlayer
	IdiotPlayer 	roleplayer.IdiotPlayer
	SeerPlayer 		roleplayer.SeerPlayer
	WitchPlayer 	roleplayer.WitchPlayer
}

var (
	Roles12 = []roleplayer.Type{
		roleplayer.Villager, roleplayer.Villager, roleplayer.Villager, roleplayer.Villager,
		roleplayer.Werewolf, roleplayer.Werewolf, roleplayer.Werewolf , roleplayer.Werewolf,
		roleplayer.Seer, roleplayer.Witch, roleplayer.Idiot, roleplayer.Hunter,
	}
)

func New() Game {
	game := Game{
		Status:			Waiting,
		PlayersCount: 	12,
		Participants: 	0,
	}

	return game
}

func (g *Game) PrintGameStatus() {
	fmt.Printf("%#v\n", g)
}

func (g *Game) JoinPlayer(player player.Player) {
	g.Participants += 1
	g.Players = append(g.Players, player)
}

func (g *Game) AssignRoles() {
	var rolePlayers []interface{}
	randomList, _ := genRandomList(g.Participants)
	for i, _ := range g.Players {
		player := roleplayer.New(&g.Players[i])
		rolePlayers = append(rolePlayers, roleplayer.NewRolePlayer(player, Roles12[randomList[i]]))
	}
	g.RolePlayers = rolePlayers
}

func (g *Game) CheckGameOver() bool {
	return true
}

func StartGame(game Game) error {
	if game.PlayersCount != game.Participants {
		return errors.New("no enough players")
	}
	game.AssignRoles()

	var i = 1
	for {
		RoundStart(i, game)
		over := game.CheckGameOver()
		if over {
			game.GameOver()
			break
		}
		i++
		continue
	}
	return nil
}

func (g *Game) GameOver() {
	fmt.Println("============= game over =============")
	g.Status = Over
}

func (g *Game) BlockGame() {
	g.Status = Block
}

func (g *Game) RestartGame() {
	g.Status = Ready
}