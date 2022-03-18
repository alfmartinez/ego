package physics

import "time"

type M interface {
	Advance(time.Duration)
}

type M2 struct {
	P Vec2
	S Vec2
	A Vec2
}

func (m M2) Advance(dt float64) M2 {
	return M2{
		P: m.P.Add(m.S.Mul(dt)),
		S: m.S.Add(m.A.Mul(dt)),
		A: m.A,
	}
}
