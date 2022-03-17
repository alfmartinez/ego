package game

import (
	"ego/pkg/context"
	"ego/pkg/object"
	"ego/pkg/observer"
	"ego/pkg/renderer"
	"ego/pkg/terrain"

	"github.com/spf13/viper"
)

func generateGame(factory func(observer.Subject, renderer.Renderer) Game) Game {
	ctx := context.GetContext()
	cfg := ctx.Get("cfg").(*viper.Viper)
	terrain.RegisterTileTypes()
	subject := observer.CreateSubject()

	for key, _ := range cfg.GetStringMap("mobs") {
		o := object.CreateObject(key)
		subject.Register(o)
	}

	grid := terrain.CreateGrid(func(t terrain.Tile) {
		subject.Register(t)
	})

	ctx.Set("terrain", grid)
	name := cfg.GetString("renderer.type")
	r := renderer.CreateRenderer(name)
	ctx.Set("renderer", r)

	game := factory(subject, r)

	return game
}
