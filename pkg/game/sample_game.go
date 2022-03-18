package game

import (
	"ego/pkg/context"
	"ego/pkg/input"
	"ego/pkg/observer"
	"ego/pkg/renderer"
	"log"
)

type sampleGame struct {
	subject  observer.Subject
	renderer renderer.Renderer
}

const (
	UPDATE_RATE = 1
	RENDER_RATE = 1
)

func CreateSampleGame(subject observer.Subject, r renderer.Renderer) Game {
	return &sampleGame{
		subject:  subject,
		renderer: r,
	}
}

func (game *sampleGame) Start() {
	game.renderer.Init()
	game.renderer.Start()
	game.loop()
}

func (game *sampleGame) Stop() {}

func (game *sampleGame) loop() {
	inputHandler := context.GetContext().Get("input").(input.InputHandler)
	log.Print("Start game loop")
	loop := true
	for loop {
		if inputHandler.IsPressed(input.ESCAPE) {
			loop = false
		}
		game.update()
		game.render()

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
