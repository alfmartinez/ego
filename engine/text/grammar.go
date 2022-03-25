package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Grammar struct {
	Pos   lexer.Position
	World World `@@ EOF`
}

type World struct {
	Pos        lexer.Position
	Statements []Statement `(@@ EOL)*`
}

type Title struct {
	Title string `@String`
}

type Statement struct {
	Pos      lexer.Position
	Title    string   `(  @String`
	Room     Room     ` | @@ `
	Item     Item     ` | @@ `
	Describe Describe ` | @@ )`
}

type Room struct {
	Pos         lexer.Position
	Designator  Designator `@@ Is Room`
	Description string     `@String?`
}

type Item struct {
	Pos         lexer.Position
	Designator  Designator `@@ Is Here`
	Description string     `@String?`
}

type Describe struct {
	Pos         lexer.Position
	Designator  Designator `Describe @@ `
	Description string     `Is @String`
}

type Designator struct {
	KeyWords []string `@Ident+`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"Comment", `//[^\n]*\n`, nil},
			{"String", `"[^"]*"`, nil},
			{"Is", `est`, nil},
			{"Room", `un lieu\.`, nil},
			{"Here", `ici\.`, nil},
			{"Describe", `La description`, nil},
			{"Ident", `\p{L}+`, nil},
			{"Punct", `[\.]+`, nil},
			{"Whitespace", `[ \t]+`, nil},
			{"EOL", `\n+`, nil},
		},
	})
)
