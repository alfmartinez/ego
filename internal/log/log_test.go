package log

import (
	"ego/pkg/configuration"
	"ego/pkg/render"
	"ego/pkg/renderer"
	"testing"
)

func TestLogRenderer(t *testing.T) {
	r := renderer.CreateRenderer(configuration.Renderer{
		Type: "log",
	})
	t.Run("should be Async", func(t *testing.T) {
		if r.IsAsync() {
			t.Error("Log renderer should not be async")
		}
	})

	r.Init()                 // Do nothing
	r.Start(make(chan bool)) // Do nothing
	r.Refresh()              // Do nothing

	t.Run("render", func(t *testing.T) {
		r.Render(render.CreateRenderTree())
	})

}
