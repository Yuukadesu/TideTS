package session

// SQLResult ExecuteSQL 的返回（INSERT / SELECT / CREATE / SHOW）。
type SQLResult struct {
	AffectedRows int          `json:"affectedRows,omitempty"`
	Rows         []Point      `json:"rows,omitempty"`
	CatalogRows  []CatalogRow `json:"catalogRows,omitempty"`
	ColumnNames  []string     `json:"columnNames,omitempty"`
}

// CatalogRow SHOW 语句返回的一行。
type CatalogRow struct {
	Columns map[string]string `json:"columns,omitempty"`
}

// IsSelect 是否为时序查询结果。
func (r *SQLResult) IsSelect() bool {
	return r != nil && len(r.Rows) > 0
}

// IsInsert 是否为写入结果。
func (r *SQLResult) IsInsert() bool {
	return r != nil && r.AffectedRows > 0 && len(r.Rows) == 0 && len(r.CatalogRows) == 0
}

// IsShow 是否为 catalog 查询结果。
func (r *SQLResult) IsShow() bool {
	return r != nil && len(r.CatalogRows) > 0
}
