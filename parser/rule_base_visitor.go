// Code generated from /home/studing/GolandProjects/src/rule-engine/parser/Rule.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Rule

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRuleVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRuleVisitor) VisitAttrPath(ctx *AttrPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitSubAttr(ctx *SubAttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitFloat(ctx *FloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitNegativeValue(ctx *NegativeValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitArithmeticWithParenthesis(ctx *ArithmeticWithParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitAttr(ctx *AttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitCompareWithParenthesis(ctx *CompareWithParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitBasicCompare(ctx *BasicCompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitBasicLogical(ctx *BasicLogicalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitTwoLogical(ctx *TwoLogicalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitLogicalWithParenthesis(ctx *LogicalWithParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}
