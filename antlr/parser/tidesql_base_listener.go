// Code generated from TideSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TideSQL
import "github.com/antlr4-go/antlr/v4"

// BaseTideSQLListener is a complete listener for a parse tree produced by TideSQLParser.
type BaseTideSQLListener struct{}

var _ TideSQLListener = &BaseTideSQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTideSQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTideSQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTideSQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTideSQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseTideSQLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseTideSQLListener) ExitStatement(ctx *StatementContext) {}

// EnterInsertStmt is called when production insertStmt is entered.
func (s *BaseTideSQLListener) EnterInsertStmt(ctx *InsertStmtContext) {}

// ExitInsertStmt is called when production insertStmt is exited.
func (s *BaseTideSQLListener) ExitInsertStmt(ctx *InsertStmtContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *BaseTideSQLListener) EnterSelectStmt(ctx *SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *BaseTideSQLListener) ExitSelectStmt(ctx *SelectStmtContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseTideSQLListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseTideSQLListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterTimePredicate is called when production timePredicate is entered.
func (s *BaseTideSQLListener) EnterTimePredicate(ctx *TimePredicateContext) {}

// ExitTimePredicate is called when production timePredicate is exited.
func (s *BaseTideSQLListener) ExitTimePredicate(ctx *TimePredicateContext) {}

// EnterCmpOp is called when production cmpOp is entered.
func (s *BaseTideSQLListener) EnterCmpOp(ctx *CmpOpContext) {}

// ExitCmpOp is called when production cmpOp is exited.
func (s *BaseTideSQLListener) ExitCmpOp(ctx *CmpOpContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseTideSQLListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseTideSQLListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterPath is called when production path is entered.
func (s *BaseTideSQLListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseTideSQLListener) ExitPath(ctx *PathContext) {}

// EnterMeasurement is called when production measurement is entered.
func (s *BaseTideSQLListener) EnterMeasurement(ctx *MeasurementContext) {}

// ExitMeasurement is called when production measurement is exited.
func (s *BaseTideSQLListener) ExitMeasurement(ctx *MeasurementContext) {}

// EnterTimestamp is called when production timestamp is entered.
func (s *BaseTideSQLListener) EnterTimestamp(ctx *TimestampContext) {}

// ExitTimestamp is called when production timestamp is exited.
func (s *BaseTideSQLListener) ExitTimestamp(ctx *TimestampContext) {}

// EnterValue is called when production value is entered.
func (s *BaseTideSQLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseTideSQLListener) ExitValue(ctx *ValueContext) {}
