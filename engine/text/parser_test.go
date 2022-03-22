package text

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want *Story
	}{
		{
			"empty",
			args{
				"testdata/empty.txt",
			},
			&Story{
				Title: "Bic Example",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.filepath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
