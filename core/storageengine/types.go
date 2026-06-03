package storageengine

import (
	"github.com/hanami/tidets/core/storageengine/model"
	"github.com/hanami/tidets/core/storageengine/tools"
)

// 对外类型别名，调用方只需 import storageengine。
type (
	SeriesKey = model.SeriesKey
	Point     = model.Point
	Value     = model.Value
	Stats     = tools.Stats
)

// DoublePoint 构造 DOUBLE 测点（测试与示例常用）。
func DoublePoint(ts int64, v float64) Point {
	return Point{Timestamp: ts, Value: model.NewDouble(v)}
}
