package game

type PlayerActions struct {
	WerewolfKill 	chan *Player
	SeerCheck 		chan *Player
	UsePoison 		chan *Player
	UseAntidote 	chan bool
}

func NewActions() PlayerActions {
	return PlayerActions{
		WerewolfKill: 	make(chan *Player),
		SeerCheck: 		make(chan *Player),
		UsePoison: 		make(chan *Player),
		UseAntidote: 	make(chan bool),
	}
}