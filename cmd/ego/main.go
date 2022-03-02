package main

import "ego/pkg/game"

func main() {
	game, err := game.CreateGame("file")
	if err != nil {
		panic(err)
	}

	game.Start()
}
