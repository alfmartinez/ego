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
	t.Run("should create log renderer", func(t *testing.T) {
		o := CreateRenderer(configuration.Renderer{
			Type: "log",
		})
		if _, ok := o.(*LogRenderer); !ok {
			t.Errorf("Should return LogRenderer, got %+v", o)
		}
	})

	t.Run("should create fyne renderer", func(t *testing.T) {
		o := CreateRenderer(configuration.Renderer{
			Type: "fyne",
			Display: configuration.Display{
				Type: "rts",
			},
		})
		if _, ok := o.(*fyneRenderer); !ok {
			t.Errorf("Should return FyneRenderer, got %+v", o)
		}
	})

}
