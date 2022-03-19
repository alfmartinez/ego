package display

import (
	"ego/engine/layer"
	"image"
)

type Displayable interface {
	Path() string
	Size() uint
	Position() image.Point
	Layer() layer.Layer
}

func CreateDisplayable(path string, size uint, position image.Point, layer layer.Layer) Displayable {
	return displayable{path, size, position, layer}
}

type displayable struct {
	path     string
	size     uint
	position image.Point
	layer    layer.Layer
}

func (d displayable) Path() string {
	return d.path
}

func (d displayable) Size() uint {
	return d.size
}

func (d displayable) Position() image.Point {
	return d.position
}

func (d displayable) Layer() layer.Layer {
	return d.layer
}
