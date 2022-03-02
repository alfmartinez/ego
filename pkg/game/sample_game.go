package game

import (
	"ego/pkg/mob"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"log"
	"time"
)

type sampleGame struct {
	Objects      []mob.GameObject
	Terrain      terrain.Terrain
	Renderer     renderer.Renderer
	ExitLoop     chan bool
	renderTicker time.Ticker
	updateTicker time.Ticker
}

const (
	UPDATE_RATE = 30
	RENDER_RATE = 30
)

func CreateSampleGame(mobs []mob.GameObject, t terrain.Terrain, r renderer.Renderer) Game {

	updateTicker := time.NewTicker(time.Second / UPDATE_RATE)
	renderTicker := time.NewTicker(time.Second / RENDER_RATE)
	return &sampleGame{
		Objects:      mobs,
		Terrain:      t,
		Renderer:     r,
		ExitLoop:     make(chan bool),
		renderTicker: *renderTicker,
		updateTicker: *updateTicker,
	}
}

func (game *sampleGame) Start() {
	go game.loop()
	defer game.Renderer.Start(game.ExitLoop)
}

func (game *sampleGame) Stop() {
	game.ExitLoop <- true
}

func (game *sampleGame) loop() {
	loop := true
	for loop {
		select {
		case <-game.ExitLoop:
			loop = false
		case <-game.updateTicker.C:
			game.update()
		case <-game.renderTicker.C:
			game.render()
		}
	}
	log.Print("End game loop")
}

func (game *sampleGame) update() {
	for _, x := range game.Objects {
		x.Update(game.Terrain)
	}
}

func (game *sampleGame) render() {
	game.Terrain.Render(game.Renderer)
	for _, x := range game.Objects {
		x.Render(game.Renderer)
	}
	game.Renderer.Refresh()
}
