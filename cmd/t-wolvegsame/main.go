package main

import (
	"fmt"
	"github.com/pascallin/go-wolvesgame/internal/cli"
	cli2 "github.com/urfave/cli/v2"
	"log"
	"strings"

	"github.com/pascallin/go-wolvesgame/internal/app"
	"github.com/pascallin/go-wolvesgame/internal/terminalview"
)

func main() {

	terminalView, printPanel := terminalview.GetTerminalView()

	err, terminal := cli.CreateCliApp(printPanel)
	if err != nil {
		log.Fatal(err)
	}

	go listen(terminal)

	if err := terminalView.Run(); err != nil {
		panic(err)
	}
}

func listen(terminal *cli2.App) {
	for {
		select {
			case command := <- app.GetApp().CommandChan:
				err := terminal.Run(strings.Fields("cmd " + command))
				if err != nil {
					fmt.Println(err)
				}
		}
	}
}