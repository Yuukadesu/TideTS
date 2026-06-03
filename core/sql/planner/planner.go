package planner

import (
	"math"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/queryengine/plan"
	"github.com/hanami/tidets/core/sql/ast"
	"github.com/hanami/tidets/core/storageengine"
)

// Build 将 AST 转为执行计划。
func Build(stmt ast.Stmt) (plan.Plan, error) {
	switch s := stmt.(type) {
	case *ast.InsertStmt:
		return buildInsert(s)
	case *ast.SelectStmt:
		return buildSelect(s)
	default:
		return plan.Plan{}, commons.ErrSQLUnsupportedStmt
	}
}

func buildInsert(s *ast.InsertStmt) (plan.Plan, error) {
	if s.DevicePath == "" || s.Measurement == "" {
		return plan.Plan{}, commons.ErrSQLDeviceMeasurementRequired
	}
	if s.Timestamp <= 0 {
		return plan.Plan{}, commons.ErrSQLTimestampInvalid
	}
	if err := s.Value.Validate(); err != nil {
		return plan.Plan{}, err
	}
	return plan.Plan{
		Kind: plan.KindInsert,
		Insert: &plan.Insert{
			Key: storageengine.SeriesKey{
				DevicePath:  s.DevicePath,
				Measurement: s.Measurement,
			},
			Point: storageengine.Point{
				Timestamp: s.Timestamp,
				Value:     s.Value,
			},
		},
	}, nil
}

func buildSelect(s *ast.SelectStmt) (plan.Plan, error) {
	if s.DevicePath == "" || s.Measurement == "" {
		return plan.Plan{}, commons.ErrSQLDeviceMeasurementRequired
	}
	start, end, err := timeRangeFromWhere(s.TimeWhere)
	if err != nil {
		return plan.Plan{}, err
	}
	return plan.Plan{
		Kind: plan.KindSelect,
		Select: &plan.Select{
			Key: storageengine.SeriesKey{
				DevicePath:  s.DevicePath,
				Measurement: s.Measurement,
			},
			Start: start,
			End:   end,
			Limit: s.Limit,
		},
	}, nil
}

func timeRangeFromWhere(preds []ast.TimePredicate) (start, end int64, err error) {
	if len(preds) == 0 {
		return 0, math.MaxInt64, nil
	}
	start = math.MinInt64
	end = math.MaxInt64
	for _, p := range preds {
		switch p.Op {
		case ast.CmpGTE:
			if p.Timestamp > start {
				start = p.Timestamp
			}
		case ast.CmpGT:
			if p.Timestamp >= start {
				start = p.Timestamp + 1
			}
		case ast.CmpLTE:
			if p.Timestamp < end {
				end = p.Timestamp
			}
		case ast.CmpLT:
			if p.Timestamp <= end {
				end = p.Timestamp - 1
			}
		case ast.CmpEQ:
			start = p.Timestamp
			end = p.Timestamp
		default:
			return 0, 0, commons.ErrSQLInvalidCmpOp
		}
	}
	if start > end {
		return 0, 0, commons.ErrSQLInvalidTimeRange
	}
	return start, end, nil
}
