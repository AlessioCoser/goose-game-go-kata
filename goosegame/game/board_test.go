package game

import (
	"testing"
)

func TestAddPlayer(t *testing.T) {
	board := NewBoard()
	err := board.AddPlayer("Pippo")

	Equal(t, nil, err)
	Equal(t, Player{"Pippo", 0}, board.Players()[0])
}

func TestDuplicatePlayer(t *testing.T) {
	board := NewBoard()
	board.AddPlayer("Pippo")

	err := board.AddPlayer("Pippo")

	NotEqual(t, nil, err)
}

func TestMovePlayerFromStart(t *testing.T) {
	board := NewBoard()
	board.AddPlayer("Pippo")

	move, _ := board.MovePlayer("Pippo", NewDice(4, 2))

	Equal(t, 0, move.From)
	Equal(t, 6, move.To)
}

func TestMoveTwoPlayers(t *testing.T) {
	board := NewBoard()
	board.AddPlayer("Pippo")
	board.AddPlayer("Pluto")
	board.MovePlayer("Pippo", NewDice(4, 2))

	move, _ := board.MovePlayer("Pippo", NewDice(1, 1))
	move_two, _ := board.MovePlayer("Pluto", NewDice(1, 2))

	Equal(t, 6, move.From)
	Equal(t, 8, move.To)
	Equal(t, 0, move_two.From)
	Equal(t, 3, move_two.To)
}

func TestWinner(t *testing.T) {
	board := NewBoard()
	board.AddPlayer("Pippo")
	board.MovePlayer("Pippo", NewDice(59, 1))
	move, _ := board.MovePlayer("Pippo", NewDice(1, 2))

	Equal(t, 63, move.To)
	NotEqual(t, nil, board.WinnerIs())
	Equal(t, "Pippo", board.WinnerName())
}


func TestFirstWinner(t *testing.T) {
	board := NewBoard()
	board.AddPlayer("Pippo")
	board.AddPlayer("Winner")
	board.MovePlayer("Pippo", NewDice(60, 1))

	move_win, _ := board.MovePlayer("Winner", NewDice(60, 3))
	_, err := board.MovePlayer("Pippo", NewDice(1, 1))

	Equal(t, "GAME_ENDED", err.Error())
	Equal(t, 63, move_win.To)
	Equal(t, "Winner", board.WinnerName())
}

func Equal(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Expect %q to be %q", actual, expected)
	}
}

func NotEqual(t *testing.T, expected interface{}, actual interface{}) {
	if actual == expected {
		t.Errorf("Expect %q to be different from %q", actual, expected)
	}
}