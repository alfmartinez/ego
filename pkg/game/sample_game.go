package game

import (
	"ego/pkg/renderer"
	"log"
	"time"
)

type sampleGame struct {
	Scene        Scene
	Renderer     renderer.Renderer
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
		Scene:        scene,
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
	game.Scene.Update()
}

func (game *sampleGame) render() {
	renderTree := game.Scene.Render()
	game.Renderer.Render(renderTree)
	game.Renderer.Refresh()
}
