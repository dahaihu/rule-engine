// Code generated from /home/studing/GolandProjects/src/rule-engine/parser/Rule.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Rule

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 74, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3, 
	2, 3, 2, 5, 2, 17, 10, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 35, 10, 5, 3, 
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 43, 10, 5, 12, 5, 14, 5, 46, 11, 
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 56, 10, 6, 3, 
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 64, 10, 7, 3, 7, 3, 7, 3, 7, 7, 
	7, 69, 10, 7, 12, 7, 14, 7, 72, 11, 7, 3, 7, 2, 4, 8, 12, 8, 2, 4, 6, 8, 
	10, 12, 2, 6, 3, 2, 10, 11, 3, 2, 12, 13, 3, 2, 14, 19, 3, 2, 8, 9, 2, 
	77, 2, 14, 3, 2, 2, 2, 4, 18, 3, 2, 2, 2, 6, 21, 3, 2, 2, 2, 8, 34, 3, 
	2, 2, 2, 10, 55, 3, 2, 2, 2, 12, 63, 3, 2, 2, 2, 14, 16, 7, 22, 2, 2, 15, 
	17, 5, 4, 3, 2, 16, 15, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 3, 3, 2, 2, 
	2, 18, 19, 7, 3, 2, 2, 19, 20, 5, 2, 2, 2, 20, 5, 3, 2, 2, 2, 21, 22, 5, 
	12, 7, 2, 22, 23, 7, 2, 2, 3, 23, 7, 3, 2, 2, 2, 24, 25, 8, 5, 1, 2, 25, 
	35, 7, 5, 2, 2, 26, 35, 7, 6, 2, 2, 27, 35, 5, 2, 2, 2, 28, 29, 7, 13, 
	2, 2, 29, 35, 5, 8, 5, 6, 30, 31, 7, 20, 2, 2, 31, 32, 5, 8, 5, 2, 32, 
	33, 7, 21, 2, 2, 33, 35, 3, 2, 2, 2, 34, 24, 3, 2, 2, 2, 34, 26, 3, 2, 
	2, 2, 34, 27, 3, 2, 2, 2, 34, 28, 3, 2, 2, 2, 34, 30, 3, 2, 2, 2, 35, 44, 
	3, 2, 2, 2, 36, 37, 12, 5, 2, 2, 37, 38, 9, 2, 2, 2, 38, 43, 5, 8, 5, 6, 
	39, 40, 12, 4, 2, 2, 40, 41, 9, 3, 2, 2, 41, 43, 5, 8, 5, 5, 42, 36, 3, 
	2, 2, 2, 42, 39, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 
	45, 3, 2, 2, 2, 45, 9, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 47, 48, 7, 20, 2, 
	2, 48, 49, 5, 10, 6, 2, 49, 50, 7, 21, 2, 2, 50, 56, 3, 2, 2, 2, 51, 52, 
	5, 8, 5, 2, 52, 53, 9, 4, 2, 2, 53, 54, 5, 8, 5, 2, 54, 56, 3, 2, 2, 2, 
	55, 47, 3, 2, 2, 2, 55, 51, 3, 2, 2, 2, 56, 11, 3, 2, 2, 2, 57, 58, 8, 
	7, 1, 2, 58, 64, 5, 10, 6, 2, 59, 60, 7, 20, 2, 2, 60, 61, 5, 12, 7, 2, 
	61, 62, 7, 21, 2, 2, 62, 64, 3, 2, 2, 2, 63, 57, 3, 2, 2, 2, 63, 59, 3, 
	2, 2, 2, 64, 70, 3, 2, 2, 2, 65, 66, 12, 4, 2, 2, 66, 67, 9, 5, 2, 2, 67, 
	69, 5, 12, 7, 5, 68, 65, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 
	2, 2, 70, 71, 3, 2, 2, 2, 71, 13, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 9, 16, 
	34, 42, 44, 55, 63, 70,
}
var literalNames = []string{
	"", "'.'", "", "", "", "", "'and'", "'or'", "'*'", "'/'", "'+'", "'-'", 
	"'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "WHITESPACE", "INT", "FLOAT", "STRING", "AND", "OR", "MUL", "DIV", 
	"ADD", "SUB", "EQ", "NE", "GT", "LT", "GE", "LE", "LEFTBRACKET", "RIGHTBRACKET", 
	"ATTRNAME",
}

