package views

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func Tab(g *gocui.Gui) error {
	if v, err := g.SetView("v1", 10, 2, 60, 12); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "View #1")
	}
	if err := MessageList(g); err != nil {
		return err
	}
	if err := TextInput(g); err != nil {
		return err
	}

	if err := bindingTabKeys(g); err != nil {
		return err
	}

	return nil
}

func bindingTabKeys(g *gocui.Gui) error {
	var err error
	err = g.SetKeybinding("", '1', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		_, err := g.SetViewOnTop("v1")
		return err
	})
	if err != nil {
		return err
	}

	err = g.SetKeybinding("", '2', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		_, err := g.SetViewOnTop("v2")
		if _, err := g.SetCurrentView("v2"); err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return err
	}

	err = g.SetKeybinding("", '3', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		_, err := g.SetViewOnTop("v3")
		if _, err := g.SetCurrentView("v3"); err != nil {
			return err
		}
		return err
	})
	return nil
}