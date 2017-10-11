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

func TestWinner(t *testing.T) {
	board := NewBoard()
	board.AddPlayer("Pippo")
	board.MovePlayer("Pippo", [2]int{6, 6})
	board.MovePlayer("Pippo", [2]int{6, 6})
	board.MovePlayer("Pippo", [2]int{6, 6})
	board.MovePlayer("Pippo", [2]int{6, 6})
	board.MovePlayer("Pippo", [2]int{6, 6})
	_, to := board.MovePlayer("Pippo", [2]int{1, 2})

	if to != 63 {
		t.Errorf("Player not moved to 63, instead moved to %d", to)
	}

	if board.WinnerIs() == nil {
		t.Errorf("Player not won")
	}

	if board.WinnerIs().Name != "Pippo" {
		t.Errorf("Player Pippo not won")
	}
}


func TestFirstWinner(t *testing.T) {
	board := NewBoard()
	board.AddPlayer("Pippo")
	board.AddPlayer("Winner")
	board.MovePlayer("Pippo", [2]int{60, 1})

	_, to_winner := board.MovePlayer("Winner", [2]int{60, 3})
	_, to_pippo := board.MovePlayer("Pippo", [2]int{1, 1})

	if to_winner != 63 {
		t.Errorf("Player Winner not moved to 63, instead moved to %d", to_winner)
	}

	if to_pippo != 61 {
		t.Errorf("Player Pippo not moved to 61, instead moved to %d", to_pippo)
	}

	if board.WinnerIs().Name != "Winner" {
		t.Errorf("Player Pippo should not win")
	}
}