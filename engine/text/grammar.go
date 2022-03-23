package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Grammar struct {
	Pos    lexer.Position
	Title  string   `@String`
	Define []Define `@@*`
}

type Define struct {
	Pos   lexer.Position
	Ident []string `@Ident+`
	Value Value    `@@`
}

type Value struct {
	Room bool `@Room`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"punct", `\.`, nil},
			{"whitespace", `\s+`, nil},
			{"nl", `\n`, nil},
			{"String", `"(\\"|[^"])*"`, nil},
			{"Ident", `[A-Z][^\s]+`, nil},
			{"Room", `est une pièce`, nil},
			{"CatchAll", `.*`, nil},
		},
	})
)
