package cui

import (
	"github.com/rivo/tview"
)

func NewFlexLayout(commandInput *tview.InputField, commandResult *tview.TextView, messageList *tview.TextView) (*tview.Flex) {
	NewSection := func (item tview.Primitive) *tview.Grid {
		return tview.NewGrid().
			SetRows(3, 0, 3).
			SetColumns(30, 0, 30).
			AddItem(item, 0, 0, 3, 3, 1, 1, false)
	}

	flex := tview.NewFlex().
		AddItem(NewSection(commandInput), 0, 1, false).
		AddItem(NewSection(commandResult), 0, 1, false).
		AddItem(NewSection(messageList), 0, 1, false)

	return flex
}
