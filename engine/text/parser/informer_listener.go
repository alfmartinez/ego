// Code generated from /Users/alex/Projects/ego/engine/text/Informer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Informer
import "github.com/antlr/antlr4/runtime/Go/antlr"

// InformerListener is a complete listener for a parse tree produced by InformerParser.
type InformerListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterRulebook is called when entering the Rulebook production.
	EnterRulebook(c *RulebookContext)

	// EnterObjectBasedRulebook is called when entering the ObjectBasedRulebook production.
	EnterObjectBasedRulebook(c *ObjectBasedRulebookContext)

	// EnterActivity is called when entering the Activity production.
	EnterActivity(c *ActivityContext)

	// EnterCertaintyOfAttribute is called when entering the CertaintyOfAttribute production.
	EnterCertaintyOfAttribute(c *CertaintyOfAttributeContext)

	// EnterObjectKind is called when entering the ObjectKind production.
	EnterObjectKind(c *ObjectKindContext)

	// EnterValueKind is called when entering the ValueKind production.
	EnterValueKind(c *ValueKindContext)

	// EnterCertainty is called when entering the certainty production.
	EnterCertainty(c *CertaintyContext)

	// EnterDesignator is called when entering the designator production.
	EnterDesignator(c *DesignatorContext)

	// EnterValues is called when entering the values production.
	EnterValues(c *ValuesContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitRulebook is called when exiting the Rulebook production.
	ExitRulebook(c *RulebookContext)

	// ExitObjectBasedRulebook is called when exiting the ObjectBasedRulebook production.
	ExitObjectBasedRulebook(c *ObjectBasedRulebookContext)

	// ExitActivity is called when exiting the Activity production.
	ExitActivity(c *ActivityContext)

	// ExitCertaintyOfAttribute is called when exiting the CertaintyOfAttribute production.
	ExitCertaintyOfAttribute(c *CertaintyOfAttributeContext)

	// ExitObjectKind is called when exiting the ObjectKind production.
	ExitObjectKind(c *ObjectKindContext)

	// ExitValueKind is called when exiting the ValueKind production.
	ExitValueKind(c *ValueKindContext)

	// ExitCertainty is called when exiting the certainty production.
	ExitCertainty(c *CertaintyContext)

	// ExitDesignator is called when exiting the designator production.
	ExitDesignator(c *DesignatorContext)

	// ExitValues is called when exiting the values production.
	ExitValues(c *ValuesContext)
}
