package display

import (
	"ego/engine/configuration"
	"ego/engine/context"
	"image"
	"testing"
)

type FakeDisplay struct {
}

func (f *FakeDisplay) Init() {}
func (f *FakeDisplay) Render() image.Image {
	return &image.RGBA{}
}
func (f *FakeDisplay) AddObject(s any) {}

func TestRegisterDisplayAllowsCreation(t *testing.T) {
	context.CreateAndRegisterContext("test")
	cfg := configuration.CreateConfiguration("").Get()
	context.Set("cfg", cfg)
	RegisterDisplay("fake", func() Display {
		return &FakeDisplay{}
	})
	t.Run("ok", func(t *testing.T) {
		cfg.Set("renderer.display.type", "fake")
		display := CreateDisplay()
		s, ok := display.(*FakeDisplay)
		if !ok {
			t.Errorf("Should return FakeDisplay, got %+v", s)
		}
	})
	t.Run("ko", func(t *testing.T) {
		cfg.Set("renderer.display.type", "foo")
		defer func() {
			if r := recover(); r == nil {
				t.Error("should panic")
			}
		}()
		CreateDisplay()

	})

}
