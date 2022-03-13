package main

import (
	_ "ego/internal/fyne"
	_ "ego/internal/glfw"
	_ "ego/internal/log"
	"ego/pkg/game"
	"runtime"
)

func main() {
	runtime.LockOSThread()
	game := game.CreateGame("file")
	game.Start()
}
