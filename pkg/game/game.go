package game

import (
	"ego/pkg/configuration"
	"ego/pkg/mob"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"log"
)

type Game struct {
	Objects  []GameObject
	Terrain  terrain.Terrain
	Renderer renderer.Renderer
}

type GameObject interface {
	Update(terrain.Terrain)
	Render(renderer.Renderer)
}

func generateGame(config configuration.Configuration) *Game {

	var mobs []GameObject
	for _, x := range config.Mobs {
		object := mob.New(x)
		mobs = append(mobs, object)
	}

	renderer := renderer.CreateRenderer("log")

	game := &Game{
		Objects:  mobs,
		Renderer: renderer,
	}

	return game
}

func (game *Game) update() {
	for _, x := range game.Objects {
		x.Update(game.Terrain)
	}
}

func (game *Game) render() {
	for _, x := range game.Objects {
		x.Render(game.Renderer)
	}
}

func (game *Game) Loop() {
	log.Print("Starting game loop")
	doLoop := true
	for doLoop {
		game.update()
		game.render()
	}
}
