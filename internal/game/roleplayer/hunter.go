package roleplayer

type IHunter interface {
	ShotPlayer()
}

type HunterPlayer struct {
	RolePlayer
}