var ruleNames = []string{
	"attrPath", "subAttr", "start", "arithmetic", "compare", "logical",
}
type RuleParser struct {
	*antlr.BaseParser
}

// NewRuleParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *RuleParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewRuleParser(input antlr.TokenStream) *RuleParser {
	this := new(RuleParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Rule.g4"

	return this
}


// RuleParser tokens.
const (
	RuleParserEOF = antlr.TokenEOF
	RuleParserT__0 = 1
	RuleParserWHITESPACE = 2
	RuleParserINT = 3
	RuleParserFLOAT = 4
	RuleParserSTRING = 5
	RuleParserAND = 6
	RuleParserOR = 7
	RuleParserMUL = 8
	RuleParserDIV = 9
	RuleParserADD = 10
	RuleParserSUB = 11
	RuleParserEQ = 12
	RuleParserNE = 13
	RuleParserGT = 14
	RuleParserLT = 15
	RuleParserGE = 16
	RuleParserLE = 17
	RuleParserLEFTBRACKET = 18
	RuleParserRIGHTBRACKET = 19
	RuleParserATTRNAME = 20
)

// RuleParser rules.
const (
	RuleParserRULE_attrPath = 0
	RuleParserRULE_subAttr = 1
	RuleParserRULE_start = 2
	RuleParserRULE_arithmetic = 3
	RuleParserRULE_compare = 4
	RuleParserRULE_logical = 5
)

// IAttrPathContext is an interface to support dynamic dispatch.
type IAttrPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrPathContext differentiates from other interfaces.
	IsAttrPathContext()
}

type AttrPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrPathContext() *AttrPathContext {
	var p = new(AttrPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_attrPath
	return p
}

func (*AttrPathContext) IsAttrPathContext() {}

func NewAttrPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrPathContext {
	var p = new(AttrPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_attrPath

	return p
}

func (s *AttrPathContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrPathContext) ATTRNAME() antlr.TerminalNode {
	return s.GetToken(RuleParserATTRNAME, 0)
}

func (s *AttrPathContext) SubAttr() ISubAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubAttrContext)
}

func (s *AttrPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AttrPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterAttrPath(s)
	}
}

func (s *AttrPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitAttrPath(s)
	}
}

func (s *AttrPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitAttrPath(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RuleParser) AttrPath() (localctx IAttrPathContext) {
	localctx = NewAttrPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RuleParserRULE_attrPath)

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
		p.SetState(12)
		p.Match(RuleParserATTRNAME)
	}
	p.SetState(14)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(13)
			p.SubAttr()
		}


	}



	return localctx
}


// ISubAttrContext is an interface to support dynamic dispatch.
type ISubAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubAttrContext differentiates from other interfaces.
	IsSubAttrContext()
}

type SubAttrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubAttrContext() *SubAttrContext {
	var p = new(SubAttrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_subAttr
	return p
}

func (*SubAttrContext) IsSubAttrContext() {}

func NewSubAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubAttrContext {
	var p = new(SubAttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_subAttr

	return p
}

func (s *SubAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *SubAttrContext) AttrPath() IAttrPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *SubAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SubAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterSubAttr(s)
	}
}

func (s *SubAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitSubAttr(s)
	}
}

func (s *SubAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitSubAttr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RuleParser) SubAttr() (localctx ISubAttrContext) {
	localctx = NewSubAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RuleParserRULE_subAttr)

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
		p.SetState(16)
		p.Match(RuleParserT__0)
	}
	{
		p.SetState(17)
		p.AttrPath()
	}



	return localctx
}


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
	p.RuleIndex = RuleParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Logical() ILogicalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(RuleParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RuleParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RuleParserRULE_start)

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
		p.SetState(19)
		p.logical(0)
	}
	{
		p.SetState(20)
		p.Match(RuleParserEOF)
	}



	return localctx
}


// IArithmeticContext is an interface to support dynamic dispatch.
type IArithmeticContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArithmeticContext differentiates from other interfaces.
	IsArithmeticContext()
}

type ArithmeticContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmeticContext() *ArithmeticContext {
	var p = new(ArithmeticContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_arithmetic
	return p
}

func (*ArithmeticContext) IsArithmeticContext() {}

func NewArithmeticContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticContext {
	var p = new(ArithmeticContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_arithmetic

	return p
}

func (s *ArithmeticContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithmeticContext) CopyFrom(ctx *ArithmeticContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ArithmeticContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type FloatContext struct {
	*ArithmeticContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	p.ArithmeticContext = NewEmptyArithmeticContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArithmeticContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(RuleParserFLOAT, 0)
}


func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitFloat(s)
	}
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}


type MulDivContext struct {
	*ArithmeticContext
	left IArithmeticContext 
	op antlr.Token
	right IArithmeticContext 
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ArithmeticContext = NewEmptyArithmeticContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArithmeticContext))

	return p
}


func (s *MulDivContext) GetOp() antlr.Token { return s.op }


func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }


func (s *MulDivContext) GetLeft() IArithmeticContext { return s.left }

func (s *MulDivContext) GetRight() IArithmeticContext { return s.right }


func (s *MulDivContext) SetLeft(v IArithmeticContext) { s.left = v }

func (s *MulDivContext) SetRight(v IArithmeticContext) { s.right = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllArithmetic() []IArithmeticContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArithmeticContext)(nil)).Elem())
	var tst = make([]IArithmeticContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArithmeticContext)
		}
	}

	return tst
}

func (s *MulDivContext) Arithmetic(i int) IArithmeticContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmeticContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArithmeticContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(RuleParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(RuleParserDIV, 0)
}


func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}


type AddSubContext struct {
	*ArithmeticContext
	left IArithmeticContext 
	op antlr.Token
	right IArithmeticContext 
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ArithmeticContext = NewEmptyArithmeticContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArithmeticContext))

	return p
}


func (s *AddSubContext) GetOp() antlr.Token { return s.op }


func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }


func (s *AddSubContext) GetLeft() IArithmeticContext { return s.left }

func (s *AddSubContext) GetRight() IArithmeticContext { return s.right }


func (s *AddSubContext) SetLeft(v IArithmeticContext) { s.left = v }

func (s *AddSubContext) SetRight(v IArithmeticContext) { s.right = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllArithmetic() []IArithmeticContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArithmeticContext)(nil)).Elem())
	var tst = make([]IArithmeticContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArithmeticContext)
		}
	}

	return tst
}

func (s *AddSubContext) Arithmetic(i int) IArithmeticContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmeticContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArithmeticContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(RuleParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(RuleParserSUB, 0)
}


func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}


type NegativeValueContext struct {
	*ArithmeticContext
}

func NewNegativeValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegativeValueContext {
	var p = new(NegativeValueContext)

	p.ArithmeticContext = NewEmptyArithmeticContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArithmeticContext))

	return p
}

func (s *NegativeValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegativeValueContext) SUB() antlr.TerminalNode {
	return s.GetToken(RuleParserSUB, 0)
}

func (s *NegativeValueContext) Arithmetic() IArithmeticContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmeticContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithmeticContext)
}


func (s *NegativeValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterNegativeValue(s)
	}
}

func (s *NegativeValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitNegativeValue(s)
	}
}

func (s *NegativeValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitNegativeValue(s)

	default:
		return t.VisitChildren(s)
	}
}


type ArithmeticWithParenthesisContext struct {
	*ArithmeticContext
}

func NewArithmeticWithParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticWithParenthesisContext {
	var p = new(ArithmeticWithParenthesisContext)

	p.ArithmeticContext = NewEmptyArithmeticContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArithmeticContext))

	return p
}

func (s *ArithmeticWithParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticWithParenthesisContext) LEFTBRACKET() antlr.TerminalNode {
	return s.GetToken(RuleParserLEFTBRACKET, 0)
}

func (s *ArithmeticWithParenthesisContext) Arithmetic() IArithmeticContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmeticContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithmeticContext)
}

func (s *ArithmeticWithParenthesisContext) RIGHTBRACKET() antlr.TerminalNode {
	return s.GetToken(RuleParserRIGHTBRACKET, 0)
}


