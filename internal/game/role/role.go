package role

type RoleClass struct {
	name string
	side string
}

func New(name string, side string) RoleClass {
	var role RoleClass
	role.name = name
	role.side = side
	return role
}