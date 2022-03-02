package game

import (
	"ego/pkg/configuration"
	"ego/pkg/object"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"testing"
	"time"
)

type fakeMob struct {
	Rendered bool
	Updated  bool
}

func (f *fakeMob) Render(r renderer.Renderer) {
	f.Rendered = true
}
func (f *fakeMob) Update(t terrain.Terrain) {
	f.Updated = true
}

func TestCreateSampleGame(t *testing.T) {
	mobs := make([]object.GameObject, 0)
	ter := terrain.CreateGrid(0, 0)
	r := renderer.CreateRenderer(configuration.Renderer{
		Type: "log",
	})
	game := CreateSampleGame(mobs, ter, r)
	if _, ok := game.(*sampleGame); !ok {
		t.Errorf("Should create sample game, got %v", game)
	}
}

func TestSampleGameSyncRenderer(t *testing.T) {
	mobs := make([]object.GameObject, 0)
	mob := &fakeMob{}
	mobs = append(mobs, mob)
	ter := terrain.CreateGrid(1, 1)
	r := renderer.CreateRenderer(configuration.Renderer{
		Type: "log",
	})
	game := CreateSampleGame(mobs, ter, r)

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
