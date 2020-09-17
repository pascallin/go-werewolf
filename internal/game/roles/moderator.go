package roles

type moderator struct {
	//法官
	name string
	life int
}

func (m moderator) IsAlive() bool {
	if m.life <= 0 {
		return false
	} else {
		return true
	}
}
func (m moderator) IsGood() bool {
	return false
}
func (m moderator) Kill() {
	m.life--
	return
}
func (m moderator) Exile() {
	m.life--
	return
}
