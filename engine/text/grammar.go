package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Story struct {
	Title string `@String`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"String", `"(\\"|[^"])*"`, lexer.Push("Story")},
		},
		"Story": {
			{"Title", `\1`, nil},
		},
	})
)
