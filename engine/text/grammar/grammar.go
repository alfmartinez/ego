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
		Test               *Test                    `  @@`
		Title              *Title                   `| @@`
		Carry              *Carry                   `| @@`
		Rulebook           *Rulebook                `| @@`
		Activity           *ActivityDefinition      `| @@`
		Instanciate        *Instanciate             `| @@`
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
		Section            *Section                 `| @@`
		SubjectIsObject    *SubjectIsObject         `| @@`
		Description        *Description             `| @@`
	}

	SubjectIsObject struct {
		Subject *Designator `@@ "is"`
		Object  *Designator `@@`
	}

	Carry struct {
		Subject *Designator   `@@ "carries"`
		Items   []*Designator `(@@ Separator?)* "."`
	}

	Title struct {
		Title  string      `@String | @Ident+`
		Author *Designator `"by" @@`
	}

	ActivityDefinition struct {
		Name *Designator `@@ "is" "activity" "."`
	}

	Rulebook struct {
		Name *Designator `@@ "is"`
		Kind string      `(@Ident "based")? "rulebook" "."`
	}

	Instanciate struct {
		Name *Designator `@@ "is"`
		Kind string      `@Ident`
		With *Property   `("with" @@)? "."`
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
		Condition *Condition `@@ (","|":")`
		Activity  *Activity  `@@ "."?`
	}

	Condition struct {
		Rulebook *Designator `@@`
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
		Direction   string      `@Ident`
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

	Designator struct {
		Elements []string `{ @Ident }`
	}
)

var (
	verbs       = []string{"is", "has", "carries", "called"}
	articles    = []string{"a", "an", "the", "The", "An", "A"}
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
