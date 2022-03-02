package game

import (
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"log"
	"time"
)

type SampleGame struct {
	Objects  []GameObject
	Terrain  terrain.Terrain
	Renderer renderer.Renderer
	ExitLoop chan bool
}

const (
	UPDATE_RATE = 30
	RENDER_RATE = 30
)

func (game *SampleGame) update() {
	for _, x := range game.Objects {
		x.Update(game.Terrain)
	}
}

func (game *SampleGame) render() {
	game.Terrain.Render(game.Renderer)
	for _, x := range game.Objects {
		x.Render(game.Renderer)
	}
	game.Renderer.Refresh()
}

func (game *SampleGame) Start() {
	if game.Renderer.IsAsync() {
		go game.loop()
		defer game.Renderer.Start(game.ExitLoop)
	} else {
		game.loop()
	}

}

func (game *SampleGame) loop() {
	log.Print("Start game loop")
	updateTicker := time.NewTicker(time.Second / UPDATE_RATE)
	renderTicker := time.NewTicker(time.Second / RENDER_RATE)
	loop := true
	for loop {
		select {
		case <-game.ExitLoop:
			loop = false
		case <-updateTicker.C:
			game.update()
			game.render()
		case <-renderTicker.C:
			//game.render()
		}
	}
	log.Print("End game loop")
}
