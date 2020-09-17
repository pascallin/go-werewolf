package roles

type WereWolf struct {
	RoleClass
}

func NewWereWolf() RoleClass {
	var role = New("狼人(WereWolf)",  "BAD")
	return role
}