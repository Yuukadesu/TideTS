package result

import "github.com/hanami/tidets/core/tsmodel"

// Kind SQL 执行结果类型。
type Kind int

const (
	KindInsert Kind = iota
	KindSelect
	KindCreateTimeseries
	KindShowDevices
	KindShowTimeseries
	KindDelete
)

// Result SQL 执行结果。
type Result struct {
	Kind         Kind
	AffectedRows int
	Rows         []Row
	CatalogRows  []CatalogRow
	ColumnNames  []string
}

// Row SELECT 返回的一行。
type Row struct {
	Timestamp int64
	Value     tsmodel.Value
}

// CatalogRow SHOW 返回的一行（列名 → 值）。
type CatalogRow struct {
	Columns map[string]string
}
