package goosegame

func NewPlayer(name string) *Player {
	player := &Player{name, 0}
	return player
}

type Player struct {
	Name string
	Position int
}

func (p *Player) MoveBy(dice [2]int) (from int, to int) {
	from = p.Position

	p.Position += dice[0] + dice[1]

	return from, p.Position
}
