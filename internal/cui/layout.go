package cui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func NewFlexLayout(tviwe *tview.Application, commandChan chan string) (*tview.Flex, *tview.TextView) {

	NewSection := func (item tview.Primitive) *tview.Grid {
		return tview.NewGrid().
			SetRows(3, 0, 3).
			SetColumns(30, 0, 30).
			// TODO: gap need to be fixed if using grid border
			//SetBorders(true).
			AddItem(item, 0, 0, 3, 3, 1, 1, false)
	}

	rightSide := tview.NewTextView().
		SetText("==== message start ==== \n")

	commandInput := tview.NewInputField().
		SetLabel(" command >> ").
		SetPlaceholder("Please enter your command").
		SetFieldBackgroundColor(tcell.Color16)

	commandInput.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter && len(commandInput.GetText()) > 0 {
			commandChan <- commandInput.GetText()
			commandInput.SetText("")
		}
	})

	flex := tview.NewFlex().
		AddItem(NewSection(commandInput), 0, 1, false).
		AddItem(NewSection(rightSide), 0, 2, false)

	return flex, rightSide
}