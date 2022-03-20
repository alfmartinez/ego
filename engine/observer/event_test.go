package observer

import (
	"reflect"
	"testing"
)

func Test_event_Data(t *testing.T) {
	type fields struct {
		eventType EventType
		data      string
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{
			"Simple",
			fields{
				eventType: RENDER,
			},
			nil,
		},
		{
			"Simple",
			fields{
				eventType: RENDER,
				data:      "foo",
			},
			[]any{"foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e Event
			if tt.fields.data != "" {
				e = CreateEvent(tt.fields.eventType, tt.fields.data)
			} else {
				e = CreateEvent(tt.fields.eventType)
			}

			if got := e.Data(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("event.Data() = %v, want %v", got, tt.want)
			}
		})
	}
}
