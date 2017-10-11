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
	expectedPlayer := Player{"Pippo"}

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