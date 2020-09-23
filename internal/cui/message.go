package cui

import "github.com/rivo/tview"

func NewMessagePanel() *tview.TextView {
	messageList := tview.NewTextView().
		SetText("==== message list ==== \n")
	return messageList
}