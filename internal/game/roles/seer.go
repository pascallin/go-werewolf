package roles

type seer struct {
	//预言家
	name string
	life int
}

func (s *seer) IsAlive() bool {
	if s.life <= 0 {
		return false
	} else {
		return true
	}
}
func (s *seer) IsGood() bool {
	return false
}
func (s *seer) Kill()  {
	s.life --
	return
}
func (s *seer) Exile()  {
	s.life --
	return
}