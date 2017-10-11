package main

import (
	"github.com/AlessioCoser/goose-game-go-kata/goosegame"
)

func main() {
	gooseGame := goosegame.NewCli()
	gooseGame.Start()
}
