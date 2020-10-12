package game

import "fmt"

func ai_werewolfAction(game *Game) {
	goodMan := game.GetGoodManLeft()
	ran, _ := genRandomList(len(goodMan))
	fmt.Println("WerewolfAction kill ------> ", goodMan[ran[0]])
	game.PlayerActions.WerewolfKill <- goodMan[ran[0]]
}

func ai_seerAction(game *Game) {
	man := game.GetManLeft()
	ran, _ := genRandomList(len(man))
	fmt.Println("Seer check ------> ", man[ran[0]].IsWerewolf())
	game.PlayerActions.SeerCheck <- man[ran[0]]
}

func ai_poisonAction(game *Game) {
	if game.RoundNumber == 1 {
		man := game.GetManLeft()
		ran, _ := genRandomList(len(man))
		fmt.Println("UsePoison kill ------> ", man[ran[0]])
		game.PlayerActions.UsePoison <- man[ran[0]]
	} else {
		game.PlayerActions.UsePoison <- nil
	}
}

func ai_antidoteAction(game *Game) {
	if game.RoundNumber == 2 {
		fmt.Println("UseAntidote save ------> ")
		game.PlayerActions.UseAntidote <- true
	} else {
		game.PlayerActions.UseAntidote <- false
	}
}

func ai_talking(game *Game) {
	for _, man := range game.GetManLeft() {
		game.PlayerActions.TalkedCount <- 1
		fmt.Println("-----" +man.Name + " talking end" + "-----")
	}
}

func ai_voting(game *Game) {
	for _, man := range game.GetManLeft() {
		people := game.GetManLeft()
		ran, _ := genRandomList(len(people))
		fmt.Println("-----" + man.Name + " vote " +  people[ran[0]].Name + "-----")
		game.PlayerActions.Voting <- people[ran[0]]
	}
}