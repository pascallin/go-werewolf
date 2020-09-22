package views

import (
	"github.com/jroimartin/gocui"
)

func TextInput(g *gocui.Gui) error {
	if v, err := g.SetView("v3", 30, 6, 100, 20); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Wrap = true
		if _, err := g.SetCurrentView("v3"); err != nil {
			return err
		}
	}
	return nil
}