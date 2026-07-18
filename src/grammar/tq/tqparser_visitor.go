// Code generated from TqParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tqgrammar // TqParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TqParser.
type TqParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TqParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by TqParser#selectStage.
	VisitSelectStage(ctx *SelectStageContext) interface{}

	// Visit a parse tree produced by TqParser#dropStage.
	VisitDropStage(ctx *DropStageContext) interface{}

	// Visit a parse tree produced by TqParser#whereStage.
	VisitWhereStage(ctx *WhereStageContext) interface{}

	// Visit a parse tree produced by TqParser#deriveStage.
	VisitDeriveStage(ctx *DeriveStageContext) interface{}

	// Visit a parse tree produced by TqParser#renameStage.
	VisitRenameStage(ctx *RenameStageContext) interface{}

	// Visit a parse tree produced by TqParser#sortStage.
	VisitSortStage(ctx *SortStageContext) interface{}

	// Visit a parse tree produced by TqParser#distinctStage.
	VisitDistinctStage(ctx *DistinctStageContext) interface{}

	// Visit a parse tree produced by TqParser#limitStage.
	VisitLimitStage(ctx *LimitStageContext) interface{}

	// Visit a parse tree produced by TqParser#offsetStage.
	VisitOffsetStage(ctx *OffsetStageContext) interface{}

	// Visit a parse tree produced by TqParser#groupStage.
	VisitGroupStage(ctx *GroupStageContext) interface{}

	// Visit a parse tree produced by TqParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by TqParser#renamePair.
	VisitRenamePair(ctx *RenamePairContext) interface{}

	// Visit a parse tree produced by TqParser#sortKey.
	VisitSortKey(ctx *SortKeyContext) interface{}

	// Visit a parse tree produced by TqParser#columnList.
	VisitColumnList(ctx *ColumnListContext) interface{}

	// Visit a parse tree produced by TqParser#columnItem.
	VisitColumnItem(ctx *ColumnItemContext) interface{}

	// Visit a parse tree produced by TqParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by TqParser#bareName.
	VisitBareName(ctx *BareNameContext) interface{}

	// Visit a parse tree produced by TqParser#verbKeyword.
	VisitVerbKeyword(ctx *VerbKeywordContext) interface{}

	// Visit a parse tree produced by TqParser#tqPowExpr.
	VisitTqPowExpr(ctx *TqPowExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqCompareExpr.
	VisitTqCompareExpr(ctx *TqCompareExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqBoolExpr.
	VisitTqBoolExpr(ctx *TqBoolExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqNumberExpr.
	VisitTqNumberExpr(ctx *TqNumberExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqPercentExpr.
	VisitTqPercentExpr(ctx *TqPercentExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqConcatExpr.
	VisitTqConcatExpr(ctx *TqConcatExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqParenExpr.
	VisitTqParenExpr(ctx *TqParenExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqUnaryExpr.
	VisitTqUnaryExpr(ctx *TqUnaryExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqCallExpr.
	VisitTqCallExpr(ctx *TqCallExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqRefExpr.
	VisitTqRefExpr(ctx *TqRefExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqColumnExpr.
	VisitTqColumnExpr(ctx *TqColumnExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqStringExpr.
	VisitTqStringExpr(ctx *TqStringExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqAddExpr.
	VisitTqAddExpr(ctx *TqAddExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqErrorExpr.
	VisitTqErrorExpr(ctx *TqErrorExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqMulExpr.
	VisitTqMulExpr(ctx *TqMulExprContext) interface{}

	// Visit a parse tree produced by TqParser#tqFunctionCall.
	VisitTqFunctionCall(ctx *TqFunctionCallContext) interface{}

	// Visit a parse tree produced by TqParser#tqArgList.
	VisitTqArgList(ctx *TqArgListContext) interface{}

	// Visit a parse tree produced by TqParser#errorExpr.
	VisitErrorExpr(ctx *ErrorExprContext) interface{}

	// Visit a parse tree produced by TqParser#pipeExpr.
	VisitPipeExpr(ctx *PipeExprContext) interface{}

	// Visit a parse tree produced by TqParser#numberExpr.
	VisitNumberExpr(ctx *NumberExprContext) interface{}

	// Visit a parse tree produced by TqParser#parenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by TqParser#concatExpr.
	VisitConcatExpr(ctx *ConcatExprContext) interface{}

	// Visit a parse tree produced by TqParser#stringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by TqParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by TqParser#addExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by TqParser#refExpr.
	VisitRefExpr(ctx *RefExprContext) interface{}

	// Visit a parse tree produced by TqParser#mulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by TqParser#percentExpr.
	VisitPercentExpr(ctx *PercentExprContext) interface{}

	// Visit a parse tree produced by TqParser#callExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by TqParser#boolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by TqParser#powExpr.
	VisitPowExpr(ctx *PowExprContext) interface{}

	// Visit a parse tree produced by TqParser#compareExpr.
	VisitCompareExpr(ctx *CompareExprContext) interface{}

	// Visit a parse tree produced by TqParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by TqParser#argList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by TqParser#reference.
	VisitReference(ctx *ReferenceContext) interface{}

	// Visit a parse tree produced by TqParser#sheetQualifier.
	VisitSheetQualifier(ctx *SheetQualifierContext) interface{}

	// Visit a parse tree produced by TqParser#cellRef.
	VisitCellRef(ctx *CellRefContext) interface{}
}
