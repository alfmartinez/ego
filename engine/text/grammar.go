package text

import (
	"github.com/alecthomas/participle/v2/lexer"
	"strings"
)

// Defines a grammar for natural text parsing.
// 1. 	S = DP + VP 			: The Chamber is a room.
// 2. 	VP = Verb + DP			: is a room
// 3a. 	DP = DP + which + VP	: The cart is a vehicle which is movable.
// 3b. 	DP = DP + who + VP		: The Janitor is a male person who has the blue key.
// 4. 	DP = DP + RP 			: An open container in the ballroom
// 5a. 	RP = Preposition + DP 	: in a container
// 5b. 	RP = Participle + DP 	: worn by Mr Darcy
// 6 	S = RP + Verb + DP 		: In the Garden is a sunflower

type (
	// Define a Grammar for parsing
	Grammar struct {
		//
		World World `@@ EOF`
	}

	World struct {
		Statements []*Statement `(@@ "." EOL)*`
	}

	Statement struct {
		DP *DescriptionPhrase `@@`
		VP *VerbPhrase        `@@`
	}

	DescriptionPhrase struct {
		Complex *ComplexPhrase `  @@`
		Simple  *SimplePhrase  `| @@`
	}

	SimplePhrase struct {
		Content string `@Ident`
	}

	ComplexPhrase struct {
		Simple *SimplePhrase `@@`
		VP     *VerbPhrase   `Determiner @@`
	}

	VerbPhrase struct {
		Verb string             `@Verb`
		DP   *DescriptionPhrase `@@`
	}
)

var (
	verbs       = []string{"is", "has"}
	articles    = []string{"a", "an", "the", "The", "An", "A"}
	determiners = []string{"which", "who"}

	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"Punct", `\.`, nil},
			{"EOL", `\n+`, nil},
			{"Whitespace", `[ \t]+`, nil},
			{"Determiner", "(" + strings.Join(determiners, "|") + ")", nil},
			{"Article", "(" + strings.Join(articles, "|") + ")", nil},
			{"Verb", "(" + strings.Join(verbs, "|") + ")", nil},
			{"Ident", `[\p{L}]+`, nil},
		},
	})
)
