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
	All   string   `@CatchAll`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"punct", `\.`, nil},
			{"nl", `\n`, nil},
			{"String", `"(\\"|[^"])*"`, nil},
			{"whitespace", `\s+`, nil},
			{"Ident", `[A-Z][^\s]+`, nil},
			{"CatchAll", `.*`, nil},
		},
	})
)
