package format

import (
	"fmt"
	"strings"

	"github.com/hanami/tidets/client/session"
	"github.com/hanami/tidets/core/tsmodel"
)

// SQLResult 将 SQL 执行结果格式化为终端输出。
func SQLResult(res *session.SQLResult) string {
	if res == nil {
		return ""
	}
	if len(res.CatalogRows) > 0 {
		return formatCatalogRows(res.ColumnNames, res.CatalogRows)
	}
	if len(res.Rows) > 0 {
		return formatSelectRows(res.Rows)
	}
	if res.AffectedRows > 0 {
		return fmt.Sprintf("OK, %d row affected", res.AffectedRows)
	}
	return "OK"
}

func formatCatalogRows(cols []string, rows []session.CatalogRow) string {
	if len(cols) == 0 && len(rows) > 0 {
		for k := range rows[0].Columns {
			cols = append(cols, k)
		}
	}
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%d row(s)\n", len(rows)))
	if len(cols) == 0 {
		return strings.TrimRight(b.String(), "\n")
	}
	widths := make([]int, len(cols))
	for i, c := range cols {
		widths[i] = len(c)
	}
	values := make([][]string, len(rows))
	for i, row := range rows {
		values[i] = make([]string, len(cols))
		for j, c := range cols {
			v := row.Columns[c]
			values[i][j] = v
			if len(v) > widths[j] {
				widths[j] = len(v)
			}
		}
	}
	header := formatTableRow(cols, widths)
	b.WriteString(header)
	b.WriteString("\n")
	b.WriteString(strings.Repeat("-", len(header)))
	b.WriteString("\n")
	for _, row := range values {
		b.WriteString(formatTableRow(row, widths))
		b.WriteString("\n")
	}
	return strings.TrimRight(b.String(), "\n")
}

func formatTableRow(cells []string, widths []int) string {
	parts := make([]string, len(cells))
	for i, c := range cells {
		parts[i] = fmt.Sprintf("%-*s", widths[i], c)
	}
	return strings.Join(parts, "  ")
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

func formatValue(v tsmodel.Value) string {
	switch v.Type {
	case tsmodel.DataTypeBoolean:
		return fmt.Sprintf("%v", v.Boolean)
	case tsmodel.DataTypeInt32:
		return fmt.Sprintf("%d", v.Int32)
	case tsmodel.DataTypeInt64:
		return fmt.Sprintf("%d", v.Int64)
	case tsmodel.DataTypeFloat:
		return fmt.Sprintf("%g", v.Float)
	case tsmodel.DataTypeDouble:
		return fmt.Sprintf("%g", v.Double)
	case tsmodel.DataTypeText:
		return fmt.Sprintf("'%s'", v.Text)
	default:
		return "?"
	}
}
