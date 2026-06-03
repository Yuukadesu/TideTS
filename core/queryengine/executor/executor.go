package executor

import (
	"context"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/queryengine/plan"
	"github.com/hanami/tidets/core/queryengine/result"
	querystore "github.com/hanami/tidets/core/queryengine/storage"
)

const defaultSelectLimit = 10000

// Executor 执行物理计划。
type Executor struct {
	Store        querystore.Backend
	DefaultLimit int
}

func (e *Executor) Execute(ctx context.Context, p plan.Plan) (*result.Result, error) {
	switch p.Kind {
	case plan.KindInsert:
		return e.execInsert(ctx, p.Insert)
	case plan.KindSelect:
		return e.execSelect(ctx, p.Select)
	default:
		return nil, commons.ErrSQLUnsupportedStmt
	}
}

func (e *Executor) execInsert(ctx context.Context, ins *plan.Insert) (*result.Result, error) {
	if ins == nil {
		return nil, commons.ErrSQLUnsupportedStmt
	}
	if err := e.Store.Insert(ctx, ins.Key, ins.Point); err != nil {
		return nil, err
	}
	return &result.Result{Kind: result.KindInsert, AffectedRows: 1}, nil
}

func (e *Executor) execSelect(ctx context.Context, sel *plan.Select) (*result.Result, error) {
	if sel == nil {
		return nil, commons.ErrSQLUnsupportedStmt
	}
	limit := sel.Limit
	if limit <= 0 {
		limit = e.DefaultLimit
	}
	if limit <= 0 {
		limit = defaultSelectLimit
	}
	points, err := e.Store.Query(ctx, sel.Key, sel.Start, sel.End, limit)
	if err != nil {
		return nil, err
	}
	rows := make([]result.Row, 0, len(points))
	for _, p := range points {
		rows = append(rows, result.Row{Timestamp: p.Timestamp, Value: p.Value})
	}
	return &result.Result{Kind: result.KindSelect, Rows: rows}, nil
}
