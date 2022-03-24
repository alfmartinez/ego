package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Grammar struct {
	Pos         lexer.Position
	Title       string       `@String EOL`
	Definitions []Definition `@@*`
	Inventaire  Inventory    `@@`
}

type Definition struct {
	Pos          lexer.Position
	Subject      Subject `@@ "est":Verb`
	Thing        Thing   `@@ "." EOL?`
	Descriptiion string  `@String EOL`
}

type Inventory struct {
	Pos        lexer.Position
	Subject    Subject `@@ "porte":Verb`
	Complement []Item  `("et"? @@ (",")?)+ "." EOL?`
}

type Item struct {
	Article  string   `@Ident`
	Keywords []string `@Ident+`
}

type Thing struct {
	Room bool `( "une" "pièce"`
	Prop bool `| "là")`
}

type Subject struct {
	Keywords []string `@Ident+`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"Comment", `//[^\n]*\n`, nil},
			{"EOL", `\n+`, nil},
			{"Whitespace", `\s+`, nil},
			{"Punct", `[\.\,]`, nil},
			{"String", `"(\\"|[^"])*"`, nil},
			{"Verb", `(est|porte)`, nil},
			{"Ident", `[\p{L}]+`, nil},
		},
	})
)
