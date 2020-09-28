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

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseSimplListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseSimplListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterMultExpr is called when production multExpr is entered.
func (s *BaseSimplListener) EnterMultExpr(ctx *MultExprContext) {}

// ExitMultExpr is called when production multExpr is exited.
func (s *BaseSimplListener) ExitMultExpr(ctx *MultExprContext) {}

// EnterNumExpr is called when production numExpr is entered.
func (s *BaseSimplListener) EnterNumExpr(ctx *NumExprContext) {}

// ExitNumExpr is called when production numExpr is exited.
func (s *BaseSimplListener) ExitNumExpr(ctx *NumExprContext) {}

// EnterParenExpr is called when production parenExpr is entered.
func (s *BaseSimplListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production parenExpr is exited.
func (s *BaseSimplListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterLetExpr is called when production letExpr is entered.
func (s *BaseSimplListener) EnterLetExpr(ctx *LetExprContext) {}

// ExitLetExpr is called when production letExpr is exited.
func (s *BaseSimplListener) ExitLetExpr(ctx *LetExprContext) {}

// EnterIdExpr is called when production idExpr is entered.
func (s *BaseSimplListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production idExpr is exited.
func (s *BaseSimplListener) ExitIdExpr(ctx *IdExprContext) {}
