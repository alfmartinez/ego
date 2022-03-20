package input

import (
	"ego/engine/slices"
	"testing"
)

func Test_inputHandler_Handle(t *testing.T) {
	type fields struct {
		keyStatus map[Key]bool
	}
	type args struct {
		evts []Event
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		pressed  Key
		released Key
	}{
		{
			"Escape",
			fields{
				keyStatus: make(map[Key]bool),
			},
			args{
				[]Event{
					{
						Key:    ESCAPE,
						Action: PRESSED,
					},
					{
						Key:    UP,
						Action: PRESSED,
					},
					{
						Key:    UP,
						Action: RELEASED,
					},
				},
			},
			ESCAPE,
			UP,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &inputHandler{
				keyStatus: tt.fields.keyStatus,
			}
			if !h.AllReleased() {
				t.Error("all keys should be released")
			}
			slices.Apply(tt.args.evts, func(e Event) {
				h.Handle(e)
			})
			if !h.IsPressed(tt.pressed) {
				t.Error("key should be pressed")
			}
			if h.IsPressed(tt.released) {
				t.Error("key should not be pressed")
			}
			if h.AllReleased() {
				t.Error("not all keys should be released")
			}
		})
	}
}
