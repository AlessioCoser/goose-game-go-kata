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

func Equal(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Expect %q to be %q", actual, expected)
	}
}
