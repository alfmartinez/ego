package layer

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestLayerMap_Add(t *testing.T) {
	type args struct {
		layer Layer
		d     interface{}
	}
	tests := []struct {
		name string
		m    LayerMap
		args args
	}{
		{
			"empty map",
			CreateLayerMap(),
			args{
				BACKGROUND,
				"bob",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Add(tt.args.layer, tt.args.d)
			if slices.IndexFunc(tt.m[BACKGROUND], func(e any) bool {
				return tt.args.d == e.(string)
			}) == -1 {
				t.Error("background layer should contain added object")
			}
			if slices.IndexFunc(tt.m[FOREGROUND], func(e any) bool {
				return tt.args.d == e.(string)
			}) != -1 {
				t.Error("foreground layer should not contain added object")
			}
		})
	}
}
