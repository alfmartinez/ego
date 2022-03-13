package fyne

import (
	"ego/pkg/configuration"
	"ego/pkg/render"
	"ego/pkg/renderer"
	"testing"
	"time"
)

func TestFyneRenderer(t *testing.T) {
	r := renderer.CreateRenderer(configuration.Renderer{
		Type: "fyne",
		Display: configuration.Display{
			Type: "rts",
		},
	})

	exit := make(chan bool)

	t.Run("should be Async", func(t *testing.T) {
		if !r.IsAsync() {
			t.Error("Fyne should be asynchronous")
		}
	})

	t.Run("should be Async", func(t *testing.T) {
		r.Init()
	})

	t.Run("Start should panic in test", func(t *testing.T) {
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
		tree := render.CreateRenderTree()
		r.Render(tree)
	})

	t.Run("Close should send message to exit", func(t *testing.T) {
		go func() {
			select {
			case v := <-exit:
				if !v {
					t.Error("exit should receive true")
				}
			case <-time.After(time.Millisecond):
				t.Error("exit chan did not receive")
			}
		}()
		r.Close()
	})

}
