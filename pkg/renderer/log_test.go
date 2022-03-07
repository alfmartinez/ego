package renderer

import (
	"ego/pkg/configuration"
	"image"
	"testing"
)

type fakeRenderable struct{}

func (f *fakeRenderable) Name() string          { return "fake" }
func (f *fakeRenderable) Doing() string         { return "fake" }
func (f *fakeRenderable) Position() image.Point { return image.Pt(0, 0) }
func (f *fakeRenderable) Path() string          { return "mario.png" }

func (f *fakeRenderable) Size() uint { return 1 }

func TestLogRenderer(t *testing.T) {
	r := CreateRenderer(configuration.Renderer{
		Type: "log",
	})

	if r.IsAsync() {
		t.Error("Log renderer should not be async")
	}
	r.Init()                 // Do nothing
	r.Start(make(chan bool)) // Do nothing
	r.Refresh()              // Do nothing

	t.Error("to fix")

}
