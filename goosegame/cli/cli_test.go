package cli

import (
	"testing"
)

func TestAddPlayerCommand(t *testing.T) {
	game := NewCli()
	response := game.handle("add player Pippo")

	Equal(t, "players: Pippo", response)
}

func TestAddTwoPlayersCommand(t *testing.T) {
	game := NewCli()
	_ = game.handle("add player Pippo")
	response := game.handle("add player Pluto")

	Equal(t, "players: Pippo, Pluto", response)
}

func TestAddAlreadyExistingPlayerCommand(t *testing.T) {
	game := NewCli()
	_ = game.handle("add player Pippo")
	response := game.handle("add player Pippo")

	Equal(t, "Pippo: already existing player", response)
}

func TestMoveAPlayerFromStart(t *testing.T) {
	game := NewCli()
	_ = game.handle("add player Pippo")
	_ = game.handle("add player Pluto")

	response := game.handle("move Pippo 4, 2")

	Equal(t, "Pippo rolls 4, 2. Pippo moves from Start to 6", response)
}

func TestMoveTwoPlayers(t *testing.T) {
	game := NewCli()
	_ = game.handle("add player Pippo")
	_ = game.handle("add player Pluto")
	_ = game.handle("move Pippo 4, 2")
	_ = game.handle("move Pluto 1, 1")

	response := game.handle("move Pluto 4, 6")

	Equal(t, "Pluto rolls 4, 6. Pluto moves from 2 to 12", response)
}

func TestPlayerWin(t *testing.T) {
	game := NewCli()
	_ = game.handle("add player Pippo")
	_ = game.handle("move Pippo 6, 6")
	_ = game.handle("move Pippo 6, 6")
	_ = game.handle("move Pippo 6, 6")
	_ = game.handle("move Pippo 6, 6")
	_ = game.handle("move Pippo 6, 6")

	response := game.handle("move Pippo 1, 2")

	Equal(t, "Pippo rolls 1, 2. Pippo moves from 60 to 63. Pippo Wins!!", response)
}

func Equal(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("\nExpect %q \nActual %q", actual, expected)
	}
}
