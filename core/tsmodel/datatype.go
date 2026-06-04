package tsmodel

import (
	"strings"

	"github.com/hanami/tidets/commons/errors"
)

// ParseDataType 解析 SQL / DDL 中的类型名（大小写不敏感）。
func ParseDataType(name string) (DataType, error) {
	switch strings.ToUpper(strings.TrimSpace(name)) {
	case "BOOLEAN":
		return DataTypeBoolean, nil
	case "INT32":
		return DataTypeInt32, nil
	case "INT64":
		return DataTypeInt64, nil
	case "FLOAT":
		return DataTypeFloat, nil
	case "DOUBLE":
		return DataTypeDouble, nil
	case "TEXT":
		return DataTypeText, nil
	default:
		return DataTypeUnknown, commons.ErrSQLDataTypeInvalid
	}
}

// DataTypeName 返回类型的大写名称（SHOW 结果用）。
func DataTypeName(dt DataType) string {
	switch dt {
	case DataTypeBoolean:
		return "BOOLEAN"
	case DataTypeInt32:
		return "INT32"
	case DataTypeInt64:
		return "INT64"
	case DataTypeFloat:
		return "FLOAT"
	case DataTypeDouble:
		return "DOUBLE"
	case DataTypeText:
		return "TEXT"
	default:
		return "UNKNOWN"
	}
}
