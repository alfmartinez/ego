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
		"", "'is'", "'a'", "'rulebook'", "'.'", "'an'", "'based'", "'activity'",
		"'action'", "'applying'", "'to'", "'Understand'", "'as'", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "NUMBER", "PUNCT",
		"STRING", "COMMENT", "WHITESPACE", "EOL", "WORD", "ANY",
	}
	staticData.ruleNames = []string{
		"start", "definition", "title", "rulebookDeclaration", "activityDeclaration",
		"actionDeclaration", "aliasDeclaration", "designator",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 90, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 4, 0, 18, 8, 0, 11, 0, 12,
		0, 19, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 29, 8, 1, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 50, 8, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 85, 8, 7, 10, 7, 12, 7, 88, 9, 7,
		1, 7, 0, 1, 14, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 1, 2, 0, 2, 2, 5, 5, 89,
		0, 17, 1, 0, 0, 0, 2, 28, 1, 0, 0, 0, 4, 30, 1, 0, 0, 0, 6, 49, 1, 0, 0,
		0, 8, 51, 1, 0, 0, 0, 10, 58, 1, 0, 0, 0, 12, 69, 1, 0, 0, 0, 14, 76, 1,
		0, 0, 0, 16, 18, 3, 2, 1, 0, 17, 16, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19,
		17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 22, 5, 0, 0,
		1, 22, 1, 1, 0, 0, 0, 23, 29, 3, 8, 4, 0, 24, 29, 3, 6, 3, 0, 25, 29, 3,
		10, 5, 0, 26, 29, 3, 12, 6, 0, 27, 29, 3, 4, 2, 0, 28, 23, 1, 0, 0, 0,
		28, 24, 1, 0, 0, 0, 28, 25, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 27, 1,
		0, 0, 0, 29, 3, 1, 0, 0, 0, 30, 31, 5, 16, 0, 0, 31, 32, 5, 19, 0, 0, 32,
		5, 1, 0, 0, 0, 33, 34, 3, 14, 7, 0, 34, 35, 5, 1, 0, 0, 35, 36, 5, 2, 0,
		0, 36, 37, 5, 3, 0, 0, 37, 38, 5, 4, 0, 0, 38, 39, 5, 19, 0, 0, 39, 50,
		1, 0, 0, 0, 40, 41, 3, 14, 7, 0, 41, 42, 5, 1, 0, 0, 42, 43, 7, 0, 0, 0,
		43, 44, 5, 20, 0, 0, 44, 45, 5, 6, 0, 0, 45, 46, 5, 3, 0, 0, 46, 47, 5,
		4, 0, 0, 47, 48, 5, 19, 0, 0, 48, 50, 1, 0, 0, 0, 49, 33, 1, 0, 0, 0, 49,
		40, 1, 0, 0, 0, 50, 7, 1, 0, 0, 0, 51, 52, 3, 14, 7, 0, 52, 53, 5, 1, 0,
		0, 53, 54, 5, 5, 0, 0, 54, 55, 5, 7, 0, 0, 55, 56, 5, 4, 0, 0, 56, 57,
		5, 19, 0, 0, 57, 9, 1, 0, 0, 0, 58, 59, 3, 14, 7, 0, 59, 60, 5, 1, 0, 0,
		60, 61, 5, 5, 0, 0, 61, 62, 5, 8, 0, 0, 62, 63, 5, 9, 0, 0, 63, 64, 5,
		10, 0, 0, 64, 65, 5, 2, 0, 0, 65, 66, 5, 20, 0, 0, 66, 67, 5, 4, 0, 0,
		67, 68, 5, 19, 0, 0, 68, 11, 1, 0, 0, 0, 69, 70, 5, 11, 0, 0, 70, 71, 5,
		16, 0, 0, 71, 72, 5, 12, 0, 0, 72, 73, 5, 20, 0, 0, 73, 74, 5, 4, 0, 0,
		74, 75, 5, 19, 0, 0, 75, 13, 1, 0, 0, 0, 76, 77, 6, 7, -1, 0, 77, 78, 5,
		20, 0, 0, 78, 86, 1, 0, 0, 0, 79, 80, 10, 3, 0, 0, 80, 85, 5, 20, 0, 0,
		81, 82, 10, 2, 0, 0, 82, 83, 5, 13, 0, 0, 83, 85, 5, 20, 0, 0, 84, 79,
		1, 0, 0, 0, 84, 81, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0,
		86, 87, 1, 0, 0, 0, 87, 15, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 5, 19, 28,
		49, 84, 86,
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
	InformerParserT__4       = 5
	InformerParserT__5       = 6
	InformerParserT__6       = 7
	InformerParserT__7       = 8
	InformerParserT__8       = 9
	InformerParserT__9       = 10
	InformerParserT__10      = 11
	InformerParserT__11      = 12
	InformerParserT__12      = 13
	InformerParserNUMBER     = 14
	InformerParserPUNCT      = 15
	InformerParserSTRING     = 16
	InformerParserCOMMENT    = 17
	InformerParserWHITESPACE = 18
	InformerParserEOL        = 19
	InformerParserWORD       = 20
	InformerParserANY        = 21
)

