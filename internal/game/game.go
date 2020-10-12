package game

import (
	"encoding/json"
	"fmt"
)

type Status int
const (
	Waiting Status = iota		// wait for start
	Ready						// game started or restart
	WaitingNightPlayerAction
	WaitingDayPlayerAction
	Over						// game over
)
func (d Status) String() string {
	return [...]string{"Waiting", "Ready", "WaitingNightPlayerAction", "WaitingDayPlayerAction","Over"}[d]
}

type Game struct {
	PlayersCount  	int 			`json:"playersCount"` // game player needed
	Participants  	int				`json:"participants"`
	PlayerActions 	PlayerActions	`json:"-"`
	Players       	[]Player		`json:"players"`
	Lifecycle     	chan Status		`json:"-"`
	RoundNumber   	int				`json:"roundNumber"`
	ErrorCatch		chan error		`json:"-"`
}

var (
	Roles12 = []RoleType{
		Villager, Villager, Villager, Villager,
		Werewolf, Werewolf, Werewolf , Werewolf,
		Seer, Witch, Idiot, Hunter,
	}
)

func New() *Game {
	return &Game{
		PlayersCount:  	12,
		Participants:  	0,
		PlayerActions: 	NewActions(),
		Lifecycle:     	make(chan Status),
		RoundNumber: 	0,
	}
}

func (g *Game) GameStatusJSON() string {
	byteArray, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	return string(byteArray)
}