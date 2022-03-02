package display

import (
	"ego/pkg/configuration"
	"ego/pkg/renderable"
	"image"
	"testing"
)

type FakeDisplay struct {
}

func (f *FakeDisplay) Init() {}
func (f *FakeDisplay) Render() image.Image {
	return &image.RGBA{}
}
func (f *FakeDisplay) AddObject(s renderable.Renderable) {}
func (f *FakeDisplay) GetSize() configuration.Size {
	return configuration.Size{}
}

func TestRegisterDisplayAllowsCreation(t *testing.T) {
	RegisterDisplay("fake", func(configuration.Display) Display {
		return &FakeDisplay{}
	})

	config := configuration.Display{
		Type:     "fake",
		Size:     configuration.Size{Width: 10, Height: 10},
		ViewPort: configuration.Size{Width: 1, Height: 1},
	}

	display := CreateDisplay(config)
	s, ok := display.(*FakeDisplay)
	if !ok {
		t.Errorf("Should return FakeDisplay, got %+v", s)
	}
}
