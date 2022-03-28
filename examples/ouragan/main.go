package main

import (
	"github.com/alfmartinez/ego/engine/configuration"
	"github.com/alfmartinez/ego/engine/context"
	"github.com/alfmartinez/ego/engine/game"
	"github.com/alfmartinez/ego/engine/state"
	"github.com/alfmartinez/ego/engine/template"
	"github.com/alfmartinez/ego/examples/ouragan/internal/states"
	"github.com/alfmartinez/ego/shared/game/text"
	"github.com/alfmartinez/ego/shared/renderer/console"
)

func init() {
	context.CreateAndRegisterContext("ouragan")
	cfg := configuration.CreateConfiguration("examples/ouragan/assets/config/")
	cfg.Init()
	context.Set("cfg", cfg.Get())
	template.InitializeTemplates("examples/ouragan/assets/templates/*")
	text.Register()
	console.Register()
	state.RegisterStatesClosure("ouragan", states.States)
}

func main() {
	var g game.Game

	g = game.CreateGame("text")
	g.Start()

}
