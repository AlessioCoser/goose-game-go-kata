package commands

import (
	"strconv"

	"github.com/AlessioCoser/goose-game-go-kata/goosegame/game"
	"github.com/AlessioCoser/goose-game-go-kata/regexp"
)

func NewMovePlayerCmd(b *game.Board) *MovePlayerCmd{
	return &MovePlayerCmd{
		board: b,
		matchPattern: "move ([a-zA-Z0-9]+) ([0-6]), ([0-6])$",
	}
}

type MovePlayerCmd struct {
	board *game.Board
	matchPattern string
}

func (Cmd *MovePlayerCmd) CanHandle(command string) bool {
	_, found := regexp.Matches(Cmd.matchPattern, command)

	return found
}

func (Cmd *MovePlayerCmd) Execute(command string) string {
	matches, _ := regexp.Matches(Cmd.matchPattern, command)

	die1, _ := strconv.Atoi(matches[1])
	die2, _ := strconv.Atoi(matches[2])
	dice := game.NewDice(die1, die2)

	move, _ := Cmd.board.MovePlayer(matches[0], dice)

	return matches[0] + " rolls " + dice.ToString() + ". " + matches[0] + " moves " + move.ToString()
}

