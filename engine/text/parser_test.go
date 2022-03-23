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
		want *Grammar
	}{
		{
			"empty",
			args{
				"testdata/empty.txt",
			},
			&Grammar{
				Title: "Bic Example",
				Definitions: []Definition{
					{[]string{"Le", "Salon", "Pourpre"}, "est", []string{"une", "pièce"}},
					{[]string{"Le", "Canapé", "Confortable"}, "est", []string{"là"}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Parse(tt.args.filepath)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Define = %#v, want %#v", got, tt.want)
			}
		})
	}
}
