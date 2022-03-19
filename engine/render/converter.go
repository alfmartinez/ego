package render

import (
	"ego/engine/display"
	"ego/engine/layer"
	"ego/engine/terrain"
	"ego/shared/object"
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
