package views

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func MessageList(g *gocui.Gui) error {
	if v, err := g.SetView("v2", 20, 4, 80, 16); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "View #2")
	}
	return nil
}