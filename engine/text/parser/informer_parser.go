// Code generated from /Users/alex/Projects/ego/engine/text/Informer.g4 by ANTLR 4.10.1. DO NOT EDIT.

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
		"", "'.'", "'is'", "'rulebook'", "'based'", "'activity'", "'kind'",
		"'of'", "'value'", "'with'", "'values'", "'usually'", "'always'", "'never'",
		"','", "'and'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ARTICLE",
		"GERUND", "WORD", "PUNCT", "COMMENT", "WHITESPACE", "EOL",
	}
	staticData.ruleNames = []string{
		"start", "statement", "definition", "definitionType", "certainty", "designator",
		"values",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 22, 88, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 3, 0, 17, 8, 0, 1, 0, 1, 0, 4, 0,
		21, 8, 0, 11, 0, 12, 0, 22, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 56, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 65, 8, 5,
		1, 5, 1, 5, 5, 5, 69, 8, 5, 10, 5, 12, 5, 72, 9, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 83, 8, 6, 10, 6, 12, 6, 86, 9, 6,
		1, 6, 0, 2, 10, 12, 7, 0, 2, 4, 6, 8, 10, 12, 0, 1, 1, 0, 11, 13, 93, 0,
		20, 1, 0, 0, 0, 2, 26, 1, 0, 0, 0, 4, 28, 1, 0, 0, 0, 6, 55, 1, 0, 0, 0,
		8, 57, 1, 0, 0, 0, 10, 64, 1, 0, 0, 0, 12, 73, 1, 0, 0, 0, 14, 16, 3, 2,
		1, 0, 15, 17, 5, 1, 0, 0, 16, 15, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 18,
		1, 0, 0, 0, 18, 19, 5, 22, 0, 0, 19, 21, 1, 0, 0, 0, 20, 14, 1, 0, 0, 0,
		21, 22, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 24, 1,
		0, 0, 0, 24, 25, 5, 0, 0, 1, 25, 1, 1, 0, 0, 0, 26, 27, 3, 4, 2, 0, 27,
		3, 1, 0, 0, 0, 28, 29, 3, 10, 5, 0, 29, 30, 5, 2, 0, 0, 30, 31, 3, 6, 3,
		0, 31, 5, 1, 0, 0, 0, 32, 33, 5, 16, 0, 0, 33, 56, 5, 3, 0, 0, 34, 35,
		5, 16, 0, 0, 35, 36, 3, 10, 5, 0, 36, 37, 5, 4, 0, 0, 37, 38, 5, 3, 0,
		0, 38, 56, 1, 0, 0, 0, 39, 40, 5, 16, 0, 0, 40, 56, 5, 5, 0, 0, 41, 42,
		3, 8, 4, 0, 42, 43, 3, 10, 5, 0, 43, 56, 1, 0, 0, 0, 44, 45, 5, 16, 0,
		0, 45, 46, 5, 6, 0, 0, 46, 47, 5, 7, 0, 0, 47, 56, 3, 10, 5, 0, 48, 49,
		5, 16, 0, 0, 49, 50, 5, 6, 0, 0, 50, 51, 5, 7, 0, 0, 51, 52, 5, 8, 0, 0,
		52, 53, 5, 9, 0, 0, 53, 54, 5, 10, 0, 0, 54, 56, 3, 12, 6, 0, 55, 32, 1,
		0, 0, 0, 55, 34, 1, 0, 0, 0, 55, 39, 1, 0, 0, 0, 55, 41, 1, 0, 0, 0, 55,
		44, 1, 0, 0, 0, 55, 48, 1, 0, 0, 0, 56, 7, 1, 0, 0, 0, 57, 58, 7, 0, 0,
		0, 58, 9, 1, 0, 0, 0, 59, 60, 6, 5, -1, 0, 60, 65, 5, 18, 0, 0, 61, 65,
		5, 16, 0, 0, 62, 65, 5, 17, 0, 0, 63, 65, 5, 7, 0, 0, 64, 59, 1, 0, 0,
		0, 64, 61, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 63, 1, 0, 0, 0, 65, 70,
		1, 0, 0, 0, 66, 67, 10, 1, 0, 0, 67, 69, 3, 10, 5, 2, 68, 66, 1, 0, 0,
		0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 11,
		1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 74, 6, 6, -1, 0, 74, 75, 5, 18, 0,
		0, 75, 84, 1, 0, 0, 0, 76, 77, 10, 2, 0, 0, 77, 78, 5, 14, 0, 0, 78, 83,
		3, 12, 6, 3, 79, 80, 10, 1, 0, 0, 80, 81, 5, 15, 0, 0, 81, 83, 3, 12, 6,
		2, 82, 76, 1, 0, 0, 0, 82, 79, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82,
		1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 13, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0,
		7, 16, 22, 55, 64, 70, 82, 84,
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
	InformerParserT__13      = 14
	InformerParserT__14      = 15
	InformerParserARTICLE    = 16
	InformerParserGERUND     = 17
	InformerParserWORD       = 18
	InformerParserPUNCT      = 19
	InformerParserCOMMENT    = 20
	InformerParserWHITESPACE = 21
	InformerParserEOL        = 22
)

