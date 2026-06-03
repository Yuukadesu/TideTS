// Code generated from TideSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TideSQL
import "github.com/antlr4-go/antlr/v4"

// TideSQLListener is a complete listener for a parse tree produced by TideSQLParser.
type TideSQLListener interface {
	antlr.ParseTreeListener

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterInsertStmt is called when entering the insertStmt production.
	EnterInsertStmt(c *InsertStmtContext)

	// EnterSelectStmt is called when entering the selectStmt production.
	EnterSelectStmt(c *SelectStmtContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterTimePredicate is called when entering the timePredicate production.
	EnterTimePredicate(c *TimePredicateContext)

	// EnterCmpOp is called when entering the cmpOp production.
	EnterCmpOp(c *CmpOpContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterMeasurement is called when entering the measurement production.
	EnterMeasurement(c *MeasurementContext)

	// EnterTimestamp is called when entering the timestamp production.
	EnterTimestamp(c *TimestampContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitInsertStmt is called when exiting the insertStmt production.
	ExitInsertStmt(c *InsertStmtContext)

	// ExitSelectStmt is called when exiting the selectStmt production.
	ExitSelectStmt(c *SelectStmtContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitTimePredicate is called when exiting the timePredicate production.
	ExitTimePredicate(c *TimePredicateContext)

	// ExitCmpOp is called when exiting the cmpOp production.
	ExitCmpOp(c *CmpOpContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitMeasurement is called when exiting the measurement production.
	ExitMeasurement(c *MeasurementContext)

	// ExitTimestamp is called when exiting the timestamp production.
	ExitTimestamp(c *TimestampContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
