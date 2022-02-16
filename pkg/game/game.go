package game

import (
	"ego/pkg/configuration"
	"ego/pkg/mob"
	"ego/pkg/terrain"
	"log"
)

type Game struct {
	Objects []GameObject
	Grid    terrain.Grid
}

type GameObject interface {
	Update(grid terrain.Grid)
	Render()
}

func generateGame(config configuration.Configuration) *Game {

	var mobs []GameObject
	for _, x := range config.Mobs {
		object := mob.New(x)
		mobs = append(mobs, object)
	}

	game := &Game{Objects: mobs}

	return game
}

func (game *Game) update() {
	for _, x := range game.Objects {
		x.Update(game.Grid)
	}
}

func (game *Game) render() {
	for _, x := range game.Objects {
		x.Render()
	}
}

func (game *Game) Loop() {
	log.Print("Starting game loop")
	doLoop := true
	turns := 120
	for doLoop {
		game.update()
		game.render()
		turns--
		doLoop = turns > 0
	}
}
