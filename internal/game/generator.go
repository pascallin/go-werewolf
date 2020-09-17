package game

import (
	gameroles "github.com/pascallin/go-wolvesgame/internal/game/roles"
)

func gen12Roles() []gameroles.Role {
	var roles []gameroles.Role
	// villagers
	for i := 0; i < 4; i++ {
		roles = append(roles, gameroles.NewVillager())
	}
	// wolves
	for i := 0; i < 4; i++ {
		roles = append(roles, gameroles.NewWereWolf())
	}
	// gods
	roles = append(roles, gameroles.NewHunter())
	roles = append(roles, gameroles.NewIdiot())
	roles = append(roles, gameroles.NewSeer())
	roles = append(roles, gameroles.NewWitch())

	return roles
}