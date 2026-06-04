package executor

import (
	"context"
	"strconv"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/queryengine/backend"
	"github.com/hanami/tidets/core/queryengine/plan"
	"github.com/hanami/tidets/core/queryengine/result"
	"github.com/hanami/tidets/core/tsmodel"
)

const (
	DefaultQueryLimit = 10000

	colDevice          = "Device"
	colTimeseriesCount = "TimeseriesCount"
	colTimeseries      = "Timeseries"
	colDataType        = "DataType"
)

// Executor 执行物理计划。
type Executor struct {
	Store        backend.Backend
	Catalog      backend.CatalogBackend
	DefaultLimit int
}

func (e *Executor) Execute(ctx context.Context, p plan.Plan) (*result.Result, error) {
	switch p.Kind {
	case plan.KindInsert:
		return e.execInsert(ctx, p.Insert)
	case plan.KindSelect:
		return e.execSelect(ctx, p.Select)
	case plan.KindDelete:
		return e.execDelete(ctx, p.Delete)
	case plan.KindCreateTimeseries:
		return e.execCreateTimeseries(ctx, p.CreateTimeseries)
	case plan.KindShowDevices:
		return e.execShowDevices(ctx, p.ShowDevices)
	case plan.KindShowTimeseries:
		return e.execShowTimeseries(ctx, p.ShowTimeseries)
	default:
		return nil, commons.ErrSQLUnsupportedStmt
	}
}

func (e *Executor) execInsert(ctx context.Context, ins *plan.Insert) (*result.Result, error) {
	if ins == nil || len(ins.Points) == 0 {
		return nil, commons.ErrSQLUnsupportedStmt
	}
	if len(ins.Points) == 1 {
		if err := e.Store.Insert(ctx, ins.Key, ins.Points[0]); err != nil {
			return nil, err
		}
		return &result.Result{Kind: result.KindInsert, AffectedRows: 1}, nil
	}
	if err := e.Store.InsertBatch(ctx, ins.Key, ins.Points); err != nil {
		return nil, err
	}
	return &result.Result{Kind: result.KindInsert, AffectedRows: len(ins.Points)}, nil
}

func (e *Executor) execSelect(ctx context.Context, sel *plan.Select) (*result.Result, error) {
	if sel == nil {
		return nil, commons.ErrSQLUnsupportedStmt
	}
	if sel.Aggregate == plan.SelectCount {
		n, err := e.Store.Count(ctx, sel.Key, sel.Start, sel.End)
		if err != nil {
			return nil, err
		}
		return &result.Result{
			Kind: result.KindSelect,
			Rows: []result.Row{{
				Timestamp: 0,
				Value:     tsmodel.NewInt64(int64(n)),
			}},
		}, nil
	}
	limit := sel.Limit
	if limit <= 0 {
		limit = e.DefaultLimit
	}
	if limit <= 0 {
		limit = DefaultQueryLimit
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

func (e *Executor) execDelete(ctx context.Context, del *plan.Delete) (*result.Result, error) {
	if del == nil {
		return nil, commons.ErrSQLUnsupportedStmt
	}
	n, err := e.Store.DeleteRange(ctx, del.Key, del.Start, del.End)
	if err != nil {
		return nil, err
	}
	return &result.Result{Kind: result.KindDelete, AffectedRows: n}, nil
}

func (e *Executor) execCreateTimeseries(ctx context.Context, create *plan.CreateTimeseries) (*result.Result, error) {
	if create == nil {
		return nil, commons.ErrSQLUnsupportedStmt
	}
	if e.Catalog == nil {
		return nil, commons.ErrSQLUnsupportedStmt
	}
	if err := e.Catalog.CreateTimeseries(ctx, create.DevicePath, create.Measurement, create.DataType); err != nil {
		return nil, err
	}
	return &result.Result{Kind: result.KindCreateTimeseries, AffectedRows: 1}, nil
}

func (e *Executor) execShowDevices(ctx context.Context, show *plan.ShowDevices) (*result.Result, error) {
	if show == nil {
		return nil, commons.ErrSQLUnsupportedStmt
	}
	if e.Catalog == nil {
		return nil, commons.ErrSQLUnsupportedStmt
	}
	devices, err := e.Catalog.ListDevices(ctx, show.Pattern)
	if err != nil {
		return nil, err
	}
	cols := []string{colDevice, colTimeseriesCount}
	rows := make([]result.CatalogRow, 0, len(devices))
	for _, d := range devices {
		rows = append(rows, result.CatalogRow{
			Columns: map[string]string{
				colDevice:          d.Path,
				colTimeseriesCount: formatInt(d.MeasurementCnt),
			},
		})
	}
	return &result.Result{
		Kind:        result.KindShowDevices,
		CatalogRows: rows,
		ColumnNames: cols,
	}, nil
}

func (e *Executor) execShowTimeseries(ctx context.Context, show *plan.ShowTimeseries) (*result.Result, error) {
	if show == nil {
		return nil, commons.ErrSQLUnsupportedStmt
	}
	if e.Catalog == nil {
		return nil, commons.ErrSQLUnsupportedStmt
	}
	series, err := e.Catalog.ListTimeseries(ctx, show.DevicePath)
	if err != nil {
		return nil, err
	}
	cols := []string{colTimeseries, colDataType}
	rows := make([]result.CatalogRow, 0, len(series))
	for _, ts := range series {
		rows = append(rows, result.CatalogRow{
			Columns: map[string]string{
				colTimeseries: ts.FullPath,
				colDataType:   tsmodel.DataTypeName(ts.DataType),
			},
		})
	}
	return &result.Result{
		Kind:        result.KindShowTimeseries,
		CatalogRows: rows,
		ColumnNames: cols,
	}, nil
}

func formatInt(n int) string {
	return strconv.Itoa(n)
}
