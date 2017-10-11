package goosegame

import (
	"errors"
)

func NewBoard() *Board {
	board := &Board{ make([]*Player, 0, 10) }

	return board
}

type Board struct {
	players []*Player
}

func (b *Board) AddPlayer (name string) error {
	player := NewPlayer(name)

	if b.playerExists(player) {
		return errors.New("ALREADY_EXISTS")
	}

	b.players = append(b.players, player)
	return nil
}

func (b *Board) Players() []Player {
	players := make([]Player, len(b.players))

	for index, player := range b.players {
		players[index] = *player
	}

	return players
}


func (b *Board) playerExists(player *Player) bool {
	exists := false

	for _, p := range b.players {
		if player.Name == p.Name {
			exists = true
		}
	}

	return exists
}
