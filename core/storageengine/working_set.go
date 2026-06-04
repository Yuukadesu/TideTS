package storageengine

import (
	"github.com/hanami/tidets/core/storageengine/memtable"
	"github.com/hanami/tidets/core/storageengine/utils"
	"github.com/hanami/tidets/core/tsmodel"
)

// workingSet 对齐 IoTDB：normal / delayed 双 MemTable + 每设备 stable time。
type workingSet struct {
	normal  *memtable.MemTable
	delayed *memtable.MemTable
	stable  map[string]int64
}

func newWorkingSet(stable map[string]int64) *workingSet {
	if stable == nil {
		stable = make(map[string]int64)
	}
	return &workingSet{
		normal:  memtable.New(),
		delayed: memtable.New(),
		stable:  stable,
	}
}

func (ws *workingSet) Insert(key SeriesKey, p Point) error {
	if p.Timestamp >= ws.stable[key.DevicePath] {
		return ws.normal.Insert(key, p)
	}
	return ws.delayed.Insert(key, p)
}

func (ws *workingSet) Delete(key SeriesKey, ts int64) {
	ws.normal.Delete(key, ts)
	ws.delayed.Delete(key, ts)
}

func (ws *workingSet) DeleteRange(key SeriesKey, start, end int64) {
	ws.normal.DeleteRange(key, start, end)
	ws.delayed.DeleteRange(key, start, end)
}

func (ws *workingSet) Query(key SeriesKey, start, end int64) []Point {
	n := ws.normal.Query(key, start, end)
	d := ws.delayed.Query(key, start, end)
	return utils.MergeSorted(n, d)
}

func (ws *workingSet) PointCount() int {
	return ws.normal.PointCount() + ws.delayed.PointCount()
}

func (ws *workingSet) IsEmpty() bool {
	return ws.normal.IsEmpty() && ws.delayed.IsEmpty()
}

func (ws *workingSet) Snapshot() map[string][]tsmodel.Point {
	a := ws.normal.Snapshot()
	b := ws.delayed.Snapshot()
	if len(b) == 0 {
		return a
	}
	out := make(map[string][]tsmodel.Point, len(a)+len(b))
	for k, pts := range a {
		out[k] = pts
	}
	for k, pts := range b {
		if existing, ok := out[k]; ok {
			out[k] = utils.MergeSorted(existing, pts)
			continue
		}
		out[k] = append([]tsmodel.Point(nil), pts...)
	}
	return out
}

func (ws *workingSet) Reset() {
	ws.normal.Reset()
	ws.delayed.Reset()
}

func (ws *workingSet) applyFlush(snap map[string][]tsmodel.Point) {
	for keyStr, pts := range snap {
		device := utils.DeviceFromSeriesKey(keyStr)
		for _, p := range pts {
			if p.Timestamp > ws.stable[device] {
				ws.stable[device] = p.Timestamp
			}
		}
	}
}

// SeriesTypes 合并 normal 与 delayed 中的序列类型。
func (ws *workingSet) SeriesTypes() map[string]tsmodel.DataType {
	out := ws.normal.SeriesTypes()
	for k, dt := range ws.delayed.SeriesTypes() {
		if _, ok := out[k]; !ok {
			out[k] = dt
		}
	}
	return out
}
