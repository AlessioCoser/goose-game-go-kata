package goosegame

import (
	"testing"
)

func TestAddPlayer(t *testing.T) {
	board := NewBoard()
	board.addPlayer("Pippo")

	actualPlayers := board.Players()
	expectedPlayer := Player{"Pippo"}

	if actualPlayers[0] != expectedPlayer {
		t.Errorf("Errore!")
	}
}
