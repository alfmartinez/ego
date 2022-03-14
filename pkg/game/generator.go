package game

import (
	"ego/pkg/configuration"
	"ego/pkg/object"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

func generateGame(config configuration.Configuration, factory func(Scene, renderer.Renderer) Game) Game {

	scene := CreateScene()
	root := scene.Root()
	tiles := root.CreateFolder("tile")
	mobs := root.CreateFolder("mobs")

	for _, x := range config.Mobs {
		o := object.CreateObject(x)
		mobs.AddObject(o)
	}

	terrain.CreateGrid(config.Grid.X, config.Grid.Y, func(t terrain.Tile) {
		tiles.AddObject(t)
	})
	r := renderer.CreateRenderer(config.Renderer)

	//r.Init()

	game := factory(scene, r)

	return game
}