// InformerParser rules.
const (
	InformerParserRULE_start          = 0
	InformerParserRULE_statement      = 1
	InformerParserRULE_definition     = 2
	InformerParserRULE_definitionType = 3
	InformerParserRULE_certainty      = 4
	InformerParserRULE_designator     = 5
	InformerParserRULE_values         = 6
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

func (s *StartContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *StartContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(InformerParserEOL)
}

func (s *StartContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(InformerParserEOL, i)
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
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<InformerParserT__6)|(1<<InformerParserARTICLE)|(1<<InformerParserGERUND)|(1<<InformerParserWORD))) != 0) {
		{
			p.SetState(14)
			p.Statement()
		}
		p.SetState(16)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == InformerParserT__0 {
			{
				p.SetState(15)
				p.Match(InformerParserT__0)
			}

		}
		{
			p.SetState(18)
			p.Match(InformerParserEOL)
		}

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(24)
		p.Match(InformerParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Definition() IDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *InformerParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InformerParserRULE_statement)

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
		p.SetState(26)
		p.Definition()
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

func (s *DefinitionContext) Designator() IDesignatorContext {
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

func (s *DefinitionContext) DefinitionType() IDefinitionTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionTypeContext)
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
	p.EnterRule(localctx, 4, InformerParserRULE_definition)

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
		p.SetState(28)
		p.designator(0)
	}
	{
		p.SetState(29)
		p.Match(InformerParserT__1)
	}
	{
		p.SetState(30)
		p.DefinitionType()
	}

	return localctx
}

// IDefinitionTypeContext is an interface to support dynamic dispatch.
type IDefinitionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionTypeContext differentiates from other interfaces.
	IsDefinitionTypeContext()
}

type DefinitionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionTypeContext() *DefinitionTypeContext {
	var p = new(DefinitionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_definitionType
	return p
}

func (*DefinitionTypeContext) IsDefinitionTypeContext() {}

func NewDefinitionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionTypeContext {
	var p = new(DefinitionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_definitionType

	return p
}

func (s *DefinitionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionTypeContext) CopyFrom(ctx *DefinitionTypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DefinitionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RulebookContext struct {
	*DefinitionTypeContext
}

func NewRulebookContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RulebookContext {
	var p = new(RulebookContext)

	p.DefinitionTypeContext = NewEmptyDefinitionTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DefinitionTypeContext))

	return p
}

func (s *RulebookContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulebookContext) ARTICLE() antlr.TerminalNode {
	return s.GetToken(InformerParserARTICLE, 0)
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

type CertaintyOfAttributeContext struct {
	*DefinitionTypeContext
}

func NewCertaintyOfAttributeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CertaintyOfAttributeContext {
	var p = new(CertaintyOfAttributeContext)

	p.DefinitionTypeContext = NewEmptyDefinitionTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DefinitionTypeContext))

	return p
}

func (s *CertaintyOfAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CertaintyOfAttributeContext) Certainty() ICertaintyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICertaintyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICertaintyContext)
}

func (s *CertaintyOfAttributeContext) Designator() IDesignatorContext {
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

func (s *CertaintyOfAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterCertaintyOfAttribute(s)
	}
}

func (s *CertaintyOfAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitCertaintyOfAttribute(s)
	}
}

type ObjectKindContext struct {
	*DefinitionTypeContext
}

func NewObjectKindContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectKindContext {
	var p = new(ObjectKindContext)

	p.DefinitionTypeContext = NewEmptyDefinitionTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DefinitionTypeContext))

	return p
}

