package goosegame

func NewPlayer(name string) *Player {
	player := &Player{name, 0}
	return player
}

type Player struct {
	Name string
	Position int
}
