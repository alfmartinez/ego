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

	terrain := terrain.CreateGrid(config.Grid.X, config.Grid.Y)

	renderer := renderer.CreateRenderer(config.Renderer)
	renderer.Init()

	game := &SampleGame{
		Objects:  mobs,
		Terrain:  terrain,
		Renderer: renderer,
		ExitLoop: make(chan bool),
	}

	return game
}
