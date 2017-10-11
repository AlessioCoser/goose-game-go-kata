package goosegame

func NewBoard() *Board {
	board := &Board{}

	return board
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

