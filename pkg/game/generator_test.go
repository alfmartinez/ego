package game

import (
	_ "ego/internal/renderer/glfw"
	"ego/pkg/object"
	"ego/pkg/render"
	"ego/pkg/renderer"
	"testing"
)

type fakeRenderer struct{}

func (r *fakeRenderer) Start(chan bool)          {}
func (r *fakeRenderer) Close()                   {}
func (r *fakeRenderer) Init()                    {}
func (r *fakeRenderer) IsAsync() bool            { return false }
func (r *fakeRenderer) Refresh()                 {}
func (r *fakeRenderer) Render(render.RenderTree) {}

type fakeGameObject struct{}

func (o *fakeGameObject) Update() {}

func TestGenerator(t *testing.T) {
	renderer.RegisterRendererFactory("fake", func() renderer.Renderer {
		return &fakeRenderer{}
	})
	object.RegisterObjectFactory("foo", func(key string) object.GameObject {
		return &fakeGameObject{}
	})
	generateGame(func(s Scene, r renderer.Renderer) Game {
		return CreateSampleGame(s, r)
	})
}
