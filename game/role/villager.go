package role

type villager struct {
	//村民
	name string
	life int
}

func (v *villager) IsAlive() bool {
	if v.life <= 0 {
		return false
	} else {
		return true
	}
}
func (v *villager) IsGood() bool {
	return false
}
func (v *villager) Kill()  {
	v.life --
	return
}
func (v *villager) Exile()  {
	v.life --
	return
}