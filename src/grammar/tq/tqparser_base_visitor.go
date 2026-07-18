// Code generated from TqParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tqgrammar // TqParser
import "github.com/antlr4-go/antlr/v4"

type BaseTqParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTqParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitSelectStage(ctx *SelectStageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitDropStage(ctx *DropStageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitWhereStage(ctx *WhereStageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitDeriveStage(ctx *DeriveStageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitRenameStage(ctx *RenameStageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitSortStage(ctx *SortStageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitDistinctStage(ctx *DistinctStageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitLimitStage(ctx *LimitStageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitOffsetStage(ctx *OffsetStageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitGroupStage(ctx *GroupStageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitRenamePair(ctx *RenamePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitSortKey(ctx *SortKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitColumnList(ctx *ColumnListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitColumnItem(ctx *ColumnItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitBareName(ctx *BareNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitVerbKeyword(ctx *VerbKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqPowExpr(ctx *TqPowExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqCompareExpr(ctx *TqCompareExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqBoolExpr(ctx *TqBoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqNumberExpr(ctx *TqNumberExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqPercentExpr(ctx *TqPercentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqConcatExpr(ctx *TqConcatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqParenExpr(ctx *TqParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqUnaryExpr(ctx *TqUnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqCallExpr(ctx *TqCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqRefExpr(ctx *TqRefExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqColumnExpr(ctx *TqColumnExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqStringExpr(ctx *TqStringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqAddExpr(ctx *TqAddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqErrorExpr(ctx *TqErrorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqMulExpr(ctx *TqMulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqFunctionCall(ctx *TqFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitTqArgList(ctx *TqArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitErrorExpr(ctx *ErrorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitPipeExpr(ctx *PipeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitNumberExpr(ctx *NumberExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitConcatExpr(ctx *ConcatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitRefExpr(ctx *RefExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitMulExpr(ctx *MulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitPercentExpr(ctx *PercentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitPowExpr(ctx *PowExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitCompareExpr(ctx *CompareExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitReference(ctx *ReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitSheetQualifier(ctx *SheetQualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTqParserVisitor) VisitCellRef(ctx *CellRefContext) interface{} {
	return v.VisitChildren(ctx)
}
