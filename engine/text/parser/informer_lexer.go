// Code generated from Informer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type InformerLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var informerlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func informerlexerLexerInit() {
	staticData := &informerlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'is'", "'a'", "'rulebook'", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "NUMBER", "IDENT", "PUNCT", "STRING", "COMMENT",
		"WHITESPACE",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "NUMBER", "IDENT", "PUNCT", "STRING",
		"COMMENT", "WHITESPACE", "ESC",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 85, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 4, 4, 41, 8, 4, 11, 4, 12, 4,
		42, 1, 5, 4, 5, 46, 8, 5, 11, 5, 12, 5, 47, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 5, 7, 55, 8, 7, 10, 7, 12, 7, 58, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 5, 8,
		64, 8, 8, 10, 8, 12, 8, 67, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 4, 9, 74,
		8, 9, 11, 9, 12, 9, 75, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10,
		84, 8, 10, 2, 56, 65, 0, 11, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 0, 1, 0, 4, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122,
		4, 0, 44, 44, 46, 46, 58, 58, 92, 92, 3, 0, 9, 10, 13, 13, 32, 32, 90,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 1, 23, 1, 0, 0, 0, 3, 26, 1, 0,
		0, 0, 5, 28, 1, 0, 0, 0, 7, 37, 1, 0, 0, 0, 9, 40, 1, 0, 0, 0, 11, 45,
		1, 0, 0, 0, 13, 49, 1, 0, 0, 0, 15, 51, 1, 0, 0, 0, 17, 61, 1, 0, 0, 0,
		19, 73, 1, 0, 0, 0, 21, 83, 1, 0, 0, 0, 23, 24, 5, 105, 0, 0, 24, 25, 5,
		115, 0, 0, 25, 2, 1, 0, 0, 0, 26, 27, 5, 97, 0, 0, 27, 4, 1, 0, 0, 0, 28,
		29, 5, 114, 0, 0, 29, 30, 5, 117, 0, 0, 30, 31, 5, 108, 0, 0, 31, 32, 5,
		101, 0, 0, 32, 33, 5, 98, 0, 0, 33, 34, 5, 111, 0, 0, 34, 35, 5, 111, 0,
		0, 35, 36, 5, 107, 0, 0, 36, 6, 1, 0, 0, 0, 37, 38, 5, 46, 0, 0, 38, 8,
		1, 0, 0, 0, 39, 41, 7, 0, 0, 0, 40, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0,
		42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 10, 1, 0, 0, 0, 44, 46, 7,
		1, 0, 0, 45, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47,
		48, 1, 0, 0, 0, 48, 12, 1, 0, 0, 0, 49, 50, 7, 2, 0, 0, 50, 14, 1, 0, 0,
		0, 51, 56, 5, 34, 0, 0, 52, 55, 3, 21, 10, 0, 53, 55, 9, 0, 0, 0, 54, 52,
		1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0,
		56, 54, 1, 0, 0, 0, 57, 59, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59, 60, 5,
		34, 0, 0, 60, 16, 1, 0, 0, 0, 61, 65, 5, 91, 0, 0, 62, 64, 9, 0, 0, 0,
		63, 62, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 65, 63, 1,
		0, 0, 0, 66, 68, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 69, 5, 93, 0, 0, 69,
		70, 1, 0, 0, 0, 70, 71, 6, 8, 0, 0, 71, 18, 1, 0, 0, 0, 72, 74, 7, 3, 0,
		0, 73, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76,
		1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78, 6, 9, 0, 0, 78, 20, 1, 0, 0, 0,
		79, 80, 5, 92, 0, 0, 80, 84, 5, 34, 0, 0, 81, 82, 5, 92, 0, 0, 82, 84,
		5, 92, 0, 0, 83, 79, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 22, 1, 0, 0, 0,
		8, 0, 42, 47, 54, 56, 65, 75, 83, 1, 6, 0, 0,
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

// InformerLexerInit initializes any static state used to implement InformerLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewInformerLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func InformerLexerInit() {
	staticData := &informerlexerLexerStaticData
	staticData.once.Do(informerlexerLexerInit)
}

// NewInformerLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewInformerLexer(input antlr.CharStream) *InformerLexer {
	InformerLexerInit()
	l := new(InformerLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &informerlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Informer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// InformerLexer tokens.
const (
	InformerLexerT__0       = 1
	InformerLexerT__1       = 2
	InformerLexerT__2       = 3
	InformerLexerT__3       = 4
	InformerLexerNUMBER     = 5
	InformerLexerIDENT      = 6
	InformerLexerPUNCT      = 7
	InformerLexerSTRING     = 8
	InformerLexerCOMMENT    = 9
	InformerLexerWHITESPACE = 10
)