func (s *ObjectKindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectKindContext) ARTICLE() antlr.TerminalNode {
	return s.GetToken(InformerParserARTICLE, 0)
}

func (s *ObjectKindContext) Designator() IDesignatorContext {
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

func (s *ObjectKindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterObjectKind(s)
	}
}

func (s *ObjectKindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitObjectKind(s)
	}
}

type ActivityContext struct {
	*DefinitionTypeContext
}

func NewActivityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ActivityContext {
	var p = new(ActivityContext)

	p.DefinitionTypeContext = NewEmptyDefinitionTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DefinitionTypeContext))

	return p
}

func (s *ActivityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActivityContext) ARTICLE() antlr.TerminalNode {
	return s.GetToken(InformerParserARTICLE, 0)
}

func (s *ActivityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterActivity(s)
	}
}

func (s *ActivityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitActivity(s)
	}
}

type ObjectBasedRulebookContext struct {
	*DefinitionTypeContext
}

func NewObjectBasedRulebookContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectBasedRulebookContext {
	var p = new(ObjectBasedRulebookContext)

	p.DefinitionTypeContext = NewEmptyDefinitionTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DefinitionTypeContext))

	return p
}

func (s *ObjectBasedRulebookContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectBasedRulebookContext) ARTICLE() antlr.TerminalNode {
	return s.GetToken(InformerParserARTICLE, 0)
}

func (s *ObjectBasedRulebookContext) Designator() IDesignatorContext {
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

func (s *ObjectBasedRulebookContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterObjectBasedRulebook(s)
	}
}

func (s *ObjectBasedRulebookContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitObjectBasedRulebook(s)
	}
}

type ValueKindContext struct {
	*DefinitionTypeContext
}

func NewValueKindContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueKindContext {
	var p = new(ValueKindContext)

	p.DefinitionTypeContext = NewEmptyDefinitionTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DefinitionTypeContext))

	return p
}

func (s *ValueKindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueKindContext) ARTICLE() antlr.TerminalNode {
	return s.GetToken(InformerParserARTICLE, 0)
}

func (s *ValueKindContext) Values() IValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *ValueKindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterValueKind(s)
	}
}

func (s *ValueKindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitValueKind(s)
	}
}

func (p *InformerParser) DefinitionType() (localctx IDefinitionTypeContext) {
	this := p
	_ = this

	localctx = NewDefinitionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, InformerParserRULE_definitionType)

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

	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRulebookContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.Match(InformerParserARTICLE)
		}
		{
			p.SetState(33)
			p.Match(InformerParserT__2)
		}

	case 2:
		localctx = NewObjectBasedRulebookContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Match(InformerParserARTICLE)
		}
		{
			p.SetState(35)
			p.designator(0)
		}
		{
			p.SetState(36)
			p.Match(InformerParserT__3)
		}
		{
			p.SetState(37)
			p.Match(InformerParserT__2)
		}

	case 3:
		localctx = NewActivityContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.Match(InformerParserARTICLE)
		}
		{
			p.SetState(40)
			p.Match(InformerParserT__4)
		}

	case 4:
		localctx = NewCertaintyOfAttributeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(41)
			p.Certainty()
		}
		{
			p.SetState(42)
			p.designator(0)
		}

	case 5:
		localctx = NewObjectKindContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(44)
			p.Match(InformerParserARTICLE)
		}
		{
			p.SetState(45)
			p.Match(InformerParserT__5)
		}
		{
			p.SetState(46)
			p.Match(InformerParserT__6)
		}
		{
			p.SetState(47)
			p.designator(0)
		}

	case 6:
		localctx = NewValueKindContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(48)
			p.Match(InformerParserARTICLE)
		}
		{
			p.SetState(49)
			p.Match(InformerParserT__5)
		}
		{
			p.SetState(50)
			p.Match(InformerParserT__6)
		}
		{
			p.SetState(51)
			p.Match(InformerParserT__7)
		}
		{
			p.SetState(52)
			p.Match(InformerParserT__8)
		}
		{
			p.SetState(53)
			p.Match(InformerParserT__9)
		}
		{
			p.SetState(54)
			p.values(0)
		}

	}

	return localctx
}

