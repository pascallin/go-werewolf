package roleplayer

import (
	uuid "github.com/satori/go.uuid"
)

type Side int

const (
	Good Side = iota
	Bad
)

func (d Side) String() string {
	return [...]string{"Good", "Bad"}[d]
}

type Type int

const (
	Seer Type = iota
	Witch
	Hunter
	Idiot
	Werewolf
	Villager
)

type RolePlayer struct {
	id  uuid.UUID
	Name string
	Type  Type
	Alias string
	Side  Side
	life int
}

func (player *RolePlayer) IsWerewolf() bool {
	return player.Side == Bad
}

func (player *RolePlayer) IsAlive() bool {
	if player.life <= 0 {
		return false
	} else {
		return true
	}
}

func (player *RolePlayer) BeKilled()  {
	player.life --
	return
}

func (player *RolePlayer) Exile() {
	player.life--
	return
}

func (player *RolePlayer) SetRole(t Type) {
	player.Type = t
	if t == Werewolf {
		player.Side = Bad
	} else {
		player.Side = Good
	}
}

func New(id uuid.UUID, name string) RolePlayer {
	var p RolePlayer
	p.id = id
	p.Name = name
	p.life = 1
	return p
}

func NewRolePlayer(p RolePlayer, t Type) interface{} {
	var rp interface{}
	switch t {
	case Werewolf:
		rp = NewWerewolfPlayer(p)
		break
	case Villager:
		rp = VillagerPlayer{p}
		break
	case Hunter:
		rp = HunterPlayer{p}
		break
	case Seer:
		rp = SeerPlayer{p}
		break
	case Idiot:
		rp = IdiotPlayer{p}
		break
	case Witch:
		rp = WitchPlayer{p}
		break
	}
	return rp
}