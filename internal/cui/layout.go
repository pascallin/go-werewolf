package cui

import (
	"github.com/rivo/tview"
)

func NewFlexLayout(commandInput *tview.InputField, commandResult *tview.TextView, messageList *tview.TextView) (*tview.Flex) {

	gamePanel := NewGamePanel()

	leftSide := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(commandInput, 0, 1, false).
		AddItem(gamePanel.table, 0, 1, false)

	flex := tview.NewFlex().
		AddItem(leftSide, 0, 1, false).
		AddItem(commandResult, 0, 1, false).
		AddItem(messageList, 0, 1, false)

	return flex
}
