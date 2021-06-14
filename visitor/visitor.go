package visitor

import (
	"fmt"
	"reflect"
	"strconv"

	"rule-engine/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tidwall/gjson"
)

type RuleVisitor struct {
	parser.BaseRuleVisitor

	Items string
}

func (v *RuleVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.BasicLogicalContext:
		return v.VisitBasicLogical(val)
	case *parser.TwoLogicalContext:
		return v.VisitTwoLogical(val)
	case *parser.LogicalWithParenthesisContext:
		return v.VisitLogicalWithParenthesis(val)
	case *parser.MulDivContext:
		return v.VisitMulDiv(val)
	case *parser.AddSubContext:
		return v.VisitAddSub(val)
	case *parser.NegativeValueContext:
		return v.VisitNegativeValue(val)
	case *parser.IntContext:
		return v.VisitInt(val)
	case *parser.FloatContext:
		return v.VisitFloat(val)
	case *parser.ArithmeticWithParenthesisContext:
		return v.VisitArithmeticWithParenthesis(val)
	// check
	case *parser.ArithmeticContext:
		return v.Visit(val)
	case *parser.BasicCompareContext:
		return v.VisitBasicCompare(val)
	case *parser.CompareWithParenthesisContext:
		return v.VisitCompareWithParenthesis(val)
	case *parser.AttrContext:
		return v.VisitAttr(val)
	default:
		panic(fmt.Errorf("usnknown context is %s", tree.GetText()))
	}
}

func (v *RuleVisitor) VisitAttrPath(ctx *parser.AttrPathContext) interface{} {
	return gjson.Get(v.Items, ctx.GetText()).Value()
}

func (v *RuleVisitor) VisitSubAttr(ctx *parser.SubAttrContext) interface{} {
	return nil
}

func (v *RuleVisitor) VisitFloat(ctx *parser.FloatContext) interface{} {
	val, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return val
}

func (v *RuleVisitor) VisitInt(ctx *parser.IntContext) interface{} {
	val, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return val
}

func (v *RuleVisitor) VisitAttr(ctx *parser.AttrContext) interface{} {
	val := gjson.Get(v.Items, ctx.GetText())
	if !val.Exists() {
		panic(
			fmt.Sprintf(
				"path:[%s] not present in %s",
				ctx.GetText(),
				v.Items,
			),
		)
	}
	return val
}

func (v *RuleVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	left := v.Visit(ctx.GetLeft())
	right := v.Visit(ctx.GetRight())

	switch ctx.GetOp().GetTokenType() {
	case parser.RuleParserADD:
		return add(left, right)
	case parser.RuleParserSUB:
		return sub(left, right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
	}
}

func (v *RuleVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	left := v.Visit(ctx.GetLeft())
	right := v.Visit(ctx.GetRight())

	switch ctx.GetOp().GetTokenType() {
	case parser.RuleParserMUL:
		return mul(left, right)
	case parser.RuleParserDIV:
		return div(left, right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
	}
}

func (v *RuleVisitor) VisitNegativeValue(
	ctx *parser.NegativeValueContext) interface{} {
	val := v.Visit(ctx.Arithmetic())
	switch k := reflect.TypeOf(val).Kind(); k {
	case reflect.Float64:
		return -1 * val.(float64)
	case reflect.Int64:
		return -1 * val.(int64)
	default:
		panic(fmt.Errorf("invalid arithmetic: %s", ctx.GetText()))
	}
}

func (v *RuleVisitor) VisitArithmeticWithParenthesis(
	ctx *parser.ArithmeticWithParenthesisContext) interface{} {
	return v.Visit(ctx.Arithmetic())
}

func (v *RuleVisitor) VisitCompareWithParenthesis(
	ctx *parser.CompareWithParenthesisContext) interface{} {
	return v.Visit(ctx.Compare())
}

func (v *RuleVisitor) VisitBasicCompare(
	ctx *parser.BasicCompareContext) interface{} {
	left := v.Visit(ctx.GetLeft())
	right := v.Visit(ctx.GetRight())
	switch op := ctx.GetOp().GetTokenType(); op {
	case parser.RuleLexerEQ:
		return eq(left, right)
	case parser.RuleLexerNE:
		return nq(left, right)
	case parser.RuleLexerGT:
		return gt(left, right)
	case parser.RuleLexerLT:
		return lt(left, right)
	case parser.RuleLexerGE:
		return ge(left, right)
	case parser.RuleLexerLE:
		return le(left, right)
	default:
		panic(fmt.Errorf("invalid compare op %s", ctx.GetOp().GetText()))
	}
}

func (v *RuleVisitor) VisitAndLogical(
	ctx *parser.TwoLogicalContext,
) interface{} {
	return v.VisitChildren(ctx)
}

func (v *RuleVisitor) VisitBasicLogical(
	ctx *parser.BasicLogicalContext,
) interface{} {
	return v.Visit(ctx.Compare())
}

func (v *RuleVisitor) VisitTwoLogical(
	ctx *parser.TwoLogicalContext,
) interface{} {
	left, right := v.Visit(ctx.GetLeft()), v.Visit(ctx.GetRight())
	leftBoolVal, leftOk := left.(bool)
	rightBoolVal, rightOk := right.(bool)
	if !leftOk || !rightOk {
		panic("invalid logical value")
	}
	switch op := ctx.GetOp().GetTokenType(); op {
	case parser.RuleLexerAND:
		return leftBoolVal && rightBoolVal
	case parser.RuleLexerOR:
		return leftBoolVal || rightBoolVal
	default:
		panic("invalid logical operation")
	}
}

func (v *RuleVisitor) VisitLogicalWithParenthesis(
	ctx *parser.LogicalWithParenthesisContext) interface{} {
	return v.Visit(ctx.Logical())
}
