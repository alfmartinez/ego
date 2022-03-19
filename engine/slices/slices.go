package slices

func Apply[T any](e []T, f func(m T)) {
	for _, o := range e {
		f(o)
	}
}
