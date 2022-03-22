package text

import (
	"ego/engine/state"
	"reflect"
	"testing"
)

func TestEngine(t *testing.T) {
	type want struct {
		states state.States
	}
	tests := []struct {
		label string
		path  string
		want  struct {
			states state.States
		}
	}{
		{
			"empty file",
			"testdata/empty.txt",
			want{
				state.States{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			engine := New()
			engine.AddText(tt.path)
			got := engine.Transcribe()
			if !reflect.DeepEqual(got, tt.want.states) {
				t.Errorf("Expected %v, got %v", tt.want.states, got)
			}
		})
	}

}
