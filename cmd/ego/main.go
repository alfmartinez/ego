package main

import (
	"ego/pkg/game/factory"
)

func main() {
	game, err := factory.CreateGame("game.yml")
	if err != nil {
		panic(err)
	}

	game.Start()
}
