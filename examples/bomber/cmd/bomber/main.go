package main

import (
	"ego/engine/configuration"
	"ego/engine/context"
	"ego/engine/game"
	"ego/shared/display/rts"
	"ego/shared/object"
	"ego/shared/renderer/glfw"
	"ego/shared/state"
	"runtime"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
	rts.Register()
	glfw.Register()
	object.Register()
	state.Register()
}

func main() {
	ctx := context.CreateContext()
	context.RegisterContext("bomber", ctx)
	cfg := configuration.CreateConfiguration("examples/bomber/assets/config/")
	cfg.Init()
	context.Set("cfg", cfg.Get())
	game := game.CreateGame("viper")
	game.Start()
}
