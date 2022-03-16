package renderer

import (
	"ego/pkg/render"
	"testing"

	"github.com/spf13/viper"
)

type fakeRenderer struct{}

func (f *fakeRenderer) Init()                    {}
func (f *fakeRenderer) IsAsync() bool            { return false }
func (f *fakeRenderer) Start(e chan bool)        {}
func (f *fakeRenderer) Render(render.RenderTree) {}
func (f *fakeRenderer) Refresh()                 {}
func (f *fakeRenderer) Close()                   {}

func TestCreateRenderer(t *testing.T) {
	RegisterRendererFactory("fake", func() Renderer {
		return &fakeRenderer{}
	})
	t.Run("should create registered renderer", func(t *testing.T) {
		viper.Set("renderer.type", "fake")
		o := CreateRenderer()
		if _, ok := o.(*fakeRenderer); !ok {
			t.Errorf("Should return fakeRenderer, got %+v", o)
		}
	})

	t.Run("should panic if not registered", func(t *testing.T) {
		viper.Set("renderer.type", "foo")
		defer func() {
			if r := recover(); r == nil {
				t.Error("Should have panicked")
			}
		}()
		CreateRenderer()

	})
}
