package visitor

import (
	"strconv"
	"strings"

	"github.com/hanami/tidets/antlr/parser"
	commons "github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/sql/ast"
	"github.com/hanami/tidets/core/storageengine/model"
)

// Builder 将 ANTLR 解析树转为 ast.Stmt。
type Builder struct {
	parser.BaseTideSQLVisitor
	err error
}

func (b *Builder) fail(err error) interface{} {
	if b.err == nil {
		b.err = err
	}
	return nil
}

func (b *Builder) Err() error { return b.err }

func (b *Builder) VisitStatement(ctx *parser.StatementContext) interface{} {
	if b.err != nil {
		return nil
	}
	if ctx.InsertStmt() != nil {
		return b.VisitInsertStmt(ctx.InsertStmt().(*parser.InsertStmtContext))
	}
	return b.VisitSelectStmt(ctx.SelectStmt().(*parser.SelectStmtContext))
}

func (b *Builder) VisitInsertStmt(ctx *parser.InsertStmtContext) interface{} {
	if b.err != nil {
		return nil
	}
	device, ok := b.visitPath(ctx.Path())
	if !ok {
		return nil
	}
	measurement, ok := b.visitMeasurement(ctx.Measurement())
	if !ok {
		return nil
	}
	ts, ok := b.visitTimestamp(ctx.Timestamp())
	if !ok {
		return nil
	}
	val, ok := b.visitValue(ctx.Value())
	if !ok {
		return nil
	}
	return &ast.InsertStmt{
		DevicePath:  device,
		Measurement: measurement,
		Timestamp:   ts,
		Value:       val,
	}
}

func (b *Builder) VisitSelectStmt(ctx *parser.SelectStmtContext) interface{} {
	if b.err != nil {
		return nil
	}
	device, ok := b.visitPath(ctx.Path())
	if !ok {
		return nil
	}
	measurement, ok := b.visitMeasurement(ctx.Measurement())
	if !ok {
		return nil
	}
	stmt := &ast.SelectStmt{
		DevicePath:  device,
		Measurement: measurement,
	}
	if wc := ctx.WhereClause(); wc != nil {
		preds, ok := b.visitWhereClause(wc)
		if !ok {
			return nil
		}
		stmt.TimeWhere = preds
	}
	if lc := ctx.LimitClause(); lc != nil {
		lim, ok := b.visitLimitClause(lc)
		if !ok {
			return nil
		}
		stmt.Limit = lim
	}
	return stmt
}

func (b *Builder) visitWhereClause(ctx parser.IWhereClauseContext) ([]ast.TimePredicate, bool) {
	var preds []ast.TimePredicate
	for _, tp := range ctx.AllTimePredicate() {
		p, ok := b.visitTimePredicate(tp)
		if !ok {
			return nil, false
		}
		preds = append(preds, p)
	}
	return preds, true
}

func (b *Builder) visitTimePredicate(ctx parser.ITimePredicateContext) (ast.TimePredicate, bool) {
	op, ok := b.visitCmpOp(ctx.CmpOp())
	if !ok {
		return ast.TimePredicate{}, false
	}
	ts, err := parseInteger(ctx.INTEGER().GetText())
	if err != nil {
		b.fail(err)
		return ast.TimePredicate{}, false
	}
	return ast.TimePredicate{Op: op, Timestamp: ts}, true
}

func (b *Builder) visitCmpOp(ctx parser.ICmpOpContext) (ast.CmpOp, bool) {
	switch {
	case ctx.GTE() != nil:
		return ast.CmpGTE, true
	case ctx.LTE() != nil:
		return ast.CmpLTE, true
	case ctx.GT() != nil:
		return ast.CmpGT, true
	case ctx.LT() != nil:
		return ast.CmpLT, true
	case ctx.EQ() != nil:
		return ast.CmpEQ, true
	default:
		b.fail(commons.ErrSQLInvalidCmpOp)
		return 0, false
	}
}

