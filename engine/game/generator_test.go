package game

import (
	"ego/engine/object"
	"ego/engine/observer"
	"ego/engine/renderer"
	"testing"
	"time"
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
	generateGame(func(s observer.Subject, r renderer.Renderer) Game {
		return CreateSampleGame(s, r)
	})
}