func (s *ArithmeticWithParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterArithmeticWithParenthesis(s)
	}
}

func (s *ArithmeticWithParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitArithmeticWithParenthesis(s)
	}
}

func (s *ArithmeticWithParenthesisContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitArithmeticWithParenthesis(s)

	default:
		return t.VisitChildren(s)
	}
}


type IntContext struct {
	*ArithmeticContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.ArithmeticContext = NewEmptyArithmeticContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArithmeticContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(RuleParserINT, 0)
}


func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitInt(s)
	}
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}


type AttrContext struct {
	*ArithmeticContext
}

func NewAttrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttrContext {
	var p = new(AttrContext)

	p.ArithmeticContext = NewEmptyArithmeticContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArithmeticContext))

	return p
}

func (s *AttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrContext) AttrPath() IAttrPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}


func (s *AttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterAttr(s)
	}
}

func (s *AttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitAttr(s)
	}
}

func (s *AttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitAttr(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *RuleParser) Arithmetic() (localctx IArithmeticContext) {
	return p.arithmetic(0)
}

func (p *RuleParser) arithmetic(_p int) (localctx IArithmeticContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewArithmeticContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IArithmeticContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, RuleParserRULE_arithmetic, _p)
	var _la int


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
	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserINT:
		localctx = NewIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(23)
			p.Match(RuleParserINT)
		}


	case RuleParserFLOAT:
		localctx = NewFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(24)
			p.Match(RuleParserFLOAT)
		}


	case RuleParserATTRNAME:
		localctx = NewAttrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(25)
			p.AttrPath()
		}


	case RuleParserSUB:
		localctx = NewNegativeValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(26)
			p.Match(RuleParserSUB)
		}
		{
			p.SetState(27)
			p.arithmetic(4)
		}


	case RuleParserLEFTBRACKET:
		localctx = NewArithmeticWithParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(28)
			p.Match(RuleParserLEFTBRACKET)
		}
		{
			p.SetState(29)
			p.arithmetic(0)
		}
		{
			p.SetState(30)
			p.Match(RuleParserRIGHTBRACKET)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(40)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewArithmeticContext(p, _parentctx, _parentState))
				localctx.(*MulDivContext).left = _prevctx


				p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_arithmetic)
				p.SetState(34)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(35)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RuleParserMUL || _la == RuleParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(36)

					var _x = p.arithmetic(4)

					localctx.(*MulDivContext).right = _x
				}


			case 2:
				localctx = NewAddSubContext(p, NewArithmeticContext(p, _parentctx, _parentState))
				localctx.(*AddSubContext).left = _prevctx


				p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_arithmetic)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(38)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RuleParserADD || _la == RuleParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(39)

					var _x = p.arithmetic(3)

					localctx.(*AddSubContext).right = _x
				}

			}

		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}



	return localctx
}


// ICompareContext is an interface to support dynamic dispatch.
type ICompareContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareContext differentiates from other interfaces.
	IsCompareContext()
}

type CompareContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareContext() *CompareContext {
	var p = new(CompareContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_compare
	return p
}

func (*CompareContext) IsCompareContext() {}

func NewCompareContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareContext {
	var p = new(CompareContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_compare

	return p
}

func (s *CompareContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareContext) CopyFrom(ctx *CompareContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type BasicCompareContext struct {
	*CompareContext
	left IArithmeticContext 
	op antlr.Token
	right IArithmeticContext 
}

func NewBasicCompareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BasicCompareContext {
	var p = new(BasicCompareContext)

	p.CompareContext = NewEmptyCompareContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CompareContext))

	return p
}


func (s *BasicCompareContext) GetOp() antlr.Token { return s.op }


func (s *BasicCompareContext) SetOp(v antlr.Token) { s.op = v }


func (s *BasicCompareContext) GetLeft() IArithmeticContext { return s.left }

func (s *BasicCompareContext) GetRight() IArithmeticContext { return s.right }


func (s *BasicCompareContext) SetLeft(v IArithmeticContext) { s.left = v }

func (s *BasicCompareContext) SetRight(v IArithmeticContext) { s.right = v }

func (s *BasicCompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicCompareContext) AllArithmetic() []IArithmeticContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArithmeticContext)(nil)).Elem())
	var tst = make([]IArithmeticContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArithmeticContext)
		}
	}

	return tst
}

