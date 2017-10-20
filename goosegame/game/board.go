package game

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

func (b *Board) PlayerNames() []string {
	playerNames := make([]string, len(b.players))

	for index, player := range b.players {
		playerNames[index] = player.name
	}

	return playerNames
}

func (b *Board) MovePlayer(name string, dice *Dice) (*Move, error) {
	player := b.getPlayer(name)

	if b.Ended {
		return nil, errors.New("GAME_ENDED")
	}

	move := player.MoveBy(dice)

	if player.IsAt(b.winPosition) {
		b.endGame()
	}

	return move, nil
}

func (b *Board) WinnerIs() *Player {
	for _, player := range b.players {
		if player.IsAt(b.winPosition) {
			return player
		}
	}
	return nil
}

func (b *Board) WinnerName() string {
	winner := b.WinnerIs()

	if winner == nil {
		return ""
	}

	return b.WinnerIs().GetName()
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
