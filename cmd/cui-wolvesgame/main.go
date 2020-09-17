package main

import (
	"log"
	"runtime"

	"github.com/jroimartin/gocui"
	"github.com/mattn/go-runewidth"

	"github.com/pascallin/go-wolvesgame/internal/cui"
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

	g.SetManagerFunc(cui.Layout)

	if err := cui.Keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}