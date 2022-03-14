package main

import (
	_ "ego/internal/renderer/glfw"
	"ego/pkg/game"
	"runtime"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	game := game.CreateGame("file")
	game.Start()
}
