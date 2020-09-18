package roleplayer

type WerewolfPlayer struct {
	RolePlayer
}

func (w WerewolfPlayer) KillPlayer(p interface{}) {
	p.(RolePlayer).BeKilled()
}

func NewWerewolfPlayer(p RolePlayer) WerewolfPlayer {
	p.Side = Bad
	return WerewolfPlayer{
		p,
	}
}

type IWerewolf interface {
	KillPlayer()
}
