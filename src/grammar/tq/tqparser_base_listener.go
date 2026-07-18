// Code generated from TqParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tqgrammar // TqParser
import "github.com/antlr4-go/antlr/v4"

// BaseTqParserListener is a complete listener for a parse tree produced by TqParser.
type BaseTqParserListener struct{}

var _ TqParserListener = &BaseTqParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTqParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTqParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTqParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTqParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseTqParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseTqParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterSelectStage is called when production selectStage is entered.
func (s *BaseTqParserListener) EnterSelectStage(ctx *SelectStageContext) {}

// ExitSelectStage is called when production selectStage is exited.
func (s *BaseTqParserListener) ExitSelectStage(ctx *SelectStageContext) {}

// EnterDropStage is called when production dropStage is entered.
func (s *BaseTqParserListener) EnterDropStage(ctx *DropStageContext) {}

// ExitDropStage is called when production dropStage is exited.
func (s *BaseTqParserListener) ExitDropStage(ctx *DropStageContext) {}

// EnterWhereStage is called when production whereStage is entered.
func (s *BaseTqParserListener) EnterWhereStage(ctx *WhereStageContext) {}

// ExitWhereStage is called when production whereStage is exited.
func (s *BaseTqParserListener) ExitWhereStage(ctx *WhereStageContext) {}

// EnterDeriveStage is called when production deriveStage is entered.
func (s *BaseTqParserListener) EnterDeriveStage(ctx *DeriveStageContext) {}

// ExitDeriveStage is called when production deriveStage is exited.
func (s *BaseTqParserListener) ExitDeriveStage(ctx *DeriveStageContext) {}

// EnterRenameStage is called when production renameStage is entered.
func (s *BaseTqParserListener) EnterRenameStage(ctx *RenameStageContext) {}

// ExitRenameStage is called when production renameStage is exited.
func (s *BaseTqParserListener) ExitRenameStage(ctx *RenameStageContext) {}

// EnterSortStage is called when production sortStage is entered.
func (s *BaseTqParserListener) EnterSortStage(ctx *SortStageContext) {}

// ExitSortStage is called when production sortStage is exited.
func (s *BaseTqParserListener) ExitSortStage(ctx *SortStageContext) {}

// EnterDistinctStage is called when production distinctStage is entered.
func (s *BaseTqParserListener) EnterDistinctStage(ctx *DistinctStageContext) {}

// ExitDistinctStage is called when production distinctStage is exited.
func (s *BaseTqParserListener) ExitDistinctStage(ctx *DistinctStageContext) {}

// EnterLimitStage is called when production limitStage is entered.
func (s *BaseTqParserListener) EnterLimitStage(ctx *LimitStageContext) {}

// ExitLimitStage is called when production limitStage is exited.
func (s *BaseTqParserListener) ExitLimitStage(ctx *LimitStageContext) {}

// EnterOffsetStage is called when production offsetStage is entered.
func (s *BaseTqParserListener) EnterOffsetStage(ctx *OffsetStageContext) {}

// ExitOffsetStage is called when production offsetStage is exited.
func (s *BaseTqParserListener) ExitOffsetStage(ctx *OffsetStageContext) {}

// EnterGroupStage is called when production groupStage is entered.
func (s *BaseTqParserListener) EnterGroupStage(ctx *GroupStageContext) {}

