// Code generated from Simpl.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Simpl

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 100,
	4, 2, 9, 2, 4, 3, 9, 3, 3, 2, 7, 2, 8, 10, 2, 12, 2, 14, 2, 11, 11, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 20, 10, 3, 12, 3, 14, 3,
	23, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 33, 10,
	3, 12, 3, 14, 3, 36, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 52, 10, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 7, 3, 95, 10, 3, 12, 3, 14, 3, 98, 11, 3, 3, 3, 2, 3, 4,
	4, 2, 4, 2, 2, 2, 120, 2, 9, 3, 2, 2, 2, 4, 51, 3, 2, 2, 2, 6, 8, 5, 4,
	3, 2, 7, 6, 3, 2, 2, 2, 8, 11, 3, 2, 2, 2, 9, 7, 3, 2, 2, 2, 9, 10, 3,
	2, 2, 2, 10, 3, 3, 2, 2, 2, 11, 9, 3, 2, 2, 2, 12, 13, 8, 3, 1, 2, 13,
	14, 7, 15, 2, 2, 14, 52, 5, 4, 3, 12, 15, 16, 7, 5, 2, 2, 16, 21, 5, 4,
	3, 2, 17, 18, 7, 16, 2, 2, 18, 20, 5, 4, 3, 2, 19, 17, 3, 2, 2, 2, 20,
	23, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 24, 3, 2, 2,
	2, 23, 21, 3, 2, 2, 2, 24, 25, 7, 6, 2, 2, 25, 52, 3, 2, 2, 2, 26, 27,
	7, 8, 2, 2, 27, 28, 5, 4, 3, 2, 28, 29, 7, 9, 2, 2, 29, 52, 3, 2, 2, 2,
	30, 34, 7, 17, 2, 2, 31, 33, 5, 4, 3, 2, 32, 31, 3, 2, 2, 2, 33, 36, 3,
	2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 37, 3, 2, 2, 2, 36,
	34, 3, 2, 2, 2, 37, 52, 7, 18, 2, 2, 38, 39, 7, 19, 2, 2, 39, 40, 5, 4,
	3, 2, 40, 41, 7, 20, 2, 2, 41, 42, 5, 4, 3, 2, 42, 43, 7, 21, 2, 2, 43,
	44, 5, 4, 3, 8, 44, 52, 3, 2, 2, 2, 45, 46, 7, 22, 2, 2, 46, 52, 5, 4,
	3, 7, 47, 52, 7, 23, 2, 2, 48, 52, 7, 24, 2, 2, 49, 52, 7, 25, 2, 2, 50,
	52, 7, 26, 2, 2, 51, 12, 3, 2, 2, 2, 51, 15, 3, 2, 2, 2, 51, 26, 3, 2,
	2, 2, 51, 30, 3, 2, 2, 2, 51, 38, 3, 2, 2, 2, 51, 45, 3, 2, 2, 2, 51, 47,
	3, 2, 2, 2, 51, 48, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2,
	52, 96, 3, 2, 2, 2, 53, 54, 12, 23, 2, 2, 54, 55, 7, 3, 2, 2, 55, 95, 5,
	4, 3, 24, 56, 57, 12, 22, 2, 2, 57, 58, 7, 4, 2, 2, 58, 95, 5, 4, 3, 23,
	59, 60, 12, 20, 2, 2, 60, 61, 7, 5, 2, 2, 61, 62, 5, 4, 3, 2, 62, 63, 7,
	6, 2, 2, 63, 64, 7, 7, 2, 2, 64, 65, 5, 4, 3, 21, 65, 95, 3, 2, 2, 2, 66,
	67, 12, 18, 2, 2, 67, 68, 7, 10, 2, 2, 68, 95, 5, 4, 3, 19, 69, 70, 12,
	17, 2, 2, 70, 71, 7, 11, 2, 2, 71, 95, 5, 4, 3, 18, 72, 73, 12, 16, 2,
	2, 73, 74, 7, 12, 2, 2, 74, 95, 5, 4, 3, 17, 75, 76, 12, 15, 2, 2, 76,
	77, 7, 13, 2, 2, 77, 95, 5, 4, 3, 16, 78, 79, 12, 14, 2, 2, 79, 80, 7,
	14, 2, 2, 80, 95, 5, 4, 3, 15, 81, 82, 12, 13, 2, 2, 82, 83, 7, 7, 2, 2,
	83, 95, 5, 4, 3, 14, 84, 85, 12, 21, 2, 2, 85, 86, 7, 5, 2, 2, 86, 87,
	5, 4, 3, 2, 87, 88, 7, 6, 2, 2, 88, 95, 3, 2, 2, 2, 89, 90, 12, 19, 2,
	2, 90, 91, 7, 8, 2, 2, 91, 92, 5, 4, 3, 2, 92, 93, 7, 9, 2, 2, 93, 95,
	3, 2, 2, 2, 94, 53, 3, 2, 2, 2, 94, 56, 3, 2, 2, 2, 94, 59, 3, 2, 2, 2,
	94, 66, 3, 2, 2, 2, 94, 69, 3, 2, 2, 2, 94, 72, 3, 2, 2, 2, 94, 75, 3,
	2, 2, 2, 94, 78, 3, 2, 2, 2, 94, 81, 3, 2, 2, 2, 94, 84, 3, 2, 2, 2, 94,
	89, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2,
	2, 97, 5, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 8, 9, 21, 34, 51, 94, 96,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'=='", "'.'", "'['", "']'", "'='", "'('", "')'", "'/'", "'*'", "'-'",
	"'+'", "'->'", "'//'", "','", "'{'", "'}'", "'if'", "'then'", "'else'",
	"'len'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "BOOL", "ID", "NUM", "STR", "WS",
}

