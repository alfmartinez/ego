package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Grammar struct {
	Pos   lexer.Position
	World World `@@ EOF`
}

type World struct {
	Pos       lexer.Position
	Statement []Statement `(@@ EOL)*`
}

type Statement struct {
	Pos   lexer.Position
	Title string `( @String`
	Room  Room   ` | @@)`
}

type Room struct {
	Pos      lexer.Position
	KeyWords []string `@Ident+ Room`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"Comment", `//[^\n]*\n`, nil},
			{"String", `"[^"]*"`, nil},
			{"Room", `est un lieu\.`, nil},
			{"Ident", `\p{L}+`, nil},
			{"Punct", `[\.]+`, nil},
			{"Whitespace", `[ \t]+`, nil},
			{"EOL", `\n+`, nil},
		},
	})
)
