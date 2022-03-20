package game

import (
	"ego/engine/context"
	"ego/engine/object"
	"ego/engine/observer"
	"ego/engine/renderer"
	"ego/engine/terrain"
	"testing"
	"time"

	"github.com/spf13/viper"
)

type fakeRenderer struct{}

func (r *fakeRenderer) Start()     {}
func (r *fakeRenderer) Close()     {}
func (r *fakeRenderer) Init()      {}
func (r *fakeRenderer) Refresh()   {}
func (r *fakeRenderer) Render(any) {}

type fakeGameObject struct{}

func (o *fakeGameObject) Update(time.Duration)    {}
func (o *fakeGameObject) OnNotify(observer.Event) {}

func TestGenerator(t *testing.T) {
	renderer.RegisterRendererFactory("fake", func() renderer.Renderer {
		return &fakeRenderer{}
	})
	object.RegisterObjectFactory("foo", func(key string) object.GameObject {
		return &fakeGameObject{}
	})

	context.CreateAndRegisterContext("test_generator")
	cfg := viper.New()
	context.Set("cfg", cfg)

	cfg.Set("mobs.bob.type", "foo")
	cfg.Set("renderer.type", "fake")

	gridData := terrain.GridData{
		Size:    1,
		Content: "XXX\nXXX\nXXX",
		Types: map[string]string{
			"x": "fake",
		},
	}
	cfg.Set("grid", gridData)
	tileTypes := map[string]struct {
		Sprite   string
		Movement int
		Collide  bool
	}{
		"fake": {
			Sprite:   "fake",
			Movement: 1,
			Collide:  true,
		},
	}
	cfg.Set("tile_types", tileTypes)
	terrain.RegisterTileTypes()
	generateGame(func(s observer.Subject, r renderer.Renderer) Game {
		return CreateSampleGame(s, r)
	})
}
