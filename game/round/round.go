package round

import (
	"github.com/pascallin/go-wolvesgame/game/role"
)

type Round struct {
	number int
	Roles []role.RoleClass // game exist roles
}

func (r Round) RoundStart() {
	// night fall
	// night actions
	// day
}
