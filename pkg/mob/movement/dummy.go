package movement

import "image"

type dummy struct {
	position image.Point
}

func (d *dummy) Position() image.Point {
	return d.position
}

func CreateDummy(x, y int) Positionnable {
	return &dummy{image.Pt(x, y)}
}
