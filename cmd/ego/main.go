package main

import (
	"ego/pkg/game"
)

func main() {
	game, err := game.CreateGame("game.yml")
	if err != nil {
		panic(err)
	}
	game.Loop()
}
