// Code generated from Informer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Informer

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type InformerParser struct {
	*antlr.BaseParser
}

var informerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func informerParserInit() {
	staticData := &informerParserStaticData
	staticData.literalNames = []string{
		"", "'is'", "'a'", "'rulebook'", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "NUMBER", "IDENT", "PUNCT", "STRING", "COMMENT",
		"WHITESPACE",
	}
	staticData.ruleNames = []string{
		"start", "definition", "rulebook", "designator",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 10, 30, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 5,
		0, 10, 8, 0, 10, 0, 12, 0, 13, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 28, 8, 3, 1, 3, 0, 0, 4,
		0, 2, 4, 6, 0, 0, 27, 0, 11, 1, 0, 0, 0, 2, 16, 1, 0, 0, 0, 4, 18, 1, 0,
		0, 0, 6, 27, 1, 0, 0, 0, 8, 10, 3, 2, 1, 0, 9, 8, 1, 0, 0, 0, 10, 13, 1,
		0, 0, 0, 11, 9, 1, 0, 0, 0, 11, 12, 1, 0, 0, 0, 12, 14, 1, 0, 0, 0, 13,
		11, 1, 0, 0, 0, 14, 15, 5, 0, 0, 1, 15, 1, 1, 0, 0, 0, 16, 17, 3, 4, 2,
		0, 17, 3, 1, 0, 0, 0, 18, 19, 3, 6, 3, 0, 19, 20, 5, 1, 0, 0, 20, 21, 5,
		2, 0, 0, 21, 22, 5, 3, 0, 0, 22, 23, 5, 4, 0, 0, 23, 5, 1, 0, 0, 0, 24,
		28, 5, 6, 0, 0, 25, 26, 5, 6, 0, 0, 26, 28, 3, 6, 3, 0, 27, 24, 1, 0, 0,
		0, 27, 25, 1, 0, 0, 0, 28, 7, 1, 0, 0, 0, 2, 11, 27,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// InformerParserInit initializes any static state used to implement InformerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewInformerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func InformerParserInit() {
	staticData := &informerParserStaticData
	staticData.once.Do(informerParserInit)
}

// NewInformerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewInformerParser(input antlr.TokenStream) *InformerParser {
	InformerParserInit()
	this := new(InformerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &informerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Informer.g4"

	return this
}

// InformerParser tokens.
const (
	InformerParserEOF        = antlr.TokenEOF
	InformerParserT__0       = 1
	InformerParserT__1       = 2
	InformerParserT__2       = 3
	InformerParserT__3       = 4
	InformerParserNUMBER     = 5
	InformerParserIDENT      = 6
	InformerParserPUNCT      = 7
	InformerParserSTRING     = 8
	InformerParserCOMMENT    = 9
	InformerParserWHITESPACE = 10
)

// InformerParser rules.
const (
	InformerParserRULE_start      = 0
	InformerParserRULE_definition = 1
	InformerParserRULE_rulebook   = 2
	InformerParserRULE_designator = 3
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(InformerParserEOF, 0)
}

func (s *StartContext) AllDefinition() []IDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefinitionContext); ok {
			tst[i] = t.(IDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) Definition(i int) IDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *InformerParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, InformerParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == InformerParserIDENT {
		{
			p.SetState(8)
			p.Definition()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(14)
		p.Match(InformerParserEOF)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) CopyFrom(ctx *DefinitionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RulebookDeclarationContext struct {
	*DefinitionContext
}

func NewRulebookDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RulebookDeclarationContext {
	var p = new(RulebookDeclarationContext)

	p.DefinitionContext = NewEmptyDefinitionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DefinitionContext))

	return p
}

func (s *RulebookDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulebookDeclarationContext) Rulebook() IRulebookContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRulebookContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRulebookContext)
}

func (s *RulebookDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterRulebookDeclaration(s)
	}
}

func (s *RulebookDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitRulebookDeclaration(s)
	}
}

func (p *InformerParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InformerParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewRulebookDeclarationContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Rulebook()
	}

	return localctx
}

// IRulebookContext is an interface to support dynamic dispatch.
type IRulebookContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulebookContext differentiates from other interfaces.
	IsRulebookContext()
}

type RulebookContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulebookContext() *RulebookContext {
	var p = new(RulebookContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_rulebook
	return p
}

func (*RulebookContext) IsRulebookContext() {}

func NewRulebookContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulebookContext {
	var p = new(RulebookContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_rulebook

	return p
}

func (s *RulebookContext) GetParser() antlr.Parser { return s.parser }

func (s *RulebookContext) Designator() IDesignatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDesignatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *RulebookContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulebookContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulebookContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterRulebook(s)
	}
}

func (s *RulebookContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitRulebook(s)
	}
}

func (p *InformerParser) Rulebook() (localctx IRulebookContext) {
	this := p
	_ = this

	localctx = NewRulebookContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, InformerParserRULE_rulebook)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Designator()
	}
	{
		p.SetState(19)
		p.Match(InformerParserT__0)
	}
	{
		p.SetState(20)
		p.Match(InformerParserT__1)
	}
	{
		p.SetState(21)
		p.Match(InformerParserT__2)
	}
	{
		p.SetState(22)
		p.Match(InformerParserT__3)
	}

	return localctx
}

// IDesignatorContext is an interface to support dynamic dispatch.
type IDesignatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDesignatorContext differentiates from other interfaces.
	IsDesignatorContext()
}

type DesignatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDesignatorContext() *DesignatorContext {
	var p = new(DesignatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_designator
	return p
}

func (*DesignatorContext) IsDesignatorContext() {}

func NewDesignatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DesignatorContext {
	var p = new(DesignatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_designator

	return p
}

func (s *DesignatorContext) GetParser() antlr.Parser { return s.parser }

func (s *DesignatorContext) IDENT() antlr.TerminalNode {
	return s.GetToken(InformerParserIDENT, 0)
}

func (s *DesignatorContext) Designator() IDesignatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDesignatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *DesignatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DesignatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DesignatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterDesignator(s)
	}
}

func (s *DesignatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitDesignator(s)
	}
}

func (p *InformerParser) Designator() (localctx IDesignatorContext) {
	this := p
	_ = this

	localctx = NewDesignatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, InformerParserRULE_designator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Match(InformerParserIDENT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Match(InformerParserIDENT)
		}
		{
			p.SetState(26)
			p.Designator()
		}

	}

	return localctx
}
