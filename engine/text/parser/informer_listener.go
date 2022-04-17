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

	// EnterDefinitionType is called when entering the definitionType production.
	EnterDefinitionType(c *DefinitionTypeContext)

	// EnterCertainty is called when entering the certainty production.
	EnterCertainty(c *CertaintyContext)

	// EnterDesignator is called when entering the designator production.
	EnterDesignator(c *DesignatorContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitDefinitionType is called when exiting the definitionType production.
	ExitDefinitionType(c *DefinitionTypeContext)

	// ExitCertainty is called when exiting the certainty production.
	ExitCertainty(c *CertaintyContext)

	// ExitDesignator is called when exiting the designator production.
	ExitDesignator(c *DesignatorContext)
}
