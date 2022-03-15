package sequence

import (
	"reflect"
	"testing"
)

func TestCreateSequence(t *testing.T) {
	type args struct {
		t SequenceType
	}
	tests := []struct {
		name string
		args args
		want SequenceFactory
	}{
		{
			name: "Evaluate",
			args: args{
				t: Evaluate,
			},
			want: CreateEvaluateCommand,
		}, {
			name: "Seek",
			args: args{
				t: Seek,
			},
			want: CreateSeekAndUseCommand,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSequence(tt.args.t); reflect.ValueOf(got).Pointer() != reflect.ValueOf(tt.want).Pointer() {
				t.Errorf("CreateSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
