package goosegame

import (
	"testing"
)

func TestAddPlayer(t *testing.T) {
	board := NewBoard()
	err := board.AddPlayer("Pippo")

	if err != nil {
		t.Errorf("%q", err)
	}

	actualPlayers := board.Players()
	expectedPlayer := Player{"Pippo", 0}

	if actualPlayers[0] != expectedPlayer {
		t.Errorf("Player not added")
	}
}

func TestDuplicatePlayer(t *testing.T) {
	board := NewBoard()
	board.AddPlayer("Pippo")

	err := board.AddPlayer("Pippo")

	if err == nil {
		t.Errorf("no error suppled")
	}
}

func TestMovePlayerFromStart(t *testing.T) {
	board := NewBoard()
	board.AddPlayer("Pippo")

	from, to := board.MovePlayer("Pippo", [2]int{4, 2})

	if from != 0 {
		t.Errorf("Player not moved from Start")
	}

	if to != 6 {
		t.Errorf("Player not moved to 6")
	}
}