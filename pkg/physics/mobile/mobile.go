package mobile

import "ego/pkg/physics/matrix"

type Mobile interface {
	GetMatrix() matrix.M
}

type Mobiles []Mobile

func (a Mobiles) Map(f func(m Mobile) interface{}) []interface{} {
	results := make([]interface{}, 0)
	for _, m := range a {
		results = append(results, f(m))
	}
	return results
}

func (a Mobiles) Apply(f func(m Mobile)) {
	for _, m := range a {
		f(m)
	}
}

func (a Mobiles) Filter(f func(m Mobile) bool) Mobiles {
	results := make([]Mobile, 0)
	for _, m := range a {
		if f(m) {
			results = append(results, m)
		}
	}
	return results
}
