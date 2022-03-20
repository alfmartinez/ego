package text

import "ego/engine/game"

func Register() {
	game.RegisterGameFactory("text", func() game.Game {
		return &textGame{}
	})
}
