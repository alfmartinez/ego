package display

import "image"

type Displayable interface {
	Path() string
	Size() uint
	Position() image.Point
}

func CreateDisplayable(path string, size uint, position image.Point) Displayable {
	return displayable{path, size, position}
}

type displayable struct {
	path     string
	size     uint
	position image.Point
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
