package game

import (
	"ego/pkg/configuration"
	"ego/pkg/mob"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"testing"
)

func TestGenerator(t *testing.T) {
	config := configuration.Configuration{
		Mobs: []configuration.Mob{
			{
				Type: "Mob",
				Name: "freddy",
			},
		},
		Renderer: configuration.Renderer{
			Type: "log",
		},
	}
	factory := func(objs []mob.GameObject, ter terrain.Terrain, r renderer.Renderer) Game {
		if _, ok := r.(*renderer.LogRenderer); !ok {
			t.Errorf("Should have LogRenderer, got %+v instead", r)
		}
		return &fakeGame{}
	}
	game := generateGame(config, factory)
	if _, ok := game.(*fakeGame); !ok {
		t.Error("Returned game should be fakeGame")
	}
}
