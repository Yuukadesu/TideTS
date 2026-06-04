package plan

import (
	"math"
	"strings"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/sql/ast"
	"github.com/hanami/tidets/core/tsmodel"
)

// Build 将 AST 转为执行计划。
func Build(stmt ast.Stmt) (Plan, error) {
	switch s := stmt.(type) {
	case *ast.InsertStmt:
		return buildInsert(s)
	case *ast.SelectStmt:
		return buildSelect(s)
	case *ast.CreateTimeseriesStmt:
		return buildCreateTimeseries(s)
	case *ast.ShowDevicesStmt:
		return buildShowDevices(s)
	case *ast.ShowTimeseriesStmt:
		return buildShowTimeseries(s)
	case *ast.DeleteStmt:
		return buildDelete(s)
	default:
		return Plan{}, commons.ErrSQLUnsupportedStmt
	}
}

func buildInsert(s *ast.InsertStmt) (Plan, error) {
	if s.DevicePath == "" || s.Measurement == "" {
		return Plan{}, commons.ErrSQLDeviceMeasurementRequired
	}
	if len(s.Rows) == 0 {
		return Plan{}, commons.ErrSQLValueRequired
	}
	points := make([]tsmodel.Point, 0, len(s.Rows))
	for _, row := range s.Rows {
		if row.Timestamp <= 0 {
			return Plan{}, commons.ErrSQLTimestampInvalid
		}
		if err := row.Value.Validate(); err != nil {
			return Plan{}, err
		}
		points = append(points, tsmodel.Point{
			Timestamp: row.Timestamp,
			Value:     row.Value,
		})
	}
	return Plan{
		Kind: KindInsert,
		Insert: &Insert{
			Key: tsmodel.SeriesKey{
				DevicePath:  s.DevicePath,
				Measurement: s.Measurement,
			},
			Points: points,
		},
	}, nil
}

func buildSelect(s *ast.SelectStmt) (Plan, error) {
	if s.DevicePath == "" || s.Measurement == "" {
		return Plan{}, commons.ErrSQLDeviceMeasurementRequired
	}
	start, end, err := timeRangeFromWhere(s.TimeWhere)
	if err != nil {
		return Plan{}, err
	}
	return Plan{
		Kind: KindSelect,
		Select: &Select{
			Key: tsmodel.SeriesKey{
				DevicePath:  s.DevicePath,
				Measurement: s.Measurement,
			},
			Start:     start,
			End:       end,
			Limit:     s.Limit,
			Aggregate: mapSelectAgg(s.Aggregate),
		},
	}, nil
}

func mapSelectAgg(a ast.SelectAgg) SelectAgg {
	if a == ast.SelectCount {
		return SelectCount
	}
	return SelectRaw
}

func buildDelete(s *ast.DeleteStmt) (Plan, error) {
	if s.DevicePath == "" || s.Measurement == "" {
		return Plan{}, commons.ErrSQLDeviceMeasurementRequired
	}
	if len(s.TimeWhere) == 0 {
		return Plan{}, commons.ErrSQLWhereRequired
	}
	start, end, err := timeRangeFromWhere(s.TimeWhere)
	if err != nil {
		return Plan{}, err
	}
	return Plan{
		Kind: KindDelete,
		Delete: &Delete{
			Key: tsmodel.SeriesKey{
				DevicePath:  s.DevicePath,
				Measurement: s.Measurement,
			},
			Start: start,
			End:   end,
		},
	}, nil
}

func buildCreateTimeseries(s *ast.CreateTimeseriesStmt) (Plan, error) {
	if err := validateDevicePath(s.DevicePath); err != nil {
		return Plan{}, err
	}
	if s.Measurement == "" {
		return Plan{}, commons.ErrSQLMeasurementRequired
	}
	if s.DataType == tsmodel.DataTypeUnknown {
		return Plan{}, commons.ErrSQLDataTypeInvalid
	}
	return Plan{
		Kind: KindCreateTimeseries,
		CreateTimeseries: &CreateTimeseries{
			DevicePath:  s.DevicePath,
			Measurement: s.Measurement,
			DataType:    s.DataType,
		},
	}, nil
}

func buildShowDevices(s *ast.ShowDevicesStmt) (Plan, error) {
	if s.Pattern != "" && !strings.Contains(s.Pattern, "*") {
		if err := validateDevicePath(s.Pattern); err != nil {
			return Plan{}, err
		}
	}
	return Plan{
		Kind: KindShowDevices,
		ShowDevices: &ShowDevices{
			Pattern: s.Pattern,
		},
	}, nil
}

func buildShowTimeseries(s *ast.ShowTimeseriesStmt) (Plan, error) {
	if err := validateDevicePath(s.DevicePath); err != nil {
		return Plan{}, err
	}
	return Plan{
		Kind: KindShowTimeseries,
		ShowTimeseries: &ShowTimeseries{
			DevicePath: s.DevicePath,
		},
	}, nil
}

func validateDevicePath(path string) error {
	if path == "" {
		return commons.ErrSQLPathRequired
	}
	if !strings.HasPrefix(path, "root") {
		return commons.ErrMetadataInvalidPath
	}
	return nil
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