// ExitGroupStage is called when production groupStage is exited.
func (s *BaseTqParserListener) ExitGroupStage(ctx *GroupStageContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseTqParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseTqParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterRenamePair is called when production renamePair is entered.
func (s *BaseTqParserListener) EnterRenamePair(ctx *RenamePairContext) {}

// ExitRenamePair is called when production renamePair is exited.
func (s *BaseTqParserListener) ExitRenamePair(ctx *RenamePairContext) {}

// EnterSortKey is called when production sortKey is entered.
func (s *BaseTqParserListener) EnterSortKey(ctx *SortKeyContext) {}

// ExitSortKey is called when production sortKey is exited.
func (s *BaseTqParserListener) ExitSortKey(ctx *SortKeyContext) {}

// EnterColumnList is called when production columnList is entered.
func (s *BaseTqParserListener) EnterColumnList(ctx *ColumnListContext) {}

// ExitColumnList is called when production columnList is exited.
func (s *BaseTqParserListener) ExitColumnList(ctx *ColumnListContext) {}

// EnterColumnItem is called when production columnItem is entered.
func (s *BaseTqParserListener) EnterColumnItem(ctx *ColumnItemContext) {}

// ExitColumnItem is called when production columnItem is exited.
func (s *BaseTqParserListener) ExitColumnItem(ctx *ColumnItemContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseTqParserListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseTqParserListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterBareName is called when production bareName is entered.
func (s *BaseTqParserListener) EnterBareName(ctx *BareNameContext) {}

// ExitBareName is called when production bareName is exited.
func (s *BaseTqParserListener) ExitBareName(ctx *BareNameContext) {}

// EnterVerbKeyword is called when production verbKeyword is entered.
func (s *BaseTqParserListener) EnterVerbKeyword(ctx *VerbKeywordContext) {}

// ExitVerbKeyword is called when production verbKeyword is exited.
func (s *BaseTqParserListener) ExitVerbKeyword(ctx *VerbKeywordContext) {}

// EnterTqPowExpr is called when production tqPowExpr is entered.
func (s *BaseTqParserListener) EnterTqPowExpr(ctx *TqPowExprContext) {}

// ExitTqPowExpr is called when production tqPowExpr is exited.
func (s *BaseTqParserListener) ExitTqPowExpr(ctx *TqPowExprContext) {}

// EnterTqCompareExpr is called when production tqCompareExpr is entered.
func (s *BaseTqParserListener) EnterTqCompareExpr(ctx *TqCompareExprContext) {}

// ExitTqCompareExpr is called when production tqCompareExpr is exited.
func (s *BaseTqParserListener) ExitTqCompareExpr(ctx *TqCompareExprContext) {}

// EnterTqBoolExpr is called when production tqBoolExpr is entered.
func (s *BaseTqParserListener) EnterTqBoolExpr(ctx *TqBoolExprContext) {}

// ExitTqBoolExpr is called when production tqBoolExpr is exited.
func (s *BaseTqParserListener) ExitTqBoolExpr(ctx *TqBoolExprContext) {}

// EnterTqNumberExpr is called when production tqNumberExpr is entered.
func (s *BaseTqParserListener) EnterTqNumberExpr(ctx *TqNumberExprContext) {}

// ExitTqNumberExpr is called when production tqNumberExpr is exited.
func (s *BaseTqParserListener) ExitTqNumberExpr(ctx *TqNumberExprContext) {}

// EnterTqPercentExpr is called when production tqPercentExpr is entered.
func (s *BaseTqParserListener) EnterTqPercentExpr(ctx *TqPercentExprContext) {}

// ExitTqPercentExpr is called when production tqPercentExpr is exited.
func (s *BaseTqParserListener) ExitTqPercentExpr(ctx *TqPercentExprContext) {}

// EnterTqConcatExpr is called when production tqConcatExpr is entered.
func (s *BaseTqParserListener) EnterTqConcatExpr(ctx *TqConcatExprContext) {}

// ExitTqConcatExpr is called when production tqConcatExpr is exited.
func (s *BaseTqParserListener) ExitTqConcatExpr(ctx *TqConcatExprContext) {}

// EnterTqParenExpr is called when production tqParenExpr is entered.
func (s *BaseTqParserListener) EnterTqParenExpr(ctx *TqParenExprContext) {}

// ExitTqParenExpr is called when production tqParenExpr is exited.
func (s *BaseTqParserListener) ExitTqParenExpr(ctx *TqParenExprContext) {}

// EnterTqUnaryExpr is called when production tqUnaryExpr is entered.
func (s *BaseTqParserListener) EnterTqUnaryExpr(ctx *TqUnaryExprContext) {}

// ExitTqUnaryExpr is called when production tqUnaryExpr is exited.
func (s *BaseTqParserListener) ExitTqUnaryExpr(ctx *TqUnaryExprContext) {}

// EnterTqCallExpr is called when production tqCallExpr is entered.
func (s *BaseTqParserListener) EnterTqCallExpr(ctx *TqCallExprContext) {}

// ExitTqCallExpr is called when production tqCallExpr is exited.
func (s *BaseTqParserListener) ExitTqCallExpr(ctx *TqCallExprContext) {}

// EnterTqRefExpr is called when production tqRefExpr is entered.
func (s *BaseTqParserListener) EnterTqRefExpr(ctx *TqRefExprContext) {}

// ExitTqRefExpr is called when production tqRefExpr is exited.
func (s *BaseTqParserListener) ExitTqRefExpr(ctx *TqRefExprContext) {}

// EnterTqColumnExpr is called when production tqColumnExpr is entered.
func (s *BaseTqParserListener) EnterTqColumnExpr(ctx *TqColumnExprContext) {}

// ExitTqColumnExpr is called when production tqColumnExpr is exited.
func (s *BaseTqParserListener) ExitTqColumnExpr(ctx *TqColumnExprContext) {}

// EnterTqStringExpr is called when production tqStringExpr is entered.
func (s *BaseTqParserListener) EnterTqStringExpr(ctx *TqStringExprContext) {}

// ExitTqStringExpr is called when production tqStringExpr is exited.
func (s *BaseTqParserListener) ExitTqStringExpr(ctx *TqStringExprContext) {}

// EnterTqAddExpr is called when production tqAddExpr is entered.
func (s *BaseTqParserListener) EnterTqAddExpr(ctx *TqAddExprContext) {}

// ExitTqAddExpr is called when production tqAddExpr is exited.
func (s *BaseTqParserListener) ExitTqAddExpr(ctx *TqAddExprContext) {}

// EnterTqErrorExpr is called when production tqErrorExpr is entered.
func (s *BaseTqParserListener) EnterTqErrorExpr(ctx *TqErrorExprContext) {}

// ExitTqErrorExpr is called when production tqErrorExpr is exited.
func (s *BaseTqParserListener) ExitTqErrorExpr(ctx *TqErrorExprContext) {}

// EnterTqMulExpr is called when production tqMulExpr is entered.
func (s *BaseTqParserListener) EnterTqMulExpr(ctx *TqMulExprContext) {}

// ExitTqMulExpr is called when production tqMulExpr is exited.
func (s *BaseTqParserListener) ExitTqMulExpr(ctx *TqMulExprContext) {}

// EnterTqFunctionCall is called when production tqFunctionCall is entered.
func (s *BaseTqParserListener) EnterTqFunctionCall(ctx *TqFunctionCallContext) {}

// ExitTqFunctionCall is called when production tqFunctionCall is exited.
func (s *BaseTqParserListener) ExitTqFunctionCall(ctx *TqFunctionCallContext) {}

// EnterTqArgList is called when production tqArgList is entered.
func (s *BaseTqParserListener) EnterTqArgList(ctx *TqArgListContext) {}

// ExitTqArgList is called when production tqArgList is exited.
func (s *BaseTqParserListener) ExitTqArgList(ctx *TqArgListContext) {}

// EnterErrorExpr is called when production errorExpr is entered.
func (s *BaseTqParserListener) EnterErrorExpr(ctx *ErrorExprContext) {}

// ExitErrorExpr is called when production errorExpr is exited.
func (s *BaseTqParserListener) ExitErrorExpr(ctx *ErrorExprContext) {}

// EnterPipeExpr is called when production pipeExpr is entered.
func (s *BaseTqParserListener) EnterPipeExpr(ctx *PipeExprContext) {}

// ExitPipeExpr is called when production pipeExpr is exited.
func (s *BaseTqParserListener) ExitPipeExpr(ctx *PipeExprContext) {}

// EnterNumberExpr is called when production numberExpr is entered.
func (s *BaseTqParserListener) EnterNumberExpr(ctx *NumberExprContext) {}

// ExitNumberExpr is called when production numberExpr is exited.
func (s *BaseTqParserListener) ExitNumberExpr(ctx *NumberExprContext) {}

// EnterParenExpr is called when production parenExpr is entered.
func (s *BaseTqParserListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production parenExpr is exited.
func (s *BaseTqParserListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterConcatExpr is called when production concatExpr is entered.
func (s *BaseTqParserListener) EnterConcatExpr(ctx *ConcatExprContext) {}

// ExitConcatExpr is called when production concatExpr is exited.
func (s *BaseTqParserListener) ExitConcatExpr(ctx *ConcatExprContext) {}

// EnterStringExpr is called when production stringExpr is entered.
func (s *BaseTqParserListener) EnterStringExpr(ctx *StringExprContext) {}

// ExitStringExpr is called when production stringExpr is exited.
func (s *BaseTqParserListener) ExitStringExpr(ctx *StringExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseTqParserListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseTqParserListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseTqParserListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseTqParserListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterRefExpr is called when production refExpr is entered.
func (s *BaseTqParserListener) EnterRefExpr(ctx *RefExprContext) {}

// ExitRefExpr is called when production refExpr is exited.
func (s *BaseTqParserListener) ExitRefExpr(ctx *RefExprContext) {}

// EnterMulExpr is called when production mulExpr is entered.
func (s *BaseTqParserListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production mulExpr is exited.
func (s *BaseTqParserListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterPercentExpr is called when production percentExpr is entered.
func (s *BaseTqParserListener) EnterPercentExpr(ctx *PercentExprContext) {}

// ExitPercentExpr is called when production percentExpr is exited.
func (s *BaseTqParserListener) ExitPercentExpr(ctx *PercentExprContext) {}

// EnterCallExpr is called when production callExpr is entered.
func (s *BaseTqParserListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production callExpr is exited.
func (s *BaseTqParserListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterBoolExpr is called when production boolExpr is entered.
func (s *BaseTqParserListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production boolExpr is exited.
func (s *BaseTqParserListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterPowExpr is called when production powExpr is entered.
func (s *BaseTqParserListener) EnterPowExpr(ctx *PowExprContext) {}

// ExitPowExpr is called when production powExpr is exited.
func (s *BaseTqParserListener) ExitPowExpr(ctx *PowExprContext) {}

// EnterCompareExpr is called when production compareExpr is entered.
func (s *BaseTqParserListener) EnterCompareExpr(ctx *CompareExprContext) {}

// ExitCompareExpr is called when production compareExpr is exited.
func (s *BaseTqParserListener) ExitCompareExpr(ctx *CompareExprContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseTqParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseTqParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseTqParserListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseTqParserListener) ExitArgList(ctx *ArgListContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseTqParserListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseTqParserListener) ExitReference(ctx *ReferenceContext) {}

// EnterSheetQualifier is called when production sheetQualifier is entered.
func (s *BaseTqParserListener) EnterSheetQualifier(ctx *SheetQualifierContext) {}

// ExitSheetQualifier is called when production sheetQualifier is exited.
func (s *BaseTqParserListener) ExitSheetQualifier(ctx *SheetQualifierContext) {}

// EnterCellRef is called when production cellRef is entered.
func (s *BaseTqParserListener) EnterCellRef(ctx *CellRefContext) {}

// ExitCellRef is called when production cellRef is exited.
func (s *BaseTqParserListener) ExitCellRef(ctx *CellRefContext) {}