var ruleNames = []string{
	"program", "expr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SimplParser struct {
	*antlr.BaseParser
}

func NewSimplParser(input antlr.TokenStream) *SimplParser {
	this := new(SimplParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Simpl.g4"

	return this
}

// SimplParser tokens.
const (
	SimplParserEOF   = antlr.TokenEOF
	SimplParserT__0  = 1
	SimplParserT__1  = 2
	SimplParserT__2  = 3
	SimplParserT__3  = 4
	SimplParserT__4  = 5
	SimplParserT__5  = 6
	SimplParserT__6  = 7
	SimplParserT__7  = 8
	SimplParserT__8  = 9
	SimplParserT__9  = 10
	SimplParserT__10 = 11
	SimplParserT__11 = 12
	SimplParserT__12 = 13
	SimplParserT__13 = 14
	SimplParserT__14 = 15
	SimplParserT__15 = 16
	SimplParserT__16 = 17
	SimplParserT__17 = 18
	SimplParserT__18 = 19
	SimplParserT__19 = 20
	SimplParserBOOL  = 21
	SimplParserID    = 22
	SimplParserNUM   = 23
	SimplParserSTR   = 24
	SimplParserWS    = 25
)

// SimplParser rules.
const (
	SimplParserRULE_program = 0
	SimplParserRULE_expr    = 1
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimplParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimplParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ProgramContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimplParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SimplParserRULE_program)
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
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimplParserT__2)|(1<<SimplParserT__5)|(1<<SimplParserT__12)|(1<<SimplParserT__14)|(1<<SimplParserT__16)|(1<<SimplParserT__19)|(1<<SimplParserBOOL)|(1<<SimplParserID)|(1<<SimplParserNUM)|(1<<SimplParserSTR))) != 0 {
		{
			p.SetState(4)
			p.expr(0)
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimplParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimplParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StrExprContext struct {
	*ExprContext
}

func NewStrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrExprContext {
	var p = new(StrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrExprContext) STR() antlr.TerminalNode {
	return s.GetToken(SimplParserSTR, 0)
}

func (s *StrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterStrExpr(s)
	}
}

func (s *StrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitStrExpr(s)
	}
}

func (s *StrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitStrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DotExprContext struct {
	*ExprContext
	id    IExprContext
	field IExprContext
}

func NewDotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DotExprContext {
	var p = new(DotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DotExprContext) GetId() IExprContext { return s.id }

func (s *DotExprContext) GetField() IExprContext { return s.field }

func (s *DotExprContext) SetId(v IExprContext) { s.id = v }

func (s *DotExprContext) SetField(v IExprContext) { s.field = v }

func (s *DotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DotExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterDotExpr(s)
	}
}

func (s *DotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitDotExpr(s)
	}
}

func (s *DotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitDotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LenExprContext struct {
	*ExprContext
	id IExprContext
}

func NewLenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LenExprContext {
	var p = new(LenExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LenExprContext) GetId() IExprContext { return s.id }

func (s *LenExprContext) SetId(v IExprContext) { s.id = v }

func (s *LenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LenExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterLenExpr(s)
	}
}

func (s *LenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitLenExpr(s)
	}
}

func (s *LenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitLenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseExprContext struct {
	*ExprContext
	con IExprContext
	t   IExprContext
	f   IExprContext
}

func NewIfElseExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseExprContext {
	var p = new(IfElseExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IfElseExprContext) GetCon() IExprContext { return s.con }

func (s *IfElseExprContext) GetT() IExprContext { return s.t }

func (s *IfElseExprContext) GetF() IExprContext { return s.f }

func (s *IfElseExprContext) SetCon(v IExprContext) { s.con = v }

func (s *IfElseExprContext) SetT(v IExprContext) { s.t = v }

func (s *IfElseExprContext) SetF(v IExprContext) { s.f = v }

func (s *IfElseExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IfElseExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfElseExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterIfElseExpr(s)
	}
}

func (s *IfElseExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitIfElseExpr(s)
	}
}

func (s *IfElseExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitIfElseExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubExprContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubExprContext {
	var p = new(SubExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SubExprContext) GetLeft() IExprContext { return s.left }

func (s *SubExprContext) GetRight() IExprContext { return s.right }

func (s *SubExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *SubExprContext) SetRight(v IExprContext) { s.right = v }

func (s *SubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SubExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterSubExpr(s)
	}
}

func (s *SubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitSubExpr(s)
	}
}

func (s *SubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultExprContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewMultExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultExprContext {
	var p = new(MultExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MultExprContext) GetLeft() IExprContext { return s.left }

func (s *MultExprContext) GetRight() IExprContext { return s.right }

func (s *MultExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *MultExprContext) SetRight(v IExprContext) { s.right = v }

func (s *MultExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MultExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MultExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterMultExpr(s)
	}
}

func (s *MultExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitMultExpr(s)
	}
}

func (s *MultExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitMultExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CommentExprContext struct {
	*ExprContext
}

func NewCommentExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CommentExprContext {
	var p = new(CommentExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CommentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CommentExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterCommentExpr(s)
	}
}

func (s *CommentExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitCommentExpr(s)
	}
}

func (s *CommentExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitCommentExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	*ExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumExprContext struct {
	*ExprContext
}

func NewNumExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumExprContext {
	var p = new(NumExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumExprContext) NUM() antlr.TerminalNode {
	return s.GetToken(SimplParserNUM, 0)
}

func (s *NumExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterNumExpr(s)
	}
}

func (s *NumExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitNumExpr(s)
	}
}

func (s *NumExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitNumExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExprContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddExprContext) GetLeft() IExprContext { return s.left }

func (s *AddExprContext) GetRight() IExprContext { return s.right }

func (s *AddExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *AddExprContext) SetRight(v IExprContext) { s.right = v }

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LambdaExprContext struct {
	*ExprContext
	param IExprContext
	body  IExprContext
}

func NewLambdaExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LambdaExprContext {
	var p = new(LambdaExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LambdaExprContext) GetParam() IExprContext { return s.param }

func (s *LambdaExprContext) GetBody() IExprContext { return s.body }

func (s *LambdaExprContext) SetParam(v IExprContext) { s.param = v }

func (s *LambdaExprContext) SetBody(v IExprContext) { s.body = v }

func (s *LambdaExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LambdaExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LambdaExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterLambdaExpr(s)
	}
}

func (s *LambdaExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitLambdaExpr(s)
	}
}

