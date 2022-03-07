package game

import (
	"ego/pkg/configuration"
	"ego/pkg/renderer"
	"testing"
	"time"
)

type fakeMob struct {
	Updated bool
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
	folder := scene.Root().CreateFolder("mobs")
	folder.AddObject(mob)
	r := renderer.CreateRenderer(configuration.Renderer{
		Type: "log",
	})
	game := CreateSampleGame(scene, r)

	if mob.Updated {
		t.Error("Mob should not have been updated")
	}

	game.Start()
	time.Sleep(time.Second)
	game.Stop()

	if !mob.Updated {
		t.Error("Mob should have been updated")
	}
}
