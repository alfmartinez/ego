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
	"fmt"
)

func generateGame(factory func(observer.Subject, renderer.Renderer) Game) Game {

	cfg := configuration.FromContext()
	terrain.RegisterTileTypes()

	inputHandler := input.CreateInputHandler()
	context.Set("input", inputHandler)

	subject := observer.CreateSubject()

	for key := range cfg.GetStringMap("mobs") {
		fmt.Printf("Creating object Key: %s\n", key)
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
