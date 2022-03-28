package text

import (
	"reflect"
	"strings"
	"testing"

	"github.com/alecthomas/participle/v2/lexer"
	"github.com/andreyvit/diff"
	"github.com/davecgh/go-spew/spew"
)

func Position(offset, line, col int) lexer.Position {
	return lexer.Position{
		Filename: "",
		Offset:   offset,
		Line:     line,
		Column:   col,
	}
}

type (
	args struct {
		content string
	}
	ParserCase struct {
		name    string
		content string
		want    World
		tokens  bool
	}
)

func TestParserEbnf(t *testing.T) {
	parser := BuildParser()
	def := parser.String()
	expected := `Grammar = World <eof> .
World = (Statement \".\" <eol>)* .
Statement = DescriptionPhrase VerbPhrase .
DescriptionPhrase = ComplexPhrase | SimplePhrase .
ComplexPhrase = <ident> <determiner> VerbPhrase .
VerbPhrase = <verb> DescriptionPhrase .
SimplePhrase = <ident> .`
	if def != expected {
		t.Errorf("Got EBNF Grammar \n%s", def)
	}
}

func TestParseReader(t *testing.T) {
	tests := make([]ParserCase, 0)
	tests = append(tests, informerSentences...)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.content)

			if tt.tokens {
				ParseTokens(reader)
				reader.Reset(tt.content)
			}

			got := ParseReader(reader)
			if !reflect.DeepEqual(got.World, tt.want) {
				t.Errorf("Result not as expected\n%v", diff.LineDiff(spew.Sprintln(got.World), spew.Sprintln(tt.want)))
			}
		})
	}
}
