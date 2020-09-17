package roles

type hunter struct {
	//猎人
	name string
	life int
}

func (h *hunter) IsAlive() bool {
	if h.life <= 0 {
		return false
	} else {
		return true
	}
}
func (h *hunter) IsGood() bool {
	return false
}
func (h *hunter) Kill() {
	h.life--
	return
}
func (h *hunter) Exile() {
	h.life--
	return
}
