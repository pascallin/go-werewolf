package main

import (
	"github.com/pascallin/go-wolvesgame/internal/terminalview"
)

func main() {
	if err := terminalview.GetTerminalView().Run(); err != nil {
		panic(err)
	}
}
