package gameround

import (
	"fmt"
	"github.com/pascallin/go-wolvesgame/internal/game/palyer"
	"github.com/pascallin/go-wolvesgame/internal/game/roles"
)

type Round struct {
	number int
	Roles []roles.Role // game exist roles
}

func New(roles []roles.Role) Round {
	return Round{
		number: 0,
		Roles: roles,
	}
}

func (r Round) RoundStart(number int, players []palyer.Player) {
	fmt.Println("============ round ============", number)

	// night fall
	fmt.Println("night fall")
	NightFall(players)
	// night actions

	// day
	fmt.Println("sunrise")
}
