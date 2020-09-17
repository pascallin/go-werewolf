package roles

type witch struct {
	//女巫
	name string
	life int
}


func (w *witch) IsAlive() bool {
	if w.life <= 0 {
		return false
	} else {
		return true
	}
}
func (w *witch) IsGood() bool {
	return false
}
func (w *witch) Kill()  {
	w.life --
	return
}
func (w *witch) Exile()  {
	w.life --
	return
}