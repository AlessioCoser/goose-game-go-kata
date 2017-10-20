package commands

import "github.com/AlessioCoser/goose-game-go-kata/goosegame/game"

type Command interface {
	CanHandle(string)bool
	Execute(string)string
}

func All(b *game.Board) []Command {
	return []Command{
		NewMovePlayerCmd(b),
		NewAddPlayerCmd(b),
		NewBaseCmd(),
	}
}
