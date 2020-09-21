package roleplayer

type SeerPlayer struct {
	RolePlayer
}

func (s SeerPlayer) CheckWerewolf(p IRolePlayer) bool {
	return p.IsWerewolf()
}

type ISeer interface {
	CheckWerewolf() Side
}
