package plan

import "github.com/hanami/tidets/core/storageengine"

// Kind 执行计划类型。
type Kind int

const (
	KindInsert Kind = iota
	KindSelect
)

// Plan 物理计划（单机最简）。
type Plan struct {
	Kind Kind

	Insert *Insert
	Select *Select
}

type Insert struct {
	Key   storageengine.SeriesKey
	Point storageengine.Point
}

type Select struct {
	Key   storageengine.SeriesKey
	Start int64
	End   int64
	Limit int
}
