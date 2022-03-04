package game

import (
	"ego/pkg/configuration"
	"ego/pkg/object"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

func init() {
	RegisterGameFactory("backend", func() Game {
		config, _ := configuration.ReadConfigurationFromFile("backend.yml")
		return generateGame(config, CreateBackendGame)
	})
}

func CreateBackendGame(objects []object.GameObject, grid terrain.Terrain, renderer renderer.Renderer) Game {
	return &backendGame{
		objects,
		grid,
		renderer,
		0,
		make(chan bool),
	}
}

type backendGame struct {
	Objects  []object.GameObject
	Grid     terrain.Terrain
	Renderer renderer.Renderer
	Speed    uint
	ExitLoop chan bool
}

func (game *backendGame) Start() {
	go game.loop()
	defer game.Renderer.Start(game.ExitLoop)
}

func (game *backendGame) Stop() {
	game.ExitLoop <- true
}

func (game *backendGame) loop() {

}
