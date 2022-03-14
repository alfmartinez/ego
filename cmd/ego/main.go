package main

import (
	_ "ego/internal/fyne"
	_ "ego/internal/glfw"
	_ "ego/internal/log"
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
