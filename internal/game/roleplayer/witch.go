package roleplayer

type WitchPlayer struct {
	RolePlayer
}

type IWitch interface {
	UsePoison()
	UseAntidote()
}