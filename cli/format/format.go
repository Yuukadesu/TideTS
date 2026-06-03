package format

import (
	"fmt"
	"strings"

	"github.com/hanami/tidets/client/session"
	"github.com/hanami/tidets/core/storageengine/model"
)

// SQLResult 将 SQL 执行结果格式化为终端输出。
func SQLResult(res *session.SQLResult) string {
	if res == nil {
		return ""
	}
	if len(res.Rows) > 0 {
		return formatSelectRows(res.Rows)
	}
	if res.AffectedRows > 0 {
		return fmt.Sprintf("OK, %d row affected", res.AffectedRows)
	}
	return "OK"
}

func formatSelectRows(rows []session.Point) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%d row(s)\n", len(rows)))
	b.WriteString("timestamp          value\n")
	b.WriteString("------------------  ----------------\n")
	for _, row := range rows {
		b.WriteString(fmt.Sprintf("%-18d  %s\n", row.Timestamp, formatValue(row.Value)))
	}
	return strings.TrimRight(b.String(), "\n")
}

func formatValue(v model.Value) string {
	switch v.Type {
	case model.DataTypeBoolean:
		return fmt.Sprintf("%v", v.Boolean)
	case model.DataTypeInt32:
		return fmt.Sprintf("%d", v.Int32)
	case model.DataTypeInt64:
		return fmt.Sprintf("%d", v.Int64)
	case model.DataTypeFloat:
		return fmt.Sprintf("%g", v.Float)
	case model.DataTypeDouble:
		return fmt.Sprintf("%g", v.Double)
	case model.DataTypeText:
		return fmt.Sprintf("'%s'", v.Text)
	default:
		return "?"
	}
}
