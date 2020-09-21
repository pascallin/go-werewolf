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
func (player *Player) IsAlive() bool {
	if player.life <= 0 {
		return false
	} else {
		return true
	}
}
func (player *Player) BeKilled()  {
	player.life--
	return
}

func (player *Player) BeSaved()  {
	player.life++
	return
}

func (player *Player) Exile() {
	player.life--
	return
}
func (player *Player) InitRole(t RoleType) {
	player.Type = t
	player.life = 1
	if t == Werewolf {
		player.Side = Bad
	} else {
		player.Side = Good
	}
	if t == Witch {
		player.Poison = true
		player.AntiDote = true
	} else {
		player.Poison = false
		player.AntiDote = false
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
