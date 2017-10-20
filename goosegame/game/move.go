package game

import "strconv"

func NewMove(from int, to int) *Move {
	return &Move{from, to}
}

type Move struct {
	From	int
	To 		int
}

func (m *Move) ToString() string {
	from := strconv.Itoa(m.From)

	if m.From == 0 {
		from = "Start"
	}

	return "from " + from + " to " + strconv.Itoa(m.To)
}