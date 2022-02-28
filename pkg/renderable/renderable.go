package renderable

import "image"

type Renderable interface {
	Position() image.Point
	Name() string
	Doing() string
	Path() string
	Size() uint
	Multiplicator() int
}
