package movement

import "image"

type location struct {
	position image.Point
}

func Loc(pt image.Point) Positionnable {
	return &location{pt}
}

func (m *location) Position() image.Point {
	return m.position
}

func (m *location) IsAt(pos Positionnable) bool {
	return m.position.Eq(pos.Position())
}
