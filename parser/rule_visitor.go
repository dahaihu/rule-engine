// Code generated from /home/studing/GolandProjects/src/rule-engine/parser/Rule.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Rule

import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by RuleParser.
type RuleVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RuleParser#attrPath.
	VisitAttrPath(ctx *AttrPathContext) interface{}

	// Visit a parse tree produced by RuleParser#subAttr.
	VisitSubAttr(ctx *SubAttrContext) interface{}

	// Visit a parse tree produced by RuleParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by RuleParser#Float.
	VisitFloat(ctx *FloatContext) interface{}

	// Visit a parse tree produced by RuleParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by RuleParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by RuleParser#NegativeValue.
	VisitNegativeValue(ctx *NegativeValueContext) interface{}

	// Visit a parse tree produced by RuleParser#ArithmeticWithParenthesis.
	VisitArithmeticWithParenthesis(ctx *ArithmeticWithParenthesisContext) interface{}

	// Visit a parse tree produced by RuleParser#Int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by RuleParser#Attr.
	VisitAttr(ctx *AttrContext) interface{}

	// Visit a parse tree produced by RuleParser#CompareWithParenthesis.
	VisitCompareWithParenthesis(ctx *CompareWithParenthesisContext) interface{}

	// Visit a parse tree produced by RuleParser#BasicCompare.
	VisitBasicCompare(ctx *BasicCompareContext) interface{}

	// Visit a parse tree produced by RuleParser#BasicLogical.
	VisitBasicLogical(ctx *BasicLogicalContext) interface{}

	// Visit a parse tree produced by RuleParser#TwoLogical.
	VisitTwoLogical(ctx *TwoLogicalContext) interface{}

	// Visit a parse tree produced by RuleParser#LogicalWithParenthesis.
	VisitLogicalWithParenthesis(ctx *LogicalWithParenthesisContext) interface{}

}