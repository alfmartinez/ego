package text

import (
	"ego/engine/game"
	"ego/engine/input"
	"ego/shared/input/prompt"
)

func Register() {
	game.RegisterGameFactory("text", func() game.Game {
		inputHandler := input.CreateInputHandler("prompt").(prompt.TextHandler)
		return &textGame{inputHandler}
	})
}
