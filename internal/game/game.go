package game

import (
	"errors"
	"fmt"
	"math/rand"

	gameroles "github.com/pascallin/go-wolvesgame/internal/game/roles"
)

type Game struct {
	Status       Status           `json:"status"`
	PlayersCount int              `json:"playersCount"`
	Participants int              `json:"participants"`
	Players      []Player         `json:"players"`
	Roles        []gameroles.Role `json:"roles"`
}


func NewGame() Game {
	game := Game{
		Status:       Ready,
		PlayersCount: 12,
		Participants: 0,
		Players:      []Player{},
		Roles:        gen12Roles(),
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

func (g *Game) AssignRoles() {
	randomList, _ := genRandomList(g.Participants)
	for i, _ := range g.Players {
		g.Players[i].SetRole(g.Roles[randomList[i]])
	}
}

func (g *Game) CheckGameOver() bool {
	var goodman uint64
	var badguy uint64
	for _, player := range g.GetPlayersLeft() {
		if player.Role.Side == gameroles.Good && player.IsAlive() {
			goodman++
		}
	}
	for _, player := range g.GetPlayersLeft() {
		if player.Role.Side == gameroles.Bad && player.IsAlive() {
			badguy++
		}
	}
	return badguy >= goodman
}

func (g *Game) GetPlayersLeft() []*Player {
	var playersLeft []*Player
	for i, player := range g.Players {
		if player.IsAlive() {
			playersLeft = append(playersLeft, &g.Players[i])
		}
	}
	return playersLeft
}

func (g *Game) GetGoodPlayersLeft() []*Player {
	var playersLeft []*Player
	for i, player := range g.Players {
		if player.IsAlive() && player.Role.IsGood() {
			playersLeft = append(playersLeft, &g.Players[i])
		}
	}
	return playersLeft
}