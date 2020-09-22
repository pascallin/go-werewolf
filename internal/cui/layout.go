package cui

import (
	"github.com/jroimartin/gocui"

	"github.com/pascallin/go-wolvesgame/internal/cui/views"
)

func Layout(g *gocui.Gui) error {
	if err := views.HelpView(g); err != nil {
		return err
	}
	if err := views.Tab(g); err != nil {
		return err
	}
	return nil
}