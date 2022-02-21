package display

import "image"

type Display interface {
	Render() image.Image
}

func CreateDisplay(name string) Display {
	displays := map[string]func() Display{
		"rts": func() Display {
			return &rts{}
		},
	}

	return displays[name]()
}
