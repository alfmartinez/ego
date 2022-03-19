package matrix

import (
	"testing"
	"time"
)

func TestEvaluatePositionV2(t *testing.T) {
	cases := []struct {
		label    string
		given    M2
		expected M2
		dt       time.Duration
	}{
		{
			"Immobile",
			M2{
				P: Vec2{0, 0},
				S: Vec2{0, 0},
				A: Vec2{0, 0},
			},
			M2{
				P: Vec2{0, 0},
				S: Vec2{0, 0},
				A: Vec2{0, 0},
			},
			time.Second,
		},
		{
			"Push forward",
			M2{
				P: Vec2{0, 0},
				S: Vec2{1, 0},
				A: Vec2{0, 0},
			},
			M2{
				P: Vec2{1, 0},
				S: Vec2{1, 0},
				A: Vec2{0, 0},
			},
			time.Second,
		},
		{
			"Accel forward",
			M2{
				P: Vec2{0, 0},
				S: Vec2{0, 0},
				A: Vec2{1, 0},
			},
			M2{
				P: Vec2{0, 0},
				S: Vec2{1, 0},
				A: Vec2{0, 0},
			},
			time.Second,
		},
		{
			"Accel forward",
			M2{
				P: Vec2{0, 0},
				S: Vec2{0, 0},
				A: Vec2{0, -10},
			},
			M2{
				P: Vec2{0, 5},
				S: Vec2{1, -10},
				A: Vec2{0, -10},
			},
			time.Second,
		},
	}

	for _, tt := range cases {
		t.Run(tt.label, func(t *testing.T) {
			tt.given.Advance(tt.dt.Seconds())
			if tt.given != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, tt.given)
			}
		})
	}
}
