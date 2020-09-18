package roles

type Seer struct {
	Role
}

func NewSeer() Role {
	var role = New("预言家(Seer)",  Good)
	return role
}