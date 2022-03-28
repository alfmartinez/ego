package text

import (
	"github.com/alfmartinez/ego/engine/configuration"
	"github.com/alfmartinez/ego/engine/context"
	"github.com/alfmartinez/ego/engine/game"
	"github.com/alfmartinez/ego/engine/input"
	"github.com/alfmartinez/ego/engine/observer"
	"github.com/alfmartinez/ego/engine/renderer"
	"github.com/alfmartinez/ego/engine/state"
	"github.com/alfmartinez/ego/shared/input/prompt"
)

func GameFactory() game.Game {
	cfg := configuration.FromContext()
	inputName := cfg.GetString("input.type")
	inputHandler := input.CreateInputHandler(inputName).(prompt.TextHandler)
	context.Set("input", inputHandler)
	subject := observer.CreateSubject()
	context.Set("subject", subject)
	renderer := renderer.CreateRenderer(cfg.GetString("renderer.type"))
	context.Set("renderer", renderer)
	statesName := cfg.GetString("states")
	logic := CreateLogic()
	states := state.CreateStates(statesName, logic)
	logic.(state.StateMachine).SetStates(states)
	subject.Register(logic)
	return &textGame{
		subject: subject,
	}
}
