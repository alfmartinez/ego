package game

import (
	"ego/pkg/context"
	"ego/pkg/input"
	"ego/pkg/observer"
	"ego/pkg/physics"
	"ego/pkg/renderer"
	"log"
	"time"
)

type sampleGame struct {
	subject  observer.Subject
	renderer renderer.Renderer
}

const (
	FPS = 30
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
	lastUpdate := time.Now()
	lastRender := time.Now()
	lastPhysics := time.Now()
	for loop {
		if inputHandler.IsPressed(input.ESCAPE) {
			loop = false
		}

		game.update(time.Since(lastUpdate))
		lastUpdate = time.Now()

		game.physics(time.Since(lastPhysics))
		lastPhysics = time.Now()

		renderWait := time.Until(lastRender.Add(time.Second / FPS))
		time.Sleep(renderWait)
		game.render()
		lastRender = time.Now()

	}
	game.renderer.Close()
	log.Print("End game loop")
}

func (game *sampleGame) update(dt time.Duration) {
	game.subject.NotifyAll(observer.CreateEvent(observer.UPDATE, dt))
}

func (game *sampleGame) physics(dt time.Duration) {
	physics.FromContext().Init()
	game.subject.NotifyAll(observer.CreateEvent(observer.PHYSICS))
	physics.FromContext().Advance(dt)
}

func (game *sampleGame) render() {
	evt := observer.CreateEvent(observer.RENDER)
	game.subject.NotifyAll(evt)
	game.renderer.Refresh()
}
