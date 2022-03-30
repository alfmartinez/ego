package grammar

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"io"
	"os"
)

var (
	parser *participle.Parser = BuildInformerParser()
)

func BuildInformerParser() *participle.Parser {
	return participle.MustBuild(
		&Grammar{},
		participle.Lexer(def),
		participle.Unquote("String"),
		participle.Elide("Article", "Whitespace", "Comment"),
		participle.CaseInsensitive("Article", "Direction"),
		participle.UseLookahead(16),
	)
}

func BuildCommandParser() *participle.Parser {
	return participle.MustBuild(
		&Command{},
		participle.Lexer(def),
		participle.Unquote("String"),
		participle.Elide("Article", "Whitespace"),
		participle.CaseInsensitive("Article", "Direction"),
		participle.UseLookahead(16),
	)
}

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

func ParseTokens(reader io.Reader) {
	tokens, _ := parser.Lex("", reader)
	fmt.Println("Tokens :")
	for _, token := range tokens {
		fmt.Printf("%#v\n", token)
	}
}
