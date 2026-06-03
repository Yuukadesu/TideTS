package session

import "github.com/hanami/tidets/core/storageengine/model"

// Value 测点值（与存储层类型一致）。
type Value = model.Value

// DataType 值类型枚举。
type DataType = model.DataType

const (
	DataTypeBoolean = model.DataTypeBoolean
	DataTypeInt32   = model.DataTypeInt32
	DataTypeInt64   = model.DataTypeInt64
	DataTypeFloat   = model.DataTypeFloat
	DataTypeDouble  = model.DataTypeDouble
	DataTypeText    = model.DataTypeText
)

func Boolean(v bool) Value   { return model.NewBoolean(v) }
func Int32(v int32) Value    { return model.NewInt32(v) }
func Int64(v int64) Value    { return model.NewInt64(v) }
func Float(v float32) Value  { return model.NewFloat(v) }
func Double(v float64) Value { return model.NewDouble(v) }
func Text(v string) Value    { return model.NewText(v) }
