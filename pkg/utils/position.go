package utils

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
