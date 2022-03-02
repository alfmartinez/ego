package display

import (
	"ego/pkg/configuration"
	"ego/pkg/renderable"
	"image"
)

type Display interface {
	Init()
	Render() image.Image
	AddObject(renderable.Renderable)
	//	GetSize() configuration.Size
}

type CropableImage interface {
	image.Image
	SubImage(image.Rectangle) image.Image
}

var displayFactories = make(map[string]func(configuration.Display) Display)

func RegisterDisplay(name string, f func(configuration.Display) Display) {
	displayFactories[name] = f
}

func CreateDisplay(config configuration.Display) Display {
	name := config.Type
	return displayFactories[name](config)
}
