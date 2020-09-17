package roles

type Role interface {
	IsAlive() bool // 是否存活
	IsGood() bool  // 预言家查验
	Kill()         // 狼人杀人/女巫毒人/猎人开枪
	Exile()        // 投票投出
}

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