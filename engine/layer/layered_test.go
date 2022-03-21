package layer

import (
	"reflect"
	"testing"
)

func TestCreateLayered(t *testing.T) {
	tests := []struct {
		name string
		want Layered
	}{
		{"simple", &layered{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateLayered()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLayered() = %v, want %v", got, tt.want)
			}
			if lGot := got.Layer(); lGot != NONE {
				t.Error("default layer should be none")
			}
			got.SetLayer(FOREGROUND)
			if lGot := got.Layer(); lGot != FOREGROUND {
				t.Error("layered should remember")
			}
		})
	}
}
