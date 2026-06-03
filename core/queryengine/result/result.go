package result

import "github.com/hanami/tidets/core/storageengine"

// Kind SQL 执行结果类型。
type Kind int

const (
	KindInsert Kind = iota
	KindSelect
)

// Result SQL 执行结果。
type Result struct {
	Kind         Kind
	AffectedRows int
	Rows         []Row
}

// Row SELECT 返回的一行。
type Row struct {
	Timestamp int64
	Value     storageengine.Value
}
