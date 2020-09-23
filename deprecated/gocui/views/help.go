package views

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func HelpView(g *gocui.Gui) error {
	maxX, _ := g.Size()

	if v, err := g.SetView("help", maxX-23, 0, maxX-1, 5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "KEYBINDINGS")
		fmt.Fprintln(v, "[1,2,3]: Change View")
		fmt.Fprintln(v, "^C: Exit")
	}
	return nil
}