package game

import "fmt"

type WaitPlayerActions int
const (
	WerewolfAction WaitPlayerActions = iota
	WitchAction
	SeerAction
	Voting
)
func (d WaitPlayerActions) String() string {
	return [...]string{"WerewolfAction", "WitchAction", "SeerAction", "Voting"}[d]
}

func RoundStart(number int, game Game) {
	fmt.Println("============ round ============", number)
	// night fall
	fmt.Println("night fall")
	NightFall(game)
	// day
	Sunrise(game)
	fmt.Println("sunrise")
}

func NightFall(game Game) {
	game.BlockGame()
	ch := make(chan WaitPlayerActions)
	var result int
	// wait werewolf action
	go func() {
		println("come into WerewolfAction")
		ch <- WerewolfAction
	}()
	<- ch
	// wait seer action
	go func() {
		println("come into SeerAction")
		ch <- SeerAction
	}()
	<- ch
	close(ch)
	println("result is:", result)
}

func Sunrise(game Game) {
	// show dead man
	// random pick speaker
	// vote
}