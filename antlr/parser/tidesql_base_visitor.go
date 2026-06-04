// Code generated from TideSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TideSQL
import "github.com/antlr4-go/antlr/v4"

type BaseTideSQLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTideSQLVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitInsertStmt(ctx *InsertStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitValueRow(ctx *ValueRowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitSelectStmt(ctx *SelectStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitDeleteStmt(ctx *DeleteStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitCreateTimeseriesStmt(ctx *CreateTimeseriesStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitShowDevicesStmt(ctx *ShowDevicesStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitShowTimeseriesStmt(ctx *ShowTimeseriesStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitShowPattern(ctx *ShowPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitTimePredicate(ctx *TimePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitCmpOp(ctx *CmpOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitPath(ctx *PathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitMeasurement(ctx *MeasurementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitDataTypeName(ctx *DataTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitTimestamp(ctx *TimestampContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTideSQLVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}
