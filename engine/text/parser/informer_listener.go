// Code generated from Informer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Informer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InformerListener is a complete listener for a parse tree produced by InformerParser.
type InformerListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterTitle is called when entering the title production.
	EnterTitle(c *TitleContext)

	// EnterGenericRulebook is called when entering the GenericRulebook production.
	EnterGenericRulebook(c *GenericRulebookContext)

	// EnterSpecificRulebook is called when entering the SpecificRulebook production.
	EnterSpecificRulebook(c *SpecificRulebookContext)

	// EnterActivityDeclaration is called when entering the activityDeclaration production.
	EnterActivityDeclaration(c *ActivityDeclarationContext)

	// EnterActionDeclaration is called when entering the actionDeclaration production.
	EnterActionDeclaration(c *ActionDeclarationContext)

	// EnterAliasDeclaration is called when entering the aliasDeclaration production.
	EnterAliasDeclaration(c *AliasDeclarationContext)

	// EnterDesignator is called when entering the designator production.
	EnterDesignator(c *DesignatorContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitTitle is called when exiting the title production.
	ExitTitle(c *TitleContext)

	// ExitGenericRulebook is called when exiting the GenericRulebook production.
	ExitGenericRulebook(c *GenericRulebookContext)

	// ExitSpecificRulebook is called when exiting the SpecificRulebook production.
	ExitSpecificRulebook(c *SpecificRulebookContext)

	// ExitActivityDeclaration is called when exiting the activityDeclaration production.
	ExitActivityDeclaration(c *ActivityDeclarationContext)

	// ExitActionDeclaration is called when exiting the actionDeclaration production.
	ExitActionDeclaration(c *ActionDeclarationContext)

	// ExitAliasDeclaration is called when exiting the aliasDeclaration production.
	ExitAliasDeclaration(c *AliasDeclarationContext)

	// ExitDesignator is called when exiting the designator production.
	ExitDesignator(c *DesignatorContext)
}
