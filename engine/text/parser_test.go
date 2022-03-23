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
				Inventory: Inventory{
					Items: []string{"une orange", "un stylo bic", "une serviette en papier"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Parse(tt.args.filepath)
			if !reflect.DeepEqual(got.Inventory.Items, tt.want.Inventory.Items) {
				t.Errorf("Parse() = %q, want %q", got.Inventory.Items, tt.want.Inventory.Items)
			}
		})
	}
}
