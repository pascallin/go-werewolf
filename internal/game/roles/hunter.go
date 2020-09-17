package roles

type Hunter struct {
	Role
}

func NewHunter() Role {
	var role = New("猎人(Hunter)",  Good)
	return role
}
