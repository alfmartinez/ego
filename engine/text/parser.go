package text

import (
	"github.com/alecthomas/participle/v2"
	"os"
)

func Parse(filepath string) *Story {
	parser := participle.MustBuild(&Story{}, participle.Lexer(def), participle.Unquote("String")) //, participle.Elide("Whitespace"))

	f, _ := os.Open(filepath)

	ast := &Story{}
	err := parser.Parse(filepath, f, ast)
	if err != nil {
		panic(err)
	}
	return ast
}
