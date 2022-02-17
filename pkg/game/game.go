package game

import (
	"ego/pkg/configuration"
	"ego/pkg/mob"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
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
	renderer.Init()

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
	game.Renderer.Refresh()
}

func (game *Game) Start() {
	if game.Renderer.IsAsync() {
		go game.Loop()
		defer game.Renderer.Start()
	} else {
		game.Loop()
	}

}

const (
	UPDATE_RATE = 10000
	RENDER_RATE = 10000
)

func (game *Game) Loop() {
	updateTicker := time.NewTicker(time.Second / UPDATE_RATE)
	renderTicker := time.NewTicker(time.Second / RENDER_RATE)
	for {
		select {
		case <-updateTicker.C:
			game.update()
		case <-renderTicker.C:
			game.render()
		}
	}
}
