package commands

import (
	"strings"

	"github.com/AlessioCoser/goose-game-go-kata/goosegame/game"
	"github.com/AlessioCoser/goose-game-go-kata/regexp"
)

func NewAddPlayerCmd(b *game.Board) *AddPlayerCmd{
	return &AddPlayerCmd{
		board: b,
		matchPattern: "add player ([a-zA-Z0-9]+)$",
	}
}

type AddPlayerCmd struct {
	board *game.Board
	matchPattern string
}

func (Cmd *AddPlayerCmd) CanHandle(command string) bool {
	_, found := regexp.Matches(Cmd.matchPattern, command)

	return found
}

func (Cmd *AddPlayerCmd) Execute(command string) string {
	matches, _ := regexp.Matches(Cmd.matchPattern, command)

	_ = Cmd.board.AddPlayer(matches[0])

	return "players: " + strings.Join(Cmd.board.PlayerNames(),", ")
}