// InformerParser rules.
const (
	InformerParserRULE_start               = 0
	InformerParserRULE_definition          = 1
	InformerParserRULE_title               = 2
	InformerParserRULE_rulebookDeclaration = 3
	InformerParserRULE_activityDeclaration = 4
	InformerParserRULE_actionDeclaration   = 5
	InformerParserRULE_aliasDeclaration    = 6
	InformerParserRULE_designator          = 7
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<InformerParserT__10)|(1<<InformerParserSTRING)|(1<<InformerParserWORD))) != 0) {
		{
			p.SetState(16)
			p.Definition()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(21)
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

func (s *DefinitionContext) ActivityDeclaration() IActivityDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActivityDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActivityDeclarationContext)
}

func (s *DefinitionContext) RulebookDeclaration() IRulebookDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRulebookDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRulebookDeclarationContext)
}

func (s *DefinitionContext) ActionDeclaration() IActionDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionDeclarationContext)
}

func (s *DefinitionContext) AliasDeclaration() IAliasDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasDeclarationContext)
}

func (s *DefinitionContext) Title() ITitleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITitleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITitleContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitDefinition(s)
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

	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.ActivityDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.RulebookDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(25)
			p.ActionDeclaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(26)
			p.AliasDeclaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(27)
			p.Title()
		}

	}

	return localctx
}

// ITitleContext is an interface to support dynamic dispatch.
type ITitleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTitleContext differentiates from other interfaces.
	IsTitleContext()
}

type TitleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTitleContext() *TitleContext {
	var p = new(TitleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_title
	return p
}

func (*TitleContext) IsTitleContext() {}

func NewTitleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TitleContext {
	var p = new(TitleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_title

	return p
}

func (s *TitleContext) GetParser() antlr.Parser { return s.parser }

func (s *TitleContext) STRING() antlr.TerminalNode {
	return s.GetToken(InformerParserSTRING, 0)
}

func (s *TitleContext) EOL() antlr.TerminalNode {
	return s.GetToken(InformerParserEOL, 0)
}

func (s *TitleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TitleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TitleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterTitle(s)
	}
}

func (s *TitleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitTitle(s)
	}
}

func (p *InformerParser) Title() (localctx ITitleContext) {
	this := p
	_ = this

	localctx = NewTitleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, InformerParserRULE_title)

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
		p.SetState(30)
		p.Match(InformerParserSTRING)
	}
	{
		p.SetState(31)
		p.Match(InformerParserEOL)
	}

	return localctx
}

// IRulebookDeclarationContext is an interface to support dynamic dispatch.
type IRulebookDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulebookDeclarationContext differentiates from other interfaces.
	IsRulebookDeclarationContext()
}

type RulebookDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulebookDeclarationContext() *RulebookDeclarationContext {
	var p = new(RulebookDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_rulebookDeclaration
	return p
}

func (*RulebookDeclarationContext) IsRulebookDeclarationContext() {}

func NewRulebookDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulebookDeclarationContext {
	var p = new(RulebookDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_rulebookDeclaration

	return p
}

func (s *RulebookDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *RulebookDeclarationContext) CopyFrom(ctx *RulebookDeclarationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *RulebookDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulebookDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GenericRulebookContext struct {
	*RulebookDeclarationContext
}

func NewGenericRulebookContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GenericRulebookContext {
	var p = new(GenericRulebookContext)

	p.RulebookDeclarationContext = NewEmptyRulebookDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RulebookDeclarationContext))

	return p
}

func (s *GenericRulebookContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericRulebookContext) Designator() IDesignatorContext {
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

func (s *GenericRulebookContext) EOL() antlr.TerminalNode {
	return s.GetToken(InformerParserEOL, 0)
}

func (s *GenericRulebookContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterGenericRulebook(s)
	}
}

func (s *GenericRulebookContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitGenericRulebook(s)
	}
}

