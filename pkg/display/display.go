package display

import (
	"ego/pkg/renderable"
	"image"
)

type Display interface {
	Render() image.Image
	AddObject(renderable.Renderable)
}

func CreateDisplay(name string) Display {
	displays := map[string]func() Display{
		"rts": func() Display {
			return &rts{}
		},
	}

	return displays[name]()
}
