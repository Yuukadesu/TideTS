package session

// SQLResult ExecuteSQL 的返回（INSERT / SELECT）。
type SQLResult struct {
	AffectedRows int     `json:"affectedRows,omitempty"`
	Rows         []Point `json:"rows,omitempty"`
}

// IsSelect 是否为查询结果。
func (r *SQLResult) IsSelect() bool {
	return r != nil && len(r.Rows) > 0
}

// IsInsert 是否为写入结果。
func (r *SQLResult) IsInsert() bool {
	return r != nil && r.AffectedRows > 0 && len(r.Rows) == 0
}
