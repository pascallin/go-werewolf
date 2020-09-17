package palyer

import (
	"github.com/pascallin/go-wolvesgame/internal/game/roles"
	uuid "github.com/satori/go.uuid"
)

type Player struct {
	Uid  uuid.UUID
	Name string
	life int
	Role roles.Role
}

func (player *Player) IsAlive() bool {
	if player.life <= 0 {
		return false
	} else {
		return true
	}
}

func (player *Player) BeKilled()  {
	player.life --
	return
}

func (player *Player) Exile() {
	player.life--
	return
}

func (player *Player) SetRole(r roles.Role) {
	player.Role = r
}

func NewPlayer(name string) Player {
	var p Player
	p.Uid = uuid.NewV4()
	p.Name = name
	p.life = 1
	return p
}