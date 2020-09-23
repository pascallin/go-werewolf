package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/urfave/cli/v2"

	cliinterface "github.com/pascallin/go-wolvesgame/internal/cli"
	"github.com/pascallin/go-wolvesgame/internal/cui"
)

func main() {

	commandChan := make(chan string)

	terminalView, printPanel := cui.GetTerminalView(commandChan)
	err, terminal := cliinterface.CreateCliApp(printPanel)
	if err != nil {
		log.Fatal(err)
	}
	go listen(terminal, commandChan)

	if err := terminalView.Run(); err != nil {
		panic(err)
	}
}

func listen(terminal *cli.App, printChan chan string) {
	for {
		select {
			case command := <- printChan:
				err := terminal.Run(strings.Fields("cmd " + command))
				if err != nil {
					fmt.Println(err)
				}
		}
	}
}