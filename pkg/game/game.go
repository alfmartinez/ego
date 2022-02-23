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
	exitLoop chan bool
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
		exitLoop: make(chan bool),
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
		defer game.Renderer.Start(game.exitLoop)
	} else {
		game.Loop()
	}

}

const (
	UPDATE_RATE = 100
	RENDER_RATE = 10
)

func (game *Game) Loop() {
	log.Print("Start game loop")
	updateTicker := time.NewTicker(time.Second / UPDATE_RATE)
	renderTicker := time.NewTicker(time.Second / RENDER_RATE)
	loop := true
	for loop {
		select {
		case <-game.exitLoop:
			loop = false
			break
		case <-updateTicker.C:
			game.update()
			game.render()
		case <-renderTicker.C:
			//game.render()
		}
	}
	log.Print("End game loop")
}
