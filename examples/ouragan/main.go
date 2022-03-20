package main

import (
	"ego/engine/configuration"
	"ego/engine/context"
	"ego/engine/game"
)

func main() {
	ctx := context.CreateContext()
	context.RegisterContext("ouragan", ctx)
	cfg := configuration.CreateConfiguration("examples/ouragan/assets/config/")
	cfg.Init()
	context.Set("cfg", cfg.Get())
	game := game.CreateGame("text")
	game.Start()
}
