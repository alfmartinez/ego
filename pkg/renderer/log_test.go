package renderer

import (
	"ego/pkg/configuration"
	"testing"
)

func TestLogRenderer(t *testing.T) {
	r := CreateRenderer(configuration.Renderer{
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

	})

}
