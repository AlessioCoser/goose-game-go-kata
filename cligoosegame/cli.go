package cligoosegame

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AlessioCoser/goose-game-go-kata/cligoosegame/commands"
	"github.com/AlessioCoser/goose-game-go-kata/goosegame"
)

func NewCli() *Cli {
	return &Cli{goosegame.NewBoard()}
}

type Cli struct {
	board *goosegame.Board
}

func (game *Cli) handle(text string) string {
	for _, command := range game.getCommands() {
		if command.CanHandle(text) {
			return command.Execute(text)
		}
	}

	return ""
}

func (game *Cli) getCommands() []commands.Command {
	return commands.All(game.board)
}

func (game *Cli) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(game.handle(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
