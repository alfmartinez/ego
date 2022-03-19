package matrix

type Vec2 struct {
	X float64
	Y float64
}

func (v Vec2) Add(d Vec2) Vec2 {
	return Vec2{
		X: v.X + d.X,
		Y: v.Y + d.Y,
	}
}

func (v Vec2) Mul(f float64) Vec2 {
	return Vec2{
		X: v.X * f,
		Y: v.Y * f,
	}
}
