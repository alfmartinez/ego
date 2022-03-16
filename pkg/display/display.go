package display

import (
	"image"

	"github.com/spf13/viper"
)

type Display interface {
	Init()
	Render() image.Image
	AddObject(Displayable)
}

type CropableImage interface {
	image.Image
	SubImage(image.Rectangle) image.Image
}

var displayFactories = make(map[string]func() Display)

func RegisterDisplay(name string, f func() Display) {
	displayFactories[name] = f
}

func CreateDisplay() Display {
	name := viper.GetString("renderer.display.type")
	return displayFactories[name]()
}