func (s *LambdaExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitLambdaExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructExprContext struct {
	*ExprContext
}

func NewStructExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructExprContext {
	var p = new(StructExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StructExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *StructExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterStructExpr(s)
	}
}

func (s *StructExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitStructExpr(s)
	}
}

func (s *StructExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitStructExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LookupExprContext struct {
	*ExprContext
	id  IExprContext
	key IExprContext
}

func NewLookupExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LookupExprContext {
	var p = new(LookupExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LookupExprContext) GetId() IExprContext { return s.id }

func (s *LookupExprContext) GetKey() IExprContext { return s.key }

func (s *LookupExprContext) SetId(v IExprContext) { s.id = v }

func (s *LookupExprContext) SetKey(v IExprContext) { s.key = v }

func (s *LookupExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LookupExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LookupExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LookupExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterLookupExpr(s)
	}
}

func (s *LookupExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitLookupExpr(s)
	}
}

func (s *LookupExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitLookupExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DivExprContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivExprContext {
	var p = new(DivExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DivExprContext) GetLeft() IExprContext { return s.left }

func (s *DivExprContext) GetRight() IExprContext { return s.right }

func (s *DivExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *DivExprContext) SetRight(v IExprContext) { s.right = v }

func (s *DivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DivExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterDivExpr(s)
	}
}

func (s *DivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitDivExpr(s)
	}
}

