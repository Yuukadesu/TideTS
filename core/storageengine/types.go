package storageengine

import (
	"github.com/hanami/tidets/core/storageengine/tools"
	"github.com/hanami/tidets/core/tsmodel"
)

// 对外类型别名，调用方也可直接 import tsmodel。
type (
	SeriesKey = tsmodel.SeriesKey
	Point     = tsmodel.Point
	Value     = tsmodel.Value
	Stats     = tools.Stats
)

// DoublePoint 构造 DOUBLE 测点（测试与示例常用）。
func DoublePoint(ts int64, v float64) Point {
	return Point{Timestamp: ts, Value: tsmodel.NewDouble(v)}
}
