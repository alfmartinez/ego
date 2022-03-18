package render

import (
	"ego/pkg/display"
	"ego/pkg/layer"
	"ego/pkg/object"
	"ego/pkg/terrain"
	"image"
	"log"
)

func ConvertObjectToDisplayable(i interface{}) display.Displayable {
	var path string
	var size uint
	var position image.Point
	var l layer.Layer

	switch v := i.(type) {
	case object.StateMob:
		path = v.Path()
		size = v.Size()
		position = v.Position().Point()
		l = layer.FOREGROUND
	case terrain.Tile:
		path = v.Path()
		size = v.Size()
		position = v.Position().Point()
		l = layer.BACKGROUND
	default:
		log.Printf("Cannot convert from %+v, return nil", i)
		return nil
	}

	return display.CreateDisplayable(path, size, position, l)
}
