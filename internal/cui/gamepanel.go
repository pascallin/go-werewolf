package cui

import (
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/pascallin/go-wolvesgame/internal/game"
)

type GamePanel struct {
	Table *tview.Table
	offset int
	players string
}

var (
	NewTitleCell = func(value string) *tview.TableCell {
		return tview.NewTableCell(value).SetTextColor(tcell.ColorOrange).SetAlign(tview.AlignCenter)
	}
	NewValueCell = func(value string) *tview.TableCell {
		return tview.NewTableCell(value).SetTextColor(tcell.ColorOrange).SetAlign(tview.AlignCenter)
	}
)

func NewGamePanel() *GamePanel {

	table := tview.NewTable().SetBorders(true)

	table.SetCell(0, 0, NewTitleCell("Game:"))
	table.SetCell(0, 1, NewValueCell("Status"))
	table.SetCell(1, 1, NewValueCell("Non"))

	table.SetCell(3, 0, NewTitleCell("Players:"))
	var rowOffset = 4

	header := strings.Split("ID Name Status", " ")
	for i := range header {
		table.SetCell(rowOffset, i,
			tview.NewTableCell(header[i]).
				SetTextColor(tcell.ColorYellow).
				SetAlign(tview.AlignCenter))
	}

	return &GamePanel{
		Table: table,
		offset: rowOffset + 1,
	}
}

func (gp *GamePanel) UpdateGameStatus(status game.Status) {
	gp.Table.SetCell(1, 1, NewValueCell(status.String()))
}

/**
playersData: column data sep with " " and row data sep with ","
example:
	NewGamePanel("1 Player1 Dead,2 Player2 Alive,3 Test Talking")
*/
func (gp *GamePanel) AddPlayerToGamePanel(playerData string) {
	row := strings.Split(playerData, " ")
	for c, v := range row {
		gp.Table.SetCell(1 + gp.offset, c,
			tview.NewTableCell(v).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter))
	}
	gp.offset++
}