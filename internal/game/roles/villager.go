package roles

type Villager struct {
	Role
}

func NewVillager() Role {
	var role = New("村民(Villager)",  Good)
	return role
}