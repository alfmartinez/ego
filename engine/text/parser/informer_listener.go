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

	// EnterRulebook is called when entering the rulebook production.
	EnterRulebook(c *RulebookContext)

	// EnterActivity is called when entering the activity production.
	EnterActivity(c *ActivityContext)

	// EnterDesignator is called when entering the designator production.
	EnterDesignator(c *DesignatorContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitRulebook is called when exiting the rulebook production.
	ExitRulebook(c *RulebookContext)

	// ExitActivity is called when exiting the activity production.
	ExitActivity(c *ActivityContext)

	// ExitDesignator is called when exiting the designator production.
	ExitDesignator(c *DesignatorContext)
}
