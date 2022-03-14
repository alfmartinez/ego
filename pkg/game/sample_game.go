package game

import (
	"ego/pkg/renderer"
	"log"
	"time"
)

type sampleGame struct {
	scene        Scene
	renderer     renderer.Renderer
	ExitLoop     chan bool
	renderTicker time.Ticker
	updateTicker time.Ticker
}

const (
	UPDATE_RATE = 30
	RENDER_RATE = 30
)

func CreateSampleGame(scene Scene, r renderer.Renderer) Game {

	updateTicker := time.NewTicker(time.Second / UPDATE_RATE)
	renderTicker := time.NewTicker(time.Second / RENDER_RATE)
	return &sampleGame{
		scene:        scene,
		renderer:     r,
		ExitLoop:     make(chan bool),
		renderTicker: *renderTicker,
		updateTicker: *updateTicker,
	}
}

func (game *sampleGame) Start() {
	game.renderer.Init()
	game.renderer.Start(game.ExitLoop)
	game.loop()
}

func (game *sampleGame) Stop() {
	game.ExitLoop <- true
}

func (game *sampleGame) loop() {
	log.Print("Start game loop")
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
	game.renderer.Close()
	log.Print("End game loop")
}

func (game *sampleGame) update() {
	game.scene.Update()
}

func (game *sampleGame) render() {
	renderTree := game.scene.Render()
	game.renderer.Render(renderTree)
	game.renderer.Refresh()
}
