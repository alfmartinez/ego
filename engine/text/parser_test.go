package text

import (
	"reflect"
	"strings"
	"testing"

	"github.com/alecthomas/participle/v2/lexer"
)

func Position(offset, line, col int) lexer.Position {
	return lexer.Position{"", offset, line, col}
}
func TestParseReader(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want World
	}{
		{
			"Title",
			args{`"Just a Title"
`},
			World{
				Position(0, 1, 1),
				[]Statement{
					{
						Pos:   Position(0, 1, 1),
						Title: "Just a Title",
					},
				},
			},
		},
		{
			"Comment",
			args{`// My Comment
"Commenting Park"
`},
			World{
				Position(14, 2, 1),
				[]Statement{
					{
						Pos:   Position(14, 2, 1),
						Title: "Commenting Park",
					},
				},
			},
		},
		{
			"A Simple Room",
			args{`"A Room With A View"
Chambord est un lieu.
`},
			World{
				Position(0, 1, 1),
				[]Statement{
					{
						Pos:   Position(0, 1, 1),
						Title: "A Room With A View",
					},
					{
						Pos: Position(21, 2, 1),
						Room: Room{
							Pos:      Position(21, 2, 1),
							KeyWords: Designator{[]string{"Chambord"}},
						},
					},
				},
			},
		},
		{
			"A Simple Room - multiple token designator",
			args{`"A Room With A View"
Château Chambord est un lieu.
`},
			World{
				Position(0, 1, 1),
				[]Statement{
					{
						Pos:   Position(0, 1, 1),
						Title: "A Room With A View",
					},
					{
						Pos: Position(21, 2, 1),
						Room: Room{
							Pos:      Position(21, 2, 1),
							KeyWords: Designator{[]string{"Château", "Chambord"}},
						},
					},
				},
			},
		},
		{
			"A Room with complex designator and a description",
			args{`"A Room With A View"
Le Château de Chambord est un lieu. "L'entrée du Château est majestueuse."
`},
			World{
				Position(0, 1, 1),
				[]Statement{
					{
						Pos:   Position(0, 1, 1),
						Title: "A Room With A View",
					},
					{
						Pos: Position(21, 2, 1),
						Room: Room{
							Pos:         Position(21, 2, 1),
							KeyWords:    Designator{[]string{"Le", "Château", "de", "Chambord"}},
							Description: "L'entrée du Château est majestueuse.",
						},
					},
				},
			},
		},
		{
			"Let's add an item",
			args{`"A Room With A View"
Le Château de Chambord est un lieu. "L'entrée du Château est majestueuse."
Un livre déchiré est ici.
`},
			World{
				Position(0, 1, 1),
				[]Statement{
					{
						Pos:   Position(0, 1, 1),
						Title: "A Room With A View",
					},
					{
						Pos: Position(21, 2, 1),
						Room: Room{
							Pos:         Position(21, 2, 1),
							KeyWords:    Designator{[]string{"Le", "Château", "de", "Chambord"}},
							Description: "L'entrée du Château est majestueuse.",
						},
					},
					{
						Pos: Position(99, 3, 1),
						Item: Item{
							Pos:      Position(99, 3, 1),
							KeyWords: []string{"Un", "livre", "déchiré"},
						},
					},
				},
			},
		},
		{
			"Let's add an item - with description",
			args{`"A Room With A View"
Le Château de Chambord est un lieu. "L'entrée du Château est majestueuse."
Un livre déchiré est ici. "Un livre déchiré semble avoir été abandonné par un lecteur peu soigneux."
`},
			World{
				Position(0, 1, 1),
				[]Statement{
					{
						Pos:   Position(0, 1, 1),
						Title: "A Room With A View",
					},
					{
						Pos: Position(21, 2, 1),
						Room: Room{
							Pos:         Position(21, 2, 1),
							KeyWords:    Designator{[]string{"Le", "Château", "de", "Chambord"}},
							Description: "L'entrée du Château est majestueuse.",
						},
					},
					{
						Pos: Position(99, 3, 1),
						Item: Item{
							Pos:         Position(99, 3, 1),
							KeyWords:    []string{"Un", "livre", "déchiré"},
							Description: "Un livre déchiré semble avoir été abandonné par un lecteur peu soigneux.",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.args.content)
			got := ParseReader(reader)
			if !reflect.DeepEqual(got.World, tt.want) {
				t.Errorf("Define = \n%#v, want \n%#v", got.World, tt.want)
			}
		})
	}
}
