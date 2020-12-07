// Code generated from Simpl.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Simpl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSimplListener is a complete listener for a parse tree produced by SimplParser.
type BaseSimplListener struct{}

var _ SimplListener = &BaseSimplListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimplListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimplListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimplListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimplListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseSimplListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseSimplListener) ExitProgram(ctx *ProgramContext) {}

// EnterStrExpr is called when production strExpr is entered.
func (s *BaseSimplListener) EnterStrExpr(ctx *StrExprContext) {}

// ExitStrExpr is called when production strExpr is exited.
func (s *BaseSimplListener) ExitStrExpr(ctx *StrExprContext) {}

// EnterDotExpr is called when production dotExpr is entered.
func (s *BaseSimplListener) EnterDotExpr(ctx *DotExprContext) {}

// ExitDotExpr is called when production dotExpr is exited.
func (s *BaseSimplListener) ExitDotExpr(ctx *DotExprContext) {}

// EnterIfElseExpr is called when production ifElseExpr is entered.
func (s *BaseSimplListener) EnterIfElseExpr(ctx *IfElseExprContext) {}

// ExitIfElseExpr is called when production ifElseExpr is exited.
func (s *BaseSimplListener) ExitIfElseExpr(ctx *IfElseExprContext) {}

// EnterSubExpr is called when production subExpr is entered.
func (s *BaseSimplListener) EnterSubExpr(ctx *SubExprContext) {}

// ExitSubExpr is called when production subExpr is exited.
func (s *BaseSimplListener) ExitSubExpr(ctx *SubExprContext) {}

// EnterMultExpr is called when production multExpr is entered.
func (s *BaseSimplListener) EnterMultExpr(ctx *MultExprContext) {}

// ExitMultExpr is called when production multExpr is exited.
func (s *BaseSimplListener) ExitMultExpr(ctx *MultExprContext) {}

// EnterCommentExpr is called when production commentExpr is entered.
func (s *BaseSimplListener) EnterCommentExpr(ctx *CommentExprContext) {}

// ExitCommentExpr is called when production commentExpr is exited.
func (s *BaseSimplListener) ExitCommentExpr(ctx *CommentExprContext) {}

// EnterParenExpr is called when production parenExpr is entered.
func (s *BaseSimplListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production parenExpr is exited.
func (s *BaseSimplListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterNumExpr is called when production numExpr is entered.
func (s *BaseSimplListener) EnterNumExpr(ctx *NumExprContext) {}

// ExitNumExpr is called when production numExpr is exited.
func (s *BaseSimplListener) ExitNumExpr(ctx *NumExprContext) {}

// EnterCommaExpr is called when production commaExpr is entered.
func (s *BaseSimplListener) EnterCommaExpr(ctx *CommaExprContext) {}

// ExitCommaExpr is called when production commaExpr is exited.
func (s *BaseSimplListener) ExitCommaExpr(ctx *CommaExprContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseSimplListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseSimplListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterLambdaExpr is called when production lambdaExpr is entered.
func (s *BaseSimplListener) EnterLambdaExpr(ctx *LambdaExprContext) {}

// ExitLambdaExpr is called when production lambdaExpr is exited.
func (s *BaseSimplListener) ExitLambdaExpr(ctx *LambdaExprContext) {}

// EnterStructExpr is called when production structExpr is entered.
func (s *BaseSimplListener) EnterStructExpr(ctx *StructExprContext) {}

// ExitStructExpr is called when production structExpr is exited.
func (s *BaseSimplListener) ExitStructExpr(ctx *StructExprContext) {}

// EnterLookupExpr is called when production lookupExpr is entered.
func (s *BaseSimplListener) EnterLookupExpr(ctx *LookupExprContext) {}

// ExitLookupExpr is called when production lookupExpr is exited.
func (s *BaseSimplListener) ExitLookupExpr(ctx *LookupExprContext) {}

// EnterDivExpr is called when production divExpr is entered.
func (s *BaseSimplListener) EnterDivExpr(ctx *DivExprContext) {}

// ExitDivExpr is called when production divExpr is exited.
func (s *BaseSimplListener) ExitDivExpr(ctx *DivExprContext) {}

// EnterBoolExpr is called when production boolExpr is entered.
func (s *BaseSimplListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production boolExpr is exited.
func (s *BaseSimplListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterCallExpr is called when production callExpr is entered.
func (s *BaseSimplListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production callExpr is exited.
func (s *BaseSimplListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterListExpr is called when production listExpr is entered.
func (s *BaseSimplListener) EnterListExpr(ctx *ListExprContext) {}

// ExitListExpr is called when production listExpr is exited.
func (s *BaseSimplListener) ExitListExpr(ctx *ListExprContext) {}

// EnterAssignExpr is called when production assignExpr is entered.
func (s *BaseSimplListener) EnterAssignExpr(ctx *AssignExprContext) {}

// ExitAssignExpr is called when production assignExpr is exited.
func (s *BaseSimplListener) ExitAssignExpr(ctx *AssignExprContext) {}

// EnterIdExpr is called when production idExpr is entered.
func (s *BaseSimplListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production idExpr is exited.
func (s *BaseSimplListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterEqualExpr is called when production equalExpr is entered.
func (s *BaseSimplListener) EnterEqualExpr(ctx *EqualExprContext) {}

// ExitEqualExpr is called when production equalExpr is exited.
func (s *BaseSimplListener) ExitEqualExpr(ctx *EqualExprContext) {}
