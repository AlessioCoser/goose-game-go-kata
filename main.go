package main

import (
	"github.com/AlessioCoser/goose-game-go-kata/goosegame/cli"
)

func main() {
	gooseGame := cli.NewCli()
	gooseGame.Start()
}
