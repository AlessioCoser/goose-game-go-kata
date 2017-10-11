package goosegame

import (
	"bufio"
	"os"
	"fmt"
	"github.com/AlessioCoser/goose-game-go-kata/goosegame/commands"
)

func NewCli() *Cli {
	cliGame := &Cli{ NewBoard() }

	return cliGame
}

type Cli struct {
	board *Board
}

func (game *Cli) Handle(text string) string {
	cmds := commands.All()

	for _, command := range cmds {
		if command.CanHandle(text) {
			return command.Execute(text)
		}
	}

	return ""
}

func (game *Cli) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(game.Handle(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}