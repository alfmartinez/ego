package render

import (
	"ego/engine/display"
	"ego/engine/layer"
	"ego/engine/movement"
	"ego/engine/terrain"
	"ego/shared/object"
	"image"
	"testing"

	"github.com/spf13/viper"
)

func TestConverter(t *testing.T) {
	var cases = []struct {
		name     string
		in       func() interface{}
		expected display.Displayable
	}{
		{
			name: "int",
			in: func() interface{} {
				return 1
			},
			expected: nil,
		},
		{
			name: "StateMob",
			in: func() interface{} {
				viper.Set("mobs.foo", object.MobData{
					Sprite: struct {
						Path string
						Size uint
					}{
						Path: "foo",
						Size: 50,
					},
					Position: movement.Location{99, 1},
				})
				return object.CreateStateMob("foo")
			},
			expected: display.CreateDisplayable("foo:0:0", 50, image.Pt(99, 1), layer.FOREGROUND),
		},
		{
			name: "Tile",
			in: func() interface{} {
				viper.Set("tile_types", map[string]struct {
					Sprite   string
					Movement int
				}{
					"plain": {
						Sprite:   "sheet:0:0",
						Movement: 10,
					},
				})
				terrain.RegisterTileTypes()
				return terrain.CreateTile(terrain.GridCoord{X: 30, Y: 30}, terrain.CreateTileType("plain"), 50)
			},
			expected: display.CreateDisplayable("sheet:0:0", 50, image.Pt(1500, 1500), layer.BACKGROUND),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := ConvertObjectToDisplayable(tt.in())
			if actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}