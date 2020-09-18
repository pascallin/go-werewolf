package roles

type Witch struct {
	Role
}

func NewWitch() Role {
	var role = New("女巫(Witch)",  Good)
	return role
}