package player

import (
	"github.com/pascallin/go-wolvesgame/game/role"
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

func (player Player) SetRole(r role.RoleClass) Player {
	player.role = r
	return player
}

func NewPlayer(name string) Player {
	var p Player
	p.uid = uuid.NewV4()
	p.name = name
	p.life = 1
	return p
}