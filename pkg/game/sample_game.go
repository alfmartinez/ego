package game

import (
	"ego/pkg/observer"
	"ego/pkg/renderer"
	"log"
	"time"
)

type sampleGame struct {
	subject      observer.Subject
	renderer     renderer.Renderer
	ExitLoop     chan bool
	renderTicker time.Ticker
	updateTicker time.Ticker
}

const (
	UPDATE_RATE = 30
	RENDER_RATE = 30
)

func CreateSampleGame(subject observer.Subject, r renderer.Renderer) Game {

	updateTicker := time.NewTicker(time.Second / UPDATE_RATE)
	renderTicker := time.NewTicker(time.Second / RENDER_RATE)
	return &sampleGame{
		subject:      subject,
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
	game.subject.NotifyAll(observer.CreateEvent(observer.UPDATE))
}

func (game *sampleGame) render() {
	evt := observer.CreateEvent(observer.RENDER)
	game.subject.NotifyAll(evt)
	game.renderer.Refresh()
}
