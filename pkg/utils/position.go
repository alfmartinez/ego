package utils

import "math"

type Position struct {
	X int
	Y int
}

func (p *Position) Relative(dx int, dy int) Position {
	return Position{
		p.X + dx,
		p.Y + dy,
	}
}

func (p *Position) DistanceTo(t Position) float64 {
	v := p.Diff(t)
	x, y := float64(v.X), float64(v.Y)
	return math.Sqrt(x*x + y*y)
}

func (p *Position) Add(a Position) Position {
	return Position{
		p.X + a.X,
		p.Y + a.Y,
	}
}

func (p *Position) IsZero() bool {
	return (p.X == 0) && (p.Y == 0)
}

func (p *Position) Diff(t Position) Position {
	return Position{
		X: t.X - p.X,
		Y: t.Y - p.Y,
	}
}

func (p *Position) UnitTowards(t *Position) Position {
	dx := t.X - p.X
	dy := t.Y - p.Y

	return Position{
		X: sign(dx),
		Y: sign(dy),
	}
}

func sign(value int) int {
	var x int
	switch {
	case value < 0:
		x = -1
	case value > 0:
		x = 1
	default:
		x = 0
	}
	return x
}

func (p *Position) Surrounding(distance int, validate func(Position) bool) []Position {
	var values []Position
	for x := -distance; x <= distance; x++ {
		for y := -distance; y <= distance; y++ {
			pos := Position{
				X: p.X + x,
				Y: p.Y + y,
			}
			if validate(pos) {
				values = append(values, pos)
			}
		}
	}
	return values
}
