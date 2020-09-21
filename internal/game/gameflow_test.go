package game

import (
	"fmt"
	"github.com/pascallin/go-wolvesgame/internal/game/player"
	uuid "github.com/satori/go.uuid"
	"os"
	"os/signal"
	"strconv"
	"testing"
)

func genPlayers(count int) []player.Player {
	var players []player.Player
	for i := 0; i < count; i++ {
		players = append(players, player.NewPlayer(uuid.NewV4(), "player"+strconv.Itoa(i)))
	}
	return players
}

func randomKill() {

}

func TestGameFlow(t *testing.T) {
	game := New()

	players := genPlayers(game.PlayersCount)
	for _, p := range players {
		game.JoinPlayer(p)
	}

	StartGame(game)

	// wait game player actions
	signalListen()
}

func signalListen() {
	c := make(chan os.Signal)
	signal.Notify(c)
	for {
		s := <-c
		//收到信号后的处理，这里只是输出信号内容，可以做一些更有意思的事
		fmt.Println("get signal:", s)
		os.Exit(1)
	}
}