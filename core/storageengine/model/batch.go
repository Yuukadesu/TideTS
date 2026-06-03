package model

// BatchRecord 一批写入中的一条。
type BatchRecord struct {
	Key   SeriesKey
	Point Point
}
