package path

import "testing"

type walkable struct {
	cost   int
	Around []Walkable
}

func CreateWalkable(cost int) *walkable {
	around := make([]Walkable, 0)
	return &walkable{cost, around}
}

func (w *walkable) MovementCost() int {
	return w.cost
}

func (w *walkable) Surrounding() []Walkable {
	return w.Around
}

func TestCreatePathMatrix(t *testing.T) {
	t.Run("Empty matrix", func(t *testing.T) {
		walkables := make([]Walkable, 0)
		matrix := CreatePathMatrix(walkables)
		if len(matrix) != 0 {
			t.Errorf("Expected matrix to be Empty")
		}
	})

	t.Run("1-1 matrix", func(t *testing.T) {
		walkables := make([]Walkable, 0)
		walkables = append(walkables, CreateWalkable(0))
		matrix := CreatePathMatrix(walkables)
		if len(matrix) != 1 {
			t.Errorf("Expected matrix to be 1 row, got %v", matrix)
		}
		if len(matrix[0]) != 1 {
			t.Errorf("Expected matrix to be 1 column, got %v", matrix)
		}
		if matrix[0][0] != 0 {
			t.Errorf("Expected matrix to be [0], got %v", matrix)
		}
	})

	t.Run("2-2 matrix, unconnected", func(t *testing.T) {
		walkables := make([]Walkable, 0)
		walkables = append(walkables, CreateWalkable(1))
		walkables = append(walkables, CreateWalkable(1))
		matrix := CreatePathMatrix(walkables)
		if len(matrix) != 2 {
			t.Errorf("Expected matrix to be 1 row, got %v", matrix)
		}
		if len(matrix[0]) != 2 {
			t.Errorf("Expected matrix to be 1 column, got %v", matrix)
		}
		for _, row := range matrix {
			for _, value := range row {
				if value != 0 {
					t.Errorf("Points are unconnected, should be 0")
				}
			}
		}
	})

	t.Run("2-2 matrix, unconnected", func(t *testing.T) {
		walkables := make([]Walkable, 0)
		walkables = append(walkables, CreateWalkable(1))
		walkables = append(walkables, CreateWalkable(1))
		connect(walkables[0], walkables[1])
		matrix := CreatePathMatrix(walkables)
		if len(matrix) != 2 {
			t.Errorf("Expected matrix to be 1 row, got %v", matrix)
		}
		if len(matrix[0]) != 2 {
			t.Errorf("Expected matrix to be 1 column, got %v", matrix)
		}
		for _, row := range matrix {
			for _, value := range row {
				if value != 0 {
					t.Errorf("Points are unconnected, should be 0, got %v", matrix)
				}
			}
		}
	})

}

func connect(aw Walkable, bw Walkable) {
	a := aw.(*walkable)
	b := bw.(*walkable)
	a.Around = append(a.Around, b)
	b.Around = append(b.Around, a)
}