func (s *DivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitDivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LookupAssignExprContext struct {
	*ExprContext
	id    IExprContext
	key   IExprContext
	value IExprContext
}

func NewLookupAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LookupAssignExprContext {
	var p = new(LookupAssignExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LookupAssignExprContext) GetId() IExprContext { return s.id }

func (s *LookupAssignExprContext) GetKey() IExprContext { return s.key }

func (s *LookupAssignExprContext) GetValue() IExprContext { return s.value }

func (s *LookupAssignExprContext) SetId(v IExprContext) { s.id = v }

func (s *LookupAssignExprContext) SetKey(v IExprContext) { s.key = v }

func (s *LookupAssignExprContext) SetValue(v IExprContext) { s.value = v }

func (s *LookupAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LookupAssignExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LookupAssignExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LookupAssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterLookupAssignExpr(s)
	}
}

func (s *LookupAssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitLookupAssignExpr(s)
	}
}

func (s *LookupAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitLookupAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolExprContext struct {
	*ExprContext
}

func NewBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExprContext {
	var p = new(BoolExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SimplParserBOOL, 0)
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExprContext struct {
	*ExprContext
	fun IExprContext
	arg IExprContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CallExprContext) GetFun() IExprContext { return s.fun }

func (s *CallExprContext) GetArg() IExprContext { return s.arg }

func (s *CallExprContext) SetFun(v IExprContext) { s.fun = v }

func (s *CallExprContext) SetArg(v IExprContext) { s.arg = v }

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CallExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListExprContext struct {
	*ExprContext
}

func NewListExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListExprContext {
	var p = new(ListExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ListExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ListExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterListExpr(s)
	}
}

func (s *ListExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitListExpr(s)
	}
}

func (s *ListExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitListExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignExprContext struct {
	*ExprContext
	id    IExprContext
	value IExprContext
}

func NewAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignExprContext {
	var p = new(AssignExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AssignExprContext) GetId() IExprContext { return s.id }

func (s *AssignExprContext) GetValue() IExprContext { return s.value }

func (s *AssignExprContext) SetId(v IExprContext) { s.id = v }

func (s *AssignExprContext) SetValue(v IExprContext) { s.value = v }

func (s *AssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AssignExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterAssignExpr(s)
	}
}

func (s *AssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitAssignExpr(s)
	}
}

func (s *AssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	*ExprContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SimplParserID, 0)
}

func (s *IdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterIdExpr(s)
	}
}

func (s *IdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitIdExpr(s)
	}
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualExprContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewEqualExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualExprContext {
	var p = new(EqualExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqualExprContext) GetLeft() IExprContext { return s.left }

func (s *EqualExprContext) GetRight() IExprContext { return s.right }

func (s *EqualExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *EqualExprContext) SetRight(v IExprContext) { s.right = v }

func (s *EqualExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqualExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.EnterEqualExpr(s)
	}
}

func (s *EqualExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimplListener); ok {
		listenerT.ExitEqualExpr(s)
	}
}

