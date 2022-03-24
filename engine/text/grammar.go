package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Grammar struct {
	Pos   lexer.Position
	World World `@@ EOF`
}

type World struct {
	Pos   lexer.Position
	Title string `@String EOL`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"Comment", `//[^\n]*\n`, nil},
			{"String", `"[^"]*"`, nil},
			{"Ident", `\p{L}+`, nil},
			{"Punct", `[\.]+`, nil},
			{"Whitespace", `[ \t]+`, nil},
			{"EOL", `\n+`, nil},
		},
	})
)
