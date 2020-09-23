package cui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func NewFlexLayout(commandChan chan string) (*tview.Flex, *tview.TextView) {

	NewSection := func (item tview.Primitive) *tview.Grid {
		return tview.NewGrid().
			SetRows(3, 0, 3).
			SetColumns(30, 0, 30).
			AddItem(item, 0, 0, 3, 3, 1, 1, false)
	}

	messageList := tview.NewTextView().
		SetText("==== message list ==== \n")

	commandResult := tview.NewTextView().
		SetText("==== command result ==== \n")

	commandInput := tview.NewInputField().
		SetLabel(" command >> ")
		//SetPlaceholder("Please enter your command")

	form := tview.NewForm().AddFormItem(commandInput)

	commandInput.SetDoneFunc(func(key tcell.Key) {
		inputText := commandInput.GetText()
		if key == tcell.KeyEnter && len(inputText) > 0 {
			// for test
			if  commandInput.GetText() == "cmd:remove" {
				removePlayerAction(form)
				return
			} else if commandInput.GetText() == "cmd:add" {
				addPlayerAction(form, func(label string) {
					messageList.SetText("kill: " + label)
				})
				return
			}
			commandChan <- inputText
			commandInput.SetText("")
		}
	})

	flex := tview.NewFlex().
		AddItem(NewSection(form), 0, 1, false).
		AddItem(NewSection(commandResult), 0, 1, false).
		AddItem(NewSection(messageList), 0, 1, false)

	return flex, commandResult
}

func addPlayerAction(form *tview.Form, selectFn func(label string)) {
	var selected string

	// player dropdown picker
	dropdown := tview.NewDropDown().
		SetLabel("Select player to kill: ").
		SetOptions([]string{"First", "Second", "Third", "Fourth", "Fifth"}, nil).
		SetSelectedFunc(func(text string, index int) {
			selected = text
		})

	form.
		AddFormItem(dropdown).
		AddButton("confirm", func() {
			selectFn(selected)
		})
}

func removePlayerAction(form *tview.Form) {
	form.
		RemoveFormItem(form.GetFormItemCount() - 1).
		RemoveButton(0)
}