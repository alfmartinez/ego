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

type Title struct {
	Title string `@String`
}

type Statement struct {
	Pos   lexer.Position
	Title string `(  @String`
	Room  Room   ` | @@ `
	Item  Item   ` | @@ )`
}

type Room struct {
	Pos         lexer.Position
	Designator  Designator `@@ Room`
	Description string     `@String?`
}

type Item struct {
	Pos         lexer.Position
	Designator  Designator `@@ Item`
	Description string     `@String?`
}

type Designator struct {
	KeyWords []string `@Ident+`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"Comment", `//[^\n]*\n`, nil},
			{"String", `"[^"]*"`, nil},
			{"Item", `est ici\.`, nil},
			{"Room", `est un lieu\.`, lexer.Push("Room")},
			{"Ident", `\p{L}+`, nil},
			{"Punct", `[\.]+`, nil},
			{"Whitespace", `[ \t]+`, nil},
			{"EOL", `\n+`, nil},
		},
		"Room": {
			lexer.Return(),
		},
	})
)
