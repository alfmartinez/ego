package text

import (
	"github.com/alecthomas/participle/v2"
	"os"
)

func Parse(filepath string) *Grammar {
	parser := participle.MustBuild(
		&Grammar{},
		participle.Lexer(def),
		participle.Unquote("String"),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)

	f, _ := os.Open(filepath)

	ast := &Grammar{}
	err := parser.Parse(filepath, f, ast) //	participle.AllowTrailing(true),

	if err != nil {
		panic(err)
	}
	return ast
}
