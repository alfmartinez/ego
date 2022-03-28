package main

import (
	"github.com/alfmartinez/ego/engine/configuration"
	"github.com/alfmartinez/ego/engine/context"
	"github.com/alfmartinez/ego/engine/game"
	"github.com/alfmartinez/ego/shared/display/rts"
	"github.com/alfmartinez/ego/shared/object"
	"github.com/alfmartinez/ego/shared/renderer/glfw"
	"github.com/alfmartinez/ego/shared/state"
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