func (s *BasicCompareContext) Arithmetic(i int) IArithmeticContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmeticContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArithmeticContext)
}

func (s *BasicCompareContext) EQ() antlr.TerminalNode {
	return s.GetToken(RuleParserEQ, 0)
}

func (s *BasicCompareContext) NE() antlr.TerminalNode {
	return s.GetToken(RuleParserNE, 0)
}

func (s *BasicCompareContext) GT() antlr.TerminalNode {
	return s.GetToken(RuleParserGT, 0)
}

func (s *BasicCompareContext) LT() antlr.TerminalNode {
	return s.GetToken(RuleParserLT, 0)
}

func (s *BasicCompareContext) GE() antlr.TerminalNode {
	return s.GetToken(RuleParserGE, 0)
}

func (s *BasicCompareContext) LE() antlr.TerminalNode {
	return s.GetToken(RuleParserLE, 0)
}


func (s *BasicCompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterBasicCompare(s)
	}
}

func (s *BasicCompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitBasicCompare(s)
	}
}

func (s *BasicCompareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitBasicCompare(s)

	default:
		return t.VisitChildren(s)
	}
}


type CompareWithParenthesisContext struct {
	*CompareContext
}

func NewCompareWithParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareWithParenthesisContext {
	var p = new(CompareWithParenthesisContext)

	p.CompareContext = NewEmptyCompareContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CompareContext))

	return p
}

func (s *CompareWithParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareWithParenthesisContext) LEFTBRACKET() antlr.TerminalNode {
	return s.GetToken(RuleParserLEFTBRACKET, 0)
}

func (s *CompareWithParenthesisContext) Compare() ICompareContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareContext)
}

func (s *CompareWithParenthesisContext) RIGHTBRACKET() antlr.TerminalNode {
	return s.GetToken(RuleParserRIGHTBRACKET, 0)
}


func (s *CompareWithParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCompareWithParenthesis(s)
	}
}

func (s *CompareWithParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCompareWithParenthesis(s)
	}
}

func (s *CompareWithParenthesisContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitCompareWithParenthesis(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *RuleParser) Compare() (localctx ICompareContext) {
	localctx = NewCompareContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RuleParserRULE_compare)
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

	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCompareWithParenthesisContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(RuleParserLEFTBRACKET)
		}
		{
			p.SetState(46)
			p.Compare()
		}
		{
			p.SetState(47)
			p.Match(RuleParserRIGHTBRACKET)
		}


	case 2:
		localctx = NewBasicCompareContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)

			var _x = p.arithmetic(0)

			localctx.(*BasicCompareContext).left = _x
		}
		{
			p.SetState(50)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*BasicCompareContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RuleParserEQ) | (1 << RuleParserNE) | (1 << RuleParserGT) | (1 << RuleParserLT) | (1 << RuleParserGE) | (1 << RuleParserLE))) != 0)) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*BasicCompareContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(51)

			var _x = p.arithmetic(0)

			localctx.(*BasicCompareContext).right = _x
		}

	}


	return localctx
}


// ILogicalContext is an interface to support dynamic dispatch.
type ILogicalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalContext differentiates from other interfaces.
	IsLogicalContext()
}

type LogicalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalContext() *LogicalContext {
	var p = new(LogicalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_logical
	return p
}

func (*LogicalContext) IsLogicalContext() {}

func NewLogicalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalContext {
	var p = new(LogicalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_logical

	return p
}

func (s *LogicalContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalContext) CopyFrom(ctx *LogicalContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type BasicLogicalContext struct {
	*LogicalContext
}

func NewBasicLogicalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BasicLogicalContext {
	var p = new(BasicLogicalContext)

	p.LogicalContext = NewEmptyLogicalContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LogicalContext))

	return p
}

func (s *BasicLogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicLogicalContext) Compare() ICompareContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareContext)
}


func (s *BasicLogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterBasicLogical(s)
	}
}

func (s *BasicLogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitBasicLogical(s)
	}
}

