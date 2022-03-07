package renderer

import (
	"ego/pkg/configuration"
	"testing"
)

func TestFyneRenderer(t *testing.T) {
	r := CreateRenderer(configuration.Renderer{
		Type: "fyne",
		Display: configuration.Display{
			Type: "rts",
		},
	})

	t.Run("should be Async", func(t *testing.T) {
		if !r.IsAsync() {
			t.Error("Fyne should be asynchronous")
		}
	})

	t.Run("should be Async", func(t *testing.T) {
		r.Init()
	})

	t.Run("Start should panic in test", func(t *testing.T) {
		exit := make(chan bool)
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Panic : %+v", r)
			}
		}()
		r.Start(exit)
	})

	t.Run("Refresh should get canvas content from display", func(t *testing.T) {
		r.Refresh()
	})

	t.Run("Render should render ^^'", func(t *testing.T) {
		tree := CreateRenderTree()
		r.Render(tree)
	})

}
