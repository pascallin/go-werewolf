package cui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/urfave/cli/v2"
	"strings"
)

func NewCommandResultPanel() *tview.TextView {
	commandResult := tview.NewTextView().
		SetText("==== command result ==== \n")
	return commandResult
}

func NewCommandInput(terminal *cli.App) *tview.InputField {
	commandInput := tview.NewInputField().
		SetLabel(" command >> ").
		SetPlaceholder("Please enter your command")

	commandInput.SetDoneFunc(func(key tcell.Key) {
		inputText := commandInput.GetText()
		if key == tcell.KeyEnter && len(inputText) > 0 {
			commandInput.SetText("")
			terminal.Run(strings.Fields("cmd " + inputText))
		}
	})
	return commandInput
}