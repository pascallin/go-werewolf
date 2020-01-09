package game

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/pascallin/go-wolvesgame/game/role"
	"github.com/pascallin/go-wolvesgame/game/player"
)

type Game struct {
	status string `json:"status"`
	round int `json:"round"`
	playersCount int `json:"playersCount"`
	participants int `json:"participants"`
	players []player.Player `json:"players"`
	roles []role.RoleClass `json:"roles"`
}

func (g Game) GameStart() (Game, error) {
	if g.playersCount != g.participants {
		return g, errors.New("no enough players")
	}
	g.assignRoles()
	return g, nil
}

func (g Game) PrintGameStatus() {
	fmt.Printf("%+v", g)
}

func (g Game) JoinPlayer(player player.Player) Game {
	g.participants += 1
	g.players = append(g.players, player)
	return g
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
func (g Game) assignRoles() Game {
	randomList, _ := genRandomList(g.participants)
	for i, _ := range g.players {
		g.players[i] = g.players[i].SetRole(g.roles[randomList[i]])
	}
	return g
}

//for temporary
func genPlayers(count int) []player.Player {
	var players []player.Player
	for i := 0; i < count; i++ {
		players = append(players, player.NewPlayer(strconv.Itoa(i)))
	}
	return players
}
//for temporary
func genRoles() []role.RoleClass {
	var roles []role.RoleClass
	// villagers
	for i := 0; i < 4; i++ {
		roles = append(roles, role.NewVillager())
	}
	// wolves
	for i := 0; i < 4; i++ {
		roles = append(roles, role.NewWereWolf())
	}
	// gods
	roles = append(roles, role.NewVillager())
	roles = append(roles, role.NewVillager())
	roles = append(roles, role.NewVillager())
	roles = append(roles, role.NewVillager())

	return roles
}

func CreateGame() Game {
	game := Game{
		status: "ready",
		round: 0,
		playersCount: 12,
		participants: 0,
		players: []player.Player{},
		roles: genRoles(),
	}

	players := genPlayers(game.playersCount)
	for _, p := range players {
		game = game.JoinPlayer(p)
	}

	game, _ = game.GameStart()
	game.PrintGameStatus()
	return game
}
