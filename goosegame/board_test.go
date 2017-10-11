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

func TestMoveTwoPlayers(t *testing.T) {
	board := NewBoard()
	board.AddPlayer("Pippo")
	board.AddPlayer("Pluto")
	board.MovePlayer("Pippo", [2]int{4, 2})

	from, to := board.MovePlayer("Pippo", [2]int{1, 1})
	from_two, to_two := board.MovePlayer("Pluto", [2]int{1, 2})

	if from != 6 {
		t.Errorf("Player not moved from 6")
	}

	if to != 8 {
		t.Errorf("Player not moved to 8")
	}

	if from_two != 0 {
		t.Errorf("Player not moved from 0")
	}

	if to_two != 3 {
		t.Errorf("Player not moved to 3")
	}
}