package text

import (
	"reflect"
	"strings"
	"testing"

	"github.com/alecthomas/participle/v2/lexer"
)

func TestParseReader(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want *Grammar
	}{
		{
			"Empty",
			args{`"Test"`},
			&Grammar{
				Pos:   lexer.Position{"", 0, 1, 1},
				World: World{"Test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.args.content)
			got := ParseReader(reader)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Define = %#v, want %#v", got, tt.want)
			}
		})
	}
}
