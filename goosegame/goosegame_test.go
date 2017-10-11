package goosegame

import (
	"testing"
)

type Player struct {
	Name string
}

func NewPlayer(name string) *Player {
	player := &Player{name}
	return player
}

type Board struct {
	players []*Player
}

func (b *Board) addPlayer (name string) {
	b.players = append(b.players, NewPlayer(name))
}

func (b *Board) Players() []Player {
	players := make([]Player, len(b.players))

	for index, player := range b.players {
		players[index] = *player
	}

	return players
}

func NewBoard() *Board {
	board := &Board{}

	return board
}

func TestAddPlayer(t *testing.T) {
	board := NewBoard()
	board.addPlayer("Pippo")

	actualPlayers := board.Players()
	expectedPlayer := Player{"Pippo"}

	if actualPlayers[0] != expectedPlayer {
		t.Errorf("Errore!")
	}
}
