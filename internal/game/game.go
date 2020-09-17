package game

import (
	"errors"
	"fmt"
	"github.com/pascallin/go-wolvesgame/internal/game/palyer"
	"math/rand"
	"strconv"

	gameroles "github.com/pascallin/go-wolvesgame/internal/game/roles"
)

type Game struct {
	status string               `json:"status"`
	round int                   `json:"round"`
	playersCount int            `json:"playersCount"`
	participants int            `json:"participants"`
	players []palyer.Player     `json:"players"`
	roles []gameroles.RoleClass `json:"roles"`
}

func (g *Game) GameStart() error {
	if g.playersCount != g.participants {
		return errors.New("no enough players")
	}
	g.assignRoles()
	return nil
}

func (g Game) PrintGameStatus() {
	fmt.Printf("%#v\n", g)
}

func (g *Game) JoinPlayer(player palyer.Player) {
	g.participants += 1
	g.players = append(g.players, player)
}

func genRandomList(length int) ([]int, error) {
	if length <= 0 {
		return nil, errors.New("the size of the parameter length illegal")
	}
	var list []int
	for i := 0; i < length; i++ {
		list = append(list, i)
	}
	for i := len(list) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		list[i], list[num] = list[num], list[i]
	}
	return list, nil
}
func (g *Game) assignRoles() {
	randomList, _ := genRandomList(g.participants)
	for i, _ := range g.players {
		g.players[i].SetRole(g.roles[randomList[i]])
	}
}

//for temporary
func genPlayers(count int) []palyer.Player {
	var players []palyer.Player
	for i := 0; i < count; i++ {
		players = append(players, palyer.NewPlayer(strconv.Itoa(i)))
	}
	return players
}
//for temporary
func genRoles() []gameroles.RoleClass {
	var roles []gameroles.RoleClass
	// villagers
	for i := 0; i < 4; i++ {
		roles = append(roles, gameroles.NewVillager())
	}
	// wolves
	for i := 0; i < 4; i++ {
		roles = append(roles, gameroles.NewWereWolf())
	}
	// gods
	roles = append(roles, gameroles.NewVillager())
	roles = append(roles, gameroles.NewVillager())
	roles = append(roles, gameroles.NewVillager())
	roles = append(roles, gameroles.NewVillager())

	return roles
}

func CreateGame() Game {
	game := Game{
		status: "ready",
		round: 0,
		playersCount: 12,
		participants: 0,
		players: []palyer.Player{},
		roles: genRoles(),
	}

	players := genPlayers(game.playersCount)
	for _, p := range players {
		game.JoinPlayer(p)
	}

	return game
}
