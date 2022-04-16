// Code generated from Informer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Informer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInformerListener is a complete listener for a parse tree produced by InformerParser.
type BaseInformerListener struct{}

var _ InformerListener = &BaseInformerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInformerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInformerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInformerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInformerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseInformerListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseInformerListener) ExitStart(ctx *StartContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseInformerListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseInformerListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterTitle is called when production title is entered.
func (s *BaseInformerListener) EnterTitle(ctx *TitleContext) {}

// ExitTitle is called when production title is exited.
func (s *BaseInformerListener) ExitTitle(ctx *TitleContext) {}

// EnterGenericRulebook is called when production GenericRulebook is entered.
func (s *BaseInformerListener) EnterGenericRulebook(ctx *GenericRulebookContext) {}

// ExitGenericRulebook is called when production GenericRulebook is exited.
func (s *BaseInformerListener) ExitGenericRulebook(ctx *GenericRulebookContext) {}

// EnterSpecificRulebook is called when production SpecificRulebook is entered.
func (s *BaseInformerListener) EnterSpecificRulebook(ctx *SpecificRulebookContext) {}

// ExitSpecificRulebook is called when production SpecificRulebook is exited.
func (s *BaseInformerListener) ExitSpecificRulebook(ctx *SpecificRulebookContext) {}

// EnterActivityDeclaration is called when production activityDeclaration is entered.
func (s *BaseInformerListener) EnterActivityDeclaration(ctx *ActivityDeclarationContext) {}

// ExitActivityDeclaration is called when production activityDeclaration is exited.
func (s *BaseInformerListener) ExitActivityDeclaration(ctx *ActivityDeclarationContext) {}

// EnterActionDeclaration is called when production actionDeclaration is entered.
func (s *BaseInformerListener) EnterActionDeclaration(ctx *ActionDeclarationContext) {}

// ExitActionDeclaration is called when production actionDeclaration is exited.
func (s *BaseInformerListener) ExitActionDeclaration(ctx *ActionDeclarationContext) {}

// EnterAliasDeclaration is called when production aliasDeclaration is entered.
func (s *BaseInformerListener) EnterAliasDeclaration(ctx *AliasDeclarationContext) {}

// ExitAliasDeclaration is called when production aliasDeclaration is exited.
func (s *BaseInformerListener) ExitAliasDeclaration(ctx *AliasDeclarationContext) {}

// EnterDesignator is called when production designator is entered.
func (s *BaseInformerListener) EnterDesignator(ctx *DesignatorContext) {}

// ExitDesignator is called when production designator is exited.
func (s *BaseInformerListener) ExitDesignator(ctx *DesignatorContext) {}
