package render

import (
	"ego/engine/display"
	"ego/engine/layer"
	"ego/engine/movement"
	"ego/engine/terrain"
	"image"
	"log"
)

type Convertible interface {
	Path() string
	Size() uint
	Position() movement.Location
	Layer() layer.Layer
}

func ConvertObjectToDisplayable(c Convertible) display.Displayable {
	var path = c.Path()
	var size = c.Size()
	var position = image.Pt(int(c.Position().X), int(c.Position().Y))
	var l = c.Layer()

	return display.CreateDisplayable(path, size, position, l)
}
