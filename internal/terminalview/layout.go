package terminalview

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/pascallin/go-wolvesgame/internal/app"
	"github.com/rivo/tview"
)

func NewFlexLayout(tviwe *tview.Application) (*tview.Flex, *tview.TextView) {

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

	var inputText string

	commandInput := tview.NewInputField().
		SetLabel(" command >> ").
		SetPlaceholder("Please enter your command").
		SetFieldBackgroundColor(tcell.Color16).
		SetChangedFunc(func(text string) {
			inputText = text
		})
	//leftSide := tview.NewGrid().
	//	SetRows(3, 0, 3).
	//	SetColumns(30, 0, 30).
	//	AddItem(command, 0, 0, 1, 3, 0, 0, false).
	//	AddItem(NewSection(rightSide), 3, 0, 1, 3, 0, 0, false)

	commandInput.SetDoneFunc(func(key tcell.Key) {
		fmt.Println("here")
		if key == tcell.KeyEnter && len(inputText) > 0 {
			app.GetApp().CommandChan <- inputText
			inputText = ""
			commandInput.SetText("")
		}
	})

	flex := tview.NewFlex().
		AddItem(NewSection(commandInput), 0, 1, false).
		AddItem(NewSection(rightSide), 0, 2, false)

	return flex, rightSide
}