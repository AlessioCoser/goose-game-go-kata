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

	from, to := Cmd.board.MovePlayer(matches[0], [2]int{die1, die2})

	fromPos := strconv.Itoa(from)
	toPos := strconv.Itoa(to)

	if fromPos == "0" {
		fromPos = "Start"
	}

	return matches[0] + " rolls " + matches[1] +
		   ", " + matches[2] + ". " + matches[0] + " moves from " +
		   fromPos + " to " + toPos
}

