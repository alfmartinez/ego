package game

import (
	"ego/pkg/configuration"
	"ego/pkg/mob"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

type Game interface {
	Start()
}

type GameObject interface {
	Update(terrain.Terrain)
	Render(renderer.Renderer)
}

var factories = make(map[string]func() (Game, error))

func RegisterGameFactory(name string, f func() (Game, error)) {
	factories[name] = f
}

func CreateGame(name string) (Game, error) {
	return factories[name]()
}

func generateGame(config configuration.Configuration) Game {

	var mobs []GameObject
	for _, x := range config.Mobs {
		object := mob.New(x)
		mobs = append(mobs, object)
	}

	t := terrain.CreateGrid(config.Grid.X, config.Grid.Y)

	r := renderer.CreateRenderer(config.Renderer)
	r.Init()

	game := CreateSampleGame(mobs, t, r)

	return game
}
