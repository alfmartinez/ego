package display

import (
	"ego/pkg/configuration"
	"ego/pkg/renderable"
	"ego/pkg/sprite/loader"
	"image"
	"log"
)

type Display interface {
	Init()
	Render() image.Image
	AddObject(renderable.Renderable)
	GetSize() configuration.Size
}

type CropableImage interface {
	image.Image
	SubImage(image.Rectangle) image.Image
}

func CreateDisplay(config configuration.Display) Display {
	name := config.Type
	log.Printf("Config Display %+v", config)
	displays := map[string]func(configuration.Display) Display{
		"rts": func(config configuration.Display) Display {
			loader := loader.CreateSpriteLoader("on_demand")
			return &rts{
				loader: loader,
				config: config,
			}
		},
	}

	return displays[name](config)
}
