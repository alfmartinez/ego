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

func (p *Position) Add(a Position) Position {
	return Position{
		p.X + a.X,
		p.Y + a.Y,
	}
}

func (p *Position) IsZero() bool {
	return (p.X == 0) && (p.Y == 0)
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
