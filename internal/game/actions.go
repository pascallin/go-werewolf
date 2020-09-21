package game

type PlayerActions struct {
	// Night Actions
	WerewolfKill 	chan *Player
	SeerCheck 		chan *Player
	UsePoison 		chan *Player
	UseAntidote 	chan bool
	// Day actions
	TalkedCount		chan int
	Voting			chan *Player
}

func NewActions() PlayerActions {
	return PlayerActions{
		WerewolfKill: 	make(chan *Player),
		SeerCheck: 		make(chan *Player),
		UsePoison: 		make(chan *Player),
		UseAntidote: 	make(chan bool),
		TalkedCount: 	make(chan int),
		Voting: 		make(chan *Player),
	}
}