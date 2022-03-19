package display

import (
	"image"
	"testing"

	"github.com/spf13/viper"
)

type FakeDisplay struct {
}

func (f *FakeDisplay) Init() {}
func (f *FakeDisplay) Render() image.Image {
	return &image.RGBA{}
}
func (f *FakeDisplay) AddObject(s Displayable) {}

func TestRegisterDisplayAllowsCreation(t *testing.T) {
	RegisterDisplay("fake", func() Display {
		return &FakeDisplay{}
	})
	viper.Set("renderer.display.type", "fake")
	display := CreateDisplay()
	s, ok := display.(*FakeDisplay)
	if !ok {
		t.Errorf("Should return FakeDisplay, got %+v", s)
	}
}