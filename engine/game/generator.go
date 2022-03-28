package game

import (
	"github.com/alfmartinez/ego/engine/configuration"
	"github.com/alfmartinez/ego/engine/context"
	"github.com/alfmartinez/ego/engine/input"
	"github.com/alfmartinez/ego/engine/object"
	"github.com/alfmartinez/ego/engine/observer"
	"github.com/alfmartinez/ego/engine/physics"
	"github.com/alfmartinez/ego/engine/renderer"
	"github.com/alfmartinez/ego/engine/terrain"
)

func generateGame(factory func(observer.Subject, renderer.Renderer) Game) Game {

	cfg := configuration.FromContext()
	terrain.RegisterTileTypes()

	inputHandler := input.CreateInputHandler("key").(input.KeyHandler)
	context.Set("input", inputHandler)

	subject := observer.CreateSubject()

	for key := range cfg.GetStringMap("mobs") {
		o := object.CreateObject(key)
		subject.Register(o)
	}

	grid := terrain.CreateGrid(func(t terrain.Tile) {
		subject.Register(t)
	})

	context.Set("terrain", grid)
	name := cfg.GetString("renderer.type")
	r := renderer.CreateRenderer(name)
	context.Set("renderer", r)

	physics := physics.CreatePhysicsEngine()
	context.Set("physics", physics)

	game := factory(subject, r)

	return game
}
