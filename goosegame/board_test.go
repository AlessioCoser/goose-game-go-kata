package goosegame

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

	from, to := board.MovePlayer("Pippo", [2]int{4, 2})

	Equal(t, 0, from)
	Equal(t, 6, to)
}

func TestMoveTwoPlayers(t *testing.T) {
	board := NewBoard()
	board.AddPlayer("Pippo")
	board.AddPlayer("Pluto")
	board.MovePlayer("Pippo", [2]int{4, 2})

	from, to := board.MovePlayer("Pippo", [2]int{1, 1})
	from_two, to_two := board.MovePlayer("Pluto", [2]int{1, 2})

	Equal(t, 6, from)
	Equal(t, 8, to)
	Equal(t, 0, from_two)
	Equal(t, 3, to_two)
}

func TestWinner(t *testing.T) {
	board := NewBoard()
	board.AddPlayer("Pippo")
	board.MovePlayer("Pippo", [2]int{59, 1})
	_, to := board.MovePlayer("Pippo", [2]int{1, 2})

	Equal(t, 63, to)
	NotEqual(t, nil, board.WinnerIs())
	Equal(t, "Pippo", board.WinnerIs().GetName())
}


func TestFirstWinner(t *testing.T) {
	board := NewBoard()
	board.AddPlayer("Pippo")
	board.AddPlayer("Winner")
	board.MovePlayer("Pippo", [2]int{60, 1})

	_, to_winner := board.MovePlayer("Winner", [2]int{60, 3})
	_, to_pippo := board.MovePlayer("Pippo", [2]int{1, 1})

	Equal(t, 63, to_winner)
	Equal(t, 61, to_pippo)
	Equal(t, "Winner", board.WinnerIs().GetName())
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