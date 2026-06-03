// Code generated from TideSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TideSQL
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TideSQLParser.
type TideSQLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TideSQLParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by TideSQLParser#insertStmt.
	VisitInsertStmt(ctx *InsertStmtContext) interface{}

	// Visit a parse tree produced by TideSQLParser#selectStmt.
	VisitSelectStmt(ctx *SelectStmtContext) interface{}

	// Visit a parse tree produced by TideSQLParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by TideSQLParser#timePredicate.
	VisitTimePredicate(ctx *TimePredicateContext) interface{}

	// Visit a parse tree produced by TideSQLParser#cmpOp.
	VisitCmpOp(ctx *CmpOpContext) interface{}

	// Visit a parse tree produced by TideSQLParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by TideSQLParser#path.
	VisitPath(ctx *PathContext) interface{}

	// Visit a parse tree produced by TideSQLParser#measurement.
	VisitMeasurement(ctx *MeasurementContext) interface{}

	// Visit a parse tree produced by TideSQLParser#timestamp.
	VisitTimestamp(ctx *TimestampContext) interface{}

	// Visit a parse tree produced by TideSQLParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