type SpecificRulebookContext struct {
	*RulebookDeclarationContext
}

func NewSpecificRulebookContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SpecificRulebookContext {
	var p = new(SpecificRulebookContext)

	p.RulebookDeclarationContext = NewEmptyRulebookDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RulebookDeclarationContext))

	return p
}

func (s *SpecificRulebookContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificRulebookContext) Designator() IDesignatorContext {
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

func (s *SpecificRulebookContext) WORD() antlr.TerminalNode {
	return s.GetToken(InformerParserWORD, 0)
}

func (s *SpecificRulebookContext) EOL() antlr.TerminalNode {
	return s.GetToken(InformerParserEOL, 0)
}

func (s *SpecificRulebookContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterSpecificRulebook(s)
	}
}

func (s *SpecificRulebookContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitSpecificRulebook(s)
	}
}

func (p *InformerParser) RulebookDeclaration() (localctx IRulebookDeclarationContext) {
	this := p
	_ = this

	localctx = NewRulebookDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, InformerParserRULE_rulebookDeclaration)
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

	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewGenericRulebookContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.designator(0)
		}
		{
			p.SetState(34)
			p.Match(InformerParserT__0)
		}
		{
			p.SetState(35)
			p.Match(InformerParserT__1)
		}
		{
			p.SetState(36)
			p.Match(InformerParserT__2)
		}
		{
			p.SetState(37)
			p.Match(InformerParserT__3)
		}
		{
			p.SetState(38)
			p.Match(InformerParserEOL)
		}

	case 2:
		localctx = NewSpecificRulebookContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.designator(0)
		}
		{
			p.SetState(41)
			p.Match(InformerParserT__0)
		}
		{
			p.SetState(42)
			_la = p.GetTokenStream().LA(1)

			if !(_la == InformerParserT__1 || _la == InformerParserT__4) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(43)
			p.Match(InformerParserWORD)
		}
		{
			p.SetState(44)
			p.Match(InformerParserT__5)
		}
		{
			p.SetState(45)
			p.Match(InformerParserT__2)
		}
		{
			p.SetState(46)
			p.Match(InformerParserT__3)
		}
		{
			p.SetState(47)
			p.Match(InformerParserEOL)
		}

	}

	return localctx
}

// IActivityDeclarationContext is an interface to support dynamic dispatch.
type IActivityDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActivityDeclarationContext differentiates from other interfaces.
	IsActivityDeclarationContext()
}

type ActivityDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActivityDeclarationContext() *ActivityDeclarationContext {
	var p = new(ActivityDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_activityDeclaration
	return p
}

func (*ActivityDeclarationContext) IsActivityDeclarationContext() {}

func NewActivityDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActivityDeclarationContext {
	var p = new(ActivityDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_activityDeclaration

	return p
}

func (s *ActivityDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ActivityDeclarationContext) Designator() IDesignatorContext {
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

func (s *ActivityDeclarationContext) EOL() antlr.TerminalNode {
	return s.GetToken(InformerParserEOL, 0)
}

func (s *ActivityDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActivityDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActivityDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterActivityDeclaration(s)
	}
}

func (s *ActivityDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitActivityDeclaration(s)
	}
}

func (p *InformerParser) ActivityDeclaration() (localctx IActivityDeclarationContext) {
	this := p
	_ = this

	localctx = NewActivityDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, InformerParserRULE_activityDeclaration)

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
		p.SetState(51)
		p.designator(0)
	}
	{
		p.SetState(52)
		p.Match(InformerParserT__0)
	}
	{
		p.SetState(53)
		p.Match(InformerParserT__4)
	}
	{
		p.SetState(54)
		p.Match(InformerParserT__6)
	}
	{
		p.SetState(55)
		p.Match(InformerParserT__3)
	}
	{
		p.SetState(56)
		p.Match(InformerParserEOL)
	}

	return localctx
}

// IActionDeclarationContext is an interface to support dynamic dispatch.
type IActionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionDeclarationContext differentiates from other interfaces.
	IsActionDeclarationContext()
}

type ActionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionDeclarationContext() *ActionDeclarationContext {
	var p = new(ActionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_actionDeclaration
	return p
}

func (*ActionDeclarationContext) IsActionDeclarationContext() {}

func NewActionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionDeclarationContext {
	var p = new(ActionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_actionDeclaration

	return p
}

func (s *ActionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionDeclarationContext) Designator() IDesignatorContext {
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

func (s *ActionDeclarationContext) WORD() antlr.TerminalNode {
	return s.GetToken(InformerParserWORD, 0)
}

func (s *ActionDeclarationContext) EOL() antlr.TerminalNode {
	return s.GetToken(InformerParserEOL, 0)
}

func (s *ActionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterActionDeclaration(s)
	}
}

func (s *ActionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitActionDeclaration(s)
	}
}

func (p *InformerParser) ActionDeclaration() (localctx IActionDeclarationContext) {
	this := p
	_ = this

	localctx = NewActionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, InformerParserRULE_actionDeclaration)

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
		p.SetState(58)
		p.designator(0)
	}
	{
		p.SetState(59)
		p.Match(InformerParserT__0)
	}
	{
		p.SetState(60)
		p.Match(InformerParserT__4)
	}
	{
		p.SetState(61)
		p.Match(InformerParserT__7)
	}
	{
		p.SetState(62)
		p.Match(InformerParserT__8)
	}
	{
		p.SetState(63)
		p.Match(InformerParserT__9)
	}
	{
		p.SetState(64)
		p.Match(InformerParserT__1)
	}
	{
		p.SetState(65)
		p.Match(InformerParserWORD)
	}
	{
		p.SetState(66)
		p.Match(InformerParserT__3)
	}
	{
		p.SetState(67)
		p.Match(InformerParserEOL)
	}

	return localctx
}

// IAliasDeclarationContext is an interface to support dynamic dispatch.
type IAliasDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasDeclarationContext differentiates from other interfaces.
	IsAliasDeclarationContext()
}

type AliasDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasDeclarationContext() *AliasDeclarationContext {
	var p = new(AliasDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_aliasDeclaration
	return p
}

func (*AliasDeclarationContext) IsAliasDeclarationContext() {}

func NewAliasDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasDeclarationContext {
	var p = new(AliasDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_aliasDeclaration

	return p
}

func (s *AliasDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasDeclarationContext) STRING() antlr.TerminalNode {
	return s.GetToken(InformerParserSTRING, 0)
}

func (s *AliasDeclarationContext) WORD() antlr.TerminalNode {
	return s.GetToken(InformerParserWORD, 0)
}

func (s *AliasDeclarationContext) EOL() antlr.TerminalNode {
	return s.GetToken(InformerParserEOL, 0)
}

func (s *AliasDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterAliasDeclaration(s)
	}
}

func (s *AliasDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitAliasDeclaration(s)
	}
}

func (p *InformerParser) AliasDeclaration() (localctx IAliasDeclarationContext) {
	this := p
	_ = this

	localctx = NewAliasDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, InformerParserRULE_aliasDeclaration)

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
		p.SetState(69)
		p.Match(InformerParserT__10)
	}
	{
		p.SetState(70)
		p.Match(InformerParserSTRING)
	}
	{
		p.SetState(71)
		p.Match(InformerParserT__11)
	}
	{
		p.SetState(72)
		p.Match(InformerParserWORD)
	}
	{
		p.SetState(73)
		p.Match(InformerParserT__3)
	}
	{
		p.SetState(74)
		p.Match(InformerParserEOL)
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

func (s *DesignatorContext) WORD() antlr.TerminalNode {
	return s.GetToken(InformerParserWORD, 0)
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
	return p.designator(0)
}

func (p *InformerParser) designator(_p int) (localctx IDesignatorContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDesignatorContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDesignatorContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, InformerParserRULE_designator, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(InformerParserWORD)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDesignatorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, InformerParserRULE_designator)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(80)
					p.Match(InformerParserWORD)
				}

			case 2:
				localctx = NewDesignatorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, InformerParserRULE_designator)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(82)
					p.Match(InformerParserT__12)
				}
				{
					p.SetState(83)
					p.Match(InformerParserWORD)
				}

			}

		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

func (p *InformerParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *DesignatorContext = nil
		if localctx != nil {
			t = localctx.(*DesignatorContext)
		}
		return p.Designator_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *InformerParser) Designator_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
