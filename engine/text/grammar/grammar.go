package grammar

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
		Statements []*Statement `(@@ EOL?)*`
	}

	Statement struct {
		Title              string                   `@String`
		ActionDefinition   *ActionDefinition        `| @@`
		Understand         *Understand              `| @@`
		WhenDeClaration    *WhenDeClaration         `| @@`
		TextPropertyKind   *SetTextPropertyOfKind   `| @@`
		TextPropertyObject *SetTextPropertyOfObject `| @@`
		EitherProperty     *EitherProperty          `| @@`
		QuickProperty      *QuickProperty           `| @@`
		CreateInPlace      *CreateInPlace           `| @@`
		Property           *PropertyDefinition      `| @@`
		Certainty          *CertaintyDefinition     `| @@`
		ValueDefinition    *ValueDefinition         `| @@`
		KindDefinition     *KindDefinition          `| @@`
		Direction          *Connector               `| @@`
		Sentence           *Sentence                `| @@`
		Section            *Section                 `| @@`
		Test               *Test                    `| @@`
		Description        *Description             `| @@`
	}

	ActionDefinition struct {
		Name   *Designator `@@ "is" "action"`
		Target *Designator `("applying" "to" @@ )?"."`
	}

	Understand struct {
		Alias  string      `"Understand" @String`
		Target *Designator `"as" @@ "."`
	}

	WhenDeClaration struct {
		Condition *Condition `"When" @@ (","|":")`
		Activity  *Activity  `@@ "."?`
	}

	Condition struct {
		Rule *Designator `@@`
	}

	Activity struct {
		Say        string      `( "say" @String`
		Launch     *Launch     `|  @@  `
		Activities []*Activity `| (EOL @@)* EOL)`
	}

	Launch struct {
		Action   string      `@Ident`
		Argument *Designator `@@`
	}

	SetTextPropertyOfKind struct {
		PropertyName *Designator `@@ "of"`
		Target       *Designator `@@ "is"`
		Certainty    string      `@Certainty`
		Value        string      `@String "."`
	}

	SetTextPropertyOfObject struct {
		Target       *Designator `@@ "has"`
		PropertyName *Designator `@@`
		Value        string      `@String "."?`
	}

	EitherProperty struct {
		Kind   *Designator `@@ "is" Either`
		Values []string    `@Ident "or" @Ident "."`
	}

	QuickProperty struct {
		Values []string `Pronoun "is" (@Ident Separator?)+ "."`
	}

	CreateInPlace struct {
		Place *Designator `"In" @@`
		Thing *Designator `"is" @@ "."`
	}

	PropertyDefinition struct {
		Kind *Designator `@@`
		Type string      `"has" "some" @Ident`
		Name *Designator `"called" @@ "."`
	}

	CertaintyDefinition struct {
		Name      *Designator `@@ "is"`
		Certainty string      `@Certainty`
		Values    []string    `(@Ident Separator?)* "."`
	}

	ValueDefinition struct {
		Name *Designator `@@ "is" Kind "of" "value"`
		With []string    `("with" "values" (@Ident Separator?)+)? "."`
	}

	KindDefinition struct {
		Name   *Designator `@@ "is" Kind "of"`
		Parent *Designator `@@`
		With   []*Property `("with" (@@ Separator?)*)? "."?`
	}

	Property struct {
		Property *Designator `@@`
		Value    string      `@String`
	}

	Description struct {
		Target      *Designator `"description" "of" @@`
		Description string      `"is" @String "."?`
	}

	Connector struct {
		Direction   *Designator `@@`
		Origin      *Designator `"of" @@ "is"`
		Target      *Designator `@@ "."`
		Description string      `@String?`
	}

	Test struct {
		Commands []string `Test {@String}`
	}

	Section struct {
		Number int    `Section @Number`
		Title  string `"-" (@String | @Ident)`
		Tag    string `("-" (@String | @Ident*))? `
	}

	Sentence struct {
		DP          *DescriptionPhrase `@@`
		VP          *VerbPhrase        `@@ "."`
		Description string             `@String?`
	}

	DescriptionPhrase struct {
		Designator *Designator     ` @@ `
		List       *List           ` @@?`
		Complex    *ComplexPhrase  ` @@?`
		Relation   *RelativePhrase ` @@?`
	}

	List struct {
		Elements []*Designator `("," @@)*`
		Last     *Designator   `"and" @@`
	}

	Designator struct {
		Elements []string `{ @Ident }`
	}

	ComplexPhrase struct {
		VP *VerbPhrase `Determiner @@`
	}

	RelativePhrase struct {
		Relation string      `@Relation`
		Related  *Designator `@@`
	}

	VerbPhrase struct {
		Verb string             `@Verb`
		DP   *DescriptionPhrase `@@`
	}
)

var (
	verbs       = []string{"is", "has", "carries", "look", "called"}
	articles    = []string{"a", "an", "the", "The", "An", "A"}
	determiners = []string{"which", "who"}
	relations   = []string{"of", "in", "with", "In"}
	certainties = []string{
		"always", "usually", "seldom", "never",
	}
	pronouns = []string{
		"he", "she", "it", "He", "She", "It",
	}

	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"String", `"[^"]*"`, nil},
			{"Separator", `(\,|and\b)`, nil},
			{"Section", `Section\b`, nil},
			{"Test", `Test me with`, nil},
			{"Kind", `kind\b`, nil},
			{"Comment", `\[(.*?)\]\n+`, nil},
			{"Either", `either\b`, nil},
			{"Pronoun", "(" + strings.Join(pronouns, "|") + `)\b`, nil},
			{"Certainty", "(" + strings.Join(certainties, "|") + `)\b`, nil},
			{"Relation", "(" + strings.Join(relations, "|") + `)\b`, nil},
			{"Determiner", "(" + strings.Join(determiners, "|") + `)\b`, nil},
			{"Article", "(" + strings.Join(articles, "|") + `)\b`, nil},
			{"Verb", "(" + strings.Join(verbs, "|") + `)\b`, nil},
			{"Ident", `[\p{L}\-]+`, nil},
			{"Number", `[0-9]+`, nil},
			{"Punct", `[\.\-\(\)\:]`, nil},
			{"EOL", `\n+`, nil},
			{"Whitespace", `[ \t]+`, nil},
		},
	})
)

func (d *Designator) Get() string {
	return strings.ToLower(strings.Join(d.Elements, " "))
}

func (d *Designator) GetCase() string {
	return strings.Join(d.Elements, " ")
}
