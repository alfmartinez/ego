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
	"fmt"
	"time"
)

func Register() {
	game.RegisterGameFactory("text", func() game.Game {
		cfg := configuration.FromContext()
		inputName := cfg.GetString("input.type")
		inputHandler := input.CreateInputHandler(inputName).(prompt.TextHandler)
		context.Set("input", inputHandler)
		subject := observer.CreateSubject()
		context.Set("subject", subject)
		logic := CreateLogic(states())

		renderer := renderer.CreateRenderer(cfg.GetString("renderer.type"))
		context.Set("renderer", renderer)
		subject.Register(logic)
		return &textGame{
			subject: subject,
		}
	})
}

func states() state.States {
	subject := observer.FromContext()
	return state.States{
		"default": func(time.Duration) string {
			inputHandler := input.FromContext().(prompt.TextHandler)
			input := inputHandler.GetText()
			fmt.Printf("Got : %s\n", input)
			if input == "exit" {
				return "exit"
			}
			return ""
		},
		"exit": func(time.Duration) string {
			subject.NotifyAll(observer.CreateEvent(observer.EXIT))
			return ""
		},
	}
}
