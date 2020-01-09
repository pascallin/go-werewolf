package role

type Villager struct {
	RoleClass
}

func NewVillager() RoleClass {
	var role = New("村民(Villager)",  "GOOD")
	return role
}