// Code generated from /Users/alex/Projects/ego/engine/text/Informer.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

// EnterStatement is called when production statement is entered.
func (s *BaseInformerListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseInformerListener) ExitStatement(ctx *StatementContext) {}

// EnterRulebook is called when production rulebook is entered.
func (s *BaseInformerListener) EnterRulebook(ctx *RulebookContext) {}

// ExitRulebook is called when production rulebook is exited.
func (s *BaseInformerListener) ExitRulebook(ctx *RulebookContext) {}

// EnterActivity is called when production activity is entered.
func (s *BaseInformerListener) EnterActivity(ctx *ActivityContext) {}

// ExitActivity is called when production activity is exited.
func (s *BaseInformerListener) ExitActivity(ctx *ActivityContext) {}

// EnterDesignator is called when production designator is entered.
func (s *BaseInformerListener) EnterDesignator(ctx *DesignatorContext) {}

// ExitDesignator is called when production designator is exited.
func (s *BaseInformerListener) ExitDesignator(ctx *DesignatorContext) {}
