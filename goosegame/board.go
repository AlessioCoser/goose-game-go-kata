package goosegame

import (
	"errors"
)

func NewBoard() *Board {
	board := &Board{ make([]*Player, 0, 10), 63, false }

	return board
}

type Board struct {
	players []*Player
	numSpaces int
	ended bool
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

	if b.ended {
		return player.Position, player.Position
	}

	from = player.Position

	player.Position = from + dice[0] + dice[1]

	if b.wins(player) {
		b.endGame()
	}

	return from, player.Position
}

func (b *Board) WinnerIs() *Player {
	for _, player := range b.players {
		if b.wins(player) {
			return player
		}
	}
	return nil
}

func (b *Board) wins(player *Player) bool {
	return player.Position == b.numSpaces
}

func (b *Board) getPlayer(playerName string) *Player {
	for _, p := range b.players {
		if playerName == p.Name {
			return p
		}
	}

	return nil
}

func (b *Board) endGame() {
	b.ended = true
}
