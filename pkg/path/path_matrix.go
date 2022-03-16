package path

import "math"

type PathMatrix [][]float64

type Walkable interface {
	MovementCost() int
	Surrounding() []Walkable
}

func CreatePathMatrix(walkables []Walkable) PathMatrix {
	n := len(walkables)
	indexes := make(map[Walkable]int)
	for idx, value := range walkables {
		indexes[value] = idx
	}
	var matrix = buildMatrix(n)
	for idx, value := range walkables {
		matrix[idx][idx] = 0
		for _, connected := range value.Surrounding() {
			cost := ExpOrZero(value.MovementCost()) * ExpOrZero(connected.MovementCost())
			matrix[indexes[connected]][idx] = math.Exp(cost)
		}
	}
	return matrix
}

func buildMatrix(n int) PathMatrix {
	matrix := make([][]float64, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, n)
	}
	return matrix
}

func ExpOrZero(value int) float64 {
	if value == 0 {
		return float64(value)
	}
	return math.Exp(float64(value))
}
