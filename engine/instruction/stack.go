package instruction

type Stack[T any] interface {
	Pop() T
	Push(T)
	Values() []T
}

func CreateStack[T any]() Stack[T] {
	values := make([]T, 0)
	return &stack[T]{values}
}

type stack[T any] struct {
	values []T
}

func (s *stack[T]) Values() []T {
	return s.values
}

func (s *stack[T]) Pop() T {
	value := s.values[0]
	s.values = s.values[1:]
	return value
}

func (s *stack[T]) Push(value T) {
	s.values = append([]T{value}, s.values...)
}
