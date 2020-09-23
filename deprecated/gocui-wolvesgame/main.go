package main

import (
	"log"
	"runtime"

	"github.com/jroimartin/gocui"
	"github.com/mattn/go-runewidth"

	"github.com/pascallin/go-wolvesgame/deprecated/cui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	// NOTE: reference: https://github.com/jroimartin/gocui/issues/217
	if runtime.GOOS == "windows" && runewidth.IsEastAsian() {
		g.ASCII = true
	}
	g.Cursor = true
	g.Mouse = false

	g.SetManagerFunc(gocui.Layout)

	if err := gocui.Keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}