package text

import (
	"github.com/alecthomas/participle/v2"
	"io"
	"os"
)

var (
	parser *participle.Parser = participle.MustBuild(
		&Grammar{},
		participle.Lexer(def),
		participle.Unquote("String"),
		participle.Elide("Comment", "Whitespace"),
		//participle.CaseInsensitive("Ident"),
		//participle.UseLookahead(1),
	)
)

func ParseFile(filepath string) *Grammar {
	f, _ := os.Open(filepath)
	return ParseReader(f)
}

func ParseReader(reader io.Reader) *Grammar {
	ast := &Grammar{}
	err := parser.Parse("", reader, ast) //	participle.AllowTrailing(true),

	if err != nil {
		panic(err)
	}
	return ast
}