func (s *BasicLogicalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitBasicLogical(s)

	default:
		return t.VisitChildren(s)
	}
}


type TwoLogicalContext struct {
	*LogicalContext
	left ILogicalContext 
	op antlr.Token
	right ILogicalContext 
}

func NewTwoLogicalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TwoLogicalContext {
	var p = new(TwoLogicalContext)

	p.LogicalContext = NewEmptyLogicalContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LogicalContext))

	return p
}


func (s *TwoLogicalContext) GetOp() antlr.Token { return s.op }


func (s *TwoLogicalContext) SetOp(v antlr.Token) { s.op = v }


func (s *TwoLogicalContext) GetLeft() ILogicalContext { return s.left }

func (s *TwoLogicalContext) GetRight() ILogicalContext { return s.right }


func (s *TwoLogicalContext) SetLeft(v ILogicalContext) { s.left = v }

func (s *TwoLogicalContext) SetRight(v ILogicalContext) { s.right = v }

func (s *TwoLogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TwoLogicalContext) AllLogical() []ILogicalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILogicalContext)(nil)).Elem())
	var tst = make([]ILogicalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILogicalContext)
		}
	}

	return tst
}

func (s *TwoLogicalContext) Logical(i int) ILogicalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILogicalContext)
}

func (s *TwoLogicalContext) AND() antlr.TerminalNode {
	return s.GetToken(RuleParserAND, 0)
}

func (s *TwoLogicalContext) OR() antlr.TerminalNode {
	return s.GetToken(RuleParserOR, 0)
}


func (s *TwoLogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterTwoLogical(s)
	}
}

func (s *TwoLogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitTwoLogical(s)
	}
}

func (s *TwoLogicalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitTwoLogical(s)

	default:
		return t.VisitChildren(s)
	}
}


type LogicalWithParenthesisContext struct {
	*LogicalContext
}

func NewLogicalWithParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalWithParenthesisContext {
	var p = new(LogicalWithParenthesisContext)

	p.LogicalContext = NewEmptyLogicalContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LogicalContext))

	return p
}

func (s *LogicalWithParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalWithParenthesisContext) LEFTBRACKET() antlr.TerminalNode {
	return s.GetToken(RuleParserLEFTBRACKET, 0)
}

func (s *LogicalWithParenthesisContext) Logical() ILogicalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalContext)
}

func (s *LogicalWithParenthesisContext) RIGHTBRACKET() antlr.TerminalNode {
	return s.GetToken(RuleParserRIGHTBRACKET, 0)
}


func (s *LogicalWithParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterLogicalWithParenthesis(s)
	}
}

func (s *LogicalWithParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitLogicalWithParenthesis(s)
	}
}

func (s *LogicalWithParenthesisContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitLogicalWithParenthesis(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *RuleParser) Logical() (localctx ILogicalContext) {
	return p.logical(0)
}

func (p *RuleParser) logical(_p int) (localctx ILogicalContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLogicalContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogicalContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, RuleParserRULE_logical, _p)
	var _la int


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
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBasicLogicalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(56)
			p.Compare()
		}


	case 2:
		localctx = NewLogicalWithParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(57)
			p.Match(RuleParserLEFTBRACKET)
		}
		{
			p.SetState(58)
			p.logical(0)
		}
		{
			p.SetState(59)
			p.Match(RuleParserRIGHTBRACKET)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTwoLogicalContext(p, NewLogicalContext(p, _parentctx, _parentState))
			localctx.(*TwoLogicalContext).left = _prevctx


			p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_logical)
			p.SetState(63)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(64)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*TwoLogicalContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == RuleParserAND || _la == RuleParserOR) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*TwoLogicalContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(65)

				var _x = p.logical(3)

				localctx.(*TwoLogicalContext).right = _x
			}


		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}



	return localctx
}


func (p *RuleParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
			var t *ArithmeticContext = nil
			if localctx != nil { t = localctx.(*ArithmeticContext) }
			return p.Arithmetic_Sempred(t, predIndex)

	case 5:
			var t *LogicalContext = nil
			if localctx != nil { t = localctx.(*LogicalContext) }
			return p.Logical_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RuleParser) Arithmetic_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
			return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleParser) Logical_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
			return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

