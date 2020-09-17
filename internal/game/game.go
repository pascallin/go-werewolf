package game

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/pascallin/go-wolvesgame/internal/game/gameround"
	"github.com/pascallin/go-wolvesgame/internal/game/palyer"
	gameroles "github.com/pascallin/go-wolvesgame/internal/game/roles"
)

type Game struct {
	status Status               	`json:"status"`
	roundNumber uint64          	`json:"roundNumber"`
	playersCount int            	`json:"playersCount"`
	participants int            	`json:"participants"`
	players []palyer.Player     	`json:"players"`
	roles []gameroles.Role 			`json:"roles"`
	round gameround.Round			`json:"round"`
}

func (g *Game) GameStart() error {
	if g.playersCount != g.participants {
		return errors.New("no enough players")
	}
	g.assignRoles()

	g.round = gameround.New(g.roles)
	var i = 1
	for {
		g.round.RoundStart(i, g.getPlayersLeft())
		gameover := g.checkGameOver()
		if gameover {
			GameOver()
			break
		}
		i++
		continue
	}
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

func (g *Game) checkGameOver() bool {
	var goodman uint64
	var badguy uint64
	for _, player := range g.getPlayersLeft() {
		if player.Role.Side == gameroles.Good && player.IsAlive() {
			goodman++
		}
	}
	for _, player := range g.getPlayersLeft() {
		if player.Role.Side == gameroles.Bad && player.IsAlive() {
			badguy++
		}
	}
	return badguy >= goodman
}

func (g *Game) getPlayersLeft() []palyer.Player {
	var playersLeft []palyer.Player
	for _, player := range g.players {
		if player.IsAlive() {
			playersLeft = append(playersLeft, player)
		}
	}
	return playersLeft
}

func CreateGame() Game {
	game := Game{
		status: Ready,
		roundNumber: 0,
		playersCount: 12,
		participants: 0,
		players: []palyer.Player{},
		roles: gen12Roles(),
	}

	return game
}


func GameOver() {
	fmt.Println("============= game over =============")
}