// ICertaintyContext is an interface to support dynamic dispatch.
type ICertaintyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCertaintyContext differentiates from other interfaces.
	IsCertaintyContext()
}

type CertaintyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCertaintyContext() *CertaintyContext {
	var p = new(CertaintyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_certainty
	return p
}

func (*CertaintyContext) IsCertaintyContext() {}

func NewCertaintyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CertaintyContext {
	var p = new(CertaintyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_certainty

	return p
}

func (s *CertaintyContext) GetParser() antlr.Parser { return s.parser }
func (s *CertaintyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CertaintyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CertaintyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterCertainty(s)
	}
}

func (s *CertaintyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitCertainty(s)
	}
}

func (p *InformerParser) Certainty() (localctx ICertaintyContext) {
	this := p
	_ = this

	localctx = NewCertaintyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, InformerParserRULE_certainty)
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
	{
		p.SetState(57)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<InformerParserT__10)|(1<<InformerParserT__11)|(1<<InformerParserT__12))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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

func (s *DesignatorContext) ARTICLE() antlr.TerminalNode {
	return s.GetToken(InformerParserARTICLE, 0)
}

func (s *DesignatorContext) GERUND() antlr.TerminalNode {
	return s.GetToken(InformerParserGERUND, 0)
}

func (s *DesignatorContext) AllDesignator() []IDesignatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDesignatorContext); ok {
			len++
		}
	}

	tst := make([]IDesignatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDesignatorContext); ok {
			tst[i] = t.(IDesignatorContext)
			i++
		}
	}

	return tst
}

func (s *DesignatorContext) Designator(i int) IDesignatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDesignatorContext); ok {
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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, InformerParserRULE_designator, _p)

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
	p.SetState(64)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case InformerParserWORD:
		{
			p.SetState(60)
			p.Match(InformerParserWORD)
		}

	case InformerParserARTICLE:
		{
			p.SetState(61)
			p.Match(InformerParserARTICLE)
		}

	case InformerParserGERUND:
		{
			p.SetState(62)
			p.Match(InformerParserGERUND)
		}

	case InformerParserT__6:
		{
			p.SetState(63)
			p.Match(InformerParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDesignatorContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, InformerParserRULE_designator)
			p.SetState(66)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(67)
				p.designator(2)
			}

		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IValuesContext is an interface to support dynamic dispatch.
type IValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValuesContext differentiates from other interfaces.
	IsValuesContext()
}

type ValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuesContext() *ValuesContext {
	var p = new(ValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_values
	return p
}

func (*ValuesContext) IsValuesContext() {}

func NewValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuesContext {
	var p = new(ValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_values

	return p
}

func (s *ValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuesContext) WORD() antlr.TerminalNode {
	return s.GetToken(InformerParserWORD, 0)
}

func (s *ValuesContext) AllValues() []IValuesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValuesContext); ok {
			len++
		}
	}

	tst := make([]IValuesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValuesContext); ok {
			tst[i] = t.(IValuesContext)
			i++
		}
	}

	return tst
}

func (s *ValuesContext) Values(i int) IValuesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuesContext); ok {
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

	return t.(IValuesContext)
}

func (s *ValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterValues(s)
	}
}

func (s *ValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitValues(s)
	}
}

func (p *InformerParser) Values() (localctx IValuesContext) {
	return p.values(0)
}

func (p *InformerParser) values(_p int) (localctx IValuesContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewValuesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValuesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, InformerParserRULE_values, _p)

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
		p.SetState(74)
		p.Match(InformerParserWORD)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewValuesContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, InformerParserRULE_values)
				p.SetState(76)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(77)
					p.Match(InformerParserT__13)
				}
				{
					p.SetState(78)
					p.values(3)
				}

			case 2:
				localctx = NewValuesContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, InformerParserRULE_values)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(80)
					p.Match(InformerParserT__14)
				}
				{
					p.SetState(81)
					p.values(2)
				}

			}

		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

func (p *InformerParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *DesignatorContext = nil
		if localctx != nil {
			t = localctx.(*DesignatorContext)
		}
		return p.Designator_Sempred(t, predIndex)

	case 6:
		var t *ValuesContext = nil
		if localctx != nil {
			t = localctx.(*ValuesContext)
		}
		return p.Values_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *InformerParser) Designator_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *InformerParser) Values_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
