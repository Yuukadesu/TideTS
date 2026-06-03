package memtable

import (
	"github.com/hanami/tidets/core/storageengine/model"
	"github.com/hanami/tidets/core/storageengine/utils"
)

// MemTable 内存表：最新写入与 WAL 回放数据（按序列时间有序）。
type MemTable struct {
	series map[string][]model.Point
}

func New() *MemTable {
	return &MemTable{series: make(map[string][]model.Point)}
}

func (m *MemTable) Insert(key model.SeriesKey, p model.Point) error {
	k := key.String()
	existing := m.series[k]
	if err := utils.CheckSeriesValueType(existing, p); err != nil {
		return err
	}
	m.series[k] = utils.InsertSorted(existing, p)
	return nil
}

func (m *MemTable) Query(key model.SeriesKey, start, end int64) []model.Point {
	return utils.QueryPointsInRange(m.series[key.String()], start, end)
}

func (m *MemTable) IsEmpty() bool {
	return len(m.series) == 0
}

func (m *MemTable) PointCount() int {
	n := 0
	for _, pts := range m.series {
		n += len(pts)
	}
	return n
}

func (m *MemTable) Snapshot() map[string][]model.Point {
	out := make(map[string][]model.Point, len(m.series))
	for k, pts := range m.series {
		out[k] = append([]model.Point(nil), pts...)
	}
	return out
}

func (m *MemTable) Reset() {
	m.series = make(map[string][]model.Point)
}
