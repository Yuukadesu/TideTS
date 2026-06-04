package storageengine

import "github.com/hanami/tidets/core/tsmodel"

// KnownSeriesTypes 返回当前存储中已知序列及其数据类型（MemTable + segment）。
func (e *Engine) KnownSeriesTypes() map[string]tsmodel.DataType {
	e.mu.RLock()
	defer e.mu.RUnlock()
	out := e.ws.SeriesTypes()
	for k, dt := range e.segments.SeriesTypes() {
		if _, ok := out[k]; !ok {
			out[k] = dt
		}
	}
	for _, job := range e.pending {
		for keyStr, pts := range job.snap {
			if len(pts) == 0 {
				continue
			}
			if _, ok := out[keyStr]; !ok {
				out[keyStr] = pts[0].Value.Type
			}
		}
	}
	return out
}
