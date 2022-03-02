package game

import (
	"ego/pkg/configuration"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

func init() {
	RegisterGameFactory("file", func() (Game, error) {
		config, _ := configuration.ReadConfigurationFromFile("game.yml")
		return generateGame(config, CreateSampleGame), nil
	})
}

type Game interface {
	Start()
	Stop()
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
