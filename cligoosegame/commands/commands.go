package commands

import "github.com/AlessioCoser/goose-game-go-kata/goosegame"

type Command interface {
	CanHandle(string)bool
	Execute(string)string
}

func All(b *goosegame.Board) []Command {
	return []Command{
		NewAddPlayerCmd(b),
		NewBaseCmd(),
	}
}
