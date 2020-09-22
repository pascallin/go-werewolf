package terminalview

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func NewFlexLayout(app *tview.Application) *tview.Flex {

	rightSide := tview.NewTextView().
		SetText("==== message start ====")

	var inputText string

	leftSide := tview.NewInputField().
		SetLabel(" command >> ").
		SetPlaceholder("Please enter your command").
		SetFieldBackgroundColor(tcell.Color16).
		SetChangedFunc(func(text string) {
			inputText = text
		})

	doneInput := func () {
		// TODO: gap need to be fixed if using grid border
		// rightSide.SetText(rightSide.GetText(false) + "  " + inputText)
		rightSide.SetText(rightSide.GetText(false) + inputText)
		leftSide.SetText("")
	}

	leftSide.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			doneInput()
		}
	})

	NewSection := func (item tview.Primitive) *tview.Grid {
		return tview.NewGrid().
			SetRows(3, 0, 3).
			SetColumns(30, 0, 30).
			// TODO: gap need to be fixed if using grid border
			//SetBorders(true).
			AddItem(item, 0, 0, 3, 3, 1, 1, false)
	}

	flex := tview.NewFlex().
		AddItem(NewSection(leftSide), 0, 1, false).
		AddItem(NewSection(rightSide), 0, 2, false)
	return flex
}