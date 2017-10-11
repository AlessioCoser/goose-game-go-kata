package goosegame

import (
	"errors"
)

func NewBoard() *Board {
	board := &Board{ make([]*Player, 0, 10), 63, false }

	return board
}

type Board struct {
	players     []*Player
	winPosition int
	Ended       bool
}

func (b *Board) AddPlayer (name string) error {
	if b.getPlayer(name) != nil {
		return errors.New("ALREADY_EXISTS")
	}

	player := NewPlayer(name)
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

	if b.Ended {
		return player.GetPosition(), player.GetPosition()
	}

	from, to = player.MoveBy(dice)

	if player.IsAt(b.winPosition) {
		b.endGame()
	}

	return from, to
}

func (b *Board) WinnerIs() *Player {
	for _, player := range b.players {
		if player.IsAt(b.winPosition) {
			return player
		}
	}
	return nil
}

func (b *Board) getPlayer(playerName string) *Player {
	for _, p := range b.players {
		if playerName == p.GetName() {
			return p
		}
	}

	return nil
}

func (b *Board) endGame() {
	b.Ended = true
}
