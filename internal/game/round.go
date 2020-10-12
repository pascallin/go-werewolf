package game

import (
	"fmt"
)

func (g *Game) RoundStart(number int) {
	manLeft := g.GetManLeft()
	for i := range manLeft {
		manLeft[i].clearRoundVoted()
	}
	// NOTE: night actions
	g.Lifecycle <- WaitingNightPlayerAction
	fmt.Println("============ round ============", number)
	killed := <-g.PlayerActions.WerewolfKill
	killed.BeKilled()

	<-g.PlayerActions.SeerCheck
	poison := <-g.PlayerActions.UsePoison
	if poison != nil {
		poison.BeKilled()
	}
	antidote := <-g.PlayerActions.UseAntidote
	if antidote {
		killed.BeSaved()
	}
	// NOTE: day actions
	g.Lifecycle <- WaitingDayPlayerAction
	g.listenTalking()
	g.listenTVoting()
	p := g.GetMostRoundVotingPlayer()
	fmt.Println("exile player", p)
	p.Exile()
}

func (g *Game) listenTalking() {
	talking := g.GetManLeft()
	var talked = 0
	var talkingEnd = false
	for !talkingEnd {
		select {
		case playerTalked := <-g.PlayerActions.TalkedCount:
			talked += playerTalked
			if talked == len(talking) {
				talkingEnd = true
			}
		}
	}
}

func (g *Game) listenTVoting() {
	voting := g.GetManLeft()
	var voted = 0
	var voteEnd = false
	for !voteEnd {
		select {
		case player := <-g.PlayerActions.Voting:
			voted++
			player.RoundVoted++
			if voted == len(voting) {
				voteEnd = true
			}
		}
	}
}