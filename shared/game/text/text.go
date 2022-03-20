package text

import (
	"ego/engine/configuration"
	"ego/engine/context"
	"ego/engine/game"
	"ego/engine/input"
	"ego/engine/observer"
	"ego/engine/renderer"
	"ego/engine/state"
	"ego/shared/input/prompt"
)

func Register() {
	game.RegisterGameFactory("text", func() game.Game {
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
	})
}
