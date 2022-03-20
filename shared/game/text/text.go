package text

import (
	"ego/engine/configuration"
	"ego/engine/game"
	"ego/engine/input"
	"ego/shared/input/prompt"
)

func Register() {
	game.RegisterGameFactory("text", func() game.Game {
		inputName := configuration.FromContext().GetString("input.type")
		inputHandler := input.CreateInputHandler(inputName).(prompt.TextHandler)
		return &textGame{inputHandler}
	})
}
