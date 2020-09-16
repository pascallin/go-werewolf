package role

type Role interface {
	IsAlive() bool // 是否存活
	IsGood() bool  // 预言家查验
	Kill()         // 狼人杀人/女巫毒人/猎人开枪
	Exile()        // 投票投出
}
