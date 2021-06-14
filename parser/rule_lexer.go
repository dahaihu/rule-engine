// Code generated from /home/studing/GolandProjects/src/rule-engine/parser/Rule.g4 by ANTLR 4.9.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 130, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 3, 2, 3, 2, 3, 3, 6, 3, 53, 10, 3, 13, 3, 14, 3, 54, 
	3, 3, 3, 3, 3, 4, 6, 4, 60, 10, 4, 13, 4, 14, 4, 61, 3, 5, 3, 5, 3, 5, 
	6, 5, 67, 10, 5, 13, 5, 14, 5, 68, 3, 6, 3, 6, 3, 6, 6, 6, 74, 10, 6, 13, 
	6, 14, 6, 75, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 
	9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 
	3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 
	18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 7, 21, 117, 
	10, 21, 12, 21, 14, 21, 120, 11, 21, 3, 22, 3, 22, 3, 22, 5, 22, 125, 10, 
	22, 3, 23, 3, 23, 3, 24, 3, 24, 2, 2, 25, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 
	17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 2, 45, 2, 47, 2, 3, 2, 
	7, 5, 2, 11, 12, 15, 15, 34, 34, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 5, 
	2, 50, 59, 67, 92, 99, 124, 4, 2, 47, 47, 97, 97, 2, 133, 2, 3, 3, 2, 2, 
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 
	2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 
	3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 
	49, 3, 2, 2, 2, 5, 52, 3, 2, 2, 2, 7, 59, 3, 2, 2, 2, 9, 63, 3, 2, 2, 2, 
	11, 70, 3, 2, 2, 2, 13, 79, 3, 2, 2, 2, 15, 83, 3, 2, 2, 2, 17, 86, 3, 
	2, 2, 2, 19, 88, 3, 2, 2, 2, 21, 90, 3, 2, 2, 2, 23, 92, 3, 2, 2, 2, 25, 
	94, 3, 2, 2, 2, 27, 97, 3, 2, 2, 2, 29, 100, 3, 2, 2, 2, 31, 102, 3, 2, 
	2, 2, 33, 104, 3, 2, 2, 2, 35, 107, 3, 2, 2, 2, 37, 110, 3, 2, 2, 2, 39, 
	112, 3, 2, 2, 2, 41, 114, 3, 2, 2, 2, 43, 124, 3, 2, 2, 2, 45, 126, 3, 
	2, 2, 2, 47, 128, 3, 2, 2, 2, 49, 50, 7, 48, 2, 2, 50, 4, 3, 2, 2, 2, 51, 
	53, 9, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 52, 3, 2, 2, 
	2, 54, 55, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 57, 8, 3, 2, 2, 57, 6, 3, 
	2, 2, 2, 58, 60, 9, 3, 2, 2, 59, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 
	59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 8, 3, 2, 2, 2, 63, 64, 5, 7, 4, 
	2, 64, 66, 7, 48, 2, 2, 65, 67, 9, 3, 2, 2, 66, 65, 3, 2, 2, 2, 67, 68, 
	3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 10, 3, 2, 2, 2, 
	70, 71, 7, 36, 2, 2, 71, 73, 9, 4, 2, 2, 72, 74, 9, 5, 2, 2, 73, 72, 3, 
	2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 
	77, 3, 2, 2, 2, 77, 78, 7, 36, 2, 2, 78, 12, 3, 2, 2, 2, 79, 80, 7, 99, 
	2, 2, 80, 81, 7, 112, 2, 2, 81, 82, 7, 102, 2, 2, 82, 14, 3, 2, 2, 2, 83, 
	84, 7, 113, 2, 2, 84, 85, 7, 116, 2, 2, 85, 16, 3, 2, 2, 2, 86, 87, 7, 
	44, 2, 2, 87, 18, 3, 2, 2, 2, 88, 89, 7, 49, 2, 2, 89, 20, 3, 2, 2, 2, 
	90, 91, 7, 45, 2, 2, 91, 22, 3, 2, 2, 2, 92, 93, 7, 47, 2, 2, 93, 24, 3, 
	2, 2, 2, 94, 95, 7, 63, 2, 2, 95, 96, 7, 63, 2, 2, 96, 26, 3, 2, 2, 2, 
	97, 98, 7, 35, 2, 2, 98, 99, 7, 63, 2, 2, 99, 28, 3, 2, 2, 2, 100, 101, 
	7, 64, 2, 2, 101, 30, 3, 2, 2, 2, 102, 103, 7, 62, 2, 2, 103, 32, 3, 2, 
	2, 2, 104, 105, 7, 64, 2, 2, 105, 106, 7, 63, 2, 2, 106, 34, 3, 2, 2, 2, 
	107, 108, 7, 62, 2, 2, 108, 109, 7, 63, 2, 2, 109, 36, 3, 2, 2, 2, 110, 
	111, 7, 42, 2, 2, 111, 38, 3, 2, 2, 2, 112, 113, 7, 43, 2, 2, 113, 40, 
	3, 2, 2, 2, 114, 118, 5, 47, 24, 2, 115, 117, 5, 43, 22, 2, 116, 115, 3, 
	2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 
	2, 119, 42, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 125, 9, 6, 2, 2, 122, 
	125, 5, 45, 23, 2, 123, 125, 5, 47, 24, 2, 124, 121, 3, 2, 2, 2, 124, 122, 
	3, 2, 2, 2, 124, 123, 3, 2, 2, 2, 125, 44, 3, 2, 2, 2, 126, 127, 4, 50, 
	59, 2, 127, 46, 3, 2, 2, 2, 128, 129, 9, 4, 2, 2, 129, 48, 3, 2, 2, 2, 
	9, 2, 54, 61, 68, 75, 118, 124, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'.'", "", "", "", "", "'and'", "'or'", "'*'", "'/'", "'+'", "'-'", 
	"'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "WHITESPACE", "INT", "FLOAT", "STRING", "AND", "OR", "MUL", "DIV", 
	"ADD", "SUB", "EQ", "NE", "GT", "LT", "GE", "LE", "LEFTBRACKET", "RIGHTBRACKET", 
	"ATTRNAME",
}

var lexerRuleNames = []string{
	"T__0", "WHITESPACE", "INT", "FLOAT", "STRING", "AND", "OR", "MUL", "DIV", 
	"ADD", "SUB", "EQ", "NE", "GT", "LT", "GE", "LE", "LEFTBRACKET", "RIGHTBRACKET", 
	"ATTRNAME", "ATTR_NAME_CHAR", "DIGIT", "ALPHA",
}
type RuleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

// NewRuleLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *RuleLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewRuleLexer(input antlr.CharStream) *RuleLexer {
	l := new(RuleLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Rule.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RuleLexer tokens.
const (
	RuleLexerT__0 = 1
	RuleLexerWHITESPACE = 2
	RuleLexerINT = 3
	RuleLexerFLOAT = 4
	RuleLexerSTRING = 5
	RuleLexerAND = 6
	RuleLexerOR = 7
	RuleLexerMUL = 8
	RuleLexerDIV = 9
	RuleLexerADD = 10
	RuleLexerSUB = 11
	RuleLexerEQ = 12
	RuleLexerNE = 13
	RuleLexerGT = 14
	RuleLexerLT = 15
	RuleLexerGE = 16
	RuleLexerLE = 17
	RuleLexerLEFTBRACKET = 18
	RuleLexerRIGHTBRACKET = 19
	RuleLexerATTRNAME = 20
)

