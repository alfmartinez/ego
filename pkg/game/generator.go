package game

import (
	"ego/pkg/configuration"
	"ego/pkg/context"
	"ego/pkg/input"
	"ego/pkg/object"
	"ego/pkg/observer"
	"ego/pkg/physics"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

func generateGame(factory func(observer.Subject, renderer.Renderer) Game) Game {

	cfg := configuration.FromContext()
	terrain.RegisterTileTypes()
	subject := observer.CreateSubject()

	for key, _ := range cfg.GetStringMap("mobs") {
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

	inputHandler := input.CreateInputHandler()
	context.Set("input", inputHandler)

	physics := physics.CreatePhysicsEngine()
	context.Set("physics", physics)

	game := factory(subject, r)

	return game
}
