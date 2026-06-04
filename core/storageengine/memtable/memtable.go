package memtable

import (
	"github.com/hanami/tidets/core/storageengine/utils"
	"github.com/hanami/tidets/core/tsmodel"
)

// MemTable 内存表：最新写入与 WAL 回放数据（按序列时间有序）。
type MemTable struct {
	series map[string][]tsmodel.Point
}

func New() *MemTable {
	return &MemTable{series: make(map[string][]tsmodel.Point)}
}

func (m *MemTable) Insert(key tsmodel.SeriesKey, p tsmodel.Point) error {
	k := key.String()
	existing := m.series[k]
	if err := utils.CheckSeriesValueType(existing, p); err != nil {
		return err
	}
	m.series[k] = utils.InsertSorted(existing, p)
	return nil
}

func (m *MemTable) Delete(key tsmodel.SeriesKey, ts int64) {
	k := key.String()
	m.series[k] = utils.DeleteFromSorted(m.series[k], ts)
}

func (m *MemTable) DeleteRange(key tsmodel.SeriesKey, start, end int64) {
	k := key.String()
	m.series[k] = utils.DeleteRangeFromSorted(m.series[k], start, end)
}

func (m *MemTable) Query(key tsmodel.SeriesKey, start, end int64) []tsmodel.Point {
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

func (m *MemTable) Snapshot() map[string][]tsmodel.Point {
	out := make(map[string][]tsmodel.Point, len(m.series))
	for k, pts := range m.series {
		out[k] = append([]tsmodel.Point(nil), pts...)
	}
	return out
}

func (m *MemTable) Reset() {
	m.series = make(map[string][]tsmodel.Point)
}

// SeriesTypes 返回内存表中各序列的数据类型（取首个点）。
func (m *MemTable) SeriesTypes() map[string]tsmodel.DataType {
	out := make(map[string]tsmodel.DataType, len(m.series))
	for k, pts := range m.series {
		if len(pts) == 0 {
			continue
		}
		out[k] = pts[0].Value.Type
	}
	return out
}
