package game

import (
	"ego/pkg/configuration"
	"ego/pkg/mob"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

func generateGame(config configuration.Configuration, factory func([]GameObject, terrain.Terrain, renderer.Renderer) Game) Game {

	var mobs []GameObject
	for _, x := range config.Mobs {
		object := mob.New(x)
		mobs = append(mobs, object)
	}

	t := terrain.CreateGrid(config.Grid.X, config.Grid.Y)

	r := renderer.CreateRenderer(config.Renderer)
	r.Init()

	game := factory(mobs, t, r)

	return game
}
