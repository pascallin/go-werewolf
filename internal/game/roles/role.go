package roles

type Side int

const (
	Good Side = iota
	Bad
)

func (d Side) String() string {
	return [...]string{"Good", "Bad"}[d]
}

type Role struct {
	Name string
	Side Side
}

func (r *Role) IsGood() bool {
	return r.Side == Good
}

func New(name string, side Side) Role {
	var role Role
	role.Name = name
	role.Side = side
	return role
}