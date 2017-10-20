package game

func NewPlayer(name string) *Player {
	player := &Player{name, 0}
	return player
}

type Player struct {
	name     string
	position int
}

func (p *Player) MoveBy(dice *Dice) (*Move) {
	from := p.position
	p.position += dice.Score()

	return NewMove(from, p.position)
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
