package game

func GetGoodManLeft(game *Game) []*Player {
	var ps []*Player
	for i, player := range game.Players {
		if !player.IsWerewolf() && player.IsAlive() {
			ps = append(ps, &game.Players[i])
		}
	}
	return ps
}

func GetManLeft(game *Game) []*Player {
	var ps []*Player
	for i, player := range game.Players {
		if player.IsAlive() {
			ps = append(ps, &game.Players[i])
		}
	}
	return ps
}