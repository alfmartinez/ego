package game

import (
	"ego/pkg/object"
	"ego/pkg/renderer"
	"ego/pkg/terrain"

	"github.com/spf13/viper"
)

func generateGame(factory func(Scene, renderer.Renderer) Game) Game {

	scene := CreateScene()
	root := scene.Root()
	tiles := root.CreateFolder("tile")
	mobs := root.CreateFolder("mobs")

	for key, _ := range viper.GetStringMap("mobs") {
		o := object.CreateObject(key)
		mobs.AddObject(o)
	}

	terrain.CreateGrid(func(t terrain.Tile) {
		tiles.AddObject(t)
	})
	r := renderer.CreateRenderer()

	//r.Init()

	game := factory(scene, r)

	return game
}
