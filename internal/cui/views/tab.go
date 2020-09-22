package views

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func Tab(g *gocui.Gui) error {
	if v, err := g.SetView("v1", 10, 2, 30, 6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "View #1")
	}
	if v, err := g.SetView("v2", 20, 4, 40, 8); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "View #2")
	}
	if v, err := g.SetView("v3", 30, 6, 50, 10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "View #3")
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
		return err
	})
	if err != nil {
		return err
	}

	err = g.SetKeybinding("", '3', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		_, err := g.SetViewOnTop("v3")
		return err
	})
	return nil
}