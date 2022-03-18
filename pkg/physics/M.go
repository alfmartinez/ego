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

const NUM_STEPS = 100

func (m *M2) Advance(dt time.Duration) {
	timeStep := dt.Seconds() / NUM_STEPS
	for i := 0; i < NUM_STEPS; i++ {
		m.S = m.S.Add(m.A.Mul(timeStep))
		m.P = m.P.Add(m.S.Mul(timeStep))
	}

}
