package main

import (
	_ "ego/internal/renderer/glfw"
	"ego/pkg/configuration"
	"ego/pkg/context"
	"ego/pkg/game"
	"runtime"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	ctx := context.CreateContext()
	context.RegisterContext("ego", ctx)
	cfg := configuration.CreateConfiguration()
	cfg.Init()
	ctx.Set("cfg", cfg.Get())
	game := game.CreateGame("viper")
	game.Start()
}
