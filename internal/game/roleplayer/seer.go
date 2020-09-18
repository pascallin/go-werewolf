package roleplayer

type SeerPlayer struct {
	RolePlayer
}

func (s SeerPlayer) CheckPlayer(p *RolePlayer) Side {
	return p.Side
}

type ISeer interface {
	CheckPlayer() Side
}
