package terminalview

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func NewFlexLayout(app *tview.Application) *tview.Flex {

	rightSide := tview.NewTextView().
		SetText("==== message start ==== ")

	var inputText string

	leftSide := tview.NewInputField().
		SetLabel(" Enter ").
		SetPlaceholder("Please enter your command").
		SetChangedFunc(func(text string) {
			inputText = text
		})

	doneInput := func () {
		rightSide.SetText(rightSide.GetText(false) + inputText)
		leftSide.SetText("")
	}

	leftSide.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			doneInput()
		}
	})

	flex := tview.NewFlex().
		AddItem(leftSide, 0, 1, false).
		AddItem(rightSide, 0, 2, false)
	return flex
}