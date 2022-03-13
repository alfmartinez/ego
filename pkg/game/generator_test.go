package game

import (
	"ego/internal/log"
	"ego/pkg/configuration"
	"ego/pkg/renderer"
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
	factory := func(scene Scene, r renderer.Renderer) Game {
		if _, ok := r.(*log.LogRenderer); !ok {
			t.Errorf("Should have LogRenderer, got %+v instead", r)
		}
		return &fakeGame{}
	}
	game := generateGame(config, factory)
	if _, ok := game.(*fakeGame); !ok {
		t.Error("Returned game should be fakeGame")
	}
}
