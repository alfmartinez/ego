package renderer

import (
	"ego/pkg/configuration"
	"ego/pkg/render"
	"testing"
)

type fakeRenderer struct{}

func (f *fakeRenderer) Init()                    {}
func (f *fakeRenderer) IsAsync() bool            { return false }
func (f *fakeRenderer) Start(e chan bool)        {}
func (f *fakeRenderer) Render(render.RenderTree) {}
func (f *fakeRenderer) Refresh()                 {}
func (f *fakeRenderer) Close()                   {}

func TestCreateRenderer(t *testing.T) {
	RegisterRendererFactory("fake", func(r configuration.Renderer) Renderer {
		return &fakeRenderer{}
	})
	t.Run("should create registered renderer", func(t *testing.T) {
		o := CreateRenderer(configuration.Renderer{
			Type: "fake",
		})
		if _, ok := o.(*fakeRenderer); !ok {
			t.Errorf("Should return fakeRenderer, got %+v", o)
		}
	})

	t.Run("should panic if not registered", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Should have panicked")
			}
		}()
		CreateRenderer(configuration.Renderer{
			Type: "panicplz",
		})

	})
}
