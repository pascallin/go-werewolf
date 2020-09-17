package roles

type RoleFactory struct {
}

func (f RoleFactory) CreateRole(roleType string) Role {
	switch roleType {
	case "seer":
		return &seer{"预言家", 1,}
	case "witch":
		return &witch{"女巫", 1,}
	case "hunter":
		return &hunter{"猎人", 1,}
	case "idiot":
		return &idiot{"白痴", 1,}
	default:
		return nil
	}
}
