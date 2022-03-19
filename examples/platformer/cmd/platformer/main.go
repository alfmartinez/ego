package main

import (
	"ego/engine/configuration"
	"ego/engine/context"
	"ego/engine/game"
	"ego/shared/display/rts"
	"ego/shared/renderer/glfw"
	"runtime"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
	rts.Register()
	glfw.Register()
}

func main() {
	ctx := context.CreateContext()
	context.RegisterContext("platformer", ctx)
	cfg := configuration.CreateConfiguration("examples/platformer/assets/config/")
	cfg.Init()
	context.Set("cfg", cfg.Get())
	game := game.CreateGame("viper")
	game.Start()
}
