package session

import "github.com/hanami/tidets/core/tsmodel"

// Value 客户端写入/读取的测点值（与存储层 tsmodel 对齐）。
type Value = tsmodel.Value

// DataType 值类型。
type DataType = tsmodel.DataType

const (
	DataTypeBoolean = tsmodel.DataTypeBoolean
	DataTypeInt32   = tsmodel.DataTypeInt32
	DataTypeInt64   = tsmodel.DataTypeInt64
	DataTypeFloat   = tsmodel.DataTypeFloat
	DataTypeDouble  = tsmodel.DataTypeDouble
	DataTypeText    = tsmodel.DataTypeText
)

func Boolean(v bool) Value   { return tsmodel.NewBoolean(v) }
func Int32(v int32) Value    { return tsmodel.NewInt32(v) }
func Int64(v int64) Value    { return tsmodel.NewInt64(v) }
func Float(v float32) Value  { return tsmodel.NewFloat(v) }
func Double(v float64) Value { return tsmodel.NewDouble(v) }
func Text(v string) Value    { return tsmodel.NewText(v) }
