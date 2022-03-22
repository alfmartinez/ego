package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Story struct {
	Title string `@Title`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"Title", `"(\\"|[^"])*"`, nil},
		},
	})
)
