// Code generated from TqParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tqgrammar // TqParser
import "github.com/antlr4-go/antlr/v4"

// TqParserListener is a complete listener for a parse tree produced by TqParser.
type TqParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterSelectStage is called when entering the selectStage production.
	EnterSelectStage(c *SelectStageContext)

	// EnterDropStage is called when entering the dropStage production.
	EnterDropStage(c *DropStageContext)

	// EnterWhereStage is called when entering the whereStage production.
	EnterWhereStage(c *WhereStageContext)

	// EnterDeriveStage is called when entering the deriveStage production.
	EnterDeriveStage(c *DeriveStageContext)

	// EnterRenameStage is called when entering the renameStage production.
	EnterRenameStage(c *RenameStageContext)

	// EnterSortStage is called when entering the sortStage production.
	EnterSortStage(c *SortStageContext)

	// EnterDistinctStage is called when entering the distinctStage production.
	EnterDistinctStage(c *DistinctStageContext)

	// EnterLimitStage is called when entering the limitStage production.
	EnterLimitStage(c *LimitStageContext)

	// EnterOffsetStage is called when entering the offsetStage production.
	EnterOffsetStage(c *OffsetStageContext)

	// EnterGroupStage is called when entering the groupStage production.
	EnterGroupStage(c *GroupStageContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterRenamePair is called when entering the renamePair production.
	EnterRenamePair(c *RenamePairContext)

	// EnterSortKey is called when entering the sortKey production.
	EnterSortKey(c *SortKeyContext)

	// EnterColumnList is called when entering the columnList production.
	EnterColumnList(c *ColumnListContext)

	// EnterColumnItem is called when entering the columnItem production.
	EnterColumnItem(c *ColumnItemContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterBareName is called when entering the bareName production.
	EnterBareName(c *BareNameContext)

	// EnterVerbKeyword is called when entering the verbKeyword production.
	EnterVerbKeyword(c *VerbKeywordContext)

	// EnterTqPowExpr is called when entering the tqPowExpr production.
	EnterTqPowExpr(c *TqPowExprContext)

	// EnterTqCompareExpr is called when entering the tqCompareExpr production.
	EnterTqCompareExpr(c *TqCompareExprContext)

	// EnterTqBoolExpr is called when entering the tqBoolExpr production.
	EnterTqBoolExpr(c *TqBoolExprContext)

	// EnterTqNumberExpr is called when entering the tqNumberExpr production.
	EnterTqNumberExpr(c *TqNumberExprContext)

	// EnterTqPercentExpr is called when entering the tqPercentExpr production.
	EnterTqPercentExpr(c *TqPercentExprContext)

	// EnterTqConcatExpr is called when entering the tqConcatExpr production.
	EnterTqConcatExpr(c *TqConcatExprContext)

	// EnterTqParenExpr is called when entering the tqParenExpr production.
	EnterTqParenExpr(c *TqParenExprContext)

	// EnterTqUnaryExpr is called when entering the tqUnaryExpr production.
	EnterTqUnaryExpr(c *TqUnaryExprContext)

	// EnterTqCallExpr is called when entering the tqCallExpr production.
	EnterTqCallExpr(c *TqCallExprContext)

	// EnterTqRefExpr is called when entering the tqRefExpr production.
	EnterTqRefExpr(c *TqRefExprContext)

	// EnterTqColumnExpr is called when entering the tqColumnExpr production.
	EnterTqColumnExpr(c *TqColumnExprContext)

	// EnterTqStringExpr is called when entering the tqStringExpr production.
	EnterTqStringExpr(c *TqStringExprContext)

	// EnterTqAddExpr is called when entering the tqAddExpr production.
	EnterTqAddExpr(c *TqAddExprContext)

	// EnterTqErrorExpr is called when entering the tqErrorExpr production.
	EnterTqErrorExpr(c *TqErrorExprContext)

	// EnterTqMulExpr is called when entering the tqMulExpr production.
	EnterTqMulExpr(c *TqMulExprContext)

	// EnterTqFunctionCall is called when entering the tqFunctionCall production.
	EnterTqFunctionCall(c *TqFunctionCallContext)

	// EnterTqArgList is called when entering the tqArgList production.
	EnterTqArgList(c *TqArgListContext)

	// EnterErrorExpr is called when entering the errorExpr production.
	EnterErrorExpr(c *ErrorExprContext)

	// EnterPipeExpr is called when entering the pipeExpr production.
	EnterPipeExpr(c *PipeExprContext)

	// EnterNumberExpr is called when entering the numberExpr production.
	EnterNumberExpr(c *NumberExprContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterConcatExpr is called when entering the concatExpr production.
	EnterConcatExpr(c *ConcatExprContext)

	// EnterStringExpr is called when entering the stringExpr production.
	EnterStringExpr(c *StringExprContext)

	// EnterUnaryExpr is called when entering the unaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterRefExpr is called when entering the refExpr production.
	EnterRefExpr(c *RefExprContext)

	// EnterMulExpr is called when entering the mulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterPercentExpr is called when entering the percentExpr production.
	EnterPercentExpr(c *PercentExprContext)

	// EnterCallExpr is called when entering the callExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterBoolExpr is called when entering the boolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterPowExpr is called when entering the powExpr production.
	EnterPowExpr(c *PowExprContext)

	// EnterCompareExpr is called when entering the compareExpr production.
	EnterCompareExpr(c *CompareExprContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterSheetQualifier is called when entering the sheetQualifier production.
	EnterSheetQualifier(c *SheetQualifierContext)

	// EnterCellRef is called when entering the cellRef production.
	EnterCellRef(c *CellRefContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitSelectStage is called when exiting the selectStage production.
	ExitSelectStage(c *SelectStageContext)

	// ExitDropStage is called when exiting the dropStage production.
	ExitDropStage(c *DropStageContext)

	// ExitWhereStage is called when exiting the whereStage production.
	ExitWhereStage(c *WhereStageContext)

	// ExitDeriveStage is called when exiting the deriveStage production.
	ExitDeriveStage(c *DeriveStageContext)

	// ExitRenameStage is called when exiting the renameStage production.
	ExitRenameStage(c *RenameStageContext)

	// ExitSortStage is called when exiting the sortStage production.
	ExitSortStage(c *SortStageContext)

	// ExitDistinctStage is called when exiting the distinctStage production.
	ExitDistinctStage(c *DistinctStageContext)

	// ExitLimitStage is called when exiting the limitStage production.
	ExitLimitStage(c *LimitStageContext)

	// ExitOffsetStage is called when exiting the offsetStage production.
	ExitOffsetStage(c *OffsetStageContext)

	// ExitGroupStage is called when exiting the groupStage production.
	ExitGroupStage(c *GroupStageContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitRenamePair is called when exiting the renamePair production.
	ExitRenamePair(c *RenamePairContext)

	// ExitSortKey is called when exiting the sortKey production.
	ExitSortKey(c *SortKeyContext)

	// ExitColumnList is called when exiting the columnList production.
	ExitColumnList(c *ColumnListContext)

	// ExitColumnItem is called when exiting the columnItem production.
	ExitColumnItem(c *ColumnItemContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitBareName is called when exiting the bareName production.
	ExitBareName(c *BareNameContext)

	// ExitVerbKeyword is called when exiting the verbKeyword production.
	ExitVerbKeyword(c *VerbKeywordContext)

	// ExitTqPowExpr is called when exiting the tqPowExpr production.
	ExitTqPowExpr(c *TqPowExprContext)

	// ExitTqCompareExpr is called when exiting the tqCompareExpr production.
	ExitTqCompareExpr(c *TqCompareExprContext)

	// ExitTqBoolExpr is called when exiting the tqBoolExpr production.
	ExitTqBoolExpr(c *TqBoolExprContext)

	// ExitTqNumberExpr is called when exiting the tqNumberExpr production.
	ExitTqNumberExpr(c *TqNumberExprContext)

	// ExitTqPercentExpr is called when exiting the tqPercentExpr production.
	ExitTqPercentExpr(c *TqPercentExprContext)

	// ExitTqConcatExpr is called when exiting the tqConcatExpr production.
	ExitTqConcatExpr(c *TqConcatExprContext)

	// ExitTqParenExpr is called when exiting the tqParenExpr production.
	ExitTqParenExpr(c *TqParenExprContext)

	// ExitTqUnaryExpr is called when exiting the tqUnaryExpr production.
	ExitTqUnaryExpr(c *TqUnaryExprContext)

	// ExitTqCallExpr is called when exiting the tqCallExpr production.
	ExitTqCallExpr(c *TqCallExprContext)

	// ExitTqRefExpr is called when exiting the tqRefExpr production.
	ExitTqRefExpr(c *TqRefExprContext)

	// ExitTqColumnExpr is called when exiting the tqColumnExpr production.
	ExitTqColumnExpr(c *TqColumnExprContext)

	// ExitTqStringExpr is called when exiting the tqStringExpr production.
	ExitTqStringExpr(c *TqStringExprContext)

	// ExitTqAddExpr is called when exiting the tqAddExpr production.
	ExitTqAddExpr(c *TqAddExprContext)

	// ExitTqErrorExpr is called when exiting the tqErrorExpr production.
	ExitTqErrorExpr(c *TqErrorExprContext)

	// ExitTqMulExpr is called when exiting the tqMulExpr production.
	ExitTqMulExpr(c *TqMulExprContext)

	// ExitTqFunctionCall is called when exiting the tqFunctionCall production.
	ExitTqFunctionCall(c *TqFunctionCallContext)

	// ExitTqArgList is called when exiting the tqArgList production.
	ExitTqArgList(c *TqArgListContext)

	// ExitErrorExpr is called when exiting the errorExpr production.
	ExitErrorExpr(c *ErrorExprContext)

	// ExitPipeExpr is called when exiting the pipeExpr production.
	ExitPipeExpr(c *PipeExprContext)

	// ExitNumberExpr is called when exiting the numberExpr production.
	ExitNumberExpr(c *NumberExprContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitConcatExpr is called when exiting the concatExpr production.
	ExitConcatExpr(c *ConcatExprContext)

	// ExitStringExpr is called when exiting the stringExpr production.
	ExitStringExpr(c *StringExprContext)

	// ExitUnaryExpr is called when exiting the unaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitRefExpr is called when exiting the refExpr production.
	ExitRefExpr(c *RefExprContext)

	// ExitMulExpr is called when exiting the mulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitPercentExpr is called when exiting the percentExpr production.
	ExitPercentExpr(c *PercentExprContext)

	// ExitCallExpr is called when exiting the callExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitBoolExpr is called when exiting the boolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitPowExpr is called when exiting the powExpr production.
	ExitPowExpr(c *PowExprContext)

	// ExitCompareExpr is called when exiting the compareExpr production.
	ExitCompareExpr(c *CompareExprContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitSheetQualifier is called when exiting the sheetQualifier production.
	ExitSheetQualifier(c *SheetQualifierContext)

	// ExitCellRef is called when exiting the cellRef production.
	ExitCellRef(c *CellRefContext)
}
