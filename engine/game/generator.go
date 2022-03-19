package game

import (
	"ego/engine/configuration"
	"ego/engine/context"
	"ego/engine/input"
	"ego/engine/object"
	"ego/engine/observer"
	"ego/engine/physics"
	"ego/engine/renderer"
	"ego/engine/terrain"
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
