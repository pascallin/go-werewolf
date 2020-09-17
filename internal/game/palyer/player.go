package palyer

import (
	"github.com/pascallin/go-wolvesgame/internal/game/roles"
	uuid "github.com/satori/go.uuid"
)

type Player struct {
	uid  uuid.UUID
	name string
	life int
	role interface{}
}

func (player Player) IsAlive() bool {
	if player.life <= 0 {
		return false
	} else {
		return true
	}
}

func (player Player) Kill()  {
	player.life --
	return
}

func (player Player) Exile() {
	player.life--
	return
}

func (player *Player) SetRole(r roles.RoleClass) {
	player.role = r
}

func NewPlayer(name string) Player {
	var p Player
	p.uid = uuid.NewV4()
	p.name = name
	p.life = 1
	return p
}