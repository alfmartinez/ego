package main

import "ego/pkg/game"

func main() {
	game := game.CreateGame("file")
	game.Start()
}
