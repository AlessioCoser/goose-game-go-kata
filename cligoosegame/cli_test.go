package cligoosegame

import (
	"testing"
)

func TestAddPlayerCommand(t *testing.T) {
	cli := NewCli()
	response := cli.handle("add player Pippo")

	Equal(t, "players: Pippo", response)
}

func TestAddTwoPlayersCommand(t *testing.T) {
	cli := NewCli()
	_ = cli.handle("add player Pippo")
	response := cli.handle("add player Pluto")

	Equal(t, "players: Pippo, Pluto", response)
}

func Equal(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Expect %q to be %q", actual, expected)
	}
}
