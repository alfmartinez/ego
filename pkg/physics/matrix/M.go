package matrix

type M interface {
	Advance(float64) M
}

type M2 struct {
	P Vec2
	S Vec2
	A Vec2
}

func (m M2) Advance(dt float64) M {
	return M2{
		P: m.P.Add(m.S.Mul(dt)),
		S: m.S.Add(m.A.Mul(dt)),
		A: m.A,
	}
}

type Matrixes []M2
type MatrixMappableFunc func(M) M

func (m Matrixes) Map(f MatrixMappableFunc) []M {
	results := make([]M, 0)
	for _, x := range m {
		results = append(results, f(x))
	}
	return results
}