func (s *EqualExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimplVisitor:
		return t.VisitEqualExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimplParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SimplParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, SimplParserRULE_expr, _p)
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
	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimplParserT__12:
		localctx = NewCommentExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(11)
			p.Match(SimplParserT__12)
		}
		{
			p.SetState(12)
			p.expr(10)
		}

	case SimplParserT__2:
		localctx = NewListExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(13)
			p.Match(SimplParserT__2)
		}
		{
			p.SetState(14)
			p.expr(0)
		}
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SimplParserT__13 {
			{
				p.SetState(15)
				p.Match(SimplParserT__13)
			}
			{
				p.SetState(16)
				p.expr(0)
			}

			p.SetState(21)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(22)
			p.Match(SimplParserT__3)
		}

	case SimplParserT__5:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(24)
			p.Match(SimplParserT__5)
		}
		{
			p.SetState(25)
			p.expr(0)
		}
		{
			p.SetState(26)
			p.Match(SimplParserT__6)
		}

	case SimplParserT__14:
		localctx = NewStructExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(28)
			p.Match(SimplParserT__14)
		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimplParserT__2)|(1<<SimplParserT__5)|(1<<SimplParserT__12)|(1<<SimplParserT__14)|(1<<SimplParserT__16)|(1<<SimplParserT__19)|(1<<SimplParserBOOL)|(1<<SimplParserID)|(1<<SimplParserNUM)|(1<<SimplParserSTR))) != 0 {
			{
				p.SetState(29)
				p.expr(0)
			}

			p.SetState(34)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(35)
			p.Match(SimplParserT__15)
		}

	case SimplParserT__16:
		localctx = NewIfElseExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(36)
			p.Match(SimplParserT__16)
		}
		{
			p.SetState(37)

			var _x = p.expr(0)

			localctx.(*IfElseExprContext).con = _x
		}
		{
			p.SetState(38)
			p.Match(SimplParserT__17)
		}
		{
			p.SetState(39)

			var _x = p.expr(0)

			localctx.(*IfElseExprContext).t = _x
		}
		{
			p.SetState(40)
			p.Match(SimplParserT__18)
		}
		{
			p.SetState(41)

			var _x = p.expr(6)

			localctx.(*IfElseExprContext).f = _x
		}

	case SimplParserT__19:
		localctx = NewLenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(43)
			p.Match(SimplParserT__19)
		}
		{
			p.SetState(44)

			var _x = p.expr(5)

			localctx.(*LenExprContext).id = _x
		}

	case SimplParserBOOL:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(45)
			p.Match(SimplParserBOOL)
		}

	case SimplParserID:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(46)
			p.Match(SimplParserID)
		}

	case SimplParserNUM:
		localctx = NewNumExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(47)
			p.Match(SimplParserNUM)
		}

	case SimplParserSTR:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)
			p.Match(SimplParserSTR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEqualExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EqualExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimplParserRULE_expr)
				p.SetState(51)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(52)
					p.Match(SimplParserT__0)
				}
				{
					p.SetState(53)

					var _x = p.expr(22)

					localctx.(*EqualExprContext).right = _x
				}

			case 2:
				localctx = NewDotExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*DotExprContext).id = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimplParserRULE_expr)
				p.SetState(54)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(55)
					p.Match(SimplParserT__1)
				}
				{
					p.SetState(56)

					var _x = p.expr(21)

					localctx.(*DotExprContext).field = _x
				}

			case 3:
				localctx = NewLookupAssignExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LookupAssignExprContext).id = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimplParserRULE_expr)
				p.SetState(57)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(58)
					p.Match(SimplParserT__2)
				}
				{
					p.SetState(59)

					var _x = p.expr(0)

					localctx.(*LookupAssignExprContext).key = _x
				}
				{
					p.SetState(60)
					p.Match(SimplParserT__3)
				}
				{
					p.SetState(61)
					p.Match(SimplParserT__4)
				}
				{
					p.SetState(62)

					var _x = p.expr(19)

					localctx.(*LookupAssignExprContext).value = _x
				}

			case 4:
				localctx = NewDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*DivExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimplParserRULE_expr)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(65)
					p.Match(SimplParserT__7)
				}
				{
					p.SetState(66)

					var _x = p.expr(17)

					localctx.(*DivExprContext).right = _x
				}

			case 5:
				localctx = NewMultExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MultExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimplParserRULE_expr)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(68)
					p.Match(SimplParserT__8)
				}
				{
					p.SetState(69)

					var _x = p.expr(16)

					localctx.(*MultExprContext).right = _x
				}

			case 6:
				localctx = NewSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*SubExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimplParserRULE_expr)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(71)
					p.Match(SimplParserT__9)
				}
				{
					p.SetState(72)

					var _x = p.expr(15)

					localctx.(*SubExprContext).right = _x
				}

			case 7:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimplParserRULE_expr)
				p.SetState(73)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(74)
					p.Match(SimplParserT__10)
				}
				{
					p.SetState(75)

					var _x = p.expr(14)

					localctx.(*AddExprContext).right = _x
				}

			case 8:
				localctx = NewLambdaExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LambdaExprContext).param = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimplParserRULE_expr)
				p.SetState(76)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(77)
					p.Match(SimplParserT__11)
				}
				{
					p.SetState(78)

					var _x = p.expr(13)

					localctx.(*LambdaExprContext).body = _x
				}

			case 9:
				localctx = NewAssignExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AssignExprContext).id = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimplParserRULE_expr)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(80)
					p.Match(SimplParserT__4)
				}
				{
					p.SetState(81)

					var _x = p.expr(12)

					localctx.(*AssignExprContext).value = _x
				}

			case 10:
				localctx = NewLookupExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LookupExprContext).id = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimplParserRULE_expr)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(83)
					p.Match(SimplParserT__2)
				}
				{
					p.SetState(84)

					var _x = p.expr(0)

					localctx.(*LookupExprContext).key = _x
				}
				{
					p.SetState(85)
					p.Match(SimplParserT__3)
				}

			case 11:
				localctx = NewCallExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*CallExprContext).fun = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimplParserRULE_expr)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(88)
					p.Match(SimplParserT__5)
				}
				{
					p.SetState(89)

					var _x = p.expr(0)

					localctx.(*CallExprContext).arg = _x
				}
				{
					p.SetState(90)
					p.Match(SimplParserT__6)
				}

			}

		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

func (p *SimplParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SimplParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 17)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
