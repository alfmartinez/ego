package main

import (
	"ego/engine/configuration"
	"ego/engine/context"
	"ego/engine/game"
	"ego/engine/state"
	"ego/engine/template"
	"ego/examples/ouragan/internal/states"
	"ego/shared/game/text"
	"ego/shared/renderer/console"
)

func init() {
	text.Register()
	console.Register()
	state.RegisterStatesClosure("ouragan", states.States)
}

func main() {
	var g game.Game
	readConfiguration()

	g = game.CreateGame("text")
	g.Start()

}

func readConfiguration() {
	ctx := context.CreateContext()
	context.RegisterContext("ouragan", ctx)
	cfg := configuration.CreateConfiguration("examples/ouragan/assets/config/")
	cfg.Init()
	context.Set("cfg", cfg.Get())
	template.InitializeTemplates("examples/ouragan/assets/templates/*")
}
