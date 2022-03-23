package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Grammar struct {
	Pos       lexer.Position
	Title     string    `@String`
	Define    []Define  `@@*`
	Inventory Inventory `@@`
}

type Define struct {
	Pos   lexer.Position
	Ident []string `@Ident+`
	Value Value    `@@`
}

type Value struct {
	Pos  lexer.Position
	Room bool `(@Room`
	Prop bool `| @Prop)`
}

type Inventory struct {
	Pos   lexer.Position
	Items []string `"Le personnage possède " @Item* "."`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"punct", `[\.\,\s]+`, nil},
			{"whitespace", `\s+`, nil},
			{"nl", `\n`, nil},
			{"String", `"(\\"|[^"])*"`, nil},
			{"Inventory", "Le personnage possède ", lexer.Push("Inventory")},
			{"Ident", `[A-Z][^\s]+`, nil},
			{"Room", `est une pièce`, nil},
			{"Prop", `est là`, nil},
			{"CatchAll", `.*`, nil},
		},
		"Inventory": {
			{"separatorComma", `(\,|\s)`, nil},
			{"And", `\bet\b`, nil},
			{"Item", `[\p{L}\s]+`, nil},
			{"InventoryEnd", `\.`, lexer.Pop()},
		},
	})
)
