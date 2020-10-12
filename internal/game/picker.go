package game

func (game *Game) GetGoodManLeft() []*Player {
	var ps []*Player
	for i, player := range game.Players {
		if !player.IsWerewolf() && player.IsAlive() {
			ps = append(ps, &game.Players[i])
		}
	}
	return ps
}

func (game *Game) GetManLeft() []*Player {
	var ps []*Player
	for i, player := range game.Players {
		if player.IsAlive() {
			ps = append(ps, &game.Players[i])
		}
	}
	return ps
}

func (game *Game) GetMostRoundVotingPlayer() *Player {
	var p *Player
	for i, player := range game.Players {
		if player.IsAlive() {
			if p == nil {
				p = &game.Players[i]
				continue
			}
			if p.RoundVoted < game.Players[i].RoundVoted {
				p = &game.Players[i]
				continue
			}
		}
	}
	return p
}