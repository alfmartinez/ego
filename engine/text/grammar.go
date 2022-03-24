package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Grammar struct {
	Pos   lexer.Position
	World World `@@ EOF`
}

type World struct {
	Title string `@String`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"String", `"[^"]*"`, nil},
		},
	})
)
