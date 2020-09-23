package cui

import "github.com/rivo/tview"

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