package terrain

import (
	"ego/pkg/motivator"
	"testing"
)

func TestGetResourcesProviding(t *testing.T) {
	cases := []struct {
		name     string
		need     motivator.Need
		expected []Resource
	}{
		{
			name:     "Health",
			need:     motivator.Health,
			expected: []Resource{Medicine},
		}, {
			name:     "Food",
			need:     motivator.Food,
			expected: []Resource{Food, Water},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := GetResourcesProviding(tt.need)
			if len(actual) != len(tt.expected) {
				t.Error("dont match len")
			}
			for i, x := range tt.expected {
				if x != actual[i] {
					t.Error("No match")
				}
			}
		})
	}

	t.Run("Unknown need", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("Should panic")
			}
		}()
		GetResourcesProviding(255)
	})
}
