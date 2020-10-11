// Code generated from Simpl.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 100,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 17, 6, 17, 85, 10, 17, 13, 17, 14, 17, 86, 3, 18,
	6, 18, 90, 10, 18, 13, 18, 14, 18, 91, 3, 19, 6, 19, 95, 10, 19, 13, 19,
	14, 19, 96, 3, 19, 3, 19, 2, 2, 20, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 3, 2, 5, 3, 2, 99, 124, 3, 2, 50, 59, 5, 2, 11,
	12, 15, 15, 34, 34, 2, 102, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2,
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2,
	2, 3, 39, 3, 2, 2, 2, 5, 41, 3, 2, 2, 2, 7, 43, 3, 2, 2, 2, 9, 46, 3, 2,
	2, 2, 11, 48, 3, 2, 2, 2, 13, 50, 3, 2, 2, 2, 15, 52, 3, 2, 2, 2, 17, 54,
	3, 2, 2, 2, 19, 57, 3, 2, 2, 2, 21, 60, 3, 2, 2, 2, 23, 65, 3, 2, 2, 2,
	25, 70, 3, 2, 2, 2, 27, 74, 3, 2, 2, 2, 29, 76, 3, 2, 2, 2, 31, 79, 3,
	2, 2, 2, 33, 84, 3, 2, 2, 2, 35, 89, 3, 2, 2, 2, 37, 94, 3, 2, 2, 2, 39,
	40, 7, 42, 2, 2, 40, 4, 3, 2, 2, 2, 41, 42, 7, 43, 2, 2, 42, 6, 3, 2, 2,
	2, 43, 44, 7, 63, 2, 2, 44, 45, 7, 63, 2, 2, 45, 8, 3, 2, 2, 2, 46, 47,
	7, 49, 2, 2, 47, 10, 3, 2, 2, 2, 48, 49, 7, 44, 2, 2, 49, 12, 3, 2, 2,
	2, 50, 51, 7, 47, 2, 2, 51, 14, 3, 2, 2, 2, 52, 53, 7, 45, 2, 2, 53, 16,
	3, 2, 2, 2, 54, 55, 7, 47, 2, 2, 55, 56, 7, 64, 2, 2, 56, 18, 3, 2, 2,
	2, 57, 58, 7, 107, 2, 2, 58, 59, 7, 104, 2, 2, 59, 20, 3, 2, 2, 2, 60,
	61, 7, 118, 2, 2, 61, 62, 7, 106, 2, 2, 62, 63, 7, 103, 2, 2, 63, 64, 7,
	112, 2, 2, 64, 22, 3, 2, 2, 2, 65, 66, 7, 103, 2, 2, 66, 67, 7, 110, 2,
	2, 67, 68, 7, 117, 2, 2, 68, 69, 7, 103, 2, 2, 69, 24, 3, 2, 2, 2, 70,
	71, 7, 110, 2, 2, 71, 72, 7, 103, 2, 2, 72, 73, 7, 118, 2, 2, 73, 26, 3,
	2, 2, 2, 74, 75, 7, 63, 2, 2, 75, 28, 3, 2, 2, 2, 76, 77, 7, 107, 2, 2,
	77, 78, 7, 112, 2, 2, 78, 30, 3, 2, 2, 2, 79, 80, 7, 103, 2, 2, 80, 81,
	7, 112, 2, 2, 81, 82, 7, 102, 2, 2, 82, 32, 3, 2, 2, 2, 83, 85, 9, 2, 2,
	2, 84, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87,
	3, 2, 2, 2, 87, 34, 3, 2, 2, 2, 88, 90, 9, 3, 2, 2, 89, 88, 3, 2, 2, 2,
	90, 91, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 36, 3,
	2, 2, 2, 93, 95, 9, 4, 2, 2, 94, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96,
	94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 8, 19,
	2, 2, 99, 38, 3, 2, 2, 2, 6, 2, 86, 91, 96, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'=='", "'/'", "'*'", "'-'", "'+'", "'->'", "'if'", "'then'",
	"'else'", "'let'", "'='", "'in'", "'end'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ID", "NUM",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "ID", "NUM", "WS",
}

type SimplLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSimplLexer(input antlr.CharStream) *SimplLexer {

	l := new(SimplLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Simpl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SimplLexer tokens.
const (
	SimplLexerT__0  = 1
	SimplLexerT__1  = 2
	SimplLexerT__2  = 3
	SimplLexerT__3  = 4
	SimplLexerT__4  = 5
	SimplLexerT__5  = 6
	SimplLexerT__6  = 7
	SimplLexerT__7  = 8
	SimplLexerT__8  = 9
	SimplLexerT__9  = 10
	SimplLexerT__10 = 11
	SimplLexerT__11 = 12
	SimplLexerT__12 = 13
	SimplLexerT__13 = 14
	SimplLexerT__14 = 15
	SimplLexerID    = 16
	SimplLexerNUM   = 17
	SimplLexerWS    = 18
)
