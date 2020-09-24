package cui

import (
	"github.com/rivo/tview"
)

func NewFlexLayout(commandInput *tview.InputField, commandResult *tview.TextView, messageList *tview.TextView, gamePanel *tview.Table) *tview.Flex {

	leftSide := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(gamePanel, 0, 1, false).
		AddItem(commandResult, 0, 2, false)
	mainContent := tview.NewFlex().
		AddItem(leftSide, 0, 1, false).
		AddItem(messageList, 0, 2, false)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainContent, 0, 10, false).
		AddItem(commandInput, 0, 1, false)

	return flex
}
