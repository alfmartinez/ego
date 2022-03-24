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
			"Title",
			args{`"Just a Title"
`},
			&Grammar{
				Pos:   lexer.Position{"", 0, 1, 1},
				World: World{"Just a Title"},
			},
		},
		{
			"Comment",
			args{`// My Comment
"Commenting Park"
`},
			&Grammar{
				Pos:   lexer.Position{"", 14, 2, 1},
				World: World{"Commenting Park"},
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
