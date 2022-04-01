package slices

func Apply[T any](e []T, f func(m T)) {
	for _, o := range e {
		f(o)
	}
}

func Filter[T any](e []T, f func(m T) bool) []T {
	var result []T = make([]T, 0)
	for _, o := range e {
		if f(o) {
			result = append(result, o)
		}
	}
	return result
}
