package movement

import "image"

type Location struct {
	X, Y float64
}

func (l Location) Point() image.Point {
	return image.Pt(int(l.X), int(l.Y))
}

func (l Location) Add(p image.Point) Location {
	return Location{
		X: l.X + float64(p.X),
		Y: l.Y + float64(p.Y),
	}
}

func (l Location) In(r image.Rectangle) bool {
	return l.Point().In(r)
}

func (l Location) Div(d int) Location {
	return Location{
		X: l.X / float64(d),
		Y: l.Y / float64(d),
	}
}

func (l Location) Sub(s Location) Location {
	return Location{
		X: l.X - s.X,
		Y: l.Y - s.Y,
	}
}

type Speed Location

type Acceleration Location
