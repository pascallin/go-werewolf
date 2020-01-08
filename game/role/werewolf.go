package role

type werewolf struct {
	name string
	life int
}

func (w *werewolf) IsAlive() bool {
	if w.life <= 0 {
		return false
	} else {
		return true
	}
}
func (w *werewolf) IsGood() bool {
	return false
}
func (w *werewolf) Kill()  {
	w.life --
	return
}
func (w *werewolf) Exile()  {
	w.life --
	return
}