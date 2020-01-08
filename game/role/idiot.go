package role

type idiot struct {
	//白痴
	name string
	life int
}

func (i *idiot) IsAlive() bool {
	if i.life <= 0 {
		return false
	} else {
		return true
	}
}
func (i *idiot) IsGood() bool {
	return false
}
func (i *idiot) Kill()  {
	i.life --
	return
}
func (i *idiot) Exile()  {
	i.life --
	return
}