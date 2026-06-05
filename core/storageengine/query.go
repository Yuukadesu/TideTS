package storageengine

import (
	"time"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/utils"
)

// Query 按时间范围查询；limit 为 0 时不限制条数。
func (e *Engine) Query(key SeriesKey, start, end int64, limit int) ([]Point, error) {
	begin := time.Now()
	if start > end {
		return nil, commons.ErrStorageInvalidRange
	}

	e.mu.RLock()
	mem := e.queryMemLocked(key, start, end)
	disk := e.segments.Query(key, start, end)
	e.mu.RUnlock()

	merged := utils.MergeSorted(mem, disk)
	merged = e.tombstones.Filter(key.String(), merged)
	out := utils.TruncatePoints(merged, limit)
	if e.hooks.OnRead != nil {
		e.hooks.OnRead("query", len(out), time.Since(begin))
	}
	return out, nil
}

func (e *Engine) queryMemLocked(key SeriesKey, start, end int64) []Point {
	mem := e.ws.Query(key, start, end)
	for _, snap := range e.pending {
		part := utils.QueryPointsInRange(snap.snap[key.String()], start, end)
		if len(part) > 0 {
			mem = utils.MergeSorted(mem, part)
		}
	}
	return mem
}
