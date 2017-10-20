package game

import "strconv"

func NewDice(first int, second int) *Dice {
	return &Dice{first, second}
}

type Dice struct {
	first	int
	second	int
}

func (d *Dice) ToString() string {
	return strconv.Itoa(d.first) + ", " + strconv.Itoa(d.second)
}

func (d *Dice) Score() (int) {
	return d.first + d.second
}
