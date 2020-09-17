package roles

type Seer struct {
	Role
}

func (s *Seer) CheckPlayerSide(playerId string) string {
	return "GOOD"
}

func NewSeer() Role {
	var role = New("预言家(Seer)",  Good)
	return role
}