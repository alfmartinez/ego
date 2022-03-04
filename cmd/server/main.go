package main

import "ego/pkg/game"

func main() {
	game := game.CreateGame("backend")
	game.Start()
}
