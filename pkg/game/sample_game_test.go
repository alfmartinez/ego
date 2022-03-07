package game

import (
	"ego/pkg/configuration"
	"ego/pkg/renderer"
	"testing"
	"time"
)

type fakeMob struct {
	Rendered bool
	Updated  bool
}

func (f *fakeMob) Update() {
	f.Updated = true
}

func TestCreateSampleGame(t *testing.T) {
	scene := CreateScene()
	r := renderer.CreateRenderer(configuration.Renderer{
		Type: "log",
	})
	game := CreateSampleGame(scene, r)
	if _, ok := game.(*sampleGame); !ok {
		t.Errorf("Should create sample game, got %v", game)
	}
}

func TestSampleGameSyncRenderer(t *testing.T) {
	mob := &fakeMob{}
	scene := CreateScene()
	scene.Root().AddObject(mob)
	r := renderer.CreateRenderer(configuration.Renderer{
		Type: "log",
	})
	game := CreateSampleGame(scene, r)

	if mob.Rendered {
		t.Error("Mob should not have been rendered")
	}
	if mob.Updated {
		t.Error("Mob should not have been updated")
	}

	game.Start()
	time.Sleep(time.Second)
	game.Stop()

	if !mob.Rendered {
		t.Error("Mob should have been rendered")
	}
	if !mob.Updated {
		t.Error("Mob should have been updated")
	}
}