func (b *Builder) visitLimitClause(ctx parser.ILimitClauseContext) (int, bool) {
	n, err := parseInteger(ctx.INTEGER().GetText())
	if err != nil {
		b.fail(err)
		return 0, false
	}
	if n <= 0 {
		b.fail(commons.ErrSQLLimitInvalid)
		return 0, false
	}
	if n > int64(^uint(0)>>1) {
		b.fail(commons.ErrSQLLimitInvalid)
		return 0, false
	}
	return int(n), true
}

func (b *Builder) visitPath(ctx parser.IPathContext) (string, bool) {
	if ctx == nil {
		b.fail(commons.ErrSQLPathRequired)
		return "", false
	}
	parts := make([]string, 0, len(ctx.AllIDENTIFIER()))
	for _, id := range ctx.AllIDENTIFIER() {
		parts = append(parts, id.GetText())
	}
	if len(parts) == 0 {
		b.fail(commons.ErrSQLPathRequired)
		return "", false
	}
	return strings.Join(parts, "."), true
}

func (b *Builder) visitMeasurement(ctx parser.IMeasurementContext) (string, bool) {
	if ctx == nil || ctx.IDENTIFIER() == nil {
		b.fail(commons.ErrSQLMeasurementRequired)
		return "", false
	}
	return ctx.IDENTIFIER().GetText(), true
}

func (b *Builder) visitTimestamp(ctx parser.ITimestampContext) (int64, bool) {
	if ctx == nil || ctx.INTEGER() == nil {
		b.fail(commons.ErrSQLTimestampRequired)
		return 0, false
	}
	ts, err := parseInteger(ctx.INTEGER().GetText())
	if err != nil {
		b.fail(err)
		return 0, false
	}
	if ts <= 0 {
		b.fail(commons.ErrSQLTimestampInvalid)
		return 0, false
	}
	return ts, true
}

func (b *Builder) visitValue(ctx parser.IValueContext) (model.Value, bool) {
	if ctx == nil {
		b.fail(commons.ErrSQLValueRequired)
		return model.Value{}, false
	}
	switch {
	case ctx.BOOLEAN() != nil:
		text := strings.ToLower(ctx.BOOLEAN().GetText())
		return model.NewBoolean(text == "true"), true
	case ctx.FLOAT() != nil:
		f, err := strconv.ParseFloat(ctx.FLOAT().GetText(), 64)
		if err != nil {
			b.fail(commons.Wrap("sql", commons.CodeInvalidArgument, "float literal", err))
			return model.Value{}, false
		}
		return model.NewDouble(f), true
	case ctx.STRING() != nil:
		s, err := unquoteString(ctx.STRING().GetText())
		if err != nil {
			b.fail(err)
			return model.Value{}, false
		}
		return model.NewText(s), true
	case ctx.INTEGER() != nil:
		n, err := parseInteger(ctx.INTEGER().GetText())
		if err != nil {
			b.fail(err)
			return model.Value{}, false
		}
		return model.NewInt64(n), true
	default:
		b.fail(commons.ErrSQLValueRequired)
		return model.Value{}, false
	}
}

func parseInteger(text string) (int64, error) {
	n, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return 0, commons.Wrap("sql", commons.CodeInvalidArgument, "integer literal", err)
	}
	return n, nil
}

func unquoteString(lit string) (string, error) {
	if len(lit) < 2 || lit[0] != '\'' || lit[len(lit)-1] != '\'' {
		return "", commons.ErrSQLStringLiteral
	}
	var b strings.Builder
	for i := 1; i < len(lit)-1; i++ {
		if lit[i] == '\\' && i+1 < len(lit)-1 {
			i++
			b.WriteByte(lit[i])
			continue
		}
		b.WriteByte(lit[i])
	}
	return b.String(), nil
}

// Build 从 Statement 根节点构建 AST。
func Build(tree *parser.StatementContext) (ast.Stmt, error) {
	v := &Builder{}
	out := tree.Accept(v)
	if v.err != nil {
		return nil, v.err
	}
	stmt, ok := out.(ast.Stmt)
	if !ok || stmt == nil {
		return nil, commons.ErrSQLParseFailed
	}
	return stmt, nil
}
