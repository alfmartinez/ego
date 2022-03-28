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
		Statements []*Statement `(@@ EOL)*`
	}

	Statement struct {
		DP          *DescriptionPhrase `(@@`
		VP          *VerbPhrase        `@@ "."`
		Description string             `@String? )`
		Title       string             `| @String`
	}

	DescriptionPhrase struct {
		Complex  *ComplexPhrase  `  @@`
		AdjNoun  *AdjNoun        `| @@`
		Relation *RelativePhrase `| @@`
		Simple   *Noun           `| @@`
	}

	AdjNoun struct {
		Adjective string `@Ident`
		Noun      *Noun  `@@`
	}

	Noun struct {
		Content string `@Ident`
	}

	ComplexPhrase struct {
		Simple *Noun       `@@`
		VP     *VerbPhrase `Determiner @@`
	}

	RelativePhrase struct {
		Simple   *Noun  `@@`
		Relation string `@Relation`
		Related  *Noun  `@@`
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
	relations   = []string{"of", "in"}

	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"String", `"[^"]*"`, nil},
			{"Relation", "(" + strings.Join(relations, "|") + `)\b`, nil},
			{"Determiner", "(" + strings.Join(determiners, "|") + `)\b`, nil},
			{"Article", "(" + strings.Join(articles, "|") + `)\b`, nil},
			{"Verb", "(" + strings.Join(verbs, "|") + `)\b`, nil},
			{"Ident", `[\p{L}]+`, nil},
			{"Punct", `\.`, nil},
			{"EOL", `\n+`, nil},
			{"Whitespace", `[ \t]+`, nil},
		},
	})
)
