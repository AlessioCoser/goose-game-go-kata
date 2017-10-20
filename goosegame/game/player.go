package game

func NewPlayer(name string) *Player {
	player := &Player{name, 0}
	return player
}

type Player struct {
	name     string
	position int
}

func (p *Player) MoveBy(dice [2]int) (from int, to int) {
	from = p.position

	p.position += dice[0] + dice[1]

	return from, p.position
}

func (p *Player) IsAt(pos int) bool {
	return p.position == pos
}

func (p *Player) GetPosition() int {
	return p.position
}

func (p *Player) GetName() string {
	return p.name
}
