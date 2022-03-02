package game

import (
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"log"
	"time"
)

func CreateSampleGame(mobs []GameObject, t terrain.Terrain, r renderer.Renderer) Game {
	return &sampleGame{
		Objects:  mobs,
		Terrain:  t,
		Renderer: r,
		ExitLoop: make(chan bool),
	}
}

type sampleGame struct {
	Objects  []GameObject
	Terrain  terrain.Terrain
	Renderer renderer.Renderer
	ExitLoop chan bool
}

const (
	UPDATE_RATE = 30
	RENDER_RATE = 30
)

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

func (game *sampleGame) Start() {
	if game.Renderer.IsAsync() {
		go game.loop()
		defer game.Renderer.Start(game.ExitLoop)
	} else {
		game.loop()
	}

}

func (game *sampleGame) loop() {
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
		case <-renderTicker.C:
			game.render()
		}
	}
	log.Print("End game loop")
}
