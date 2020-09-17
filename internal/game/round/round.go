package round

import (
	"github.com/pascallin/go-wolvesgame/internal/game/roles"
)

type Round struct {
	number int
	Roles []roles.RoleClass // game exist roles
}

func (r Round) RoundStart() {
	// night fall
	// night actions
	// day
}
