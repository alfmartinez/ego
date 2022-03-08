package render

import (
	"ego/pkg/configuration"
	"ego/pkg/display"
	"ego/pkg/object"
	"ego/pkg/terrain"
	"image"
	"testing"
)

func TestConverter(t *testing.T) {
	var cases = []struct {
		name     string
		in       interface{}
		expected display.Displayable
	}{
		{
			name:     "int",
			in:       1,
			expected: nil,
		},
		{
			name: "StateMob",
			in: object.CreateStateMob(configuration.Mob{
				Position: configuration.Position{
					X: 99,
					Y: 1,
				},
				Sprite: "foo.png",
			}),
			expected: display.CreateDisplayable("foo.png", 50, image.Pt(99, 1)),
		},
		{
			name:     "Tile",
			in:       terrain.CreateTile(image.Pt(30, 30), terrain.CreateTileType("plain"), 50),
			expected: display.CreateDisplayable("mario.png", 50, image.Pt(1500, 1500)),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := ConvertObjectToDisplayable(tt.in)
			if actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
