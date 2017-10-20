package main

import (
	"github.com/AlessioCoser/goose-game-go-kata/cligoosegame"
)

func main() {
	gooseGame := cligoosegame.NewCli()
	gooseGame.Start()
}
