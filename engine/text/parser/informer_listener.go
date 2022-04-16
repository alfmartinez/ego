// Code generated from Informer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Informer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InformerListener is a complete listener for a parse tree produced by InformerParser.
type InformerListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterRulebookDeclaration is called when entering the RulebookDeclaration production.
	EnterRulebookDeclaration(c *RulebookDeclarationContext)

	// EnterRulebook is called when entering the rulebook production.
	EnterRulebook(c *RulebookContext)

	// EnterDesignator is called when entering the designator production.
	EnterDesignator(c *DesignatorContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitRulebookDeclaration is called when exiting the RulebookDeclaration production.
	ExitRulebookDeclaration(c *RulebookDeclarationContext)

	// ExitRulebook is called when exiting the rulebook production.
	ExitRulebook(c *RulebookContext)

	// ExitDesignator is called when exiting the designator production.
	ExitDesignator(c *DesignatorContext)
}
