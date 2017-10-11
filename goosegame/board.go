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

	if b.getPlayer(player.Name) != nil {
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

func (b *Board) MovePlayer(name string, dice [2]int) (from int, to int) {
	player := b.getPlayer(name)

	from = player.Position
	to = from + dice[0] + dice[1]

	player.Position = dice[0] + dice[1]
	return
}


func (b *Board) getPlayer(playerName string) *Player {
	for _, p := range b.players {
		if playerName == p.Name {
			return p
		}
	}

	return nil
}
