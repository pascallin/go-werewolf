package roles

type WereWolf struct {
	Role
}

func NewWereWolf() Role {
	var role = New("狼人(WereWolf)",  Bad)
	return role
}