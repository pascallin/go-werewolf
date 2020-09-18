package game

import (
	"fmt"
	"github.com/pascallin/go-wolvesgame/internal/game/roleplayer"
)

type Game struct {
	Status       Status			`json:"status"`
	PlayersCount int			`json:"playersCount"` // game player neeed
	Participants int			`json:"participants"`
	Players      []Player		`json:"players"`
	RolePlayers	 []interface{}	`json:"rolePlayers"`
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
		Status:       Ready,
		PlayersCount: 12,
		Participants: 0,
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
	var rolePlayers []interface{}
	randomList, _ := genRandomList(g.Participants)
	for i, _ := range g.Players {
		player := roleplayer.New(g.Players[i].ID, g.Players[i].Name)
		rolePlayers = append(rolePlayers, roleplayer.NewRolePlayer(player, Roles12[randomList[i]]))
	}
	g.RolePlayers = rolePlayers
}

//func (g *Game) CheckGameOver() bool {
//	goodman := len(g.GetGoodPlayersLeft())
//	badguy := len(g.GetWerewolfPlayersLeft())
//	if badguy == 0 {
//		return true
//	}
//	return badguy >= goodman
//}