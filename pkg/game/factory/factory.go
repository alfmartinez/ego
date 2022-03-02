package factory

import (
	"ego/pkg/configuration"
	"ego/pkg/game"
	"ego/pkg/mob"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

func CreateGame(filename string) (game.GameInterface, error) {
	return CreateGameFromConfiguration(filename)
}

func generateGame(config configuration.Configuration) game.GameInterface {

	var mobs []game.GameObject
	for _, x := range config.Mobs {
		object := mob.New(x)
		mobs = append(mobs, object)
	}

	terrain := terrain.CreateGrid(config.Grid.X, config.Grid.Y)

	renderer := renderer.CreateRenderer(config.Renderer)
	renderer.Init()

	game := &game.SampleGame{
		Objects:  mobs,
		Terrain:  terrain,
		Renderer: renderer,
		ExitLoop: make(chan bool),
	}

	return game
}
