package game

import (
	"ego/pkg/configuration"
	"ego/pkg/mob"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"log"
	"time"
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

	terrain := terrain.CreateGrid(config.Grid.X, config.Grid.Y)

	renderer := renderer.CreateRenderer(config.Renderer)

	game := &Game{
		Objects:  mobs,
		Terrain:  terrain,
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
		time.Sleep(100 * time.Millisecond)
	}
}
