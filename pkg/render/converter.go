package render

import (
	"ego/pkg/display"
	"ego/pkg/object"
	"image"
	"log"
)

func ConvertObjectToDisplayable(i interface{}) display.Displayable {
	var path string
	var size uint
	var position image.Point

	switch v := i.(type) {
	case object.StateMob:
		path = v.Path()
		size = v.Size()
		position = v.Position()
	default:
		log.Fatalf("Cannot convert %+v", i)
	}

	return display.CreateDisplayable(path, size, position)
}
