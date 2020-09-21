package game

// Player good or bad
type Side int
const (
	Good Side = iota
	Bad
)

func (d Side) String() string {
	return [...]string{"Good", "Bad"}[d]
}

// Player role type
type RoleType int
const (
	Seer RoleType = iota
	Witch
	Hunter
	Idiot
	Werewolf
	Villager
)

type Player struct {
	ID			int			`json:"id"`
	Name 		string		`json:"name"`
	Type  		RoleType	`json:"type"`
	Alias 		string		`json:"alias"`
	Side  		Side		`json:"side"`
	life 		int			`json:"life"`
	Poison  	bool		`json:"poison"`
	AntiDote 	bool		`json:"antidote"`
	RoundVoted	int			`json:"-"`
}

func (player *Player) IsWerewolf() bool {
	return player.Side == Bad
}
func (p *Player) IsAlive() bool {
	if p.life <= 0 {
		return false
	} else {
		return true
	}
}
func (p *Player) BeKilled()  {
	p.life--
	return
}

func (p *Player) BeSaved()  {
	p.life++
	return
}

func (p *Player) Exile() {
	p.life--
	return
}
func (p *Player) InitRole(t RoleType) {
	p.Type = t
	p.life = 1
	if t == Werewolf {
		p.Side = Bad
	} else {
		p.Side = Good
	}
	if t == Witch {
		p.Poison = true
		p.AntiDote = true
	} else {
		p.Poison = false
		p.AntiDote = false
	}
}

func (p *Player) clearRoundVoted() {
	p.RoundVoted = 0
}

func NewPlayer(id int, name string) Player {
	return Player{
		ID: id,
		Name: name,
	}
}
