package goosegame

func NewPlayer(name string) *Player {
	player := &Player{name}
	return player
}

type Player struct {
	Name string
}
