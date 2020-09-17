package roles

type Idiot struct {
	Role
}

func NewIdiot() Role {
	var role = New("白痴(Idiot)",  Good)
	return role
}